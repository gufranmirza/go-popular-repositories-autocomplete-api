// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufranmirza/go-popular-repositories-autocomplete-api/dal/repositorydal (interfaces: RepositoryDal)

// Package middlewaresmock is a generated GoMock package.
package middlewaresmock

import (
	gomock "github.com/golang/mock/gomock"
	dbmodels "github.com/gufranmirza/go-popular-repositories-autocomplete-api/database/dbmodels"
	reflect "reflect"
)

// MockRepositoryDal is a mock of RepositoryDal interface
type MockRepositoryDal struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryDalMockRecorder
}

// MockRepositoryDalMockRecorder is the mock recorder for MockRepositoryDal
type MockRepositoryDalMockRecorder struct {
	mock *MockRepositoryDal
}

// NewMockRepositoryDal creates a new mock instance
func NewMockRepositoryDal(ctrl *gomock.Controller) *MockRepositoryDal {
	mock := &MockRepositoryDal{ctrl: ctrl}
	mock.recorder = &MockRepositoryDalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositoryDal) EXPECT() *MockRepositoryDalMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRepositoryDal) Create(arg0 *dbmodels.Repository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockRepositoryDalMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositoryDal)(nil).Create), arg0)
}

// Search mocks base method
func (m *MockRepositoryDal) Search(arg0 string) ([]dbmodels.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].([]dbmodels.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryDalMockRecorder) Search(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepositoryDal)(nil).Search), arg0)
}
