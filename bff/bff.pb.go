// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: ping-pong/bff/bff.proto

package bff

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

type BFFRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BFFRequest) Reset() {
	*x = BFFRequest{}
	mi := &file_ping_pong_bff_bff_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BFFRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFFRequest) ProtoMessage() {}

func (x *BFFRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ping_pong_bff_bff_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFFRequest.ProtoReflect.Descriptor instead.
func (*BFFRequest) Descriptor() ([]byte, []int) {
	return file_ping_pong_bff_bff_proto_rawDescGZIP(), []int{0}
}

func (x *BFFRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BFFResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      string                 `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BFFResponse) Reset() {
	*x = BFFResponse{}
	mi := &file_ping_pong_bff_bff_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BFFResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFFResponse) ProtoMessage() {}

func (x *BFFResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ping_pong_bff_bff_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFFResponse.ProtoReflect.Descriptor instead.
func (*BFFResponse) Descriptor() ([]byte, []int) {
	return file_ping_pong_bff_bff_proto_rawDescGZIP(), []int{1}
}

func (x *BFFResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_ping_pong_bff_bff_proto protoreflect.FileDescriptor

var file_ping_pong_bff_bff_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x6f, 0x6e, 0x67, 0x2f, 0x62, 0x66, 0x66, 0x2f,
	0x62, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x62, 0x66, 0x66, 0x22, 0x26,
	0x0a, 0x0a, 0x42, 0x46, 0x46, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0b, 0x42, 0x46, 0x46, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x3e, 0x0a, 0x0a, 0x42, 0x46, 0x46, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x0b, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0f,
	0x2e, 0x62, 0x66, 0x66, 0x2e, 0x42, 0x46, 0x46, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x42, 0x46, 0x46, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x6f, 0x6e, 0x67, 0x2f, 0x62,
	0x66, 0x66, 0x3b, 0x62, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ping_pong_bff_bff_proto_rawDescOnce sync.Once
	file_ping_pong_bff_bff_proto_rawDescData = file_ping_pong_bff_bff_proto_rawDesc
)

func file_ping_pong_bff_bff_proto_rawDescGZIP() []byte {
	file_ping_pong_bff_bff_proto_rawDescOnce.Do(func() {
		file_ping_pong_bff_bff_proto_rawDescData = protoimpl.X.CompressGZIP(file_ping_pong_bff_bff_proto_rawDescData)
	})
	return file_ping_pong_bff_bff_proto_rawDescData
}

var file_ping_pong_bff_bff_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ping_pong_bff_bff_proto_goTypes = []any{
	(*BFFRequest)(nil),  // 0: bff.BFFRequest
	(*BFFResponse)(nil), // 1: bff.BFFResponse
}
var file_ping_pong_bff_bff_proto_depIdxs = []int32{
	0, // 0: bff.BFFService.ForwardPing:input_type -> bff.BFFRequest
	1, // 1: bff.BFFService.ForwardPing:output_type -> bff.BFFResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ping_pong_bff_bff_proto_init() }
func file_ping_pong_bff_bff_proto_init() {
	if File_ping_pong_bff_bff_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ping_pong_bff_bff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ping_pong_bff_bff_proto_goTypes,
		DependencyIndexes: file_ping_pong_bff_bff_proto_depIdxs,
		MessageInfos:      file_ping_pong_bff_bff_proto_msgTypes,
	}.Build()
	File_ping_pong_bff_bff_proto = out.File
	file_ping_pong_bff_bff_proto_rawDesc = nil
	file_ping_pong_bff_bff_proto_goTypes = nil
	file_ping_pong_bff_bff_proto_depIdxs = nil
}
