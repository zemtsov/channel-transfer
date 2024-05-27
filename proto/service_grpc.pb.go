// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: service.proto

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

const (
	API_TransferByCustomer_FullMethodName = "/core.channel_transfer.API/TransferByCustomer"
	API_TransferByAdmin_FullMethodName    = "/core.channel_transfer.API/TransferByAdmin"
	API_TransferStatus_FullMethodName     = "/core.channel_transfer.API/TransferStatus"
)

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	TransferByCustomer(ctx context.Context, in *TransferBeginCustomerRequest, opts ...grpc.CallOption) (*TransferStatusResponse, error)
	TransferByAdmin(ctx context.Context, in *TransferBeginAdminRequest, opts ...grpc.CallOption) (*TransferStatusResponse, error)
	TransferStatus(ctx context.Context, in *TransferStatusRequest, opts ...grpc.CallOption) (*TransferStatusResponse, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) TransferByCustomer(ctx context.Context, in *TransferBeginCustomerRequest, opts ...grpc.CallOption) (*TransferStatusResponse, error) {
	out := new(TransferStatusResponse)
	err := c.cc.Invoke(ctx, API_TransferByCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) TransferByAdmin(ctx context.Context, in *TransferBeginAdminRequest, opts ...grpc.CallOption) (*TransferStatusResponse, error) {
	out := new(TransferStatusResponse)
	err := c.cc.Invoke(ctx, API_TransferByAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) TransferStatus(ctx context.Context, in *TransferStatusRequest, opts ...grpc.CallOption) (*TransferStatusResponse, error) {
	out := new(TransferStatusResponse)
	err := c.cc.Invoke(ctx, API_TransferStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations must embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	TransferByCustomer(context.Context, *TransferBeginCustomerRequest) (*TransferStatusResponse, error)
	TransferByAdmin(context.Context, *TransferBeginAdminRequest) (*TransferStatusResponse, error)
	TransferStatus(context.Context, *TransferStatusRequest) (*TransferStatusResponse, error)
	mustEmbedUnimplementedAPIServer()
}

// UnimplementedAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) TransferByCustomer(context.Context, *TransferBeginCustomerRequest) (*TransferStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferByCustomer not implemented")
}
func (UnimplementedAPIServer) TransferByAdmin(context.Context, *TransferBeginAdminRequest) (*TransferStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferByAdmin not implemented")
}
func (UnimplementedAPIServer) TransferStatus(context.Context, *TransferStatusRequest) (*TransferStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferStatus not implemented")
}
func (UnimplementedAPIServer) mustEmbedUnimplementedAPIServer() {}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_TransferByCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferBeginCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).TransferByCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: API_TransferByCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).TransferByCustomer(ctx, req.(*TransferBeginCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_TransferByAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferBeginAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).TransferByAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: API_TransferByAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).TransferByAdmin(ctx, req.(*TransferBeginAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_TransferStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).TransferStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: API_TransferStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).TransferStatus(ctx, req.(*TransferStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.channel_transfer.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransferByCustomer",
			Handler:    _API_TransferByCustomer_Handler,
		},
		{
			MethodName: "TransferByAdmin",
			Handler:    _API_TransferByAdmin_Handler,
		},
		{
			MethodName: "TransferStatus",
			Handler:    _API_TransferStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}