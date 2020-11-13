// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pbf/metric/delete.proto

package metric

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DeleteI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteI) Reset() {
	*x = DeleteI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_metric_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI) ProtoMessage() {}

func (x *DeleteI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metric_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI.ProtoReflect.Descriptor instead.
func (*DeleteI) Descriptor() ([]byte, []int) {
	return file_pbf_metric_delete_proto_rawDescGZIP(), []int{0}
}

type DeleteO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteO) Reset() {
	*x = DeleteO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_metric_delete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO) ProtoMessage() {}

func (x *DeleteO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metric_delete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO.ProtoReflect.Descriptor instead.
func (*DeleteO) Descriptor() ([]byte, []int) {
	return file_pbf_metric_delete_proto_rawDescGZIP(), []int{1}
}

var File_pbf_metric_delete_proto protoreflect.FileDescriptor

var file_pbf_metric_delete_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x22, 0x09, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x22, 0x09, 0x0a, 0x07,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_metric_delete_proto_rawDescOnce sync.Once
	file_pbf_metric_delete_proto_rawDescData = file_pbf_metric_delete_proto_rawDesc
)

func file_pbf_metric_delete_proto_rawDescGZIP() []byte {
	file_pbf_metric_delete_proto_rawDescOnce.Do(func() {
		file_pbf_metric_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_metric_delete_proto_rawDescData)
	})
	return file_pbf_metric_delete_proto_rawDescData
}

var file_pbf_metric_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbf_metric_delete_proto_goTypes = []interface{}{
	(*DeleteI)(nil), // 0: metric.DeleteI
	(*DeleteO)(nil), // 1: metric.DeleteO
}
var file_pbf_metric_delete_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbf_metric_delete_proto_init() }
func file_pbf_metric_delete_proto_init() {
	if File_pbf_metric_delete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_metric_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI); i {
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
		file_pbf_metric_delete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO); i {
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
			RawDescriptor: file_pbf_metric_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_metric_delete_proto_goTypes,
		DependencyIndexes: file_pbf_metric_delete_proto_depIdxs,
		MessageInfos:      file_pbf_metric_delete_proto_msgTypes,
	}.Build()
	File_pbf_metric_delete_proto = out.File
	file_pbf_metric_delete_proto_rawDesc = nil
	file_pbf_metric_delete_proto_goTypes = nil
	file_pbf_metric_delete_proto_depIdxs = nil
}
