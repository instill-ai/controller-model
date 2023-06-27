package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/instill-ai/controller-model/internal/util"
	"github.com/instill-ai/controller-model/pkg/logger"

	controllerPB "github.com/instill-ai/protogen-go/model/controller/v1alpha"
	modelPB "github.com/instill-ai/protogen-go/model/model/v1alpha"
)

func (s *service) ProbeModels(ctx context.Context, cancel context.CancelFunc) error {
	defer cancel()

	logger, _ := logger.GetZapLogger(ctx)

	var wg sync.WaitGroup

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

	wg.Add(len(models))

	for _, model := range models {

		go func(model *modelPB.Model) {
			defer wg.Done()

			resourcePermalink := util.ConvertUIDToResourcePermalink(model.Uid, resourceType)

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
						return
					}
				}
			} else {
				if resp, err := s.modelPrivateClient.CheckModel(ctx, &modelPB.CheckModelRequest{
					ModelPermalink: fmt.Sprintf("%s/%s", resourceType, model.Uid),
				}); err == nil {
					if err = s.UpdateResourceState(ctx, &controllerPB.Resource{
						ResourcePermalink: resourcePermalink,
						State: &controllerPB.Resource_ModelState{
							ModelState: resp.State,
						},
					}); err != nil {
						logger.Error(err.Error())
						return
					}
				} else {
					logger.Error(err.Error())
					return
				}
			}

			logResp, _ := s.GetResourceState(ctx, resourcePermalink)
			logger.Info(fmt.Sprintf("[Controller] Got %v", logResp))
		}(model)

	}

	wg.Wait()

	return nil
}
