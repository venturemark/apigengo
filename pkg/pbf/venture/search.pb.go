// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pbf/venture/search.proto

package venture

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

// SearchI is the input for searching ventures. Below is an example to fetch a
// particular venture by its ID.
//
//     {
//         "api": {
//             "chunking": {
//                 "perpage": "100",
//                 "pointer": "300"
//             }
//         },
//         "obj": [
//             {
//                 "metadata": {
//                     "venture.venturemark.co/id": "<id>"
//                 }
//             }
//         }
//     }
//
// Ventures can also be searched by user. Below is an example to fetch a all
// ventures a user is part of. Note that the user ID must be provided with the
// subject ID metadata.
//
//     {
//         "api": {
//             "chunking": {
//                 "perpage": "100",
//                 "pointer": "300"
//             }
//         },
//         "obj": [
//             {
//                 "metadata": {
//                     "subject.venturemark.co/id": "<id>"
//                 }
//             }
//         }
//     }
//
type SearchI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *SearchI_API   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Obj []*SearchI_Obj `protobuf:"bytes,2,rep,name=obj,proto3" json:"obj,omitempty"`
}

func (x *SearchI) Reset() {
	*x = SearchI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI) ProtoMessage() {}

func (x *SearchI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[0]
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
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchI) GetApi() *SearchI_API {
	if x != nil {
		return x.Api
	}
	return nil
}

func (x *SearchI) GetObj() []*SearchI_Obj {
	if x != nil {
		return x.Obj
	}
	return nil
}

type SearchI_API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunking *SearchI_API_Chunking `protobuf:"bytes,1,opt,name=chunking,proto3" json:"chunking,omitempty"`
	Operator []string              `protobuf:"bytes,2,rep,name=operator,proto3" json:"operator,omitempty"`
}

func (x *SearchI_API) Reset() {
	*x = SearchI_API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_API) ProtoMessage() {}

func (x *SearchI_API) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_API.ProtoReflect.Descriptor instead.
func (*SearchI_API) Descriptor() ([]byte, []int) {
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchI_API) GetChunking() *SearchI_API_Chunking {
	if x != nil {
		return x.Chunking
	}
	return nil
}

func (x *SearchI_API) GetOperator() []string {
	if x != nil {
		return x.Operator
	}
	return nil
}

type SearchI_API_Chunking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pointer string `protobuf:"bytes,1,opt,name=pointer,proto3" json:"pointer,omitempty"`
	Perpage string `protobuf:"bytes,2,opt,name=perpage,proto3" json:"perpage,omitempty"`
}

func (x *SearchI_API_Chunking) Reset() {
	*x = SearchI_API_Chunking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_API_Chunking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_API_Chunking) ProtoMessage() {}

func (x *SearchI_API_Chunking) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_API_Chunking.ProtoReflect.Descriptor instead.
func (*SearchI_API_Chunking) Descriptor() ([]byte, []int) {
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchI_API_Chunking) GetPointer() string {
	if x != nil {
		return x.Pointer
	}
	return ""
}

func (x *SearchI_API_Chunking) GetPerpage() string {
	if x != nil {
		return x.Perpage
	}
	return ""
}

type SearchI_Obj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string     `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Property *SearchI_Obj_Property `protobuf:"bytes,2,opt,name=property,proto3" json:"property,omitempty"`
}

func (x *SearchI_Obj) Reset() {
	*x = SearchI_Obj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Obj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Obj) ProtoMessage() {}

func (x *SearchI_Obj) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Obj.ProtoReflect.Descriptor instead.
func (*SearchI_Obj) Descriptor() ([]byte, []int) {
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{3}
}

func (x *SearchI_Obj) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SearchI_Obj) GetProperty() *SearchI_Obj_Property {
	if x != nil {
		return x.Property
	}
	return nil
}

type SearchI_Obj_Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchI_Obj_Property) Reset() {
	*x = SearchI_Obj_Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Obj_Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Obj_Property) ProtoMessage() {}

func (x *SearchI_Obj_Property) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Obj_Property.ProtoReflect.Descriptor instead.
func (*SearchI_Obj_Property) Descriptor() ([]byte, []int) {
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{4}
}

// SearchO is the output for searching ventures. Below is an example showing a
// list of objects associated with the requested user.
//
//     {
//         "api": {
//             "chunking": {
//                 "perpage": "100",
//                 "pointer": "300"
//             }
//         },
//         "obj": [
//             {
//                 "metadata": {
//                     "venture.venturemark.co/id": "<id>"
//                 },
//                 "property": {
//                     "desc": "Lorem ipsum ...",
//                     "link": [
//                         {
//                             "addr": "https://twitter.com/ibm",
//                             "text": "Twitter"
//                         }
//                     ],
//                     "name": "IBM"
//                 }
//             }
//         ]
//     }
//
type SearchO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *SearchO_API   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Obj []*SearchO_Obj `protobuf:"bytes,2,rep,name=obj,proto3" json:"obj,omitempty"`
}

func (x *SearchO) Reset() {
	*x = SearchO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO) ProtoMessage() {}

func (x *SearchO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[5]
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
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{5}
}

func (x *SearchO) GetApi() *SearchO_API {
	if x != nil {
		return x.Api
	}
	return nil
}

func (x *SearchO) GetObj() []*SearchO_Obj {
	if x != nil {
		return x.Obj
	}
	return nil
}

type SearchO_API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunking *SearchO_API_Chunking `protobuf:"bytes,1,opt,name=chunking,proto3" json:"chunking,omitempty"`
}

func (x *SearchO_API) Reset() {
	*x = SearchO_API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_API) ProtoMessage() {}

func (x *SearchO_API) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_API.ProtoReflect.Descriptor instead.
func (*SearchO_API) Descriptor() ([]byte, []int) {
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{6}
}

func (x *SearchO_API) GetChunking() *SearchO_API_Chunking {
	if x != nil {
		return x.Chunking
	}
	return nil
}

type SearchO_API_Chunking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pointer string `protobuf:"bytes,1,opt,name=pointer,proto3" json:"pointer,omitempty"`
	Perpage string `protobuf:"bytes,2,opt,name=perpage,proto3" json:"perpage,omitempty"`
}

func (x *SearchO_API_Chunking) Reset() {
	*x = SearchO_API_Chunking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_API_Chunking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_API_Chunking) ProtoMessage() {}

func (x *SearchO_API_Chunking) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_API_Chunking.ProtoReflect.Descriptor instead.
func (*SearchO_API_Chunking) Descriptor() ([]byte, []int) {
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{7}
}

func (x *SearchO_API_Chunking) GetPointer() string {
	if x != nil {
		return x.Pointer
	}
	return ""
}

func (x *SearchO_API_Chunking) GetPerpage() string {
	if x != nil {
		return x.Perpage
	}
	return ""
}

type SearchO_Obj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string     `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Property *SearchO_Obj_Property `protobuf:"bytes,2,opt,name=property,proto3" json:"property,omitempty"`
}

func (x *SearchO_Obj) Reset() {
	*x = SearchO_Obj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Obj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Obj) ProtoMessage() {}

func (x *SearchO_Obj) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Obj.ProtoReflect.Descriptor instead.
func (*SearchO_Obj) Descriptor() ([]byte, []int) {
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{8}
}

func (x *SearchO_Obj) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SearchO_Obj) GetProperty() *SearchO_Obj_Property {
	if x != nil {
		return x.Property
	}
	return nil
}

type SearchO_Obj_Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc string                       `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	Link []*SearchO_Obj_Property_Link `protobuf:"bytes,2,rep,name=link,proto3" json:"link,omitempty"`
	Name string                       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SearchO_Obj_Property) Reset() {
	*x = SearchO_Obj_Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Obj_Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Obj_Property) ProtoMessage() {}

func (x *SearchO_Obj_Property) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Obj_Property.ProtoReflect.Descriptor instead.
func (*SearchO_Obj_Property) Descriptor() ([]byte, []int) {
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{9}
}

func (x *SearchO_Obj_Property) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *SearchO_Obj_Property) GetLink() []*SearchO_Obj_Property_Link {
	if x != nil {
		return x.Link
	}
	return nil
}

func (x *SearchO_Obj_Property) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SearchO_Obj_Property_Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SearchO_Obj_Property_Link) Reset() {
	*x = SearchO_Obj_Property_Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_search_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Obj_Property_Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Obj_Property_Link) ProtoMessage() {}

func (x *SearchO_Obj_Property_Link) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_search_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Obj_Property_Link.ProtoReflect.Descriptor instead.
func (*SearchO_Obj_Property_Link) Descriptor() ([]byte, []int) {
	return file_pbf_venture_search_proto_rawDescGZIP(), []int{10}
}

func (x *SearchO_Obj_Property_Link) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *SearchO_Obj_Property_Link) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_pbf_venture_search_proto protoreflect.FileDescriptor

var file_pbf_venture_search_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x66, 0x2f, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x65, 0x6e, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x59, 0x0a, 0x07, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x12, 0x26,
	0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x65,
	0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x41, 0x50,
	0x49, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x26, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x52, 0x03, 0x6f, 0x62, 0x6a, 0x22, 0x64,
	0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x41, 0x50, 0x49, 0x12, 0x39, 0x0a,
	0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x49, 0x5f, 0x41, 0x50, 0x49, 0x5f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x22, 0x4a, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f,
	0x41, 0x50, 0x49, 0x5f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x72, 0x70, 0x61, 0x67, 0x65,
	0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a,
	0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x39, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x22, 0x59, 0x0a, 0x07, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x12, 0x26, 0x0a, 0x03, 0x61,
	0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x41, 0x50, 0x49, 0x52, 0x03,
	0x61, 0x70, 0x69, 0x12, 0x26, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x52, 0x03, 0x6f, 0x62, 0x6a, 0x22, 0x48, 0x0a, 0x0b, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x41, 0x50, 0x49, 0x12, 0x39, 0x0a, 0x08, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76,
	0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x41,
	0x50, 0x49, 0x5f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x4a, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f,
	0x5f, 0x41, 0x50, 0x49, 0x5f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x72, 0x70, 0x61, 0x67,
	0x65, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62,
	0x6a, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x39, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x76, 0x0a, 0x14, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x36, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x5f, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x43, 0x0a, 0x19, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a,
	0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x76, 0x65, 0x6e, 0x74,
	0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_venture_search_proto_rawDescOnce sync.Once
	file_pbf_venture_search_proto_rawDescData = file_pbf_venture_search_proto_rawDesc
)

func file_pbf_venture_search_proto_rawDescGZIP() []byte {
	file_pbf_venture_search_proto_rawDescOnce.Do(func() {
		file_pbf_venture_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_venture_search_proto_rawDescData)
	})
	return file_pbf_venture_search_proto_rawDescData
}

var file_pbf_venture_search_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pbf_venture_search_proto_goTypes = []interface{}{
	(*SearchI)(nil),                   // 0: venture.SearchI
	(*SearchI_API)(nil),               // 1: venture.SearchI_API
	(*SearchI_API_Chunking)(nil),      // 2: venture.SearchI_API_Chunking
	(*SearchI_Obj)(nil),               // 3: venture.SearchI_Obj
	(*SearchI_Obj_Property)(nil),      // 4: venture.SearchI_Obj_Property
	(*SearchO)(nil),                   // 5: venture.SearchO
	(*SearchO_API)(nil),               // 6: venture.SearchO_API
	(*SearchO_API_Chunking)(nil),      // 7: venture.SearchO_API_Chunking
	(*SearchO_Obj)(nil),               // 8: venture.SearchO_Obj
	(*SearchO_Obj_Property)(nil),      // 9: venture.SearchO_Obj_Property
	(*SearchO_Obj_Property_Link)(nil), // 10: venture.SearchO_Obj_Property_Link
	nil,                               // 11: venture.SearchI_Obj.MetadataEntry
	nil,                               // 12: venture.SearchO_Obj.MetadataEntry
}
var file_pbf_venture_search_proto_depIdxs = []int32{
	1,  // 0: venture.SearchI.api:type_name -> venture.SearchI_API
	3,  // 1: venture.SearchI.obj:type_name -> venture.SearchI_Obj
	2,  // 2: venture.SearchI_API.chunking:type_name -> venture.SearchI_API_Chunking
	11, // 3: venture.SearchI_Obj.metadata:type_name -> venture.SearchI_Obj.MetadataEntry
	4,  // 4: venture.SearchI_Obj.property:type_name -> venture.SearchI_Obj_Property
	6,  // 5: venture.SearchO.api:type_name -> venture.SearchO_API
	8,  // 6: venture.SearchO.obj:type_name -> venture.SearchO_Obj
	7,  // 7: venture.SearchO_API.chunking:type_name -> venture.SearchO_API_Chunking
	12, // 8: venture.SearchO_Obj.metadata:type_name -> venture.SearchO_Obj.MetadataEntry
	9,  // 9: venture.SearchO_Obj.property:type_name -> venture.SearchO_Obj_Property
	10, // 10: venture.SearchO_Obj_Property.link:type_name -> venture.SearchO_Obj_Property_Link
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pbf_venture_search_proto_init() }
func file_pbf_venture_search_proto_init() {
	if File_pbf_venture_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_venture_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_venture_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_API); i {
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
		file_pbf_venture_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_API_Chunking); i {
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
		file_pbf_venture_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Obj); i {
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
		file_pbf_venture_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Obj_Property); i {
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
		file_pbf_venture_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_venture_search_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_API); i {
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
		file_pbf_venture_search_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_API_Chunking); i {
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
		file_pbf_venture_search_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Obj); i {
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
		file_pbf_venture_search_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Obj_Property); i {
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
		file_pbf_venture_search_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Obj_Property_Link); i {
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
			RawDescriptor: file_pbf_venture_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_venture_search_proto_goTypes,
		DependencyIndexes: file_pbf_venture_search_proto_depIdxs,
		MessageInfos:      file_pbf_venture_search_proto_msgTypes,
	}.Build()
	File_pbf_venture_search_proto = out.File
	file_pbf_venture_search_proto_rawDesc = nil
	file_pbf_venture_search_proto_goTypes = nil
	file_pbf_venture_search_proto_depIdxs = nil
}
