// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pbf/metupd/update.proto

package metupd

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

// UpdateI is the input for updating metric updates. That is persisting
// datapoints and text which belong together in order to form a metric update.
// Metrics and updates are separate resources. For lifecycle events like
// updating and deleting we want to manage a single entity. That is to reduce
// operational complexity and thus minimize failure scenarious. Otherwise
// updating metrics and updates separately would imply the technical challenge
// of reliably associating each resource with the other after the fact.
//
// A timeline can have many metric updates. A timeline thus can have many
// metrics and many updates each while any metric is only associated with a
// single update respectively. Below is an example JSON representation showing
// an emitted metric update for the dimensions x and y, where x represents the
// unix timestamp of server side creation and y represents the numeric values
// given by the user at update time. Two y axis values imply two visualized
// graphs in the UI. Additionally the user's written update in natural language
// is provided. Updating metric updates does only allow to modify the yaxis and
// text properties.
//
//     {
//         "yaxis": [
//             32,
//             85
//         ],
//         "text": "Lorem ipsum ...",
//         "timeline": "tml-kn433",
//         "timestamp": 1604959525
//     }
//
type UpdateI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// yaxis contains the metric part of the metric update. The datapoints
	// provided are coordinates relevant for the user. Providing [y1, y2] may be
	// used to visualize two graphs at the same time in the UI.
	Yaxis []int64 `protobuf:"varint,1,rep,packed,name=yaxis,proto3" json:"yaxis,omitempty"`
	// text contains the update part of the metric update. The string provided is
	// the user's natural language in written form.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// timeline is the identifier used to associate the metric update with the
	// specified timeline.
	Timeline string `protobuf:"bytes,3,opt,name=timeline,proto3" json:"timeline,omitempty"`
	// timestamp is the time at which the metric update got created. The timestamp
	// is a unix timestamp in seconds normalized to the UTC timezone.
	Timestamp int64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *UpdateI) Reset() {
	*x = UpdateI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_metupd_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI) ProtoMessage() {}

func (x *UpdateI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metupd_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateI.ProtoReflect.Descriptor instead.
func (*UpdateI) Descriptor() ([]byte, []int) {
	return file_pbf_metupd_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateI) GetYaxis() []int64 {
	if x != nil {
		return x.Yaxis
	}
	return nil
}

func (x *UpdateI) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UpdateI) GetTimeline() string {
	if x != nil {
		return x.Timeline
	}
	return ""
}

func (x *UpdateI) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// UpdateO is the output for updating metric updates. Only the exact unix
// timestamp of creation is returned when successfully updating a metric update.
//
//     {
//         "timestamp": 1604959525
//     }
//
type UpdateO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// timestamp is the time at which the metric update got created. The timestamp
	// is a unix timestamp in seconds normalized to the UTC timezone.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *UpdateO) Reset() {
	*x = UpdateO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_metupd_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateO) ProtoMessage() {}

func (x *UpdateO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metupd_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateO.ProtoReflect.Descriptor instead.
func (*UpdateO) Descriptor() ([]byte, []int) {
	return file_pbf_metupd_update_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateO) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_pbf_metupd_update_proto protoreflect.FileDescriptor

var file_pbf_metupd_update_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x75, 0x70, 0x64, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x74, 0x75, 0x70,
	0x64, 0x22, 0x6d, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x12, 0x14, 0x0a, 0x05,
	0x79, 0x61, 0x78, 0x69, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x79, 0x61, 0x78,
	0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x27, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x6d,
	0x65, 0x74, 0x75, 0x70, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_metupd_update_proto_rawDescOnce sync.Once
	file_pbf_metupd_update_proto_rawDescData = file_pbf_metupd_update_proto_rawDesc
)

func file_pbf_metupd_update_proto_rawDescGZIP() []byte {
	file_pbf_metupd_update_proto_rawDescOnce.Do(func() {
		file_pbf_metupd_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_metupd_update_proto_rawDescData)
	})
	return file_pbf_metupd_update_proto_rawDescData
}

var file_pbf_metupd_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbf_metupd_update_proto_goTypes = []interface{}{
	(*UpdateI)(nil), // 0: metupd.UpdateI
	(*UpdateO)(nil), // 1: metupd.UpdateO
}
var file_pbf_metupd_update_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbf_metupd_update_proto_init() }
func file_pbf_metupd_update_proto_init() {
	if File_pbf_metupd_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_metupd_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateI); i {
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
		file_pbf_metupd_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateO); i {
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
			RawDescriptor: file_pbf_metupd_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_metupd_update_proto_goTypes,
		DependencyIndexes: file_pbf_metupd_update_proto_depIdxs,
		MessageInfos:      file_pbf_metupd_update_proto_msgTypes,
	}.Build()
	File_pbf_metupd_update_proto = out.File
	file_pbf_metupd_update_proto_rawDesc = nil
	file_pbf_metupd_update_proto_goTypes = nil
	file_pbf_metupd_update_proto_depIdxs = nil
}
