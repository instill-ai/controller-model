package service_test

//go:generate mockgen -destination mock_model_client_test.go -package $GOPACKAGE github.com/instill-ai/protogen-go/model/model/v1alpha ModelPublicServiceClient

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	etcdv3 "go.etcd.io/etcd/client/v3"

	"github.com/instill-ai/controller-model/pkg/service"

	healthcheckPB "github.com/instill-ai/protogen-go/common/healthcheck/v1alpha"
	controllerPB "github.com/instill-ai/protogen-go/model/controller/v1alpha"
	modelPB "github.com/instill-ai/protogen-go/model/model/v1alpha"
)

const serviceResourceName = "resources/name/types/services"
const modelResourceName = "resources/name/types/models"

type Client struct {
	etcdv3.Cluster
	etcdv3.KV
	etcdv3.Lease
	etcdv3.Watcher
	etcdv3.Auth
	etcdv3.Maintenance

	// Username is a user name for authentication.
	Username string
	// Password is a password for authentication.
	Password string
	// contains filtered or unexported fields
}

func TestGetResourceState(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	t.Run("service", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockCluster := NewMockCluster(ctrl)
		mockKV := NewMockKV(ctrl)
		mockLease := NewMockLease(ctrl)
		mockWatcher := NewMockWatcher(ctrl)
		mockAuth := NewMockAuth(ctrl)
		mockMaintenance := NewMockMaintenance(ctrl)

		mockEtcdClient := etcdv3.Client{
			Cluster:     mockCluster,
			KV:          mockKV,
			Lease:       mockLease,
			Watcher:     mockWatcher,
			Auth:        mockAuth,
			Maintenance: mockMaintenance,
		}

		var resp *etcdv3.GetResponse

		mockKV.
			EXPECT().
			Get(ctx, serviceResourceName).
			Return(resp, nil).
			Times(1)

		s := service.NewService(nil, nil, nil, nil, mockEtcdClient)

		resource, err := s.GetResourceState(ctx, serviceResourceName)

		assert.Equal(t, healthcheckPB.HealthCheckResponse_SERVING_STATUS_UNSPECIFIED, resource.GetBackendState())

		assert.NoError(t, err)
	})
	t.Run("model", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockCluster := NewMockCluster(ctrl)
		mockKV := NewMockKV(ctrl)
		mockLease := NewMockLease(ctrl)
		mockWatcher := NewMockWatcher(ctrl)
		mockAuth := NewMockAuth(ctrl)
		mockMaintenance := NewMockMaintenance(ctrl)

		mockEtcdClient := etcdv3.Client{
			Cluster:     mockCluster,
			KV:          mockKV,
			Lease:       mockLease,
			Watcher:     mockWatcher,
			Auth:        mockAuth,
			Maintenance: mockMaintenance,
		}

		var resp *etcdv3.GetResponse

		mockKV.
			EXPECT().
			Get(ctx, modelResourceName).
			Return(resp, nil).
			Times(1)

		s := service.NewService(nil, nil, nil, nil, mockEtcdClient)

		resource, err := s.GetResourceState(ctx, modelResourceName)

		assert.Equal(t, modelPB.Model_STATE_UNSPECIFIED, resource.GetModelState())

		assert.NoError(t, err)
	})

}

func TestUpdateResourceState(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	t.Run("service", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockCluster := NewMockCluster(ctrl)
		mockKV := NewMockKV(ctrl)
		mockLease := NewMockLease(ctrl)
		mockWatcher := NewMockWatcher(ctrl)
		mockAuth := NewMockAuth(ctrl)
		mockMaintenance := NewMockMaintenance(ctrl)

		mockEtcdClient := etcdv3.Client{
			Cluster:     mockCluster,
			KV:          mockKV,
			Lease:       mockLease,
			Watcher:     mockWatcher,
			Auth:        mockAuth,
			Maintenance: mockMaintenance,
		}

		resource := controllerPB.Resource{
			ResourcePermalink: serviceResourceName,
			State: &controllerPB.Resource_BackendState{
				BackendState: healthcheckPB.HealthCheckResponse_SERVING_STATUS_UNSPECIFIED,
			},
		}

		mockKV.
			EXPECT().
			Put(ctx, serviceResourceName, string("0")).
			Return(&etcdv3.PutResponse{}, nil).
			Times(1)

		s := service.NewService(nil, nil, nil, nil, mockEtcdClient)

		err := s.UpdateResourceState(ctx, &resource)

		assert.NoError(t, err)
	})

	t.Run("model", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockCluster := NewMockCluster(ctrl)
		mockKV := NewMockKV(ctrl)
		mockLease := NewMockLease(ctrl)
		mockWatcher := NewMockWatcher(ctrl)
		mockAuth := NewMockAuth(ctrl)
		mockMaintenance := NewMockMaintenance(ctrl)

		mockEtcdClient := etcdv3.Client{
			Cluster:     mockCluster,
			KV:          mockKV,
			Lease:       mockLease,
			Watcher:     mockWatcher,
			Auth:        mockAuth,
			Maintenance: mockMaintenance,
		}

		resource := controllerPB.Resource{
			ResourcePermalink: modelResourceName,
			State: &controllerPB.Resource_ModelState{
				ModelState: modelPB.Model_STATE_UNSPECIFIED,
			},
		}

		mockKV.
			EXPECT().
			Put(ctx, modelResourceName, string("0")).
			Return(&etcdv3.PutResponse{}, nil).
			Times(1)

		s := service.NewService(nil, nil, nil, nil, mockEtcdClient)

		err := s.UpdateResourceState(ctx, &resource)

		assert.NoError(t, err)
	})

}

func TestDeleteResourceState(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	t.Run("service", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockCluster := NewMockCluster(ctrl)
		mockKV := NewMockKV(ctrl)
		mockLease := NewMockLease(ctrl)
		mockWatcher := NewMockWatcher(ctrl)
		mockAuth := NewMockAuth(ctrl)
		mockMaintenance := NewMockMaintenance(ctrl)

		mockEtcdClient := etcdv3.Client{
			Cluster:     mockCluster,
			KV:          mockKV,
			Lease:       mockLease,
			Watcher:     mockWatcher,
			Auth:        mockAuth,
			Maintenance: mockMaintenance,
		}

		var resp *etcdv3.DeleteResponse

		mockKV.
			EXPECT().
			Delete(ctx, serviceResourceName).
			Return(resp, nil).
			Times(1)

		s := service.NewService(nil, nil, nil, nil, mockEtcdClient)

		err := s.DeleteResourceState(ctx, serviceResourceName)

		assert.NoError(t, err)
	})
	t.Run("model", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockCluster := NewMockCluster(ctrl)
		mockKV := NewMockKV(ctrl)
		mockLease := NewMockLease(ctrl)
		mockWatcher := NewMockWatcher(ctrl)
		mockAuth := NewMockAuth(ctrl)
		mockMaintenance := NewMockMaintenance(ctrl)

		mockEtcdClient := etcdv3.Client{
			Cluster:     mockCluster,
			KV:          mockKV,
			Lease:       mockLease,
			Watcher:     mockWatcher,
			Auth:        mockAuth,
			Maintenance: mockMaintenance,
		}

		var resp *etcdv3.DeleteResponse

		mockKV.
			EXPECT().
			Delete(ctx, modelResourceName).
			Return(resp, nil).
			Times(1)

		s := service.NewService(nil, nil, nil, nil, mockEtcdClient)

		err := s.DeleteResourceState(ctx, modelResourceName)

		assert.NoError(t, err)
	})

}
