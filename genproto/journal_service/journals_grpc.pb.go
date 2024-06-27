// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: journals.proto

package journal_service

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

// JournalServiceClient is the client API for JournalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JournalServiceClient interface {
	Create(ctx context.Context, in *CreateJournal, opts ...grpc.CallOption) (*GetJournal, error)
	GetByID(ctx context.Context, in *JournalPrimaryKey, opts ...grpc.CallOption) (*GetJournal, error)
	GetList(ctx context.Context, in *GetListJournalRequest, opts ...grpc.CallOption) (*GetListJournalResponse, error)
	Update(ctx context.Context, in *UpdateJournal, opts ...grpc.CallOption) (*GetJournal, error)
	Delete(ctx context.Context, in *JournalPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type journalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJournalServiceClient(cc grpc.ClientConnInterface) JournalServiceClient {
	return &journalServiceClient{cc}
}

func (c *journalServiceClient) Create(ctx context.Context, in *CreateJournal, opts ...grpc.CallOption) (*GetJournal, error) {
	out := new(GetJournal)
	err := c.cc.Invoke(ctx, "/journal_service_go.JournalService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) GetByID(ctx context.Context, in *JournalPrimaryKey, opts ...grpc.CallOption) (*GetJournal, error) {
	out := new(GetJournal)
	err := c.cc.Invoke(ctx, "/journal_service_go.JournalService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) GetList(ctx context.Context, in *GetListJournalRequest, opts ...grpc.CallOption) (*GetListJournalResponse, error) {
	out := new(GetListJournalResponse)
	err := c.cc.Invoke(ctx, "/journal_service_go.JournalService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) Update(ctx context.Context, in *UpdateJournal, opts ...grpc.CallOption) (*GetJournal, error) {
	out := new(GetJournal)
	err := c.cc.Invoke(ctx, "/journal_service_go.JournalService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) Delete(ctx context.Context, in *JournalPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/journal_service_go.JournalService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JournalServiceServer is the server API for JournalService service.
// All implementations must embed UnimplementedJournalServiceServer
// for forward compatibility
type JournalServiceServer interface {
	Create(context.Context, *CreateJournal) (*GetJournal, error)
	GetByID(context.Context, *JournalPrimaryKey) (*GetJournal, error)
	GetList(context.Context, *GetListJournalRequest) (*GetListJournalResponse, error)
	Update(context.Context, *UpdateJournal) (*GetJournal, error)
	Delete(context.Context, *JournalPrimaryKey) (*empty.Empty, error)
	mustEmbedUnimplementedJournalServiceServer()
}

// UnimplementedJournalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJournalServiceServer struct {
}

func (UnimplementedJournalServiceServer) Create(context.Context, *CreateJournal) (*GetJournal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedJournalServiceServer) GetByID(context.Context, *JournalPrimaryKey) (*GetJournal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedJournalServiceServer) GetList(context.Context, *GetListJournalRequest) (*GetListJournalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedJournalServiceServer) Update(context.Context, *UpdateJournal) (*GetJournal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedJournalServiceServer) Delete(context.Context, *JournalPrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedJournalServiceServer) mustEmbedUnimplementedJournalServiceServer() {}

// UnsafeJournalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JournalServiceServer will
// result in compilation errors.
type UnsafeJournalServiceServer interface {
	mustEmbedUnimplementedJournalServiceServer()
}

func RegisterJournalServiceServer(s grpc.ServiceRegistrar, srv JournalServiceServer) {
	s.RegisterService(&JournalService_ServiceDesc, srv)
}

func _JournalService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJournal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal_service_go.JournalService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).Create(ctx, req.(*CreateJournal))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JournalPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal_service_go.JournalService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).GetByID(ctx, req.(*JournalPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListJournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal_service_go.JournalService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).GetList(ctx, req.(*GetListJournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJournal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal_service_go.JournalService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).Update(ctx, req.(*UpdateJournal))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JournalPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal_service_go.JournalService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).Delete(ctx, req.(*JournalPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// JournalService_ServiceDesc is the grpc.ServiceDesc for JournalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JournalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "journal_service_go.JournalService",
	HandlerType: (*JournalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _JournalService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _JournalService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _JournalService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _JournalService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _JournalService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "journals.proto",
}
