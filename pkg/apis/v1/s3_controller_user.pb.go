// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: metalstack/io/s3/api/v1/s3_controller_user.proto

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

type S3ControllerUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant      string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project     string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	MaxBuckets  int64  `protobuf:"varint,5,opt,name=max_buckets,json=maxBuckets,proto3" json:"max_buckets,omitempty"`
	DisplayName string `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *S3ControllerUser) Reset() {
	*x = S3ControllerUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ControllerUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ControllerUser) ProtoMessage() {}

func (x *S3ControllerUser) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ControllerUser.ProtoReflect.Descriptor instead.
func (*S3ControllerUser) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescGZIP(), []int{0}
}

func (x *S3ControllerUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3ControllerUser) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3ControllerUser) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3ControllerUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *S3ControllerUser) GetMaxBuckets() int64 {
	if x != nil {
		return x.MaxBuckets
	}
	return 0
}

func (x *S3ControllerUser) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type S3ControllerUserWithCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant      string       `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project     string       `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Email       string       `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	MaxBuckets  int64        `protobuf:"varint,5,opt,name=max_buckets,json=maxBuckets,proto3" json:"max_buckets,omitempty"`
	DisplayName string       `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Keys        []*S3UserKey `protobuf:"bytes,9,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *S3ControllerUserWithCredentials) Reset() {
	*x = S3ControllerUserWithCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ControllerUserWithCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ControllerUserWithCredentials) ProtoMessage() {}

func (x *S3ControllerUserWithCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ControllerUserWithCredentials.ProtoReflect.Descriptor instead.
func (*S3ControllerUserWithCredentials) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescGZIP(), []int{1}
}

func (x *S3ControllerUserWithCredentials) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3ControllerUserWithCredentials) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3ControllerUserWithCredentials) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3ControllerUserWithCredentials) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *S3ControllerUserWithCredentials) GetMaxBuckets() int64 {
	if x != nil {
		return x.MaxBuckets
	}
	return 0
}

func (x *S3ControllerUserWithCredentials) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *S3ControllerUserWithCredentials) GetKeys() []*S3UserKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

type S3ControllerUserCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant      string     `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project     string     `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Email       string     `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	MaxBuckets  int64      `protobuf:"varint,5,opt,name=max_buckets,json=maxBuckets,proto3" json:"max_buckets,omitempty"`
	DisplayName string     `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Key         *S3UserKey `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *S3ControllerUserCreate) Reset() {
	*x = S3ControllerUserCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ControllerUserCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ControllerUserCreate) ProtoMessage() {}

func (x *S3ControllerUserCreate) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ControllerUserCreate.ProtoReflect.Descriptor instead.
func (*S3ControllerUserCreate) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescGZIP(), []int{2}
}

func (x *S3ControllerUserCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3ControllerUserCreate) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3ControllerUserCreate) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3ControllerUserCreate) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *S3ControllerUserCreate) GetMaxBuckets() int64 {
	if x != nil {
		return x.MaxBuckets
	}
	return 0
}

func (x *S3ControllerUserCreate) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *S3ControllerUserCreate) GetKey() *S3UserKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type S3ControllerUserUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant           string       `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project          string       `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	RemoveAccessKeys []string     `protobuf:"bytes,4,rep,name=remove_access_keys,json=removeAccessKeys,proto3" json:"remove_access_keys,omitempty"`
	AddKeys          []*S3UserKey `protobuf:"bytes,5,rep,name=add_keys,json=addKeys,proto3" json:"add_keys,omitempty"`
}

func (x *S3ControllerUserUpdate) Reset() {
	*x = S3ControllerUserUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ControllerUserUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ControllerUserUpdate) ProtoMessage() {}

func (x *S3ControllerUserUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ControllerUserUpdate.ProtoReflect.Descriptor instead.
func (*S3ControllerUserUpdate) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescGZIP(), []int{3}
}

func (x *S3ControllerUserUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3ControllerUserUpdate) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3ControllerUserUpdate) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3ControllerUserUpdate) GetRemoveAccessKeys() []string {
	if x != nil {
		return x.RemoveAccessKeys
	}
	return nil
}

func (x *S3ControllerUserUpdate) GetAddKeys() []*S3UserKey {
	if x != nil {
		return x.AddKeys
	}
	return nil
}

type S3ControllerUserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*S3ControllerUser `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *S3ControllerUserList) Reset() {
	*x = S3ControllerUserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ControllerUserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ControllerUserList) ProtoMessage() {}

func (x *S3ControllerUserList) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ControllerUserList.ProtoReflect.Descriptor instead.
func (*S3ControllerUserList) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescGZIP(), []int{4}
}

func (x *S3ControllerUserList) GetUsers() []*S3ControllerUser {
	if x != nil {
		return x.Users
	}
	return nil
}

type S3ControllerUserDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant  string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Project string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Force   bool   `protobuf:"varint,4,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *S3ControllerUserDelete) Reset() {
	*x = S3ControllerUserDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ControllerUserDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ControllerUserDelete) ProtoMessage() {}

func (x *S3ControllerUserDelete) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ControllerUserDelete.ProtoReflect.Descriptor instead.
func (*S3ControllerUserDelete) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescGZIP(), []int{5}
}

func (x *S3ControllerUserDelete) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3ControllerUserDelete) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *S3ControllerUserDelete) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *S3ControllerUserDelete) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

var File_metalstack_io_s3_api_v1_s3_controller_user_proto protoreflect.FileDescriptor

var file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f,
	0x73, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x33, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69,
	0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f, 0x73, 0x33, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x10, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf9, 0x01, 0x0a, 0x1f, 0x53, 0x33, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x78,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x22, 0xee, 0x01, 0x0a, 0x16, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78,
	0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x6d, 0x61, 0x78, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0xcb, 0x01, 0x0a, 0x16, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65,
	0x79, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x33, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x61, 0x64, 0x64, 0x4b, 0x65, 0x79,
	0x73, 0x22, 0x57, 0x0a, 0x14, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x74, 0x0a, 0x16, 0x53, 0x33,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x32, 0xa1, 0x04, 0x0a, 0x17, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x38,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73,
	0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x55, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f,
	0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x2d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f,
	0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x73, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x38, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x59, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2f,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73,
	0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a,
	0x1e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e,
	0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x73, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x38, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescOnce sync.Once
	file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescData = file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDesc
)

func file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescGZIP() []byte {
	file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescOnce.Do(func() {
		file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescData)
	})
	return file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDescData
}

var file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metalstack_io_s3_api_v1_s3_controller_user_proto_goTypes = []interface{}{
	(*S3ControllerUser)(nil),                // 0: metalstack.io.s3.api.v1.S3ControllerUser
	(*S3ControllerUserWithCredentials)(nil), // 1: metalstack.io.s3.api.v1.S3ControllerUserWithCredentials
	(*S3ControllerUserCreate)(nil),          // 2: metalstack.io.s3.api.v1.S3ControllerUserCreate
	(*S3ControllerUserUpdate)(nil),          // 3: metalstack.io.s3.api.v1.S3ControllerUserUpdate
	(*S3ControllerUserList)(nil),            // 4: metalstack.io.s3.api.v1.S3ControllerUserList
	(*S3ControllerUserDelete)(nil),          // 5: metalstack.io.s3.api.v1.S3ControllerUserDelete
	(*S3UserKey)(nil),                       // 6: metalstack.io.s3.api.v1.S3UserKey
	(*Empty)(nil),                           // 7: metalstack.io.s3.api.v1.Empty
}
var file_metalstack_io_s3_api_v1_s3_controller_user_proto_depIdxs = []int32{
	6, // 0: metalstack.io.s3.api.v1.S3ControllerUserWithCredentials.keys:type_name -> metalstack.io.s3.api.v1.S3UserKey
	6, // 1: metalstack.io.s3.api.v1.S3ControllerUserCreate.key:type_name -> metalstack.io.s3.api.v1.S3UserKey
	6, // 2: metalstack.io.s3.api.v1.S3ControllerUserUpdate.add_keys:type_name -> metalstack.io.s3.api.v1.S3UserKey
	0, // 3: metalstack.io.s3.api.v1.S3ControllerUserList.users:type_name -> metalstack.io.s3.api.v1.S3ControllerUser
	0, // 4: metalstack.io.s3.api.v1.S3ControllerUserService.Get:input_type -> metalstack.io.s3.api.v1.S3ControllerUser
	7, // 5: metalstack.io.s3.api.v1.S3ControllerUserService.List:input_type -> metalstack.io.s3.api.v1.Empty
	2, // 6: metalstack.io.s3.api.v1.S3ControllerUserService.Create:input_type -> metalstack.io.s3.api.v1.S3ControllerUserCreate
	5, // 7: metalstack.io.s3.api.v1.S3ControllerUserService.Delete:input_type -> metalstack.io.s3.api.v1.S3ControllerUserDelete
	3, // 8: metalstack.io.s3.api.v1.S3ControllerUserService.Update:input_type -> metalstack.io.s3.api.v1.S3ControllerUserUpdate
	1, // 9: metalstack.io.s3.api.v1.S3ControllerUserService.Get:output_type -> metalstack.io.s3.api.v1.S3ControllerUserWithCredentials
	4, // 10: metalstack.io.s3.api.v1.S3ControllerUserService.List:output_type -> metalstack.io.s3.api.v1.S3ControllerUserList
	1, // 11: metalstack.io.s3.api.v1.S3ControllerUserService.Create:output_type -> metalstack.io.s3.api.v1.S3ControllerUserWithCredentials
	7, // 12: metalstack.io.s3.api.v1.S3ControllerUserService.Delete:output_type -> metalstack.io.s3.api.v1.Empty
	1, // 13: metalstack.io.s3.api.v1.S3ControllerUserService.Update:output_type -> metalstack.io.s3.api.v1.S3ControllerUserWithCredentials
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_metalstack_io_s3_api_v1_s3_controller_user_proto_init() }
func file_metalstack_io_s3_api_v1_s3_controller_user_proto_init() {
	if File_metalstack_io_s3_api_v1_s3_controller_user_proto != nil {
		return
	}
	file_metalstack_io_s3_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ControllerUser); i {
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
		file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ControllerUserWithCredentials); i {
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
		file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ControllerUserCreate); i {
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
		file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ControllerUserUpdate); i {
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
		file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ControllerUserList); i {
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
		file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ControllerUserDelete); i {
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
			RawDescriptor: file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_io_s3_api_v1_s3_controller_user_proto_goTypes,
		DependencyIndexes: file_metalstack_io_s3_api_v1_s3_controller_user_proto_depIdxs,
		MessageInfos:      file_metalstack_io_s3_api_v1_s3_controller_user_proto_msgTypes,
	}.Build()
	File_metalstack_io_s3_api_v1_s3_controller_user_proto = out.File
	file_metalstack_io_s3_api_v1_s3_controller_user_proto_rawDesc = nil
	file_metalstack_io_s3_api_v1_s3_controller_user_proto_goTypes = nil
	file_metalstack_io_s3_api_v1_s3_controller_user_proto_depIdxs = nil
}
