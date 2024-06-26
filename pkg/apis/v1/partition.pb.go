// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: metalstack/io/s3/api/v1/partition.proto

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

type S3Partition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ready           bool   `protobuf:"varint,2,opt,name=ready,proto3" json:"ready,omitempty"`
	StorageEndpoint string `protobuf:"bytes,3,opt,name=storage_endpoint,json=storageEndpoint,proto3" json:"storage_endpoint,omitempty"`
}

func (x *S3Partition) Reset() {
	*x = S3Partition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_partition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Partition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Partition) ProtoMessage() {}

func (x *S3Partition) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_partition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Partition.ProtoReflect.Descriptor instead.
func (*S3Partition) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_partition_proto_rawDescGZIP(), []int{0}
}

func (x *S3Partition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *S3Partition) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

func (x *S3Partition) GetStorageEndpoint() string {
	if x != nil {
		return x.StorageEndpoint
	}
	return ""
}

type S3PartitionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partition *S3Partition `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
}

func (x *S3PartitionResponse) Reset() {
	*x = S3PartitionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_partition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3PartitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3PartitionResponse) ProtoMessage() {}

func (x *S3PartitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_partition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3PartitionResponse.ProtoReflect.Descriptor instead.
func (*S3PartitionResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_partition_proto_rawDescGZIP(), []int{1}
}

func (x *S3PartitionResponse) GetPartition() *S3Partition {
	if x != nil {
		return x.Partition
	}
	return nil
}

type S3PartitionListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partitions []*S3PartitionResponse `protobuf:"bytes,1,rep,name=partitions,proto3" json:"partitions,omitempty"`
}

func (x *S3PartitionListResponse) Reset() {
	*x = S3PartitionListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_s3_api_v1_partition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3PartitionListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3PartitionListResponse) ProtoMessage() {}

func (x *S3PartitionListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_s3_api_v1_partition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3PartitionListResponse.ProtoReflect.Descriptor instead.
func (*S3PartitionListResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_io_s3_api_v1_partition_proto_rawDescGZIP(), []int{2}
}

func (x *S3PartitionListResponse) GetPartitions() []*S3PartitionResponse {
	if x != nil {
		return x.Partitions
	}
	return nil
}

var File_metalstack_io_s3_api_v1_partition_proto protoreflect.FileDescriptor

var file_metalstack_io_s3_api_v1_partition_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f,
	0x73, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x1a, 0x24, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69,
	0x6f, 0x2f, 0x73, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0b, 0x53, 0x33, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x29, 0x0a,
	0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x13, 0x53, 0x33, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x67, 0x0a, 0x17, 0x53, 0x33, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c,
	0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x6e, 0x0a, 0x12,
	0x53, 0x33, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x58, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x30, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x33, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x33, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_metalstack_io_s3_api_v1_partition_proto_rawDescOnce sync.Once
	file_metalstack_io_s3_api_v1_partition_proto_rawDescData = file_metalstack_io_s3_api_v1_partition_proto_rawDesc
)

func file_metalstack_io_s3_api_v1_partition_proto_rawDescGZIP() []byte {
	file_metalstack_io_s3_api_v1_partition_proto_rawDescOnce.Do(func() {
		file_metalstack_io_s3_api_v1_partition_proto_rawDescData = protoimpl.X.CompressGZIP(file_metalstack_io_s3_api_v1_partition_proto_rawDescData)
	})
	return file_metalstack_io_s3_api_v1_partition_proto_rawDescData
}

var file_metalstack_io_s3_api_v1_partition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_metalstack_io_s3_api_v1_partition_proto_goTypes = []any{
	(*S3Partition)(nil),             // 0: metalstack.io.s3.api.v1.S3Partition
	(*S3PartitionResponse)(nil),     // 1: metalstack.io.s3.api.v1.S3PartitionResponse
	(*S3PartitionListResponse)(nil), // 2: metalstack.io.s3.api.v1.S3PartitionListResponse
	(*Empty)(nil),                   // 3: metalstack.io.s3.api.v1.Empty
}
var file_metalstack_io_s3_api_v1_partition_proto_depIdxs = []int32{
	0, // 0: metalstack.io.s3.api.v1.S3PartitionResponse.partition:type_name -> metalstack.io.s3.api.v1.S3Partition
	1, // 1: metalstack.io.s3.api.v1.S3PartitionListResponse.partitions:type_name -> metalstack.io.s3.api.v1.S3PartitionResponse
	3, // 2: metalstack.io.s3.api.v1.S3PartitionService.List:input_type -> metalstack.io.s3.api.v1.Empty
	2, // 3: metalstack.io.s3.api.v1.S3PartitionService.List:output_type -> metalstack.io.s3.api.v1.S3PartitionListResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_metalstack_io_s3_api_v1_partition_proto_init() }
func file_metalstack_io_s3_api_v1_partition_proto_init() {
	if File_metalstack_io_s3_api_v1_partition_proto != nil {
		return
	}
	file_metalstack_io_s3_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metalstack_io_s3_api_v1_partition_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*S3Partition); i {
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
		file_metalstack_io_s3_api_v1_partition_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*S3PartitionResponse); i {
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
		file_metalstack_io_s3_api_v1_partition_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*S3PartitionListResponse); i {
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
			RawDescriptor: file_metalstack_io_s3_api_v1_partition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_io_s3_api_v1_partition_proto_goTypes,
		DependencyIndexes: file_metalstack_io_s3_api_v1_partition_proto_depIdxs,
		MessageInfos:      file_metalstack_io_s3_api_v1_partition_proto_msgTypes,
	}.Build()
	File_metalstack_io_s3_api_v1_partition_proto = out.File
	file_metalstack_io_s3_api_v1_partition_proto_rawDesc = nil
	file_metalstack_io_s3_api_v1_partition_proto_goTypes = nil
	file_metalstack_io_s3_api_v1_partition_proto_depIdxs = nil
}
