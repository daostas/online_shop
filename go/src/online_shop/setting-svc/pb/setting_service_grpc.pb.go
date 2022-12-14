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

// SettingServiceClient is the client API for SettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingServiceClient interface {
	SetDefaultLanguage(ctx context.Context, in *SetDefaultLanguageReq, opts ...grpc.CallOption) (*SettRes, error)
	SetChangingParentStatus(ctx context.Context, in *EmptySettReq, opts ...grpc.CallOption) (*SettRes, error)
}

type settingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingServiceClient(cc grpc.ClientConnInterface) SettingServiceClient {
	return &settingServiceClient{cc}
}

func (c *settingServiceClient) SetDefaultLanguage(ctx context.Context, in *SetDefaultLanguageReq, opts ...grpc.CallOption) (*SettRes, error) {
	out := new(SettRes)
	err := c.cc.Invoke(ctx, "/proto.SettingService/SetDefaultLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingServiceClient) SetChangingParentStatus(ctx context.Context, in *EmptySettReq, opts ...grpc.CallOption) (*SettRes, error) {
	out := new(SettRes)
	err := c.cc.Invoke(ctx, "/proto.SettingService/SetChangingParentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingServiceServer is the server API for SettingService service.
// All implementations must embed UnimplementedSettingServiceServer
// for forward compatibility
type SettingServiceServer interface {
	SetDefaultLanguage(context.Context, *SetDefaultLanguageReq) (*SettRes, error)
	SetChangingParentStatus(context.Context, *EmptySettReq) (*SettRes, error)
	mustEmbedUnimplementedSettingServiceServer()
}

// UnimplementedSettingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSettingServiceServer struct {
}

func (UnimplementedSettingServiceServer) SetDefaultLanguage(context.Context, *SetDefaultLanguageReq) (*SettRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultLanguage not implemented")
}
func (UnimplementedSettingServiceServer) SetChangingParentStatus(context.Context, *EmptySettReq) (*SettRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChangingParentStatus not implemented")
}
func (UnimplementedSettingServiceServer) mustEmbedUnimplementedSettingServiceServer() {}

// UnsafeSettingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingServiceServer will
// result in compilation errors.
type UnsafeSettingServiceServer interface {
	mustEmbedUnimplementedSettingServiceServer()
}

func RegisterSettingServiceServer(s grpc.ServiceRegistrar, srv SettingServiceServer) {
	s.RegisterService(&SettingService_ServiceDesc, srv)
}

func _SettingService_SetDefaultLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDefaultLanguageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).SetDefaultLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SettingService/SetDefaultLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).SetDefaultLanguage(ctx, req.(*SetDefaultLanguageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingService_SetChangingParentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptySettReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).SetChangingParentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SettingService/SetChangingParentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).SetChangingParentStatus(ctx, req.(*EmptySettReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SettingService_ServiceDesc is the grpc.ServiceDesc for SettingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SettingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SettingService",
	HandlerType: (*SettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDefaultLanguage",
			Handler:    _SettingService_SetDefaultLanguage_Handler,
		},
		{
			MethodName: "SetChangingParentStatus",
			Handler:    _SettingService_SetChangingParentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "setting_service.proto",
}
