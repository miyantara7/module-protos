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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UsermanagementServiceClient is the client API for UsermanagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsermanagementServiceClient interface {
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetDetailUserInformation(ctx context.Context, in *UserInformationReq, opts ...grpc.CallOption) (*UserInformationRes, error)
}

type usermanagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsermanagementServiceClient(cc grpc.ClientConnInterface) UsermanagementServiceClient {
	return &usermanagementServiceClient{cc}
}

func (c *usermanagementServiceClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/bussiness.UsermanagementService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usermanagementServiceClient) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bussiness.UsermanagementService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usermanagementServiceClient) GetDetailUserInformation(ctx context.Context, in *UserInformationReq, opts ...grpc.CallOption) (*UserInformationRes, error) {
	out := new(UserInformationRes)
	err := c.cc.Invoke(ctx, "/bussiness.UsermanagementService/GetDetailUserInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsermanagementServiceServer is the server API for UsermanagementService service.
// All implementations must embed UnimplementedUsermanagementServiceServer
// for forward compatibility
type UsermanagementServiceServer interface {
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	RegisterUser(context.Context, *RegisterRequest) (*emptypb.Empty, error)
	GetDetailUserInformation(context.Context, *UserInformationReq) (*UserInformationRes, error)
	mustEmbedUnimplementedUsermanagementServiceServer()
}

// UnimplementedUsermanagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsermanagementServiceServer struct {
}

func (UnimplementedUsermanagementServiceServer) LoginUser(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedUsermanagementServiceServer) RegisterUser(context.Context, *RegisterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUsermanagementServiceServer) GetDetailUserInformation(context.Context, *UserInformationReq) (*UserInformationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailUserInformation not implemented")
}
func (UnimplementedUsermanagementServiceServer) mustEmbedUnimplementedUsermanagementServiceServer() {}

// UnsafeUsermanagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsermanagementServiceServer will
// result in compilation errors.
type UnsafeUsermanagementServiceServer interface {
	mustEmbedUnimplementedUsermanagementServiceServer()
}

func RegisterUsermanagementServiceServer(s grpc.ServiceRegistrar, srv UsermanagementServiceServer) {
	s.RegisterService(&UsermanagementService_ServiceDesc, srv)
}

func _UsermanagementService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsermanagementServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bussiness.UsermanagementService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsermanagementServiceServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsermanagementService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsermanagementServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bussiness.UsermanagementService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsermanagementServiceServer).RegisterUser(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsermanagementService_GetDetailUserInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInformationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsermanagementServiceServer).GetDetailUserInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bussiness.UsermanagementService/GetDetailUserInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsermanagementServiceServer).GetDetailUserInformation(ctx, req.(*UserInformationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UsermanagementService_ServiceDesc is the grpc.ServiceDesc for UsermanagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsermanagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bussiness.UsermanagementService",
	HandlerType: (*UsermanagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _UsermanagementService_LoginUser_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _UsermanagementService_RegisterUser_Handler,
		},
		{
			MethodName: "GetDetailUserInformation",
			Handler:    _UsermanagementService_GetDetailUserInformation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/proto/bussiness/b_user_management.proto",
}
