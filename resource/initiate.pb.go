// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: com/itcoursee/core/resource/initiate.proto

package resource

import (
	kernel "github.com/itcoursee/core/kernel"
	core "gitlab.com/ruijc/storage/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 初始化分块上传请求
type InitiateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 媒体类型
	// @gotags: validate:"required"
	Mime string `protobuf:"bytes,4,opt,name=mime,proto3" json:"mime,omitempty" validate:"required"`
	// 文件名
	// @gotags: validate:"required,max=128"
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty" validate:"required,max=128"`
	// 类型
	// @gotags: validate:"required"
	Type kernel.Filetype `protobuf:"varint,6,opt,name=type,proto3,enum=com.itcoursee.core.kernel.Filetype" json:"type,omitempty" validate:"required"`
	// 分组
	Suite int64 `protobuf:"varint,7,opt,name=suite,proto3" json:"suite,omitempty"`
	// 分片数量
	// @gotags: validate:"required"
	Parts int32 `protobuf:"varint,9,opt,name=parts,proto3" json:"parts,omitempty" validate:"required"`
	// 分片起始值
	// @gotags: default:"1"
	Start int32 `protobuf:"varint,10,opt,name=start,proto3" json:"start,omitempty" default:"1"`
	// 过期时间
	// 默认为5个小时
	// 超过该时间，上传地址将不可用
	// @gotags: default:"{'seconds': 18000}"
	Expires *durationpb.Duration `protobuf:"bytes,15,opt,name=expires,proto3" json:"expires,omitempty" default:"{'seconds': 18000}"`
}

func (x *InitiateReq) Reset() {
	*x = InitiateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_resource_initiate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateReq) ProtoMessage() {}

func (x *InitiateReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_resource_initiate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateReq.ProtoReflect.Descriptor instead.
func (*InitiateReq) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_resource_initiate_proto_rawDescGZIP(), []int{0}
}

func (x *InitiateReq) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *InitiateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InitiateReq) GetType() kernel.Filetype {
	if x != nil {
		return x.Type
	}
	return kernel.Filetype(0)
}

func (x *InitiateReq) GetSuite() int64 {
	if x != nil {
		return x.Suite
	}
	return 0
}

func (x *InitiateReq) GetParts() int32 {
	if x != nil {
		return x.Parts
	}
	return 0
}

func (x *InitiateReq) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *InitiateReq) GetExpires() *durationpb.Duration {
	if x != nil {
		return x.Expires
	}
	return nil
}

// 初始化分块上传返回
type InitiateRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	// 后续所有操作都使用此编号作为唯一标识
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 上传地址
	Urls []*core.Url `protobuf:"bytes,4,rep,name=urls,proto3" json:"urls,omitempty"`
	// 过期时间
	Expired *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=expired,proto3" json:"expired,omitempty"`
}

func (x *InitiateRsp) Reset() {
	*x = InitiateRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_resource_initiate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiateRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateRsp) ProtoMessage() {}

func (x *InitiateRsp) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_resource_initiate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateRsp.ProtoReflect.Descriptor instead.
func (*InitiateRsp) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_resource_initiate_proto_rawDescGZIP(), []int{1}
}

func (x *InitiateRsp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InitiateRsp) GetUrls() []*core.Url {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *InitiateRsp) GetExpired() *timestamppb.Timestamp {
	if x != nil {
		return x.Expired
	}
	return nil
}

var File_com_itcoursee_core_resource_initiate_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_resource_initiate_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x75, 0x69, 0x6a, 0x63, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x72, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x75, 0x69, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x75, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61,
	0x72, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x22, 0x84,
	0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f,
	0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x72, 0x75, 0x69, 0x6a, 0x63, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12,
	0x34, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x42, 0x4c, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_resource_initiate_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_resource_initiate_proto_rawDescData = file_com_itcoursee_core_resource_initiate_proto_rawDesc
)

func file_com_itcoursee_core_resource_initiate_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_resource_initiate_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_resource_initiate_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_resource_initiate_proto_rawDescData)
	})
	return file_com_itcoursee_core_resource_initiate_proto_rawDescData
}

var file_com_itcoursee_core_resource_initiate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_itcoursee_core_resource_initiate_proto_goTypes = []interface{}{
	(*InitiateReq)(nil),           // 0: com.itcoursee.core.resource.InitiateReq
	(*InitiateRsp)(nil),           // 1: com.itcoursee.core.resource.InitiateRsp
	(kernel.Filetype)(0),          // 2: com.itcoursee.core.kernel.Filetype
	(*durationpb.Duration)(nil),   // 3: google.protobuf.Duration
	(*core.Url)(nil),              // 4: com.ruijc.storage.core.Url
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_com_itcoursee_core_resource_initiate_proto_depIdxs = []int32{
	2, // 0: com.itcoursee.core.resource.InitiateReq.type:type_name -> com.itcoursee.core.kernel.Filetype
	3, // 1: com.itcoursee.core.resource.InitiateReq.expires:type_name -> google.protobuf.Duration
	4, // 2: com.itcoursee.core.resource.InitiateRsp.urls:type_name -> com.ruijc.storage.core.Url
	5, // 3: com.itcoursee.core.resource.InitiateRsp.expired:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_resource_initiate_proto_init() }
func file_com_itcoursee_core_resource_initiate_proto_init() {
	if File_com_itcoursee_core_resource_initiate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_resource_initiate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiateReq); i {
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
		file_com_itcoursee_core_resource_initiate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiateRsp); i {
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
			RawDescriptor: file_com_itcoursee_core_resource_initiate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_resource_initiate_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_resource_initiate_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_resource_initiate_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_resource_initiate_proto = out.File
	file_com_itcoursee_core_resource_initiate_proto_rawDesc = nil
	file_com_itcoursee_core_resource_initiate_proto_goTypes = nil
	file_com_itcoursee_core_resource_initiate_proto_depIdxs = nil
}
