// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authService.proto

/*
Package authServiceProto is a generated protocol buffer package.

It is generated from these files:
	authService.proto

It has these top-level messages:
	AuthRequest
	AuthReply
*/
package authServiceProto

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

// The AuthRequest request contains two parameters.
type AuthRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *AuthRequest) Reset()                    { *m = AuthRequest{} }
func (m *AuthRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()               {}
func (*AuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// The auth response contains the result of the authentication.
type AuthReply struct {
	Response string `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *AuthReply) Reset()                    { *m = AuthReply{} }
func (m *AuthReply) String() string            { return proto.CompactTextString(m) }
func (*AuthReply) ProtoMessage()               {}
func (*AuthReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthReply) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "authServiceProto.AuthRequest")
	proto.RegisterType((*AuthReply)(nil), "authServiceProto.AuthReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	// Checks user Credentials.
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := grpc.Invoke(ctx, "/authServiceProto.Auth/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	// Checks user Credentials.
	Authenticate(context.Context, *AuthRequest) (*AuthReply, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authServiceProto.Auth/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authServiceProto.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Auth_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authService.proto",
}

func init() { proto.RegisterFile("authService.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2c, 0x2d, 0xc9,
	0x08, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40,
	0x12, 0x0a, 0x00, 0x89, 0x28, 0xb9, 0x72, 0x71, 0x3b, 0x96, 0x96, 0x64, 0x04, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x49, 0x71, 0x71, 0x94, 0x16, 0xa7, 0x16, 0xf9, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x20, 0xb9, 0x82, 0xc4, 0xe2, 0xe2, 0xf2, 0xfc,
	0xa2, 0x14, 0x09, 0x26, 0x88, 0x1c, 0x8c, 0xaf, 0xa4, 0xce, 0xc5, 0x09, 0x31, 0xa6, 0x20, 0xa7,
	0x12, 0xa4, 0xb0, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x18, 0x6e, 0x08, 0x8c, 0x6f, 0x14, 0xc2,
	0xc5, 0x02, 0x52, 0x28, 0xe4, 0xc3, 0xc5, 0x03, 0xa2, 0x53, 0xf3, 0x4a, 0x32, 0x93, 0x13, 0x4b,
	0x52, 0x85, 0x64, 0xf5, 0xd0, 0x9d, 0xa6, 0x87, 0xe4, 0x2e, 0x29, 0x69, 0x5c, 0xd2, 0x05, 0x39,
	0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xef, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x99,
	0xc0, 0x5b, 0xf3, 0x00, 0x00, 0x00,
}
