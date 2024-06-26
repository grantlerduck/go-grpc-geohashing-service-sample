// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: v1/poi.proto

package poiv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PoIQueryInfo_GetPoI_FullMethodName              = "/api.v1.PoIQueryInfo/GetPoI"
	PoIQueryInfo_GetPoIsInProximity_FullMethodName  = "/api.v1.PoIQueryInfo/GetPoIsInProximity"
	PoIQueryInfo_PoISearchAlongRoute_FullMethodName = "/api.v1.PoIQueryInfo/PoISearchAlongRoute"
)

// PoIQueryInfoClient is the client API for PoIQueryInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// see https://github.com/grpc-ecosystem/grpc-gateway
type PoIQueryInfoClient interface {
	GetPoI(ctx context.Context, in *GetPoIRequest, opts ...grpc.CallOption) (*GetPoIResponse, error)
	GetPoIsInProximity(ctx context.Context, in *GetPoIsInProximityRequest, opts ...grpc.CallOption) (*PoIListResponse, error)
	PoISearchAlongRoute(ctx context.Context, in *GetPoIsAlongRouteRequest, opts ...grpc.CallOption) (*PoIListResponse, error)
}

type poIQueryInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewPoIQueryInfoClient(cc grpc.ClientConnInterface) PoIQueryInfoClient {
	return &poIQueryInfoClient{cc}
}

func (c *poIQueryInfoClient) GetPoI(ctx context.Context, in *GetPoIRequest, opts ...grpc.CallOption) (*GetPoIResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPoIResponse)
	err := c.cc.Invoke(ctx, PoIQueryInfo_GetPoI_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poIQueryInfoClient) GetPoIsInProximity(ctx context.Context, in *GetPoIsInProximityRequest, opts ...grpc.CallOption) (*PoIListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PoIListResponse)
	err := c.cc.Invoke(ctx, PoIQueryInfo_GetPoIsInProximity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poIQueryInfoClient) PoISearchAlongRoute(ctx context.Context, in *GetPoIsAlongRouteRequest, opts ...grpc.CallOption) (*PoIListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PoIListResponse)
	err := c.cc.Invoke(ctx, PoIQueryInfo_PoISearchAlongRoute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoIQueryInfoServer is the server API for PoIQueryInfo service.
// All implementations should embed UnimplementedPoIQueryInfoServer
// for forward compatibility
//
// see https://github.com/grpc-ecosystem/grpc-gateway
type PoIQueryInfoServer interface {
	GetPoI(context.Context, *GetPoIRequest) (*GetPoIResponse, error)
	GetPoIsInProximity(context.Context, *GetPoIsInProximityRequest) (*PoIListResponse, error)
	PoISearchAlongRoute(context.Context, *GetPoIsAlongRouteRequest) (*PoIListResponse, error)
}

// UnimplementedPoIQueryInfoServer should be embedded to have forward compatible implementations.
type UnimplementedPoIQueryInfoServer struct {
}

func (UnimplementedPoIQueryInfoServer) GetPoI(context.Context, *GetPoIRequest) (*GetPoIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoI not implemented")
}
func (UnimplementedPoIQueryInfoServer) GetPoIsInProximity(context.Context, *GetPoIsInProximityRequest) (*PoIListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoIsInProximity not implemented")
}
func (UnimplementedPoIQueryInfoServer) PoISearchAlongRoute(context.Context, *GetPoIsAlongRouteRequest) (*PoIListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PoISearchAlongRoute not implemented")
}

// UnsafePoIQueryInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoIQueryInfoServer will
// result in compilation errors.
type UnsafePoIQueryInfoServer interface {
	mustEmbedUnimplementedPoIQueryInfoServer()
}

func RegisterPoIQueryInfoServer(s grpc.ServiceRegistrar, srv PoIQueryInfoServer) {
	s.RegisterService(&PoIQueryInfo_ServiceDesc, srv)
}

func _PoIQueryInfo_GetPoI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoIQueryInfoServer).GetPoI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoIQueryInfo_GetPoI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoIQueryInfoServer).GetPoI(ctx, req.(*GetPoIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoIQueryInfo_GetPoIsInProximity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoIsInProximityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoIQueryInfoServer).GetPoIsInProximity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoIQueryInfo_GetPoIsInProximity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoIQueryInfoServer).GetPoIsInProximity(ctx, req.(*GetPoIsInProximityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoIQueryInfo_PoISearchAlongRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoIsAlongRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoIQueryInfoServer).PoISearchAlongRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoIQueryInfo_PoISearchAlongRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoIQueryInfoServer).PoISearchAlongRoute(ctx, req.(*GetPoIsAlongRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PoIQueryInfo_ServiceDesc is the grpc.ServiceDesc for PoIQueryInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PoIQueryInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.PoIQueryInfo",
	HandlerType: (*PoIQueryInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPoI",
			Handler:    _PoIQueryInfo_GetPoI_Handler,
		},
		{
			MethodName: "GetPoIsInProximity",
			Handler:    _PoIQueryInfo_GetPoIsInProximity_Handler,
		},
		{
			MethodName: "PoISearchAlongRoute",
			Handler:    _PoIQueryInfo_PoISearchAlongRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/poi.proto",
}
