// Code generated by MockGen. DO NOT EDIT.
// Source: repository/querier.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// GetAllMembers mocks base method.
func (m *MockQuerier) GetAllMembers(ctx context.Context, db DBTX) ([]Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMembers", ctx, db)
	ret0, _ := ret[0].([]Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMembers indicates an expected call of GetAllMembers.
func (mr *MockQuerierMockRecorder) GetAllMembers(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMembers", reflect.TypeOf((*MockQuerier)(nil).GetAllMembers), ctx, db)
}

// GetAllScenes mocks base method.
func (m *MockQuerier) GetAllScenes(ctx context.Context, db DBTX) ([]GetAllScenesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllScenes", ctx, db)
	ret0, _ := ret[0].([]GetAllScenesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllScenes indicates an expected call of GetAllScenes.
func (mr *MockQuerierMockRecorder) GetAllScenes(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllScenes", reflect.TypeOf((*MockQuerier)(nil).GetAllScenes), ctx, db)
}

// GetGroup mocks base method.
func (m *MockQuerier) GetGroup(ctx context.Context, db DBTX) ([]GetGroupRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", ctx, db)
	ret0, _ := ret[0].([]GetGroupRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockQuerierMockRecorder) GetGroup(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockQuerier)(nil).GetGroup), ctx, db)
}

// GetGroupNameById mocks base method.
func (m *MockQuerier) GetGroupNameById(ctx context.Context, db DBTX, id int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupNameById", ctx, db, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupNameById indicates an expected call of GetGroupNameById.
func (mr *MockQuerierMockRecorder) GetGroupNameById(ctx, db, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupNameById", reflect.TypeOf((*MockQuerier)(nil).GetGroupNameById), ctx, db, id)
}

// GetLiveList mocks base method.
func (m *MockQuerier) GetLiveList(ctx context.Context, db DBTX) ([]GetLiveListRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLiveList", ctx, db)
	ret0, _ := ret[0].([]GetLiveListRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLiveList indicates an expected call of GetLiveList.
func (mr *MockQuerierMockRecorder) GetLiveList(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLiveList", reflect.TypeOf((*MockQuerier)(nil).GetLiveList), ctx, db)
}

// GetMembers mocks base method.
func (m *MockQuerier) GetMembers(ctx context.Context, db DBTX) ([]GetMembersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembers", ctx, db)
	ret0, _ := ret[0].([]GetMembersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembers indicates an expected call of GetMembers.
func (mr *MockQuerierMockRecorder) GetMembers(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockQuerier)(nil).GetMembers), ctx, db)
}

// GetMembersByGroup mocks base method.
func (m *MockQuerier) GetMembersByGroup(ctx context.Context, db DBTX, groupID int64) ([]GetMembersByGroupRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembersByGroup", ctx, db, groupID)
	ret0, _ := ret[0].([]GetMembersByGroupRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembersByGroup indicates an expected call of GetMembersByGroup.
func (mr *MockQuerierMockRecorder) GetMembersByGroup(ctx, db, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembersByGroup", reflect.TypeOf((*MockQuerier)(nil).GetMembersByGroup), ctx, db, groupID)
}

// GetMusicList mocks base method.
func (m *MockQuerier) GetMusicList(ctx context.Context, db DBTX) ([]GetMusicListRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusicList", ctx, db)
	ret0, _ := ret[0].([]GetMusicListRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusicList indicates an expected call of GetMusicList.
func (mr *MockQuerierMockRecorder) GetMusicList(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusicList", reflect.TypeOf((*MockQuerier)(nil).GetMusicList), ctx, db)
}

// GetMusicListWithColor mocks base method.
func (m *MockQuerier) GetMusicListWithColor(ctx context.Context, db DBTX, name string) ([]GetMusicListWithColorRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusicListWithColor", ctx, db, name)
	ret0, _ := ret[0].([]GetMusicListWithColorRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusicListWithColor indicates an expected call of GetMusicListWithColor.
func (mr *MockQuerierMockRecorder) GetMusicListWithColor(ctx, db, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusicListWithColor", reflect.TypeOf((*MockQuerier)(nil).GetMusicListWithColor), ctx, db, name)
}

// GetPhotographByGroupId mocks base method.
func (m *MockQuerier) GetPhotographByGroupId(ctx context.Context, db DBTX, groupID int64) ([]GetPhotographByGroupIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotographByGroupId", ctx, db, groupID)
	ret0, _ := ret[0].([]GetPhotographByGroupIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotographByGroupId indicates an expected call of GetPhotographByGroupId.
func (mr *MockQuerierMockRecorder) GetPhotographByGroupId(ctx, db, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotographByGroupId", reflect.TypeOf((*MockQuerier)(nil).GetPhotographByGroupId), ctx, db, groupID)
}

// GetPhotographList mocks base method.
func (m *MockQuerier) GetPhotographList(ctx context.Context, db DBTX, arg GetPhotographListParams) ([]GetPhotographListRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotographList", ctx, db, arg)
	ret0, _ := ret[0].([]GetPhotographListRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotographList indicates an expected call of GetPhotographList.
func (mr *MockQuerierMockRecorder) GetPhotographList(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotographList", reflect.TypeOf((*MockQuerier)(nil).GetPhotographList), ctx, db, arg)
}

// GetPhotographListAll mocks base method.
func (m *MockQuerier) GetPhotographListAll(ctx context.Context, db DBTX) ([]GetPhotographListAllRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotographListAll", ctx, db)
	ret0, _ := ret[0].([]GetPhotographListAllRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotographListAll indicates an expected call of GetPhotographListAll.
func (mr *MockQuerierMockRecorder) GetPhotographListAll(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotographListAll", reflect.TypeOf((*MockQuerier)(nil).GetPhotographListAll), ctx, db)
}

// GetPhotographListByPhotoType mocks base method.
func (m *MockQuerier) GetPhotographListByPhotoType(ctx context.Context, db DBTX, photoType string) ([]GetPhotographListByPhotoTypeRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotographListByPhotoType", ctx, db, photoType)
	ret0, _ := ret[0].([]GetPhotographListByPhotoTypeRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotographListByPhotoType indicates an expected call of GetPhotographListByPhotoType.
func (mr *MockQuerierMockRecorder) GetPhotographListByPhotoType(ctx, db, photoType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotographListByPhotoType", reflect.TypeOf((*MockQuerier)(nil).GetPhotographListByPhotoType), ctx, db, photoType)
}

// GetProducerMember mocks base method.
func (m *MockQuerier) GetProducerMember(ctx context.Context, db DBTX) ([]GetProducerMemberRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducerMember", ctx, db)
	ret0, _ := ret[0].([]GetProducerMemberRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducerMember indicates an expected call of GetProducerMember.
func (mr *MockQuerierMockRecorder) GetProducerMember(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducerMember", reflect.TypeOf((*MockQuerier)(nil).GetProducerMember), ctx, db)
}

// GetProducerOffice mocks base method.
func (m *MockQuerier) GetProducerOffice(ctx context.Context, db DBTX, producerID int64) (ProducerOffice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducerOffice", ctx, db, producerID)
	ret0, _ := ret[0].(ProducerOffice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducerOffice indicates an expected call of GetProducerOffice.
func (mr *MockQuerierMockRecorder) GetProducerOffice(ctx, db, producerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducerOffice", reflect.TypeOf((*MockQuerier)(nil).GetProducerOffice), ctx, db, producerID)
}

// GetProducerScenes mocks base method.
func (m *MockQuerier) GetProducerScenes(ctx context.Context, db DBTX, arg GetProducerScenesParams) ([]GetProducerScenesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducerScenes", ctx, db, arg)
	ret0, _ := ret[0].([]GetProducerScenesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducerScenes indicates an expected call of GetProducerScenes.
func (mr *MockQuerierMockRecorder) GetProducerScenes(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducerScenes", reflect.TypeOf((*MockQuerier)(nil).GetProducerScenes), ctx, db, arg)
}

// GetProducerScenesByGroupId mocks base method.
func (m *MockQuerier) GetProducerScenesByGroupId(ctx context.Context, db DBTX, groupID int64) ([]GetProducerScenesByGroupIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducerScenesByGroupId", ctx, db, groupID)
	ret0, _ := ret[0].([]GetProducerScenesByGroupIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducerScenesByGroupId indicates an expected call of GetProducerScenesByGroupId.
func (mr *MockQuerierMockRecorder) GetProducerScenesByGroupId(ctx, db, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducerScenesByGroupId", reflect.TypeOf((*MockQuerier)(nil).GetProducerScenesByGroupId), ctx, db, groupID)
}

// GetScenesWithColor mocks base method.
func (m *MockQuerier) GetScenesWithColor(ctx context.Context, db DBTX, arg GetScenesWithColorParams) ([]GetScenesWithColorRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScenesWithColor", ctx, db, arg)
	ret0, _ := ret[0].([]GetScenesWithColorRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScenesWithColor indicates an expected call of GetScenesWithColor.
func (mr *MockQuerierMockRecorder) GetScenesWithColor(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScenesWithColor", reflect.TypeOf((*MockQuerier)(nil).GetScenesWithColor), ctx, db, arg)
}

// GetScenesWithGroupId mocks base method.
func (m *MockQuerier) GetScenesWithGroupId(ctx context.Context, db DBTX, groupID int64) ([]GetScenesWithGroupIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScenesWithGroupId", ctx, db, groupID)
	ret0, _ := ret[0].([]GetScenesWithGroupIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScenesWithGroupId indicates an expected call of GetScenesWithGroupId.
func (mr *MockQuerierMockRecorder) GetScenesWithGroupId(ctx, db, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScenesWithGroupId", reflect.TypeOf((*MockQuerier)(nil).GetScenesWithGroupId), ctx, db, groupID)
}

// GetSsrPlusReleasedPhotographList mocks base method.
func (m *MockQuerier) GetSsrPlusReleasedPhotographList(ctx context.Context, db DBTX) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSsrPlusReleasedPhotographList", ctx, db)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSsrPlusReleasedPhotographList indicates an expected call of GetSsrPlusReleasedPhotographList.
func (mr *MockQuerierMockRecorder) GetSsrPlusReleasedPhotographList(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSsrPlusReleasedPhotographList", reflect.TypeOf((*MockQuerier)(nil).GetSsrPlusReleasedPhotographList), ctx, db)
}

// InitProducerSceneAll mocks base method.
func (m *MockQuerier) InitProducerSceneAll(ctx context.Context, db DBTX, arg InitProducerSceneAllParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitProducerSceneAll", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitProducerSceneAll indicates an expected call of InitProducerSceneAll.
func (mr *MockQuerierMockRecorder) InitProducerSceneAll(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitProducerSceneAll", reflect.TypeOf((*MockQuerier)(nil).InitProducerSceneAll), ctx, db, arg)
}

// RegistLive mocks base method.
func (m *MockQuerier) RegistLive(ctx context.Context, db DBTX, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistLive", ctx, db, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistLive indicates an expected call of RegistLive.
func (mr *MockQuerierMockRecorder) RegistLive(ctx, db, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistLive", reflect.TypeOf((*MockQuerier)(nil).RegistLive), ctx, db, name)
}

// RegistMusic mocks base method.
func (m *MockQuerier) RegistMusic(ctx context.Context, db DBTX, arg RegistMusicParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistMusic", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistMusic indicates an expected call of RegistMusic.
func (mr *MockQuerierMockRecorder) RegistMusic(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistMusic", reflect.TypeOf((*MockQuerier)(nil).RegistMusic), ctx, db, arg)
}

// RegistPhotograph mocks base method.
func (m *MockQuerier) RegistPhotograph(ctx context.Context, db DBTX, arg RegistPhotographParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistPhotograph", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistPhotograph indicates an expected call of RegistPhotograph.
func (mr *MockQuerierMockRecorder) RegistPhotograph(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistPhotograph", reflect.TypeOf((*MockQuerier)(nil).RegistPhotograph), ctx, db, arg)
}

// RegistProducerMember mocks base method.
func (m *MockQuerier) RegistProducerMember(ctx context.Context, db DBTX, arg RegistProducerMemberParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistProducerMember", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistProducerMember indicates an expected call of RegistProducerMember.
func (mr *MockQuerierMockRecorder) RegistProducerMember(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistProducerMember", reflect.TypeOf((*MockQuerier)(nil).RegistProducerMember), ctx, db, arg)
}

// RegistProducerOffice mocks base method.
func (m *MockQuerier) RegistProducerOffice(ctx context.Context, db DBTX, producerID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistProducerOffice", ctx, db, producerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistProducerOffice indicates an expected call of RegistProducerOffice.
func (mr *MockQuerierMockRecorder) RegistProducerOffice(ctx, db, producerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistProducerOffice", reflect.TypeOf((*MockQuerier)(nil).RegistProducerOffice), ctx, db, producerID)
}

// RegistProducerScene mocks base method.
func (m *MockQuerier) RegistProducerScene(ctx context.Context, db DBTX, arg RegistProducerSceneParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistProducerScene", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistProducerScene indicates an expected call of RegistProducerScene.
func (mr *MockQuerierMockRecorder) RegistProducerScene(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistProducerScene", reflect.TypeOf((*MockQuerier)(nil).RegistProducerScene), ctx, db, arg)
}

// RegistScene mocks base method.
func (m *MockQuerier) RegistScene(ctx context.Context, db DBTX, arg RegistSceneParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistScene", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistScene indicates an expected call of RegistScene.
func (mr *MockQuerierMockRecorder) RegistScene(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistScene", reflect.TypeOf((*MockQuerier)(nil).RegistScene), ctx, db, arg)
}

// UpdateProducerMember mocks base method.
func (m *MockQuerier) UpdateProducerMember(ctx context.Context, db DBTX, arg UpdateProducerMemberParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProducerMember", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProducerMember indicates an expected call of UpdateProducerMember.
func (mr *MockQuerierMockRecorder) UpdateProducerMember(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProducerMember", reflect.TypeOf((*MockQuerier)(nil).UpdateProducerMember), ctx, db, arg)
}

// UpdateProducerOffice mocks base method.
func (m *MockQuerier) UpdateProducerOffice(ctx context.Context, db DBTX, arg UpdateProducerOfficeParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProducerOffice", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProducerOffice indicates an expected call of UpdateProducerOffice.
func (mr *MockQuerierMockRecorder) UpdateProducerOffice(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProducerOffice", reflect.TypeOf((*MockQuerier)(nil).UpdateProducerOffice), ctx, db, arg)
}

// UpdateProducerScene mocks base method.
func (m *MockQuerier) UpdateProducerScene(ctx context.Context, db DBTX, arg UpdateProducerSceneParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProducerScene", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProducerScene indicates an expected call of UpdateProducerScene.
func (mr *MockQuerierMockRecorder) UpdateProducerScene(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProducerScene", reflect.TypeOf((*MockQuerier)(nil).UpdateProducerScene), ctx, db, arg)
}
