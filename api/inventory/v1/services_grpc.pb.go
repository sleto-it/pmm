// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: inventory/v1/services.proto

package inventoryv1

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
	ServicesService_ListServices_FullMethodName           = "/inventory.v1.ServicesService/ListServices"
	ServicesService_ListActiveServiceTypes_FullMethodName = "/inventory.v1.ServicesService/ListActiveServiceTypes"
	ServicesService_GetService_FullMethodName             = "/inventory.v1.ServicesService/GetService"
	ServicesService_AddService_FullMethodName             = "/inventory.v1.ServicesService/AddService"
	ServicesService_RemoveService_FullMethodName          = "/inventory.v1.ServicesService/RemoveService"
	ServicesService_ChangeService_FullMethodName          = "/inventory.v1.ServicesService/ChangeService"
)

// ServicesServiceClient is the client API for ServicesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesServiceClient interface {
	// ListServices returns a list of Services filtered by type.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	// ListActiveServiceTypes returns a list of active Services.
	ListActiveServiceTypes(ctx context.Context, in *ListActiveServiceTypesRequest, opts ...grpc.CallOption) (*ListActiveServiceTypesResponse, error)
	// GetService returns a single Service by ID.
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
	// AddService adds any type of Service.
	AddService(ctx context.Context, in *AddServiceRequest, opts ...grpc.CallOption) (*AddServiceResponse, error)
	// RemoveService removes a Service.
	RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error)
	// ChangeService allows changing configuration of a Service.
	ChangeService(ctx context.Context, in *ChangeServiceRequest, opts ...grpc.CallOption) (*ChangeServiceResponse, error)
}

type servicesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesServiceClient(cc grpc.ClientConnInterface) ServicesServiceClient {
	return &servicesServiceClient{cc}
}

func (c *servicesServiceClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, ServicesService_ListServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) ListActiveServiceTypes(ctx context.Context, in *ListActiveServiceTypesRequest, opts ...grpc.CallOption) (*ListActiveServiceTypesResponse, error) {
	out := new(ListActiveServiceTypesResponse)
	err := c.cc.Invoke(ctx, ServicesService_ListActiveServiceTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_GetService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) AddService(ctx context.Context, in *AddServiceRequest, opts ...grpc.CallOption) (*AddServiceResponse, error) {
	out := new(AddServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_AddService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error) {
	out := new(RemoveServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_RemoveService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) ChangeService(ctx context.Context, in *ChangeServiceRequest, opts ...grpc.CallOption) (*ChangeServiceResponse, error) {
	out := new(ChangeServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_ChangeService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesServiceServer is the server API for ServicesService service.
// All implementations must embed UnimplementedServicesServiceServer
// for forward compatibility
type ServicesServiceServer interface {
	// ListServices returns a list of Services filtered by type.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	// ListActiveServiceTypes returns a list of active Services.
	ListActiveServiceTypes(context.Context, *ListActiveServiceTypesRequest) (*ListActiveServiceTypesResponse, error)
	// GetService returns a single Service by ID.
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
	// AddService adds any type of Service.
	AddService(context.Context, *AddServiceRequest) (*AddServiceResponse, error)
	// RemoveService removes a Service.
	RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error)
	// ChangeService allows changing configuration of a Service.
	ChangeService(context.Context, *ChangeServiceRequest) (*ChangeServiceResponse, error)
	mustEmbedUnimplementedServicesServiceServer()
}

// UnimplementedServicesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServicesServiceServer struct{}

func (UnimplementedServicesServiceServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}

func (UnimplementedServicesServiceServer) ListActiveServiceTypes(context.Context, *ListActiveServiceTypesRequest) (*ListActiveServiceTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActiveServiceTypes not implemented")
}

func (UnimplementedServicesServiceServer) GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}

func (UnimplementedServicesServiceServer) AddService(context.Context, *AddServiceRequest) (*AddServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddService not implemented")
}

func (UnimplementedServicesServiceServer) RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveService not implemented")
}

func (UnimplementedServicesServiceServer) ChangeService(context.Context, *ChangeServiceRequest) (*ChangeServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeService not implemented")
}
func (UnimplementedServicesServiceServer) mustEmbedUnimplementedServicesServiceServer() {}

// UnsafeServicesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServiceServer will
// result in compilation errors.
type UnsafeServicesServiceServer interface {
	mustEmbedUnimplementedServicesServiceServer()
}

func RegisterServicesServiceServer(s grpc.ServiceRegistrar, srv ServicesServiceServer) {
	s.RegisterService(&ServicesService_ServiceDesc, srv)
}

func _ServicesService_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_ListServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_ListActiveServiceTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActiveServiceTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).ListActiveServiceTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_ListActiveServiceTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).ListActiveServiceTypes(ctx, req.(*ListActiveServiceTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_GetService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_AddService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).AddService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_AddService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).AddService(ctx, req.(*AddServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_RemoveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).RemoveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_RemoveService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).RemoveService(ctx, req.(*RemoveServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_ChangeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).ChangeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_ChangeService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).ChangeService(ctx, req.(*ChangeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicesService_ServiceDesc is the grpc.ServiceDesc for ServicesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.v1.ServicesService",
	HandlerType: (*ServicesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _ServicesService_ListServices_Handler,
		},
		{
			MethodName: "ListActiveServiceTypes",
			Handler:    _ServicesService_ListActiveServiceTypes_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _ServicesService_GetService_Handler,
		},
		{
			MethodName: "AddService",
			Handler:    _ServicesService_AddService_Handler,
		},
		{
			MethodName: "RemoveService",
			Handler:    _ServicesService_RemoveService_Handler,
		},
		{
			MethodName: "ChangeService",
			Handler:    _ServicesService_ChangeService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/v1/services.proto",
}
