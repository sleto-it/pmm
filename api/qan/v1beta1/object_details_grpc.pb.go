// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: qan/v1beta1/object_details.proto

package qanv1beta1

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
	ObjectDetailsService_GetMetrics_FullMethodName                  = "/qan.v1beta1.ObjectDetailsService/GetMetrics"
	ObjectDetailsService_GetQueryExample_FullMethodName             = "/qan.v1beta1.ObjectDetailsService/GetQueryExample"
	ObjectDetailsService_GetLabels_FullMethodName                   = "/qan.v1beta1.ObjectDetailsService/GetLabels"
	ObjectDetailsService_GetQueryPlan_FullMethodName                = "/qan.v1beta1.ObjectDetailsService/GetQueryPlan"
	ObjectDetailsService_GetHistogram_FullMethodName                = "/qan.v1beta1.ObjectDetailsService/GetHistogram"
	ObjectDetailsService_QueryExists_FullMethodName                 = "/qan.v1beta1.ObjectDetailsService/QueryExists"
	ObjectDetailsService_ExplainFingerprintByQueryID_FullMethodName = "/qan.v1beta1.ObjectDetailsService/ExplainFingerprintByQueryID"
	ObjectDetailsService_SchemaByQueryID_FullMethodName             = "/qan.v1beta1.ObjectDetailsService/SchemaByQueryID"
)

// ObjectDetailsServiceClient is the client API for ObjectDetailsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectDetailsServiceClient interface {
	// GetMetrics gets map of metrics for specific filtering.
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	// GetQueryExample gets list of query examples.
	GetQueryExample(ctx context.Context, in *GetQueryExampleRequest, opts ...grpc.CallOption) (*GetQueryExampleResponse, error)
	// GetLabels gets list of labels for object details.
	GetLabels(ctx context.Context, in *GetLabelsRequest, opts ...grpc.CallOption) (*GetLabelsResponse, error)
	// GetQueryPlan gets query plan and plan id for specific filtering.
	GetQueryPlan(ctx context.Context, in *GetQueryPlanRequest, opts ...grpc.CallOption) (*GetQueryPlanResponse, error)
	// GetHistogram gets histogram items for specific filtering.
	GetHistogram(ctx context.Context, in *GetHistogramRequest, opts ...grpc.CallOption) (*GetHistogramResponse, error)
	// QueryExists check if query exists in clickhouse.
	QueryExists(ctx context.Context, in *QueryExistsRequest, opts ...grpc.CallOption) (*QueryExistsResponse, error)
	// ExplainFingerprintByQueryID get explain fingerprint for given query ID.
	ExplainFingerprintByQueryID(ctx context.Context, in *ExplainFingerprintByQueryIDRequest, opts ...grpc.CallOption) (*ExplainFingerprintByQueryIDResponse, error)
	// SchemaByQueryID returns schema for given queryID and serviceID.
	SchemaByQueryID(ctx context.Context, in *SchemaByQueryIDRequest, opts ...grpc.CallOption) (*SchemaByQueryIDResponse, error)
}

type objectDetailsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectDetailsServiceClient(cc grpc.ClientConnInterface) ObjectDetailsServiceClient {
	return &objectDetailsServiceClient{cc}
}

func (c *objectDetailsServiceClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, ObjectDetailsService_GetMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsServiceClient) GetQueryExample(ctx context.Context, in *GetQueryExampleRequest, opts ...grpc.CallOption) (*GetQueryExampleResponse, error) {
	out := new(GetQueryExampleResponse)
	err := c.cc.Invoke(ctx, ObjectDetailsService_GetQueryExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsServiceClient) GetLabels(ctx context.Context, in *GetLabelsRequest, opts ...grpc.CallOption) (*GetLabelsResponse, error) {
	out := new(GetLabelsResponse)
	err := c.cc.Invoke(ctx, ObjectDetailsService_GetLabels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsServiceClient) GetQueryPlan(ctx context.Context, in *GetQueryPlanRequest, opts ...grpc.CallOption) (*GetQueryPlanResponse, error) {
	out := new(GetQueryPlanResponse)
	err := c.cc.Invoke(ctx, ObjectDetailsService_GetQueryPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsServiceClient) GetHistogram(ctx context.Context, in *GetHistogramRequest, opts ...grpc.CallOption) (*GetHistogramResponse, error) {
	out := new(GetHistogramResponse)
	err := c.cc.Invoke(ctx, ObjectDetailsService_GetHistogram_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsServiceClient) QueryExists(ctx context.Context, in *QueryExistsRequest, opts ...grpc.CallOption) (*QueryExistsResponse, error) {
	out := new(QueryExistsResponse)
	err := c.cc.Invoke(ctx, ObjectDetailsService_QueryExists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsServiceClient) ExplainFingerprintByQueryID(ctx context.Context, in *ExplainFingerprintByQueryIDRequest, opts ...grpc.CallOption) (*ExplainFingerprintByQueryIDResponse, error) {
	out := new(ExplainFingerprintByQueryIDResponse)
	err := c.cc.Invoke(ctx, ObjectDetailsService_ExplainFingerprintByQueryID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsServiceClient) SchemaByQueryID(ctx context.Context, in *SchemaByQueryIDRequest, opts ...grpc.CallOption) (*SchemaByQueryIDResponse, error) {
	out := new(SchemaByQueryIDResponse)
	err := c.cc.Invoke(ctx, ObjectDetailsService_SchemaByQueryID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectDetailsServiceServer is the server API for ObjectDetailsService service.
// All implementations must embed UnimplementedObjectDetailsServiceServer
// for forward compatibility
type ObjectDetailsServiceServer interface {
	// GetMetrics gets map of metrics for specific filtering.
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	// GetQueryExample gets list of query examples.
	GetQueryExample(context.Context, *GetQueryExampleRequest) (*GetQueryExampleResponse, error)
	// GetLabels gets list of labels for object details.
	GetLabels(context.Context, *GetLabelsRequest) (*GetLabelsResponse, error)
	// GetQueryPlan gets query plan and plan id for specific filtering.
	GetQueryPlan(context.Context, *GetQueryPlanRequest) (*GetQueryPlanResponse, error)
	// GetHistogram gets histogram items for specific filtering.
	GetHistogram(context.Context, *GetHistogramRequest) (*GetHistogramResponse, error)
	// QueryExists check if query exists in clickhouse.
	QueryExists(context.Context, *QueryExistsRequest) (*QueryExistsResponse, error)
	// ExplainFingerprintByQueryID get explain fingerprint for given query ID.
	ExplainFingerprintByQueryID(context.Context, *ExplainFingerprintByQueryIDRequest) (*ExplainFingerprintByQueryIDResponse, error)
	// SchemaByQueryID returns schema for given queryID and serviceID.
	SchemaByQueryID(context.Context, *SchemaByQueryIDRequest) (*SchemaByQueryIDResponse, error)
	mustEmbedUnimplementedObjectDetailsServiceServer()
}

// UnimplementedObjectDetailsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedObjectDetailsServiceServer struct{}

func (UnimplementedObjectDetailsServiceServer) GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}

func (UnimplementedObjectDetailsServiceServer) GetQueryExample(context.Context, *GetQueryExampleRequest) (*GetQueryExampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryExample not implemented")
}

func (UnimplementedObjectDetailsServiceServer) GetLabels(context.Context, *GetLabelsRequest) (*GetLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabels not implemented")
}

func (UnimplementedObjectDetailsServiceServer) GetQueryPlan(context.Context, *GetQueryPlanRequest) (*GetQueryPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryPlan not implemented")
}

func (UnimplementedObjectDetailsServiceServer) GetHistogram(context.Context, *GetHistogramRequest) (*GetHistogramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistogram not implemented")
}

func (UnimplementedObjectDetailsServiceServer) QueryExists(context.Context, *QueryExistsRequest) (*QueryExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryExists not implemented")
}

func (UnimplementedObjectDetailsServiceServer) ExplainFingerprintByQueryID(context.Context, *ExplainFingerprintByQueryIDRequest) (*ExplainFingerprintByQueryIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExplainFingerprintByQueryID not implemented")
}

func (UnimplementedObjectDetailsServiceServer) SchemaByQueryID(context.Context, *SchemaByQueryIDRequest) (*SchemaByQueryIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchemaByQueryID not implemented")
}
func (UnimplementedObjectDetailsServiceServer) mustEmbedUnimplementedObjectDetailsServiceServer() {}

// UnsafeObjectDetailsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectDetailsServiceServer will
// result in compilation errors.
type UnsafeObjectDetailsServiceServer interface {
	mustEmbedUnimplementedObjectDetailsServiceServer()
}

func RegisterObjectDetailsServiceServer(s grpc.ServiceRegistrar, srv ObjectDetailsServiceServer) {
	s.RegisterService(&ObjectDetailsService_ServiceDesc, srv)
}

func _ObjectDetailsService_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServiceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectDetailsService_GetMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServiceServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetailsService_GetQueryExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueryExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServiceServer).GetQueryExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectDetailsService_GetQueryExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServiceServer).GetQueryExample(ctx, req.(*GetQueryExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetailsService_GetLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServiceServer).GetLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectDetailsService_GetLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServiceServer).GetLabels(ctx, req.(*GetLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetailsService_GetQueryPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueryPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServiceServer).GetQueryPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectDetailsService_GetQueryPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServiceServer).GetQueryPlan(ctx, req.(*GetQueryPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetailsService_GetHistogram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistogramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServiceServer).GetHistogram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectDetailsService_GetHistogram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServiceServer).GetHistogram(ctx, req.(*GetHistogramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetailsService_QueryExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServiceServer).QueryExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectDetailsService_QueryExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServiceServer).QueryExists(ctx, req.(*QueryExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetailsService_ExplainFingerprintByQueryID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExplainFingerprintByQueryIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServiceServer).ExplainFingerprintByQueryID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectDetailsService_ExplainFingerprintByQueryID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServiceServer).ExplainFingerprintByQueryID(ctx, req.(*ExplainFingerprintByQueryIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetailsService_SchemaByQueryID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaByQueryIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServiceServer).SchemaByQueryID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectDetailsService_SchemaByQueryID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServiceServer).SchemaByQueryID(ctx, req.(*SchemaByQueryIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectDetailsService_ServiceDesc is the grpc.ServiceDesc for ObjectDetailsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectDetailsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qan.v1beta1.ObjectDetailsService",
	HandlerType: (*ObjectDetailsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetrics",
			Handler:    _ObjectDetailsService_GetMetrics_Handler,
		},
		{
			MethodName: "GetQueryExample",
			Handler:    _ObjectDetailsService_GetQueryExample_Handler,
		},
		{
			MethodName: "GetLabels",
			Handler:    _ObjectDetailsService_GetLabels_Handler,
		},
		{
			MethodName: "GetQueryPlan",
			Handler:    _ObjectDetailsService_GetQueryPlan_Handler,
		},
		{
			MethodName: "GetHistogram",
			Handler:    _ObjectDetailsService_GetHistogram_Handler,
		},
		{
			MethodName: "QueryExists",
			Handler:    _ObjectDetailsService_QueryExists_Handler,
		},
		{
			MethodName: "ExplainFingerprintByQueryID",
			Handler:    _ObjectDetailsService_ExplainFingerprintByQueryID_Handler,
		},
		{
			MethodName: "SchemaByQueryID",
			Handler:    _ObjectDetailsService_SchemaByQueryID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qan/v1beta1/object_details.proto",
}
