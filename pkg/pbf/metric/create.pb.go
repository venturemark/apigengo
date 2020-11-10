// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pbf/metric/create.proto

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

// CreateI is the input to create metrics. That is persisting datapoints
// associated with an update. Below is an example JSON representation showing an
// emitted datapoint for the dimensions x and y, where x represents a Unix
// Timestamp and y represents a numeric value given by the user at that time.
//
//     {
//         "datapoint": [
//             1604959525,
//             85
//         ],
//         "metric_id": "met-d289g",
//         "timestamp": "1604959525",
//         "update_id": "upd-j3o45"
//     }
//
type CreateI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// datapoint is a list of datapoints provided by convention. Providing [x, y]
	// may be interpreted as datapoints for x and y dimensions on a graph.
	Datapoint []int64 `protobuf:"varint,1,rep,packed,name=datapoint,proto3" json:"datapoint,omitempty"`
	// metric_id is the identifier used to associate the given datapoints with the
	// specified metric.
	MetricId string `protobuf:"bytes,2,opt,name=metric_id,json=metricId,proto3" json:"metric_id,omitempty"`
	// timestamp is the time at which the datapoints got emitted. With conventions
	// of the x axis to be time, the first datapoint and timestamp must be equal.
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// update_id is the identifier used to associate the created metric with the
	// specified update.
	UpdateId string `protobuf:"bytes,4,opt,name=update_id,json=updateId,proto3" json:"update_id,omitempty"`
}

func (x *CreateI) Reset() {
	*x = CreateI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_metric_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI) ProtoMessage() {}

func (x *CreateI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metric_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateI.ProtoReflect.Descriptor instead.
func (*CreateI) Descriptor() ([]byte, []int) {
	return file_pbf_metric_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateI) GetDatapoint() []int64 {
	if x != nil {
		return x.Datapoint
	}
	return nil
}

func (x *CreateI) GetMetricId() string {
	if x != nil {
		return x.MetricId
	}
	return ""
}

func (x *CreateI) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *CreateI) GetUpdateId() string {
	if x != nil {
		return x.UpdateId
	}
	return ""
}

type CreateO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateO) Reset() {
	*x = CreateO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_metric_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO) ProtoMessage() {}

func (x *CreateO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metric_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateO.ProtoReflect.Descriptor instead.
func (*CreateO) Descriptor() ([]byte, []int) {
	return file_pbf_metric_create_proto_rawDescGZIP(), []int{1}
}

var File_pbf_metric_create_proto protoreflect.FileDescriptor

var file_pbf_metric_create_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x22, 0x7f, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x22, 0x09, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x3b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pbf_metric_create_proto_rawDescOnce sync.Once
	file_pbf_metric_create_proto_rawDescData = file_pbf_metric_create_proto_rawDesc
)

func file_pbf_metric_create_proto_rawDescGZIP() []byte {
	file_pbf_metric_create_proto_rawDescOnce.Do(func() {
		file_pbf_metric_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_metric_create_proto_rawDescData)
	})
	return file_pbf_metric_create_proto_rawDescData
}

var file_pbf_metric_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbf_metric_create_proto_goTypes = []interface{}{
	(*CreateI)(nil), // 0: metric.CreateI
	(*CreateO)(nil), // 1: metric.CreateO
}
var file_pbf_metric_create_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbf_metric_create_proto_init() }
func file_pbf_metric_create_proto_init() {
	if File_pbf_metric_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_metric_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateI); i {
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
		file_pbf_metric_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateO); i {
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
			RawDescriptor: file_pbf_metric_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_metric_create_proto_goTypes,
		DependencyIndexes: file_pbf_metric_create_proto_depIdxs,
		MessageInfos:      file_pbf_metric_create_proto_msgTypes,
	}.Build()
	File_pbf_metric_create_proto = out.File
	file_pbf_metric_create_proto_rawDesc = nil
	file_pbf_metric_create_proto_goTypes = nil
	file_pbf_metric_create_proto_depIdxs = nil
}
