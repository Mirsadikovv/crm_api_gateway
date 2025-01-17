// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: support_teachers.proto

package support_teacher_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SupportTeacherServiceClient is the client API for SupportTeacherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupportTeacherServiceClient interface {
	Create(ctx context.Context, in *CreateSupportTeacher, opts ...grpc.CallOption) (*GetSupportTeacher, error)
	GetByID(ctx context.Context, in *SupportTeacherPrimaryKey, opts ...grpc.CallOption) (*GetSupportTeacher, error)
	GetList(ctx context.Context, in *GetListSupportTeacherRequest, opts ...grpc.CallOption) (*GetListSupportTeacherResponse, error)
	Update(ctx context.Context, in *UpdateSupportTeacher, opts ...grpc.CallOption) (*GetSupportTeacher, error)
	Delete(ctx context.Context, in *SupportTeacherPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
	Login(ctx context.Context, in *SupportTeacherLoginRequest, opts ...grpc.CallOption) (*SupportTeacherLoginResponse, error)
	Register(ctx context.Context, in *SupportTeacherRegisterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RegisterConfirm(ctx context.Context, in *SupportTeacherRegisterConfRequest, opts ...grpc.CallOption) (*SupportTeacherLoginResponse, error)
	ChangePassword(ctx context.Context, in *SupportTeacherChangePassword, opts ...grpc.CallOption) (*SupportTeacherChangePasswordResp, error)
}

type supportTeacherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupportTeacherServiceClient(cc grpc.ClientConnInterface) SupportTeacherServiceClient {
	return &supportTeacherServiceClient{cc}
}

func (c *supportTeacherServiceClient) Create(ctx context.Context, in *CreateSupportTeacher, opts ...grpc.CallOption) (*GetSupportTeacher, error) {
	out := new(GetSupportTeacher)
	err := c.cc.Invoke(ctx, "/support_teacher_service_go.SupportTeacherService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportTeacherServiceClient) GetByID(ctx context.Context, in *SupportTeacherPrimaryKey, opts ...grpc.CallOption) (*GetSupportTeacher, error) {
	out := new(GetSupportTeacher)
	err := c.cc.Invoke(ctx, "/support_teacher_service_go.SupportTeacherService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportTeacherServiceClient) GetList(ctx context.Context, in *GetListSupportTeacherRequest, opts ...grpc.CallOption) (*GetListSupportTeacherResponse, error) {
	out := new(GetListSupportTeacherResponse)
	err := c.cc.Invoke(ctx, "/support_teacher_service_go.SupportTeacherService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportTeacherServiceClient) Update(ctx context.Context, in *UpdateSupportTeacher, opts ...grpc.CallOption) (*GetSupportTeacher, error) {
	out := new(GetSupportTeacher)
	err := c.cc.Invoke(ctx, "/support_teacher_service_go.SupportTeacherService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportTeacherServiceClient) Delete(ctx context.Context, in *SupportTeacherPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/support_teacher_service_go.SupportTeacherService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportTeacherServiceClient) Login(ctx context.Context, in *SupportTeacherLoginRequest, opts ...grpc.CallOption) (*SupportTeacherLoginResponse, error) {
	out := new(SupportTeacherLoginResponse)
	err := c.cc.Invoke(ctx, "/support_teacher_service_go.SupportTeacherService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportTeacherServiceClient) Register(ctx context.Context, in *SupportTeacherRegisterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/support_teacher_service_go.SupportTeacherService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportTeacherServiceClient) RegisterConfirm(ctx context.Context, in *SupportTeacherRegisterConfRequest, opts ...grpc.CallOption) (*SupportTeacherLoginResponse, error) {
	out := new(SupportTeacherLoginResponse)
	err := c.cc.Invoke(ctx, "/support_teacher_service_go.SupportTeacherService/RegisterConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportTeacherServiceClient) ChangePassword(ctx context.Context, in *SupportTeacherChangePassword, opts ...grpc.CallOption) (*SupportTeacherChangePasswordResp, error) {
	out := new(SupportTeacherChangePasswordResp)
	err := c.cc.Invoke(ctx, "/support_teacher_service_go.SupportTeacherService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupportTeacherServiceServer is the server API for SupportTeacherService service.
// All implementations must embed UnimplementedSupportTeacherServiceServer
// for forward compatibility
type SupportTeacherServiceServer interface {
	Create(context.Context, *CreateSupportTeacher) (*GetSupportTeacher, error)
	GetByID(context.Context, *SupportTeacherPrimaryKey) (*GetSupportTeacher, error)
	GetList(context.Context, *GetListSupportTeacherRequest) (*GetListSupportTeacherResponse, error)
	Update(context.Context, *UpdateSupportTeacher) (*GetSupportTeacher, error)
	Delete(context.Context, *SupportTeacherPrimaryKey) (*empty.Empty, error)
	Login(context.Context, *SupportTeacherLoginRequest) (*SupportTeacherLoginResponse, error)
	Register(context.Context, *SupportTeacherRegisterRequest) (*empty.Empty, error)
	RegisterConfirm(context.Context, *SupportTeacherRegisterConfRequest) (*SupportTeacherLoginResponse, error)
	ChangePassword(context.Context, *SupportTeacherChangePassword) (*SupportTeacherChangePasswordResp, error)
	mustEmbedUnimplementedSupportTeacherServiceServer()
}

// UnimplementedSupportTeacherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupportTeacherServiceServer struct {
}

func (UnimplementedSupportTeacherServiceServer) Create(context.Context, *CreateSupportTeacher) (*GetSupportTeacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSupportTeacherServiceServer) GetByID(context.Context, *SupportTeacherPrimaryKey) (*GetSupportTeacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedSupportTeacherServiceServer) GetList(context.Context, *GetListSupportTeacherRequest) (*GetListSupportTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedSupportTeacherServiceServer) Update(context.Context, *UpdateSupportTeacher) (*GetSupportTeacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSupportTeacherServiceServer) Delete(context.Context, *SupportTeacherPrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSupportTeacherServiceServer) Login(context.Context, *SupportTeacherLoginRequest) (*SupportTeacherLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSupportTeacherServiceServer) Register(context.Context, *SupportTeacherRegisterRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedSupportTeacherServiceServer) RegisterConfirm(context.Context, *SupportTeacherRegisterConfRequest) (*SupportTeacherLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterConfirm not implemented")
}
func (UnimplementedSupportTeacherServiceServer) ChangePassword(context.Context, *SupportTeacherChangePassword) (*SupportTeacherChangePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedSupportTeacherServiceServer) mustEmbedUnimplementedSupportTeacherServiceServer() {}

// UnsafeSupportTeacherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupportTeacherServiceServer will
// result in compilation errors.
type UnsafeSupportTeacherServiceServer interface {
	mustEmbedUnimplementedSupportTeacherServiceServer()
}

func RegisterSupportTeacherServiceServer(s grpc.ServiceRegistrar, srv SupportTeacherServiceServer) {
	s.RegisterService(&SupportTeacherService_ServiceDesc, srv)
}

func _SupportTeacherService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSupportTeacher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportTeacherServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/support_teacher_service_go.SupportTeacherService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportTeacherServiceServer).Create(ctx, req.(*CreateSupportTeacher))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportTeacherService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportTeacherPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportTeacherServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/support_teacher_service_go.SupportTeacherService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportTeacherServiceServer).GetByID(ctx, req.(*SupportTeacherPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportTeacherService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListSupportTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportTeacherServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/support_teacher_service_go.SupportTeacherService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportTeacherServiceServer).GetList(ctx, req.(*GetListSupportTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportTeacherService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSupportTeacher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportTeacherServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/support_teacher_service_go.SupportTeacherService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportTeacherServiceServer).Update(ctx, req.(*UpdateSupportTeacher))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportTeacherService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportTeacherPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportTeacherServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/support_teacher_service_go.SupportTeacherService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportTeacherServiceServer).Delete(ctx, req.(*SupportTeacherPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportTeacherService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportTeacherLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportTeacherServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/support_teacher_service_go.SupportTeacherService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportTeacherServiceServer).Login(ctx, req.(*SupportTeacherLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportTeacherService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportTeacherRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportTeacherServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/support_teacher_service_go.SupportTeacherService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportTeacherServiceServer).Register(ctx, req.(*SupportTeacherRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportTeacherService_RegisterConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportTeacherRegisterConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportTeacherServiceServer).RegisterConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/support_teacher_service_go.SupportTeacherService/RegisterConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportTeacherServiceServer).RegisterConfirm(ctx, req.(*SupportTeacherRegisterConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportTeacherService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportTeacherChangePassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportTeacherServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/support_teacher_service_go.SupportTeacherService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportTeacherServiceServer).ChangePassword(ctx, req.(*SupportTeacherChangePassword))
	}
	return interceptor(ctx, in, info, handler)
}

// SupportTeacherService_ServiceDesc is the grpc.ServiceDesc for SupportTeacherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupportTeacherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "support_teacher_service_go.SupportTeacherService",
	HandlerType: (*SupportTeacherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SupportTeacherService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _SupportTeacherService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _SupportTeacherService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SupportTeacherService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SupportTeacherService_Delete_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _SupportTeacherService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _SupportTeacherService_Register_Handler,
		},
		{
			MethodName: "RegisterConfirm",
			Handler:    _SupportTeacherService_RegisterConfirm_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _SupportTeacherService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "support_teachers.proto",
}
