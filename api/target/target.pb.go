// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: api/target/target.proto

package target

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Conduit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkServiceName string `protobuf:"bytes,1,opt,name=networkServiceName,proto3" json:"networkServiceName,omitempty"`
	Trench             string `protobuf:"bytes,2,opt,name=trench,proto3" json:"trench,omitempty"`
}

func (x *Conduit) Reset() {
	*x = Conduit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_target_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conduit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conduit) ProtoMessage() {}

func (x *Conduit) ProtoReflect() protoreflect.Message {
	mi := &file_api_target_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conduit.ProtoReflect.Descriptor instead.
func (*Conduit) Descriptor() ([]byte, []int) {
	return file_api_target_target_proto_rawDescGZIP(), []int{0}
}

func (x *Conduit) GetNetworkServiceName() string {
	if x != nil {
		return x.NetworkServiceName
	}
	return ""
}

func (x *Conduit) GetTrench() string {
	if x != nil {
		return x.Trench
	}
	return ""
}

type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conduit *Conduit `protobuf:"bytes,1,opt,name=conduit,proto3" json:"conduit,omitempty"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_target_target_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_api_target_target_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_api_target_target_proto_rawDescGZIP(), []int{1}
}

func (x *Stream) GetConduit() *Conduit {
	if x != nil {
		return x.Conduit
	}
	return nil
}

var File_api_target_target_proto protoreflect.FileDescriptor

var file_api_target_target_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x65,
	0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x65, 0x6e, 0x63,
	0x68, 0x22, 0x33, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x29, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x32, 0xe3, 0x01, 0x0a, 0x0a, 0x41, 0x6d, 0x62, 0x61, 0x73,
	0x73, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x12, 0x0f, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x12, 0x0e, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x72, 0x64, 0x69,
	0x78, 0x2f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_target_target_proto_rawDescOnce sync.Once
	file_api_target_target_proto_rawDescData = file_api_target_target_proto_rawDesc
)

func file_api_target_target_proto_rawDescGZIP() []byte {
	file_api_target_target_proto_rawDescOnce.Do(func() {
		file_api_target_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_target_target_proto_rawDescData)
	})
	return file_api_target_target_proto_rawDescData
}

var file_api_target_target_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_target_target_proto_goTypes = []interface{}{
	(*Conduit)(nil),     // 0: target.Conduit
	(*Stream)(nil),      // 1: target.Stream
	(*empty.Empty)(nil), // 2: google.protobuf.Empty
}
var file_api_target_target_proto_depIdxs = []int32{
	0, // 0: target.Stream.conduit:type_name -> target.Conduit
	0, // 1: target.Ambassador.Connect:input_type -> target.Conduit
	0, // 2: target.Ambassador.Disconnect:input_type -> target.Conduit
	1, // 3: target.Ambassador.Request:input_type -> target.Stream
	1, // 4: target.Ambassador.Close:input_type -> target.Stream
	2, // 5: target.Ambassador.Connect:output_type -> google.protobuf.Empty
	2, // 6: target.Ambassador.Disconnect:output_type -> google.protobuf.Empty
	2, // 7: target.Ambassador.Request:output_type -> google.protobuf.Empty
	2, // 8: target.Ambassador.Close:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_target_target_proto_init() }
func file_api_target_target_proto_init() {
	if File_api_target_target_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_target_target_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conduit); i {
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
		file_api_target_target_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
			RawDescriptor: file_api_target_target_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_target_target_proto_goTypes,
		DependencyIndexes: file_api_target_target_proto_depIdxs,
		MessageInfos:      file_api_target_target_proto_msgTypes,
	}.Build()
	File_api_target_target_proto = out.File
	file_api_target_target_proto_rawDesc = nil
	file_api_target_target_proto_goTypes = nil
	file_api_target_target_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AmbassadorClient is the client API for Ambassador service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AmbassadorClient interface {
	Connect(ctx context.Context, in *Conduit, opts ...grpc.CallOption) (*empty.Empty, error)
	Disconnect(ctx context.Context, in *Conduit, opts ...grpc.CallOption) (*empty.Empty, error)
	Request(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*empty.Empty, error)
	Close(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*empty.Empty, error)
}

type ambassadorClient struct {
	cc grpc.ClientConnInterface
}

func NewAmbassadorClient(cc grpc.ClientConnInterface) AmbassadorClient {
	return &ambassadorClient{cc}
}

func (c *ambassadorClient) Connect(ctx context.Context, in *Conduit, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/target.Ambassador/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambassadorClient) Disconnect(ctx context.Context, in *Conduit, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/target.Ambassador/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambassadorClient) Request(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/target.Ambassador/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambassadorClient) Close(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/target.Ambassador/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AmbassadorServer is the server API for Ambassador service.
type AmbassadorServer interface {
	Connect(context.Context, *Conduit) (*empty.Empty, error)
	Disconnect(context.Context, *Conduit) (*empty.Empty, error)
	Request(context.Context, *Stream) (*empty.Empty, error)
	Close(context.Context, *Stream) (*empty.Empty, error)
}

// UnimplementedAmbassadorServer can be embedded to have forward compatible implementations.
type UnimplementedAmbassadorServer struct {
}

func (*UnimplementedAmbassadorServer) Connect(context.Context, *Conduit) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedAmbassadorServer) Disconnect(context.Context, *Conduit) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (*UnimplementedAmbassadorServer) Request(context.Context, *Stream) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (*UnimplementedAmbassadorServer) Close(context.Context, *Stream) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterAmbassadorServer(s *grpc.Server, srv AmbassadorServer) {
	s.RegisterService(&_Ambassador_serviceDesc, srv)
}

func _Ambassador_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conduit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbassadorServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/target.Ambassador/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbassadorServer).Connect(ctx, req.(*Conduit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambassador_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conduit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbassadorServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/target.Ambassador/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbassadorServer).Disconnect(ctx, req.(*Conduit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambassador_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stream)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbassadorServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/target.Ambassador/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbassadorServer).Request(ctx, req.(*Stream))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambassador_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stream)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbassadorServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/target.Ambassador/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbassadorServer).Close(ctx, req.(*Stream))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ambassador_serviceDesc = grpc.ServiceDesc{
	ServiceName: "target.Ambassador",
	HandlerType: (*AmbassadorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Ambassador_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _Ambassador_Disconnect_Handler,
		},
		{
			MethodName: "Request",
			Handler:    _Ambassador_Request_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Ambassador_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/target/target.proto",
}
