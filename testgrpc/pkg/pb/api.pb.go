// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	GetValRequest
	GetValReply
	PutValRequest
	PutValReply
	Msg
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetValRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetValRequest) Reset()                    { *m = GetValRequest{} }
func (m *GetValRequest) String() string            { return proto.CompactTextString(m) }
func (*GetValRequest) ProtoMessage()               {}
func (*GetValRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetValRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetValReply struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
}

func (m *GetValReply) Reset()                    { *m = GetValReply{} }
func (m *GetValReply) String() string            { return proto.CompactTextString(m) }
func (*GetValReply) ProtoMessage()               {}
func (*GetValReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetValReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetValReply) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type PutValRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
}

func (m *PutValRequest) Reset()                    { *m = PutValRequest{} }
func (m *PutValRequest) String() string            { return proto.CompactTextString(m) }
func (*PutValRequest) ProtoMessage()               {}
func (*PutValRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PutValRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutValRequest) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type PutValReply struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *PutValReply) Reset()                    { *m = PutValReply{} }
func (m *PutValReply) String() string            { return proto.CompactTextString(m) }
func (*PutValReply) ProtoMessage()               {}
func (*PutValReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PutValReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type Msg struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Msg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*GetValRequest)(nil), "pb.GetValRequest")
	proto.RegisterType((*GetValReply)(nil), "pb.GetValReply")
	proto.RegisterType((*PutValRequest)(nil), "pb.PutValRequest")
	proto.RegisterType((*PutValReply)(nil), "pb.PutValReply")
	proto.RegisterType((*Msg)(nil), "pb.Msg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Demo service

type DemoClient interface {
	Get(ctx context.Context, in *GetValRequest, opts ...grpc.CallOption) (*GetValReply, error)
	Put(ctx context.Context, in *PutValRequest, opts ...grpc.CallOption) (*PutValReply, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Demo_StreamClient, error)
}

type demoClient struct {
	cc *grpc.ClientConn
}

func NewDemoClient(cc *grpc.ClientConn) DemoClient {
	return &demoClient{cc}
}

func (c *demoClient) Get(ctx context.Context, in *GetValRequest, opts ...grpc.CallOption) (*GetValReply, error) {
	out := new(GetValReply)
	err := grpc.Invoke(ctx, "/pb.Demo/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) Put(ctx context.Context, in *PutValRequest, opts ...grpc.CallOption) (*PutValReply, error) {
	out := new(PutValReply)
	err := grpc.Invoke(ctx, "/pb.Demo/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Demo_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Demo_serviceDesc.Streams[0], c.cc, "/pb.Demo/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoStreamClient{stream}
	return x, nil
}

type Demo_StreamClient interface {
	Send(*Msg) error
	Recv() (*Msg, error)
	grpc.ClientStream
}

type demoStreamClient struct {
	grpc.ClientStream
}

func (x *demoStreamClient) Send(m *Msg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demoStreamClient) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Demo service

type DemoServer interface {
	Get(context.Context, *GetValRequest) (*GetValReply, error)
	Put(context.Context, *PutValRequest) (*PutValReply, error)
	Stream(Demo_StreamServer) error
}

func RegisterDemoServer(s *grpc.Server, srv DemoServer) {
	s.RegisterService(&_Demo_serviceDesc, srv)
}

func _Demo_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Demo/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).Get(ctx, req.(*GetValRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutValRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Demo/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).Put(ctx, req.(*PutValRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServer).Stream(&demoStreamServer{stream})
}

type Demo_StreamServer interface {
	Send(*Msg) error
	Recv() (*Msg, error)
	grpc.ServerStream
}

type demoStreamServer struct {
	grpc.ServerStream
}

func (x *demoStreamServer) Send(m *Msg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demoStreamServer) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Demo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Demo",
	HandlerType: (*DemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Demo_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Demo_Put_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Demo_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe4, 0xe2, 0x75, 0x4f, 0x2d,
	0x09, 0x4b, 0xcc, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x4e,
	0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0x0c, 0xb9, 0xb8, 0x61, 0x4a,
	0x0a, 0x72, 0x2a, 0x31, 0x15, 0x80, 0x44, 0xca, 0x12, 0x73, 0x24, 0x98, 0x20, 0x22, 0x65, 0x89,
	0x39, 0x4a, 0xc6, 0x5c, 0xbc, 0x01, 0xa5, 0x78, 0x4d, 0xc5, 0xa2, 0x49, 0x96, 0x8b, 0x1b, 0xa6,
	0x09, 0x64, 0x0f, 0x1f, 0x17, 0x53, 0x7e, 0x36, 0x58, 0x07, 0x47, 0x10, 0x53, 0x7e, 0xb6, 0x92,
	0x38, 0x17, 0xb3, 0x6f, 0x71, 0x3a, 0x48, 0x5f, 0x6e, 0x71, 0x3a, 0xcc, 0xa4, 0xdc, 0xe2, 0x74,
	0xa3, 0x1a, 0x2e, 0x16, 0x97, 0xd4, 0xdc, 0x7c, 0x21, 0x4d, 0x2e, 0x66, 0xf7, 0xd4, 0x12, 0x21,
	0x41, 0xbd, 0x82, 0x24, 0x3d, 0x14, 0x3f, 0x49, 0xf1, 0x23, 0x0b, 0x81, 0xcc, 0xd6, 0xe4, 0x62,
	0x0e, 0x28, 0x85, 0x2a, 0x45, 0x71, 0x28, 0x44, 0x29, 0xb2, 0x33, 0xe4, 0xb8, 0xd8, 0x82, 0x4b,
	0x8a, 0x52, 0x13, 0x73, 0x85, 0xd8, 0x41, 0x52, 0xbe, 0xc5, 0xe9, 0x52, 0x30, 0x86, 0x06, 0xa3,
	0x01, 0x63, 0x12, 0x1b, 0x38, 0x2c, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x59, 0xac,
	0xe1, 0x58, 0x01, 0x00, 0x00,
}
