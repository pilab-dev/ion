// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: room/v1/room.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RoomService_CreateRoom_FullMethodName = "/room.RoomService/CreateRoom"
	RoomService_UpdateRoom_FullMethodName = "/room.RoomService/UpdateRoom"
	RoomService_EndRoom_FullMethodName    = "/room.RoomService/EndRoom"
	RoomService_GetRooms_FullMethodName   = "/room.RoomService/GetRooms"
	RoomService_AddPeer_FullMethodName    = "/room.RoomService/AddPeer"
	RoomService_UpdatePeer_FullMethodName = "/room.RoomService/UpdatePeer"
	RoomService_RemovePeer_FullMethodName = "/room.RoomService/RemovePeer"
	RoomService_GetPeers_FullMethodName   = "/room.RoomService/GetPeers"
)

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomServiceClient interface {
	// Manager API
	// Room API
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomReply, error)
	UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomReply, error)
	EndRoom(ctx context.Context, in *EndRoomRequest, opts ...grpc.CallOption) (*EndRoomReply, error)
	GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsReply, error)
	// Peer API
	AddPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*AddPeerReply, error)
	UpdatePeer(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*UpdatePeerReply, error)
	RemovePeer(ctx context.Context, in *RemovePeerRequest, opts ...grpc.CallOption) (*RemovePeerReply, error)
	GetPeers(ctx context.Context, in *GetPeersRequest, opts ...grpc.CallOption) (*GetPeersReply, error)
}

type roomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomServiceClient(cc grpc.ClientConnInterface) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRoomReply)
	err := c.cc.Invoke(ctx, RoomService_CreateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRoomReply)
	err := c.cc.Invoke(ctx, RoomService_UpdateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) EndRoom(ctx context.Context, in *EndRoomRequest, opts ...grpc.CallOption) (*EndRoomReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EndRoomReply)
	err := c.cc.Invoke(ctx, RoomService_EndRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoomsReply)
	err := c.cc.Invoke(ctx, RoomService_GetRooms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) AddPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*AddPeerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPeerReply)
	err := c.cc.Invoke(ctx, RoomService_AddPeer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) UpdatePeer(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*UpdatePeerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePeerReply)
	err := c.cc.Invoke(ctx, RoomService_UpdatePeer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) RemovePeer(ctx context.Context, in *RemovePeerRequest, opts ...grpc.CallOption) (*RemovePeerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemovePeerReply)
	err := c.cc.Invoke(ctx, RoomService_RemovePeer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetPeers(ctx context.Context, in *GetPeersRequest, opts ...grpc.CallOption) (*GetPeersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPeersReply)
	err := c.cc.Invoke(ctx, RoomService_GetPeers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
// All implementations must embed UnimplementedRoomServiceServer
// for forward compatibility.
type RoomServiceServer interface {
	// Manager API
	// Room API
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomReply, error)
	UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomReply, error)
	EndRoom(context.Context, *EndRoomRequest) (*EndRoomReply, error)
	GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsReply, error)
	// Peer API
	AddPeer(context.Context, *AddPeerRequest) (*AddPeerReply, error)
	UpdatePeer(context.Context, *UpdatePeerRequest) (*UpdatePeerReply, error)
	RemovePeer(context.Context, *RemovePeerRequest) (*RemovePeerReply, error)
	GetPeers(context.Context, *GetPeersRequest) (*GetPeersReply, error)
	mustEmbedUnimplementedRoomServiceServer()
}

// UnimplementedRoomServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoomServiceServer struct{}

func (UnimplementedRoomServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoomServiceServer) UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}
func (UnimplementedRoomServiceServer) EndRoom(context.Context, *EndRoomRequest) (*EndRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndRoom not implemented")
}
func (UnimplementedRoomServiceServer) GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRooms not implemented")
}
func (UnimplementedRoomServiceServer) AddPeer(context.Context, *AddPeerRequest) (*AddPeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeer not implemented")
}
func (UnimplementedRoomServiceServer) UpdatePeer(context.Context, *UpdatePeerRequest) (*UpdatePeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePeer not implemented")
}
func (UnimplementedRoomServiceServer) RemovePeer(context.Context, *RemovePeerRequest) (*RemovePeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePeer not implemented")
}
func (UnimplementedRoomServiceServer) GetPeers(context.Context, *GetPeersRequest) (*GetPeersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeers not implemented")
}
func (UnimplementedRoomServiceServer) mustEmbedUnimplementedRoomServiceServer() {}
func (UnimplementedRoomServiceServer) testEmbeddedByValue()                     {}

// UnsafeRoomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServiceServer will
// result in compilation errors.
type UnsafeRoomServiceServer interface {
	mustEmbedUnimplementedRoomServiceServer()
}

func RegisterRoomServiceServer(s grpc.ServiceRegistrar, srv RoomServiceServer) {
	// If the following call pancis, it indicates UnimplementedRoomServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RoomService_ServiceDesc, srv)
}

func _RoomService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_UpdateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).UpdateRoom(ctx, req.(*UpdateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_EndRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).EndRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_EndRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).EndRoom(ctx, req.(*EndRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_GetRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetRooms(ctx, req.(*GetRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_AddPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).AddPeer(ctx, req.(*AddPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_UpdatePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).UpdatePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_UpdatePeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).UpdatePeer(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_RemovePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).RemovePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_RemovePeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).RemovePeer(ctx, req.(*RemovePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_GetPeers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetPeers(ctx, req.(*GetPeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoomService_ServiceDesc is the grpc.ServiceDesc for RoomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "room.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _RoomService_CreateRoom_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _RoomService_UpdateRoom_Handler,
		},
		{
			MethodName: "EndRoom",
			Handler:    _RoomService_EndRoom_Handler,
		},
		{
			MethodName: "GetRooms",
			Handler:    _RoomService_GetRooms_Handler,
		},
		{
			MethodName: "AddPeer",
			Handler:    _RoomService_AddPeer_Handler,
		},
		{
			MethodName: "UpdatePeer",
			Handler:    _RoomService_UpdatePeer_Handler,
		},
		{
			MethodName: "RemovePeer",
			Handler:    _RoomService_RemovePeer_Handler,
		},
		{
			MethodName: "GetPeers",
			Handler:    _RoomService_GetPeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "room/v1/room.proto",
}

const (
	RoomSignal_Signal_FullMethodName = "/room.RoomSignal/Signal"
)

// RoomSignalClient is the client API for RoomSignal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomSignalClient interface {
	// Signal API for room
	Signal(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SignalRequest, SignalReply], error)
}

type roomSignalClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomSignalClient(cc grpc.ClientConnInterface) RoomSignalClient {
	return &roomSignalClient{cc}
}

func (c *roomSignalClient) Signal(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SignalRequest, SignalReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RoomSignal_ServiceDesc.Streams[0], RoomSignal_Signal_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SignalRequest, SignalReply]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RoomSignal_SignalClient = grpc.BidiStreamingClient[SignalRequest, SignalReply]

// RoomSignalServer is the server API for RoomSignal service.
// All implementations must embed UnimplementedRoomSignalServer
// for forward compatibility.
type RoomSignalServer interface {
	// Signal API for room
	Signal(grpc.BidiStreamingServer[SignalRequest, SignalReply]) error
	mustEmbedUnimplementedRoomSignalServer()
}

// UnimplementedRoomSignalServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoomSignalServer struct{}

func (UnimplementedRoomSignalServer) Signal(grpc.BidiStreamingServer[SignalRequest, SignalReply]) error {
	return status.Errorf(codes.Unimplemented, "method Signal not implemented")
}
func (UnimplementedRoomSignalServer) mustEmbedUnimplementedRoomSignalServer() {}
func (UnimplementedRoomSignalServer) testEmbeddedByValue()                    {}

// UnsafeRoomSignalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomSignalServer will
// result in compilation errors.
type UnsafeRoomSignalServer interface {
	mustEmbedUnimplementedRoomSignalServer()
}

func RegisterRoomSignalServer(s grpc.ServiceRegistrar, srv RoomSignalServer) {
	// If the following call pancis, it indicates UnimplementedRoomSignalServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RoomSignal_ServiceDesc, srv)
}

func _RoomSignal_Signal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RoomSignalServer).Signal(&grpc.GenericServerStream[SignalRequest, SignalReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RoomSignal_SignalServer = grpc.BidiStreamingServer[SignalRequest, SignalReply]

// RoomSignal_ServiceDesc is the grpc.ServiceDesc for RoomSignal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomSignal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "room.RoomSignal",
	HandlerType: (*RoomSignalServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Signal",
			Handler:       _RoomSignal_Signal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "room/v1/room.proto",
}
