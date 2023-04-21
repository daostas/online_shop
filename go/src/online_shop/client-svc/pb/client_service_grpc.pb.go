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

// ClientsClient is the client API for Clients service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientsClient interface {
	GetClientInfo(ctx context.Context, in *ClientReq, opts ...grpc.CallOption) (*GetClientInfoRes, error)
	UpdateClientInfo(ctx context.Context, in *UpdateClientInfoReq, opts ...grpc.CallOption) (*ClientRes, error)
	UpdateClientPass(ctx context.Context, in *UpdateClientPassReq, opts ...grpc.CallOption) (*ClientRes, error)
	DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*ClientRes, error)
}

type clientsClient struct {
	cc grpc.ClientConnInterface
}

func NewClientsClient(cc grpc.ClientConnInterface) ClientsClient {
	return &clientsClient{cc}
}

func (c *clientsClient) GetClientInfo(ctx context.Context, in *ClientReq, opts ...grpc.CallOption) (*GetClientInfoRes, error) {
	out := new(GetClientInfoRes)
	err := c.cc.Invoke(ctx, "/proto.Clients/GetClientInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsClient) UpdateClientInfo(ctx context.Context, in *UpdateClientInfoReq, opts ...grpc.CallOption) (*ClientRes, error) {
	out := new(ClientRes)
	err := c.cc.Invoke(ctx, "/proto.Clients/UpdateClientInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsClient) UpdateClientPass(ctx context.Context, in *UpdateClientPassReq, opts ...grpc.CallOption) (*ClientRes, error) {
	out := new(ClientRes)
	err := c.cc.Invoke(ctx, "/proto.Clients/UpdateClientPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsClient) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*ClientRes, error) {
	out := new(ClientRes)
	err := c.cc.Invoke(ctx, "/proto.Clients/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientsServer is the server API for Clients service.
// All implementations must embed UnimplementedClientsServer
// for forward compatibility
type ClientsServer interface {
	GetClientInfo(context.Context, *ClientReq) (*GetClientInfoRes, error)
	UpdateClientInfo(context.Context, *UpdateClientInfoReq) (*ClientRes, error)
	UpdateClientPass(context.Context, *UpdateClientPassReq) (*ClientRes, error)
	DeleteClient(context.Context, *DeleteClientReq) (*ClientRes, error)
	mustEmbedUnimplementedClientsServer()
}

// UnimplementedClientsServer must be embedded to have forward compatible implementations.
type UnimplementedClientsServer struct {
}

func (UnimplementedClientsServer) GetClientInfo(context.Context, *ClientReq) (*GetClientInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientInfo not implemented")
}
func (UnimplementedClientsServer) UpdateClientInfo(context.Context, *UpdateClientInfoReq) (*ClientRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientInfo not implemented")
}
func (UnimplementedClientsServer) UpdateClientPass(context.Context, *UpdateClientPassReq) (*ClientRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientPass not implemented")
}
func (UnimplementedClientsServer) DeleteClient(context.Context, *DeleteClientReq) (*ClientRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedClientsServer) mustEmbedUnimplementedClientsServer() {}

// UnsafeClientsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientsServer will
// result in compilation errors.
type UnsafeClientsServer interface {
	mustEmbedUnimplementedClientsServer()
}

func RegisterClientsServer(s grpc.ServiceRegistrar, srv ClientsServer) {
	s.RegisterService(&Clients_ServiceDesc, srv)
}

func _Clients_GetClientInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServer).GetClientInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clients/GetClientInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServer).GetClientInfo(ctx, req.(*ClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clients_UpdateClientInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServer).UpdateClientInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clients/UpdateClientInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServer).UpdateClientInfo(ctx, req.(*UpdateClientInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clients_UpdateClientPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientPassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServer).UpdateClientPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clients/UpdateClientPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServer).UpdateClientPass(ctx, req.(*UpdateClientPassReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clients_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clients/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServer).DeleteClient(ctx, req.(*DeleteClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Clients_ServiceDesc is the grpc.ServiceDesc for Clients service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Clients_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Clients",
	HandlerType: (*ClientsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClientInfo",
			Handler:    _Clients_GetClientInfo_Handler,
		},
		{
			MethodName: "UpdateClientInfo",
			Handler:    _Clients_UpdateClientInfo_Handler,
		},
		{
			MethodName: "UpdateClientPass",
			Handler:    _Clients_UpdateClientPass_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _Clients_DeleteClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_service.proto",
}

// ClientGroupsClient is the client API for ClientGroups service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientGroupsClient interface {
	GetGroups(ctx context.Context, in *GetGroupsReq, opts ...grpc.CallOption) (*GetGroupsRes, error)
}

type clientGroupsClient struct {
	cc grpc.ClientConnInterface
}

func NewClientGroupsClient(cc grpc.ClientConnInterface) ClientGroupsClient {
	return &clientGroupsClient{cc}
}

func (c *clientGroupsClient) GetGroups(ctx context.Context, in *GetGroupsReq, opts ...grpc.CallOption) (*GetGroupsRes, error) {
	out := new(GetGroupsRes)
	err := c.cc.Invoke(ctx, "/proto.ClientGroups/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientGroupsServer is the server API for ClientGroups service.
// All implementations must embed UnimplementedClientGroupsServer
// for forward compatibility
type ClientGroupsServer interface {
	GetGroups(context.Context, *GetGroupsReq) (*GetGroupsRes, error)
	mustEmbedUnimplementedClientGroupsServer()
}

// UnimplementedClientGroupsServer must be embedded to have forward compatible implementations.
type UnimplementedClientGroupsServer struct {
}

func (UnimplementedClientGroupsServer) GetGroups(context.Context, *GetGroupsReq) (*GetGroupsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedClientGroupsServer) mustEmbedUnimplementedClientGroupsServer() {}

// UnsafeClientGroupsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientGroupsServer will
// result in compilation errors.
type UnsafeClientGroupsServer interface {
	mustEmbedUnimplementedClientGroupsServer()
}

func RegisterClientGroupsServer(s grpc.ServiceRegistrar, srv ClientGroupsServer) {
	s.RegisterService(&ClientGroups_ServiceDesc, srv)
}

func _ClientGroups_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientGroupsServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ClientGroups/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientGroupsServer).GetGroups(ctx, req.(*GetGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientGroups_ServiceDesc is the grpc.ServiceDesc for ClientGroups service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientGroups_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ClientGroups",
	HandlerType: (*ClientGroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroups",
			Handler:    _ClientGroups_GetGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_service.proto",
}

// ClientLanguagesClient is the client API for ClientLanguages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientLanguagesClient interface {
	GetLanguages(ctx context.Context, in *GetLanguagesReq, opts ...grpc.CallOption) (*GetLanguagesRes, error)
}

type clientLanguagesClient struct {
	cc grpc.ClientConnInterface
}

func NewClientLanguagesClient(cc grpc.ClientConnInterface) ClientLanguagesClient {
	return &clientLanguagesClient{cc}
}

func (c *clientLanguagesClient) GetLanguages(ctx context.Context, in *GetLanguagesReq, opts ...grpc.CallOption) (*GetLanguagesRes, error) {
	out := new(GetLanguagesRes)
	err := c.cc.Invoke(ctx, "/proto.ClientLanguages/GetLanguages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientLanguagesServer is the server API for ClientLanguages service.
// All implementations must embed UnimplementedClientLanguagesServer
// for forward compatibility
type ClientLanguagesServer interface {
	GetLanguages(context.Context, *GetLanguagesReq) (*GetLanguagesRes, error)
	mustEmbedUnimplementedClientLanguagesServer()
}

// UnimplementedClientLanguagesServer must be embedded to have forward compatible implementations.
type UnimplementedClientLanguagesServer struct {
}

func (UnimplementedClientLanguagesServer) GetLanguages(context.Context, *GetLanguagesReq) (*GetLanguagesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLanguages not implemented")
}
func (UnimplementedClientLanguagesServer) mustEmbedUnimplementedClientLanguagesServer() {}

// UnsafeClientLanguagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientLanguagesServer will
// result in compilation errors.
type UnsafeClientLanguagesServer interface {
	mustEmbedUnimplementedClientLanguagesServer()
}

func RegisterClientLanguagesServer(s grpc.ServiceRegistrar, srv ClientLanguagesServer) {
	s.RegisterService(&ClientLanguages_ServiceDesc, srv)
}

func _ClientLanguages_GetLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLanguagesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientLanguagesServer).GetLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ClientLanguages/GetLanguages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientLanguagesServer).GetLanguages(ctx, req.(*GetLanguagesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientLanguages_ServiceDesc is the grpc.ServiceDesc for ClientLanguages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientLanguages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ClientLanguages",
	HandlerType: (*ClientLanguagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLanguages",
			Handler:    _ClientLanguages_GetLanguages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_service.proto",
}
