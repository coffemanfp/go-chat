// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shirasudon/go-chat/domain (interfaces: Repositories)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/shirasudon/go-chat/domain"
	reflect "reflect"
)

// MockRepositories is a mock of Repositories interface
type MockRepositories struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoriesMockRecorder
}

// MockRepositoriesMockRecorder is the mock recorder for MockRepositories
type MockRepositoriesMockRecorder struct {
	mock *MockRepositories
}

// NewMockRepositories creates a new mock instance
func NewMockRepositories(ctrl *gomock.Controller) *MockRepositories {
	mock := &MockRepositories{ctrl: ctrl}
	mock.recorder = &MockRepositoriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositories) EXPECT() *MockRepositoriesMockRecorder {
	return m.recorder
}

// Events mocks base method
func (m *MockRepositories) Events() domain.EventRepository {
	ret := m.ctrl.Call(m, "Events")
	ret0, _ := ret[0].(domain.EventRepository)
	return ret0
}

// Events indicates an expected call of Events
func (mr *MockRepositoriesMockRecorder) Events() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockRepositories)(nil).Events))
}

// Messages mocks base method
func (m *MockRepositories) Messages() domain.MessageRepository {
	ret := m.ctrl.Call(m, "Messages")
	ret0, _ := ret[0].(domain.MessageRepository)
	return ret0
}

// Messages indicates an expected call of Messages
func (mr *MockRepositoriesMockRecorder) Messages() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Messages", reflect.TypeOf((*MockRepositories)(nil).Messages))
}

// Rooms mocks base method
func (m *MockRepositories) Rooms() domain.RoomRepository {
	ret := m.ctrl.Call(m, "Rooms")
	ret0, _ := ret[0].(domain.RoomRepository)
	return ret0
}

// Rooms indicates an expected call of Rooms
func (mr *MockRepositoriesMockRecorder) Rooms() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rooms", reflect.TypeOf((*MockRepositories)(nil).Rooms))
}

// Users mocks base method
func (m *MockRepositories) Users() domain.UserRepository {
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(domain.UserRepository)
	return ret0
}

// Users indicates an expected call of Users
func (mr *MockRepositoriesMockRecorder) Users() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockRepositories)(nil).Users))
}
