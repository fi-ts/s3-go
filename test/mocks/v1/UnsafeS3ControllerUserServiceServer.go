// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeS3ControllerUserServiceServer is an autogenerated mock type for the UnsafeS3ControllerUserServiceServer type
type UnsafeS3ControllerUserServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedS3ControllerUserServiceServer provides a mock function with given fields:
func (_m *UnsafeS3ControllerUserServiceServer) mustEmbedUnimplementedS3ControllerUserServiceServer() {
	_m.Called()
}

type NewUnsafeS3ControllerUserServiceServerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeS3ControllerUserServiceServer creates a new instance of UnsafeS3ControllerUserServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeS3ControllerUserServiceServer(t NewUnsafeS3ControllerUserServiceServerT) *UnsafeS3ControllerUserServiceServer {
	mock := &UnsafeS3ControllerUserServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
