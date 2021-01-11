// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pbf/texupd/update.proto

package texupd

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

// UpdateI is the input for updating text updates. For more information about
// text updates as a separate resource see create.proto.
//
//     {
//         "obj": {
//             "metadata": {
//                 "audience.venturemark.co/id": "aud-klc5p",
//                 "organization.venturemark.co/id": "org-pl142",
//                 "timeline.venturemark.co/id": "1606396471",
//                 "update.venturemark.co/id": "1606400781",
//                 "user.venturemark.co/id": "usr-al9qy"
//             }
//             "property": {
//                 "text": "Lorem ipsum ..."
//             }
//         }
//     }
//
type UpdateI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *UpdateI_API `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Obj *UpdateI_Obj `protobuf:"bytes,2,opt,name=obj,proto3" json:"obj,omitempty"`
}

func (x *UpdateI) Reset() {
	*x = UpdateI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_texupd_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI) ProtoMessage() {}

func (x *UpdateI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_texupd_update_proto_msgTypes[0]
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
	return file_pbf_texupd_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateI) GetApi() *UpdateI_API {
	if x != nil {
		return x.Api
	}
	return nil
}

func (x *UpdateI) GetObj() *UpdateI_Obj {
	if x != nil {
		return x.Obj
	}
	return nil
}

type UpdateI_API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateI_API) Reset() {
	*x = UpdateI_API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_texupd_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI_API) ProtoMessage() {}

func (x *UpdateI_API) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_texupd_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateI_API.ProtoReflect.Descriptor instead.
func (*UpdateI_API) Descriptor() ([]byte, []int) {
	return file_pbf_texupd_update_proto_rawDescGZIP(), []int{1}
}

type UpdateI_Obj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string     `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Property *UpdateI_Obj_Property `protobuf:"bytes,2,opt,name=property,proto3" json:"property,omitempty"`
}

func (x *UpdateI_Obj) Reset() {
	*x = UpdateI_Obj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_texupd_update_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI_Obj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI_Obj) ProtoMessage() {}

func (x *UpdateI_Obj) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_texupd_update_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateI_Obj.ProtoReflect.Descriptor instead.
func (*UpdateI_Obj) Descriptor() ([]byte, []int) {
	return file_pbf_texupd_update_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateI_Obj) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateI_Obj) GetProperty() *UpdateI_Obj_Property {
	if x != nil {
		return x.Property
	}
	return nil
}

type UpdateI_Obj_Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *UpdateI_Obj_Property) Reset() {
	*x = UpdateI_Obj_Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_texupd_update_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI_Obj_Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI_Obj_Property) ProtoMessage() {}

func (x *UpdateI_Obj_Property) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_texupd_update_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateI_Obj_Property.ProtoReflect.Descriptor instead.
func (*UpdateI_Obj_Property) Descriptor() ([]byte, []int) {
	return file_pbf_texupd_update_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateI_Obj_Property) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// UpdateO is the output for updating text updates. The response will contain
// object metadata to indicate the text update got in fact updated. The example
// below indicates that the update resource got updated. That is, the text
// changed.
//
//     {
//         "obj": {
//             "metadata": {
//                 "update.venturemark.co/status": "updated"
//             }
//     }
//
type UpdateO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *UpdateO_API `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Obj *UpdateO_Obj `protobuf:"bytes,2,opt,name=obj,proto3" json:"obj,omitempty"`
}

func (x *UpdateO) Reset() {
	*x = UpdateO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_texupd_update_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateO) ProtoMessage() {}

func (x *UpdateO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_texupd_update_proto_msgTypes[4]
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
	return file_pbf_texupd_update_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateO) GetApi() *UpdateO_API {
	if x != nil {
		return x.Api
	}
	return nil
}

func (x *UpdateO) GetObj() *UpdateO_Obj {
	if x != nil {
		return x.Obj
	}
	return nil
}

type UpdateO_API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateO_API) Reset() {
	*x = UpdateO_API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_texupd_update_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateO_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateO_API) ProtoMessage() {}

func (x *UpdateO_API) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_texupd_update_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateO_API.ProtoReflect.Descriptor instead.
func (*UpdateO_API) Descriptor() ([]byte, []int) {
	return file_pbf_texupd_update_proto_rawDescGZIP(), []int{5}
}

type UpdateO_Obj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpdateO_Obj) Reset() {
	*x = UpdateO_Obj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_texupd_update_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateO_Obj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateO_Obj) ProtoMessage() {}

func (x *UpdateO_Obj) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_texupd_update_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateO_Obj.ProtoReflect.Descriptor instead.
func (*UpdateO_Obj) Descriptor() ([]byte, []int) {
	return file_pbf_texupd_update_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateO_Obj) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_pbf_texupd_update_proto protoreflect.FileDescriptor

var file_pbf_texupd_update_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x66, 0x2f, 0x74, 0x65, 0x78, 0x75, 0x70, 0x64, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x78, 0x75, 0x70,
	0x64, 0x22, 0x57, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x12, 0x25, 0x0a, 0x03,
	0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x78, 0x75,
	0x70, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x41, 0x50, 0x49, 0x52, 0x03,
	0x61, 0x70, 0x69, 0x12, 0x25, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x74, 0x65, 0x78, 0x75, 0x70, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x52, 0x03, 0x6f, 0x62, 0x6a, 0x22, 0x0d, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x41, 0x50, 0x49, 0x22, 0xc3, 0x01, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65,
	0x78, 0x75, 0x70, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x78,
	0x75, 0x70, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x5f,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x2a, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x5f, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x57, 0x0a, 0x07, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x12, 0x25, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x78, 0x75, 0x70, 0x64, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x5f, 0x41, 0x50, 0x49, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x25, 0x0a,
	0x03, 0x6f, 0x62, 0x6a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x78,
	0x75, 0x70, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x52,
	0x03, 0x6f, 0x62, 0x6a, 0x22, 0x0d, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x5f,
	0x41, 0x50, 0x49, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x5f,
	0x4f, 0x62, 0x6a, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x78, 0x75, 0x70, 0x64, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x74, 0x65, 0x78, 0x75, 0x70, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pbf_texupd_update_proto_rawDescOnce sync.Once
	file_pbf_texupd_update_proto_rawDescData = file_pbf_texupd_update_proto_rawDesc
)

func file_pbf_texupd_update_proto_rawDescGZIP() []byte {
	file_pbf_texupd_update_proto_rawDescOnce.Do(func() {
		file_pbf_texupd_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_texupd_update_proto_rawDescData)
	})
	return file_pbf_texupd_update_proto_rawDescData
}

var file_pbf_texupd_update_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pbf_texupd_update_proto_goTypes = []interface{}{
	(*UpdateI)(nil),              // 0: texupd.UpdateI
	(*UpdateI_API)(nil),          // 1: texupd.UpdateI_API
	(*UpdateI_Obj)(nil),          // 2: texupd.UpdateI_Obj
	(*UpdateI_Obj_Property)(nil), // 3: texupd.UpdateI_Obj_Property
	(*UpdateO)(nil),              // 4: texupd.UpdateO
	(*UpdateO_API)(nil),          // 5: texupd.UpdateO_API
	(*UpdateO_Obj)(nil),          // 6: texupd.UpdateO_Obj
	nil,                          // 7: texupd.UpdateI_Obj.MetadataEntry
	nil,                          // 8: texupd.UpdateO_Obj.MetadataEntry
}
var file_pbf_texupd_update_proto_depIdxs = []int32{
	1, // 0: texupd.UpdateI.api:type_name -> texupd.UpdateI_API
	2, // 1: texupd.UpdateI.obj:type_name -> texupd.UpdateI_Obj
	7, // 2: texupd.UpdateI_Obj.metadata:type_name -> texupd.UpdateI_Obj.MetadataEntry
	3, // 3: texupd.UpdateI_Obj.property:type_name -> texupd.UpdateI_Obj_Property
	5, // 4: texupd.UpdateO.api:type_name -> texupd.UpdateO_API
	6, // 5: texupd.UpdateO.obj:type_name -> texupd.UpdateO_Obj
	8, // 6: texupd.UpdateO_Obj.metadata:type_name -> texupd.UpdateO_Obj.MetadataEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pbf_texupd_update_proto_init() }
func file_pbf_texupd_update_proto_init() {
	if File_pbf_texupd_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_texupd_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_texupd_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateI_API); i {
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
		file_pbf_texupd_update_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateI_Obj); i {
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
		file_pbf_texupd_update_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateI_Obj_Property); i {
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
		file_pbf_texupd_update_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_texupd_update_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateO_API); i {
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
		file_pbf_texupd_update_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateO_Obj); i {
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
			RawDescriptor: file_pbf_texupd_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_texupd_update_proto_goTypes,
		DependencyIndexes: file_pbf_texupd_update_proto_depIdxs,
		MessageInfos:      file_pbf_texupd_update_proto_msgTypes,
	}.Build()
	File_pbf_texupd_update_proto = out.File
	file_pbf_texupd_update_proto_rawDesc = nil
	file_pbf_texupd_update_proto_goTypes = nil
	file_pbf_texupd_update_proto_depIdxs = nil
}