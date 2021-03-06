// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shirasudon/go-chat/chat (interfaces: Pubsub)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	event "github.com/shirasudon/go-chat/domain/event"
	reflect "reflect"
)

// MockPubsub is a mock of Pubsub interface
type MockPubsub struct {
	ctrl     *gomock.Controller
	recorder *MockPubsubMockRecorder
}

// MockPubsubMockRecorder is the mock recorder for MockPubsub
type MockPubsubMockRecorder struct {
	mock *MockPubsub
}

// NewMockPubsub creates a new mock instance
func NewMockPubsub(ctrl *gomock.Controller) *MockPubsub {
	mock := &MockPubsub{ctrl: ctrl}
	mock.recorder = &MockPubsubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPubsub) EXPECT() *MockPubsubMockRecorder {
	return m.recorder
}

// Pub mocks base method
func (m *MockPubsub) Pub(arg0 ...event.Event) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Pub", varargs...)
}

// Pub indicates an expected call of Pub
func (mr *MockPubsubMockRecorder) Pub(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pub", reflect.TypeOf((*MockPubsub)(nil).Pub), arg0...)
}

// Sub mocks base method
func (m *MockPubsub) Sub(arg0 ...event.Type) chan interface{} {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sub", varargs...)
	ret0, _ := ret[0].(chan interface{})
	return ret0
}

// Sub indicates an expected call of Sub
func (mr *MockPubsubMockRecorder) Sub(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sub", reflect.TypeOf((*MockPubsub)(nil).Sub), arg0...)
}
