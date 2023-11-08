package service

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/instill-ai/controller-model/internal/util"
	"github.com/instill-ai/controller-model/pkg/logger"

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

			workflowID, _ := s.GetResourceWorkflowID(ctx, resourcePermalink)

			var currentState modelPB.Model_State
			var desireState modelPB.Model_State

			curResp, _ := s.GetResourceState(ctx, resourcePermalink)
			currentState = curResp.GetModelState()

			if currentState == modelPB.Model_STATE_ERROR {
				logger.Warn(fmt.Sprintf("[Controller] %s: %v", model.Name, currentState))
				return
			}

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
			if resp, err := s.modelPrivateClient.CheckModelAdmin(ctx, &modelPB.CheckModelAdminRequest{
				ModelPermalink: modelPermalink,
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

			desireState = model.State

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
			logger.Info(fmt.Sprintf("[Controller] Moving %v from %v to %v", modelPermalink, currentState, desireState))
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
			logger.Info(fmt.Sprintf("[Controller] Moving %v from %v to %v", modelPermalink, currentState, desireState))
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
