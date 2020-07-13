// Code generated by MockGen. DO NOT EDIT.
// Source: store/memcache.go

// Package mocks is a generated GoMock package.
package mocks

import (
	memcache "github.com/bradfitz/gomemcache/memcache"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMemcacheClientInterface is a mock of MemcacheClientInterface interface
type MockMemcacheClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMemcacheClientInterfaceMockRecorder
}

// MockMemcacheClientInterfaceMockRecorder is the mock recorder for MockMemcacheClientInterface
type MockMemcacheClientInterfaceMockRecorder struct {
	mock *MockMemcacheClientInterface
}

// NewMockMemcacheClientInterface creates a new mock instance
func NewMockMemcacheClientInterface(ctrl *gomock.Controller) *MockMemcacheClientInterface {
	mock := &MockMemcacheClientInterface{ctrl: ctrl}
	mock.recorder = &MockMemcacheClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMemcacheClientInterface) EXPECT() *MockMemcacheClientInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockMemcacheClientInterface) Get(key string) (*memcache.Item, error) {
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*memcache.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMemcacheClientInterfaceMockRecorder) Get(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMemcacheClientInterface)(nil).Get), key)
}

// Set mocks base method
func (m *MockMemcacheClientInterface) Set(item *memcache.Item) error {
	ret := m.ctrl.Call(m, "Set", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockMemcacheClientInterfaceMockRecorder) Set(item interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockMemcacheClientInterface)(nil).Set), item)
}

// Delete mocks base method
func (m *MockMemcacheClientInterface) Delete(item string) error {
	ret := m.ctrl.Call(m, "Delete", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockMemcacheClientInterfaceMockRecorder) Delete(item interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMemcacheClientInterface)(nil).Delete), item)
}

// FlushAll mocks base method
func (m *MockMemcacheClientInterface) FlushAll() error {
	ret := m.ctrl.Call(m, "FlushAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushAll indicates an expected call of FlushAll
func (mr *MockMemcacheClientInterfaceMockRecorder) FlushAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAll", reflect.TypeOf((*MockMemcacheClientInterface)(nil).FlushAll))
}
