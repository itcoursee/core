// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: com/itcoursee/core/vo/playback.proto

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

// 播放数据
type Playback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 地址
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// 当前分辨率
	Current kernel.Resolution `protobuf:"varint,4,opt,name=current,proto3,enum=com.itcoursee.core.kernel.Resolution" json:"current,omitempty"`
	// 所有的分辨率
	Resolutions []kernel.Resolution `protobuf:"varint,5,rep,packed,name=resolutions,proto3,enum=com.itcoursee.core.kernel.Resolution" json:"resolutions,omitempty"`
}

func (x *Playback) Reset() {
	*x = Playback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_vo_playback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Playback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Playback) ProtoMessage() {}

func (x *Playback) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_vo_playback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Playback.ProtoReflect.Descriptor instead.
func (*Playback) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_vo_playback_proto_rawDescGZIP(), []int{0}
}

func (x *Playback) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Playback) GetCurrent() kernel.Resolution {
	if x != nil {
		return x.Current
	}
	return kernel.Resolution(0)
}

func (x *Playback) GetResolutions() []kernel.Resolution {
	if x != nil {
		return x.Resolutions
	}
	return nil
}

var File_com_itcoursee_core_vo_playback_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_vo_playback_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x1a, 0x2a, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x08, 0x50, 0x6c,
	0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x3f, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x25,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x3a, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_vo_playback_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_vo_playback_proto_rawDescData = file_com_itcoursee_core_vo_playback_proto_rawDesc
)

func file_com_itcoursee_core_vo_playback_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_vo_playback_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_vo_playback_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_vo_playback_proto_rawDescData)
	})
	return file_com_itcoursee_core_vo_playback_proto_rawDescData
}

var file_com_itcoursee_core_vo_playback_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_itcoursee_core_vo_playback_proto_goTypes = []interface{}{
	(*Playback)(nil),       // 0: com.itcoursee.core.vo.Playback
	(kernel.Resolution)(0), // 1: com.itcoursee.core.kernel.Resolution
}
var file_com_itcoursee_core_vo_playback_proto_depIdxs = []int32{
	1, // 0: com.itcoursee.core.vo.Playback.current:type_name -> com.itcoursee.core.kernel.Resolution
	1, // 1: com.itcoursee.core.vo.Playback.resolutions:type_name -> com.itcoursee.core.kernel.Resolution
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_vo_playback_proto_init() }
func file_com_itcoursee_core_vo_playback_proto_init() {
	if File_com_itcoursee_core_vo_playback_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_vo_playback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Playback); i {
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
			RawDescriptor: file_com_itcoursee_core_vo_playback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_vo_playback_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_vo_playback_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_vo_playback_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_vo_playback_proto = out.File
	file_com_itcoursee_core_vo_playback_proto_rawDesc = nil
	file_com_itcoursee_core_vo_playback_proto_goTypes = nil
	file_com_itcoursee_core_vo_playback_proto_depIdxs = nil
}
