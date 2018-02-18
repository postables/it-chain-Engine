// Code generated by protoc-gen-go. DO NOT EDIT.
// source: webhook.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	webhook.proto

It has these top-level messages:
	WebhookRequest
	WebhookResponse
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

type WebhookRequest struct {
	Urls string `protobuf:"bytes,1,opt,name=urls" json:"urls,omitempty"`
}

func (m *WebhookRequest) Reset()                    { *m = WebhookRequest{} }
func (m *WebhookRequest) String() string            { return proto1.CompactTextString(m) }
func (*WebhookRequest) ProtoMessage()               {}
func (*WebhookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WebhookRequest) GetUrls() string {
	if m != nil {
		return m.Urls
	}
	return ""
}

type WebhookResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *WebhookResponse) Reset()                    { *m = WebhookResponse{} }
func (m *WebhookResponse) String() string            { return proto1.CompactTextString(m) }
func (*WebhookResponse) ProtoMessage()               {}
func (*WebhookResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WebhookResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*WebhookRequest)(nil), "proto.WebhookRequest")
	proto1.RegisterType((*WebhookResponse)(nil), "proto.WebhookResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Webhook service

type WebhookClient interface {
	Register(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookResponse, error)
}

type webhookClient struct {
	cc *grpc.ClientConn
}

func NewWebhookClient(cc *grpc.ClientConn) WebhookClient {
	return &webhookClient{cc}
}

func (c *webhookClient) Register(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookResponse, error) {
	out := new(WebhookResponse)
	err := grpc.Invoke(ctx, "/proto.Webhook/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Webhook service

type WebhookServer interface {
	Register(context.Context, *WebhookRequest) (*WebhookResponse, error)
}

func RegisterWebhookServer(s *grpc.Server, srv WebhookServer) {
	s.RegisterService(&_Webhook_serviceDesc, srv)
}

func _Webhook_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Webhook/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Register(ctx, req.(*WebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Webhook_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Webhook",
	HandlerType: (*WebhookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Webhook_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webhook.proto",
}

func init() { proto1.RegisterFile("webhook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4f, 0x4d, 0xca,
	0xc8, 0xcf, 0xcf, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x2a, 0x5c,
	0x7c, 0xe1, 0x10, 0xf1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xd2,
	0xa2, 0x9c, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x9b, 0x8b, 0x1f,
	0xae, 0xaa, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38,
	0x31, 0x3d, 0x15, 0xaa, 0x12, 0xc6, 0x35, 0x72, 0xe3, 0x62, 0x87, 0x2a, 0x16, 0xb2, 0xe6, 0xe2,
	0x08, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0x12, 0x12, 0x85, 0x58, 0xac, 0x87, 0x6a, 0x9d,
	0x94, 0x18, 0xba, 0x30, 0xc4, 0x7c, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x84, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xaf, 0x23, 0x64, 0xb5, 0xb9, 0x00, 0x00, 0x00,
}
