// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: metalstack/io/s3/api/v1/s3_user.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type S3User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant      string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project     string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	MaxBuckets  int64  `protobuf:"varint,5,opt,name=max_buckets,json=maxBuckets,proto3" json:"max_buckets,omitempty"`
	DisplayName string `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Partition   string `protobuf:"bytes,7,opt,name=partition,proto3" json:"partition,omitempty"`
	Endpoint    string `protobuf:"bytes,8,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *S3User) Reset() {
	*x = S3User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3User) ProtoMessage() {}

func (x *S3User) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3User.ProtoReflect.Descriptor instead.
func (*S3User) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_user_proto_rawDescGZIP(), []int{0}
}

func (x *S3User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3User) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3User) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *S3User) GetMaxBuckets() int64 {
	if x != nil {
		return x.MaxBuckets
	}
	return 0
}

func (x *S3User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *S3User) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *S3User) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type S3UserWithCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant      string       `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project     string       `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Email       string       `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	MaxBuckets  int64        `protobuf:"varint,5,opt,name=max_buckets,json=maxBuckets,proto3" json:"max_buckets,omitempty"`
	DisplayName string       `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Partition   string       `protobuf:"bytes,9,opt,name=partition,proto3" json:"partition,omitempty"`
	Endpoint    string       `protobuf:"bytes,10,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Keys        []*S3UserKey `protobuf:"bytes,11,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *S3UserWithCredentials) Reset() {
	*x = S3UserWithCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3UserWithCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3UserWithCredentials) ProtoMessage() {}

func (x *S3UserWithCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3UserWithCredentials.ProtoReflect.Descriptor instead.
func (*S3UserWithCredentials) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_user_proto_rawDescGZIP(), []int{1}
}

func (x *S3UserWithCredentials) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3UserWithCredentials) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3UserWithCredentials) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3UserWithCredentials) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *S3UserWithCredentials) GetMaxBuckets() int64 {
	if x != nil {
		return x.MaxBuckets
	}
	return 0
}

func (x *S3UserWithCredentials) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *S3UserWithCredentials) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *S3UserWithCredentials) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *S3UserWithCredentials) GetKeys() []*S3UserKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

type S3UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*S3User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *S3UserList) Reset() {
	*x = S3UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3UserList) ProtoMessage() {}

func (x *S3UserList) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3UserList.ProtoReflect.Descriptor instead.
func (*S3UserList) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_user_proto_rawDescGZIP(), []int{2}
}

func (x *S3UserList) GetUsers() []*S3User {
	if x != nil {
		return x.Users
	}
	return nil
}

type S3UserDescribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant    string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project   string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Partition string `protobuf:"bytes,4,opt,name=partition,proto3" json:"partition,omitempty"`
}

func (x *S3UserDescribeRequest) Reset() {
	*x = S3UserDescribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3UserDescribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3UserDescribeRequest) ProtoMessage() {}

func (x *S3UserDescribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3UserDescribeRequest.ProtoReflect.Descriptor instead.
func (*S3UserDescribeRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_user_proto_rawDescGZIP(), []int{3}
}

func (x *S3UserDescribeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3UserDescribeRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3UserDescribeRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3UserDescribeRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

type S3UserDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant    string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project   string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Partition string `protobuf:"bytes,4,opt,name=partition,proto3" json:"partition,omitempty"`
	Force     bool   `protobuf:"varint,5,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *S3UserDeleteRequest) Reset() {
	*x = S3UserDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3UserDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3UserDeleteRequest) ProtoMessage() {}

func (x *S3UserDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3UserDeleteRequest.ProtoReflect.Descriptor instead.
func (*S3UserDeleteRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_user_proto_rawDescGZIP(), []int{4}
}

func (x *S3UserDeleteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3UserDeleteRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3UserDeleteRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3UserDeleteRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *S3UserDeleteRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type S3UserCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant      string     `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project     string     `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Email       string     `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	MaxBuckets  int64      `protobuf:"varint,5,opt,name=max_buckets,json=maxBuckets,proto3" json:"max_buckets,omitempty"`
	DisplayName string     `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Partition   string     `protobuf:"bytes,7,opt,name=partition,proto3" json:"partition,omitempty"`
	Key         *S3UserKey `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *S3UserCreateRequest) Reset() {
	*x = S3UserCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3UserCreateRequest) ProtoMessage() {}

func (x *S3UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3UserCreateRequest.ProtoReflect.Descriptor instead.
func (*S3UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_user_proto_rawDescGZIP(), []int{5}
}

func (x *S3UserCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3UserCreateRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3UserCreateRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3UserCreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *S3UserCreateRequest) GetMaxBuckets() int64 {
	if x != nil {
		return x.MaxBuckets
	}
	return 0
}

func (x *S3UserCreateRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *S3UserCreateRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *S3UserCreateRequest) GetKey() *S3UserKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type S3UserUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant           string       `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project          string       `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	RemoveAccessKeys []string     `protobuf:"bytes,4,rep,name=remove_access_keys,json=removeAccessKeys,proto3" json:"remove_access_keys,omitempty"`
	AddKeys          []*S3UserKey `protobuf:"bytes,5,rep,name=add_keys,json=addKeys,proto3" json:"add_keys,omitempty"`
	Partition        string       `protobuf:"bytes,6,opt,name=partition,proto3" json:"partition,omitempty"`
}

func (x *S3UserUpdateRequest) Reset() {
	*x = S3UserUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3UserUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3UserUpdateRequest) ProtoMessage() {}

func (x *S3UserUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3UserUpdateRequest.ProtoReflect.Descriptor instead.
func (*S3UserUpdateRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_user_proto_rawDescGZIP(), []int{6}
}

func (x *S3UserUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3UserUpdateRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3UserUpdateRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3UserUpdateRequest) GetRemoveAccessKeys() []string {
	if x != nil {
		return x.RemoveAccessKeys
	}
	return nil
}

func (x *S3UserUpdateRequest) GetAddKeys() []*S3UserKey {
	if x != nil {
		return x.AddKeys
	}
	return nil
}

func (x *S3UserUpdateRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

type S3UserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partition string `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
}

func (x *S3UserListRequest) Reset() {
	*x = S3UserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3UserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3UserListRequest) ProtoMessage() {}

func (x *S3UserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3UserListRequest.ProtoReflect.Descriptor instead.
func (*S3UserListRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_user_proto_rawDescGZIP(), []int{7}
}

func (x *S3UserListRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

var File_metalstack_io_s3_api_v1_s3_user_proto protoreflect.FileDescriptor

var file_metalstack_io_s3_api_v1_s3_user_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f,
	0x73, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x33, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x1a, 0x24, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f,
	0x73, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x06, 0x53, 0x33, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xa9, 0x02, 0x0a, 0x15,
	0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x36, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65,
	0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x43, 0x0a, 0x0a, 0x53, 0x33, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x33, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x7b, 0x0a, 0x15,
	0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8f, 0x01, 0x0a, 0x13, 0x53, 0x33,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x89, 0x02, 0x0a, 0x13,
	0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73,
	0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xe6, 0x01, 0x0a, 0x13, 0x53, 0x33, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b,
	0x65, 0x79, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x61, 0x64, 0x64, 0x4b, 0x65,
	0x79, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x31, 0x0a, 0x11, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x32, 0xf8, 0x03, 0x0a, 0x0d, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2e, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74,
	0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x57, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x33, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f,
	0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x66, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x2c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e,
	0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69,
	0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x57, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x2c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f,
	0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73,
	0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x57,
	0x69, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x0d,
	0x5a, 0x0b, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metalstack_io_s3_api_v1_s3_user_proto_rawDescOnce sync.Once
	file_metalstack_io_s3_api_v1_s3_user_proto_rawDescData = file_metalstack_io_s3_api_v1_s3_user_proto_rawDesc
)

func file_metalstack_io_s3_api_v1_s3_user_proto_rawDescGZIP() []byte {
	file_metalstack_io_s3_api_v1_s3_user_proto_rawDescOnce.Do(func() {
		file_metalstack_io_s3_api_v1_s3_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_metalstack_io_s3_api_v1_s3_user_proto_rawDescData)
	})
	return file_metalstack_io_s3_api_v1_s3_user_proto_rawDescData
}

var file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_metalstack_io_s3_api_v1_s3_user_proto_goTypes = []interface{}{
	(*S3User)(nil),                // 0: metalstack.io.s3.api.v1.S3User
	(*S3UserWithCredentials)(nil), // 1: metalstack.io.s3.api.v1.S3UserWithCredentials
	(*S3UserList)(nil),            // 2: metalstack.io.s3.api.v1.S3UserList
	(*S3UserDescribeRequest)(nil), // 3: metalstack.io.s3.api.v1.S3UserDescribeRequest
	(*S3UserDeleteRequest)(nil),   // 4: metalstack.io.s3.api.v1.S3UserDeleteRequest
	(*S3UserCreateRequest)(nil),   // 5: metalstack.io.s3.api.v1.S3UserCreateRequest
	(*S3UserUpdateRequest)(nil),   // 6: metalstack.io.s3.api.v1.S3UserUpdateRequest
	(*S3UserListRequest)(nil),     // 7: metalstack.io.s3.api.v1.S3UserListRequest
	(*S3UserKey)(nil),             // 8: metalstack.io.s3.api.v1.S3UserKey
}
var file_metalstack_io_s3_api_v1_s3_user_proto_depIdxs = []int32{
	8, // 0: metalstack.io.s3.api.v1.S3UserWithCredentials.keys:type_name -> metalstack.io.s3.api.v1.S3UserKey
	0, // 1: metalstack.io.s3.api.v1.S3UserList.users:type_name -> metalstack.io.s3.api.v1.S3User
	8, // 2: metalstack.io.s3.api.v1.S3UserCreateRequest.key:type_name -> metalstack.io.s3.api.v1.S3UserKey
	8, // 3: metalstack.io.s3.api.v1.S3UserUpdateRequest.add_keys:type_name -> metalstack.io.s3.api.v1.S3UserKey
	3, // 4: metalstack.io.s3.api.v1.S3UserService.Get:input_type -> metalstack.io.s3.api.v1.S3UserDescribeRequest
	7, // 5: metalstack.io.s3.api.v1.S3UserService.List:input_type -> metalstack.io.s3.api.v1.S3UserListRequest
	5, // 6: metalstack.io.s3.api.v1.S3UserService.Create:input_type -> metalstack.io.s3.api.v1.S3UserCreateRequest
	4, // 7: metalstack.io.s3.api.v1.S3UserService.Delete:input_type -> metalstack.io.s3.api.v1.S3UserDeleteRequest
	6, // 8: metalstack.io.s3.api.v1.S3UserService.Update:input_type -> metalstack.io.s3.api.v1.S3UserUpdateRequest
	1, // 9: metalstack.io.s3.api.v1.S3UserService.Get:output_type -> metalstack.io.s3.api.v1.S3UserWithCredentials
	2, // 10: metalstack.io.s3.api.v1.S3UserService.List:output_type -> metalstack.io.s3.api.v1.S3UserList
	1, // 11: metalstack.io.s3.api.v1.S3UserService.Create:output_type -> metalstack.io.s3.api.v1.S3UserWithCredentials
	0, // 12: metalstack.io.s3.api.v1.S3UserService.Delete:output_type -> metalstack.io.s3.api.v1.S3User
	1, // 13: metalstack.io.s3.api.v1.S3UserService.Update:output_type -> metalstack.io.s3.api.v1.S3UserWithCredentials
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_metalstack_io_s3_api_v1_s3_user_proto_init() }
func file_metalstack_io_s3_api_v1_s3_user_proto_init() {
	if File_metalstack_io_s3_api_v1_s3_user_proto != nil {
		return
	}
	file_metalstack_io_s3_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3UserWithCredentials); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3UserList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3UserDescribeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3UserDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3UserCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3UserUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3UserListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metalstack_io_s3_api_v1_s3_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_io_s3_api_v1_s3_user_proto_goTypes,
		DependencyIndexes: file_metalstack_io_s3_api_v1_s3_user_proto_depIdxs,
		MessageInfos:      file_metalstack_io_s3_api_v1_s3_user_proto_msgTypes,
	}.Build()
	File_metalstack_io_s3_api_v1_s3_user_proto = out.File
	file_metalstack_io_s3_api_v1_s3_user_proto_rawDesc = nil
	file_metalstack_io_s3_api_v1_s3_user_proto_goTypes = nil
	file_metalstack_io_s3_api_v1_s3_user_proto_depIdxs = nil
}
