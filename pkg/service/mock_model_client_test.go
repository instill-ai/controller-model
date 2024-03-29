// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/instill-ai/protogen-go/model/model/v1alpha (interfaces: ModelPublicServiceClient)

// Package service_test is a generated GoMock package.
package service_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	modelv1alpha "github.com/instill-ai/protogen-go/model/model/v1alpha"
	grpc "google.golang.org/grpc"
)

// MockModelPublicServiceClient is a mock of ModelPublicServiceClient interface.
type MockModelPublicServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockModelPublicServiceClientMockRecorder
}

// MockModelPublicServiceClientMockRecorder is the mock recorder for MockModelPublicServiceClient.
type MockModelPublicServiceClientMockRecorder struct {
	mock *MockModelPublicServiceClient
}

// NewMockModelPublicServiceClient creates a new mock instance.
func NewMockModelPublicServiceClient(ctrl *gomock.Controller) *MockModelPublicServiceClient {
	mock := &MockModelPublicServiceClient{ctrl: ctrl}
	mock.recorder = &MockModelPublicServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelPublicServiceClient) EXPECT() *MockModelPublicServiceClientMockRecorder {
	return m.recorder
}

// CreateOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) CreateOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.CreateOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.CreateOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.CreateOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganizationModel indicates an expected call of CreateOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) CreateOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).CreateOrganizationModel), varargs...)
}

// CreateOrganizationModelBinaryFileUpload mocks base method.
func (m *MockModelPublicServiceClient) CreateOrganizationModelBinaryFileUpload(arg0 context.Context, arg1 ...grpc.CallOption) (modelv1alpha.ModelPublicService_CreateOrganizationModelBinaryFileUploadClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrganizationModelBinaryFileUpload", varargs...)
	ret0, _ := ret[0].(modelv1alpha.ModelPublicService_CreateOrganizationModelBinaryFileUploadClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganizationModelBinaryFileUpload indicates an expected call of CreateOrganizationModelBinaryFileUpload.
func (mr *MockModelPublicServiceClientMockRecorder) CreateOrganizationModelBinaryFileUpload(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganizationModelBinaryFileUpload", reflect.TypeOf((*MockModelPublicServiceClient)(nil).CreateOrganizationModelBinaryFileUpload), varargs...)
}

// CreateUserModel mocks base method.
func (m *MockModelPublicServiceClient) CreateUserModel(arg0 context.Context, arg1 *modelv1alpha.CreateUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.CreateUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.CreateUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserModel indicates an expected call of CreateUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) CreateUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).CreateUserModel), varargs...)
}

// CreateUserModelBinaryFileUpload mocks base method.
func (m *MockModelPublicServiceClient) CreateUserModelBinaryFileUpload(arg0 context.Context, arg1 ...grpc.CallOption) (modelv1alpha.ModelPublicService_CreateUserModelBinaryFileUploadClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUserModelBinaryFileUpload", varargs...)
	ret0, _ := ret[0].(modelv1alpha.ModelPublicService_CreateUserModelBinaryFileUploadClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserModelBinaryFileUpload indicates an expected call of CreateUserModelBinaryFileUpload.
func (mr *MockModelPublicServiceClientMockRecorder) CreateUserModelBinaryFileUpload(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserModelBinaryFileUpload", reflect.TypeOf((*MockModelPublicServiceClient)(nil).CreateUserModelBinaryFileUpload), varargs...)
}

// DeleteOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) DeleteOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.DeleteOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.DeleteOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.DeleteOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrganizationModel indicates an expected call of DeleteOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) DeleteOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).DeleteOrganizationModel), varargs...)
}

// DeleteUserModel mocks base method.
func (m *MockModelPublicServiceClient) DeleteUserModel(arg0 context.Context, arg1 *modelv1alpha.DeleteUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.DeleteUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.DeleteUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserModel indicates an expected call of DeleteUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) DeleteUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).DeleteUserModel), varargs...)
}

// DeployOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) DeployOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.DeployOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.DeployOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeployOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.DeployOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployOrganizationModel indicates an expected call of DeployOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) DeployOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).DeployOrganizationModel), varargs...)
}

// DeployUserModel mocks base method.
func (m *MockModelPublicServiceClient) DeployUserModel(arg0 context.Context, arg1 *modelv1alpha.DeployUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.DeployUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeployUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.DeployUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployUserModel indicates an expected call of DeployUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) DeployUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).DeployUserModel), varargs...)
}

// GetModelDefinition mocks base method.
func (m *MockModelPublicServiceClient) GetModelDefinition(arg0 context.Context, arg1 *modelv1alpha.GetModelDefinitionRequest, arg2 ...grpc.CallOption) (*modelv1alpha.GetModelDefinitionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetModelDefinition", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.GetModelDefinitionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelDefinition indicates an expected call of GetModelDefinition.
func (mr *MockModelPublicServiceClientMockRecorder) GetModelDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelDefinition", reflect.TypeOf((*MockModelPublicServiceClient)(nil).GetModelDefinition), varargs...)
}

// GetModelOperation mocks base method.
func (m *MockModelPublicServiceClient) GetModelOperation(arg0 context.Context, arg1 *modelv1alpha.GetModelOperationRequest, arg2 ...grpc.CallOption) (*modelv1alpha.GetModelOperationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetModelOperation", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.GetModelOperationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelOperation indicates an expected call of GetModelOperation.
func (mr *MockModelPublicServiceClientMockRecorder) GetModelOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelOperation", reflect.TypeOf((*MockModelPublicServiceClient)(nil).GetModelOperation), varargs...)
}

// GetOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) GetOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.GetOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.GetOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.GetOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationModel indicates an expected call of GetOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) GetOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).GetOrganizationModel), varargs...)
}

// GetOrganizationModelCard mocks base method.
func (m *MockModelPublicServiceClient) GetOrganizationModelCard(arg0 context.Context, arg1 *modelv1alpha.GetOrganizationModelCardRequest, arg2 ...grpc.CallOption) (*modelv1alpha.GetOrganizationModelCardResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrganizationModelCard", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.GetOrganizationModelCardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationModelCard indicates an expected call of GetOrganizationModelCard.
func (mr *MockModelPublicServiceClientMockRecorder) GetOrganizationModelCard(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationModelCard", reflect.TypeOf((*MockModelPublicServiceClient)(nil).GetOrganizationModelCard), varargs...)
}

// GetUserModel mocks base method.
func (m *MockModelPublicServiceClient) GetUserModel(arg0 context.Context, arg1 *modelv1alpha.GetUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.GetUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.GetUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserModel indicates an expected call of GetUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) GetUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).GetUserModel), varargs...)
}

// GetUserModelCard mocks base method.
func (m *MockModelPublicServiceClient) GetUserModelCard(arg0 context.Context, arg1 *modelv1alpha.GetUserModelCardRequest, arg2 ...grpc.CallOption) (*modelv1alpha.GetUserModelCardResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserModelCard", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.GetUserModelCardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserModelCard indicates an expected call of GetUserModelCard.
func (mr *MockModelPublicServiceClientMockRecorder) GetUserModelCard(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserModelCard", reflect.TypeOf((*MockModelPublicServiceClient)(nil).GetUserModelCard), varargs...)
}

// ListModelDefinitions mocks base method.
func (m *MockModelPublicServiceClient) ListModelDefinitions(arg0 context.Context, arg1 *modelv1alpha.ListModelDefinitionsRequest, arg2 ...grpc.CallOption) (*modelv1alpha.ListModelDefinitionsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListModelDefinitions", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.ListModelDefinitionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModelDefinitions indicates an expected call of ListModelDefinitions.
func (mr *MockModelPublicServiceClientMockRecorder) ListModelDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelDefinitions", reflect.TypeOf((*MockModelPublicServiceClient)(nil).ListModelDefinitions), varargs...)
}

// ListModels mocks base method.
func (m *MockModelPublicServiceClient) ListModels(arg0 context.Context, arg1 *modelv1alpha.ListModelsRequest, arg2 ...grpc.CallOption) (*modelv1alpha.ListModelsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListModels", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.ListModelsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModels indicates an expected call of ListModels.
func (mr *MockModelPublicServiceClientMockRecorder) ListModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModels", reflect.TypeOf((*MockModelPublicServiceClient)(nil).ListModels), varargs...)
}

// ListOrganizationModels mocks base method.
func (m *MockModelPublicServiceClient) ListOrganizationModels(arg0 context.Context, arg1 *modelv1alpha.ListOrganizationModelsRequest, arg2 ...grpc.CallOption) (*modelv1alpha.ListOrganizationModelsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrganizationModels", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.ListOrganizationModelsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizationModels indicates an expected call of ListOrganizationModels.
func (mr *MockModelPublicServiceClientMockRecorder) ListOrganizationModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizationModels", reflect.TypeOf((*MockModelPublicServiceClient)(nil).ListOrganizationModels), varargs...)
}

// ListUserModels mocks base method.
func (m *MockModelPublicServiceClient) ListUserModels(arg0 context.Context, arg1 *modelv1alpha.ListUserModelsRequest, arg2 ...grpc.CallOption) (*modelv1alpha.ListUserModelsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUserModels", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.ListUserModelsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserModels indicates an expected call of ListUserModels.
func (mr *MockModelPublicServiceClientMockRecorder) ListUserModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserModels", reflect.TypeOf((*MockModelPublicServiceClient)(nil).ListUserModels), varargs...)
}

// Liveness mocks base method.
func (m *MockModelPublicServiceClient) Liveness(arg0 context.Context, arg1 *modelv1alpha.LivenessRequest, arg2 ...grpc.CallOption) (*modelv1alpha.LivenessResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Liveness", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.LivenessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Liveness indicates an expected call of Liveness.
func (mr *MockModelPublicServiceClientMockRecorder) Liveness(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Liveness", reflect.TypeOf((*MockModelPublicServiceClient)(nil).Liveness), varargs...)
}

// LookUpModel mocks base method.
func (m *MockModelPublicServiceClient) LookUpModel(arg0 context.Context, arg1 *modelv1alpha.LookUpModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.LookUpModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LookUpModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.LookUpModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookUpModel indicates an expected call of LookUpModel.
func (mr *MockModelPublicServiceClientMockRecorder) LookUpModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookUpModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).LookUpModel), varargs...)
}

// PublishOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) PublishOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.PublishOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.PublishOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PublishOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.PublishOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishOrganizationModel indicates an expected call of PublishOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) PublishOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).PublishOrganizationModel), varargs...)
}

// PublishUserModel mocks base method.
func (m *MockModelPublicServiceClient) PublishUserModel(arg0 context.Context, arg1 *modelv1alpha.PublishUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.PublishUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PublishUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.PublishUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishUserModel indicates an expected call of PublishUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) PublishUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).PublishUserModel), varargs...)
}

// Readiness mocks base method.
func (m *MockModelPublicServiceClient) Readiness(arg0 context.Context, arg1 *modelv1alpha.ReadinessRequest, arg2 ...grpc.CallOption) (*modelv1alpha.ReadinessResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Readiness", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.ReadinessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Readiness indicates an expected call of Readiness.
func (mr *MockModelPublicServiceClientMockRecorder) Readiness(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Readiness", reflect.TypeOf((*MockModelPublicServiceClient)(nil).Readiness), varargs...)
}

// RenameOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) RenameOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.RenameOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.RenameOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RenameOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.RenameOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameOrganizationModel indicates an expected call of RenameOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) RenameOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).RenameOrganizationModel), varargs...)
}

// RenameUserModel mocks base method.
func (m *MockModelPublicServiceClient) RenameUserModel(arg0 context.Context, arg1 *modelv1alpha.RenameUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.RenameUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RenameUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.RenameUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameUserModel indicates an expected call of RenameUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) RenameUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).RenameUserModel), varargs...)
}

// TriggerOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) TriggerOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.TriggerOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.TriggerOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TriggerOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.TriggerOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerOrganizationModel indicates an expected call of TriggerOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) TriggerOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).TriggerOrganizationModel), varargs...)
}

// TriggerOrganizationModelBinaryFileUpload mocks base method.
func (m *MockModelPublicServiceClient) TriggerOrganizationModelBinaryFileUpload(arg0 context.Context, arg1 ...grpc.CallOption) (modelv1alpha.ModelPublicService_TriggerOrganizationModelBinaryFileUploadClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TriggerOrganizationModelBinaryFileUpload", varargs...)
	ret0, _ := ret[0].(modelv1alpha.ModelPublicService_TriggerOrganizationModelBinaryFileUploadClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerOrganizationModelBinaryFileUpload indicates an expected call of TriggerOrganizationModelBinaryFileUpload.
func (mr *MockModelPublicServiceClientMockRecorder) TriggerOrganizationModelBinaryFileUpload(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerOrganizationModelBinaryFileUpload", reflect.TypeOf((*MockModelPublicServiceClient)(nil).TriggerOrganizationModelBinaryFileUpload), varargs...)
}

// TriggerUserModel mocks base method.
func (m *MockModelPublicServiceClient) TriggerUserModel(arg0 context.Context, arg1 *modelv1alpha.TriggerUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.TriggerUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TriggerUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.TriggerUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerUserModel indicates an expected call of TriggerUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) TriggerUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).TriggerUserModel), varargs...)
}

// TriggerUserModelBinaryFileUpload mocks base method.
func (m *MockModelPublicServiceClient) TriggerUserModelBinaryFileUpload(arg0 context.Context, arg1 ...grpc.CallOption) (modelv1alpha.ModelPublicService_TriggerUserModelBinaryFileUploadClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TriggerUserModelBinaryFileUpload", varargs...)
	ret0, _ := ret[0].(modelv1alpha.ModelPublicService_TriggerUserModelBinaryFileUploadClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerUserModelBinaryFileUpload indicates an expected call of TriggerUserModelBinaryFileUpload.
func (mr *MockModelPublicServiceClientMockRecorder) TriggerUserModelBinaryFileUpload(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerUserModelBinaryFileUpload", reflect.TypeOf((*MockModelPublicServiceClient)(nil).TriggerUserModelBinaryFileUpload), varargs...)
}

// UndeployOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) UndeployOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.UndeployOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.UndeployOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UndeployOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.UndeployOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UndeployOrganizationModel indicates an expected call of UndeployOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) UndeployOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndeployOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).UndeployOrganizationModel), varargs...)
}

// UndeployUserModel mocks base method.
func (m *MockModelPublicServiceClient) UndeployUserModel(arg0 context.Context, arg1 *modelv1alpha.UndeployUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.UndeployUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UndeployUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.UndeployUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UndeployUserModel indicates an expected call of UndeployUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) UndeployUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndeployUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).UndeployUserModel), varargs...)
}

// UnpublishOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) UnpublishOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.UnpublishOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.UnpublishOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnpublishOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.UnpublishOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpublishOrganizationModel indicates an expected call of UnpublishOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) UnpublishOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpublishOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).UnpublishOrganizationModel), varargs...)
}

// UnpublishUserModel mocks base method.
func (m *MockModelPublicServiceClient) UnpublishUserModel(arg0 context.Context, arg1 *modelv1alpha.UnpublishUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.UnpublishUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnpublishUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.UnpublishUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpublishUserModel indicates an expected call of UnpublishUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) UnpublishUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpublishUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).UnpublishUserModel), varargs...)
}

// UpdateOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) UpdateOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.UpdateOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.UpdateOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.UpdateOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrganizationModel indicates an expected call of UpdateOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) UpdateOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).UpdateOrganizationModel), varargs...)
}

// UpdateUserModel mocks base method.
func (m *MockModelPublicServiceClient) UpdateUserModel(arg0 context.Context, arg1 *modelv1alpha.UpdateUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.UpdateUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.UpdateUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserModel indicates an expected call of UpdateUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) UpdateUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).UpdateUserModel), varargs...)
}

// WatchOrganizationModel mocks base method.
func (m *MockModelPublicServiceClient) WatchOrganizationModel(arg0 context.Context, arg1 *modelv1alpha.WatchOrganizationModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.WatchOrganizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WatchOrganizationModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.WatchOrganizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchOrganizationModel indicates an expected call of WatchOrganizationModel.
func (mr *MockModelPublicServiceClientMockRecorder) WatchOrganizationModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOrganizationModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).WatchOrganizationModel), varargs...)
}

// WatchUserModel mocks base method.
func (m *MockModelPublicServiceClient) WatchUserModel(arg0 context.Context, arg1 *modelv1alpha.WatchUserModelRequest, arg2 ...grpc.CallOption) (*modelv1alpha.WatchUserModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WatchUserModel", varargs...)
	ret0, _ := ret[0].(*modelv1alpha.WatchUserModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUserModel indicates an expected call of WatchUserModel.
func (mr *MockModelPublicServiceClientMockRecorder) WatchUserModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUserModel", reflect.TypeOf((*MockModelPublicServiceClient)(nil).WatchUserModel), varargs...)
}
