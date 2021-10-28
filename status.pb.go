// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/core/status.proto

package core

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

// 状态
type Status int32

const (
	// 未知，不要使用
	Status_STATUS_UNSPECIFIED Status = 0
	// 已创建
	Status_CREATED Status = 1
	// 上传中
	Status_UPLOADING Status = 10
	// 已上传
	Status_UPLOADED Status = 11
	// 转码中
	Status_CONVERTING Status = 20
	// 转码完成
	Status_CONVERTED Status = 21
	// 售卖中
	Status_SELLING Status = 30
	// 已下架
	Status_REMOVED Status = 31
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "STATUS_UNSPECIFIED",
		1:  "CREATED",
		10: "UPLOADING",
		11: "UPLOADED",
		20: "CONVERTING",
		21: "CONVERTED",
		30: "SELLING",
		31: "REMOVED",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATED":            1,
		"UPLOADING":          10,
		"UPLOADED":           11,
		"CONVERTING":         20,
		"CONVERTED":          21,
		"SELLING":            30,
		"REMOVED":            31,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_itcoursee_core_status_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_itcoursee_core_status_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_itcoursee_core_status_proto_rawDescGZIP(), []int{0}
}

var File_itcoursee_core_status_proto protoreflect.FileDescriptor

var file_itcoursee_core_status_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0x83, 0x01,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08,
	0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x14, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x54, 0x45, 0x44, 0x10, 0x15, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x4c,
	0x4c, 0x49, 0x4e, 0x47, 0x10, 0x1e, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45,
	0x44, 0x10, 0x1f, 0x42, 0x36, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x1e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_itcoursee_core_status_proto_rawDescOnce sync.Once
	file_itcoursee_core_status_proto_rawDescData = file_itcoursee_core_status_proto_rawDesc
)

func file_itcoursee_core_status_proto_rawDescGZIP() []byte {
	file_itcoursee_core_status_proto_rawDescOnce.Do(func() {
		file_itcoursee_core_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_core_status_proto_rawDescData)
	})
	return file_itcoursee_core_status_proto_rawDescData
}

var file_itcoursee_core_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_itcoursee_core_status_proto_goTypes = []interface{}{
	(Status)(0), // 0: itcoursee.core.Status
}
var file_itcoursee_core_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_itcoursee_core_status_proto_init() }
func file_itcoursee_core_status_proto_init() {
	if File_itcoursee_core_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_itcoursee_core_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_core_status_proto_goTypes,
		DependencyIndexes: file_itcoursee_core_status_proto_depIdxs,
		EnumInfos:         file_itcoursee_core_status_proto_enumTypes,
	}.Build()
	File_itcoursee_core_status_proto = out.File
	file_itcoursee_core_status_proto_rawDesc = nil
	file_itcoursee_core_status_proto_goTypes = nil
	file_itcoursee_core_status_proto_depIdxs = nil
}
