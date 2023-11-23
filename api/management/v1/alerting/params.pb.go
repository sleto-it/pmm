// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        (unknown)
// source: management/v1/alerting/params.proto

package alertingv1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ParamUnit represents template parameter unit.
type ParamUnit int32

const (
	// Invalid, unknown or absent.
	ParamUnit_PARAM_UNIT_UNSPECIFIED ParamUnit = 0
	// %
	ParamUnit_PARAM_UNIT_PERCENTAGE ParamUnit = 1
	// s
	ParamUnit_PARAM_UNIT_SECONDS ParamUnit = 2
)

// Enum value maps for ParamUnit.
var (
	ParamUnit_name = map[int32]string{
		0: "PARAM_UNIT_UNSPECIFIED",
		1: "PARAM_UNIT_PERCENTAGE",
		2: "PARAM_UNIT_SECONDS",
	}
	ParamUnit_value = map[string]int32{
		"PARAM_UNIT_UNSPECIFIED": 0,
		"PARAM_UNIT_PERCENTAGE":  1,
		"PARAM_UNIT_SECONDS":     2,
	}
)

func (x ParamUnit) Enum() *ParamUnit {
	p := new(ParamUnit)
	*p = x
	return p
}

func (x ParamUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParamUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_management_v1_alerting_params_proto_enumTypes[0].Descriptor()
}

func (ParamUnit) Type() protoreflect.EnumType {
	return &file_management_v1_alerting_params_proto_enumTypes[0]
}

func (x ParamUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParamUnit.Descriptor instead.
func (ParamUnit) EnumDescriptor() ([]byte, []int) {
	return file_management_v1_alerting_params_proto_rawDescGZIP(), []int{0}
}

// ParamType represents template parameter type.
type ParamType int32

const (
	ParamType_PARAM_TYPE_UNSPECIFIED ParamType = 0
	ParamType_PARAM_TYPE_BOOL        ParamType = 1
	ParamType_PARAM_TYPE_FLOAT       ParamType = 2
	ParamType_PARAM_TYPE_STRING      ParamType = 3
)

// Enum value maps for ParamType.
var (
	ParamType_name = map[int32]string{
		0: "PARAM_TYPE_UNSPECIFIED",
		1: "PARAM_TYPE_BOOL",
		2: "PARAM_TYPE_FLOAT",
		3: "PARAM_TYPE_STRING",
	}
	ParamType_value = map[string]int32{
		"PARAM_TYPE_UNSPECIFIED": 0,
		"PARAM_TYPE_BOOL":        1,
		"PARAM_TYPE_FLOAT":       2,
		"PARAM_TYPE_STRING":      3,
	}
)

func (x ParamType) Enum() *ParamType {
	p := new(ParamType)
	*p = x
	return p
}

func (x ParamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParamType) Descriptor() protoreflect.EnumDescriptor {
	return file_management_v1_alerting_params_proto_enumTypes[1].Descriptor()
}

func (ParamType) Type() protoreflect.EnumType {
	return &file_management_v1_alerting_params_proto_enumTypes[1]
}

func (x ParamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParamType.Descriptor instead.
func (ParamType) EnumDescriptor() ([]byte, []int) {
	return file_management_v1_alerting_params_proto_rawDescGZIP(), []int{1}
}

var File_management_v1_alerting_params_proto protoreflect.FileDescriptor

var file_management_v1_alerting_params_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2a, 0x5a, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x12,
	0x1a, 0x0a, 0x16, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e,
	0x54, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f,
	0x55, 0x4e, 0x49, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x53, 0x10, 0x02, 0x2a, 0x69,
	0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54,
	0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x42, 0xa9, 0x01, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61,
	0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x3b,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58,
	0xaa, 0x02, 0x0b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_v1_alerting_params_proto_rawDescOnce sync.Once
	file_management_v1_alerting_params_proto_rawDescData = file_management_v1_alerting_params_proto_rawDesc
)

func file_management_v1_alerting_params_proto_rawDescGZIP() []byte {
	file_management_v1_alerting_params_proto_rawDescOnce.Do(func() {
		file_management_v1_alerting_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_v1_alerting_params_proto_rawDescData)
	})
	return file_management_v1_alerting_params_proto_rawDescData
}

var (
	file_management_v1_alerting_params_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
	file_management_v1_alerting_params_proto_goTypes   = []interface{}{
		(ParamUnit)(0), // 0: alerting.v1.ParamUnit
		(ParamType)(0), // 1: alerting.v1.ParamType
	}
)

var file_management_v1_alerting_params_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_management_v1_alerting_params_proto_init() }
func file_management_v1_alerting_params_proto_init() {
	if File_management_v1_alerting_params_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_management_v1_alerting_params_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_management_v1_alerting_params_proto_goTypes,
		DependencyIndexes: file_management_v1_alerting_params_proto_depIdxs,
		EnumInfos:         file_management_v1_alerting_params_proto_enumTypes,
	}.Build()
	File_management_v1_alerting_params_proto = out.File
	file_management_v1_alerting_params_proto_rawDesc = nil
	file_management_v1_alerting_params_proto_goTypes = nil
	file_management_v1_alerting_params_proto_depIdxs = nil
}
