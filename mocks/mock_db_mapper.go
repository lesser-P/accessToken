// Code generated by MockGen. DO NOT EDIT.
// Source: i_token.go

// Package mock_model is a generated GoMock package.
package mocks

import (
	model "accessToken_go_zero/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTokenDao is a mock of TokenDao interface.
type MockTokenDao struct {
	ctrl     *gomock.Controller
	recorder *MockTokenDaoMockRecorder
}

// MockTokenDaoMockRecorder is the mock recorder for MockTokenDao.
type MockTokenDaoMockRecorder struct {
	mock *MockTokenDao
}

// NewMockTokenDao creates a new mock instance.
func NewMockTokenDao(ctrl *gomock.Controller) *MockTokenDao {
	mock := &MockTokenDao{ctrl: ctrl}
	mock.recorder = &MockTokenDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenDao) EXPECT() *MockTokenDaoMockRecorder {
	return m.recorder
}

// DeleteByUserId mocks base method.
func (m *MockTokenDao) DeleteByUserId(userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUserId", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByUserId indicates an expected call of DeleteByUserId.
func (mr *MockTokenDaoMockRecorder) DeleteByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUserId", reflect.TypeOf((*MockTokenDao)(nil).DeleteByUserId), userId)
}

// FreezeToken mocks base method.
func (m *MockTokenDao) FreezeToken(token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FreezeToken", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// FreezeToken indicates an expected call of FreezeToken.
func (mr *MockTokenDaoMockRecorder) FreezeToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreezeToken", reflect.TypeOf((*MockTokenDao)(nil).FreezeToken), token)
}

// SaveOrUpdate mocks base method.
func (m *MockTokenDao) SaveOrUpdate(info *model.TokenDetail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOrUpdate", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOrUpdate indicates an expected call of SaveOrUpdate.
func (mr *MockTokenDaoMockRecorder) SaveOrUpdate(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOrUpdate", reflect.TypeOf((*MockTokenDao)(nil).SaveOrUpdate), info)
}

// SelectByToken mocks base method.
func (m *MockTokenDao) SelectByToken(token string) (*model.TokenDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByToken", token)
	ret0, _ := ret[0].(*model.TokenDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByToken indicates an expected call of SelectByToken.
func (mr *MockTokenDaoMockRecorder) SelectByToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByToken", reflect.TypeOf((*MockTokenDao)(nil).SelectByToken), token)
}

// SelectByUserId mocks base method.
func (m *MockTokenDao) SelectByUserId(userId string) (*model.TokenDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByUserId", userId)
	ret0, _ := ret[0].(*model.TokenDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByUserId indicates an expected call of SelectByUserId.
func (mr *MockTokenDaoMockRecorder) SelectByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByUserId", reflect.TypeOf((*MockTokenDao)(nil).SelectByUserId), userId)
}