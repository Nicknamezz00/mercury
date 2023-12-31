// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v2/message_service.proto

package apiv2

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
	MercuryService_ListMessages_FullMethodName = "/mercury.api.v2.MercuryService/ListMessages"
	MercuryService_GetMessage_FullMethodName   = "/mercury.api.v2.MercuryService/GetMessage"
)

// MercuryServiceClient is the client API for MercuryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MercuryServiceClient interface {
	ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error)
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error)
}

type mercuryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMercuryServiceClient(cc grpc.ClientConnInterface) MercuryServiceClient {
	return &mercuryServiceClient{cc}
}

func (c *mercuryServiceClient) ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error) {
	out := new(ListMessagesResponse)
	err := c.cc.Invoke(ctx, MercuryService_ListMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mercuryServiceClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error) {
	out := new(GetMessageResponse)
	err := c.cc.Invoke(ctx, MercuryService_GetMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MercuryServiceServer is the server API for MercuryService service.
// All implementations must embed UnimplementedMercuryServiceServer
// for forward compatibility
type MercuryServiceServer interface {
	ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error)
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)
	mustEmbedUnimplementedMercuryServiceServer()
}

// UnimplementedMercuryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMercuryServiceServer struct {
}

func (UnimplementedMercuryServiceServer) ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessages not implemented")
}
func (UnimplementedMercuryServiceServer) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedMercuryServiceServer) mustEmbedUnimplementedMercuryServiceServer() {}

// UnsafeMercuryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MercuryServiceServer will
// result in compilation errors.
type UnsafeMercuryServiceServer interface {
	mustEmbedUnimplementedMercuryServiceServer()
}

func RegisterMercuryServiceServer(s grpc.ServiceRegistrar, srv MercuryServiceServer) {
	s.RegisterService(&MercuryService_ServiceDesc, srv)
}

func _MercuryService_ListMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryServiceServer).ListMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MercuryService_ListMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryServiceServer).ListMessages(ctx, req.(*ListMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MercuryService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MercuryService_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryServiceServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MercuryService_ServiceDesc is the grpc.ServiceDesc for MercuryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MercuryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mercury.api.v2.MercuryService",
	HandlerType: (*MercuryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMessages",
			Handler:    _MercuryService_ListMessages_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _MercuryService_GetMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v2/message_service.proto",
}
