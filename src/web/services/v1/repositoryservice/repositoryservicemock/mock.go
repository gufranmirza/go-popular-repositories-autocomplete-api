// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufranmirza/go-popular-repositories-autocomplete-api/web/services/v1/repositoryservice (interfaces: RepositoryService)

// Package repositoryservicemock is a generated GoMock package.
package repositoryservicemock

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockRepositoryService is a mock of RepositoryService interface
type MockRepositoryService struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryServiceMockRecorder
}

// MockRepositoryServiceMockRecorder is the mock recorder for MockRepositoryService
type MockRepositoryServiceMockRecorder struct {
	mock *MockRepositoryService
}

// NewMockRepositoryService creates a new mock instance
func NewMockRepositoryService(ctrl *gomock.Controller) *MockRepositoryService {
	mock := &MockRepositoryService{ctrl: ctrl}
	mock.recorder = &MockRepositoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositoryService) EXPECT() *MockRepositoryServiceMockRecorder {
	return m.recorder
}

// Search mocks base method
func (m *MockRepositoryService) Search(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Search", arg0, arg1)
}

// Search indicates an expected call of Search
func (mr *MockRepositoryServiceMockRecorder) Search(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepositoryService)(nil).Search), arg0, arg1)
}