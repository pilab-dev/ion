// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: rtc/v1/rtc.proto

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
	RTC_Signal_FullMethodName = "/rtc.RTC/Signal"
)

// RTCClient is the client API for RTC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RTCClient interface {
	Signal(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Signalling, Signalling], error)
}

type rTCClient struct {
	cc grpc.ClientConnInterface
}

func NewRTCClient(cc grpc.ClientConnInterface) RTCClient {
	return &rTCClient{cc}
}

func (c *rTCClient) Signal(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Signalling, Signalling], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RTC_ServiceDesc.Streams[0], RTC_Signal_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Signalling, Signalling]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RTC_SignalClient = grpc.BidiStreamingClient[Signalling, Signalling]

// RTCServer is the server API for RTC service.
// All implementations must embed UnimplementedRTCServer
// for forward compatibility.
type RTCServer interface {
	Signal(grpc.BidiStreamingServer[Signalling, Signalling]) error
	mustEmbedUnimplementedRTCServer()
}

// UnimplementedRTCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRTCServer struct{}

func (UnimplementedRTCServer) Signal(grpc.BidiStreamingServer[Signalling, Signalling]) error {
	return status.Errorf(codes.Unimplemented, "method Signal not implemented")
}
func (UnimplementedRTCServer) mustEmbedUnimplementedRTCServer() {}
func (UnimplementedRTCServer) testEmbeddedByValue()             {}

// UnsafeRTCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RTCServer will
// result in compilation errors.
type UnsafeRTCServer interface {
	mustEmbedUnimplementedRTCServer()
}

func RegisterRTCServer(s grpc.ServiceRegistrar, srv RTCServer) {
	// If the following call pancis, it indicates UnimplementedRTCServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RTC_ServiceDesc, srv)
}

func _RTC_Signal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RTCServer).Signal(&grpc.GenericServerStream[Signalling, Signalling]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RTC_SignalServer = grpc.BidiStreamingServer[Signalling, Signalling]

// RTC_ServiceDesc is the grpc.ServiceDesc for RTC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RTC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rtc.RTC",
	HandlerType: (*RTCServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Signal",
			Handler:       _RTC_Signal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rtc/v1/rtc.proto",
}
