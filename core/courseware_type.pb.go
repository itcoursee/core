// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/core/courseware_type.proto

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

// 课件类型
type CoursewareType int32

const (
	// 未知，不要使用
	CoursewareType_COURSEWARE_TYPE_UNSPECIFIED CoursewareType = 0
	// 课程
	CoursewareType_COURSEWARE_TYPE_COURSE CoursewareType = 1
	// 讲次
	CoursewareType_COURSEWARE_TYPE_LECTURE CoursewareType = 2
)

// Enum value maps for CoursewareType.
var (
	CoursewareType_name = map[int32]string{
		0: "COURSEWARE_TYPE_UNSPECIFIED",
		1: "COURSEWARE_TYPE_COURSE",
		2: "COURSEWARE_TYPE_LECTURE",
	}
	CoursewareType_value = map[string]int32{
		"COURSEWARE_TYPE_UNSPECIFIED": 0,
		"COURSEWARE_TYPE_COURSE":      1,
		"COURSEWARE_TYPE_LECTURE":     2,
	}
)

func (x CoursewareType) Enum() *CoursewareType {
	p := new(CoursewareType)
	*p = x
	return p
}

func (x CoursewareType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoursewareType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_itcoursee_core_courseware_type_proto_enumTypes[0].Descriptor()
}

func (CoursewareType) Type() protoreflect.EnumType {
	return &file_com_itcoursee_core_courseware_type_proto_enumTypes[0]
}

func (x CoursewareType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoursewareType.Descriptor instead.
func (CoursewareType) EnumDescriptor() ([]byte, []int) {
	return file_com_itcoursee_core_courseware_type_proto_rawDescGZIP(), []int{0}
}

var File_com_itcoursee_core_courseware_type_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_courseware_type_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0x6a,
	0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x57, 0x41, 0x52, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x57, 0x41, 0x52, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a,
	0x17, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x57, 0x41, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4c, 0x45, 0x43, 0x54, 0x55, 0x52, 0x45, 0x10, 0x02, 0x42, 0x3b, 0x0a, 0x12, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_courseware_type_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_courseware_type_proto_rawDescData = file_com_itcoursee_core_courseware_type_proto_rawDesc
)

func file_com_itcoursee_core_courseware_type_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_courseware_type_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_courseware_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_courseware_type_proto_rawDescData)
	})
	return file_com_itcoursee_core_courseware_type_proto_rawDescData
}

var file_com_itcoursee_core_courseware_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_itcoursee_core_courseware_type_proto_goTypes = []interface{}{
	(CoursewareType)(0), // 0: com.itcoursee.core.CoursewareType
}
var file_com_itcoursee_core_courseware_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_courseware_type_proto_init() }
func file_com_itcoursee_core_courseware_type_proto_init() {
	if File_com_itcoursee_core_courseware_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_core_courseware_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_courseware_type_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_courseware_type_proto_depIdxs,
		EnumInfos:         file_com_itcoursee_core_courseware_type_proto_enumTypes,
	}.Build()
	File_com_itcoursee_core_courseware_type_proto = out.File
	file_com_itcoursee_core_courseware_type_proto_rawDesc = nil
	file_com_itcoursee_core_courseware_type_proto_goTypes = nil
	file_com_itcoursee_core_courseware_type_proto_depIdxs = nil
}
