// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: schedule_tracking.proto

package __

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

// ScheduleTrackingClient is the client API for ScheduleTracking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleTrackingClient interface {
	AddContainersOnTrack(ctx context.Context, in *AddOnTrackRequest, opts ...grpc.CallOption) (*AddOnTrackResponse, error)
	AddBillNosOnTrack(ctx context.Context, in *AddOnTrackRequest, opts ...grpc.CallOption) (*AddOnTrackResponse, error)
	UpdateContainer(ctx context.Context, in *AddOnTrackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateBill(ctx context.Context, in *AddOnTrackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteContainersFromTracking(ctx context.Context, in *DeleteFromTrackingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteBillsFromTracking(ctx context.Context, in *DeleteFromTrackingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetInfoAboutTrack(ctx context.Context, in *GetInfoAboutTrackRequest, opts ...grpc.CallOption) (*GetInfoAboutTrackResponse, error)
	GetTimeZone(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTimeZoneResponse, error)
}

type scheduleTrackingClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleTrackingClient(cc grpc.ClientConnInterface) ScheduleTrackingClient {
	return &scheduleTrackingClient{cc}
}

func (c *scheduleTrackingClient) AddContainersOnTrack(ctx context.Context, in *AddOnTrackRequest, opts ...grpc.CallOption) (*AddOnTrackResponse, error) {
	out := new(AddOnTrackResponse)
	err := c.cc.Invoke(ctx, "/schedule_tacking.ScheduleTracking/AddContainersOnTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleTrackingClient) AddBillNosOnTrack(ctx context.Context, in *AddOnTrackRequest, opts ...grpc.CallOption) (*AddOnTrackResponse, error) {
	out := new(AddOnTrackResponse)
	err := c.cc.Invoke(ctx, "/schedule_tacking.ScheduleTracking/AddBillNosOnTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleTrackingClient) UpdateContainer(ctx context.Context, in *AddOnTrackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/schedule_tacking.ScheduleTracking/UpdateContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleTrackingClient) UpdateBill(ctx context.Context, in *AddOnTrackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/schedule_tacking.ScheduleTracking/UpdateBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleTrackingClient) DeleteContainersFromTracking(ctx context.Context, in *DeleteFromTrackingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/schedule_tacking.ScheduleTracking/DeleteContainersFromTracking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleTrackingClient) DeleteBillsFromTracking(ctx context.Context, in *DeleteFromTrackingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/schedule_tacking.ScheduleTracking/DeleteBillsFromTracking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleTrackingClient) GetInfoAboutTrack(ctx context.Context, in *GetInfoAboutTrackRequest, opts ...grpc.CallOption) (*GetInfoAboutTrackResponse, error) {
	out := new(GetInfoAboutTrackResponse)
	err := c.cc.Invoke(ctx, "/schedule_tacking.ScheduleTracking/GetInfoAboutTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleTrackingClient) GetTimeZone(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTimeZoneResponse, error) {
	out := new(GetTimeZoneResponse)
	err := c.cc.Invoke(ctx, "/schedule_tacking.ScheduleTracking/GetTimeZone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleTrackingServer is the server API for ScheduleTracking service.
// All implementations must embed UnimplementedScheduleTrackingServer
// for forward compatibility
type ScheduleTrackingServer interface {
	AddContainersOnTrack(context.Context, *AddOnTrackRequest) (*AddOnTrackResponse, error)
	AddBillNosOnTrack(context.Context, *AddOnTrackRequest) (*AddOnTrackResponse, error)
	UpdateContainer(context.Context, *AddOnTrackRequest) (*emptypb.Empty, error)
	UpdateBill(context.Context, *AddOnTrackRequest) (*emptypb.Empty, error)
	DeleteContainersFromTracking(context.Context, *DeleteFromTrackingRequest) (*emptypb.Empty, error)
	DeleteBillsFromTracking(context.Context, *DeleteFromTrackingRequest) (*emptypb.Empty, error)
	GetInfoAboutTrack(context.Context, *GetInfoAboutTrackRequest) (*GetInfoAboutTrackResponse, error)
	GetTimeZone(context.Context, *emptypb.Empty) (*GetTimeZoneResponse, error)
	mustEmbedUnimplementedScheduleTrackingServer()
}

// UnimplementedScheduleTrackingServer must be embedded to have forward compatible implementations.
type UnimplementedScheduleTrackingServer struct {
}

func (UnimplementedScheduleTrackingServer) AddContainersOnTrack(context.Context, *AddOnTrackRequest) (*AddOnTrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContainersOnTrack not implemented")
}
func (UnimplementedScheduleTrackingServer) AddBillNosOnTrack(context.Context, *AddOnTrackRequest) (*AddOnTrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBillNosOnTrack not implemented")
}
func (UnimplementedScheduleTrackingServer) UpdateContainer(context.Context, *AddOnTrackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContainer not implemented")
}
func (UnimplementedScheduleTrackingServer) UpdateBill(context.Context, *AddOnTrackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBill not implemented")
}
func (UnimplementedScheduleTrackingServer) DeleteContainersFromTracking(context.Context, *DeleteFromTrackingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContainersFromTracking not implemented")
}
func (UnimplementedScheduleTrackingServer) DeleteBillsFromTracking(context.Context, *DeleteFromTrackingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBillsFromTracking not implemented")
}
func (UnimplementedScheduleTrackingServer) GetInfoAboutTrack(context.Context, *GetInfoAboutTrackRequest) (*GetInfoAboutTrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoAboutTrack not implemented")
}
func (UnimplementedScheduleTrackingServer) GetTimeZone(context.Context, *emptypb.Empty) (*GetTimeZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeZone not implemented")
}
func (UnimplementedScheduleTrackingServer) mustEmbedUnimplementedScheduleTrackingServer() {}

// UnsafeScheduleTrackingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleTrackingServer will
// result in compilation errors.
type UnsafeScheduleTrackingServer interface {
	mustEmbedUnimplementedScheduleTrackingServer()
}

func RegisterScheduleTrackingServer(s grpc.ServiceRegistrar, srv ScheduleTrackingServer) {
	s.RegisterService(&ScheduleTracking_ServiceDesc, srv)
}

func _ScheduleTracking_AddContainersOnTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOnTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleTrackingServer).AddContainersOnTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_tacking.ScheduleTracking/AddContainersOnTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleTrackingServer).AddContainersOnTrack(ctx, req.(*AddOnTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleTracking_AddBillNosOnTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOnTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleTrackingServer).AddBillNosOnTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_tacking.ScheduleTracking/AddBillNosOnTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleTrackingServer).AddBillNosOnTrack(ctx, req.(*AddOnTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleTracking_UpdateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOnTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleTrackingServer).UpdateContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_tacking.ScheduleTracking/UpdateContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleTrackingServer).UpdateContainer(ctx, req.(*AddOnTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleTracking_UpdateBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOnTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleTrackingServer).UpdateBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_tacking.ScheduleTracking/UpdateBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleTrackingServer).UpdateBill(ctx, req.(*AddOnTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleTracking_DeleteContainersFromTracking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFromTrackingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleTrackingServer).DeleteContainersFromTracking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_tacking.ScheduleTracking/DeleteContainersFromTracking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleTrackingServer).DeleteContainersFromTracking(ctx, req.(*DeleteFromTrackingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleTracking_DeleteBillsFromTracking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFromTrackingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleTrackingServer).DeleteBillsFromTracking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_tacking.ScheduleTracking/DeleteBillsFromTracking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleTrackingServer).DeleteBillsFromTracking(ctx, req.(*DeleteFromTrackingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleTracking_GetInfoAboutTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoAboutTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleTrackingServer).GetInfoAboutTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_tacking.ScheduleTracking/GetInfoAboutTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleTrackingServer).GetInfoAboutTrack(ctx, req.(*GetInfoAboutTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleTracking_GetTimeZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleTrackingServer).GetTimeZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_tacking.ScheduleTracking/GetTimeZone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleTrackingServer).GetTimeZone(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ScheduleTracking_ServiceDesc is the grpc.ServiceDesc for ScheduleTracking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScheduleTracking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedule_tacking.ScheduleTracking",
	HandlerType: (*ScheduleTrackingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddContainersOnTrack",
			Handler:    _ScheduleTracking_AddContainersOnTrack_Handler,
		},
		{
			MethodName: "AddBillNosOnTrack",
			Handler:    _ScheduleTracking_AddBillNosOnTrack_Handler,
		},
		{
			MethodName: "UpdateContainer",
			Handler:    _ScheduleTracking_UpdateContainer_Handler,
		},
		{
			MethodName: "UpdateBill",
			Handler:    _ScheduleTracking_UpdateBill_Handler,
		},
		{
			MethodName: "DeleteContainersFromTracking",
			Handler:    _ScheduleTracking_DeleteContainersFromTracking_Handler,
		},
		{
			MethodName: "DeleteBillsFromTracking",
			Handler:    _ScheduleTracking_DeleteBillsFromTracking_Handler,
		},
		{
			MethodName: "GetInfoAboutTrack",
			Handler:    _ScheduleTracking_GetInfoAboutTrack_Handler,
		},
		{
			MethodName: "GetTimeZone",
			Handler:    _ScheduleTracking_GetTimeZone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedule_tracking.proto",
}

// ArchiveClient is the client API for Archive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchiveClient interface {
	GetArchive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllBillsContainerResponse, error)
}

type archiveClient struct {
	cc grpc.ClientConnInterface
}

func NewArchiveClient(cc grpc.ClientConnInterface) ArchiveClient {
	return &archiveClient{cc}
}

func (c *archiveClient) GetArchive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllBillsContainerResponse, error) {
	out := new(GetAllBillsContainerResponse)
	err := c.cc.Invoke(ctx, "/schedule_tacking.Archive/GetArchive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchiveServer is the server API for Archive service.
// All implementations must embed UnimplementedArchiveServer
// for forward compatibility
type ArchiveServer interface {
	GetArchive(context.Context, *emptypb.Empty) (*GetAllBillsContainerResponse, error)
	mustEmbedUnimplementedArchiveServer()
}

// UnimplementedArchiveServer must be embedded to have forward compatible implementations.
type UnimplementedArchiveServer struct {
}

func (UnimplementedArchiveServer) GetArchive(context.Context, *emptypb.Empty) (*GetAllBillsContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArchive not implemented")
}
func (UnimplementedArchiveServer) mustEmbedUnimplementedArchiveServer() {}

// UnsafeArchiveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchiveServer will
// result in compilation errors.
type UnsafeArchiveServer interface {
	mustEmbedUnimplementedArchiveServer()
}

func RegisterArchiveServer(s grpc.ServiceRegistrar, srv ArchiveServer) {
	s.RegisterService(&Archive_ServiceDesc, srv)
}

func _Archive_GetArchive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveServer).GetArchive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_tacking.Archive/GetArchive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveServer).GetArchive(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Archive_ServiceDesc is the grpc.ServiceDesc for Archive service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Archive_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedule_tacking.Archive",
	HandlerType: (*ArchiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArchive",
			Handler:    _Archive_GetArchive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedule_tracking.proto",
}
