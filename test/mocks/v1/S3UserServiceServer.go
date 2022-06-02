// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	v1 "github.com/fi-ts/s3-go/pkg/apis/v1"
	mock "github.com/stretchr/testify/mock"
)

// S3UserServiceServer is an autogenerated mock type for the S3UserServiceServer type
type S3UserServiceServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *S3UserServiceServer) Create(_a0 context.Context, _a1 *v1.S3UserCreateRequest) (*v1.S3UserWithCredentials, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.S3UserWithCredentials
	if rf, ok := ret.Get(0).(func(context.Context, *v1.S3UserCreateRequest) *v1.S3UserWithCredentials); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.S3UserWithCredentials)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.S3UserCreateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *S3UserServiceServer) Delete(_a0 context.Context, _a1 *v1.S3UserDeleteRequest) (*v1.S3User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.S3User
	if rf, ok := ret.Get(0).(func(context.Context, *v1.S3UserDeleteRequest) *v1.S3User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.S3User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.S3UserDeleteRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *S3UserServiceServer) Get(_a0 context.Context, _a1 *v1.S3UserDescribeRequest) (*v1.S3UserWithCredentials, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.S3UserWithCredentials
	if rf, ok := ret.Get(0).(func(context.Context, *v1.S3UserDescribeRequest) *v1.S3UserWithCredentials); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.S3UserWithCredentials)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.S3UserDescribeRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *S3UserServiceServer) List(_a0 context.Context, _a1 *v1.S3UserListRequest) (*v1.S3UserList, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.S3UserList
	if rf, ok := ret.Get(0).(func(context.Context, *v1.S3UserListRequest) *v1.S3UserList); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.S3UserList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.S3UserListRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *S3UserServiceServer) Update(_a0 context.Context, _a1 *v1.S3UserUpdateRequest) (*v1.S3UserWithCredentials, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.S3UserWithCredentials
	if rf, ok := ret.Get(0).(func(context.Context, *v1.S3UserUpdateRequest) *v1.S3UserWithCredentials); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.S3UserWithCredentials)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.S3UserUpdateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewS3UserServiceServerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewS3UserServiceServer creates a new instance of S3UserServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewS3UserServiceServer(t NewS3UserServiceServerT) *S3UserServiceServer {
	mock := &S3UserServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
