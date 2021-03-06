// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// BsbLookupClient is the client API for BsbLookup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BsbLookupClient interface {
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type bsbLookupClient struct {
	cc grpc.ClientConnInterface
}

func NewBsbLookupClient(cc grpc.ClientConnInterface) BsbLookupClient {
	return &bsbLookupClient{cc}
}

func (c *bsbLookupClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/bsbLookup.bsbLookup/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BsbLookupServer is the server API for BsbLookup service.
// All implementations must embed UnimplementedBsbLookupServer
// for forward compatibility
type BsbLookupServer interface {
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	mustEmbedUnimplementedBsbLookupServer()
}

// UnimplementedBsbLookupServer must be embedded to have forward compatible implementations.
type UnimplementedBsbLookupServer struct {
}

func (UnimplementedBsbLookupServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedBsbLookupServer) mustEmbedUnimplementedBsbLookupServer() {}

// UnsafeBsbLookupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BsbLookupServer will
// result in compilation errors.
type UnsafeBsbLookupServer interface {
	mustEmbedUnimplementedBsbLookupServer()
}

func RegisterBsbLookupServer(s grpc.ServiceRegistrar, srv BsbLookupServer) {
	s.RegisterService(&BsbLookup_ServiceDesc, srv)
}

func _BsbLookup_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BsbLookupServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bsbLookup.bsbLookup/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BsbLookupServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BsbLookup_ServiceDesc is the grpc.ServiceDesc for BsbLookup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BsbLookup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bsbLookup.bsbLookup",
	HandlerType: (*BsbLookupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _BsbLookup_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bsbLookup.proto",
}
