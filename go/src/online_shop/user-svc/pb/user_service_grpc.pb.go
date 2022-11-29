// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	RegisterUser(ctx context.Context, in *RegUserReq, opts ...grpc.CallOption) (*UserRes, error)
	SignInUser(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*UserRes, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UserRes, error)
	UpdateUserPass(ctx context.Context, in *UpdateUserPassReq, opts ...grpc.CallOption) (*UserRes, error)
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*UserRes, error)
	AddToBasket(ctx context.Context, in *AddToBasketReq, opts ...grpc.CallOption) (*UserRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *RegUserReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/proto.UserService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignInUser(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/proto.UserService/SignInUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/proto.UserService/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserPass(ctx context.Context, in *UpdateUserPassReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/proto.UserService/UpdateUserPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/proto.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddToBasket(ctx context.Context, in *AddToBasketReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/proto.UserService/AddToBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	RegisterUser(context.Context, *RegUserReq) (*UserRes, error)
	SignInUser(context.Context, *SignInReq) (*UserRes, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoReq) (*UserRes, error)
	UpdateUserPass(context.Context, *UpdateUserPassReq) (*UserRes, error)
	DeleteUser(context.Context, *DeleteUserReq) (*UserRes, error)
	AddToBasket(context.Context, *AddToBasketReq) (*UserRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) RegisterUser(context.Context, *RegUserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserServiceServer) SignInUser(context.Context, *SignInReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserInfo(context.Context, *UpdateUserInfoReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserPass(context.Context, *UpdateUserPassReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPass not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) AddToBasket(context.Context, *AddToBasketReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToBasket not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*RegUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignInUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignInUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/SignInUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignInUser(ctx, req.(*SignInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/UpdateUserPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserPass(ctx, req.(*UpdateUserPassReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddToBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToBasketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddToBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/AddToBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddToBasket(ctx, req.(*AddToBasketReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
		{
			MethodName: "SignInUser",
			Handler:    _UserService_SignInUser_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserService_UpdateUserInfo_Handler,
		},
		{
			MethodName: "UpdateUserPass",
			Handler:    _UserService_UpdateUserPass_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "AddToBasket",
			Handler:    _UserService_AddToBasket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}