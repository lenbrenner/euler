// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package resources

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

// MatildaClient is the client API for Matilda service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatildaClient interface {
	GetLocation(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Location, error)
}

type matildaClient struct {
	cc grpc.ClientConnInterface
}

func NewMatildaClient(cc grpc.ClientConnInterface) MatildaClient {
	return &matildaClient{cc}
}

func (c *matildaClient) GetLocation(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/resources.Matilda/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatildaServer is the server API for Matilda service.
// All implementations must embed UnimplementedMatildaServer
// for forward compatibility
type MatildaServer interface {
	GetLocation(context.Context, *Point) (*Location, error)
	mustEmbedUnimplementedMatildaServer()
}

// UnimplementedMatildaServer must be embedded to have forward compatible implementations.
type UnimplementedMatildaServer struct {
}

func (UnimplementedMatildaServer) GetLocation(context.Context, *Point) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedMatildaServer) mustEmbedUnimplementedMatildaServer() {}

// UnsafeMatildaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatildaServer will
// result in compilation errors.
type UnsafeMatildaServer interface {
	mustEmbedUnimplementedMatildaServer()
}

func RegisterMatildaServer(s grpc.ServiceRegistrar, srv MatildaServer) {
	s.RegisterService(&Matilda_ServiceDesc, srv)
}

func _Matilda_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatildaServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Matilda/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatildaServer).GetLocation(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

// Matilda_ServiceDesc is the grpc.ServiceDesc for Matilda service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Matilda_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resources.Matilda",
	HandlerType: (*MatildaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocation",
			Handler:    _Matilda_GetLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources/matilda.proto",
}