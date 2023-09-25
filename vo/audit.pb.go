// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: com/itcoursee/core/vo/audit.proto

package vo

import (
	audit "github.com/itcoursee/core/audit"
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

// 审计
type Audit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 名称
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// 状态
	Status audit.Status `protobuf:"varint,6,opt,name=status,proto3,enum=com.itcoursee.core.audit.Status" json:"status,omitempty"`
	// 当前审核节点数
	Current int32 `protobuf:"varint,7,opt,name=current,proto3" json:"current,omitempty"`
	// 审核者
	Reviewer *Reviewer `protobuf:"bytes,8,opt,name=reviewer,proto3" json:"reviewer,omitempty"`
	// 提交记录
	Commit *Commit `protobuf:"bytes,9,opt,name=commit,proto3" json:"commit,omitempty"`
	// 审核
	// ! 通过字段 fields=review 返回
	// @gotags: str:"-"
	Review *Review `protobuf:"bytes,15,opt,name=review,proto3" json:"review,omitempty" str:"-"`
}

func (x *Audit) Reset() {
	*x = Audit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_core_vo_audit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audit) ProtoMessage() {}

func (x *Audit) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_core_vo_audit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audit.ProtoReflect.Descriptor instead.
func (*Audit) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_core_vo_audit_proto_rawDescGZIP(), []int{0}
}

func (x *Audit) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Audit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Audit) GetStatus() audit.Status {
	if x != nil {
		return x.Status
	}
	return audit.Status(0)
}

func (x *Audit) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *Audit) GetReviewer() *Reviewer {
	if x != nil {
		return x.Reviewer
	}
	return nil
}

func (x *Audit) GetCommit() *Commit {
	if x != nil {
		return x.Commit
	}
	return nil
}

func (x *Audit) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

var File_com_itcoursee_core_vo_audit_proto protoreflect.FileDescriptor

var file_com_itcoursee_core_vo_audit_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f,
	0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaa, 0x02, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x65, 0x72, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x12, 0x35,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x2e, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x42, 0x3a, 0x0a, 0x15,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_core_vo_audit_proto_rawDescOnce sync.Once
	file_com_itcoursee_core_vo_audit_proto_rawDescData = file_com_itcoursee_core_vo_audit_proto_rawDesc
)

func file_com_itcoursee_core_vo_audit_proto_rawDescGZIP() []byte {
	file_com_itcoursee_core_vo_audit_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_core_vo_audit_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_core_vo_audit_proto_rawDescData)
	})
	return file_com_itcoursee_core_vo_audit_proto_rawDescData
}

var file_com_itcoursee_core_vo_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_itcoursee_core_vo_audit_proto_goTypes = []interface{}{
	(*Audit)(nil),     // 0: com.itcoursee.core.vo.Audit
	(audit.Status)(0), // 1: com.itcoursee.core.audit.Status
	(*Reviewer)(nil),  // 2: com.itcoursee.core.vo.Reviewer
	(*Commit)(nil),    // 3: com.itcoursee.core.vo.Commit
	(*Review)(nil),    // 4: com.itcoursee.core.vo.Review
}
var file_com_itcoursee_core_vo_audit_proto_depIdxs = []int32{
	1, // 0: com.itcoursee.core.vo.Audit.status:type_name -> com.itcoursee.core.audit.Status
	2, // 1: com.itcoursee.core.vo.Audit.reviewer:type_name -> com.itcoursee.core.vo.Reviewer
	3, // 2: com.itcoursee.core.vo.Audit.commit:type_name -> com.itcoursee.core.vo.Commit
	4, // 3: com.itcoursee.core.vo.Audit.review:type_name -> com.itcoursee.core.vo.Review
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_itcoursee_core_vo_audit_proto_init() }
func file_com_itcoursee_core_vo_audit_proto_init() {
	if File_com_itcoursee_core_vo_audit_proto != nil {
		return
	}
	file_com_itcoursee_core_vo_commit_proto_init()
	file_com_itcoursee_core_vo_review_proto_init()
	file_com_itcoursee_core_vo_reviewer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_core_vo_audit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audit); i {
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
			RawDescriptor: file_com_itcoursee_core_vo_audit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_core_vo_audit_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_core_vo_audit_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_core_vo_audit_proto_msgTypes,
	}.Build()
	File_com_itcoursee_core_vo_audit_proto = out.File
	file_com_itcoursee_core_vo_audit_proto_rawDesc = nil
	file_com_itcoursee_core_vo_audit_proto_goTypes = nil
	file_com_itcoursee_core_vo_audit_proto_depIdxs = nil
}
