// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/core/course/add.proto

package course

import (
	kernel "github.com/itcoursee/core/kernel"
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

// 创建课程请求
type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 名称
	// @gotags: validate:"required,max=64"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 类型
	// @gotags: validate:"required,oneof=1 2 3"
	Type kernel.Type `protobuf:"varint,4,opt,name=type,proto3,enum=com.itcoursee.core.kernel.Type" json:"type,omitempty"`
	// 标签列表
	// 最大十个标签
	// @gotags: validate:"required,max=10"
	Tags []int64 `protobuf:"varint,5,rep,packed,name=tags,proto3" json:"tags,omitempty"`
	// 方向
	Direction int64 `protobuf:"varint,6,opt,name=direction,proto3" json:"direction,omitempty"`
	// 分类
	Category int64 `protobuf:"varint,7,opt,name=category,proto3" json:"category,omitempty"`
	// 观众
	Audience kernel.Audience `protobuf:"varint,8,opt,name=audience,proto3,enum=com.itcoursee.core.kernel.Audience" json:"audience,omitempty"`
	// 封面
	Covers []int64 `protobuf:"varint,9,rep,packed,name=covers,proto3" json:"covers,omitempty"`
	// 价格
	Price *kernel.Price `protobuf:"bytes,10,opt,name=price,proto3" json:"price,omitempty"`
	// 简介
	// @gotags: validate:"required,max=4096"
	Profile string `protobuf:"bytes,15,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_course_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_course_add_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_course_add_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddReq) GetType() kernel.Type {
	if x != nil {
		return x.Type
	}
	return kernel.Type(0)
}

func (x *AddReq) GetTags() []int64 {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AddReq) GetDirection() int64 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *AddReq) GetCategory() int64 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *AddReq) GetAudience() kernel.Audience {
	if x != nil {
		return x.Audience
	}
	return kernel.Audience(0)
}

func (x *AddReq) GetCovers() []int64 {
	if x != nil {
		return x.Covers
	}
	return nil
}

func (x *AddReq) GetPrice() *kernel.Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *AddReq) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

// 创建课程返回
type AddRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程
	Course *vo.Course `protobuf:"bytes,3,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *AddRsp) Reset() {
	*x = AddRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_course_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRsp) ProtoMessage() {}

func (x *AddRsp) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_course_add_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRsp.ProtoReflect.Descriptor instead.
func (*AddRsp) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_course_add_proto_rawDescGZIP(), []int{1}
}

func (x *AddRsp) GetCourse() *vo.Course {
	if x != nil {
		return x.Course
	}
	return nil
}

var File_com_itcoursee_core_course_add_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_course_add_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x02, 0x0a, 0x06,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65,
	0x72, 0x6e, 0x65, 0x6c, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x3f, 0x0a,
	0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x3f, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52,
	0x73, 0x70, 0x12, 0x35, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x42, 0x46, 0x0a, 0x19, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_course_add_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_course_add_proto_rawDescData = file_com_itcoursee_core_course_add_proto_rawDesc
)

func file_com_itcoursee_core_course_add_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_course_add_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_course_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_course_add_proto_rawDescData)
	})
	return file_com_itcoursee_core_course_add_proto_rawDescData
}

var file_com_itcoursee_core_course_add_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_itcoursee_core_course_add_proto_goTypes = []interface{}{
	(*AddReq)(nil),       // 0: com.itcoursee.core.course.AddReq
	(*AddRsp)(nil),       // 1: com.itcoursee.core.course.AddRsp
	(kernel.Type)(0),     // 2: com.itcoursee.core.kernel.Type
	(kernel.Audience)(0), // 3: com.itcoursee.core.kernel.Audience
	(*kernel.Price)(nil), // 4: com.itcoursee.core.kernel.Price
	(*vo.Course)(nil),    // 5: com.itcoursee.core.vo.Course
}
var file_com_itcoursee_core_course_add_proto_depIdxs = []int32{
	2, // 0: com.itcoursee.core.course.AddReq.type:type_name -> com.itcoursee.core.kernel.Type
	3, // 1: com.itcoursee.core.course.AddReq.audience:type_name -> com.itcoursee.core.kernel.Audience
	4, // 2: com.itcoursee.core.course.AddReq.price:type_name -> com.itcoursee.core.kernel.Price
	5, // 3: com.itcoursee.core.course.AddRsp.course:type_name -> com.itcoursee.core.vo.Course
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_course_add_proto_init() }
func file_com_itcoursee_core_course_add_proto_init() {
	if File_com_itcoursee_core_course_add_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_course_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_com_itcoursee_core_course_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRsp); i {
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
			RawDescriptor: file_com_itcoursee_core_course_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_course_add_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_course_add_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_course_add_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_course_add_proto = out.File
	file_com_itcoursee_core_course_add_proto_rawDesc = nil
	file_com_itcoursee_core_course_add_proto_goTypes = nil
	file_com_itcoursee_core_course_add_proto_depIdxs = nil
}
