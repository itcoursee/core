// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/core/courseware_subtype.proto

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
type CoursewareSubtype int32

const (
	// 未知，不要使用
	CoursewareSubtype_COURSEWARE_SUBTYPE_UNSPECIFIED CoursewareSubtype = 0
	// 附件
	CoursewareSubtype_ATTACHMENT CoursewareSubtype = 1
	// 封面
	CoursewareSubtype_COVER CoursewareSubtype = 2
	// 试听
	CoursewareSubtype_AUDITION CoursewareSubtype = 3
)

// Enum value maps for CoursewareSubtype.
var (
	CoursewareSubtype_name = map[int32]string{
		0: "COURSEWARE_SUBTYPE_UNSPECIFIED",
		1: "ATTACHMENT",
		2: "COVER",
		3: "AUDITION",
	}
	CoursewareSubtype_value = map[string]int32{
		"COURSEWARE_SUBTYPE_UNSPECIFIED": 0,
		"ATTACHMENT":                     1,
		"COVER":                          2,
		"AUDITION":                       3,
	}
)

func (x CoursewareSubtype) Enum() *CoursewareSubtype {
	p := new(CoursewareSubtype)
	*p = x
	return p
}

func (x CoursewareSubtype) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoursewareSubtype) Descriptor() protoreflect.EnumDescriptor {
	return file_itcoursee_core_courseware_subtype_proto_enumTypes[0].Descriptor()
}

func (CoursewareSubtype) Type() protoreflect.EnumType {
	return &file_itcoursee_core_courseware_subtype_proto_enumTypes[0]
}

func (x CoursewareSubtype) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoursewareSubtype.Descriptor instead.
func (CoursewareSubtype) EnumDescriptor() ([]byte, []int) {
	return file_itcoursee_core_courseware_subtype_proto_rawDescGZIP(), []int{0}
}

var File_itcoursee_core_courseware_subtype_proto protoreflect.FileDescriptor

var file_itcoursee_core_courseware_subtype_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0x60, 0x0a, 0x11, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22,
	0x0a, 0x1e, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x57, 0x41, 0x52, 0x45, 0x5f, 0x53, 0x55, 0x42,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x54, 0x41, 0x43, 0x48, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x55, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x36, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x50, 0x01, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itcoursee_core_courseware_subtype_proto_rawDescOnce sync.Once
	file_itcoursee_core_courseware_subtype_proto_rawDescData = file_itcoursee_core_courseware_subtype_proto_rawDesc
)

func file_itcoursee_core_courseware_subtype_proto_rawDescGZIP() []byte {
	file_itcoursee_core_courseware_subtype_proto_rawDescOnce.Do(func() {
		file_itcoursee_core_courseware_subtype_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_core_courseware_subtype_proto_rawDescData)
	})
	return file_itcoursee_core_courseware_subtype_proto_rawDescData
}

var file_itcoursee_core_courseware_subtype_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_itcoursee_core_courseware_subtype_proto_goTypes = []interface{}{
	(CoursewareSubtype)(0), // 0: itcoursee.core.CoursewareSubtype
}
var file_itcoursee_core_courseware_subtype_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_itcoursee_core_courseware_subtype_proto_init() }
func file_itcoursee_core_courseware_subtype_proto_init() {
	if File_itcoursee_core_courseware_subtype_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_itcoursee_core_courseware_subtype_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_core_courseware_subtype_proto_goTypes,
		DependencyIndexes: file_itcoursee_core_courseware_subtype_proto_depIdxs,
		EnumInfos:         file_itcoursee_core_courseware_subtype_proto_enumTypes,
	}.Build()
	File_itcoursee_core_courseware_subtype_proto = out.File
	file_itcoursee_core_courseware_subtype_proto_rawDesc = nil
	file_itcoursee_core_courseware_subtype_proto_goTypes = nil
	file_itcoursee_core_courseware_subtype_proto_depIdxs = nil
}
