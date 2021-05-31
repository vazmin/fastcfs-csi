// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: fcfs_fused_mount.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MountFcfsFusedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasePath  string   `protobuf:"bytes,1,opt,name=basePath,proto3" json:"basePath,omitempty"`
	MountArgs []string `protobuf:"bytes,2,rep,name=mountArgs,proto3" json:"mountArgs,omitempty"`
	// Secrets required by plugin to complete volume creation request.
	// This field is OPTIONAL. Refer to the `Secrets Requirements`
	// section on how to use this field.
	Secrets map[string]string `protobuf:"bytes,3,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MountFcfsFusedRequest) Reset() {
	*x = MountFcfsFusedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fcfs_fused_mount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountFcfsFusedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountFcfsFusedRequest) ProtoMessage() {}

func (x *MountFcfsFusedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fcfs_fused_mount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountFcfsFusedRequest.ProtoReflect.Descriptor instead.
func (*MountFcfsFusedRequest) Descriptor() ([]byte, []int) {
	return file_fcfs_fused_mount_proto_rawDescGZIP(), []int{0}
}

func (x *MountFcfsFusedRequest) GetBasePath() string {
	if x != nil {
		return x.BasePath
	}
	return ""
}

func (x *MountFcfsFusedRequest) GetMountArgs() []string {
	if x != nil {
		return x.MountArgs
	}
	return nil
}

func (x *MountFcfsFusedRequest) GetSecrets() map[string]string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

type MountFcfsFusedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *MountFcfsFusedResponse) Reset() {
	*x = MountFcfsFusedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fcfs_fused_mount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountFcfsFusedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountFcfsFusedResponse) ProtoMessage() {}

func (x *MountFcfsFusedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fcfs_fused_mount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountFcfsFusedResponse.ProtoReflect.Descriptor instead.
func (*MountFcfsFusedResponse) Descriptor() ([]byte, []int) {
	return file_fcfs_fused_mount_proto_rawDescGZIP(), []int{1}
}

func (x *MountFcfsFusedResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var file_fcfs_fused_mount_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1059,
		Name:          "csi_secret",
		Tag:           "varint,1059,opt,name=csi_secret",
		Filename:      "fcfs_fused_mount.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// Indicates that a field MAY contain information that is sensitive
	// and MUST be treated as such (e.g. not logged).
	//
	// optional bool csi_secret = 1059;
	E_CsiSecret = &file_fcfs_fused_mount_proto_extTypes[0]
)

var File_fcfs_fused_mount_proto protoreflect.FileDescriptor

var file_fcfs_fused_mount_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x63, 0x66, 0x73, 0x5f, 0x66, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x15, 0x4d,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x63, 0x66, 0x73, 0x46, 0x75, 0x73, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x42,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x63, 0x66, 0x73, 0x46, 0x75, 0x73, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0x98, 0x42, 0x01, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x30,
	0x0a, 0x16, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x63, 0x66, 0x73, 0x46, 0x75, 0x73, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x32, 0x53, 0x0a, 0x0c, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x0e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x63, 0x66, 0x73, 0x46, 0x75, 0x73,
	0x65, 0x64, 0x12, 0x16, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x63, 0x66, 0x73, 0x46, 0x75,
	0x73, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x4d, 0x6f, 0x75,
	0x6e, 0x74, 0x46, 0x63, 0x66, 0x73, 0x46, 0x75, 0x73, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x3a, 0x3d, 0x0a, 0x0a, 0x63, 0x73, 0x69, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xa3, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x73, 0x69, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fcfs_fused_mount_proto_rawDescOnce sync.Once
	file_fcfs_fused_mount_proto_rawDescData = file_fcfs_fused_mount_proto_rawDesc
)

func file_fcfs_fused_mount_proto_rawDescGZIP() []byte {
	file_fcfs_fused_mount_proto_rawDescOnce.Do(func() {
		file_fcfs_fused_mount_proto_rawDescData = protoimpl.X.CompressGZIP(file_fcfs_fused_mount_proto_rawDescData)
	})
	return file_fcfs_fused_mount_proto_rawDescData
}

var file_fcfs_fused_mount_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_fcfs_fused_mount_proto_goTypes = []interface{}{
	(*MountFcfsFusedRequest)(nil),     // 0: MountFcfsFusedRequest
	(*MountFcfsFusedResponse)(nil),    // 1: MountFcfsFusedResponse
	nil,                               // 2: MountFcfsFusedRequest.SecretsEntry
	(*descriptorpb.FieldOptions)(nil), // 3: google.protobuf.FieldOptions
}
var file_fcfs_fused_mount_proto_depIdxs = []int32{
	2, // 0: MountFcfsFusedRequest.secrets:type_name -> MountFcfsFusedRequest.SecretsEntry
	3, // 1: csi_secret:extendee -> google.protobuf.FieldOptions
	0, // 2: MountService.MountFcfsFused:input_type -> MountFcfsFusedRequest
	1, // 3: MountService.MountFcfsFused:output_type -> MountFcfsFusedResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fcfs_fused_mount_proto_init() }
func file_fcfs_fused_mount_proto_init() {
	if File_fcfs_fused_mount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fcfs_fused_mount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountFcfsFusedRequest); i {
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
		file_fcfs_fused_mount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountFcfsFusedResponse); i {
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
			RawDescriptor: file_fcfs_fused_mount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_fcfs_fused_mount_proto_goTypes,
		DependencyIndexes: file_fcfs_fused_mount_proto_depIdxs,
		MessageInfos:      file_fcfs_fused_mount_proto_msgTypes,
		ExtensionInfos:    file_fcfs_fused_mount_proto_extTypes,
	}.Build()
	File_fcfs_fused_mount_proto = out.File
	file_fcfs_fused_mount_proto_rawDesc = nil
	file_fcfs_fused_mount_proto_goTypes = nil
	file_fcfs_fused_mount_proto_depIdxs = nil
}