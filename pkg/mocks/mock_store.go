// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bryanl/clientkube/pkg/cluster (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	"github.com/bryanl/clientkube/pkg/cluster"
	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockStore) Add(arg0 schema.GroupVersionResource, arg1 runtime.Object) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0, arg1)
}

// Add indicates an expected call of Add
func (mr *MockStoreMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStore)(nil).Add), arg0, arg1)
}

// Delete mocks base method
func (m *MockStore) Delete(arg0 schema.GroupVersionResource, arg1 runtime.Object) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", arg0, arg1)
}

// Delete indicates an expected call of Delete
func (mr *MockStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), arg0, arg1)
}

// List mocks base method
func (m *MockStore) List(arg0 schema.GroupVersionResource, arg1 cluster.ListOptions) (*unstructured.UnstructuredList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.UnstructuredList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStore)(nil).List), arg0, arg1)
}

// Update mocks base method
func (m *MockStore) Update(arg0 schema.GroupVersionResource, arg1 runtime.Object) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", arg0, arg1)
}

// Update indicates an expected call of Update
func (mr *MockStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), arg0, arg1)
}

// Watch mocks base method
func (m *MockStore) Watch(arg0 schema.GroupVersionResource, arg1 cluster.ListOptions) (cluster.Watch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(cluster.Watch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockStoreMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockStore)(nil).Watch), arg0, arg1)
}
