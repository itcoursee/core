// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/core/courseware/course.proto

package courseware

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

// 创建课程课件请求
type AddByCourseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程编号
	// @gotags: validate:"required"
	Course int64 `protobuf:"varint,3,opt,name=course,proto3" json:"course,omitempty" validate:"required"`
	// 课件名称
	// @gotags: validate:"required,max=255"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" validate:"required,max=255"`
}

func (x *AddByCourseReq) Reset() {
	*x = AddByCourseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_courseware_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddByCourseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddByCourseReq) ProtoMessage() {}

func (x *AddByCourseReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_courseware_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddByCourseReq.ProtoReflect.Descriptor instead.
func (*AddByCourseReq) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_courseware_course_proto_rawDescGZIP(), []int{0}
}

func (x *AddByCourseReq) GetCourse() int64 {
	if x != nil {
		return x.Course
	}
	return 0
}

func (x *AddByCourseReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 获得课程的所有课件请求
type GetsByCourseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程编号
	// @gotags: validate:"required"
	Course int64 `protobuf:"varint,3,opt,name=course,proto3" json:"course,omitempty" validate:"required"`
}

func (x *GetsByCourseReq) Reset() {
	*x = GetsByCourseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_courseware_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByCourseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByCourseReq) ProtoMessage() {}

func (x *GetsByCourseReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_courseware_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsByCourseReq.ProtoReflect.Descriptor instead.
func (*GetsByCourseReq) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_courseware_course_proto_rawDescGZIP(), []int{1}
}

func (x *GetsByCourseReq) GetCourse() int64 {
	if x != nil {
		return x.Course
	}
	return 0
}

var File_com_itcoursee_core_courseware_course_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_courseware_course_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x22, 0x3c, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x42, 0x52, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x3b, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_courseware_course_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_courseware_course_proto_rawDescData = file_com_itcoursee_core_courseware_course_proto_rawDesc
)

func file_com_itcoursee_core_courseware_course_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_courseware_course_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_courseware_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_courseware_course_proto_rawDescData)
	})
	return file_com_itcoursee_core_courseware_course_proto_rawDescData
}

var file_com_itcoursee_core_courseware_course_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_itcoursee_core_courseware_course_proto_goTypes = []interface{}{
	(*AddByCourseReq)(nil),  // 0: com.itcoursee.core.courseware.AddByCourseReq
	(*GetsByCourseReq)(nil), // 1: com.itcoursee.core.courseware.GetsByCourseReq
}
var file_com_itcoursee_core_courseware_course_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_courseware_course_proto_init() }
func file_com_itcoursee_core_courseware_course_proto_init() {
	if File_com_itcoursee_core_courseware_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_courseware_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddByCourseReq); i {
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
		file_com_itcoursee_core_courseware_course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsByCourseReq); i {
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
			RawDescriptor: file_com_itcoursee_core_courseware_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_courseware_course_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_courseware_course_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_courseware_course_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_courseware_course_proto = out.File
	file_com_itcoursee_core_courseware_course_proto_rawDesc = nil
	file_com_itcoursee_core_courseware_course_proto_goTypes = nil
	file_com_itcoursee_core_courseware_course_proto_depIdxs = nil
}
