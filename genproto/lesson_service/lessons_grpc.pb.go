// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: lessons.proto

package lesson_service

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

// LessonServiceClient is the client API for LessonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LessonServiceClient interface {
	Create(ctx context.Context, in *CreateLesson, opts ...grpc.CallOption) (*GetLesson, error)
	GetByID(ctx context.Context, in *LessonPrimaryKey, opts ...grpc.CallOption) (*GetLesson, error)
	GetList(ctx context.Context, in *GetListLessonRequest, opts ...grpc.CallOption) (*GetListLessonResponse, error)
	Update(ctx context.Context, in *UpdateLesson, opts ...grpc.CallOption) (*GetLesson, error)
	Delete(ctx context.Context, in *LessonPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type lessonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLessonServiceClient(cc grpc.ClientConnInterface) LessonServiceClient {
	return &lessonServiceClient{cc}
}

func (c *lessonServiceClient) Create(ctx context.Context, in *CreateLesson, opts ...grpc.CallOption) (*GetLesson, error) {
	out := new(GetLesson)
	err := c.cc.Invoke(ctx, "/lesson_service_go.LessonService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lessonServiceClient) GetByID(ctx context.Context, in *LessonPrimaryKey, opts ...grpc.CallOption) (*GetLesson, error) {
	out := new(GetLesson)
	err := c.cc.Invoke(ctx, "/lesson_service_go.LessonService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lessonServiceClient) GetList(ctx context.Context, in *GetListLessonRequest, opts ...grpc.CallOption) (*GetListLessonResponse, error) {
	out := new(GetListLessonResponse)
	err := c.cc.Invoke(ctx, "/lesson_service_go.LessonService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lessonServiceClient) Update(ctx context.Context, in *UpdateLesson, opts ...grpc.CallOption) (*GetLesson, error) {
	out := new(GetLesson)
	err := c.cc.Invoke(ctx, "/lesson_service_go.LessonService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lessonServiceClient) Delete(ctx context.Context, in *LessonPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/lesson_service_go.LessonService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LessonServiceServer is the server API for LessonService service.
// All implementations must embed UnimplementedLessonServiceServer
// for forward compatibility
type LessonServiceServer interface {
	Create(context.Context, *CreateLesson) (*GetLesson, error)
	GetByID(context.Context, *LessonPrimaryKey) (*GetLesson, error)
	GetList(context.Context, *GetListLessonRequest) (*GetListLessonResponse, error)
	Update(context.Context, *UpdateLesson) (*GetLesson, error)
	Delete(context.Context, *LessonPrimaryKey) (*empty.Empty, error)
	mustEmbedUnimplementedLessonServiceServer()
}

// UnimplementedLessonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLessonServiceServer struct {
}

func (UnimplementedLessonServiceServer) Create(context.Context, *CreateLesson) (*GetLesson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLessonServiceServer) GetByID(context.Context, *LessonPrimaryKey) (*GetLesson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedLessonServiceServer) GetList(context.Context, *GetListLessonRequest) (*GetListLessonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedLessonServiceServer) Update(context.Context, *UpdateLesson) (*GetLesson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLessonServiceServer) Delete(context.Context, *LessonPrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLessonServiceServer) mustEmbedUnimplementedLessonServiceServer() {}

// UnsafeLessonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LessonServiceServer will
// result in compilation errors.
type UnsafeLessonServiceServer interface {
	mustEmbedUnimplementedLessonServiceServer()
}

func RegisterLessonServiceServer(s grpc.ServiceRegistrar, srv LessonServiceServer) {
	s.RegisterService(&LessonService_ServiceDesc, srv)
}

func _LessonService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLesson)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lesson_service_go.LessonService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonServiceServer).Create(ctx, req.(*CreateLesson))
	}
	return interceptor(ctx, in, info, handler)
}

func _LessonService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LessonPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lesson_service_go.LessonService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonServiceServer).GetByID(ctx, req.(*LessonPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _LessonService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListLessonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lesson_service_go.LessonService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonServiceServer).GetList(ctx, req.(*GetListLessonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LessonService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLesson)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lesson_service_go.LessonService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonServiceServer).Update(ctx, req.(*UpdateLesson))
	}
	return interceptor(ctx, in, info, handler)
}

func _LessonService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LessonPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lesson_service_go.LessonService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonServiceServer).Delete(ctx, req.(*LessonPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// LessonService_ServiceDesc is the grpc.ServiceDesc for LessonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LessonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lesson_service_go.LessonService",
	HandlerType: (*LessonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LessonService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _LessonService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _LessonService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LessonService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LessonService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lessons.proto",
}