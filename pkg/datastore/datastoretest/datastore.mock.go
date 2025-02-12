// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pipe-cd/pipecd/pkg/datastore (interfaces: ProjectStore,PipedStore,ApplicationStore,DeploymentStore,CommandStore)

// Package datastoretest is a generated GoMock package.
package datastoretest

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	datastore "github.com/pipe-cd/pipecd/pkg/datastore"
	model "github.com/pipe-cd/pipecd/pkg/model"
)

// MockProjectStore is a mock of ProjectStore interface.
type MockProjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockProjectStoreMockRecorder
}

// MockProjectStoreMockRecorder is the mock recorder for MockProjectStore.
type MockProjectStoreMockRecorder struct {
	mock *MockProjectStore
}

// NewMockProjectStore creates a new mock instance.
func NewMockProjectStore(ctrl *gomock.Controller) *MockProjectStore {
	mock := &MockProjectStore{ctrl: ctrl}
	mock.recorder = &MockProjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectStore) EXPECT() *MockProjectStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockProjectStore) Add(arg0 context.Context, arg1 *model.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockProjectStoreMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockProjectStore)(nil).Add), arg0, arg1)
}

// DisableStaticAdmin mocks base method.
func (m *MockProjectStore) DisableStaticAdmin(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableStaticAdmin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableStaticAdmin indicates an expected call of DisableStaticAdmin.
func (mr *MockProjectStoreMockRecorder) DisableStaticAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableStaticAdmin", reflect.TypeOf((*MockProjectStore)(nil).DisableStaticAdmin), arg0, arg1)
}

// EnableStaticAdmin mocks base method.
func (m *MockProjectStore) EnableStaticAdmin(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableStaticAdmin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableStaticAdmin indicates an expected call of EnableStaticAdmin.
func (mr *MockProjectStoreMockRecorder) EnableStaticAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableStaticAdmin", reflect.TypeOf((*MockProjectStore)(nil).EnableStaticAdmin), arg0, arg1)
}

// Get mocks base method.
func (m *MockProjectStore) Get(arg0 context.Context, arg1 string) (*model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProjectStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProjectStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockProjectStore) List(arg0 context.Context, arg1 datastore.ListOptions) ([]model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProjectStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProjectStore)(nil).List), arg0, arg1)
}

// MigrateFromProjectRBACConfig mocks base method.
func (m *MockProjectStore) MigrateFromProjectRBACConfig(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateFromProjectRBACConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MigrateFromProjectRBACConfig indicates an expected call of MigrateFromProjectRBACConfig.
func (mr *MockProjectStoreMockRecorder) MigrateFromProjectRBACConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateFromProjectRBACConfig", reflect.TypeOf((*MockProjectStore)(nil).MigrateFromProjectRBACConfig), arg0, arg1)
}

// UpdateProjectRBACConfig mocks base method.
func (m *MockProjectStore) UpdateProjectRBACConfig(arg0 context.Context, arg1 string, arg2 *model.ProjectRBACConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectRBACConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectRBACConfig indicates an expected call of UpdateProjectRBACConfig.
func (mr *MockProjectStoreMockRecorder) UpdateProjectRBACConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectRBACConfig", reflect.TypeOf((*MockProjectStore)(nil).UpdateProjectRBACConfig), arg0, arg1, arg2)
}

// UpdateProjectRBACRoles mocks base method.
func (m *MockProjectStore) UpdateProjectRBACRoles(arg0 context.Context, arg1 string, arg2 []*model.ProjectRBACRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectRBACRoles", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectRBACRoles indicates an expected call of UpdateProjectRBACRoles.
func (mr *MockProjectStoreMockRecorder) UpdateProjectRBACRoles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectRBACRoles", reflect.TypeOf((*MockProjectStore)(nil).UpdateProjectRBACRoles), arg0, arg1, arg2)
}

// UpdateProjectSSOConfig mocks base method.
func (m *MockProjectStore) UpdateProjectSSOConfig(arg0 context.Context, arg1 string, arg2 *model.ProjectSSOConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectSSOConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectSSOConfig indicates an expected call of UpdateProjectSSOConfig.
func (mr *MockProjectStoreMockRecorder) UpdateProjectSSOConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectSSOConfig", reflect.TypeOf((*MockProjectStore)(nil).UpdateProjectSSOConfig), arg0, arg1, arg2)
}

// UpdateProjectStaticAdmin mocks base method.
func (m *MockProjectStore) UpdateProjectStaticAdmin(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectStaticAdmin", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectStaticAdmin indicates an expected call of UpdateProjectStaticAdmin.
func (mr *MockProjectStoreMockRecorder) UpdateProjectStaticAdmin(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectStaticAdmin", reflect.TypeOf((*MockProjectStore)(nil).UpdateProjectStaticAdmin), arg0, arg1, arg2, arg3)
}

// UpdateProjectUserGroups mocks base method.
func (m *MockProjectStore) UpdateProjectUserGroups(arg0 context.Context, arg1 string, arg2 []*model.ProjectUserGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectUserGroups", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectUserGroups indicates an expected call of UpdateProjectUserGroups.
func (mr *MockProjectStoreMockRecorder) UpdateProjectUserGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectUserGroups", reflect.TypeOf((*MockProjectStore)(nil).UpdateProjectUserGroups), arg0, arg1, arg2)
}

// MockPipedStore is a mock of PipedStore interface.
type MockPipedStore struct {
	ctrl     *gomock.Controller
	recorder *MockPipedStoreMockRecorder
}

// MockPipedStoreMockRecorder is the mock recorder for MockPipedStore.
type MockPipedStoreMockRecorder struct {
	mock *MockPipedStore
}

// NewMockPipedStore creates a new mock instance.
func NewMockPipedStore(ctrl *gomock.Controller) *MockPipedStore {
	mock := &MockPipedStore{ctrl: ctrl}
	mock.recorder = &MockPipedStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipedStore) EXPECT() *MockPipedStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPipedStore) Add(arg0 context.Context, arg1 *model.Piped) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockPipedStoreMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPipedStore)(nil).Add), arg0, arg1)
}

// AddKey mocks base method.
func (m *MockPipedStore) AddKey(arg0 context.Context, arg1, arg2, arg3 string, arg4 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddKey", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddKey indicates an expected call of AddKey.
func (mr *MockPipedStoreMockRecorder) AddKey(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKey", reflect.TypeOf((*MockPipedStore)(nil).AddKey), arg0, arg1, arg2, arg3, arg4)
}

// DeleteOldKeys mocks base method.
func (m *MockPipedStore) DeleteOldKeys(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOldKeys", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOldKeys indicates an expected call of DeleteOldKeys.
func (mr *MockPipedStoreMockRecorder) DeleteOldKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOldKeys", reflect.TypeOf((*MockPipedStore)(nil).DeleteOldKeys), arg0, arg1)
}

// DisablePiped mocks base method.
func (m *MockPipedStore) DisablePiped(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisablePiped", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisablePiped indicates an expected call of DisablePiped.
func (mr *MockPipedStoreMockRecorder) DisablePiped(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisablePiped", reflect.TypeOf((*MockPipedStore)(nil).DisablePiped), arg0, arg1)
}

// EnablePiped mocks base method.
func (m *MockPipedStore) EnablePiped(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePiped", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnablePiped indicates an expected call of EnablePiped.
func (mr *MockPipedStoreMockRecorder) EnablePiped(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePiped", reflect.TypeOf((*MockPipedStore)(nil).EnablePiped), arg0, arg1)
}

// Get mocks base method.
func (m *MockPipedStore) Get(arg0 context.Context, arg1 string) (*model.Piped, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.Piped)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPipedStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPipedStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockPipedStore) List(arg0 context.Context, arg1 datastore.ListOptions) ([]*model.Piped, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*model.Piped)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPipedStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPipedStore)(nil).List), arg0, arg1)
}

// UpdateDesiredVersion mocks base method.
func (m *MockPipedStore) UpdateDesiredVersion(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDesiredVersion", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDesiredVersion indicates an expected call of UpdateDesiredVersion.
func (mr *MockPipedStoreMockRecorder) UpdateDesiredVersion(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDesiredVersion", reflect.TypeOf((*MockPipedStore)(nil).UpdateDesiredVersion), arg0, arg1, arg2)
}

// UpdateInfo mocks base method.
func (m *MockPipedStore) UpdateInfo(arg0 context.Context, arg1, arg2, arg3 string, arg4 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfo", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInfo indicates an expected call of UpdateInfo.
func (mr *MockPipedStoreMockRecorder) UpdateInfo(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfo", reflect.TypeOf((*MockPipedStore)(nil).UpdateInfo), arg0, arg1, arg2, arg3, arg4)
}

// UpdateMetadata mocks base method.
func (m *MockPipedStore) UpdateMetadata(arg0 context.Context, arg1, arg2 string, arg3 []*model.Piped_CloudProvider, arg4 []*model.ApplicationGitRepository, arg5 *model.Piped_SecretEncryption, arg6 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMetadata", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetadata indicates an expected call of UpdateMetadata.
func (mr *MockPipedStoreMockRecorder) UpdateMetadata(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetadata", reflect.TypeOf((*MockPipedStore)(nil).UpdateMetadata), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// MockApplicationStore is a mock of ApplicationStore interface.
type MockApplicationStore struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationStoreMockRecorder
}

// MockApplicationStoreMockRecorder is the mock recorder for MockApplicationStore.
type MockApplicationStoreMockRecorder struct {
	mock *MockApplicationStore
}

// NewMockApplicationStore creates a new mock instance.
func NewMockApplicationStore(ctrl *gomock.Controller) *MockApplicationStore {
	mock := &MockApplicationStore{ctrl: ctrl}
	mock.recorder = &MockApplicationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationStore) EXPECT() *MockApplicationStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockApplicationStore) Add(arg0 context.Context, arg1 *model.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockApplicationStoreMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockApplicationStore)(nil).Add), arg0, arg1)
}

// Delete mocks base method.
func (m *MockApplicationStore) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockApplicationStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockApplicationStore)(nil).Delete), arg0, arg1)
}

// Disable mocks base method.
func (m *MockApplicationStore) Disable(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disable indicates an expected call of Disable.
func (mr *MockApplicationStoreMockRecorder) Disable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockApplicationStore)(nil).Disable), arg0, arg1)
}

// Enable mocks base method.
func (m *MockApplicationStore) Enable(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockApplicationStoreMockRecorder) Enable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockApplicationStore)(nil).Enable), arg0, arg1)
}

// Get mocks base method.
func (m *MockApplicationStore) Get(arg0 context.Context, arg1 string) (*model.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockApplicationStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApplicationStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockApplicationStore) List(arg0 context.Context, arg1 datastore.ListOptions) ([]*model.Application, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*model.Application)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockApplicationStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockApplicationStore)(nil).List), arg0, arg1)
}

// UpdateBasicInfo mocks base method.
func (m *MockApplicationStore) UpdateBasicInfo(arg0 context.Context, arg1, arg2, arg3 string, arg4 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBasicInfo", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBasicInfo indicates an expected call of UpdateBasicInfo.
func (mr *MockApplicationStoreMockRecorder) UpdateBasicInfo(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBasicInfo", reflect.TypeOf((*MockApplicationStore)(nil).UpdateBasicInfo), arg0, arg1, arg2, arg3, arg4)
}

// UpdateConfigFilename mocks base method.
func (m *MockApplicationStore) UpdateConfigFilename(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfigFilename", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfigFilename indicates an expected call of UpdateConfigFilename.
func (mr *MockApplicationStoreMockRecorder) UpdateConfigFilename(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfigFilename", reflect.TypeOf((*MockApplicationStore)(nil).UpdateConfigFilename), arg0, arg1, arg2)
}

// UpdateConfiguration mocks base method.
func (m *MockApplicationStore) UpdateConfiguration(arg0 context.Context, arg1, arg2, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfiguration", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfiguration indicates an expected call of UpdateConfiguration.
func (mr *MockApplicationStoreMockRecorder) UpdateConfiguration(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfiguration", reflect.TypeOf((*MockApplicationStore)(nil).UpdateConfiguration), arg0, arg1, arg2, arg3, arg4)
}

// UpdateDeployingStatus mocks base method.
func (m *MockApplicationStore) UpdateDeployingStatus(arg0 context.Context, arg1 string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeployingStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeployingStatus indicates an expected call of UpdateDeployingStatus.
func (mr *MockApplicationStoreMockRecorder) UpdateDeployingStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeployingStatus", reflect.TypeOf((*MockApplicationStore)(nil).UpdateDeployingStatus), arg0, arg1, arg2)
}

// UpdateMostRecentDeployment mocks base method.
func (m *MockApplicationStore) UpdateMostRecentDeployment(arg0 context.Context, arg1 string, arg2 model.DeploymentStatus, arg3 *model.ApplicationDeploymentReference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMostRecentDeployment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMostRecentDeployment indicates an expected call of UpdateMostRecentDeployment.
func (mr *MockApplicationStoreMockRecorder) UpdateMostRecentDeployment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMostRecentDeployment", reflect.TypeOf((*MockApplicationStore)(nil).UpdateMostRecentDeployment), arg0, arg1, arg2, arg3)
}

// UpdateSyncState mocks base method.
func (m *MockApplicationStore) UpdateSyncState(arg0 context.Context, arg1 string, arg2 *model.ApplicationSyncState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSyncState", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSyncState indicates an expected call of UpdateSyncState.
func (mr *MockApplicationStoreMockRecorder) UpdateSyncState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSyncState", reflect.TypeOf((*MockApplicationStore)(nil).UpdateSyncState), arg0, arg1, arg2)
}

// MockDeploymentStore is a mock of DeploymentStore interface.
type MockDeploymentStore struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentStoreMockRecorder
}

// MockDeploymentStoreMockRecorder is the mock recorder for MockDeploymentStore.
type MockDeploymentStoreMockRecorder struct {
	mock *MockDeploymentStore
}

// NewMockDeploymentStore creates a new mock instance.
func NewMockDeploymentStore(ctrl *gomock.Controller) *MockDeploymentStore {
	mock := &MockDeploymentStore{ctrl: ctrl}
	mock.recorder = &MockDeploymentStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentStore) EXPECT() *MockDeploymentStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockDeploymentStore) Add(arg0 context.Context, arg1 *model.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockDeploymentStoreMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockDeploymentStore)(nil).Add), arg0, arg1)
}

// Get mocks base method.
func (m *MockDeploymentStore) Get(arg0 context.Context, arg1 string) (*model.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDeploymentStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeploymentStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockDeploymentStore) List(arg0 context.Context, arg1 datastore.ListOptions) ([]*model.Deployment, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*model.Deployment)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockDeploymentStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeploymentStore)(nil).List), arg0, arg1)
}

// UpdateMetadata mocks base method.
func (m *MockDeploymentStore) UpdateMetadata(arg0 context.Context, arg1 string, arg2 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMetadata", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetadata indicates an expected call of UpdateMetadata.
func (mr *MockDeploymentStoreMockRecorder) UpdateMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetadata", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateMetadata), arg0, arg1, arg2)
}

// UpdateStageMetadata mocks base method.
func (m *MockDeploymentStore) UpdateStageMetadata(arg0 context.Context, arg1, arg2 string, arg3 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStageMetadata", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStageMetadata indicates an expected call of UpdateStageMetadata.
func (mr *MockDeploymentStoreMockRecorder) UpdateStageMetadata(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStageMetadata", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateStageMetadata), arg0, arg1, arg2, arg3)
}

// UpdateStageStatus mocks base method.
func (m *MockDeploymentStore) UpdateStageStatus(arg0 context.Context, arg1, arg2 string, arg3 model.StageStatus, arg4 string, arg5 []string, arg6 bool, arg7 int32, arg8 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStageStatus", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStageStatus indicates an expected call of UpdateStageStatus.
func (mr *MockDeploymentStoreMockRecorder) UpdateStageStatus(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStageStatus", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateStageStatus), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// UpdateStatus mocks base method.
func (m *MockDeploymentStore) UpdateStatus(arg0 context.Context, arg1 string, arg2 model.DeploymentStatus, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockDeploymentStoreMockRecorder) UpdateStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateStatus), arg0, arg1, arg2, arg3)
}

// UpdateToCompleted mocks base method.
func (m *MockDeploymentStore) UpdateToCompleted(arg0 context.Context, arg1 string, arg2 model.DeploymentStatus, arg3 map[string]model.StageStatus, arg4 string, arg5 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToCompleted", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToCompleted indicates an expected call of UpdateToCompleted.
func (mr *MockDeploymentStoreMockRecorder) UpdateToCompleted(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToCompleted", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateToCompleted), arg0, arg1, arg2, arg3, arg4, arg5)
}

// UpdateToPlanned mocks base method.
func (m *MockDeploymentStore) UpdateToPlanned(arg0 context.Context, arg1, arg2, arg3, arg4, arg5, arg6 string, arg7 []*model.ArtifactVersion, arg8 []*model.PipelineStage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToPlanned", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToPlanned indicates an expected call of UpdateToPlanned.
func (mr *MockDeploymentStoreMockRecorder) UpdateToPlanned(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToPlanned", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateToPlanned), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// MockCommandStore is a mock of CommandStore interface.
type MockCommandStore struct {
	ctrl     *gomock.Controller
	recorder *MockCommandStoreMockRecorder
}

// MockCommandStoreMockRecorder is the mock recorder for MockCommandStore.
type MockCommandStoreMockRecorder struct {
	mock *MockCommandStore
}

// NewMockCommandStore creates a new mock instance.
func NewMockCommandStore(ctrl *gomock.Controller) *MockCommandStore {
	mock := &MockCommandStore{ctrl: ctrl}
	mock.recorder = &MockCommandStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandStore) EXPECT() *MockCommandStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCommandStore) Add(arg0 context.Context, arg1 *model.Command) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockCommandStoreMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCommandStore)(nil).Add), arg0, arg1)
}

// Get mocks base method.
func (m *MockCommandStore) Get(arg0 context.Context, arg1 string) (*model.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCommandStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCommandStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockCommandStore) List(arg0 context.Context, arg1 datastore.ListOptions) ([]*model.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*model.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCommandStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCommandStore)(nil).List), arg0, arg1)
}

// UpdateStatus mocks base method.
func (m *MockCommandStore) UpdateStatus(arg0 context.Context, arg1 string, arg2 model.CommandStatus, arg3 map[string]string, arg4 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockCommandStoreMockRecorder) UpdateStatus(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockCommandStore)(nil).UpdateStatus), arg0, arg1, arg2, arg3, arg4)
}
