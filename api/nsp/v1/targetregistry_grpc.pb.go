// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: api/nsp/v1/targetregistry.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TargetRegistryClient is the client API for TargetRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TargetRegistryClient interface {
	Register(ctx context.Context, in *Target, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Unregister(ctx context.Context, in *Target, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Watch(ctx context.Context, in *Target, opts ...grpc.CallOption) (TargetRegistry_WatchClient, error)
}

type targetRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewTargetRegistryClient(cc grpc.ClientConnInterface) TargetRegistryClient {
	return &targetRegistryClient{cc}
}

func (c *targetRegistryClient) Register(ctx context.Context, in *Target, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/nsp.v1.TargetRegistry/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetRegistryClient) Unregister(ctx context.Context, in *Target, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/nsp.v1.TargetRegistry/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetRegistryClient) Watch(ctx context.Context, in *Target, opts ...grpc.CallOption) (TargetRegistry_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &TargetRegistry_ServiceDesc.Streams[0], "/nsp.v1.TargetRegistry/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &targetRegistryWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TargetRegistry_WatchClient interface {
	Recv() (*TargetResponse, error)
	grpc.ClientStream
}

type targetRegistryWatchClient struct {
	grpc.ClientStream
}

func (x *targetRegistryWatchClient) Recv() (*TargetResponse, error) {
	m := new(TargetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TargetRegistryServer is the server API for TargetRegistry service.
// All implementations must embed UnimplementedTargetRegistryServer
// for forward compatibility
type TargetRegistryServer interface {
	Register(context.Context, *Target) (*emptypb.Empty, error)
	Unregister(context.Context, *Target) (*emptypb.Empty, error)
	Watch(*Target, TargetRegistry_WatchServer) error
	mustEmbedUnimplementedTargetRegistryServer()
}

// UnimplementedTargetRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedTargetRegistryServer struct {
}

func (UnimplementedTargetRegistryServer) Register(context.Context, *Target) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedTargetRegistryServer) Unregister(context.Context, *Target) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (UnimplementedTargetRegistryServer) Watch(*Target, TargetRegistry_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedTargetRegistryServer) mustEmbedUnimplementedTargetRegistryServer() {}

// UnsafeTargetRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TargetRegistryServer will
// result in compilation errors.
type UnsafeTargetRegistryServer interface {
	mustEmbedUnimplementedTargetRegistryServer()
}

func RegisterTargetRegistryServer(s grpc.ServiceRegistrar, srv TargetRegistryServer) {
	s.RegisterService(&TargetRegistry_ServiceDesc, srv)
}

func _TargetRegistry_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetRegistryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsp.v1.TargetRegistry/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetRegistryServer).Register(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetRegistry_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetRegistryServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsp.v1.TargetRegistry/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetRegistryServer).Unregister(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetRegistry_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Target)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TargetRegistryServer).Watch(m, &targetRegistryWatchServer{stream})
}

type TargetRegistry_WatchServer interface {
	Send(*TargetResponse) error
	grpc.ServerStream
}

type targetRegistryWatchServer struct {
	grpc.ServerStream
}

func (x *targetRegistryWatchServer) Send(m *TargetResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TargetRegistry_ServiceDesc is the grpc.ServiceDesc for TargetRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TargetRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nsp.v1.TargetRegistry",
	HandlerType: (*TargetRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _TargetRegistry_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _TargetRegistry_Unregister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _TargetRegistry_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/nsp/v1/targetregistry.proto",
}
