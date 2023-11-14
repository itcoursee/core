// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: com/itcoursee/core/vo/recommend.proto

package vo

import (
	kernel "github.com/itcoursee/core/kernel"
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

// 推荐
type Recommend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 标识
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 目标
	Target int64 `protobuf:"varint,5,opt,name=target,proto3" json:"target,omitempty"`
	// 类型
	Type kernel.RecommendType `protobuf:"varint,6,opt,name=type,proto3,enum=com.itcoursee.core.kernel.RecommendType" json:"type,omitempty"`
	// 封面
	// @gotags: str:"-"
	Cover *File `protobuf:"bytes,10,opt,name=cover,proto3" json:"cover,omitempty" str:"-"`
}

func (x *Recommend) Reset() {
	*x = Recommend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_vo_recommend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recommend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recommend) ProtoMessage() {}

func (x *Recommend) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_vo_recommend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recommend.ProtoReflect.Descriptor instead.
func (*Recommend) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_vo_recommend_proto_rawDescGZIP(), []int{0}
}

func (x *Recommend) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Recommend) GetTarget() int64 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *Recommend) GetType() kernel.RecommendType {
	if x != nil {
		return x.Type
	}
	return kernel.RecommendType(0)
}

func (x *Recommend) GetCover() *File {
	if x != nil {
		return x.Cover
	}
	return nil
}

var File_com_itcoursee_core_vo_recommend_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_vo_recommend_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x1a, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa4, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x3a, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f,
	0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f,
	0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_vo_recommend_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_vo_recommend_proto_rawDescData = file_com_itcoursee_core_vo_recommend_proto_rawDesc
)

func file_com_itcoursee_core_vo_recommend_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_vo_recommend_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_vo_recommend_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_vo_recommend_proto_rawDescData)
	})
	return file_com_itcoursee_core_vo_recommend_proto_rawDescData
}

var file_com_itcoursee_core_vo_recommend_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_itcoursee_core_vo_recommend_proto_goTypes = []interface{}{
	(*Recommend)(nil),         // 0: com.itcoursee.core.vo.Recommend
	(kernel.RecommendType)(0), // 1: com.itcoursee.core.kernel.RecommendType
	(*File)(nil),              // 2: com.itcoursee.core.vo.File
}
var file_com_itcoursee_core_vo_recommend_proto_depIdxs = []int32{
	1, // 0: com.itcoursee.core.vo.Recommend.type:type_name -> com.itcoursee.core.kernel.RecommendType
	2, // 1: com.itcoursee.core.vo.Recommend.cover:type_name -> com.itcoursee.core.vo.File
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_vo_recommend_proto_init() }
func file_com_itcoursee_core_vo_recommend_proto_init() {
	if File_com_itcoursee_core_vo_recommend_proto != nil {
		return
	}
	file_com_itcoursee_core_vo_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_vo_recommend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recommend); i {
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
			RawDescriptor: file_com_itcoursee_core_vo_recommend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_vo_recommend_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_vo_recommend_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_vo_recommend_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_vo_recommend_proto = out.File
	file_com_itcoursee_core_vo_recommend_proto_rawDesc = nil
	file_com_itcoursee_core_vo_recommend_proto_goTypes = nil
	file_com_itcoursee_core_vo_recommend_proto_depIdxs = nil
}
