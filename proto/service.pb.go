// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Request
	Response
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	A      int64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B      int64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	Number int64 `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Request) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *Request) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type Response struct {
	Result int64 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Prime  bool  `protobuf:"varint,2,opt,name=prime" json:"prime,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *Response) GetPrime() bool {
	if m != nil {
		return m.Prime
	}
	return false
}

func init() {
	proto1.RegisterType((*Request)(nil), "proto.Request")
	proto1.RegisterType((*Response)(nil), "proto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AddService service

type AddServiceClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	PrimeNumbers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.AddService/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.AddService/Multiply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) PrimeNumbers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.AddService/PrimeNumbers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AddService service

type AddServiceServer interface {
	Add(context.Context, *Request) (*Response, error)
	Multiply(context.Context, *Request) (*Response, error)
	PrimeNumbers(context.Context, *Request) (*Response, error)
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Multiply(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_PrimeNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).PrimeNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/PrimeNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).PrimeNumbers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AddService_Add_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _AddService_Multiply_Handler,
		},
		{
			MethodName: "PrimeNumbers",
			Handler:    _AddService_PrimeNumbers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto1.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xb6, 0x5c,
	0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x3c, 0x5c, 0x8c, 0x89, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xcc, 0x41, 0x8c, 0x89, 0x20, 0x5e, 0x92, 0x04, 0x13, 0x84, 0x97, 0x24, 0x24, 0xc6,
	0xc5, 0x96, 0x57, 0x9a, 0x9b, 0x94, 0x5a, 0x24, 0xc1, 0x0c, 0x16, 0x82, 0xf2, 0x94, 0x2c, 0xb8,
	0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x41, 0x6a, 0x8a, 0x52, 0x8b, 0x4b, 0x73,
	0x4a, 0xa0, 0x86, 0x40, 0x79, 0x42, 0x22, 0x5c, 0xac, 0x05, 0x45, 0x99, 0xb9, 0xa9, 0x60, 0xd3,
	0x38, 0x82, 0x20, 0x1c, 0xa3, 0x49, 0x8c, 0x5c, 0x5c, 0x8e, 0x29, 0x29, 0xc1, 0x10, 0x47, 0x09,
	0xa9, 0x71, 0x31, 0x3b, 0xa6, 0xa4, 0x08, 0xf1, 0x41, 0x5c, 0xa7, 0x07, 0x75, 0x93, 0x14, 0x3f,
	0x9c, 0x0f, 0xb5, 0x44, 0x9b, 0x8b, 0xc3, 0xb7, 0x34, 0xa7, 0x24, 0xb3, 0x20, 0xa7, 0x92, 0xb0,
	0x62, 0x7d, 0x2e, 0x9e, 0x00, 0x90, 0x65, 0x7e, 0x60, 0xc7, 0x16, 0x13, 0xd4, 0x90, 0xc4, 0x06,
	0xe6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x05, 0xe4, 0x2c, 0x2c, 0x01, 0x00, 0x00,
}