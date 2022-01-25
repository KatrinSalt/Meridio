// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/ambassador/tap/types/trench.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/nordix/meridio/api/nsp/v1"
	types "github.com/nordix/meridio/pkg/ambassador/tap/types"
)

// MockTrench is a mock of Trench interface.
type MockTrench struct {
	ctrl     *gomock.Controller
	recorder *MockTrenchMockRecorder
}

// MockTrenchMockRecorder is the mock recorder for MockTrench.
type MockTrenchMockRecorder struct {
	mock *MockTrench
}

// NewMockTrench creates a new mock instance.
func NewMockTrench(ctrl *gomock.Controller) *MockTrench {
	mock := &MockTrench{ctrl: ctrl}
	mock.recorder = &MockTrenchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrench) EXPECT() *MockTrenchMockRecorder {
	return m.recorder
}

// AddConduit mocks base method.
func (m *MockTrench) AddConduit(arg0 context.Context, arg1 *v1.Conduit) (types.Conduit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddConduit", arg0, arg1)
	ret0, _ := ret[0].(types.Conduit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddConduit indicates an expected call of AddConduit.
func (mr *MockTrenchMockRecorder) AddConduit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConduit", reflect.TypeOf((*MockTrench)(nil).AddConduit), arg0, arg1)
}

// Delete mocks base method.
func (m *MockTrench) Delete(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTrenchMockRecorder) Delete(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrench)(nil).Delete), ctx)
}

// Equals mocks base method.
func (m *MockTrench) Equals(arg0 *v1.Trench) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equals", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equals indicates an expected call of Equals.
func (mr *MockTrenchMockRecorder) Equals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equals", reflect.TypeOf((*MockTrench)(nil).Equals), arg0)
}

// GetConduit mocks base method.
func (m *MockTrench) GetConduit(arg0 *v1.Conduit) types.Conduit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConduit", arg0)
	ret0, _ := ret[0].(types.Conduit)
	return ret0
}

// GetConduit indicates an expected call of GetConduit.
func (mr *MockTrenchMockRecorder) GetConduit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConduit", reflect.TypeOf((*MockTrench)(nil).GetConduit), arg0)
}

// GetConduits mocks base method.
func (m *MockTrench) GetConduits() []types.Conduit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConduits")
	ret0, _ := ret[0].([]types.Conduit)
	return ret0
}

// GetConduits indicates an expected call of GetConduits.
func (mr *MockTrenchMockRecorder) GetConduits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConduits", reflect.TypeOf((*MockTrench)(nil).GetConduits))
}

// RemoveConduit mocks base method.
func (m *MockTrench) RemoveConduit(arg0 context.Context, arg1 *v1.Conduit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveConduit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveConduit indicates an expected call of RemoveConduit.
func (mr *MockTrenchMockRecorder) RemoveConduit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveConduit", reflect.TypeOf((*MockTrench)(nil).RemoveConduit), arg0, arg1)
}
