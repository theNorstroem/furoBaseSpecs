// Code generated by furo-proto-gen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: furo/furo.proto

package furopb

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// a single fieldconstraint
type FieldConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the constraint value as string, even it is a number
	Is string `protobuf:"bytes,1,opt,name=is,proto3" json:"is,omitempty"`
	// The message to display on constraint violation
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FieldConstraint) Reset() {
	*x = FieldConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_furo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldConstraint) ProtoMessage() {}

func (x *FieldConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_furo_furo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldConstraint.ProtoReflect.Descriptor instead.
func (*FieldConstraint) Descriptor() ([]byte, []int) {
	return file_furo_furo_proto_rawDescGZIP(), []int{0}
}

func (x *FieldConstraint) GetIs() string {
	if x != nil {
		return x.Is
	}
	return ""
}

func (x *FieldConstraint) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Metas for a field
type FieldMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The label
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// A hint
	Hint string `protobuf:"bytes,2,opt,name=hint,proto3" json:"hint,omitempty"`
	// The default value as JSON string
	Default string `protobuf:"bytes,3,opt,name=default,proto3" json:"default,omitempty"`
	// readonly
	Readonly bool `protobuf:"varint,4,opt,name=readonly,proto3" json:"readonly,omitempty"`
	// repeated
	Repeated bool `protobuf:"varint,5,opt,name=repeated,proto3" json:"repeated,omitempty"`
	// Fieldoptions
	Options *Fieldoption `protobuf:"bytes,6,opt,name=options,proto3" json:"options,omitempty"`
	// Put in type specific metas for your fields here
	Typespecific *any.Any `protobuf:"bytes,7,opt,name=typespecific,proto3" json:"typespecific,omitempty"`
}

func (x *FieldMeta) Reset() {
	*x = FieldMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_furo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMeta) ProtoMessage() {}

func (x *FieldMeta) ProtoReflect() protoreflect.Message {
	mi := &file_furo_furo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMeta.ProtoReflect.Descriptor instead.
func (*FieldMeta) Descriptor() ([]byte, []int) {
	return file_furo_furo_proto_rawDescGZIP(), []int{1}
}

func (x *FieldMeta) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *FieldMeta) GetHint() string {
	if x != nil {
		return x.Hint
	}
	return ""
}

func (x *FieldMeta) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *FieldMeta) GetReadonly() bool {
	if x != nil {
		return x.Readonly
	}
	return false
}

func (x *FieldMeta) GetRepeated() bool {
	if x != nil {
		return x.Repeated
	}
	return false
}

func (x *FieldMeta) GetOptions() *Fieldoption {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *FieldMeta) GetTypespecific() *any.Any {
	if x != nil {
		return x.Typespecific
	}
	return nil
}

// Metas for a field
type Fieldoption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a list with options, use furo.optionitem or your own
	List []*any.Any `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	// Add flags for your field. This can be something like "searchable".
	// //The flags can be used by generators, ui components,...
	//
	Flags []string `protobuf:"bytes,2,rep,name=flags,proto3" json:"flags,omitempty"`
}

func (x *Fieldoption) Reset() {
	*x = Fieldoption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_furo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fieldoption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fieldoption) ProtoMessage() {}

func (x *Fieldoption) ProtoReflect() protoreflect.Message {
	mi := &file_furo_furo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fieldoption.ProtoReflect.Descriptor instead.
func (*Fieldoption) Descriptor() ([]byte, []int) {
	return file_furo_furo_proto_rawDescGZIP(), []int{2}
}

func (x *Fieldoption) GetList() []*any.Any {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *Fieldoption) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

// link
type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the relationship like self...
	Rel string `protobuf:"bytes,1,opt,name=rel,proto3" json:"rel,omitempty"`
	// method of curl v1.0.0
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// link
	Href string `protobuf:"bytes,3,opt,name=href,proto3" json:"href,omitempty"`
	// mime type
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// name of the service which can handle this link
	Service string `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_furo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_furo_furo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_furo_furo_proto_rawDescGZIP(), []int{3}
}

func (x *Link) GetRel() string {
	if x != nil {
		return x.Rel
	}
	return ""
}

func (x *Link) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Link) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

func (x *Link) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Link) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

// meta info
type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fields of meta info
	Fields map[string]*MetaField `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_furo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_furo_furo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_furo_furo_proto_rawDescGZIP(), []int{4}
}

func (x *Meta) GetFields() map[string]*MetaField {
	if x != nil {
		return x.Fields
	}
	return nil
}

// fields of meta info
type MetaField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// meta informatioxn of a field
	Meta *FieldMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// constraints for a field
	Constraints map[string]*FieldConstraint `protobuf:"bytes,2,rep,name=constraints,proto3" json:"constraints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetaField) Reset() {
	*x = MetaField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_furo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaField) ProtoMessage() {}

func (x *MetaField) ProtoReflect() protoreflect.Message {
	mi := &file_furo_furo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaField.ProtoReflect.Descriptor instead.
func (*MetaField) Descriptor() ([]byte, []int) {
	return file_furo_furo_proto_rawDescGZIP(), []int{5}
}

func (x *MetaField) GetMeta() *FieldMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *MetaField) GetConstraints() map[string]*FieldConstraint {
	if x != nil {
		return x.Constraints
	}
	return nil
}

// Items for fieldoption.list
type Optionitem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// String representation
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// is the item selected
	Selected bool `protobuf:"varint,3,opt,name=selected,proto3" json:"selected,omitempty"`
}

func (x *Optionitem) Reset() {
	*x = Optionitem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_furo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Optionitem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Optionitem) ProtoMessage() {}

func (x *Optionitem) ProtoReflect() protoreflect.Message {
	mi := &file_furo_furo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Optionitem.ProtoReflect.Descriptor instead.
func (*Optionitem) Descriptor() ([]byte, []int) {
	return file_furo_furo_proto_rawDescGZIP(), []int{6}
}

func (x *Optionitem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Optionitem) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Optionitem) GetSelected() bool {
	if x != nil {
		return x.Selected
	}
	return false
}

// reference
type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String representation of the reference
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Id of the reference
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Hateoas link
	Link *Link `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_furo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_furo_furo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_furo_furo_proto_rawDescGZIP(), []int{7}
}

func (x *Reference) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Reference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reference) GetLink() *Link {
	if x != nil {
		return x.Link
	}
	return nil
}

var File_furo_furo_proto protoreflect.FileDescriptor

var file_furo_furo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x75, 0x72, 0x6f, 0x2f, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x66, 0x75, 0x72, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xee, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x75, 0x72,
	0x6f, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x22, 0x4d, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x22,
	0x72, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x72, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66,
	0x75, 0x72, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x4a, 0x0a, 0x0b,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66,
	0x75, 0x72, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcb, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x74,
	0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x1a,
	0x55, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b, 0x0a, 0x0a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x69, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x22, 0x5e, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x42, 0x62, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x42,
	0x09, 0x46, 0x75, 0x72, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x4e, 0x6f, 0x72, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x6d, 0x2f, 0x46, 0x75, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x75, 0x72, 0x6f,
	0x3b, 0x66, 0x75, 0x72, 0x6f, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x46, 0x50, 0x42,
	0xaa, 0x02, 0x04, 0x46, 0x75, 0x72, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_furo_furo_proto_rawDescOnce sync.Once
	file_furo_furo_proto_rawDescData = file_furo_furo_proto_rawDesc
)

func file_furo_furo_proto_rawDescGZIP() []byte {
	file_furo_furo_proto_rawDescOnce.Do(func() {
		file_furo_furo_proto_rawDescData = protoimpl.X.CompressGZIP(file_furo_furo_proto_rawDescData)
	})
	return file_furo_furo_proto_rawDescData
}

var file_furo_furo_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_furo_furo_proto_goTypes = []interface{}{
	(*FieldConstraint)(nil), // 0: furo.FieldConstraint
	(*FieldMeta)(nil),       // 1: furo.FieldMeta
	(*Fieldoption)(nil),     // 2: furo.Fieldoption
	(*Link)(nil),            // 3: furo.Link
	(*Meta)(nil),            // 4: furo.Meta
	(*MetaField)(nil),       // 5: furo.MetaField
	(*Optionitem)(nil),      // 6: furo.Optionitem
	(*Reference)(nil),       // 7: furo.Reference
	nil,                     // 8: furo.Meta.FieldsEntry
	nil,                     // 9: furo.MetaField.ConstraintsEntry
	(*any.Any)(nil),         // 10: google.protobuf.Any
}
var file_furo_furo_proto_depIdxs = []int32{
	2,  // 0: furo.FieldMeta.options:type_name -> furo.Fieldoption
	10, // 1: furo.FieldMeta.typespecific:type_name -> google.protobuf.Any
	10, // 2: furo.Fieldoption.list:type_name -> google.protobuf.Any
	8,  // 3: furo.Meta.fields:type_name -> furo.Meta.FieldsEntry
	1,  // 4: furo.MetaField.meta:type_name -> furo.FieldMeta
	9,  // 5: furo.MetaField.constraints:type_name -> furo.MetaField.ConstraintsEntry
	3,  // 6: furo.Reference.link:type_name -> furo.Link
	5,  // 7: furo.Meta.FieldsEntry.value:type_name -> furo.MetaField
	0,  // 8: furo.MetaField.ConstraintsEntry.value:type_name -> furo.FieldConstraint
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_furo_furo_proto_init() }
func file_furo_furo_proto_init() {
	if File_furo_furo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_furo_furo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldConstraint); i {
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
		file_furo_furo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldMeta); i {
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
		file_furo_furo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fieldoption); i {
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
		file_furo_furo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
		file_furo_furo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_furo_furo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaField); i {
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
		file_furo_furo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Optionitem); i {
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
		file_furo_furo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
			RawDescriptor: file_furo_furo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_furo_furo_proto_goTypes,
		DependencyIndexes: file_furo_furo_proto_depIdxs,
		MessageInfos:      file_furo_furo_proto_msgTypes,
	}.Build()
	File_furo_furo_proto = out.File
	file_furo_furo_proto_rawDesc = nil
	file_furo_furo_proto_goTypes = nil
	file_furo_furo_proto_depIdxs = nil
}
