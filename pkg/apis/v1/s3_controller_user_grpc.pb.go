// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// S3ControllerUserServiceClient is the client API for S3ControllerUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S3ControllerUserServiceClient interface {
	Get(ctx context.Context, in *S3ControllerUser, opts ...grpc.CallOption) (*S3ControllerUserWithCredentials, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*S3ControllerUserList, error)
	Create(ctx context.Context, in *S3ControllerUserCreate, opts ...grpc.CallOption) (*S3ControllerUserWithCredentials, error)
	Delete(ctx context.Context, in *S3ControllerUserDelete, opts ...grpc.CallOption) (*Empty, error)
	Update(ctx context.Context, in *S3ControllerUserUpdate, opts ...grpc.CallOption) (*S3ControllerUserWithCredentials, error)
}

type s3ControllerUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewS3ControllerUserServiceClient(cc grpc.ClientConnInterface) S3ControllerUserServiceClient {
	return &s3ControllerUserServiceClient{cc}
}

func (c *s3ControllerUserServiceClient) Get(ctx context.Context, in *S3ControllerUser, opts ...grpc.CallOption) (*S3ControllerUserWithCredentials, error) {
	out := new(S3ControllerUserWithCredentials)
	err := c.cc.Invoke(ctx, "/metalstack.io.s3.api.v1.S3ControllerUserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3ControllerUserServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*S3ControllerUserList, error) {
	out := new(S3ControllerUserList)
	err := c.cc.Invoke(ctx, "/metalstack.io.s3.api.v1.S3ControllerUserService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3ControllerUserServiceClient) Create(ctx context.Context, in *S3ControllerUserCreate, opts ...grpc.CallOption) (*S3ControllerUserWithCredentials, error) {
	out := new(S3ControllerUserWithCredentials)
	err := c.cc.Invoke(ctx, "/metalstack.io.s3.api.v1.S3ControllerUserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3ControllerUserServiceClient) Delete(ctx context.Context, in *S3ControllerUserDelete, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/metalstack.io.s3.api.v1.S3ControllerUserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3ControllerUserServiceClient) Update(ctx context.Context, in *S3ControllerUserUpdate, opts ...grpc.CallOption) (*S3ControllerUserWithCredentials, error) {
	out := new(S3ControllerUserWithCredentials)
	err := c.cc.Invoke(ctx, "/metalstack.io.s3.api.v1.S3ControllerUserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S3ControllerUserServiceServer is the server API for S3ControllerUserService service.
// All implementations should embed UnimplementedS3ControllerUserServiceServer
// for forward compatibility
type S3ControllerUserServiceServer interface {
	Get(context.Context, *S3ControllerUser) (*S3ControllerUserWithCredentials, error)
	List(context.Context, *Empty) (*S3ControllerUserList, error)
	Create(context.Context, *S3ControllerUserCreate) (*S3ControllerUserWithCredentials, error)
	Delete(context.Context, *S3ControllerUserDelete) (*Empty, error)
	Update(context.Context, *S3ControllerUserUpdate) (*S3ControllerUserWithCredentials, error)
}

// UnimplementedS3ControllerUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedS3ControllerUserServiceServer struct {
}

func (UnimplementedS3ControllerUserServiceServer) Get(context.Context, *S3ControllerUser) (*S3ControllerUserWithCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedS3ControllerUserServiceServer) List(context.Context, *Empty) (*S3ControllerUserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedS3ControllerUserServiceServer) Create(context.Context, *S3ControllerUserCreate) (*S3ControllerUserWithCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedS3ControllerUserServiceServer) Delete(context.Context, *S3ControllerUserDelete) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedS3ControllerUserServiceServer) Update(context.Context, *S3ControllerUserUpdate) (*S3ControllerUserWithCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// UnsafeS3ControllerUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S3ControllerUserServiceServer will
// result in compilation errors.
type UnsafeS3ControllerUserServiceServer interface {
	mustEmbedUnimplementedS3ControllerUserServiceServer()
}

func RegisterS3ControllerUserServiceServer(s grpc.ServiceRegistrar, srv S3ControllerUserServiceServer) {
	s.RegisterService(&S3ControllerUserService_ServiceDesc, srv)
}

func _S3ControllerUserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3ControllerUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3ControllerUserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metalstack.io.s3.api.v1.S3ControllerUserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3ControllerUserServiceServer).Get(ctx, req.(*S3ControllerUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3ControllerUserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3ControllerUserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metalstack.io.s3.api.v1.S3ControllerUserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3ControllerUserServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3ControllerUserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3ControllerUserCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3ControllerUserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metalstack.io.s3.api.v1.S3ControllerUserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3ControllerUserServiceServer).Create(ctx, req.(*S3ControllerUserCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3ControllerUserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3ControllerUserDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3ControllerUserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metalstack.io.s3.api.v1.S3ControllerUserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3ControllerUserServiceServer).Delete(ctx, req.(*S3ControllerUserDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3ControllerUserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3ControllerUserUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3ControllerUserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metalstack.io.s3.api.v1.S3ControllerUserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3ControllerUserServiceServer).Update(ctx, req.(*S3ControllerUserUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// S3ControllerUserService_ServiceDesc is the grpc.ServiceDesc for S3ControllerUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S3ControllerUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metalstack.io.s3.api.v1.S3ControllerUserService",
	HandlerType: (*S3ControllerUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _S3ControllerUserService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _S3ControllerUserService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _S3ControllerUserService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _S3ControllerUserService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _S3ControllerUserService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metalstack/io/s3/api/v1/s3_controller_user.proto",
}
