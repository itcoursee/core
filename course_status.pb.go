// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/core/course_status.proto

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

// 课程状态
type CourseStatus int32

const (
	// 未知，不要使用
	CourseStatus_COURSE_STATUS_UNSPECIFIED CourseStatus = 0
	// 已创建
	CourseStatus_COURSE_STATUS_CREATED CourseStatus = 1
	// 审核中
	CourseStatus_COURSE_STATUS_REVIEWING CourseStatus = 10
	// 审核通过
	CourseStatus_COURSE_STATUS_REVIEWED CourseStatus = 20
	// 已上架
	CourseStatus_COURSE_STATUS_SALES CourseStatus = 21
	// 审核失败
	CourseStatus_COURSE_STATUS_FAILED CourseStatus = 30
	// 已下架
	CourseStatus_COURSE_STATUS_REMOVED CourseStatus = 31
)

// Enum value maps for CourseStatus.
var (
	CourseStatus_name = map[int32]string{
		0:  "COURSE_STATUS_UNSPECIFIED",
		1:  "COURSE_STATUS_CREATED",
		10: "COURSE_STATUS_REVIEWING",
		20: "COURSE_STATUS_REVIEWED",
		21: "COURSE_STATUS_SALES",
		30: "COURSE_STATUS_FAILED",
		31: "COURSE_STATUS_REMOVED",
	}
	CourseStatus_value = map[string]int32{
		"COURSE_STATUS_UNSPECIFIED": 0,
		"COURSE_STATUS_CREATED":     1,
		"COURSE_STATUS_REVIEWING":   10,
		"COURSE_STATUS_REVIEWED":    20,
		"COURSE_STATUS_SALES":       21,
		"COURSE_STATUS_FAILED":      30,
		"COURSE_STATUS_REMOVED":     31,
	}
)

func (x CourseStatus) Enum() *CourseStatus {
	p := new(CourseStatus)
	*p = x
	return p
}

func (x CourseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CourseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_com_itcoursee_core_course_status_proto_enumTypes[0].Descriptor()
}

func (CourseStatus) Type() protoreflect.EnumType {
	return &file_com_itcoursee_core_course_status_proto_enumTypes[0]
}

func (x CourseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CourseStatus.Descriptor instead.
func (CourseStatus) EnumDescriptor() ([]byte, []int) {
	return file_com_itcoursee_core_course_status_proto_rawDescGZIP(), []int{0}
}

var File_com_itcoursee_core_course_status_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_course_status_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0xcf, 0x01, 0x0a,
	0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a,
	0x19, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15,
	0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x55, 0x52, 0x53,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x49,
	0x4e, 0x47, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x45, 0x44, 0x10, 0x14,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x10, 0x15, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x55,
	0x52, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x1e, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x1f, 0x42, 0x36,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_course_status_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_course_status_proto_rawDescData = file_com_itcoursee_core_course_status_proto_rawDesc
)

func file_com_itcoursee_core_course_status_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_course_status_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_course_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_course_status_proto_rawDescData)
	})
	return file_com_itcoursee_core_course_status_proto_rawDescData
}

var file_com_itcoursee_core_course_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_itcoursee_core_course_status_proto_goTypes = []interface{}{
	(CourseStatus)(0), // 0: com.itcoursee.core.CourseStatus
}
var file_com_itcoursee_core_course_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_course_status_proto_init() }
func file_com_itcoursee_core_course_status_proto_init() {
	if File_com_itcoursee_core_course_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_core_course_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_course_status_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_course_status_proto_depIdxs,
		EnumInfos:         file_com_itcoursee_core_course_status_proto_enumTypes,
	}.Build()
	File_com_itcoursee_core_course_status_proto = out.File
	file_com_itcoursee_core_course_status_proto_rawDesc = nil
	file_com_itcoursee_core_course_status_proto_goTypes = nil
	file_com_itcoursee_core_course_status_proto_depIdxs = nil
}
