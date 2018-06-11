// Code generated by MockGen. DO NOT EDIT.
// Source: controller/internal/enforcer/enforcer.go

// Package mockenforcer is a generated GoMock package.
package mockenforcer

import (
	context "context"
	reflect "reflect"

	portset "go.aporeto.io/trireme-lib/controller/internal/portset"
	fqconfig "go.aporeto.io/trireme-lib/controller/pkg/fqconfig"
	secrets "go.aporeto.io/trireme-lib/controller/pkg/secrets"
	policy "go.aporeto.io/trireme-lib/policy"
	gomock "github.com/golang/mock/gomock"
)

// MockEnforcer is a mock of Enforcer interface
// nolint
type MockEnforcer struct {
	ctrl     *gomock.Controller
	recorder *MockEnforcerMockRecorder
}

// MockEnforcerMockRecorder is the mock recorder for MockEnforcer
// nolint
type MockEnforcerMockRecorder struct {
	mock *MockEnforcer
}

// NewMockEnforcer creates a new mock instance
// nolint
func NewMockEnforcer(ctrl *gomock.Controller) *MockEnforcer {
	mock := &MockEnforcer{ctrl: ctrl}
	mock.recorder = &MockEnforcerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockEnforcer) EXPECT() *MockEnforcerMockRecorder {
	return m.recorder
}

// Enforce mocks base method
// nolint
func (m *MockEnforcer) Enforce(contextID string, puInfo *policy.PUInfo) error {
	ret := m.ctrl.Call(m, "Enforce", contextID, puInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enforce indicates an expected call of Enforce
// nolint
func (mr *MockEnforcerMockRecorder) Enforce(contextID, puInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enforce", reflect.TypeOf((*MockEnforcer)(nil).Enforce), contextID, puInfo)
}

// Unenforce mocks base method
// nolint
func (m *MockEnforcer) Unenforce(contextID string) error {
	ret := m.ctrl.Call(m, "Unenforce", contextID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unenforce indicates an expected call of Unenforce
// nolint
func (mr *MockEnforcerMockRecorder) Unenforce(contextID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unenforce", reflect.TypeOf((*MockEnforcer)(nil).Unenforce), contextID)
}

// GetFilterQueue mocks base method
// nolint
func (m *MockEnforcer) GetFilterQueue() *fqconfig.FilterQueue {
	ret := m.ctrl.Call(m, "GetFilterQueue")
	ret0, _ := ret[0].(*fqconfig.FilterQueue)
	return ret0
}

// GetFilterQueue indicates an expected call of GetFilterQueue
// nolint
func (mr *MockEnforcerMockRecorder) GetFilterQueue() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilterQueue", reflect.TypeOf((*MockEnforcer)(nil).GetFilterQueue))
}

// GetPortSetInstance mocks base method
// nolint
func (m *MockEnforcer) GetPortSetInstance() portset.PortSet {
	ret := m.ctrl.Call(m, "GetPortSetInstance")
	ret0, _ := ret[0].(portset.PortSet)
	return ret0
}

// GetPortSetInstance indicates an expected call of GetPortSetInstance
// nolint
func (mr *MockEnforcerMockRecorder) GetPortSetInstance() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPortSetInstance", reflect.TypeOf((*MockEnforcer)(nil).GetPortSetInstance))
}

// Run mocks base method
// nolint
func (m *MockEnforcer) Run(ctx context.Context) error {
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
// nolint
func (mr *MockEnforcerMockRecorder) Run(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockEnforcer)(nil).Run), ctx)
}

// UpdateSecrets mocks base method
// nolint
func (m *MockEnforcer) UpdateSecrets(secrets secrets.Secrets) error {
	ret := m.ctrl.Call(m, "UpdateSecrets", secrets)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecrets indicates an expected call of UpdateSecrets
// nolint
func (mr *MockEnforcerMockRecorder) UpdateSecrets(secrets interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecrets", reflect.TypeOf((*MockEnforcer)(nil).UpdateSecrets), secrets)
}
