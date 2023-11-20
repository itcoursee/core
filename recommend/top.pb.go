// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: com/itcoursee/core/recommend/top.proto

package recommend

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

// 最佳请求
type TopReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 个数
	// @gotags: default:"10"
	Size int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty" default:"10"`
	// 类型
	// @gotags: validate:"oneof=1"
	Type kernel.RecommendType `protobuf:"varint,4,opt,name=type,proto3,enum=com.itcoursee.core.kernel.RecommendType" json:"type,omitempty" validate:"oneof=1"`
	// 字段筛选
	// @gotags: validate:"dive,oneof=cover"
	Fields []string `protobuf:"bytes,10,rep,name=fields,proto3" json:"fields,omitempty" validate:"dive,oneof=cover"`
}

func (x *TopReq) Reset() {
	*x = TopReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_recommend_top_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopReq) ProtoMessage() {}

func (x *TopReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_recommend_top_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopReq.ProtoReflect.Descriptor instead.
func (*TopReq) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_recommend_top_proto_rawDescGZIP(), []int{0}
}

func (x *TopReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *TopReq) GetType() kernel.RecommendType {
	if x != nil {
		return x.Type
	}
	return kernel.RecommendType(0)
}

func (x *TopReq) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

// 最佳响应
type TopRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据
	Recommends []*vo.Recommend `protobuf:"bytes,4,rep,name=recommends,proto3" json:"recommends,omitempty"`
}

func (x *TopRsp) Reset() {
	*x = TopRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_recommend_top_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopRsp) ProtoMessage() {}

func (x *TopRsp) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_recommend_top_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopRsp.ProtoReflect.Descriptor instead.
func (*TopRsp) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_recommend_top_proto_rawDescGZIP(), []int{1}
}

func (x *TopRsp) GetRecommends() []*vo.Recommend {
	if x != nil {
		return x.Recommends
	}
	return nil
}

var File_com_itcoursee_core_recommend_top_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_recommend_top_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2f, 0x74,
	0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a,
	0x06, 0x54, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x22, 0x4a, 0x0a, 0x06, 0x54, 0x6f, 0x70, 0x52, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x0a, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x73, 0x42, 0x4f, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x50, 0x01, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x3b, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_recommend_top_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_recommend_top_proto_rawDescData = file_com_itcoursee_core_recommend_top_proto_rawDesc
)

func file_com_itcoursee_core_recommend_top_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_recommend_top_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_recommend_top_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_recommend_top_proto_rawDescData)
	})
	return file_com_itcoursee_core_recommend_top_proto_rawDescData
}

var file_com_itcoursee_core_recommend_top_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_itcoursee_core_recommend_top_proto_goTypes = []interface{}{
	(*TopReq)(nil),            // 0: com.itcoursee.core.recommend.TopReq
	(*TopRsp)(nil),            // 1: com.itcoursee.core.recommend.TopRsp
	(kernel.RecommendType)(0), // 2: com.itcoursee.core.kernel.RecommendType
	(*vo.Recommend)(nil),      // 3: com.itcoursee.core.vo.Recommend
}
var file_com_itcoursee_core_recommend_top_proto_depIdxs = []int32{
	2, // 0: com.itcoursee.core.recommend.TopReq.type:type_name -> com.itcoursee.core.kernel.RecommendType
	3, // 1: com.itcoursee.core.recommend.TopRsp.recommends:type_name -> com.itcoursee.core.vo.Recommend
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_recommend_top_proto_init() }
func file_com_itcoursee_core_recommend_top_proto_init() {
	if File_com_itcoursee_core_recommend_top_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_recommend_top_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopReq); i {
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
		file_com_itcoursee_core_recommend_top_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopRsp); i {
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
			RawDescriptor: file_com_itcoursee_core_recommend_top_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_recommend_top_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_recommend_top_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_recommend_top_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_recommend_top_proto = out.File
	file_com_itcoursee_core_recommend_top_proto_rawDesc = nil
	file_com_itcoursee_core_recommend_top_proto_goTypes = nil
	file_com_itcoursee_core_recommend_top_proto_depIdxs = nil
}
