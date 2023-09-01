// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/core/review/self.proto

package review

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

// 获取请求
type SelfReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 关键字
	Keyword string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	// 每页个数
	// @gotags: default:"20"
	Size int32 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty" default:"20"`
	// 页数
	// @gotags: default:"1"
	Page int32 `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty" default:"1"`
	// 排序
	Order string `protobuf:"bytes,10,opt,name=order,proto3" json:"order,omitempty"`
	// 排序字段
	// @gotags: default:"id"
	Sort string `protobuf:"bytes,11,opt,name=sort,proto3" json:"sort,omitempty" default:"id"`
}

func (x *SelfReq) Reset() {
	*x = SelfReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_review_self_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfReq) ProtoMessage() {}

func (x *SelfReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_review_self_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfReq.ProtoReflect.Descriptor instead.
func (*SelfReq) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_review_self_proto_rawDescGZIP(), []int{0}
}

func (x *SelfReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SelfReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SelfReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SelfReq) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *SelfReq) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

// 获取返回
type SelfRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 审核
	Review *vo.Review `protobuf:"bytes,3,opt,name=review,proto3" json:"review,omitempty"`
}

func (x *SelfRsp) Reset() {
	*x = SelfRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_review_self_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfRsp) ProtoMessage() {}

func (x *SelfRsp) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_review_self_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfRsp.ProtoReflect.Descriptor instead.
func (*SelfRsp) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_review_self_proto_rawDescGZIP(), []int{1}
}

func (x *SelfRsp) GetReview() *vo.Review {
	if x != nil {
		return x.Review
	}
	return nil
}

var File_com_itcoursee_core_review_self_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_review_self_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x73, 0x65, 0x6c, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x07, 0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x40, 0x0a, 0x07,
	0x53, 0x65, 0x6c, 0x66, 0x52, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x42, 0x46,
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x01, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x3b,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_review_self_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_review_self_proto_rawDescData = file_com_itcoursee_core_review_self_proto_rawDesc
)

func file_com_itcoursee_core_review_self_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_review_self_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_review_self_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_review_self_proto_rawDescData)
	})
	return file_com_itcoursee_core_review_self_proto_rawDescData
}

var file_com_itcoursee_core_review_self_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_itcoursee_core_review_self_proto_goTypes = []interface{}{
	(*SelfReq)(nil),   // 0: com.itcoursee.core.review.SelfReq
	(*SelfRsp)(nil),   // 1: com.itcoursee.core.review.SelfRsp
	(*vo.Review)(nil), // 2: com.itcoursee.core.vo.Review
}
var file_com_itcoursee_core_review_self_proto_depIdxs = []int32{
	2, // 0: com.itcoursee.core.review.SelfRsp.review:type_name -> com.itcoursee.core.vo.Review
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_review_self_proto_init() }
func file_com_itcoursee_core_review_self_proto_init() {
	if File_com_itcoursee_core_review_self_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_review_self_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfReq); i {
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
		file_com_itcoursee_core_review_self_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfRsp); i {
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
			RawDescriptor: file_com_itcoursee_core_review_self_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_review_self_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_review_self_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_review_self_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_review_self_proto = out.File
	file_com_itcoursee_core_review_self_proto_rawDesc = nil
	file_com_itcoursee_core_review_self_proto_goTypes = nil
	file_com_itcoursee_core_review_self_proto_depIdxs = nil
}
