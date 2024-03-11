// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: takcontrol.proto

package cotproto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TAK Protocol control message
// This specifies to a recipient what versions
// of protocol elements this sender supports during
// decoding.
type TakControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lowest TAK protocol version supported
	// If not filled in (reads as 0), version 1 is assumed
	MinProtoVersion uint32 `protobuf:"varint,1,opt,name=minProtoVersion,proto3" json:"minProtoVersion,omitempty"`
	// Highest TAK protocol version supported
	// If not filled in (reads as 0), version 1 is assumed
	MaxProtoVersion uint32 `protobuf:"varint,2,opt,name=maxProtoVersion,proto3" json:"maxProtoVersion,omitempty"`
}

func (x *TakControl) Reset() {
	*x = TakControl{}

	if protoimpl.UnsafeEnabled {
		mi := &file_takcontrol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakControl) ProtoMessage() {}

func (x *TakControl) ProtoReflect() protoreflect.Message {
	mi := &file_takcontrol_proto_msgTypes[0]

	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}

		return ms
	}

	return mi.MessageOf(x)
}

// Deprecated: Use TakControl.ProtoReflect.Descriptor instead.
func (*TakControl) Descriptor() ([]byte, []int) {
	return file_takcontrol_proto_rawDescGZIP(), []int{0}
}

func (x *TakControl) GetMinProtoVersion() uint32 {
	if x != nil {
		return x.MinProtoVersion
	}

	return 0
}

func (x *TakControl) GetMaxProtoVersion() uint32 {
	if x != nil {
		return x.MaxProtoVersion
	}

	return 0
}

var File_takcontrol_proto protoreflect.FileDescriptor

var file_takcontrol_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x61, 0x6b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x60, 0x0a, 0x0a, 0x54, 0x61, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x61,
	0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x26, 0x48, 0x03, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x64, 0x75, 0x64, 0x6b, 0x6f, 0x76, 0x2f, 0x67, 0x6f, 0x61,
	0x74, 0x61, 0x6b, 0x2f, 0x63, 0x6f, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_takcontrol_proto_rawDescOnce sync.Once
	file_takcontrol_proto_rawDescData = file_takcontrol_proto_rawDesc
)

func file_takcontrol_proto_rawDescGZIP() []byte {
	file_takcontrol_proto_rawDescOnce.Do(func() {
		file_takcontrol_proto_rawDescData = protoimpl.X.CompressGZIP(file_takcontrol_proto_rawDescData)
	})

	return file_takcontrol_proto_rawDescData
}

var file_takcontrol_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_takcontrol_proto_goTypes = []interface{}{
	(*TakControl)(nil), // 0: TakControl
}
var file_takcontrol_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_takcontrol_proto_init() }
func file_takcontrol_proto_init() {
	if File_takcontrol_proto != nil {
		return
	}

	if !protoimpl.UnsafeEnabled {
		file_takcontrol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakControl); i {
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
			RawDescriptor: file_takcontrol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_takcontrol_proto_goTypes,
		DependencyIndexes: file_takcontrol_proto_depIdxs,
		MessageInfos:      file_takcontrol_proto_msgTypes,
	}.Build()
	File_takcontrol_proto = out.File
	file_takcontrol_proto_rawDesc = nil
	file_takcontrol_proto_goTypes = nil
	file_takcontrol_proto_depIdxs = nil
}
