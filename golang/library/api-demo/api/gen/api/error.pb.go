// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: error.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 错误状态码
type ECode int32

const (
	// 通用 0 - 99
	ECode_EOK      ECode = 0 // 成功
	ECode_EUnknown ECode = 1 // 未知错误
	ECode_EService ECode = 2 // 服务错误
	ECode_EParam   ECode = 3 // 参数错误
	// 用户  100 - 199
	ECode_EAuthNotLogin  ECode = 100 // 未授权登录
	ECode_EAuthExpire    ECode = 101 // 授权过期
	ECode_EUserExist     ECode = 102 // 用户已存在
	ECode_EUserNotExist  ECode = 103 // 用户不存在
	ECode_EEmailExist    ECode = 104 // 邮箱已存在
	ECode_EEmailNotExist ECode = 105 // 邮箱不存在
	ECode_EPasswordError ECode = 106 // 密码错误
)

// Enum value maps for ECode.
var (
	ECode_name = map[int32]string{
		0:   "EOK",
		1:   "EUnknown",
		2:   "EService",
		3:   "EParam",
		100: "EAuthNotLogin",
		101: "EAuthExpire",
		102: "EUserExist",
		103: "EUserNotExist",
		104: "EEmailExist",
		105: "EEmailNotExist",
		106: "EPasswordError",
	}
	ECode_value = map[string]int32{
		"EOK":            0,
		"EUnknown":       1,
		"EService":       2,
		"EParam":         3,
		"EAuthNotLogin":  100,
		"EAuthExpire":    101,
		"EUserExist":     102,
		"EUserNotExist":  103,
		"EEmailExist":    104,
		"EEmailNotExist": 105,
		"EPasswordError": 106,
	}
)

func (x ECode) Enum() *ECode {
	p := new(ECode)
	*p = x
	return p
}

func (x ECode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ECode) Descriptor() protoreflect.EnumDescriptor {
	return file_error_proto_enumTypes[0].Descriptor()
}

func (ECode) Type() protoreflect.EnumType {
	return &file_error_proto_enumTypes[0]
}

func (x ECode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ECode.Descriptor instead.
func (ECode) EnumDescriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

var File_error_proto protoreflect.FileDescriptor

var file_error_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x2a, 0xb8, 0x01, 0x0a, 0x05, 0x45, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x45, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x10, 0x03, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x41, 0x75, 0x74, 0x68, 0x4e, 0x6f, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x64,
	0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x41, 0x75, 0x74, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x10,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10,
	0x66, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x10, 0x67, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x10, 0x68, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e,
	0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0x69, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x6a, 0x42, 0x11, 0x5a,
	0x0f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_proto_rawDescOnce sync.Once
	file_error_proto_rawDescData = file_error_proto_rawDesc
)

func file_error_proto_rawDescGZIP() []byte {
	file_error_proto_rawDescOnce.Do(func() {
		file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_proto_rawDescData)
	})
	return file_error_proto_rawDescData
}

var file_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_proto_goTypes = []interface{}{
	(ECode)(0), // 0: api.ECode
}
var file_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_proto_init() }
func file_error_proto_init() {
	if File_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_proto_goTypes,
		DependencyIndexes: file_error_proto_depIdxs,
		EnumInfos:         file_error_proto_enumTypes,
	}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}
