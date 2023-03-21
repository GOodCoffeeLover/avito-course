// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: auth/protos/auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AutherClient is the client API for Auther service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutherClient interface {
	AuthByName(ctx context.Context, in *AuthByNameRequest, opts ...grpc.CallOption) (*AuthByNameResponse, error)
}

type autherClient struct {
	cc grpc.ClientConnInterface
}

func NewAutherClient(cc grpc.ClientConnInterface) AutherClient {
	return &autherClient{cc}
}

func (c *autherClient) AuthByName(ctx context.Context, in *AuthByNameRequest, opts ...grpc.CallOption) (*AuthByNameResponse, error) {
	out := new(AuthByNameResponse)
	err := c.cc.Invoke(ctx, "/auth.Auther/AuthByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutherServer is the server API for Auther service.
// All implementations must embed UnimplementedAutherServer
// for forward compatibility
type AutherServer interface {
	AuthByName(context.Context, *AuthByNameRequest) (*AuthByNameResponse, error)
	mustEmbedUnimplementedAutherServer()
}

// UnimplementedAutherServer must be embedded to have forward compatible implementations.
type UnimplementedAutherServer struct {
}

func (UnimplementedAutherServer) AuthByName(context.Context, *AuthByNameRequest) (*AuthByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthByName not implemented")
}
func (UnimplementedAutherServer) mustEmbedUnimplementedAutherServer() {}

// UnsafeAutherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutherServer will
// result in compilation errors.
type UnsafeAutherServer interface {
	mustEmbedUnimplementedAutherServer()
}

func RegisterAutherServer(s grpc.ServiceRegistrar, srv AutherServer) {
	s.RegisterService(&Auther_ServiceDesc, srv)
}

func _Auther_AuthByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutherServer).AuthByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auther/AuthByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutherServer).AuthByName(ctx, req.(*AuthByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auther_ServiceDesc is the grpc.ServiceDesc for Auther service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auther_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auther",
	HandlerType: (*AutherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthByName",
			Handler:    _Auther_AuthByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/protos/auth.proto",
}
