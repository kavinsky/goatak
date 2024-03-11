// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: track.proto

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

// All items are required unless otherwise noted!
// "required" means if they are missing on send, the conversion
// to the message format will be rejected and fall back to opaque
// XML representation
type Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speed  float64 `protobuf:"fixed64,1,opt,name=speed,proto3" json:"speed,omitempty"`   // speed=
	Course float64 `protobuf:"fixed64,2,opt,name=course,proto3" json:"course,omitempty"` // course=
}

func (x *Track) Reset() {
	*x = Track{}

	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[0]

	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}

		return ms
	}

	return mi.MessageOf(x)
}

// Deprecated: Use Track.ProtoReflect.Descriptor instead.
func (*Track) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{0}
}

func (x *Track) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}

	return 0
}

func (x *Track) GetCourse() float64 {
	if x != nil {
		return x.Course
	}

	return 0
}

var File_track_proto protoreflect.FileDescriptor

var file_track_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a,
	0x05, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x42, 0x26, 0x48, 0x03, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x64, 0x75, 0x64, 0x6b, 0x6f, 0x76, 0x2f, 0x67, 0x6f, 0x61,
	0x74, 0x61, 0x6b, 0x2f, 0x63, 0x6f, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_track_proto_rawDescOnce sync.Once
	file_track_proto_rawDescData = file_track_proto_rawDesc
)

func file_track_proto_rawDescGZIP() []byte {
	file_track_proto_rawDescOnce.Do(func() {
		file_track_proto_rawDescData = protoimpl.X.CompressGZIP(file_track_proto_rawDescData)
	})

	return file_track_proto_rawDescData
}

var file_track_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_track_proto_goTypes = []interface{}{
	(*Track)(nil), // 0: Track
}
var file_track_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_track_proto_init() }
func file_track_proto_init() {
	if File_track_proto != nil {
		return
	}

	if !protoimpl.UnsafeEnabled {
		file_track_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Track); i {
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
			RawDescriptor: file_track_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_track_proto_goTypes,
		DependencyIndexes: file_track_proto_depIdxs,
		MessageInfos:      file_track_proto_msgTypes,
	}.Build()
	File_track_proto = out.File
	file_track_proto_rawDesc = nil
	file_track_proto_goTypes = nil
	file_track_proto_depIdxs = nil
}
