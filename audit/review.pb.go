// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: com/itcoursee/core/audit/review.proto

package audit

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

// 按审核获取请求
type GetsByReviewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 审核标识
	// @gotags: validate:"required"
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty" validate:"required"`
}

func (x *GetsByReviewReq) Reset() {
	*x = GetsByReviewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_audit_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByReviewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByReviewReq) ProtoMessage() {}

func (x *GetsByReviewReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_audit_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsByReviewReq.ProtoReflect.Descriptor instead.
func (*GetsByReviewReq) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_audit_review_proto_rawDescGZIP(), []int{0}
}

func (x *GetsByReviewReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 按审核获取响应
type GetsByReviewRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据
	Audits []*vo.Audit `protobuf:"bytes,3,rep,name=audits,proto3" json:"audits,omitempty"`
}

func (x *GetsByReviewRsp) Reset() {
	*x = GetsByReviewRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_audit_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByReviewRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByReviewRsp) ProtoMessage() {}

func (x *GetsByReviewRsp) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_audit_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsByReviewRsp.ProtoReflect.Descriptor instead.
func (*GetsByReviewRsp) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_audit_review_proto_rawDescGZIP(), []int{1}
}

func (x *GetsByReviewRsp) GetAudits() []*vo.Audit {
	if x != nil {
		return x.Audits
	}
	return nil
}

var File_com_itcoursee_core_audit_review_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_audit_review_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x73, 0x42,
	0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x06, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x06, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73,
	0x42, 0x43, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x50, 0x01, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x3b,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_audit_review_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_audit_review_proto_rawDescData = file_com_itcoursee_core_audit_review_proto_rawDesc
)

func file_com_itcoursee_core_audit_review_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_audit_review_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_audit_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_audit_review_proto_rawDescData)
	})
	return file_com_itcoursee_core_audit_review_proto_rawDescData
}

var file_com_itcoursee_core_audit_review_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_itcoursee_core_audit_review_proto_goTypes = []interface{}{
	(*GetsByReviewReq)(nil), // 0: com.itcoursee.core.audit.GetsByReviewReq
	(*GetsByReviewRsp)(nil), // 1: com.itcoursee.core.audit.GetsByReviewRsp
	(*vo.Audit)(nil),        // 2: com.itcoursee.core.vo.Audit
}
var file_com_itcoursee_core_audit_review_proto_depIdxs = []int32{
	2, // 0: com.itcoursee.core.audit.GetsByReviewRsp.audits:type_name -> com.itcoursee.core.vo.Audit
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_audit_review_proto_init() }
func file_com_itcoursee_core_audit_review_proto_init() {
	if File_com_itcoursee_core_audit_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_audit_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsByReviewReq); i {
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
		file_com_itcoursee_core_audit_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsByReviewRsp); i {
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
			RawDescriptor: file_com_itcoursee_core_audit_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_audit_review_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_audit_review_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_audit_review_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_audit_review_proto = out.File
	file_com_itcoursee_core_audit_review_proto_rawDesc = nil
	file_com_itcoursee_core_audit_review_proto_goTypes = nil
	file_com_itcoursee_core_audit_review_proto_depIdxs = nil
}
