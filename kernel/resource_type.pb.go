// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: com/itcoursee/core/kernel/resource_type.proto

package kernel

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

// 资源类型
type ResourceType int32

const (
	// 未知，不要使用
	ResourceType_RESOURCE_TYPE_UNSPECIFIED ResourceType = 0
	// 图片
	ResourceType_RESOURCE_TYPE_IMAGE ResourceType = 1
	// 可播放可观看
	// 视频
	// 音频
	// 未来可能支持的如我们自己开发的课程格式
	ResourceType_RESOURCE_TYPE_PLAYBACK ResourceType = 2
	// 文件
	ResourceType_RESOURCE_TYPE_FILE ResourceType = 3
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "RESOURCE_TYPE_UNSPECIFIED",
		1: "RESOURCE_TYPE_IMAGE",
		2: "RESOURCE_TYPE_PLAYBACK",
		3: "RESOURCE_TYPE_FILE",
	}
	ResourceType_value = map[string]int32{
		"RESOURCE_TYPE_UNSPECIFIED": 0,
		"RESOURCE_TYPE_IMAGE":       1,
		"RESOURCE_TYPE_PLAYBACK":    2,
		"RESOURCE_TYPE_FILE":        3,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_itcoursee_core_kernel_resource_type_proto_enumTypes[0].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_com_itcoursee_core_kernel_resource_type_proto_enumTypes[0]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_com_itcoursee_core_kernel_resource_type_proto_rawDescGZIP(), []int{0}
}

var File_com_itcoursee_core_kernel_resource_type_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_kernel_resource_type_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2a, 0x7a, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45,
	0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x16,
	0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x46, 0x49, 0x4c, 0x45, 0x10, 0x03, 0x42, 0x46, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x3b, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_kernel_resource_type_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_kernel_resource_type_proto_rawDescData = file_com_itcoursee_core_kernel_resource_type_proto_rawDesc
)

func file_com_itcoursee_core_kernel_resource_type_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_kernel_resource_type_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_kernel_resource_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_kernel_resource_type_proto_rawDescData)
	})
	return file_com_itcoursee_core_kernel_resource_type_proto_rawDescData
}

var file_com_itcoursee_core_kernel_resource_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_itcoursee_core_kernel_resource_type_proto_goTypes = []interface{}{
	(ResourceType)(0), // 0: com.itcoursee.core.kernel.ResourceType
}
var file_com_itcoursee_core_kernel_resource_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_kernel_resource_type_proto_init() }
func file_com_itcoursee_core_kernel_resource_type_proto_init() {
	if File_com_itcoursee_core_kernel_resource_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_core_kernel_resource_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_kernel_resource_type_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_kernel_resource_type_proto_depIdxs,
		EnumInfos:         file_com_itcoursee_core_kernel_resource_type_proto_enumTypes,
	}.Build()
	File_com_itcoursee_core_kernel_resource_type_proto = out.File
	file_com_itcoursee_core_kernel_resource_type_proto_rawDesc = nil
	file_com_itcoursee_core_kernel_resource_type_proto_goTypes = nil
	file_com_itcoursee_core_kernel_resource_type_proto_depIdxs = nil
}
