// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: gsloc/services/gslb/v1/gslb.proto

package gslbsvc

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
	GSLB_SetEntry_FullMethodName                 = "/gsloc.services.gslb.v1.GSLB/SetEntry"
	GSLB_DeleteEntry_FullMethodName              = "/gsloc.services.gslb.v1.GSLB/DeleteEntry"
	GSLB_GetEntry_FullMethodName                 = "/gsloc.services.gslb.v1.GSLB/GetEntry"
	GSLB_ListEntries_FullMethodName              = "/gsloc.services.gslb.v1.GSLB/ListEntries"
	GSLB_AddMember_FullMethodName                = "/gsloc.services.gslb.v1.GSLB/AddMember"
	GSLB_DeleteMember_FullMethodName             = "/gsloc.services.gslb.v1.GSLB/DeleteMember"
	GSLB_SetMemberStatus_FullMethodName          = "/gsloc.services.gslb.v1.GSLB/SetMemberStatus"
	GSLB_SetMembersStatusByFilter_FullMethodName = "/gsloc.services.gslb.v1.GSLB/SetMembersStatusByFilter"
	GSLB_SetHealthCheck_FullMethodName           = "/gsloc.services.gslb.v1.GSLB/SetHealthCheck"
	GSLB_ListDcs_FullMethodName                  = "/gsloc.services.gslb.v1.GSLB/ListDcs"
)

// GSLBClient is the client API for GSLB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GSLBClient interface {
	SetEntry(ctx context.Context, in *SetEntryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error)
	ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error)
	AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetMemberStatus(ctx context.Context, in *SetMemberStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetMembersStatusByFilter(ctx context.Context, in *SetMembersStatusByFilterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetHealthCheck(ctx context.Context, in *SetHealthCheckRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListDcs(ctx context.Context, in *ListDcsRequest, opts ...grpc.CallOption) (*ListDcsResponse, error)
}

type gSLBClient struct {
	cc grpc.ClientConnInterface
}

func NewGSLBClient(cc grpc.ClientConnInterface) GSLBClient {
	return &gSLBClient{cc}
}

func (c *gSLBClient) SetEntry(ctx context.Context, in *SetEntryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GSLB_SetEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSLBClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GSLB_DeleteEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSLBClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error) {
	out := new(GetEntryResponse)
	err := c.cc.Invoke(ctx, GSLB_GetEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSLBClient) ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error) {
	out := new(ListEntriesResponse)
	err := c.cc.Invoke(ctx, GSLB_ListEntries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSLBClient) AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GSLB_AddMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSLBClient) DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GSLB_DeleteMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSLBClient) SetMemberStatus(ctx context.Context, in *SetMemberStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GSLB_SetMemberStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSLBClient) SetMembersStatusByFilter(ctx context.Context, in *SetMembersStatusByFilterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GSLB_SetMembersStatusByFilter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSLBClient) SetHealthCheck(ctx context.Context, in *SetHealthCheckRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GSLB_SetHealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSLBClient) ListDcs(ctx context.Context, in *ListDcsRequest, opts ...grpc.CallOption) (*ListDcsResponse, error) {
	out := new(ListDcsResponse)
	err := c.cc.Invoke(ctx, GSLB_ListDcs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GSLBServer is the server API for GSLB service.
// All implementations must embed UnimplementedGSLBServer
// for forward compatibility
type GSLBServer interface {
	SetEntry(context.Context, *SetEntryRequest) (*emptypb.Empty, error)
	DeleteEntry(context.Context, *DeleteEntryRequest) (*emptypb.Empty, error)
	GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error)
	ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error)
	AddMember(context.Context, *AddMemberRequest) (*emptypb.Empty, error)
	DeleteMember(context.Context, *DeleteMemberRequest) (*emptypb.Empty, error)
	SetMemberStatus(context.Context, *SetMemberStatusRequest) (*emptypb.Empty, error)
	SetMembersStatusByFilter(context.Context, *SetMembersStatusByFilterRequest) (*emptypb.Empty, error)
	SetHealthCheck(context.Context, *SetHealthCheckRequest) (*emptypb.Empty, error)
	ListDcs(context.Context, *ListDcsRequest) (*ListDcsResponse, error)
	mustEmbedUnimplementedGSLBServer()
}

// UnimplementedGSLBServer must be embedded to have forward compatible implementations.
type UnimplementedGSLBServer struct {
}

func (UnimplementedGSLBServer) SetEntry(context.Context, *SetEntryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEntry not implemented")
}
func (UnimplementedGSLBServer) DeleteEntry(context.Context, *DeleteEntryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntry not implemented")
}
func (UnimplementedGSLBServer) GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntry not implemented")
}
func (UnimplementedGSLBServer) ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntries not implemented")
}
func (UnimplementedGSLBServer) AddMember(context.Context, *AddMemberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMember not implemented")
}
func (UnimplementedGSLBServer) DeleteMember(context.Context, *DeleteMemberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMember not implemented")
}
func (UnimplementedGSLBServer) SetMemberStatus(context.Context, *SetMemberStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMemberStatus not implemented")
}
func (UnimplementedGSLBServer) SetMembersStatusByFilter(context.Context, *SetMembersStatusByFilterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMembersStatusByFilter not implemented")
}
func (UnimplementedGSLBServer) SetHealthCheck(context.Context, *SetHealthCheckRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHealthCheck not implemented")
}
func (UnimplementedGSLBServer) ListDcs(context.Context, *ListDcsRequest) (*ListDcsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDcs not implemented")
}
func (UnimplementedGSLBServer) mustEmbedUnimplementedGSLBServer() {}

// UnsafeGSLBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GSLBServer will
// result in compilation errors.
type UnsafeGSLBServer interface {
	mustEmbedUnimplementedGSLBServer()
}

func RegisterGSLBServer(s grpc.ServiceRegistrar, srv GSLBServer) {
	s.RegisterService(&GSLB_ServiceDesc, srv)
}

func _GSLB_SetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSLBServer).SetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSLB_SetEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSLBServer).SetEntry(ctx, req.(*SetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSLB_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSLBServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSLB_DeleteEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSLBServer).DeleteEntry(ctx, req.(*DeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSLB_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSLBServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSLB_GetEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSLBServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSLB_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSLBServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSLB_ListEntries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSLBServer).ListEntries(ctx, req.(*ListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSLB_AddMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSLBServer).AddMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSLB_AddMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSLBServer).AddMember(ctx, req.(*AddMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSLB_DeleteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSLBServer).DeleteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSLB_DeleteMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSLBServer).DeleteMember(ctx, req.(*DeleteMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSLB_SetMemberStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMemberStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSLBServer).SetMemberStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSLB_SetMemberStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSLBServer).SetMemberStatus(ctx, req.(*SetMemberStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSLB_SetMembersStatusByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMembersStatusByFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSLBServer).SetMembersStatusByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSLB_SetMembersStatusByFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSLBServer).SetMembersStatusByFilter(ctx, req.(*SetMembersStatusByFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSLB_SetHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSLBServer).SetHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSLB_SetHealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSLBServer).SetHealthCheck(ctx, req.(*SetHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSLB_ListDcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDcsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSLBServer).ListDcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSLB_ListDcs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSLBServer).ListDcs(ctx, req.(*ListDcsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GSLB_ServiceDesc is the grpc.ServiceDesc for GSLB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GSLB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gsloc.services.gslb.v1.GSLB",
	HandlerType: (*GSLBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetEntry",
			Handler:    _GSLB_SetEntry_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _GSLB_DeleteEntry_Handler,
		},
		{
			MethodName: "GetEntry",
			Handler:    _GSLB_GetEntry_Handler,
		},
		{
			MethodName: "ListEntries",
			Handler:    _GSLB_ListEntries_Handler,
		},
		{
			MethodName: "AddMember",
			Handler:    _GSLB_AddMember_Handler,
		},
		{
			MethodName: "DeleteMember",
			Handler:    _GSLB_DeleteMember_Handler,
		},
		{
			MethodName: "SetMemberStatus",
			Handler:    _GSLB_SetMemberStatus_Handler,
		},
		{
			MethodName: "SetMembersStatusByFilter",
			Handler:    _GSLB_SetMembersStatusByFilter_Handler,
		},
		{
			MethodName: "SetHealthCheck",
			Handler:    _GSLB_SetHealthCheck_Handler,
		},
		{
			MethodName: "ListDcs",
			Handler:    _GSLB_ListDcs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gsloc/services/gslb/v1/gslb.proto",
}
