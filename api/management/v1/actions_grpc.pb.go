// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: management/v1/actions.proto

package managementv1

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
	ActionsService_GetAction_FullMethodName                              = "/management.v1.ActionsService/GetAction"
	ActionsService_StartMySQLExplainAction_FullMethodName                = "/management.v1.ActionsService/StartMySQLExplainAction"
	ActionsService_StartMySQLExplainJSONAction_FullMethodName            = "/management.v1.ActionsService/StartMySQLExplainJSONAction"
	ActionsService_StartMySQLExplainTraditionalJSONAction_FullMethodName = "/management.v1.ActionsService/StartMySQLExplainTraditionalJSONAction"
	ActionsService_StartMySQLShowCreateTableAction_FullMethodName        = "/management.v1.ActionsService/StartMySQLShowCreateTableAction"
	ActionsService_StartMySQLShowTableStatusAction_FullMethodName        = "/management.v1.ActionsService/StartMySQLShowTableStatusAction"
	ActionsService_StartMySQLShowIndexAction_FullMethodName              = "/management.v1.ActionsService/StartMySQLShowIndexAction"
	ActionsService_StartPostgreSQLShowCreateTableAction_FullMethodName   = "/management.v1.ActionsService/StartPostgreSQLShowCreateTableAction"
	ActionsService_StartPostgreSQLShowIndexAction_FullMethodName         = "/management.v1.ActionsService/StartPostgreSQLShowIndexAction"
	ActionsService_StartMongoDBExplainAction_FullMethodName              = "/management.v1.ActionsService/StartMongoDBExplainAction"
	ActionsService_StartPTSummaryAction_FullMethodName                   = "/management.v1.ActionsService/StartPTSummaryAction"
	ActionsService_StartPTPgSummaryAction_FullMethodName                 = "/management.v1.ActionsService/StartPTPgSummaryAction"
	ActionsService_StartPTMongoDBSummaryAction_FullMethodName            = "/management.v1.ActionsService/StartPTMongoDBSummaryAction"
	ActionsService_StartPTMySQLSummaryAction_FullMethodName              = "/management.v1.ActionsService/StartPTMySQLSummaryAction"
	ActionsService_CancelAction_FullMethodName                           = "/management.v1.ActionsService/CancelAction"
)

// ActionsServiceClient is the client API for ActionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionsServiceClient interface {
	// GetAction gets result of a given Action.
	GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error)
	// StartMySQLExplainAction starts MySQL EXPLAIN Action with traditional output.
	StartMySQLExplainAction(ctx context.Context, in *StartMySQLExplainActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainActionResponse, error)
	// StartMySQLExplainJSONAction starts MySQL EXPLAIN Action with JSON output.
	StartMySQLExplainJSONAction(ctx context.Context, in *StartMySQLExplainJSONActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainJSONActionResponse, error)
	// StartMySQLExplainTraditionalJSONAction starts MySQL EXPLAIN Action with traditional JSON output.
	StartMySQLExplainTraditionalJSONAction(ctx context.Context, in *StartMySQLExplainTraditionalJSONActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainTraditionalJSONActionResponse, error)
	// StartMySQLShowCreateTableAction starts MySQL SHOW CREATE TABLE Action.
	StartMySQLShowCreateTableAction(ctx context.Context, in *StartMySQLShowCreateTableActionRequest, opts ...grpc.CallOption) (*StartMySQLShowCreateTableActionResponse, error)
	// StartMySQLShowTableStatusAction starts MySQL SHOW TABLE STATUS Action.
	StartMySQLShowTableStatusAction(ctx context.Context, in *StartMySQLShowTableStatusActionRequest, opts ...grpc.CallOption) (*StartMySQLShowTableStatusActionResponse, error)
	// StartMySQLShowIndexAction starts MySQL SHOW INDEX Action.
	StartMySQLShowIndexAction(ctx context.Context, in *StartMySQLShowIndexActionRequest, opts ...grpc.CallOption) (*StartMySQLShowIndexActionResponse, error)
	// StartPostgreSQLShowCreateTableAction starts PostgreSQL SHOW CREATE TABLE Action.
	StartPostgreSQLShowCreateTableAction(ctx context.Context, in *StartPostgreSQLShowCreateTableActionRequest, opts ...grpc.CallOption) (*StartPostgreSQLShowCreateTableActionResponse, error)
	// StartPostgreSQLShowIndexAction starts PostgreSQL SHOW INDEX Action.
	StartPostgreSQLShowIndexAction(ctx context.Context, in *StartPostgreSQLShowIndexActionRequest, opts ...grpc.CallOption) (*StartPostgreSQLShowIndexActionResponse, error)
	// StartMongoDBExplainAction starts MongoDB EXPLAIN Action.
	StartMongoDBExplainAction(ctx context.Context, in *StartMongoDBExplainActionRequest, opts ...grpc.CallOption) (*StartMongoDBExplainActionResponse, error)
	// StartPTSummaryAction starts pt-summary Action.
	StartPTSummaryAction(ctx context.Context, in *StartPTSummaryActionRequest, opts ...grpc.CallOption) (*StartPTSummaryActionResponse, error)
	// StartPTPgSummaryAction starts pt-pg-summary Action.
	StartPTPgSummaryAction(ctx context.Context, in *StartPTPgSummaryActionRequest, opts ...grpc.CallOption) (*StartPTPgSummaryActionResponse, error)
	// StartPTMongoDBSummaryAction starts pt-mongodb-summary Action.
	StartPTMongoDBSummaryAction(ctx context.Context, in *StartPTMongoDBSummaryActionRequest, opts ...grpc.CallOption) (*StartPTMongoDBSummaryActionResponse, error)
	// StartPTMySQLSummaryAction starts pt-mysql-summary Action.
	StartPTMySQLSummaryAction(ctx context.Context, in *StartPTMySQLSummaryActionRequest, opts ...grpc.CallOption) (*StartPTMySQLSummaryActionResponse, error)
	// CancelAction stops an Action.
	CancelAction(ctx context.Context, in *CancelActionRequest, opts ...grpc.CallOption) (*CancelActionResponse, error)
}

type actionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionsServiceClient(cc grpc.ClientConnInterface) ActionsServiceClient {
	return &actionsServiceClient{cc}
}

func (c *actionsServiceClient) GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error) {
	out := new(GetActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_GetAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartMySQLExplainAction(ctx context.Context, in *StartMySQLExplainActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainActionResponse, error) {
	out := new(StartMySQLExplainActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartMySQLExplainAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartMySQLExplainJSONAction(ctx context.Context, in *StartMySQLExplainJSONActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainJSONActionResponse, error) {
	out := new(StartMySQLExplainJSONActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartMySQLExplainJSONAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartMySQLExplainTraditionalJSONAction(ctx context.Context, in *StartMySQLExplainTraditionalJSONActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainTraditionalJSONActionResponse, error) {
	out := new(StartMySQLExplainTraditionalJSONActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartMySQLExplainTraditionalJSONAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartMySQLShowCreateTableAction(ctx context.Context, in *StartMySQLShowCreateTableActionRequest, opts ...grpc.CallOption) (*StartMySQLShowCreateTableActionResponse, error) {
	out := new(StartMySQLShowCreateTableActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartMySQLShowCreateTableAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartMySQLShowTableStatusAction(ctx context.Context, in *StartMySQLShowTableStatusActionRequest, opts ...grpc.CallOption) (*StartMySQLShowTableStatusActionResponse, error) {
	out := new(StartMySQLShowTableStatusActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartMySQLShowTableStatusAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartMySQLShowIndexAction(ctx context.Context, in *StartMySQLShowIndexActionRequest, opts ...grpc.CallOption) (*StartMySQLShowIndexActionResponse, error) {
	out := new(StartMySQLShowIndexActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartMySQLShowIndexAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartPostgreSQLShowCreateTableAction(ctx context.Context, in *StartPostgreSQLShowCreateTableActionRequest, opts ...grpc.CallOption) (*StartPostgreSQLShowCreateTableActionResponse, error) {
	out := new(StartPostgreSQLShowCreateTableActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartPostgreSQLShowCreateTableAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartPostgreSQLShowIndexAction(ctx context.Context, in *StartPostgreSQLShowIndexActionRequest, opts ...grpc.CallOption) (*StartPostgreSQLShowIndexActionResponse, error) {
	out := new(StartPostgreSQLShowIndexActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartPostgreSQLShowIndexAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartMongoDBExplainAction(ctx context.Context, in *StartMongoDBExplainActionRequest, opts ...grpc.CallOption) (*StartMongoDBExplainActionResponse, error) {
	out := new(StartMongoDBExplainActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartMongoDBExplainAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartPTSummaryAction(ctx context.Context, in *StartPTSummaryActionRequest, opts ...grpc.CallOption) (*StartPTSummaryActionResponse, error) {
	out := new(StartPTSummaryActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartPTSummaryAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartPTPgSummaryAction(ctx context.Context, in *StartPTPgSummaryActionRequest, opts ...grpc.CallOption) (*StartPTPgSummaryActionResponse, error) {
	out := new(StartPTPgSummaryActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartPTPgSummaryAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartPTMongoDBSummaryAction(ctx context.Context, in *StartPTMongoDBSummaryActionRequest, opts ...grpc.CallOption) (*StartPTMongoDBSummaryActionResponse, error) {
	out := new(StartPTMongoDBSummaryActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartPTMongoDBSummaryAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartPTMySQLSummaryAction(ctx context.Context, in *StartPTMySQLSummaryActionRequest, opts ...grpc.CallOption) (*StartPTMySQLSummaryActionResponse, error) {
	out := new(StartPTMySQLSummaryActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartPTMySQLSummaryAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) CancelAction(ctx context.Context, in *CancelActionRequest, opts ...grpc.CallOption) (*CancelActionResponse, error) {
	out := new(CancelActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_CancelAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionsServiceServer is the server API for ActionsService service.
// All implementations must embed UnimplementedActionsServiceServer
// for forward compatibility
type ActionsServiceServer interface {
	// GetAction gets result of a given Action.
	GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error)
	// StartMySQLExplainAction starts MySQL EXPLAIN Action with traditional output.
	StartMySQLExplainAction(context.Context, *StartMySQLExplainActionRequest) (*StartMySQLExplainActionResponse, error)
	// StartMySQLExplainJSONAction starts MySQL EXPLAIN Action with JSON output.
	StartMySQLExplainJSONAction(context.Context, *StartMySQLExplainJSONActionRequest) (*StartMySQLExplainJSONActionResponse, error)
	// StartMySQLExplainTraditionalJSONAction starts MySQL EXPLAIN Action with traditional JSON output.
	StartMySQLExplainTraditionalJSONAction(context.Context, *StartMySQLExplainTraditionalJSONActionRequest) (*StartMySQLExplainTraditionalJSONActionResponse, error)
	// StartMySQLShowCreateTableAction starts MySQL SHOW CREATE TABLE Action.
	StartMySQLShowCreateTableAction(context.Context, *StartMySQLShowCreateTableActionRequest) (*StartMySQLShowCreateTableActionResponse, error)
	// StartMySQLShowTableStatusAction starts MySQL SHOW TABLE STATUS Action.
	StartMySQLShowTableStatusAction(context.Context, *StartMySQLShowTableStatusActionRequest) (*StartMySQLShowTableStatusActionResponse, error)
	// StartMySQLShowIndexAction starts MySQL SHOW INDEX Action.
	StartMySQLShowIndexAction(context.Context, *StartMySQLShowIndexActionRequest) (*StartMySQLShowIndexActionResponse, error)
	// StartPostgreSQLShowCreateTableAction starts PostgreSQL SHOW CREATE TABLE Action.
	StartPostgreSQLShowCreateTableAction(context.Context, *StartPostgreSQLShowCreateTableActionRequest) (*StartPostgreSQLShowCreateTableActionResponse, error)
	// StartPostgreSQLShowIndexAction starts PostgreSQL SHOW INDEX Action.
	StartPostgreSQLShowIndexAction(context.Context, *StartPostgreSQLShowIndexActionRequest) (*StartPostgreSQLShowIndexActionResponse, error)
	// StartMongoDBExplainAction starts MongoDB EXPLAIN Action.
	StartMongoDBExplainAction(context.Context, *StartMongoDBExplainActionRequest) (*StartMongoDBExplainActionResponse, error)
	// StartPTSummaryAction starts pt-summary Action.
	StartPTSummaryAction(context.Context, *StartPTSummaryActionRequest) (*StartPTSummaryActionResponse, error)
	// StartPTPgSummaryAction starts pt-pg-summary Action.
	StartPTPgSummaryAction(context.Context, *StartPTPgSummaryActionRequest) (*StartPTPgSummaryActionResponse, error)
	// StartPTMongoDBSummaryAction starts pt-mongodb-summary Action.
	StartPTMongoDBSummaryAction(context.Context, *StartPTMongoDBSummaryActionRequest) (*StartPTMongoDBSummaryActionResponse, error)
	// StartPTMySQLSummaryAction starts pt-mysql-summary Action.
	StartPTMySQLSummaryAction(context.Context, *StartPTMySQLSummaryActionRequest) (*StartPTMySQLSummaryActionResponse, error)
	// CancelAction stops an Action.
	CancelAction(context.Context, *CancelActionRequest) (*CancelActionResponse, error)
	mustEmbedUnimplementedActionsServiceServer()
}

// UnimplementedActionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActionsServiceServer struct{}

func (UnimplementedActionsServiceServer) GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAction not implemented")
}

func (UnimplementedActionsServiceServer) StartMySQLExplainAction(context.Context, *StartMySQLExplainActionRequest) (*StartMySQLExplainActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLExplainAction not implemented")
}

func (UnimplementedActionsServiceServer) StartMySQLExplainJSONAction(context.Context, *StartMySQLExplainJSONActionRequest) (*StartMySQLExplainJSONActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLExplainJSONAction not implemented")
}

func (UnimplementedActionsServiceServer) StartMySQLExplainTraditionalJSONAction(context.Context, *StartMySQLExplainTraditionalJSONActionRequest) (*StartMySQLExplainTraditionalJSONActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLExplainTraditionalJSONAction not implemented")
}

func (UnimplementedActionsServiceServer) StartMySQLShowCreateTableAction(context.Context, *StartMySQLShowCreateTableActionRequest) (*StartMySQLShowCreateTableActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLShowCreateTableAction not implemented")
}

func (UnimplementedActionsServiceServer) StartMySQLShowTableStatusAction(context.Context, *StartMySQLShowTableStatusActionRequest) (*StartMySQLShowTableStatusActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLShowTableStatusAction not implemented")
}

func (UnimplementedActionsServiceServer) StartMySQLShowIndexAction(context.Context, *StartMySQLShowIndexActionRequest) (*StartMySQLShowIndexActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLShowIndexAction not implemented")
}

func (UnimplementedActionsServiceServer) StartPostgreSQLShowCreateTableAction(context.Context, *StartPostgreSQLShowCreateTableActionRequest) (*StartPostgreSQLShowCreateTableActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPostgreSQLShowCreateTableAction not implemented")
}

func (UnimplementedActionsServiceServer) StartPostgreSQLShowIndexAction(context.Context, *StartPostgreSQLShowIndexActionRequest) (*StartPostgreSQLShowIndexActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPostgreSQLShowIndexAction not implemented")
}

func (UnimplementedActionsServiceServer) StartMongoDBExplainAction(context.Context, *StartMongoDBExplainActionRequest) (*StartMongoDBExplainActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMongoDBExplainAction not implemented")
}

func (UnimplementedActionsServiceServer) StartPTSummaryAction(context.Context, *StartPTSummaryActionRequest) (*StartPTSummaryActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPTSummaryAction not implemented")
}

func (UnimplementedActionsServiceServer) StartPTPgSummaryAction(context.Context, *StartPTPgSummaryActionRequest) (*StartPTPgSummaryActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPTPgSummaryAction not implemented")
}

func (UnimplementedActionsServiceServer) StartPTMongoDBSummaryAction(context.Context, *StartPTMongoDBSummaryActionRequest) (*StartPTMongoDBSummaryActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPTMongoDBSummaryAction not implemented")
}

func (UnimplementedActionsServiceServer) StartPTMySQLSummaryAction(context.Context, *StartPTMySQLSummaryActionRequest) (*StartPTMySQLSummaryActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPTMySQLSummaryAction not implemented")
}

func (UnimplementedActionsServiceServer) CancelAction(context.Context, *CancelActionRequest) (*CancelActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAction not implemented")
}
func (UnimplementedActionsServiceServer) mustEmbedUnimplementedActionsServiceServer() {}

// UnsafeActionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionsServiceServer will
// result in compilation errors.
type UnsafeActionsServiceServer interface {
	mustEmbedUnimplementedActionsServiceServer()
}

func RegisterActionsServiceServer(s grpc.ServiceRegistrar, srv ActionsServiceServer) {
	s.RegisterService(&ActionsService_ServiceDesc, srv)
}

func _ActionsService_GetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).GetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_GetAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).GetAction(ctx, req.(*GetActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartMySQLExplainAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLExplainActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartMySQLExplainAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartMySQLExplainAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartMySQLExplainAction(ctx, req.(*StartMySQLExplainActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartMySQLExplainJSONAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLExplainJSONActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartMySQLExplainJSONAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartMySQLExplainJSONAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartMySQLExplainJSONAction(ctx, req.(*StartMySQLExplainJSONActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartMySQLExplainTraditionalJSONAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLExplainTraditionalJSONActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartMySQLExplainTraditionalJSONAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartMySQLExplainTraditionalJSONAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartMySQLExplainTraditionalJSONAction(ctx, req.(*StartMySQLExplainTraditionalJSONActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartMySQLShowCreateTableAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLShowCreateTableActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartMySQLShowCreateTableAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartMySQLShowCreateTableAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartMySQLShowCreateTableAction(ctx, req.(*StartMySQLShowCreateTableActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartMySQLShowTableStatusAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLShowTableStatusActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartMySQLShowTableStatusAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartMySQLShowTableStatusAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartMySQLShowTableStatusAction(ctx, req.(*StartMySQLShowTableStatusActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartMySQLShowIndexAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLShowIndexActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartMySQLShowIndexAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartMySQLShowIndexAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartMySQLShowIndexAction(ctx, req.(*StartMySQLShowIndexActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartPostgreSQLShowCreateTableAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPostgreSQLShowCreateTableActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartPostgreSQLShowCreateTableAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartPostgreSQLShowCreateTableAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartPostgreSQLShowCreateTableAction(ctx, req.(*StartPostgreSQLShowCreateTableActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartPostgreSQLShowIndexAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPostgreSQLShowIndexActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartPostgreSQLShowIndexAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartPostgreSQLShowIndexAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartPostgreSQLShowIndexAction(ctx, req.(*StartPostgreSQLShowIndexActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartMongoDBExplainAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMongoDBExplainActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartMongoDBExplainAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartMongoDBExplainAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartMongoDBExplainAction(ctx, req.(*StartMongoDBExplainActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartPTSummaryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPTSummaryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartPTSummaryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartPTSummaryAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartPTSummaryAction(ctx, req.(*StartPTSummaryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartPTPgSummaryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPTPgSummaryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartPTPgSummaryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartPTPgSummaryAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartPTPgSummaryAction(ctx, req.(*StartPTPgSummaryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartPTMongoDBSummaryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPTMongoDBSummaryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartPTMongoDBSummaryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartPTMongoDBSummaryAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartPTMongoDBSummaryAction(ctx, req.(*StartPTMongoDBSummaryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartPTMySQLSummaryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPTMySQLSummaryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartPTMySQLSummaryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartPTMySQLSummaryAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartPTMySQLSummaryAction(ctx, req.(*StartPTMySQLSummaryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_CancelAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).CancelAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_CancelAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).CancelAction(ctx, req.(*CancelActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActionsService_ServiceDesc is the grpc.ServiceDesc for ActionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.v1.ActionsService",
	HandlerType: (*ActionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAction",
			Handler:    _ActionsService_GetAction_Handler,
		},
		{
			MethodName: "StartMySQLExplainAction",
			Handler:    _ActionsService_StartMySQLExplainAction_Handler,
		},
		{
			MethodName: "StartMySQLExplainJSONAction",
			Handler:    _ActionsService_StartMySQLExplainJSONAction_Handler,
		},
		{
			MethodName: "StartMySQLExplainTraditionalJSONAction",
			Handler:    _ActionsService_StartMySQLExplainTraditionalJSONAction_Handler,
		},
		{
			MethodName: "StartMySQLShowCreateTableAction",
			Handler:    _ActionsService_StartMySQLShowCreateTableAction_Handler,
		},
		{
			MethodName: "StartMySQLShowTableStatusAction",
			Handler:    _ActionsService_StartMySQLShowTableStatusAction_Handler,
		},
		{
			MethodName: "StartMySQLShowIndexAction",
			Handler:    _ActionsService_StartMySQLShowIndexAction_Handler,
		},
		{
			MethodName: "StartPostgreSQLShowCreateTableAction",
			Handler:    _ActionsService_StartPostgreSQLShowCreateTableAction_Handler,
		},
		{
			MethodName: "StartPostgreSQLShowIndexAction",
			Handler:    _ActionsService_StartPostgreSQLShowIndexAction_Handler,
		},
		{
			MethodName: "StartMongoDBExplainAction",
			Handler:    _ActionsService_StartMongoDBExplainAction_Handler,
		},
		{
			MethodName: "StartPTSummaryAction",
			Handler:    _ActionsService_StartPTSummaryAction_Handler,
		},
		{
			MethodName: "StartPTPgSummaryAction",
			Handler:    _ActionsService_StartPTPgSummaryAction_Handler,
		},
		{
			MethodName: "StartPTMongoDBSummaryAction",
			Handler:    _ActionsService_StartPTMongoDBSummaryAction_Handler,
		},
		{
			MethodName: "StartPTMySQLSummaryAction",
			Handler:    _ActionsService_StartPTMySQLSummaryAction_Handler,
		},
		{
			MethodName: "CancelAction",
			Handler:    _ActionsService_CancelAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "management/v1/actions.proto",
}
