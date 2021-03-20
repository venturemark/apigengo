// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: pbf/role/delete.proto

package role

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

// DeleteI is the input for deleting roles.
//
//     {
//         "obj": [
//             {
//                 "metadata": {
//                     "resource.venturemark.co/kind": "venture",
//                     "role.venturemark.co/id": "<id>",
//                     "venture.venturemark.co/id": "<id>"
//                 }
//             }
//         ]
//     }
//
type DeleteI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *DeleteI_API   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Obj []*DeleteI_Obj `protobuf:"bytes,2,rep,name=obj,proto3" json:"obj,omitempty"`
}

func (x *DeleteI) Reset() {
	*x = DeleteI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_role_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI) ProtoMessage() {}

func (x *DeleteI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_role_delete_proto_msgTypes[0]
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
	return file_pbf_role_delete_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteI) GetApi() *DeleteI_API {
	if x != nil {
		return x.Api
	}
	return nil
}

func (x *DeleteI) GetObj() []*DeleteI_Obj {
	if x != nil {
		return x.Obj
	}
	return nil
}

type DeleteI_API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteI_API) Reset() {
	*x = DeleteI_API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_role_delete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI_API) ProtoMessage() {}

func (x *DeleteI_API) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_role_delete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI_API.ProtoReflect.Descriptor instead.
func (*DeleteI_API) Descriptor() ([]byte, []int) {
	return file_pbf_role_delete_proto_rawDescGZIP(), []int{1}
}

type DeleteI_Obj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeleteI_Obj) Reset() {
	*x = DeleteI_Obj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_role_delete_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI_Obj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI_Obj) ProtoMessage() {}

func (x *DeleteI_Obj) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_role_delete_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI_Obj.ProtoReflect.Descriptor instead.
func (*DeleteI_Obj) Descriptor() ([]byte, []int) {
	return file_pbf_role_delete_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteI_Obj) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// DeleteO is the output for deleting roles.
//
//     {
//         "obj": [
//             {
//                 "metadata": {
//                     "role.venturemark.co/status": "deleted"
//                 }
//             }
//         ]
//     }
//
type DeleteO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *DeleteO_API   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Obj []*DeleteO_Obj `protobuf:"bytes,2,rep,name=obj,proto3" json:"obj,omitempty"`
}

func (x *DeleteO) Reset() {
	*x = DeleteO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_role_delete_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO) ProtoMessage() {}

func (x *DeleteO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_role_delete_proto_msgTypes[3]
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
	return file_pbf_role_delete_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteO) GetApi() *DeleteO_API {
	if x != nil {
		return x.Api
	}
	return nil
}

func (x *DeleteO) GetObj() []*DeleteO_Obj {
	if x != nil {
		return x.Obj
	}
	return nil
}

type DeleteO_API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteO_API) Reset() {
	*x = DeleteO_API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_role_delete_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO_API) ProtoMessage() {}

func (x *DeleteO_API) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_role_delete_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO_API.ProtoReflect.Descriptor instead.
func (*DeleteO_API) Descriptor() ([]byte, []int) {
	return file_pbf_role_delete_proto_rawDescGZIP(), []int{4}
}

type DeleteO_Obj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeleteO_Obj) Reset() {
	*x = DeleteO_Obj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_role_delete_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO_Obj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO_Obj) ProtoMessage() {}

func (x *DeleteO_Obj) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_role_delete_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO_Obj.ProtoReflect.Descriptor instead.
func (*DeleteO_Obj) Descriptor() ([]byte, []int) {
	return file_pbf_role_delete_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteO_Obj) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_pbf_role_delete_proto protoreflect.FileDescriptor

var file_pbf_role_delete_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x66, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x53, 0x0a,
	0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x12, 0x23, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x5f, 0x41, 0x50, 0x49, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x23, 0x0a,
	0x03, 0x6f, 0x62, 0x6a, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x52, 0x03, 0x6f,
	0x62, 0x6a, 0x22, 0x0d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x41, 0x50,
	0x49, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62,
	0x6a, 0x12, 0x3b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x53, 0x0a, 0x07, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x12, 0x23, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x5f, 0x41, 0x50, 0x49, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x23, 0x0a, 0x03, 0x6f,
	0x62, 0x6a, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x52, 0x03, 0x6f, 0x62, 0x6a,
	0x22, 0x0d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x41, 0x50, 0x49, 0x22,
	0x87, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x12,
	0x3b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x5f, 0x4f, 0x62, 0x6a, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b,
	0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_role_delete_proto_rawDescOnce sync.Once
	file_pbf_role_delete_proto_rawDescData = file_pbf_role_delete_proto_rawDesc
)

func file_pbf_role_delete_proto_rawDescGZIP() []byte {
	file_pbf_role_delete_proto_rawDescOnce.Do(func() {
		file_pbf_role_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_role_delete_proto_rawDescData)
	})
	return file_pbf_role_delete_proto_rawDescData
}

var file_pbf_role_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pbf_role_delete_proto_goTypes = []interface{}{
	(*DeleteI)(nil),     // 0: role.DeleteI
	(*DeleteI_API)(nil), // 1: role.DeleteI_API
	(*DeleteI_Obj)(nil), // 2: role.DeleteI_Obj
	(*DeleteO)(nil),     // 3: role.DeleteO
	(*DeleteO_API)(nil), // 4: role.DeleteO_API
	(*DeleteO_Obj)(nil), // 5: role.DeleteO_Obj
	nil,                 // 6: role.DeleteI_Obj.MetadataEntry
	nil,                 // 7: role.DeleteO_Obj.MetadataEntry
}
var file_pbf_role_delete_proto_depIdxs = []int32{
	1, // 0: role.DeleteI.api:type_name -> role.DeleteI_API
	2, // 1: role.DeleteI.obj:type_name -> role.DeleteI_Obj
	6, // 2: role.DeleteI_Obj.metadata:type_name -> role.DeleteI_Obj.MetadataEntry
	4, // 3: role.DeleteO.api:type_name -> role.DeleteO_API
	5, // 4: role.DeleteO.obj:type_name -> role.DeleteO_Obj
	7, // 5: role.DeleteO_Obj.metadata:type_name -> role.DeleteO_Obj.MetadataEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pbf_role_delete_proto_init() }
func file_pbf_role_delete_proto_init() {
	if File_pbf_role_delete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_role_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_role_delete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI_API); i {
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
		file_pbf_role_delete_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI_Obj); i {
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
		file_pbf_role_delete_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_role_delete_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO_API); i {
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
		file_pbf_role_delete_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO_Obj); i {
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
			RawDescriptor: file_pbf_role_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_role_delete_proto_goTypes,
		DependencyIndexes: file_pbf_role_delete_proto_depIdxs,
		MessageInfos:      file_pbf_role_delete_proto_msgTypes,
	}.Build()
	File_pbf_role_delete_proto = out.File
	file_pbf_role_delete_proto_rawDesc = nil
	file_pbf_role_delete_proto_goTypes = nil
	file_pbf_role_delete_proto_depIdxs = nil
}
