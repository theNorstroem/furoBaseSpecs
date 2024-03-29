// Code generated by furo-proto-gen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: furo/property.proto

package propertypb

import (
	furo "github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Integer type with embedded meta
type IntegerProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Integer data part
	Data int32 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *IntegerProperty) Reset() {
	*x = IntegerProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_property_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegerProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegerProperty) ProtoMessage() {}

func (x *IntegerProperty) ProtoReflect() protoreflect.Message {
	mi := &file_furo_property_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegerProperty.ProtoReflect.Descriptor instead.
func (*IntegerProperty) Descriptor() ([]byte, []int) {
	return file_furo_property_proto_rawDescGZIP(), []int{0}
}

func (x *IntegerProperty) GetData() int32 {
	if x != nil {
		return x.Data
	}
	return 0
}

// Number type with embedded meta
type NumberProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data part
	Data float32 `protobuf:"fixed32,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *NumberProperty) Reset() {
	*x = NumberProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_property_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberProperty) ProtoMessage() {}

func (x *NumberProperty) ProtoReflect() protoreflect.Message {
	mi := &file_furo_property_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberProperty.ProtoReflect.Descriptor instead.
func (*NumberProperty) Descriptor() ([]byte, []int) {
	return file_furo_property_proto_rawDescGZIP(), []int{1}
}

func (x *NumberProperty) GetData() float32 {
	if x != nil {
		return x.Data
	}
	return 0
}

// Type to define property values with type information
type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the property
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// String representation of the property
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// data part of the property
	Data *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// Meta for the response
	Meta *furo.Meta `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	// property code for additional settings
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	// Optional attribute flags e.g. is-overwritable
	Flags []string `protobuf:"bytes,6,rep,name=flags,proto3" json:"flags,omitempty"`
	// Optional flag indicating that the property differs from the original value
	IsOverwritten bool `protobuf:"varint,7,opt,name=is_overwritten,json=isOverwritten,proto3" json:"is_overwritten,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_property_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_furo_property_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_furo_property_proto_rawDescGZIP(), []int{2}
}

func (x *Property) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Property) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Property) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Property) GetMeta() *furo.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Property) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Property) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *Property) GetIsOverwritten() bool {
	if x != nil {
		return x.IsOverwritten
	}
	return false
}

// String type to use in property
type StringOptionProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String representation of val
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The value, Id is used to make working with data-inputs easier
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StringOptionProperty) Reset() {
	*x = StringOptionProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_property_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringOptionProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringOptionProperty) ProtoMessage() {}

func (x *StringOptionProperty) ProtoReflect() protoreflect.Message {
	mi := &file_furo_property_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringOptionProperty.ProtoReflect.Descriptor instead.
func (*StringOptionProperty) Descriptor() ([]byte, []int) {
	return file_furo_property_proto_rawDescGZIP(), []int{3}
}

func (x *StringOptionProperty) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *StringOptionProperty) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// String type to use in property
type StringProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data part
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StringProperty) Reset() {
	*x = StringProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_property_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringProperty) ProtoMessage() {}

func (x *StringProperty) ProtoReflect() protoreflect.Message {
	mi := &file_furo_property_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringProperty.ProtoReflect.Descriptor instead.
func (*StringProperty) Descriptor() ([]byte, []int) {
	return file_furo_property_proto_rawDescGZIP(), []int{4}
}

func (x *StringProperty) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_furo_property_proto protoreflect.FileDescriptor

var file_furo_property_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x75, 0x72, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x75, 0x72, 0x6f, 0x1a, 0x0f, 0x66, 0x75, 0x72,
	0x6f, 0x2f, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24,
	0x0a, 0x0e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xd8, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66,
	0x75, 0x72, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x6f,
	0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x69, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x22,
	0x49, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x82, 0x01, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x42, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x4e, 0x6f, 0x72, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x6d,
	0x2f, 0x46, 0x75, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x75, 0x72, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x3b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x70, 0x62, 0xa2,
	0x02, 0x03, 0x46, 0x50, 0x42, 0xaa, 0x02, 0x0d, 0x46, 0x75, 0x72, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_furo_property_proto_rawDescOnce sync.Once
	file_furo_property_proto_rawDescData = file_furo_property_proto_rawDesc
)

func file_furo_property_proto_rawDescGZIP() []byte {
	file_furo_property_proto_rawDescOnce.Do(func() {
		file_furo_property_proto_rawDescData = protoimpl.X.CompressGZIP(file_furo_property_proto_rawDescData)
	})
	return file_furo_property_proto_rawDescData
}

var file_furo_property_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_furo_property_proto_goTypes = []interface{}{
	(*IntegerProperty)(nil),      // 0: furo.IntegerProperty
	(*NumberProperty)(nil),       // 1: furo.NumberProperty
	(*Property)(nil),             // 2: furo.Property
	(*StringOptionProperty)(nil), // 3: furo.StringOptionProperty
	(*StringProperty)(nil),       // 4: furo.StringProperty
	(*anypb.Any)(nil),            // 5: google.protobuf.Any
	(*furo.Meta)(nil),            // 6: furo.Meta
}
var file_furo_property_proto_depIdxs = []int32{
	5, // 0: furo.Property.data:type_name -> google.protobuf.Any
	6, // 1: furo.Property.meta:type_name -> furo.Meta
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_furo_property_proto_init() }
func file_furo_property_proto_init() {
	if File_furo_property_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_furo_property_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegerProperty); i {
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
		file_furo_property_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberProperty); i {
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
		file_furo_property_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_furo_property_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringOptionProperty); i {
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
		file_furo_property_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringProperty); i {
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
			RawDescriptor: file_furo_property_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_furo_property_proto_goTypes,
		DependencyIndexes: file_furo_property_proto_depIdxs,
		MessageInfos:      file_furo_property_proto_msgTypes,
	}.Build()
	File_furo_property_proto = out.File
	file_furo_property_proto_rawDesc = nil
	file_furo_property_proto_goTypes = nil
	file_furo_property_proto_depIdxs = nil
}
