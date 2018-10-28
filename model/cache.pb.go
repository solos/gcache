// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache.proto

/*
Package cache is a generated protocol buffer package.

It is generated from these files:
	cache.proto

It has these top-level messages:
	GetReq
	GetResp
	SetReq
	SetResp
*/
package cache

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

// 请求
type GetReq struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResp struct {
	Status bool   `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *GetResp) Reset()                    { *m = GetResp{} }
func (m *GetResp) String() string            { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()               {}
func (*GetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetResp) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *GetResp) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetReq struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Ttl   int64  `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *SetReq) Reset()                    { *m = SetReq{} }
func (m *SetReq) String() string            { return proto.CompactTextString(m) }
func (*SetReq) ProtoMessage()               {}
func (*SetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetReq) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *SetReq) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type SetResp struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *SetResp) Reset()                    { *m = SetResp{} }
func (m *SetResp) String() string            { return proto.CompactTextString(m) }
func (*SetResp) ProtoMessage()               {}
func (*SetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SetResp) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*GetReq)(nil), "cache.GetReq")
	proto.RegisterType((*GetResp)(nil), "cache.GetResp")
	proto.RegisterType((*SetReq)(nil), "cache.SetReq")
	proto.RegisterType((*SetResp)(nil), "cache.SetResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Data service

type DataClient interface {
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
	Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetResp, error)
}

type dataClient struct {
	cc *grpc.ClientConn
}

func NewDataClient(cc *grpc.ClientConn) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := grpc.Invoke(ctx, "/cache.Data/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetResp, error) {
	out := new(SetResp)
	err := grpc.Invoke(ctx, "/cache.Data/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Data service

type DataServer interface {
	Get(context.Context, *GetReq) (*GetResp, error)
	Set(context.Context, *SetReq) (*SetResp, error)
}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.Data/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.Data/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Set(ctx, req.(*SetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cache.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Data_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Data_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache.proto",
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4e, 0x4c, 0xce,
	0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xa4, 0xb8, 0xd8, 0xdc,
	0x53, 0x4b, 0x82, 0x52, 0x0b, 0x85, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x40, 0x4c, 0x25, 0x73, 0x2e, 0x76, 0xb0, 0x5c, 0x71, 0x81, 0x90, 0x18, 0x17,
	0x5b, 0x71, 0x49, 0x62, 0x49, 0x69, 0x31, 0x58, 0x9e, 0x23, 0x08, 0xca, 0x13, 0x12, 0xe1, 0x62,
	0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x02, 0x6b, 0x83, 0x70, 0x94, 0x9c, 0xb8, 0xd8, 0x82,
	0x71, 0x18, 0x8a, 0x5d, 0x07, 0x48, 0x5d, 0x49, 0x49, 0x8e, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x73,
	0x10, 0x88, 0xa9, 0xa4, 0xc8, 0xc5, 0x1e, 0x8c, 0xdf, 0x72, 0xa3, 0x20, 0x2e, 0x16, 0x97, 0xc4,
	0x92, 0x44, 0x21, 0x15, 0x2e, 0x66, 0xf7, 0xd4, 0x12, 0x21, 0x5e, 0x3d, 0x88, 0xff, 0x20, 0xfe,
	0x91, 0xe2, 0x43, 0xe6, 0x16, 0x17, 0x80, 0x54, 0x05, 0x23, 0xa9, 0x0a, 0x46, 0x55, 0x05, 0xb5,
	0x2b, 0x89, 0x0d, 0x1c, 0x3a, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0xf3, 0xca, 0xad,
	0x2c, 0x01, 0x00, 0x00,
}