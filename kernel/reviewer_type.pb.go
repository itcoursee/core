// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/core/kernel/reviewer_type.proto

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

// 审核者类型
type ReviewerType int32

const (
	// 占位，不要使用
	ReviewerType_REVIEWER_TYPE_UNSPECIFIED ReviewerType = 0
	// 任意审核者
	ReviewerType_REVIEWER_TYPE_ANY ReviewerType = 1
)

// Enum value maps for ReviewerType.
var (
	ReviewerType_name = map[int32]string{
		0: "REVIEWER_TYPE_UNSPECIFIED",
		1: "REVIEWER_TYPE_ANY",
	}
	ReviewerType_value = map[string]int32{
		"REVIEWER_TYPE_UNSPECIFIED": 0,
		"REVIEWER_TYPE_ANY":         1,
	}
)

func (x ReviewerType) Enum() *ReviewerType {
	p := new(ReviewerType)
	*p = x
	return p
}

func (x ReviewerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReviewerType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_itcoursee_core_kernel_reviewer_type_proto_enumTypes[0].Descriptor()
}

func (ReviewerType) Type() protoreflect.EnumType {
	return &file_com_itcoursee_core_kernel_reviewer_type_proto_enumTypes[0]
}

func (x ReviewerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReviewerType.Descriptor instead.
func (ReviewerType) EnumDescriptor() ([]byte, []int) {
	return file_com_itcoursee_core_kernel_reviewer_type_proto_rawDescGZIP(), []int{0}
}

var File_com_itcoursee_core_kernel_reviewer_type_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_kernel_reviewer_type_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2a, 0x44, 0x0a, 0x0c, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45,
	0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x56,
	0x49, 0x45, 0x57, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x59, 0x10, 0x01,
	0x42, 0x46, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x50, 0x01, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x3b, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_kernel_reviewer_type_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_kernel_reviewer_type_proto_rawDescData = file_com_itcoursee_core_kernel_reviewer_type_proto_rawDesc
)

func file_com_itcoursee_core_kernel_reviewer_type_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_kernel_reviewer_type_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_kernel_reviewer_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_kernel_reviewer_type_proto_rawDescData)
	})
	return file_com_itcoursee_core_kernel_reviewer_type_proto_rawDescData
}

var file_com_itcoursee_core_kernel_reviewer_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_itcoursee_core_kernel_reviewer_type_proto_goTypes = []interface{}{
	(ReviewerType)(0), // 0: com.itcoursee.core.kernel.ReviewerType
}
var file_com_itcoursee_core_kernel_reviewer_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_kernel_reviewer_type_proto_init() }
func file_com_itcoursee_core_kernel_reviewer_type_proto_init() {
	if File_com_itcoursee_core_kernel_reviewer_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_core_kernel_reviewer_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_kernel_reviewer_type_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_kernel_reviewer_type_proto_depIdxs,
		EnumInfos:         file_com_itcoursee_core_kernel_reviewer_type_proto_enumTypes,
	}.Build()
	File_com_itcoursee_core_kernel_reviewer_type_proto = out.File
	file_com_itcoursee_core_kernel_reviewer_type_proto_rawDesc = nil
	file_com_itcoursee_core_kernel_reviewer_type_proto_goTypes = nil
	file_com_itcoursee_core_kernel_reviewer_type_proto_depIdxs = nil
}
