// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: timestamp.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TimestampService_GetTimestampData_FullMethodName = "/TimestampService/GetTimestampData"
)

// TimestampServiceClient is the client API for TimestampService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimestampServiceClient interface {
	GetTimestampData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TimestampData, error)
}

type timestampServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimestampServiceClient(cc grpc.ClientConnInterface) TimestampServiceClient {
	return &timestampServiceClient{cc}
}

func (c *timestampServiceClient) GetTimestampData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TimestampData, error) {
	out := new(TimestampData)
	err := c.cc.Invoke(ctx, TimestampService_GetTimestampData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimestampServiceServer is the server API for TimestampService service.
// All implementations must embed UnimplementedTimestampServiceServer
// for forward compatibility
type TimestampServiceServer interface {
	GetTimestampData(context.Context, *emptypb.Empty) (*TimestampData, error)
	mustEmbedUnimplementedTimestampServiceServer()
}

// UnimplementedTimestampServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimestampServiceServer struct {
}

func (UnimplementedTimestampServiceServer) GetTimestampData(context.Context, *emptypb.Empty) (*TimestampData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimestampData not implemented")
}
func (UnimplementedTimestampServiceServer) mustEmbedUnimplementedTimestampServiceServer() {}

// UnsafeTimestampServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimestampServiceServer will
// result in compilation errors.
type UnsafeTimestampServiceServer interface {
	mustEmbedUnimplementedTimestampServiceServer()
}

func RegisterTimestampServiceServer(s grpc.ServiceRegistrar, srv TimestampServiceServer) {
	s.RegisterService(&TimestampService_ServiceDesc, srv)
}

func _TimestampService_GetTimestampData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimestampServiceServer).GetTimestampData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimestampService_GetTimestampData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimestampServiceServer).GetTimestampData(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TimestampService_ServiceDesc is the grpc.ServiceDesc for TimestampService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimestampService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TimestampService",
	HandlerType: (*TimestampServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimestampData",
			Handler:    _TimestampService_GetTimestampData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timestamp.proto",
}
