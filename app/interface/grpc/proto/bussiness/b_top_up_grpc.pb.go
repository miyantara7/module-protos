// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bussiness

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TopUpServicesClient is the client API for TopUpServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopUpServicesClient interface {
	TopUp(ctx context.Context, in *TopUpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Payment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type topUpServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewTopUpServicesClient(cc grpc.ClientConnInterface) TopUpServicesClient {
	return &topUpServicesClient{cc}
}

func (c *topUpServicesClient) TopUp(ctx context.Context, in *TopUpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bussiness.TopUpServices/TopUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topUpServicesClient) Payment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bussiness.TopUpServices/Payment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopUpServicesServer is the server API for TopUpServices service.
// All implementations should embed UnimplementedTopUpServicesServer
// for forward compatibility
type TopUpServicesServer interface {
	TopUp(context.Context, *TopUpRequest) (*emptypb.Empty, error)
	Payment(context.Context, *PaymentRequest) (*emptypb.Empty, error)
}

// UnimplementedTopUpServicesServer should be embedded to have forward compatible implementations.
type UnimplementedTopUpServicesServer struct {
}

func (UnimplementedTopUpServicesServer) TopUp(context.Context, *TopUpRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopUp not implemented")
}
func (UnimplementedTopUpServicesServer) Payment(context.Context, *PaymentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Payment not implemented")
}

// UnsafeTopUpServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopUpServicesServer will
// result in compilation errors.
type UnsafeTopUpServicesServer interface {
	mustEmbedUnimplementedTopUpServicesServer()
}

func RegisterTopUpServicesServer(s grpc.ServiceRegistrar, srv TopUpServicesServer) {
	s.RegisterService(&_TopUpServices_serviceDesc, srv)
}

func _TopUpServices_TopUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopUpServicesServer).TopUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bussiness.TopUpServices/TopUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopUpServicesServer).TopUp(ctx, req.(*TopUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopUpServices_Payment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopUpServicesServer).Payment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bussiness.TopUpServices/Payment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopUpServicesServer).Payment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopUpServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bussiness.TopUpServices",
	HandlerType: (*TopUpServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TopUp",
			Handler:    _TopUpServices_TopUp_Handler,
		},
		{
			MethodName: "Payment",
			Handler:    _TopUpServices_Payment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "b_top_up.proto",
}
