// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pbf/role/api.proto

package role

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_pbf_role_api_proto protoreflect.FileDescriptor

var file_pbf_role_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x66, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x15, 0x70, 0x62, 0x66, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x70, 0x62, 0x66, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x62, 0x66, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x70, 0x62, 0x66, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xad, 0x01, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12, 0x28,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x22, 0x00, 0x12, 0x28, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0d, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x1a, 0x0d, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x72, 0x6f, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pbf_role_api_proto_goTypes = []interface{}{
	(*CreateI)(nil), // 0: role.CreateI
	(*DeleteI)(nil), // 1: role.DeleteI
	(*SearchI)(nil), // 2: role.SearchI
	(*UpdateI)(nil), // 3: role.UpdateI
	(*CreateO)(nil), // 4: role.CreateO
	(*DeleteO)(nil), // 5: role.DeleteO
	(*SearchO)(nil), // 6: role.SearchO
	(*UpdateO)(nil), // 7: role.UpdateO
}
var file_pbf_role_api_proto_depIdxs = []int32{
	0, // 0: role.API.Create:input_type -> role.CreateI
	1, // 1: role.API.Delete:input_type -> role.DeleteI
	2, // 2: role.API.Search:input_type -> role.SearchI
	3, // 3: role.API.Update:input_type -> role.UpdateI
	4, // 4: role.API.Create:output_type -> role.CreateO
	5, // 5: role.API.Delete:output_type -> role.DeleteO
	6, // 6: role.API.Search:output_type -> role.SearchO
	7, // 7: role.API.Update:output_type -> role.UpdateO
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbf_role_api_proto_init() }
func file_pbf_role_api_proto_init() {
	if File_pbf_role_api_proto != nil {
		return
	}
	file_pbf_role_create_proto_init()
	file_pbf_role_delete_proto_init()
	file_pbf_role_search_proto_init()
	file_pbf_role_update_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbf_role_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pbf_role_api_proto_goTypes,
		DependencyIndexes: file_pbf_role_api_proto_depIdxs,
	}.Build()
	File_pbf_role_api_proto = out.File
	file_pbf_role_api_proto_rawDesc = nil
	file_pbf_role_api_proto_goTypes = nil
	file_pbf_role_api_proto_depIdxs = nil
}
