// Code generated by furo-proto-gen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: furo/type/type.proto

package furotypepb

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

// Represents a whole calendar date, e.g. date of birth. The time of day and
//  time zone are either specified elsewhere or are not significant. The date
//  is relative to the Proleptic Gregorian Calendar. The day may be 0 to
//  represent a year and month where the day is not significant, e.g. credit card
//  expiration date. The year may be 0 to represent a month and day independent
//  of year, e.g. anniversary date. Related types are [google.type.TimeOfDay][google.type.TimeOfDay]
//  and `google.protobuf.Timestamp`.
type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Year of date. Must be from 1 to 9999, or 0 if specifying a date without
	//  a year.
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Month of year. Must be from 1 to 12.
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Day of month. Must be from 1 to 31 and valid for the year and month, or 0
	//  if specifying a year/month where the day is not significant.
	Day int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	// String representation of the reference
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_type_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_furo_type_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_furo_type_type_proto_rawDescGZIP(), []int{0}
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *Date) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// Represents an amount of money with its currency type.
type Money struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The whole units of the amount.
	//  For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	Units int64 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// Number of nano (10^-9) units of the amount.
	//  The value must be between -999,999,999 and +999,999,999 inclusive.
	//  If `units` is positive, `nanos` must be positive or zero.
	//  If `units` is zero, `nanos` can be positive, zero, or negative.
	//  If `units` is negative, `nanos` must be negative or zero.
	//  For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos int32 `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
	// String representation of the reference
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *Money) Reset() {
	*x = Money{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furo_type_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Money) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Money) ProtoMessage() {}

func (x *Money) ProtoReflect() protoreflect.Message {
	mi := &file_furo_type_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Money.ProtoReflect.Descriptor instead.
func (*Money) Descriptor() ([]byte, []int) {
	return file_furo_type_type_proto_rawDescGZIP(), []int{1}
}

func (x *Money) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Money) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Money) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

func (x *Money) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_furo_type_type_proto protoreflect.FileDescriptor

var file_furo_type_type_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x75, 0x72, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x65, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7b, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e,
	0x6f, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x69, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x2e, 0x66, 0x75, 0x72,
	0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x42, 0x09, 0x44, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x68, 0x65, 0x4e, 0x6f, 0x72, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x6d, 0x2f, 0x46, 0x75, 0x72,
	0x6f, 0x42, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f,
	0x70, 0x62, 0x2f, 0x66, 0x75, 0x72, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x3b, 0x66, 0x75, 0x72,
	0x6f, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x46, 0x50, 0x42,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_furo_type_type_proto_rawDescOnce sync.Once
	file_furo_type_type_proto_rawDescData = file_furo_type_type_proto_rawDesc
)

func file_furo_type_type_proto_rawDescGZIP() []byte {
	file_furo_type_type_proto_rawDescOnce.Do(func() {
		file_furo_type_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_furo_type_type_proto_rawDescData)
	})
	return file_furo_type_type_proto_rawDescData
}

var file_furo_type_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_furo_type_type_proto_goTypes = []interface{}{
	(*Date)(nil),  // 0: furo.type.Date
	(*Money)(nil), // 1: furo.type.Money
}
var file_furo_type_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_furo_type_type_proto_init() }
func file_furo_type_type_proto_init() {
	if File_furo_type_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_furo_type_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
		file_furo_type_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Money); i {
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
			RawDescriptor: file_furo_type_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_furo_type_type_proto_goTypes,
		DependencyIndexes: file_furo_type_type_proto_depIdxs,
		MessageInfos:      file_furo_type_type_proto_msgTypes,
	}.Build()
	File_furo_type_type_proto = out.File
	file_furo_type_type_proto_rawDesc = nil
	file_furo_type_type_proto_goTypes = nil
	file_furo_type_type_proto_depIdxs = nil
}
