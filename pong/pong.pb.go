// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: ping-pong/pong/pong.proto

package pong

import (
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

type PongRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PongRequest) Reset() {
	*x = PongRequest{}
	mi := &file_ping_pong_pong_pong_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongRequest) ProtoMessage() {}

func (x *PongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ping_pong_pong_pong_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongRequest.ProtoReflect.Descriptor instead.
func (*PongRequest) Descriptor() ([]byte, []int) {
	return file_ping_pong_pong_pong_proto_rawDescGZIP(), []int{0}
}

func (x *PongRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PongResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      string                 `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PongResponse) Reset() {
	*x = PongResponse{}
	mi := &file_ping_pong_pong_pong_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongResponse) ProtoMessage() {}

func (x *PongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ping_pong_pong_pong_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongResponse.ProtoReflect.Descriptor instead.
func (*PongResponse) Descriptor() ([]byte, []int) {
	return file_ping_pong_pong_pong_proto_rawDescGZIP(), []int{1}
}

func (x *PongResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_ping_pong_pong_pong_proto protoreflect.FileDescriptor

var file_ping_pong_pong_pong_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x6f, 0x6e, 0x67, 0x2f, 0x70, 0x6f, 0x6e, 0x67,
	0x2f, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x50, 0x6f,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x3f, 0x0a, 0x0b, 0x50, 0x6f, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6e, 0x67,
	0x12, 0x11, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x70, 0x69, 0x6e, 0x67, 0x2d,
	0x70, 0x6f, 0x6e, 0x67, 0x2f, 0x70, 0x6f, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3b, 0x70, 0x6f, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ping_pong_pong_pong_proto_rawDescOnce sync.Once
	file_ping_pong_pong_pong_proto_rawDescData = file_ping_pong_pong_pong_proto_rawDesc
)

func file_ping_pong_pong_pong_proto_rawDescGZIP() []byte {
	file_ping_pong_pong_pong_proto_rawDescOnce.Do(func() {
		file_ping_pong_pong_pong_proto_rawDescData = protoimpl.X.CompressGZIP(file_ping_pong_pong_pong_proto_rawDescData)
	})
	return file_ping_pong_pong_pong_proto_rawDescData
}

var file_ping_pong_pong_pong_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ping_pong_pong_pong_proto_goTypes = []any{
	(*PongRequest)(nil),  // 0: pong.PongRequest
	(*PongResponse)(nil), // 1: pong.PongResponse
}
var file_ping_pong_pong_pong_proto_depIdxs = []int32{
	0, // 0: pong.PongService.GetPong:input_type -> pong.PongRequest
	1, // 1: pong.PongService.GetPong:output_type -> pong.PongResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ping_pong_pong_pong_proto_init() }
func file_ping_pong_pong_pong_proto_init() {
	if File_ping_pong_pong_pong_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ping_pong_pong_pong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ping_pong_pong_pong_proto_goTypes,
		DependencyIndexes: file_ping_pong_pong_pong_proto_depIdxs,
		MessageInfos:      file_ping_pong_pong_pong_proto_msgTypes,
	}.Build()
	File_ping_pong_pong_pong_proto = out.File
	file_ping_pong_pong_pong_proto_rawDesc = nil
	file_ping_pong_pong_pong_proto_goTypes = nil
	file_ping_pong_pong_pong_proto_depIdxs = nil
}
