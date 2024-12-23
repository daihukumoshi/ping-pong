// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: ping-pong/bff/bff.proto

package bff

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BFFService_ForwardPing_FullMethodName = "/bff.BFFService/ForwardPing"
)

// BFFServiceClient is the client API for BFFService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BFFServiceClient interface {
	ForwardPing(ctx context.Context, in *BFFRequest, opts ...grpc.CallOption) (*BFFResponse, error)
}

type bFFServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBFFServiceClient(cc grpc.ClientConnInterface) BFFServiceClient {
	return &bFFServiceClient{cc}
}

func (c *bFFServiceClient) ForwardPing(ctx context.Context, in *BFFRequest, opts ...grpc.CallOption) (*BFFResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BFFResponse)
	err := c.cc.Invoke(ctx, BFFService_ForwardPing_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BFFServiceServer is the server API for BFFService service.
// All implementations must embed UnimplementedBFFServiceServer
// for forward compatibility.
type BFFServiceServer interface {
	ForwardPing(context.Context, *BFFRequest) (*BFFResponse, error)
	mustEmbedUnimplementedBFFServiceServer()
}

// UnimplementedBFFServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBFFServiceServer struct{}

func (UnimplementedBFFServiceServer) ForwardPing(context.Context, *BFFRequest) (*BFFResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardPing not implemented")
}
func (UnimplementedBFFServiceServer) mustEmbedUnimplementedBFFServiceServer() {}
func (UnimplementedBFFServiceServer) testEmbeddedByValue()                    {}

// UnsafeBFFServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BFFServiceServer will
// result in compilation errors.
type UnsafeBFFServiceServer interface {
	mustEmbedUnimplementedBFFServiceServer()
}

func RegisterBFFServiceServer(s grpc.ServiceRegistrar, srv BFFServiceServer) {
	// If the following call pancis, it indicates UnimplementedBFFServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BFFService_ServiceDesc, srv)
}

func _BFFService_ForwardPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BFFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServiceServer).ForwardPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFFService_ForwardPing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServiceServer).ForwardPing(ctx, req.(*BFFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BFFService_ServiceDesc is the grpc.ServiceDesc for BFFService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BFFService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bff.BFFService",
	HandlerType: (*BFFServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForwardPing",
			Handler:    _BFFService_ForwardPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping-pong/bff/bff.proto",
}
