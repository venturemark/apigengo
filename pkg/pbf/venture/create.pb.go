// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: pbf/venture/create.proto

package venture

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

// CreateI is the input for creating ventures. Below is an example JSON
// representation showing an emitted venture. The name in the example expresses
// the context of this particular venture.
//
//     {
//         "obj": [
//             {
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
type CreateI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *CreateI_API   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Obj []*CreateI_Obj `protobuf:"bytes,2,rep,name=obj,proto3" json:"obj,omitempty"`
}

func (x *CreateI) Reset() {
	*x = CreateI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI) ProtoMessage() {}

func (x *CreateI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_create_proto_msgTypes[0]
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
	return file_pbf_venture_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateI) GetApi() *CreateI_API {
	if x != nil {
		return x.Api
	}
	return nil
}

func (x *CreateI) GetObj() []*CreateI_Obj {
	if x != nil {
		return x.Obj
	}
	return nil
}

type CreateI_API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateI_API) Reset() {
	*x = CreateI_API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_API) ProtoMessage() {}

func (x *CreateI_API) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateI_API.ProtoReflect.Descriptor instead.
func (*CreateI_API) Descriptor() ([]byte, []int) {
	return file_pbf_venture_create_proto_rawDescGZIP(), []int{1}
}

type CreateI_Obj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string     `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Property *CreateI_Obj_Property `protobuf:"bytes,2,opt,name=property,proto3" json:"property,omitempty"`
}

func (x *CreateI_Obj) Reset() {
	*x = CreateI_Obj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_Obj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_Obj) ProtoMessage() {}

func (x *CreateI_Obj) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_create_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateI_Obj.ProtoReflect.Descriptor instead.
func (*CreateI_Obj) Descriptor() ([]byte, []int) {
	return file_pbf_venture_create_proto_rawDescGZIP(), []int{2}
}

func (x *CreateI_Obj) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateI_Obj) GetProperty() *CreateI_Obj_Property {
	if x != nil {
		return x.Property
	}
	return nil
}

type CreateI_Obj_Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc string                       `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	Link []*CreateI_Obj_Property_Link `protobuf:"bytes,2,rep,name=link,proto3" json:"link,omitempty"`
	Name string                       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateI_Obj_Property) Reset() {
	*x = CreateI_Obj_Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_create_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_Obj_Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_Obj_Property) ProtoMessage() {}

func (x *CreateI_Obj_Property) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_create_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateI_Obj_Property.ProtoReflect.Descriptor instead.
func (*CreateI_Obj_Property) Descriptor() ([]byte, []int) {
	return file_pbf_venture_create_proto_rawDescGZIP(), []int{3}
}

func (x *CreateI_Obj_Property) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateI_Obj_Property) GetLink() []*CreateI_Obj_Property_Link {
	if x != nil {
		return x.Link
	}
	return nil
}

func (x *CreateI_Obj_Property) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateI_Obj_Property_Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateI_Obj_Property_Link) Reset() {
	*x = CreateI_Obj_Property_Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_create_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_Obj_Property_Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_Obj_Property_Link) ProtoMessage() {}

func (x *CreateI_Obj_Property_Link) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_create_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateI_Obj_Property_Link.ProtoReflect.Descriptor instead.
func (*CreateI_Obj_Property_Link) Descriptor() ([]byte, []int) {
	return file_pbf_venture_create_proto_rawDescGZIP(), []int{4}
}

func (x *CreateI_Obj_Property_Link) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *CreateI_Obj_Property_Link) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// CreateO is the output for creating ventures. Only the exact unix timestamp of
// creation is returned with the object metadata when successfully creating a
// venture.
//
//     {
//         "obj": [
//             {
//                 "metadata": {
//                     "venture.venturemark.co/id": "<id>"
//                 }
//             }
//         ]
//     }
//
type CreateO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *CreateO_API   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Obj []*CreateO_Obj `protobuf:"bytes,2,rep,name=obj,proto3" json:"obj,omitempty"`
}

func (x *CreateO) Reset() {
	*x = CreateO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_create_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO) ProtoMessage() {}

func (x *CreateO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_create_proto_msgTypes[5]
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
	return file_pbf_venture_create_proto_rawDescGZIP(), []int{5}
}

func (x *CreateO) GetApi() *CreateO_API {
	if x != nil {
		return x.Api
	}
	return nil
}

func (x *CreateO) GetObj() []*CreateO_Obj {
	if x != nil {
		return x.Obj
	}
	return nil
}

type CreateO_API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateO_API) Reset() {
	*x = CreateO_API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_create_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO_API) ProtoMessage() {}

func (x *CreateO_API) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_create_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateO_API.ProtoReflect.Descriptor instead.
func (*CreateO_API) Descriptor() ([]byte, []int) {
	return file_pbf_venture_create_proto_rawDescGZIP(), []int{6}
}

type CreateO_Obj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateO_Obj) Reset() {
	*x = CreateO_Obj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_venture_create_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO_Obj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO_Obj) ProtoMessage() {}

func (x *CreateO_Obj) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_venture_create_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateO_Obj.ProtoReflect.Descriptor instead.
func (*CreateO_Obj) Descriptor() ([]byte, []int) {
	return file_pbf_venture_create_proto_rawDescGZIP(), []int{7}
}

func (x *CreateO_Obj) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_pbf_venture_create_proto protoreflect.FileDescriptor

var file_pbf_venture_create_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x66, 0x2f, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x65, 0x6e, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x59, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x12, 0x26,
	0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x65,
	0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x41, 0x50,
	0x49, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x26, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x52, 0x03, 0x6f, 0x62, 0x6a, 0x22, 0x0d,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x41, 0x50, 0x49, 0x22, 0xc5, 0x01,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x12, 0x3e, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x76, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x5f, 0x4f, 0x62, 0x6a, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x12, 0x36, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x5f, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x59, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x12, 0x26, 0x0a,
	0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x65, 0x6e,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x41, 0x50, 0x49,
	0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x26, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x52, 0x03, 0x6f, 0x62, 0x6a, 0x22, 0x0d, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x41, 0x50, 0x49, 0x22, 0x8a, 0x01, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x12, 0x3e, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x5f, 0x4f, 0x62, 0x6a, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b,
	0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_venture_create_proto_rawDescOnce sync.Once
	file_pbf_venture_create_proto_rawDescData = file_pbf_venture_create_proto_rawDesc
)

func file_pbf_venture_create_proto_rawDescGZIP() []byte {
	file_pbf_venture_create_proto_rawDescOnce.Do(func() {
		file_pbf_venture_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_venture_create_proto_rawDescData)
	})
	return file_pbf_venture_create_proto_rawDescData
}

var file_pbf_venture_create_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pbf_venture_create_proto_goTypes = []interface{}{
	(*CreateI)(nil),                   // 0: venture.CreateI
	(*CreateI_API)(nil),               // 1: venture.CreateI_API
	(*CreateI_Obj)(nil),               // 2: venture.CreateI_Obj
	(*CreateI_Obj_Property)(nil),      // 3: venture.CreateI_Obj_Property
	(*CreateI_Obj_Property_Link)(nil), // 4: venture.CreateI_Obj_Property_Link
	(*CreateO)(nil),                   // 5: venture.CreateO
	(*CreateO_API)(nil),               // 6: venture.CreateO_API
	(*CreateO_Obj)(nil),               // 7: venture.CreateO_Obj
	nil,                               // 8: venture.CreateI_Obj.MetadataEntry
	nil,                               // 9: venture.CreateO_Obj.MetadataEntry
}
var file_pbf_venture_create_proto_depIdxs = []int32{
	1, // 0: venture.CreateI.api:type_name -> venture.CreateI_API
	2, // 1: venture.CreateI.obj:type_name -> venture.CreateI_Obj
	8, // 2: venture.CreateI_Obj.metadata:type_name -> venture.CreateI_Obj.MetadataEntry
	3, // 3: venture.CreateI_Obj.property:type_name -> venture.CreateI_Obj_Property
	4, // 4: venture.CreateI_Obj_Property.link:type_name -> venture.CreateI_Obj_Property_Link
	6, // 5: venture.CreateO.api:type_name -> venture.CreateO_API
	7, // 6: venture.CreateO.obj:type_name -> venture.CreateO_Obj
	9, // 7: venture.CreateO_Obj.metadata:type_name -> venture.CreateO_Obj.MetadataEntry
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pbf_venture_create_proto_init() }
func file_pbf_venture_create_proto_init() {
	if File_pbf_venture_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_venture_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_venture_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateI_API); i {
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
		file_pbf_venture_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateI_Obj); i {
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
		file_pbf_venture_create_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateI_Obj_Property); i {
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
		file_pbf_venture_create_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateI_Obj_Property_Link); i {
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
		file_pbf_venture_create_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_venture_create_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateO_API); i {
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
		file_pbf_venture_create_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateO_Obj); i {
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
			RawDescriptor: file_pbf_venture_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_venture_create_proto_goTypes,
		DependencyIndexes: file_pbf_venture_create_proto_depIdxs,
		MessageInfos:      file_pbf_venture_create_proto_msgTypes,
	}.Build()
	File_pbf_venture_create_proto = out.File
	file_pbf_venture_create_proto_rawDesc = nil
	file_pbf_venture_create_proto_goTypes = nil
	file_pbf_venture_create_proto_depIdxs = nil
}