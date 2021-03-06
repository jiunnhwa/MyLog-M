// Code generated by MockGen. DO NOT EDIT.
// Source: ./mylog.go

// Package mock is a generated GoMock package.
package mock

import (
	domain "MyLog-M/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockrepository is a mock of repository interface.
type Mockrepository struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryMockRecorder
}

// MockrepositoryMockRecorder is the mock recorder for Mockrepository.
type MockrepositoryMockRecorder struct {
	mock *Mockrepository
}

// NewMockrepository creates a new mock instance.
func NewMockrepository(ctrl *gomock.Controller) *Mockrepository {
	mock := &Mockrepository{ctrl: ctrl}
	mock.recorder = &MockrepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepository) EXPECT() *MockrepositoryMockRecorder {
	return m.recorder
}

// Tail mocks base method.
func (m *Mockrepository) Tail(limit int64) (*[]domain.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tail", limit)
	ret0, _ := ret[0].(*[]domain.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tail indicates an expected call of Tail.
func (mr *MockrepositoryMockRecorder) Tail(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tail", reflect.TypeOf((*Mockrepository)(nil).Tail), limit)
}
