// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: desc/error.proto

package rpc

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// CommonClient is the client API for Common service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonClient interface {
}

type commonClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonClient(cc grpc.ClientConnInterface) CommonClient {
	return &commonClient{cc}
}

// CommonServer is the server API for Common service.
// All implementations must embed UnimplementedCommonServer
// for forward compatibility
type CommonServer interface {
	mustEmbedUnimplementedCommonServer()
}

// UnimplementedCommonServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServer struct {
}

func (UnimplementedCommonServer) mustEmbedUnimplementedCommonServer() {}

// UnsafeCommonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommonServer will
// result in compilation errors.
type UnsafeCommonServer interface {
	mustEmbedUnimplementedCommonServer()
}

func RegisterCommonServer(s grpc.ServiceRegistrar, srv CommonServer) {
	s.RegisterService(&Common_ServiceDesc, srv)
}

// Common_ServiceDesc is the grpc.ServiceDesc for Common service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Common_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.Common",
	HandlerType: (*CommonServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "desc/error.proto",
}
