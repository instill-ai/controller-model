package handler

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/instill-ai/controller-model/pkg/logger"
	"github.com/instill-ai/controller-model/pkg/service"

	custom_otel "github.com/instill-ai/controller-model/pkg/logger/otel"
	commonPB "github.com/instill-ai/protogen-go/common"
	controllerPB "github.com/instill-ai/protogen-go/model/controller/v1alpha"
)

// PrivateHandler is the handler for private controller service
type PrivateHandler struct {
	controllerPB.UnimplementedControllerPrivateServiceServer
	service service.Service
}

// NewPrivateHandler returns a new private handler instance
func NewPrivateHandler(s service.Service) controllerPB.ControllerPrivateServiceServer {
	return &PrivateHandler{
		service: s,
	}
}

var tracer = otel.Tracer("controller.private-handler.tracer")

// Liveness checks the liveness of the server
func (h *PrivateHandler) Liveness(ctx context.Context, in *commonPB.LivenessRequest) (*commonPB.LivenessResponse, error) {
	return &commonPB.LivenessResponse{
		HealthCheckResponse: &commonPB.HealthCheckResponse{
			Status: commonPB.HealthCheckResponse_SERVING_STATUS_SERVING,
		},
	}, nil

}

// Readiness checks the readiness of the server
func (h *PrivateHandler) Readiness(ctx context.Context, in *commonPB.ReadinessRequest) (*commonPB.ReadinessResponse, error) {
	return &commonPB.ReadinessResponse{
		HealthCheckResponse: &commonPB.HealthCheckResponse{
			Status: commonPB.HealthCheckResponse_SERVING_STATUS_SERVING,
		},
	}, nil
}

func (h *PrivateHandler) GetResource(ctx context.Context, req *controllerPB.GetResourceRequest) (*controllerPB.GetResourceResponse, error) {

	ctx, span := tracer.Start(ctx, "GetResource",
		trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	logger, _ := logger.GetZapLogger(ctx)

	resource, err := h.service.GetResourceState(ctx, req.ResourcePermalink)
	if err != nil {
		return nil, err
	}

	logger.Info(string(custom_otel.NewLogMessage(
		span,
		false,
		"GetResource",
		"request",
		"GetResource done",
		false,
		custom_otel.SetEventResource(resource),
	)))

	return &controllerPB.GetResourceResponse{
		Resource: resource,
	}, nil
}

func (h *PrivateHandler) UpdateResource(ctx context.Context, req *controllerPB.UpdateResourceRequest) (*controllerPB.UpdateResourceResponse, error) {

	ctx, span := tracer.Start(ctx, "UpdateResource",
		trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	logger, _ := logger.GetZapLogger(ctx)

	if req.WorkflowId != nil {
		err := h.service.UpdateResourceWorkflowID(ctx, req.Resource.ResourcePermalink, *req.WorkflowId)

		if err != nil {
			return nil, err
		}
	}

	if err := h.service.UpdateResourceState(ctx, req.Resource); err != nil {
		return nil, err
	}

	logger.Info(string(custom_otel.NewLogMessage(
		span,
		false,
		"UpdateResource",
		"request",
		"UpdateResource done",
		false,
		custom_otel.SetEventResource(req.Resource),
	)))

	return &controllerPB.UpdateResourceResponse{
		Resource: req.Resource,
	}, nil
}

func (h *PrivateHandler) DeleteResource(ctx context.Context, req *controllerPB.DeleteResourceRequest) (*controllerPB.DeleteResourceResponse, error) {

	ctx, span := tracer.Start(ctx, "UpdateResource",
		trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	logger, _ := logger.GetZapLogger(ctx)

	if err := h.service.DeleteResourceState(ctx, req.ResourcePermalink); err != nil {
		return nil, err
	}

	if err := h.service.DeleteResourceWorkflowID(ctx, req.ResourcePermalink); err != nil {
		return nil, err
	}

	logger.Info(string(custom_otel.NewLogMessage(
		span,
		false,
		"UpdateResource",
		"request",
		"UpdateResource done",
		false,
		custom_otel.SetEventResource(req.ResourcePermalink),
	)))

	return &controllerPB.DeleteResourceResponse{}, nil
}
