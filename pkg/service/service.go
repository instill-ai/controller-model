package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"

	etcdv3 "go.etcd.io/etcd/client/v3"

	"github.com/instill-ai/controller-model/config"
	"github.com/instill-ai/controller-model/internal/util"
	"github.com/instill-ai/controller-model/pkg/logger"

	inferenceserver "github.com/instill-ai/controller-model/internal/triton"
	mgmtPB "github.com/instill-ai/protogen-go/base/mgmt/v1alpha"
	healthcheckPB "github.com/instill-ai/protogen-go/common/healthcheck/v1alpha"
	controllerPB "github.com/instill-ai/protogen-go/model/controller/v1alpha"
	modelPB "github.com/instill-ai/protogen-go/model/model/v1alpha"
)

// Service is the interface for the controller service
type Service interface {
	GetResourceState(ctx context.Context, resourcePermalink string) (*controllerPB.Resource, error)
	UpdateResourceState(ctx context.Context, resource *controllerPB.Resource) error
	DeleteResourceState(ctx context.Context, resourcePermalink string) error
	GetResourceWorkflowID(ctx context.Context, resourcePermalink string) (*string, error)
	UpdateResourceWorkflowID(ctx context.Context, resourcePermalink string, workflowID string) error
	DeleteResourceWorkflowID(ctx context.Context, resourcePermalink string) error
	ProbeBackend(ctx context.Context, cancel context.CancelFunc) error
	ProbeModels(ctx context.Context, cancel context.CancelFunc) error
}

type service struct {
	modelPublicClient  modelPB.ModelPublicServiceClient
	modelPrivateClient modelPB.ModelPrivateServiceClient
	mgmtPublicClient   mgmtPB.MgmtPublicServiceClient
	tritonClient       inferenceserver.GRPCInferenceServiceClient
	etcdClient         etcdv3.Client
}

// NewService returns a new service instance
func NewService(
	mp modelPB.ModelPublicServiceClient,
	m modelPB.ModelPrivateServiceClient,
	mg mgmtPB.MgmtPublicServiceClient,
	t inferenceserver.GRPCInferenceServiceClient,
	e etcdv3.Client) Service {
	return &service{
		modelPublicClient:  mp,
		modelPrivateClient: m,
		mgmtPublicClient:   mg,
		tritonClient:       t,
		etcdClient:         e,
	}
}

func (s *service) GetResourceState(ctx context.Context, resourcePermalink string) (*controllerPB.Resource, error) {
	resp, err := s.etcdClient.Get(ctx, resourcePermalink)

	if err != nil {
		return nil, err
	}

	kvs := resp.Kvs

	if len(kvs) == 0 {
		return nil, fmt.Errorf(fmt.Sprintf("resource %v not found in etcd storage", resourcePermalink))
	}

	resourceType := strings.SplitN(resourcePermalink, "/", 4)[3]

	stateEnumValue, _ := strconv.ParseInt(string(kvs[0].Value[:]), 10, 32)

	switch resourceType {
	case util.RESOURCE_TYPE_MODEL:
		return &controllerPB.Resource{
			ResourcePermalink: resourcePermalink,
			State: &controllerPB.Resource_ModelState{
				ModelState: modelPB.Model_State(stateEnumValue),
			},
			Progress: nil,
		}, nil
	case util.RESOURCE_TYPE_SERVICE:
		return &controllerPB.Resource{
			ResourcePermalink: resourcePermalink,
			State: &controllerPB.Resource_BackendState{
				BackendState: healthcheckPB.HealthCheckResponse_ServingStatus(stateEnumValue),
			},
		}, nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("get resource type %s not implemented", resourceType))
	}
}

func (s *service) UpdateResourceState(ctx context.Context, resource *controllerPB.Resource) error {
	resourceType := strings.SplitN(resource.ResourcePermalink, "/", 4)[3]

	state := 0

	switch resourceType {
	case util.RESOURCE_TYPE_MODEL:
		state = int(resource.GetModelState())
	case util.RESOURCE_TYPE_SERVICE:
		state = int(resource.GetBackendState())
	default:
		return fmt.Errorf(fmt.Sprintf("update resource type %s not implemented", resourceType))
	}

	if _, err := s.etcdClient.Put(ctx, resource.ResourcePermalink, fmt.Sprint(state)); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteResourceState(ctx context.Context, resourcePermalink string) error {
	_, err := s.etcdClient.Delete(ctx, resourcePermalink)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetResourceWorkflowID(ctx context.Context, resourcePermalink string) (*string, error) {
	resourceWorkflowID := util.ConvertResourcePermalinkToWorkflowName(resourcePermalink)

	resp, err := s.etcdClient.Get(ctx, resourceWorkflowID)

	if err != nil {
		return nil, err
	}

	kvs := resp.Kvs

	if len(kvs) == 0 {
		return nil, fmt.Errorf("workflowID not found in etcd storage")
	}

	workflowID := string(kvs[0].Value[:])

	return &workflowID, nil
}

func (s *service) UpdateResourceWorkflowID(ctx context.Context, resourcePermalink string, workflowID string) error {
	resourceWorkflowID := util.ConvertResourcePermalinkToWorkflowName(resourcePermalink)

	_, err := s.etcdClient.Put(ctx, resourceWorkflowID, workflowID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteResourceWorkflowID(ctx context.Context, resourcePermalink string) error {
	resourceWorkflowID := util.ConvertResourcePermalinkToWorkflowName(resourcePermalink)

	_, err := s.etcdClient.Delete(ctx, resourceWorkflowID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) ProbeBackend(ctx context.Context, cancel context.CancelFunc) error {
	defer cancel()

	logger, _ := logger.GetZapLogger(ctx)

	var wg sync.WaitGroup

	healthcheck := healthcheckPB.HealthCheckResponse{
		Status: healthcheckPB.HealthCheckResponse_SERVING_STATUS_UNSPECIFIED,
	}

	var backendServices = [...]string{
		config.Config.ModelBackend.Host,
		config.Config.MgmtBackend.Host,
		config.Config.TritonServer.Host,
	}

	wg.Add(len(backendServices))

	for _, hostname := range backendServices {
		go func(hostname string) {
			defer wg.Done()

			switch hostname {
			case config.Config.ModelBackend.Host:
				resp, err := s.modelPublicClient.Liveness(ctx, &modelPB.LivenessRequest{})

				if err != nil {
					healthcheck = healthcheckPB.HealthCheckResponse{
						Status: healthcheckPB.HealthCheckResponse_SERVING_STATUS_NOT_SERVING,
					}
				} else {
					healthcheck = *resp.GetHealthCheckResponse()
				}
			case config.Config.MgmtBackend.Host:
				resp, err := s.mgmtPublicClient.Liveness(ctx, &mgmtPB.LivenessRequest{})

				if err != nil {
					healthcheck = healthcheckPB.HealthCheckResponse{
						Status: healthcheckPB.HealthCheckResponse_SERVING_STATUS_NOT_SERVING,
					}
				} else {
					healthcheck = *resp.GetHealthCheckResponse()
				}
			case config.Config.TritonServer.Host:
				resp, err := s.tritonClient.ServerLive(ctx, &inferenceserver.ServerLiveRequest{})

				if err != nil {
					healthcheck = healthcheckPB.HealthCheckResponse{
						Status: healthcheckPB.HealthCheckResponse_SERVING_STATUS_NOT_SERVING,
					}
				} else {
					if resp.GetLive() {
						healthcheck = healthcheckPB.HealthCheckResponse{
							Status: healthcheckPB.HealthCheckResponse_SERVING_STATUS_SERVING,
						}
					} else {
						healthcheck = healthcheckPB.HealthCheckResponse{
							Status: healthcheckPB.HealthCheckResponse_SERVING_STATUS_NOT_SERVING,
						}
					}
				}
			}

			err := s.UpdateResourceState(ctx, &controllerPB.Resource{
				ResourcePermalink: util.ConvertServiceToResourceName(hostname),
				State: &controllerPB.Resource_BackendState{
					BackendState: healthcheck.Status,
				},
			})

			if err != nil {
				logger.Error(err.Error())
				return
			}

			resp, _ := s.GetResourceState(ctx, util.ConvertServiceToResourceName(hostname))

			logger.Info(fmt.Sprintf("[Controller] Got %v", resp))
		}(hostname)
	}

	wg.Wait()

	return nil
}

func (s *service) getOperationInfo(workflowID string, resourceType string) (*longrunningpb.Operation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Server.Timeout*time.Second)
	defer cancel()

	var operation *longrunningpb.Operation

	switch resourceType {
	case util.RESOURCE_TYPE_MODEL:
		op, err := s.modelPublicClient.GetModelOperation(ctx, &modelPB.GetModelOperationRequest{
			Name: fmt.Sprintf("operations/%s", workflowID),
		})
		if err != nil {
			return nil, err
		}
		operation = op.Operation
	}

	return operation, nil
}
