// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/devpablocristo/nanlabs/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockService) CreateTask(arg0 context.Context, arg1 *domain.Task) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0, arg1)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockServiceMockRecorder) CreateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockService)(nil).CreateTask), arg0, arg1)
}

// CreateTrelloCard mocks base method.
func (m *MockService) CreateTrelloCard(arg0 context.Context, arg1 *domain.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrelloCard", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTrelloCard indicates an expected call of CreateTrelloCard.
func (mr *MockServiceMockRecorder) CreateTrelloCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrelloCard", reflect.TypeOf((*MockService)(nil).CreateTrelloCard), arg0, arg1)
}

// DeleteTask mocks base method.
func (m *MockService) DeleteTask(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockServiceMockRecorder) DeleteTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockService)(nil).DeleteTask), arg0, arg1)
}

// GetTask mocks base method.
func (m *MockService) GetTask(arg0 context.Context, arg1 string) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0, arg1)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockServiceMockRecorder) GetTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockService)(nil).GetTask), arg0, arg1)
}

// GetTasks mocks base method.
func (m *MockService) GetTasks(arg0 context.Context) (map[string]*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks", arg0)
	ret0, _ := ret[0].(map[string]*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockServiceMockRecorder) GetTasks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockService)(nil).GetTasks), arg0)
}

// UpdateTask mocks base method.
func (m *MockService) UpdateTask(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockServiceMockRecorder) UpdateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockService)(nil).UpdateTask), arg0, arg1)
}