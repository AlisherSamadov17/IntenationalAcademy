// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: teacher.proto

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
	TeacherService_Create_FullMethodName        = "/user_service.TeacherService/Create"
	TeacherService_GetByID_FullMethodName       = "/user_service.TeacherService/GetByID"
	TeacherService_GetList_FullMethodName       = "/user_service.TeacherService/GetList"
	TeacherService_Update_FullMethodName        = "/user_service.TeacherService/Update"
	TeacherService_Delete_FullMethodName        = "/user_service.TeacherService/Delete"
	TeacherService_GetReportList_FullMethodName = "/user_service.TeacherService/GetReportList"
)

// TeacherServiceClient is the client API for TeacherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeacherServiceClient interface {
	Create(ctx context.Context, in *CreateTeacherRequest, opts ...grpc.CallOption) (*TeacherResponse, error)
	GetByID(ctx context.Context, in *TeacherID, opts ...grpc.CallOption) (*GetTeacherResponse, error)
	GetList(ctx context.Context, in *GetListTeacherRequest, opts ...grpc.CallOption) (*GetListTeacherResponse, error)
	Update(ctx context.Context, in *UpdateTeacherRequest, opts ...grpc.CallOption) (*GetTeacherResponse, error)
	Delete(ctx context.Context, in *TeacherID, opts ...grpc.CallOption) (*TeacherEmpty, error)
	GetReportList(ctx context.Context, in *GetReportListTeacherRequest, opts ...grpc.CallOption) (*GetReportListTeacherResponse, error)
}

type teacherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeacherServiceClient(cc grpc.ClientConnInterface) TeacherServiceClient {
	return &teacherServiceClient{cc}
}

func (c *teacherServiceClient) Create(ctx context.Context, in *CreateTeacherRequest, opts ...grpc.CallOption) (*TeacherResponse, error) {
	out := new(TeacherResponse)
	err := c.cc.Invoke(ctx, TeacherService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) GetByID(ctx context.Context, in *TeacherID, opts ...grpc.CallOption) (*GetTeacherResponse, error) {
	out := new(GetTeacherResponse)
	err := c.cc.Invoke(ctx, TeacherService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) GetList(ctx context.Context, in *GetListTeacherRequest, opts ...grpc.CallOption) (*GetListTeacherResponse, error) {
	out := new(GetListTeacherResponse)
	err := c.cc.Invoke(ctx, TeacherService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) Update(ctx context.Context, in *UpdateTeacherRequest, opts ...grpc.CallOption) (*GetTeacherResponse, error) {
	out := new(GetTeacherResponse)
	err := c.cc.Invoke(ctx, TeacherService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) Delete(ctx context.Context, in *TeacherID, opts ...grpc.CallOption) (*TeacherEmpty, error) {
	out := new(TeacherEmpty)
	err := c.cc.Invoke(ctx, TeacherService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) GetReportList(ctx context.Context, in *GetReportListTeacherRequest, opts ...grpc.CallOption) (*GetReportListTeacherResponse, error) {
	out := new(GetReportListTeacherResponse)
	err := c.cc.Invoke(ctx, TeacherService_GetReportList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeacherServiceServer is the server API for TeacherService service.
// All implementations must embed UnimplementedTeacherServiceServer
// for forward compatibility
type TeacherServiceServer interface {
	Create(context.Context, *CreateTeacherRequest) (*TeacherResponse, error)
	GetByID(context.Context, *TeacherID) (*GetTeacherResponse, error)
	GetList(context.Context, *GetListTeacherRequest) (*GetListTeacherResponse, error)
	Update(context.Context, *UpdateTeacherRequest) (*GetTeacherResponse, error)
	Delete(context.Context, *TeacherID) (*TeacherEmpty, error)
	GetReportList(context.Context, *GetReportListTeacherRequest) (*GetReportListTeacherResponse, error)
	mustEmbedUnimplementedTeacherServiceServer()
}

// UnimplementedTeacherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeacherServiceServer struct {
}

func (UnimplementedTeacherServiceServer) Create(context.Context, *CreateTeacherRequest) (*TeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTeacherServiceServer) GetByID(context.Context, *TeacherID) (*GetTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedTeacherServiceServer) GetList(context.Context, *GetListTeacherRequest) (*GetListTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedTeacherServiceServer) Update(context.Context, *UpdateTeacherRequest) (*GetTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTeacherServiceServer) Delete(context.Context, *TeacherID) (*TeacherEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTeacherServiceServer) GetReportList(context.Context, *GetReportListTeacherRequest) (*GetReportListTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReportList not implemented")
}
func (UnimplementedTeacherServiceServer) mustEmbedUnimplementedTeacherServiceServer() {}

// UnsafeTeacherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeacherServiceServer will
// result in compilation errors.
type UnsafeTeacherServiceServer interface {
	mustEmbedUnimplementedTeacherServiceServer()
}

func RegisterTeacherServiceServer(s grpc.ServiceRegistrar, srv TeacherServiceServer) {
	s.RegisterService(&TeacherService_ServiceDesc, srv)
}

func _TeacherService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).Create(ctx, req.(*CreateTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeacherID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).GetByID(ctx, req.(*TeacherID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).GetList(ctx, req.(*GetListTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).Update(ctx, req.(*UpdateTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeacherID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).Delete(ctx, req.(*TeacherID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_GetReportList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportListTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).GetReportList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_GetReportList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).GetReportList(ctx, req.(*GetReportListTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeacherService_ServiceDesc is the grpc.ServiceDesc for TeacherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeacherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.TeacherService",
	HandlerType: (*TeacherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TeacherService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _TeacherService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _TeacherService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TeacherService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TeacherService_Delete_Handler,
		},
		{
			MethodName: "GetReportList",
			Handler:    _TeacherService_GetReportList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teacher.proto",
}
