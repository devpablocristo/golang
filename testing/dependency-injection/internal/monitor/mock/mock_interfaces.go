// Code generated by MockGen. DO NOT EDIT.
// Source: internal/monitor/interfaces.go

// Package mock_monitor is a generated GoMock package.
package mock_monitor

import (
	context "context"
	monitor "demo-project/internal/monitor"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMonitorRepositoryPort is a mock of MonitorRepositoryPort interface.
type MockMonitorRepositoryPort struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorRepositoryPortMockRecorder
}

// MockMonitorRepositoryPortMockRecorder is the mock recorder for MockMonitorRepositoryPort.
type MockMonitorRepositoryPortMockRecorder struct {
	mock *MockMonitorRepositoryPort
}

// NewMockMonitorRepositoryPort creates a new mock instance.
func NewMockMonitorRepositoryPort(ctrl *gomock.Controller) *MockMonitorRepositoryPort {
	mock := &MockMonitorRepositoryPort{ctrl: ctrl}
	mock.recorder = &MockMonitorRepositoryPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMonitorRepositoryPort) EXPECT() *MockMonitorRepositoryPortMockRecorder {
	return m.recorder
}

// CheckMonitorExists mocks base method.
func (m *MockMonitorRepositoryPort) CheckMonitorExists(arg0 *monitor.Monitor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckMonitorExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckMonitorExists indicates an expected call of CheckMonitorExists.
func (mr *MockMonitorRepositoryPortMockRecorder) CheckMonitorExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckMonitorExists", reflect.TypeOf((*MockMonitorRepositoryPort)(nil).CheckMonitorExists), arg0)
}

// CreateMonitor mocks base method.
func (m *MockMonitorRepositoryPort) CreateMonitor(arg0 *monitor.Monitor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMonitor", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMonitor indicates an expected call of CreateMonitor.
func (mr *MockMonitorRepositoryPortMockRecorder) CreateMonitor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMonitor", reflect.TypeOf((*MockMonitorRepositoryPort)(nil).CreateMonitor), arg0)
}

// MockMonitorDatadogPort is a mock of MonitorDatadogPort interface.
type MockMonitorDatadogPort struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorDatadogPortMockRecorder
}

// MockMonitorDatadogPortMockRecorder is the mock recorder for MockMonitorDatadogPort.
type MockMonitorDatadogPortMockRecorder struct {
	mock *MockMonitorDatadogPort
}

// NewMockMonitorDatadogPort creates a new mock instance.
func NewMockMonitorDatadogPort(ctrl *gomock.Controller) *MockMonitorDatadogPort {
	mock := &MockMonitorDatadogPort{ctrl: ctrl}
	mock.recorder = &MockMonitorDatadogPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMonitorDatadogPort) EXPECT() *MockMonitorDatadogPortMockRecorder {
	return m.recorder
}

// CreateMonitor mocks base method.
func (m *MockMonitorDatadogPort) CreateMonitor(ctx context.Context, options ...interface{}) (*monitor.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMonitor", varargs...)
	ret0, _ := ret[0].(*monitor.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMonitor indicates an expected call of CreateMonitor.
func (mr *MockMonitorDatadogPortMockRecorder) CreateMonitor(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMonitor", reflect.TypeOf((*MockMonitorDatadogPort)(nil).CreateMonitor), varargs...)
}

// MockMonitorUsecasePort is a mock of MonitorUsecasePort interface.
type MockMonitorUsecasePort struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorUsecasePortMockRecorder
}

// MockMonitorUsecasePortMockRecorder is the mock recorder for MockMonitorUsecasePort.
type MockMonitorUsecasePortMockRecorder struct {
	mock *MockMonitorUsecasePort
}

// NewMockMonitorUsecasePort creates a new mock instance.
func NewMockMonitorUsecasePort(ctrl *gomock.Controller) *MockMonitorUsecasePort {
	mock := &MockMonitorUsecasePort{ctrl: ctrl}
	mock.recorder = &MockMonitorUsecasePortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMonitorUsecasePort) EXPECT() *MockMonitorUsecasePortMockRecorder {
	return m.recorder
}

// CreateMonitor mocks base method.
func (m_2 *MockMonitorUsecasePort) CreateMonitor(ctx context.Context, m *monitor.Monitor) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "CreateMonitor", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMonitor indicates an expected call of CreateMonitor.
func (mr *MockMonitorUsecasePortMockRecorder) CreateMonitor(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMonitor", reflect.TypeOf((*MockMonitorUsecasePort)(nil).CreateMonitor), ctx, m)
}

// PlatformOrBrand mocks base method.
func (m_2 *MockMonitorUsecasePort) PlatformOrBrand(ctx context.Context, m *monitor.Monitor, query []string) ([]string, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "PlatformOrBrand", ctx, m, query)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlatformOrBrand indicates an expected call of PlatformOrBrand.
func (mr *MockMonitorUsecasePortMockRecorder) PlatformOrBrand(ctx, m, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlatformOrBrand", reflect.TypeOf((*MockMonitorUsecasePort)(nil).PlatformOrBrand), ctx, m, query)
}
