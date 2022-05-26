// Code generated by MockGen. DO NOT EDIT.
// Source: client/storage/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	model "advdiploma/client/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddSecret mocks base method.
func (m *MockStorage) AddSecret(v model.Secret) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSecret", v)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSecret indicates an expected call of AddSecret.
func (mr *MockStorageMockRecorder) AddSecret(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecret", reflect.TypeOf((*MockStorage)(nil).AddSecret), v)
}

// Close mocks base method.
func (m *MockStorage) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}

// GetMetaList mocks base method.
func (m *MockStorage) GetMetaList() ([]model.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetaList")
	ret0, _ := ret[0].([]model.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetaList indicates an expected call of GetMetaList.
func (mr *MockStorageMockRecorder) GetMetaList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaList", reflect.TypeOf((*MockStorage)(nil).GetMetaList))
}

// GetSecret mocks base method.
func (m *MockStorage) GetSecret(id int64) (model.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", id)
	ret0, _ := ret[0].(model.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockStorageMockRecorder) GetSecret(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockStorage)(nil).GetSecret), id)
}

// UpdateSecret mocks base method.
func (m *MockStorage) UpdateSecret(v model.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockStorageMockRecorder) UpdateSecret(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockStorage)(nil).UpdateSecret), v)
}
