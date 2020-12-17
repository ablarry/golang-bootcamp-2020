// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/repo.go

// Package repo is a generated GoMock package.
package repo

import (
	model "github.com/ablarry/golang-bootcamp-2020/pkg/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// FindOne mocks base method
func (m *MockRepo) FindOne(id int64) (*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", id)
	ret0, _ := ret[0].(*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockRepoMockRecorder) FindOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockRepo)(nil).FindOne), id)
}

// Find mocks base method
func (m *MockRepo) Find(paginator, limit int64) ([]*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", paginator, limit)
	ret0, _ := ret[0].([]*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockRepoMockRecorder) Find(paginator, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepo)(nil).Find), paginator, limit)
}

// Create mocks base method
func (m *MockRepo) Create(p *model.Pokemon) (*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", p)
	ret0, _ := ret[0].(*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRepoMockRecorder) Create(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepo)(nil).Create), p)
}

// Delete mocks base method
func (m *MockRepo) Delete(id int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockRepoMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepo)(nil).Delete), id)
}