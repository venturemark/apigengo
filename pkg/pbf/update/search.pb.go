// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pbf/update/search.proto

package update

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

// SearchI is the input for searching updates. That is looking up the user's
// natural language in written form associated with a timeline. Below is an
// example JSON representation showing a filter query meant to fetch all updates
// associated with the specified timeline.
//
//     {
//         "filter": {
//             "property": [
//                 {
//                     "timeline": "tml-kjw44"
//                 }
//             ]
//         }
//     }
//
type SearchI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *SearchI_Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *SearchI) Reset() {
	*x = SearchI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_update_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI) ProtoMessage() {}

func (x *SearchI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_update_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI.ProtoReflect.Descriptor instead.
func (*SearchI) Descriptor() ([]byte, []int) {
	return file_pbf_update_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchI) GetFilter() *SearchI_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

// SearchO is the input for searching updates. That is looking up the user's
// natural language in written form associated with a timeline. Below is an
// example JSON representation showing a list of results when fetching all
// updates associated with the specified timeline.
//
//     {
//         "result": [
//             {
//                 "text": "Lorem ipsum ...",
//                 "timeline": "tml-kn433",
//                 "timestamp": 1604959525
//             },
//             ...
//         ]
//     }
//
type SearchO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *SearchO_Filter   `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Result []*SearchO_Result `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *SearchO) Reset() {
	*x = SearchO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_update_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO) ProtoMessage() {}

func (x *SearchO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_update_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO.ProtoReflect.Descriptor instead.
func (*SearchO) Descriptor() ([]byte, []int) {
	return file_pbf_update_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchO) GetFilter() *SearchO_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *SearchO) GetResult() []*SearchO_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type SearchI_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunking *SearchI_Filter_Chunking   `protobuf:"bytes,1,opt,name=chunking,proto3" json:"chunking,omitempty"`
	Operator []string                   `protobuf:"bytes,2,rep,name=operator,proto3" json:"operator,omitempty"`
	Property []*SearchI_Filter_Property `protobuf:"bytes,3,rep,name=property,proto3" json:"property,omitempty"`
}

func (x *SearchI_Filter) Reset() {
	*x = SearchI_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_update_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Filter) ProtoMessage() {}

func (x *SearchI_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_update_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Filter.ProtoReflect.Descriptor instead.
func (*SearchI_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_update_search_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SearchI_Filter) GetChunking() *SearchI_Filter_Chunking {
	if x != nil {
		return x.Chunking
	}
	return nil
}

func (x *SearchI_Filter) GetOperator() []string {
	if x != nil {
		return x.Operator
	}
	return nil
}

func (x *SearchI_Filter) GetProperty() []*SearchI_Filter_Property {
	if x != nil {
		return x.Property
	}
	return nil
}

type SearchI_Filter_Chunking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pointer string `protobuf:"bytes,1,opt,name=pointer,proto3" json:"pointer,omitempty"`
	Perpage string `protobuf:"bytes,2,opt,name=perpage,proto3" json:"perpage,omitempty"`
}

func (x *SearchI_Filter_Chunking) Reset() {
	*x = SearchI_Filter_Chunking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_update_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Filter_Chunking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Filter_Chunking) ProtoMessage() {}

func (x *SearchI_Filter_Chunking) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_update_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Filter_Chunking.ProtoReflect.Descriptor instead.
func (*SearchI_Filter_Chunking) Descriptor() ([]byte, []int) {
	return file_pbf_update_search_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SearchI_Filter_Chunking) GetPointer() string {
	if x != nil {
		return x.Pointer
	}
	return ""
}

func (x *SearchI_Filter_Chunking) GetPerpage() string {
	if x != nil {
		return x.Perpage
	}
	return ""
}

type SearchI_Filter_Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeline  string `protobuf:"bytes,1,opt,name=timeline,proto3" json:"timeline,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SearchI_Filter_Property) Reset() {
	*x = SearchI_Filter_Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_update_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Filter_Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Filter_Property) ProtoMessage() {}

func (x *SearchI_Filter_Property) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_update_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Filter_Property.ProtoReflect.Descriptor instead.
func (*SearchI_Filter_Property) Descriptor() ([]byte, []int) {
	return file_pbf_update_search_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *SearchI_Filter_Property) GetTimeline() string {
	if x != nil {
		return x.Timeline
	}
	return ""
}

func (x *SearchI_Filter_Property) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type SearchO_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunking *SearchO_Filter_Chunking `protobuf:"bytes,1,opt,name=chunking,proto3" json:"chunking,omitempty"`
}

func (x *SearchO_Filter) Reset() {
	*x = SearchO_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_update_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Filter) ProtoMessage() {}

func (x *SearchO_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_update_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Filter.ProtoReflect.Descriptor instead.
func (*SearchO_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_update_search_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SearchO_Filter) GetChunking() *SearchO_Filter_Chunking {
	if x != nil {
		return x.Chunking
	}
	return nil
}

type SearchO_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text      string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Timeline  string `protobuf:"bytes,2,opt,name=timeline,proto3" json:"timeline,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SearchO_Result) Reset() {
	*x = SearchO_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_update_search_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Result) ProtoMessage() {}

func (x *SearchO_Result) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_update_search_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Result.ProtoReflect.Descriptor instead.
func (*SearchO_Result) Descriptor() ([]byte, []int) {
	return file_pbf_update_search_proto_rawDescGZIP(), []int{1, 1}
}

func (x *SearchO_Result) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SearchO_Result) GetTimeline() string {
	if x != nil {
		return x.Timeline
	}
	return ""
}

func (x *SearchO_Result) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type SearchO_Filter_Chunking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pointer string `protobuf:"bytes,1,opt,name=pointer,proto3" json:"pointer,omitempty"`
	Perpage string `protobuf:"bytes,2,opt,name=perpage,proto3" json:"perpage,omitempty"`
}

func (x *SearchO_Filter_Chunking) Reset() {
	*x = SearchO_Filter_Chunking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_update_search_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Filter_Chunking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Filter_Chunking) ProtoMessage() {}

func (x *SearchO_Filter_Chunking) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_update_search_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Filter_Chunking.ProtoReflect.Descriptor instead.
func (*SearchO_Filter_Chunking) Descriptor() ([]byte, []int) {
	return file_pbf_update_search_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *SearchO_Filter_Chunking) GetPointer() string {
	if x != nil {
		return x.Pointer
	}
	return ""
}

func (x *SearchO_Filter_Chunking) GetPerpage() string {
	if x != nil {
		return x.Perpage
	}
	return ""
}

var File_pbf_update_search_proto protoreflect.FileDescriptor

var file_pbf_update_search_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x66, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x22, 0xe0, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x12, 0x2e, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0xa4, 0x02,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x3b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x49, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x3e,
	0x0a, 0x08, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x72, 0x70, 0x61, 0x67, 0x65, 0x1a, 0x44,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0xc9, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f,
	0x12, 0x2e, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x1a, 0x85, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x08, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x3e, 0x0a, 0x08, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x65, 0x72, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x65, 0x72, 0x70, 0x61, 0x67, 0x65, 0x1a, 0x56, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_update_search_proto_rawDescOnce sync.Once
	file_pbf_update_search_proto_rawDescData = file_pbf_update_search_proto_rawDesc
)

func file_pbf_update_search_proto_rawDescGZIP() []byte {
	file_pbf_update_search_proto_rawDescOnce.Do(func() {
		file_pbf_update_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_update_search_proto_rawDescData)
	})
	return file_pbf_update_search_proto_rawDescData
}

var file_pbf_update_search_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pbf_update_search_proto_goTypes = []interface{}{
	(*SearchI)(nil),                 // 0: update.SearchI
	(*SearchO)(nil),                 // 1: update.SearchO
	(*SearchI_Filter)(nil),          // 2: update.SearchI.Filter
	(*SearchI_Filter_Chunking)(nil), // 3: update.SearchI.Filter.Chunking
	(*SearchI_Filter_Property)(nil), // 4: update.SearchI.Filter.Property
	(*SearchO_Filter)(nil),          // 5: update.SearchO.Filter
	(*SearchO_Result)(nil),          // 6: update.SearchO.Result
	(*SearchO_Filter_Chunking)(nil), // 7: update.SearchO.Filter.Chunking
}
var file_pbf_update_search_proto_depIdxs = []int32{
	2, // 0: update.SearchI.filter:type_name -> update.SearchI.Filter
	5, // 1: update.SearchO.filter:type_name -> update.SearchO.Filter
	6, // 2: update.SearchO.result:type_name -> update.SearchO.Result
	3, // 3: update.SearchI.Filter.chunking:type_name -> update.SearchI.Filter.Chunking
	4, // 4: update.SearchI.Filter.property:type_name -> update.SearchI.Filter.Property
	7, // 5: update.SearchO.Filter.chunking:type_name -> update.SearchO.Filter.Chunking
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pbf_update_search_proto_init() }
func file_pbf_update_search_proto_init() {
	if File_pbf_update_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_update_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI); i {
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
		file_pbf_update_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO); i {
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
		file_pbf_update_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Filter); i {
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
		file_pbf_update_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Filter_Chunking); i {
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
		file_pbf_update_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Filter_Property); i {
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
		file_pbf_update_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Filter); i {
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
		file_pbf_update_search_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Result); i {
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
		file_pbf_update_search_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Filter_Chunking); i {
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
			RawDescriptor: file_pbf_update_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_update_search_proto_goTypes,
		DependencyIndexes: file_pbf_update_search_proto_depIdxs,
		MessageInfos:      file_pbf_update_search_proto_msgTypes,
	}.Build()
	File_pbf_update_search_proto = out.File
	file_pbf_update_search_proto_rawDesc = nil
	file_pbf_update_search_proto_goTypes = nil
	file_pbf_update_search_proto_depIdxs = nil
}
