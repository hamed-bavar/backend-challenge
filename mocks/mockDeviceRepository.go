// Code generated by MockGen. DO NOT EDIT.
// Source: challenge/domain (interfaces: DeviceRepository)

// Package domain is a generated GoMock package.
package domain

import (
	domain "challenge/domain"
	errors "challenge/lib/errors"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDeviceRepository is a mock of DeviceRepository interface.
type MockDeviceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceRepositoryMockRecorder
}

// MockDeviceRepositoryMockRecorder is the mock recorder for MockDeviceRepository.
type MockDeviceRepositoryMockRecorder struct {
	mock *MockDeviceRepository
}

// NewMockDeviceRepository creates a new mock instance.
func NewMockDeviceRepository(ctrl *gomock.Controller) *MockDeviceRepository {
	mock := &MockDeviceRepository{ctrl: ctrl}
	mock.recorder = &MockDeviceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceRepository) EXPECT() *MockDeviceRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDeviceRepository) Create(arg0 *domain.Device) (*domain.Device, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*domain.Device)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDeviceRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDeviceRepository)(nil).Create), arg0)
}

// FindById mocks base method.
func (m *MockDeviceRepository) FindById(arg0 string) (*domain.Device, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*domain.Device)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockDeviceRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockDeviceRepository)(nil).FindById), arg0)
}