package service

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/instill-ai/controller-model/config"
	"github.com/instill-ai/controller-model/internal/util"
	"github.com/instill-ai/controller-model/pkg/logger"

	healthcheckPB "github.com/instill-ai/protogen-go/common/healthcheck/v1alpha"
	controllerPB "github.com/instill-ai/protogen-go/model/controller/v1alpha"
	modelPB "github.com/instill-ai/protogen-go/model/model/v1alpha"
)

type ReconcileModel struct {
	ModelPermalink    string
	ResourcePermalink string
	CurrentState      modelPB.Model_State
	DesireState       modelPB.Model_State
}

func (s *service) ProbeModels(ctx context.Context, cancel context.CancelFunc) error {
	defer cancel()

	logger, _ := logger.GetZapLogger(ctx)

	var wg sync.WaitGroup

	if modelBackendResource, err := s.GetResourceState(ctx, util.ConvertServiceToResourceName(config.Config.ModelBackend.Host)); err != nil {
		return err
	} else if modelBackendResource.GetBackendState() != healthcheckPB.HealthCheckResponse_SERVING_STATUS_SERVING {
		return fmt.Errorf("[Controller] model-backend is not serving")
	}
	if tritonServerResource, err := s.GetResourceState(ctx, util.ConvertServiceToResourceName(config.Config.TritonServer.Host)); err != nil {
		return err
	} else if tritonServerResource.GetBackendState() != healthcheckPB.HealthCheckResponse_SERVING_STATUS_SERVING {
		return fmt.Errorf("[Controller] triton-server is not serving")
	}

	resp, err := s.modelPrivateClient.ListModelsAdmin(ctx, &modelPB.ListModelsAdminRequest{})

	if err != nil {
		return err
	}

	models := resp.Models
	nextPageToken := &resp.NextPageToken
	totalSize := resp.TotalSize

	for totalSize > util.DefaultPageSize {
		resp, err := s.modelPrivateClient.ListModelsAdmin(ctx, &modelPB.ListModelsAdminRequest{
			PageToken: nextPageToken,
		})

		if err != nil {
			return err
		}

		nextPageToken = &resp.NextPageToken
		totalSize -= util.DefaultPageSize
		models = append(models, resp.Models...)
	}

	resourceType := "models"

	reconcileModelChannel := make(chan ReconcileModel, len(models))

	for _, model := range models {
		wg.Add(1)
		go func(model *modelPB.Model) {
			defer wg.Done()

			resourcePermalink := util.ConvertUIDToResourcePermalink(model.Uid, resourceType)
			modelPermalink := fmt.Sprintf("%s/%s", "models", model.Uid)

			// model in transition state
			workflowID, _ := s.GetResourceWorkflowID(ctx, resourcePermalink)

			if workflowID != nil {
				opInfo, err := s.getOperationInfo(*workflowID, util.RESOURCE_TYPE_MODEL)
				if err != nil {
					logger.Error(err.Error())
					return
				}
				if opInfo.Done {
					if err := s.DeleteResourceWorkflowID(ctx, resourcePermalink); err != nil {
						logger.Error(err.Error())
					}
					if opInfo.GetError() != nil {
						if err = s.UpdateResourceState(ctx, &controllerPB.Resource{
							ResourcePermalink: resourcePermalink,
							State: &controllerPB.Resource_ModelState{
								ModelState: modelPB.Model_STATE_ERROR,
							},
						}); err != nil {
							logger.Error(err.Error())
						}
					}
				}
				return
			}

			// model in end state
			lastResp, _ := s.GetResourceState(ctx, resourcePermalink)
			curResp, err := s.modelPrivateClient.CheckModelAdmin(ctx, &modelPB.CheckModelAdminRequest{
				ModelPermalink: modelPermalink,
			})
			if err != nil {
				logger.Error(err.Error())
				return
			}

			lastProbeState := lastResp.GetModelState()
			currentState := curResp.GetState()
			desireState := model.State

			if currentState != modelPB.Model_STATE_ERROR && lastProbeState != modelPB.Model_STATE_ERROR {
				if err = s.UpdateResourceRetryCount(ctx, resourcePermalink, 0); err != nil {
					return
				}
			}

			if err = s.UpdateResourceState(ctx, &controllerPB.Resource{
				ResourcePermalink: resourcePermalink,
				State: &controllerPB.Resource_ModelState{
					ModelState: currentState,
				},
			}); err != nil {
				logger.Error(err.Error())
				return
			}

			if err := s.checkRetry(ctx, resourcePermalink); err != nil {
				if e := s.UpdateResourceState(ctx, &controllerPB.Resource{
					ResourcePermalink: resourcePermalink,
					State: &controllerPB.Resource_ModelState{
						ModelState: modelPB.Model_STATE_ERROR,
					},
				}); e != nil {
					return
				}
				currentState = lastProbeState
				logger.Warn(err.Error())
			} else {
				if lastProbeState == modelPB.Model_STATE_ERROR {
					logger.Warn(fmt.Sprintf("[Controller] %s last op errored, trigger retry", model.Name))
				}
			}

			rModel := ReconcileModel{
				ModelPermalink:    modelPermalink,
				ResourcePermalink: resourcePermalink,
				CurrentState:      currentState,
				DesireState:       desireState,
			}

			reconcileModelChannel <- rModel

		}(model)

	}

	wg.Wait()

	for i := 0; i < len(models); i++ {
		select {
		case rModel := <-reconcileModelChannel:
			if err = s.moveCurrentStateToDesireState(
				ctx,
				rModel.ModelPermalink,
				rModel.ResourcePermalink,
				rModel.CurrentState,
				rModel.DesireState,
			); err != nil {
				logger.Error(err.Error())
			}
		default:
			logger.Info(fmt.Sprintf("[Controller] %v not in a valid state for operation", models[i].Name))
		}
	}

	close(reconcileModelChannel)

	return nil
}

func (s *service) moveCurrentStateToDesireState(ctx context.Context, modelPermalink string, resourcePermalink string, currentState modelPB.Model_State, desireState modelPB.Model_State) (err error) {
	logger, _ := logger.GetZapLogger(ctx)

	switch desireState {
	case modelPB.Model_STATE_ONLINE:
		switch currentState {
		case modelPB.Model_STATE_OFFLINE:
			logger.Info(fmt.Sprintf("[Controller] moving %v from %v to %v", modelPermalink, currentState, desireState))
			resp, err := s.modelPrivateClient.DeployModelAdmin(ctx, &modelPB.DeployModelAdminRequest{
				ModelPermalink: modelPermalink,
			})
			if err != nil {
				if err = s.UpdateResourceState(ctx, &controllerPB.Resource{
					ResourcePermalink: resourcePermalink,
					State: &controllerPB.Resource_ModelState{
						ModelState: modelPB.Model_STATE_ERROR,
					},
				}); err != nil {
					return err
				}
				return err
			}
			if err = s.UpdateResourceWorkflowID(ctx, resourcePermalink, strings.Split(resp.GetOperation().GetName(), "/")[1]); err != nil {
				return err
			}
			if err = s.UpdateResourceState(ctx, &controllerPB.Resource{
				ResourcePermalink: resourcePermalink,
				State: &controllerPB.Resource_ModelState{
					ModelState: modelPB.Model_STATE_UNSPECIFIED,
				},
			}); err != nil {
				return err
			}
		}
	case modelPB.Model_STATE_OFFLINE:
		switch currentState {
		case modelPB.Model_STATE_ONLINE:
			logger.Info(fmt.Sprintf("[Controller] moving %v from %v to %v", modelPermalink, currentState, desireState))
			resp, err := s.modelPrivateClient.UndeployModelAdmin(ctx, &modelPB.UndeployModelAdminRequest{
				ModelPermalink: modelPermalink,
			})
			if err != nil {
				if err = s.UpdateResourceState(ctx, &controllerPB.Resource{
					ResourcePermalink: resourcePermalink,
					State: &controllerPB.Resource_ModelState{
						ModelState: modelPB.Model_STATE_ERROR,
					},
				}); err != nil {
					return err
				}
				return err
			}
			if err = s.UpdateResourceWorkflowID(ctx, resourcePermalink, strings.Split(resp.GetOperation().GetName(), "/")[1]); err != nil {
				return err
			}
			if err = s.UpdateResourceState(ctx, &controllerPB.Resource{
				ResourcePermalink: resourcePermalink,
				State: &controllerPB.Resource_ModelState{
					ModelState: modelPB.Model_STATE_UNSPECIFIED,
				},
			}); err != nil {
				return err
			}
		}
	}
	return err
}

func (s *service) checkRetry(ctx context.Context, resourcePermalink string) error {
	// check retry count
	if retryCount, err := s.GetResourceRetryCount(ctx, resourcePermalink); err != nil {
		return err
	} else if *retryCount >= util.DefaultRetryCount {
		return fmt.Errorf(fmt.Sprintf("[Controller] retry limit reached for %s", resourcePermalink))
	} else {
		if err = s.UpdateResourceRetryCount(ctx, resourcePermalink, *retryCount+int64(1)); err != nil {
			return err
		}
	}
	return nil
}
