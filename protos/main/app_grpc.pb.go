// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: app.proto

package protosGenerated

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
	Accommodation_GetAccommodation_FullMethodName    = "/accommodation/GetAccommodation"
	Accommodation_SetAccommodation_FullMethodName    = "/accommodation/SetAccommodation"
	Accommodation_UpdateAccommodation_FullMethodName = "/accommodation/UpdateAccommodation"
)

// AccommodationClient is the client API for Accommodation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationClient interface {
	GetAccommodation(ctx context.Context, in *AccommodationRequest, opts ...grpc.CallOption) (*DummyList, error)
	SetAccommodation(ctx context.Context, in *AccommodationResponse, opts ...grpc.CallOption) (*Emptya, error)
	UpdateAccommodation(ctx context.Context, in *AccommodationResponse, opts ...grpc.CallOption) (*Emptya, error)
}

type accommodationClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationClient(cc grpc.ClientConnInterface) AccommodationClient {
	return &accommodationClient{cc}
}

func (c *accommodationClient) GetAccommodation(ctx context.Context, in *AccommodationRequest, opts ...grpc.CallOption) (*DummyList, error) {
	out := new(DummyList)
	err := c.cc.Invoke(ctx, Accommodation_GetAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationClient) SetAccommodation(ctx context.Context, in *AccommodationResponse, opts ...grpc.CallOption) (*Emptya, error) {
	out := new(Emptya)
	err := c.cc.Invoke(ctx, Accommodation_SetAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationClient) UpdateAccommodation(ctx context.Context, in *AccommodationResponse, opts ...grpc.CallOption) (*Emptya, error) {
	out := new(Emptya)
	err := c.cc.Invoke(ctx, Accommodation_UpdateAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServer is the server API for Accommodation service.
// All implementations must embed UnimplementedAccommodationServer
// for forward compatibility
type AccommodationServer interface {
	GetAccommodation(context.Context, *AccommodationRequest) (*DummyList, error)
	SetAccommodation(context.Context, *AccommodationResponse) (*Emptya, error)
	UpdateAccommodation(context.Context, *AccommodationResponse) (*Emptya, error)
	mustEmbedUnimplementedAccommodationServer()
}

// UnimplementedAccommodationServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServer struct {
}

func (UnimplementedAccommodationServer) GetAccommodation(context.Context, *AccommodationRequest) (*DummyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodation not implemented")
}
func (UnimplementedAccommodationServer) SetAccommodation(context.Context, *AccommodationResponse) (*Emptya, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccommodation not implemented")
}
func (UnimplementedAccommodationServer) UpdateAccommodation(context.Context, *AccommodationResponse) (*Emptya, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccommodation not implemented")
}
func (UnimplementedAccommodationServer) mustEmbedUnimplementedAccommodationServer() {}

// UnsafeAccommodationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServer will
// result in compilation errors.
type UnsafeAccommodationServer interface {
	mustEmbedUnimplementedAccommodationServer()
}

func RegisterAccommodationServer(s grpc.ServiceRegistrar, srv AccommodationServer) {
	s.RegisterService(&Accommodation_ServiceDesc, srv)
}

func _Accommodation_GetAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServer).GetAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Accommodation_GetAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServer).GetAccommodation(ctx, req.(*AccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accommodation_SetAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServer).SetAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Accommodation_SetAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServer).SetAccommodation(ctx, req.(*AccommodationResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accommodation_UpdateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServer).UpdateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Accommodation_UpdateAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServer).UpdateAccommodation(ctx, req.(*AccommodationResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// Accommodation_ServiceDesc is the grpc.ServiceDesc for Accommodation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Accommodation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accommodation",
	HandlerType: (*AccommodationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccommodation",
			Handler:    _Accommodation_GetAccommodation_Handler,
		},
		{
			MethodName: "SetAccommodation",
			Handler:    _Accommodation_SetAccommodation_Handler,
		},
		{
			MethodName: "UpdateAccommodation",
			Handler:    _Accommodation_UpdateAccommodation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}
