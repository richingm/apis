// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: users/v1/users.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Users_CreateUsers_FullMethodName = "/api.users.v1.Users/CreateUsers"
	Users_UpdateUsers_FullMethodName = "/api.users.v1.Users/UpdateUsers"
	Users_DeleteUsers_FullMethodName = "/api.users.v1.Users/DeleteUsers"
	Users_GetUsers_FullMethodName    = "/api.users.v1.Users/GetUsers"
	Users_ListUsers_FullMethodName   = "/api.users.v1.Users/ListUsers"
)

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	CreateUsers(ctx context.Context, in *CreateUsersRequest, opts ...grpc.CallOption) (*CreateUsersReply, error)
	UpdateUsers(ctx context.Context, in *UpdateUsersRequest, opts ...grpc.CallOption) (*UpdateUsersReply, error)
	DeleteUsers(ctx context.Context, in *DeleteUsersRequest, opts ...grpc.CallOption) (*DeleteUsersReply, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersReply, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersReply, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) CreateUsers(ctx context.Context, in *CreateUsersRequest, opts ...grpc.CallOption) (*CreateUsersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUsersReply)
	err := c.cc.Invoke(ctx, Users_CreateUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUsers(ctx context.Context, in *UpdateUsersRequest, opts ...grpc.CallOption) (*UpdateUsersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUsersReply)
	err := c.cc.Invoke(ctx, Users_UpdateUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DeleteUsers(ctx context.Context, in *DeleteUsersRequest, opts ...grpc.CallOption) (*DeleteUsersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUsersReply)
	err := c.cc.Invoke(ctx, Users_DeleteUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUsersReply)
	err := c.cc.Invoke(ctx, Users_GetUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUsersReply)
	err := c.cc.Invoke(ctx, Users_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	CreateUsers(context.Context, *CreateUsersRequest) (*CreateUsersReply, error)
	UpdateUsers(context.Context, *UpdateUsersRequest) (*UpdateUsersReply, error)
	DeleteUsers(context.Context, *DeleteUsersRequest) (*DeleteUsersReply, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersReply, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersReply, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) CreateUsers(context.Context, *CreateUsersRequest) (*CreateUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUsers not implemented")
}
func (UnimplementedUsersServer) UpdateUsers(context.Context, *UpdateUsersRequest) (*UpdateUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUsers not implemented")
}
func (UnimplementedUsersServer) DeleteUsers(context.Context, *DeleteUsersRequest) (*DeleteUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUsers not implemented")
}
func (UnimplementedUsersServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUsersServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_CreateUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_CreateUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUsers(ctx, req.(*CreateUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_UpdateUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUsers(ctx, req.(*UpdateUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DeleteUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_DeleteUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteUsers(ctx, req.(*DeleteUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.users.v1.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUsers",
			Handler:    _Users_CreateUsers_Handler,
		},
		{
			MethodName: "UpdateUsers",
			Handler:    _Users_UpdateUsers_Handler,
		},
		{
			MethodName: "DeleteUsers",
			Handler:    _Users_DeleteUsers_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Users_GetUsers_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Users_ListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/v1/users.proto",
}