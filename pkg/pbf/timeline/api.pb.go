// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pbf/timeline/api.proto

package timeline

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

var File_pbf_timeline_api_proto protoreflect.FileDescriptor

var file_pbf_timeline_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x1a, 0x19, 0x70, 0x62, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70,
	0x62, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x62, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x62, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcd,
	0x01, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x1a, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x1a, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x1a, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x1a, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x3b, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_pbf_timeline_api_proto_goTypes = []interface{}{
	(*CreateI)(nil), // 0: timeline.CreateI
	(*DeleteI)(nil), // 1: timeline.DeleteI
	(*SearchI)(nil), // 2: timeline.SearchI
	(*UpdateI)(nil), // 3: timeline.UpdateI
	(*CreateO)(nil), // 4: timeline.CreateO
	(*DeleteO)(nil), // 5: timeline.DeleteO
	(*SearchO)(nil), // 6: timeline.SearchO
	(*UpdateO)(nil), // 7: timeline.UpdateO
}
var file_pbf_timeline_api_proto_depIdxs = []int32{
	0, // 0: timeline.API.Create:input_type -> timeline.CreateI
	1, // 1: timeline.API.Delete:input_type -> timeline.DeleteI
	2, // 2: timeline.API.Search:input_type -> timeline.SearchI
	3, // 3: timeline.API.Update:input_type -> timeline.UpdateI
	4, // 4: timeline.API.Create:output_type -> timeline.CreateO
	5, // 5: timeline.API.Delete:output_type -> timeline.DeleteO
	6, // 6: timeline.API.Search:output_type -> timeline.SearchO
	7, // 7: timeline.API.Update:output_type -> timeline.UpdateO
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbf_timeline_api_proto_init() }
func file_pbf_timeline_api_proto_init() {
	if File_pbf_timeline_api_proto != nil {
		return
	}
	file_pbf_timeline_create_proto_init()
	file_pbf_timeline_delete_proto_init()
	file_pbf_timeline_search_proto_init()
	file_pbf_timeline_update_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbf_timeline_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pbf_timeline_api_proto_goTypes,
		DependencyIndexes: file_pbf_timeline_api_proto_depIdxs,
	}.Build()
	File_pbf_timeline_api_proto = out.File
	file_pbf_timeline_api_proto_rawDesc = nil
	file_pbf_timeline_api_proto_goTypes = nil
	file_pbf_timeline_api_proto_depIdxs = nil
}