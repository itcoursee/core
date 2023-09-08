// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/core/resource/complete.proto

package resource

import (
	vo "github.com/itcoursee/core/vo"
	core "gitlab.com/ruijc/storage/core"
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

// 完成请求
type CompleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	// @gotags: validate:"required"
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 分片上传回传的对象信息列表
	// @gotags: validate:"omitempty,dive"
	Parts []*core.Part `protobuf:"bytes,10,rep,name=parts,proto3" json:"parts,omitempty"`
}

func (x *CompleteReq) Reset() {
	*x = CompleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_resource_complete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteReq) ProtoMessage() {}

func (x *CompleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_resource_complete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteReq.ProtoReflect.Descriptor instead.
func (*CompleteReq) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_resource_complete_proto_rawDescGZIP(), []int{0}
}

func (x *CompleteReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CompleteReq) GetParts() []*core.Part {
	if x != nil {
		return x.Parts
	}
	return nil
}

// 完成返回
type CompleteRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据
	Resource *vo.Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *CompleteRsp) Reset() {
	*x = CompleteRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_resource_complete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteRsp) ProtoMessage() {}

func (x *CompleteRsp) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_resource_complete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteRsp.ProtoReflect.Descriptor instead.
func (*CompleteRsp) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_resource_complete_proto_rawDescGZIP(), []int{1}
}

func (x *CompleteRsp) GetResource() *vo.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

var File_com_itcoursee_core_resource_complete_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_resource_complete_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x75, 0x69, 0x6a, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x75, 0x69, 0x6a, 0x63, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x52, 0x05,
	0x70, 0x61, 0x72, 0x74, 0x73, 0x22, 0x4a, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x4c, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_resource_complete_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_resource_complete_proto_rawDescData = file_com_itcoursee_core_resource_complete_proto_rawDesc
)

func file_com_itcoursee_core_resource_complete_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_resource_complete_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_resource_complete_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_resource_complete_proto_rawDescData)
	})
	return file_com_itcoursee_core_resource_complete_proto_rawDescData
}

var file_com_itcoursee_core_resource_complete_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_itcoursee_core_resource_complete_proto_goTypes = []interface{}{
	(*CompleteReq)(nil), // 0: com.itcoursee.core.resource.CompleteReq
	(*CompleteRsp)(nil), // 1: com.itcoursee.core.resource.CompleteRsp
	(*core.Part)(nil),   // 2: com.ruijc.storage.core.Part
	(*vo.Resource)(nil), // 3: com.itcoursee.core.vo.Resource
}
var file_com_itcoursee_core_resource_complete_proto_depIdxs = []int32{
	2, // 0: com.itcoursee.core.resource.CompleteReq.parts:type_name -> com.ruijc.storage.core.Part
	3, // 1: com.itcoursee.core.resource.CompleteRsp.resource:type_name -> com.itcoursee.core.vo.Resource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_resource_complete_proto_init() }
func file_com_itcoursee_core_resource_complete_proto_init() {
	if File_com_itcoursee_core_resource_complete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_resource_complete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteReq); i {
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
		file_com_itcoursee_core_resource_complete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteRsp); i {
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
			RawDescriptor: file_com_itcoursee_core_resource_complete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_resource_complete_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_resource_complete_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_resource_complete_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_resource_complete_proto = out.File
	file_com_itcoursee_core_resource_complete_proto_rawDesc = nil
	file_com_itcoursee_core_resource_complete_proto_goTypes = nil
	file_com_itcoursee_core_resource_complete_proto_depIdxs = nil
}
