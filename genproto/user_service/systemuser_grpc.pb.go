// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: systemuser.proto

package user_service

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

const (
	SystemUserService_Create_FullMethodName     = "/user_service.SystemUserService/Create"
	SystemUserService_GetByID_FullMethodName    = "/user_service.SystemUserService/GetByID"
	SystemUserService_GetList_FullMethodName    = "/user_service.SystemUserService/GetList"
	SystemUserService_Update_FullMethodName     = "/user_service.SystemUserService/Update"
	SystemUserService_Delete_FullMethodName     = "/user_service.SystemUserService/Delete"
	SystemUserService_GetByGmail_FullMethodName = "/user_service.SystemUserService/GetByGmail"
)

// SystemUserServiceClient is the client API for SystemUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemUserServiceClient interface {
	Create(ctx context.Context, in *CreateSystemUser, opts ...grpc.CallOption) (*SystemUserPrimaryKey, error)
	GetByID(ctx context.Context, in *SystemUserPrimaryKey, opts ...grpc.CallOption) (*SystemUser, error)
	GetList(ctx context.Context, in *GetListSystemUserRequest, opts ...grpc.CallOption) (*GetListSystemUserResponse, error)
	Update(ctx context.Context, in *UpdateSystemUserRequest, opts ...grpc.CallOption) (*UpdateSystemUserResponse, error)
	Delete(ctx context.Context, in *SystemUserPrimaryKey, opts ...grpc.CallOption) (*SystemUserEmpty, error)
	GetByGmail(ctx context.Context, in *SystemUserGmail, opts ...grpc.CallOption) (*SystemUserPrimaryKey, error)
}

type systemUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemUserServiceClient(cc grpc.ClientConnInterface) SystemUserServiceClient {
	return &systemUserServiceClient{cc}
}

func (c *systemUserServiceClient) Create(ctx context.Context, in *CreateSystemUser, opts ...grpc.CallOption) (*SystemUserPrimaryKey, error) {
	out := new(SystemUserPrimaryKey)
	err := c.cc.Invoke(ctx, SystemUserService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) GetByID(ctx context.Context, in *SystemUserPrimaryKey, opts ...grpc.CallOption) (*SystemUser, error) {
	out := new(SystemUser)
	err := c.cc.Invoke(ctx, SystemUserService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) GetList(ctx context.Context, in *GetListSystemUserRequest, opts ...grpc.CallOption) (*GetListSystemUserResponse, error) {
	out := new(GetListSystemUserResponse)
	err := c.cc.Invoke(ctx, SystemUserService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) Update(ctx context.Context, in *UpdateSystemUserRequest, opts ...grpc.CallOption) (*UpdateSystemUserResponse, error) {
	out := new(UpdateSystemUserResponse)
	err := c.cc.Invoke(ctx, SystemUserService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) Delete(ctx context.Context, in *SystemUserPrimaryKey, opts ...grpc.CallOption) (*SystemUserEmpty, error) {
	out := new(SystemUserEmpty)
	err := c.cc.Invoke(ctx, SystemUserService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) GetByGmail(ctx context.Context, in *SystemUserGmail, opts ...grpc.CallOption) (*SystemUserPrimaryKey, error) {
	out := new(SystemUserPrimaryKey)
	err := c.cc.Invoke(ctx, SystemUserService_GetByGmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemUserServiceServer is the server API for SystemUserService service.
// All implementations should embed UnimplementedSystemUserServiceServer
// for forward compatibility
type SystemUserServiceServer interface {
	Create(context.Context, *CreateSystemUser) (*SystemUserPrimaryKey, error)
	GetByID(context.Context, *SystemUserPrimaryKey) (*SystemUser, error)
	GetList(context.Context, *GetListSystemUserRequest) (*GetListSystemUserResponse, error)
	Update(context.Context, *UpdateSystemUserRequest) (*UpdateSystemUserResponse, error)
	Delete(context.Context, *SystemUserPrimaryKey) (*SystemUserEmpty, error)
	GetByGmail(context.Context, *SystemUserGmail) (*SystemUserPrimaryKey, error)
}

// UnimplementedSystemUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSystemUserServiceServer struct {
}

func (UnimplementedSystemUserServiceServer) Create(context.Context, *CreateSystemUser) (*SystemUserPrimaryKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSystemUserServiceServer) GetByID(context.Context, *SystemUserPrimaryKey) (*SystemUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedSystemUserServiceServer) GetList(context.Context, *GetListSystemUserRequest) (*GetListSystemUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedSystemUserServiceServer) Update(context.Context, *UpdateSystemUserRequest) (*UpdateSystemUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSystemUserServiceServer) Delete(context.Context, *SystemUserPrimaryKey) (*SystemUserEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSystemUserServiceServer) GetByGmail(context.Context, *SystemUserGmail) (*SystemUserPrimaryKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByGmail not implemented")
}

// UnsafeSystemUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemUserServiceServer will
// result in compilation errors.
type UnsafeSystemUserServiceServer interface {
	mustEmbedUnimplementedSystemUserServiceServer()
}

func RegisterSystemUserServiceServer(s grpc.ServiceRegistrar, srv SystemUserServiceServer) {
	s.RegisterService(&SystemUserService_ServiceDesc, srv)
}

func _SystemUserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSystemUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Create(ctx, req.(*CreateSystemUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).GetByID(ctx, req.(*SystemUserPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListSystemUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).GetList(ctx, req.(*GetListSystemUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSystemUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Update(ctx, req.(*UpdateSystemUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Delete(ctx, req.(*SystemUserPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_GetByGmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserGmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).GetByGmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_GetByGmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).GetByGmail(ctx, req.(*SystemUserGmail))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemUserService_ServiceDesc is the grpc.ServiceDesc for SystemUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.SystemUserService",
	HandlerType: (*SystemUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SystemUserService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _SystemUserService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _SystemUserService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SystemUserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SystemUserService_Delete_Handler,
		},
		{
			MethodName: "GetByGmail",
			Handler:    _SystemUserService_GetByGmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "systemuser.proto",
}
