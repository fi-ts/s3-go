// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeS3UserServiceServer is an autogenerated mock type for the UnsafeS3UserServiceServer type
type UnsafeS3UserServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedS3UserServiceServer provides a mock function with given fields:
func (_m *UnsafeS3UserServiceServer) mustEmbedUnimplementedS3UserServiceServer() {
	_m.Called()
}

// NewUnsafeS3UserServiceServer creates a new instance of UnsafeS3UserServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeS3UserServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeS3UserServiceServer {
	mock := &UnsafeS3UserServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
