// Code generated by MockGen. DO NOT EDIT.
// Source: task.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	dto "ggl_test/models/dto"
	entity "ggl_test/models/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITaskRepo is a mock of ITaskRepo interface.
type MockITaskRepo struct {
	ctrl     *gomock.Controller
	recorder *MockITaskRepoMockRecorder
}

// MockITaskRepoMockRecorder is the mock recorder for MockITaskRepo.
type MockITaskRepoMockRecorder struct {
	mock *MockITaskRepo
}

// NewMockITaskRepo creates a new mock instance.
func NewMockITaskRepo(ctrl *gomock.Controller) *MockITaskRepo {
	mock := &MockITaskRepo{ctrl: ctrl}
	mock.recorder = &MockITaskRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITaskRepo) EXPECT() *MockITaskRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockITaskRepo) Add(c *dto.AppContext, data *entity.Task) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", c, data)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockITaskRepoMockRecorder) Add(c, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockITaskRepo)(nil).Add), c, data)
}

// GetById mocks base method.
func (m *MockITaskRepo) GetById(c *dto.AppContext, id int64) (*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", c, id)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockITaskRepoMockRecorder) GetById(c, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockITaskRepo)(nil).GetById), c, id)
}

// GetList mocks base method.
func (m *MockITaskRepo) GetList(c *dto.AppContext) (*[]entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", c)
	ret0, _ := ret[0].(*[]entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockITaskRepoMockRecorder) GetList(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockITaskRepo)(nil).GetList), c)
}
