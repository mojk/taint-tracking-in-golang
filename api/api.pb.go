// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PingMessage struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (m *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(m, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type Request struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Request) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type Response struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*PingMessage)(nil), "api.PingMessage")
	proto.RegisterType((*Request)(nil), "api.Request")
	proto.RegisterType((*Response)(nil), "api.Response")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xcf, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x06, 0xe0, 0xc6, 0x6a, 0x6d, 0xc7, 0x0a, 0x92, 0x83, 0x94, 0x9e, 0x24, 0x20, 0x28, 0x4a,
	0x0f, 0xd5, 0x17, 0xe8, 0xcd, 0x8b, 0x22, 0xd9, 0x27, 0x48, 0xc8, 0x10, 0x02, 0xdd, 0x26, 0x9b,
	0xa4, 0x87, 0x7d, 0xfb, 0xa5, 0x6d, 0x58, 0x96, 0x3d, 0x7e, 0xcc, 0x0c, 0xff, 0x3f, 0x50, 0x09,
	0x67, 0x3a, 0xe7, 0x6d, 0xb4, 0x34, 0x17, 0xce, 0xb0, 0x77, 0x78, 0xf8, 0x37, 0x93, 0xfe, 0xc5,
	0x10, 0x84, 0x46, 0xda, 0x42, 0xa9, 0x3d, 0x62, 0x34, 0x93, 0x6e, 0xc8, 0x0b, 0x79, 0xab, 0xf8,
	0xd9, 0xec, 0x15, 0xee, 0x39, 0x1e, 0x66, 0x0c, 0x91, 0xd6, 0x40, 0xc4, 0x3a, 0xbf, 0xe3, 0x44,
	0x2c, 0x92, 0xcd, 0xcd, 0x26, 0xc9, 0x18, 0x94, 0x1c, 0x83, 0xb3, 0x53, 0x40, 0xfa, 0x0c, 0x85,
	0xc7, 0x30, 0x8f, 0x31, 0x2d, 0x27, 0xf5, 0xdf, 0x70, 0xbb, 0xa4, 0xd2, 0x4f, 0x28, 0x77, 0xe2,
	0xf8, 0x83, 0xe3, 0x68, 0x69, 0xdd, 0x2d, 0xd5, 0x52, 0x42, 0xfb, 0xb4, 0xea, 0xa2, 0x1a, 0xcb,
	0xfa, 0x1e, 0xf2, 0x41, 0x29, 0xfa, 0x01, 0x30, 0x28, 0xf5, 0x37, 0xef, 0x25, 0xfa, 0x70, 0x75,
	0xf6, 0x98, 0xb4, 0xe5, 0xb3, 0x4c, 0x16, 0xeb, 0xaf, 0x5f, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x3d, 0x51, 0xa1, 0x44, 0xf8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingClient is the client API for Ping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingClient interface {
	SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PingMessage, error)
}

type pingClient struct {
	cc *grpc.ClientConn
}

func NewPingClient(cc *grpc.ClientConn) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PingMessage, error) {
	out := new(PingMessage)
	err := c.cc.Invoke(ctx, "/api.Ping/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServer is the server API for Ping service.
type PingServer interface {
	SayHello(context.Context, *Request) (*PingMessage, error)
}

// UnimplementedPingServer can be embedded to have forward compatible implementations.
type UnimplementedPingServer struct {
}

func (*UnimplementedPingServer) SayHello(ctx context.Context, req *Request) (*PingMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterPingServer(s *grpc.Server, srv PingServer) {
	s.RegisterService(&_Ping_serviceDesc, srv)
}

func _Ping_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Ping/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).SayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ping_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Ping_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// AddClient is the client API for Add service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddClient interface {
	AddNumbers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type addClient struct {
	cc *grpc.ClientConn
}

func NewAddClient(cc *grpc.ClientConn) AddClient {
	return &addClient{cc}
}

func (c *addClient) AddNumbers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Add/AddNumbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServer is the server API for Add service.
type AddServer interface {
	AddNumbers(context.Context, *Request) (*Response, error)
}

// UnimplementedAddServer can be embedded to have forward compatible implementations.
type UnimplementedAddServer struct {
}

func (*UnimplementedAddServer) AddNumbers(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNumbers not implemented")
}

func RegisterAddServer(s *grpc.Server, srv AddServer) {
	s.RegisterService(&_Add_serviceDesc, srv)
}

func _Add_AddNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).AddNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Add/AddNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).AddNumbers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Add_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Add",
	HandlerType: (*AddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNumbers",
			Handler:    _Add_AddNumbers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}