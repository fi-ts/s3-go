// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeS3PartitionServiceServer is an autogenerated mock type for the UnsafeS3PartitionServiceServer type
type UnsafeS3PartitionServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedS3PartitionServiceServer provides a mock function with given fields:
func (_m *UnsafeS3PartitionServiceServer) mustEmbedUnimplementedS3PartitionServiceServer() {
	_m.Called()
}

type NewUnsafeS3PartitionServiceServerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeS3PartitionServiceServer creates a new instance of UnsafeS3PartitionServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeS3PartitionServiceServer(t NewUnsafeS3PartitionServiceServerT) *UnsafeS3PartitionServiceServer {
	mock := &UnsafeS3PartitionServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
