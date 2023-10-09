// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: com/itcoursee/core/course/review.proto

package course

import (
	vo "github.com/itcoursee/core/vo"
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

// 提交审核请求
type ReviewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	// @gotags: validate:"required"
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty" validate:"required"`
}

func (x *ReviewReq) Reset() {
	*x = ReviewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_course_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewReq) ProtoMessage() {}

func (x *ReviewReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_course_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewReq.ProtoReflect.Descriptor instead.
func (*ReviewReq) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_course_review_proto_rawDescGZIP(), []int{0}
}

func (x *ReviewReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 提交审核响应
type ReviewRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据
	Course *vo.Course `protobuf:"bytes,3,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *ReviewRsp) Reset() {
	*x = ReviewRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_course_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewRsp) ProtoMessage() {}

func (x *ReviewRsp) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_course_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewRsp.ProtoReflect.Descriptor instead.
func (*ReviewRsp) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_course_review_proto_rawDescGZIP(), []int{1}
}

func (x *ReviewRsp) GetCourse() *vo.Course {
	if x != nil {
		return x.Course
	}
	return nil
}

var File_com_itcoursee_core_course_review_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_course_review_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x09, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x73,
	0x70, 0x12, 0x35, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x42, 0x46, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_course_review_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_course_review_proto_rawDescData = file_com_itcoursee_core_course_review_proto_rawDesc
)

func file_com_itcoursee_core_course_review_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_course_review_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_course_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_course_review_proto_rawDescData)
	})
	return file_com_itcoursee_core_course_review_proto_rawDescData
}

var file_com_itcoursee_core_course_review_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_itcoursee_core_course_review_proto_goTypes = []interface{}{
	(*ReviewReq)(nil), // 0: com.itcoursee.core.course.ReviewReq
	(*ReviewRsp)(nil), // 1: com.itcoursee.core.course.ReviewRsp
	(*vo.Course)(nil), // 2: com.itcoursee.core.vo.Course
}
var file_com_itcoursee_core_course_review_proto_depIdxs = []int32{
	2, // 0: com.itcoursee.core.course.ReviewRsp.course:type_name -> com.itcoursee.core.vo.Course
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_course_review_proto_init() }
func file_com_itcoursee_core_course_review_proto_init() {
	if File_com_itcoursee_core_course_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_course_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewReq); i {
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
		file_com_itcoursee_core_course_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewRsp); i {
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
			RawDescriptor: file_com_itcoursee_core_course_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_course_review_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_course_review_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_course_review_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_course_review_proto = out.File
	file_com_itcoursee_core_course_review_proto_rawDesc = nil
	file_com_itcoursee_core_course_review_proto_goTypes = nil
	file_com_itcoursee_core_course_review_proto_depIdxs = nil
}
