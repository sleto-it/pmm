// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: management/v1/azure.proto

package managementv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DiscoverAzureDatabaseType describes supported Azure Database instance engines.
type DiscoverAzureDatabaseType int32

const (
	DiscoverAzureDatabaseType_DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED DiscoverAzureDatabaseType = 0
	// MySQL type: microsoft.dbformysql or MariaDB type: microsoft.dbformariadb
	DiscoverAzureDatabaseType_DISCOVER_AZURE_DATABASE_TYPE_MYSQL DiscoverAzureDatabaseType = 1
	// PostgreSQL type: microsoft.dbformysql
	DiscoverAzureDatabaseType_DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL DiscoverAzureDatabaseType = 2
)

// Enum value maps for DiscoverAzureDatabaseType.
var (
	DiscoverAzureDatabaseType_name = map[int32]string{
		0: "DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED",
		1: "DISCOVER_AZURE_DATABASE_TYPE_MYSQL",
		2: "DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL",
	}
	DiscoverAzureDatabaseType_value = map[string]int32{
		"DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED": 0,
		"DISCOVER_AZURE_DATABASE_TYPE_MYSQL":       1,
		"DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL":  2,
	}
)

func (x DiscoverAzureDatabaseType) Enum() *DiscoverAzureDatabaseType {
	p := new(DiscoverAzureDatabaseType)
	*p = x
	return p
}

func (x DiscoverAzureDatabaseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiscoverAzureDatabaseType) Descriptor() protoreflect.EnumDescriptor {
	return file_management_v1_azure_proto_enumTypes[0].Descriptor()
}

func (DiscoverAzureDatabaseType) Type() protoreflect.EnumType {
	return &file_management_v1_azure_proto_enumTypes[0]
}

func (x DiscoverAzureDatabaseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiscoverAzureDatabaseType.Descriptor instead.
func (DiscoverAzureDatabaseType) EnumDescriptor() ([]byte, []int) {
	return file_management_v1_azure_proto_rawDescGZIP(), []int{0}
}

// DiscoverAzureDatabaseRequest discover azure databases request.
type DiscoverAzureDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Azure client ID.
	AzureClientId string `protobuf:"bytes,1,opt,name=azure_client_id,json=azureClientId,proto3" json:"azure_client_id,omitempty"`
	// Azure client secret.
	AzureClientSecret string `protobuf:"bytes,2,opt,name=azure_client_secret,json=azureClientSecret,proto3" json:"azure_client_secret,omitempty"`
	// Azure tanant ID.
	AzureTenantId string `protobuf:"bytes,3,opt,name=azure_tenant_id,json=azureTenantId,proto3" json:"azure_tenant_id,omitempty"`
	// Azure subscription ID.
	AzureSubscriptionId string `protobuf:"bytes,4,opt,name=azure_subscription_id,json=azureSubscriptionId,proto3" json:"azure_subscription_id,omitempty"`
}

func (x *DiscoverAzureDatabaseRequest) Reset() {
	*x = DiscoverAzureDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_v1_azure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverAzureDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverAzureDatabaseRequest) ProtoMessage() {}

func (x *DiscoverAzureDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_azure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverAzureDatabaseRequest.ProtoReflect.Descriptor instead.
func (*DiscoverAzureDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_management_v1_azure_proto_rawDescGZIP(), []int{0}
}

func (x *DiscoverAzureDatabaseRequest) GetAzureClientId() string {
	if x != nil {
		return x.AzureClientId
	}
	return ""
}

func (x *DiscoverAzureDatabaseRequest) GetAzureClientSecret() string {
	if x != nil {
		return x.AzureClientSecret
	}
	return ""
}

func (x *DiscoverAzureDatabaseRequest) GetAzureTenantId() string {
	if x != nil {
		return x.AzureTenantId
	}
	return ""
}

func (x *DiscoverAzureDatabaseRequest) GetAzureSubscriptionId() string {
	if x != nil {
		return x.AzureSubscriptionId
	}
	return ""
}

// DiscoverAzureDatabaseInstance models an unique Azure Database instance for the list of instances returned by Discovery.
type DiscoverAzureDatabaseInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Azure database instance ID.
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Azure database location.
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// Azure database server name.
	ServiceName string `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Database username.
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	// Address used to connect to it.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Azure Resource group.
	AzureResourceGroup string `protobuf:"bytes,6,opt,name=azure_resource_group,json=azureResourceGroup,proto3" json:"azure_resource_group,omitempty"`
	// Environment tag.
	Environment string `protobuf:"bytes,7,opt,name=environment,proto3" json:"environment,omitempty"`
	// Database type.
	Type DiscoverAzureDatabaseType `protobuf:"varint,8,opt,name=type,proto3,enum=management.v1.DiscoverAzureDatabaseType" json:"type,omitempty"`
	// Azure database availability zone.
	Az string `protobuf:"bytes,9,opt,name=az,proto3" json:"az,omitempty"`
	// Represents a purchasable Stock Keeping Unit (SKU) under a product.
	// https://docs.microsoft.com/en-us/partner-center/develop/product-resources#sku.
	NodeModel string `protobuf:"bytes,10,opt,name=node_model,json=nodeModel,proto3" json:"node_model,omitempty"`
}

func (x *DiscoverAzureDatabaseInstance) Reset() {
	*x = DiscoverAzureDatabaseInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_v1_azure_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverAzureDatabaseInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverAzureDatabaseInstance) ProtoMessage() {}

func (x *DiscoverAzureDatabaseInstance) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_azure_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverAzureDatabaseInstance.ProtoReflect.Descriptor instead.
func (*DiscoverAzureDatabaseInstance) Descriptor() ([]byte, []int) {
	return file_management_v1_azure_proto_rawDescGZIP(), []int{1}
}

func (x *DiscoverAzureDatabaseInstance) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *DiscoverAzureDatabaseInstance) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *DiscoverAzureDatabaseInstance) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *DiscoverAzureDatabaseInstance) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DiscoverAzureDatabaseInstance) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DiscoverAzureDatabaseInstance) GetAzureResourceGroup() string {
	if x != nil {
		return x.AzureResourceGroup
	}
	return ""
}

func (x *DiscoverAzureDatabaseInstance) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *DiscoverAzureDatabaseInstance) GetType() DiscoverAzureDatabaseType {
	if x != nil {
		return x.Type
	}
	return DiscoverAzureDatabaseType_DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED
}

func (x *DiscoverAzureDatabaseInstance) GetAz() string {
	if x != nil {
		return x.Az
	}
	return ""
}

func (x *DiscoverAzureDatabaseInstance) GetNodeModel() string {
	if x != nil {
		return x.NodeModel
	}
	return ""
}

// DiscoverAzureDatabaseResponse discover azure databases response.
type DiscoverAzureDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AzureDatabaseInstance []*DiscoverAzureDatabaseInstance `protobuf:"bytes,1,rep,name=azure_database_instance,json=azureDatabaseInstance,proto3" json:"azure_database_instance,omitempty"`
}

func (x *DiscoverAzureDatabaseResponse) Reset() {
	*x = DiscoverAzureDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_v1_azure_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverAzureDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverAzureDatabaseResponse) ProtoMessage() {}

func (x *DiscoverAzureDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_azure_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverAzureDatabaseResponse.ProtoReflect.Descriptor instead.
func (*DiscoverAzureDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_management_v1_azure_proto_rawDescGZIP(), []int{2}
}

func (x *DiscoverAzureDatabaseResponse) GetAzureDatabaseInstance() []*DiscoverAzureDatabaseInstance {
	if x != nil {
		return x.AzureDatabaseInstance
	}
	return nil
}

type AddAzureDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Azure database location.
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Azure database availability zone.
	Az string `protobuf:"bytes,2,opt,name=az,proto3" json:"az,omitempty"`
	// Azure database instance ID.
	InstanceId string `protobuf:"bytes,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Represents a purchasable Stock Keeping Unit (SKU) under a product.
	// https://docs.microsoft.com/en-us/partner-center/develop/product-resources#sku.
	NodeModel string `protobuf:"bytes,4,opt,name=node_model,json=nodeModel,proto3" json:"node_model,omitempty"`
	// Address used to connect to it.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Access port.
	Port uint32 `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	// Unique across all Nodes user-defined name. Defaults to Azure Database instance ID.
	NodeName string `protobuf:"bytes,7,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Unique across all Services user-defined name. Defaults to Azure Database instance ID.
	ServiceName string `protobuf:"bytes,8,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,9,opt,name=environment,proto3" json:"environment,omitempty"`
	// Username for scraping metrics.
	Username string `protobuf:"bytes,10,opt,name=username,proto3" json:"username,omitempty"`
	// Password for scraping metrics.
	Password string `protobuf:"bytes,11,opt,name=password,proto3" json:"password,omitempty"`
	// Azure client ID.
	AzureClientId string `protobuf:"bytes,12,opt,name=azure_client_id,json=azureClientId,proto3" json:"azure_client_id,omitempty"`
	// Azure client secret.
	AzureClientSecret string `protobuf:"bytes,13,opt,name=azure_client_secret,json=azureClientSecret,proto3" json:"azure_client_secret,omitempty"`
	// Azure tanant ID.
	AzureTenantId string `protobuf:"bytes,14,opt,name=azure_tenant_id,json=azureTenantId,proto3" json:"azure_tenant_id,omitempty"`
	// Azure subscription ID.
	AzureSubscriptionId string `protobuf:"bytes,15,opt,name=azure_subscription_id,json=azureSubscriptionId,proto3" json:"azure_subscription_id,omitempty"`
	// Azure resource group.
	AzureResourceGroup string `protobuf:"bytes,16,opt,name=azure_resource_group,json=azureResourceGroup,proto3" json:"azure_resource_group,omitempty"`
	// If true, adds azure_database_exporter.
	AzureDatabaseExporter bool `protobuf:"varint,17,opt,name=azure_database_exporter,json=azureDatabaseExporter,proto3" json:"azure_database_exporter,omitempty"`
	// If true, adds qan-mysql-perfschema-agent or qan-postgresql-pgstatements-agent.
	Qan bool `protobuf:"varint,18,opt,name=qan,proto3" json:"qan,omitempty"`
	// Custom user-assigned labels for Node and Service.
	CustomLabels map[string]string `protobuf:"bytes,19,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Skip connection check.
	SkipConnectionCheck bool `protobuf:"varint,20,opt,name=skip_connection_check,json=skipConnectionCheck,proto3" json:"skip_connection_check,omitempty"`
	// Use TLS for database connections.
	Tls bool `protobuf:"varint,21,opt,name=tls,proto3" json:"tls,omitempty"`
	// Skip TLS certificate and hostname validation.
	TlsSkipVerify bool `protobuf:"varint,22,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty"`
	// Disable query examples.
	DisableQueryExamples bool `protobuf:"varint,23,opt,name=disable_query_examples,json=disableQueryExamples,proto3" json:"disable_query_examples,omitempty"`
	// Tablestats group collectors will be disabled if there are more than that number of tables.
	// If zero, server's default value is used.
	// Use negative value to disable them.
	TablestatsGroupTableLimit int32 `protobuf:"varint,24,opt,name=tablestats_group_table_limit,json=tablestatsGroupTableLimit,proto3" json:"tablestats_group_table_limit,omitempty"`
	// Azure database resource type (mysql, maria, postgres)
	Type DiscoverAzureDatabaseType `protobuf:"varint,25,opt,name=type,proto3,enum=management.v1.DiscoverAzureDatabaseType" json:"type,omitempty"`
}

func (x *AddAzureDatabaseRequest) Reset() {
	*x = AddAzureDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_v1_azure_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAzureDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAzureDatabaseRequest) ProtoMessage() {}

func (x *AddAzureDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_azure_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAzureDatabaseRequest.ProtoReflect.Descriptor instead.
func (*AddAzureDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_management_v1_azure_proto_rawDescGZIP(), []int{3}
}

func (x *AddAzureDatabaseRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAz() string {
	if x != nil {
		return x.Az
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetNodeModel() string {
	if x != nil {
		return x.NodeModel
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *AddAzureDatabaseRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureClientId() string {
	if x != nil {
		return x.AzureClientId
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureClientSecret() string {
	if x != nil {
		return x.AzureClientSecret
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureTenantId() string {
	if x != nil {
		return x.AzureTenantId
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureSubscriptionId() string {
	if x != nil {
		return x.AzureSubscriptionId
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureResourceGroup() string {
	if x != nil {
		return x.AzureResourceGroup
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureDatabaseExporter() bool {
	if x != nil {
		return x.AzureDatabaseExporter
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetQan() bool {
	if x != nil {
		return x.Qan
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

func (x *AddAzureDatabaseRequest) GetSkipConnectionCheck() bool {
	if x != nil {
		return x.SkipConnectionCheck
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetTls() bool {
	if x != nil {
		return x.Tls
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetTlsSkipVerify() bool {
	if x != nil {
		return x.TlsSkipVerify
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetDisableQueryExamples() bool {
	if x != nil {
		return x.DisableQueryExamples
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetTablestatsGroupTableLimit() int32 {
	if x != nil {
		return x.TablestatsGroupTableLimit
	}
	return 0
}

func (x *AddAzureDatabaseRequest) GetType() DiscoverAzureDatabaseType {
	if x != nil {
		return x.Type
	}
	return DiscoverAzureDatabaseType_DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED
}

type AddAzureDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddAzureDatabaseResponse) Reset() {
	*x = AddAzureDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_v1_azure_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAzureDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAzureDatabaseResponse) ProtoMessage() {}

func (x *AddAzureDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_azure_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAzureDatabaseResponse.ProtoReflect.Descriptor instead.
func (*AddAzureDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_management_v1_azure_proto_rawDescGZIP(), []int{4}
}

var File_management_v1_azure_proto protoreflect.FileDescriptor

var file_management_v1_azure_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x1c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x41, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x13, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x11, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2f,
	0x0a, 0x0f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x0d, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x3b, 0x0a, 0x15, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x13, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xf2, 0x02, 0x0a,
	0x1d, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x30, 0x0a, 0x14, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x7a, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x61, 0x7a, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x22, 0x85, 0x01, 0x0a, 0x1d, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x7a,
	0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x17, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x15, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x96, 0x09, 0x0a, 0x17, 0x41, 0x64,
	0x64, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x7a, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x61, 0x7a, 0x12, 0x28, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x21, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x2f, 0x0a, 0x0f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x13, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x11, 0x61, 0x7a, 0x75, 0x72, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x0f,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0d,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3b, 0x0a,
	0x15, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x13, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x14, 0x61, 0x7a,
	0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x12, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x36, 0x0a, 0x17, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x71, 0x61, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x71, 0x61, 0x6e, 0x12,
	0x5d, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x32,
	0x0a, 0x15, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73,
	0x6b, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x74, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74,
	0x6c, 0x73, 0x53, 0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x34, 0x0a, 0x16,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x12, 0x3f, 0x0a, 0x1c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x28, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x1a, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x9e,
	0x01, 0x0a, 0x19, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x28,
	0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x41, 0x5a, 0x55, 0x52, 0x45, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x44, 0x49,
	0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x41, 0x5a, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x41, 0x54,
	0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x59, 0x53, 0x51, 0x4c,
	0x10, 0x01, 0x12, 0x2b, 0x0a, 0x27, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x41,
	0x5a, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x51, 0x4c, 0x10, 0x02, 0x42,
	0xab, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58,
	0xaa, 0x02, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x19, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_v1_azure_proto_rawDescOnce sync.Once
	file_management_v1_azure_proto_rawDescData = file_management_v1_azure_proto_rawDesc
)

func file_management_v1_azure_proto_rawDescGZIP() []byte {
	file_management_v1_azure_proto_rawDescOnce.Do(func() {
		file_management_v1_azure_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_v1_azure_proto_rawDescData)
	})
	return file_management_v1_azure_proto_rawDescData
}

var (
	file_management_v1_azure_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_management_v1_azure_proto_msgTypes  = make([]protoimpl.MessageInfo, 6)
	file_management_v1_azure_proto_goTypes   = []any{
		(DiscoverAzureDatabaseType)(0),        // 0: management.v1.DiscoverAzureDatabaseType
		(*DiscoverAzureDatabaseRequest)(nil),  // 1: management.v1.DiscoverAzureDatabaseRequest
		(*DiscoverAzureDatabaseInstance)(nil), // 2: management.v1.DiscoverAzureDatabaseInstance
		(*DiscoverAzureDatabaseResponse)(nil), // 3: management.v1.DiscoverAzureDatabaseResponse
		(*AddAzureDatabaseRequest)(nil),       // 4: management.v1.AddAzureDatabaseRequest
		(*AddAzureDatabaseResponse)(nil),      // 5: management.v1.AddAzureDatabaseResponse
		nil,                                   // 6: management.v1.AddAzureDatabaseRequest.CustomLabelsEntry
	}
)

var file_management_v1_azure_proto_depIdxs = []int32{
	0, // 0: management.v1.DiscoverAzureDatabaseInstance.type:type_name -> management.v1.DiscoverAzureDatabaseType
	2, // 1: management.v1.DiscoverAzureDatabaseResponse.azure_database_instance:type_name -> management.v1.DiscoverAzureDatabaseInstance
	6, // 2: management.v1.AddAzureDatabaseRequest.custom_labels:type_name -> management.v1.AddAzureDatabaseRequest.CustomLabelsEntry
	0, // 3: management.v1.AddAzureDatabaseRequest.type:type_name -> management.v1.DiscoverAzureDatabaseType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_management_v1_azure_proto_init() }
func file_management_v1_azure_proto_init() {
	if File_management_v1_azure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_management_v1_azure_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DiscoverAzureDatabaseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_management_v1_azure_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DiscoverAzureDatabaseInstance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_management_v1_azure_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DiscoverAzureDatabaseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_management_v1_azure_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddAzureDatabaseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_management_v1_azure_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddAzureDatabaseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_management_v1_azure_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_management_v1_azure_proto_goTypes,
		DependencyIndexes: file_management_v1_azure_proto_depIdxs,
		EnumInfos:         file_management_v1_azure_proto_enumTypes,
		MessageInfos:      file_management_v1_azure_proto_msgTypes,
	}.Build()
	File_management_v1_azure_proto = out.File
	file_management_v1_azure_proto_rawDesc = nil
	file_management_v1_azure_proto_goTypes = nil
	file_management_v1_azure_proto_depIdxs = nil
}
