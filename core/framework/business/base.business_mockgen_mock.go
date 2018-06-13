// Code generated by MockGen. DO NOT EDIT.
// Source: core/framework/business/base.business.go

// Package mock_business is a generated GoMock package.
package business

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBaseBusinessModel is a mock of BaseBusinessModel interface
type MockBaseBusinessModel struct {
	ctrl     *gomock.Controller
	recorder *MockBaseBusinessModelMockRecorder
}

// MockBaseBusinessModelMockRecorder is the mock recorder for MockBaseBusinessModel
type MockBaseBusinessModelMockRecorder struct {
	mock *MockBaseBusinessModel
}

// NewMockBaseBusinessModel creates a new mock instance
func NewMockBaseBusinessModel(ctrl *gomock.Controller) *MockBaseBusinessModel {
	mock := &MockBaseBusinessModel{ctrl: ctrl}
	mock.recorder = &MockBaseBusinessModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBaseBusinessModel) EXPECT() *MockBaseBusinessModelMockRecorder {
	return m.recorder
}

// GetModel mocks base method
func (m *MockBaseBusinessModel) GetModel() (interface{}, error) {
	ret := m.ctrl.Call(m, "GetModel")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModel indicates an expected call of GetModel
func (mr *MockBaseBusinessModelMockRecorder) GetModel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModel", reflect.TypeOf((*MockBaseBusinessModel)(nil).GetModel))
}

// Validate mocks base method
func (m *MockBaseBusinessModel) Validate(obj interface{}) error {
	ret := m.ctrl.Call(m, "Validate", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockBaseBusinessModelMockRecorder) Validate(obj interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockBaseBusinessModel)(nil).Validate), obj)
}

// MockBaseBusinessModelWithContext is a mock of BaseBusinessModelWithContext interface
type MockBaseBusinessModelWithContext struct {
	ctrl     *gomock.Controller
	recorder *MockBaseBusinessModelWithContextMockRecorder
}

// MockBaseBusinessModelWithContextMockRecorder is the mock recorder for MockBaseBusinessModelWithContext
type MockBaseBusinessModelWithContextMockRecorder struct {
	mock *MockBaseBusinessModelWithContext
}

// NewMockBaseBusinessModelWithContext creates a new mock instance
func NewMockBaseBusinessModelWithContext(ctrl *gomock.Controller) *MockBaseBusinessModelWithContext {
	mock := &MockBaseBusinessModelWithContext{ctrl: ctrl}
	mock.recorder = &MockBaseBusinessModelWithContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBaseBusinessModelWithContext) EXPECT() *MockBaseBusinessModelWithContextMockRecorder {
	return m.recorder
}

// GetModelCtx mocks base method
func (m *MockBaseBusinessModelWithContext) GetModelCtx(ctx context.Context) (interface{}, error) {
	ret := m.ctrl.Call(m, "GetModelCtx", ctx)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelCtx indicates an expected call of GetModelCtx
func (mr *MockBaseBusinessModelWithContextMockRecorder) GetModelCtx(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelCtx", reflect.TypeOf((*MockBaseBusinessModelWithContext)(nil).GetModelCtx), ctx)
}

// Validate mocks base method
func (m *MockBaseBusinessModelWithContext) Validate(obj interface{}) error {
	ret := m.ctrl.Call(m, "Validate", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockBaseBusinessModelWithContextMockRecorder) Validate(obj interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockBaseBusinessModelWithContext)(nil).Validate), obj)
}
