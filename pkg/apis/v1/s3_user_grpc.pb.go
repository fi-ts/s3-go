// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: metalstack/io/s3/api/v1/s3_user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	S3UserService_Get_FullMethodName    = "/metalstack.io.s3.api.v1.S3UserService/Get"
	S3UserService_List_FullMethodName   = "/metalstack.io.s3.api.v1.S3UserService/List"
	S3UserService_Create_FullMethodName = "/metalstack.io.s3.api.v1.S3UserService/Create"
	S3UserService_Delete_FullMethodName = "/metalstack.io.s3.api.v1.S3UserService/Delete"
	S3UserService_Update_FullMethodName = "/metalstack.io.s3.api.v1.S3UserService/Update"
)

// S3UserServiceClient is the client API for S3UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S3UserServiceClient interface {
	Get(ctx context.Context, in *S3UserDescribeRequest, opts ...grpc.CallOption) (*S3UserWithCredentials, error)
	List(ctx context.Context, in *S3UserListRequest, opts ...grpc.CallOption) (*S3UserList, error)
	Create(ctx context.Context, in *S3UserCreateRequest, opts ...grpc.CallOption) (*S3UserWithCredentials, error)
	Delete(ctx context.Context, in *S3UserDeleteRequest, opts ...grpc.CallOption) (*S3User, error)
	Update(ctx context.Context, in *S3UserUpdateRequest, opts ...grpc.CallOption) (*S3UserWithCredentials, error)
}

type s3UserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewS3UserServiceClient(cc grpc.ClientConnInterface) S3UserServiceClient {
	return &s3UserServiceClient{cc}
}

func (c *s3UserServiceClient) Get(ctx context.Context, in *S3UserDescribeRequest, opts ...grpc.CallOption) (*S3UserWithCredentials, error) {
	out := new(S3UserWithCredentials)
	err := c.cc.Invoke(ctx, S3UserService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3UserServiceClient) List(ctx context.Context, in *S3UserListRequest, opts ...grpc.CallOption) (*S3UserList, error) {
	out := new(S3UserList)
	err := c.cc.Invoke(ctx, S3UserService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3UserServiceClient) Create(ctx context.Context, in *S3UserCreateRequest, opts ...grpc.CallOption) (*S3UserWithCredentials, error) {
	out := new(S3UserWithCredentials)
	err := c.cc.Invoke(ctx, S3UserService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3UserServiceClient) Delete(ctx context.Context, in *S3UserDeleteRequest, opts ...grpc.CallOption) (*S3User, error) {
	out := new(S3User)
	err := c.cc.Invoke(ctx, S3UserService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3UserServiceClient) Update(ctx context.Context, in *S3UserUpdateRequest, opts ...grpc.CallOption) (*S3UserWithCredentials, error) {
	out := new(S3UserWithCredentials)
	err := c.cc.Invoke(ctx, S3UserService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S3UserServiceServer is the server API for S3UserService service.
// All implementations should embed UnimplementedS3UserServiceServer
// for forward compatibility
type S3UserServiceServer interface {
	Get(context.Context, *S3UserDescribeRequest) (*S3UserWithCredentials, error)
	List(context.Context, *S3UserListRequest) (*S3UserList, error)
	Create(context.Context, *S3UserCreateRequest) (*S3UserWithCredentials, error)
	Delete(context.Context, *S3UserDeleteRequest) (*S3User, error)
	Update(context.Context, *S3UserUpdateRequest) (*S3UserWithCredentials, error)
}

// UnimplementedS3UserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedS3UserServiceServer struct {
}

func (UnimplementedS3UserServiceServer) Get(context.Context, *S3UserDescribeRequest) (*S3UserWithCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedS3UserServiceServer) List(context.Context, *S3UserListRequest) (*S3UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedS3UserServiceServer) Create(context.Context, *S3UserCreateRequest) (*S3UserWithCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedS3UserServiceServer) Delete(context.Context, *S3UserDeleteRequest) (*S3User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedS3UserServiceServer) Update(context.Context, *S3UserUpdateRequest) (*S3UserWithCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// UnsafeS3UserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S3UserServiceServer will
// result in compilation errors.
type UnsafeS3UserServiceServer interface {
	mustEmbedUnimplementedS3UserServiceServer()
}

func RegisterS3UserServiceServer(s grpc.ServiceRegistrar, srv S3UserServiceServer) {
	s.RegisterService(&S3UserService_ServiceDesc, srv)
}

func _S3UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3UserDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3UserService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3UserServiceServer).Get(ctx, req.(*S3UserDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3UserService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3UserServiceServer).List(ctx, req.(*S3UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3UserService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3UserServiceServer).Create(ctx, req.(*S3UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3UserService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3UserServiceServer).Delete(ctx, req.(*S3UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3UserService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3UserServiceServer).Update(ctx, req.(*S3UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// S3UserService_ServiceDesc is the grpc.ServiceDesc for S3UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S3UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metalstack.io.s3.api.v1.S3UserService",
	HandlerType: (*S3UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _S3UserService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _S3UserService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _S3UserService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _S3UserService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _S3UserService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metalstack/io/s3/api/v1/s3_user.proto",
}
