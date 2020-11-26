// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: vector_tile.proto

package mvt

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
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

// GeomType is described in section 4.3.4 of the specification
type Tile_GeomType int32

const (
	Tile_UNKNOWN    Tile_GeomType = 0
	Tile_POINT      Tile_GeomType = 1
	Tile_LINESTRING Tile_GeomType = 2
	Tile_POLYGON    Tile_GeomType = 3
)

// Enum value maps for Tile_GeomType.
var (
	Tile_GeomType_name = map[int32]string{
		0: "UNKNOWN",
		1: "POINT",
		2: "LINESTRING",
		3: "POLYGON",
	}
	Tile_GeomType_value = map[string]int32{
		"UNKNOWN":    0,
		"POINT":      1,
		"LINESTRING": 2,
		"POLYGON":    3,
	}
)

func (x Tile_GeomType) Enum() *Tile_GeomType {
	p := new(Tile_GeomType)
	*p = x
	return p
}

func (x Tile_GeomType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tile_GeomType) Descriptor() protoreflect.EnumDescriptor {
	return file_vector_tile_proto_enumTypes[0].Descriptor()
}

func (Tile_GeomType) Type() protoreflect.EnumType {
	return &file_vector_tile_proto_enumTypes[0]
}

func (x Tile_GeomType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Tile_GeomType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Tile_GeomType(num)
	return nil
}

// Deprecated: Use Tile_GeomType.Descriptor instead.
func (Tile_GeomType) EnumDescriptor() ([]byte, []int) {
	return file_vector_tile_proto_rawDescGZIP(), []int{0, 0}
}

type Tile struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	Layers []*Tile_Layer `protobuf:"bytes,3,rep,name=layers" json:"layers,omitempty"`
}

func (x *Tile) Reset() {
	*x = Tile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vector_tile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tile) ProtoMessage() {}

func (x *Tile) ProtoReflect() protoreflect.Message {
	mi := &file_vector_tile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tile.ProtoReflect.Descriptor instead.
func (*Tile) Descriptor() ([]byte, []int) {
	return file_vector_tile_proto_rawDescGZIP(), []int{0}
}

var extRange_Tile = []protoiface.ExtensionRangeV1{
	{Start: 16, End: 8191},
}

// Deprecated: Use Tile.ProtoReflect.Descriptor.ExtensionRanges instead.
func (*Tile) ExtensionRangeArray() []protoiface.ExtensionRangeV1 {
	return extRange_Tile
}

func (x *Tile) GetLayers() []*Tile_Layer {
	if x != nil {
		return x.Layers
	}
	return nil
}

// Variant type encoding
// The use of values is described in section 4.1 of the specification
type Tile_Value struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	// Exactly one of these values must be present in a valid message
	StringValue *string  `protobuf:"bytes,1,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	FloatValue  *float32 `protobuf:"fixed32,2,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	DoubleValue *float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
	IntValue    *int64   `protobuf:"varint,4,opt,name=int_value,json=intValue" json:"int_value,omitempty"`
	UintValue   *uint64  `protobuf:"varint,5,opt,name=uint_value,json=uintValue" json:"uint_value,omitempty"`
	SintValue   *int64   `protobuf:"zigzag64,6,opt,name=sint_value,json=sintValue" json:"sint_value,omitempty"`
	BoolValue   *bool    `protobuf:"varint,7,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
}

func (x *Tile_Value) Reset() {
	*x = Tile_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vector_tile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tile_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tile_Value) ProtoMessage() {}

func (x *Tile_Value) ProtoReflect() protoreflect.Message {
	mi := &file_vector_tile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tile_Value.ProtoReflect.Descriptor instead.
func (*Tile_Value) Descriptor() ([]byte, []int) {
	return file_vector_tile_proto_rawDescGZIP(), []int{0, 0}
}

var extRange_Tile_Value = []protoiface.ExtensionRangeV1{
	{Start: 8, End: 536870911},
}

// Deprecated: Use Tile_Value.ProtoReflect.Descriptor.ExtensionRanges instead.
func (*Tile_Value) ExtensionRangeArray() []protoiface.ExtensionRangeV1 {
	return extRange_Tile_Value
}

func (x *Tile_Value) GetStringValue() string {
	if x != nil && x.StringValue != nil {
		return *x.StringValue
	}
	return ""
}

func (x *Tile_Value) GetFloatValue() float32 {
	if x != nil && x.FloatValue != nil {
		return *x.FloatValue
	}
	return 0
}

func (x *Tile_Value) GetDoubleValue() float64 {
	if x != nil && x.DoubleValue != nil {
		return *x.DoubleValue
	}
	return 0
}

func (x *Tile_Value) GetIntValue() int64 {
	if x != nil && x.IntValue != nil {
		return *x.IntValue
	}
	return 0
}

func (x *Tile_Value) GetUintValue() uint64 {
	if x != nil && x.UintValue != nil {
		return *x.UintValue
	}
	return 0
}

func (x *Tile_Value) GetSintValue() int64 {
	if x != nil && x.SintValue != nil {
		return *x.SintValue
	}
	return 0
}

func (x *Tile_Value) GetBoolValue() bool {
	if x != nil && x.BoolValue != nil {
		return *x.BoolValue
	}
	return false
}

// Features are described in section 4.2 of the specification
type Tile_Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *uint64 `protobuf:"varint,1,opt,name=id,def=0" json:"id,omitempty"`
	// Tags of this feature are encoded as repeated pairs of
	// integers.
	// A detailed description of tags is located in sections
	// 4.2 and 4.4 of the specification
	Tags []uint32 `protobuf:"varint,2,rep,packed,name=tags" json:"tags,omitempty"`
	// The type of geometry stored in this feature.
	Type *Tile_GeomType `protobuf:"varint,3,opt,name=type,enum=mvt.Tile_GeomType,def=0" json:"type,omitempty"`
	// Contains a stream of commands and parameters (vertices).
	// A detailed description on geometry encoding is located in
	// section 4.3 of the specification.
	Geometry []uint32 `protobuf:"varint,4,rep,packed,name=geometry" json:"geometry,omitempty"`
}

// Default values for Tile_Feature fields.
const (
	Default_Tile_Feature_Id   = uint64(0)
	Default_Tile_Feature_Type = Tile_UNKNOWN
)

func (x *Tile_Feature) Reset() {
	*x = Tile_Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vector_tile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tile_Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tile_Feature) ProtoMessage() {}

func (x *Tile_Feature) ProtoReflect() protoreflect.Message {
	mi := &file_vector_tile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tile_Feature.ProtoReflect.Descriptor instead.
func (*Tile_Feature) Descriptor() ([]byte, []int) {
	return file_vector_tile_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Tile_Feature) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return Default_Tile_Feature_Id
}

func (x *Tile_Feature) GetTags() []uint32 {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Tile_Feature) GetType() Tile_GeomType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_Tile_Feature_Type
}

func (x *Tile_Feature) GetGeometry() []uint32 {
	if x != nil {
		return x.Geometry
	}
	return nil
}

// Layers are described in section 4.1 of the specification
type Tile_Layer struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	// Any compliant implementation must first read the version
	// number encoded in this message and choose the correct
	// implementation for this version number before proceeding to
	// decode other parts of this message.
	Version *uint32 `protobuf:"varint,15,req,name=version,def=1" json:"version,omitempty"`
	Name    *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// The actual features in this tile.
	Features []*Tile_Feature `protobuf:"bytes,2,rep,name=features" json:"features,omitempty"`
	// Dictionary encoding for keys
	Keys []string `protobuf:"bytes,3,rep,name=keys" json:"keys,omitempty"`
	// Dictionary encoding for values
	Values []*Tile_Value `protobuf:"bytes,4,rep,name=values" json:"values,omitempty"`
	// Although this is an "optional" field it is required by the specification.
	// See https://github.com/mapbox/vector-tile-spec/issues/47
	Extent *uint32 `protobuf:"varint,5,opt,name=extent,def=4096" json:"extent,omitempty"`
}

// Default values for Tile_Layer fields.
const (
	Default_Tile_Layer_Version = uint32(1)
	Default_Tile_Layer_Extent  = uint32(4096)
)

func (x *Tile_Layer) Reset() {
	*x = Tile_Layer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vector_tile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tile_Layer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tile_Layer) ProtoMessage() {}

func (x *Tile_Layer) ProtoReflect() protoreflect.Message {
	mi := &file_vector_tile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tile_Layer.ProtoReflect.Descriptor instead.
func (*Tile_Layer) Descriptor() ([]byte, []int) {
	return file_vector_tile_proto_rawDescGZIP(), []int{0, 2}
}

var extRange_Tile_Layer = []protoiface.ExtensionRangeV1{
	{Start: 16, End: 536870911},
}

// Deprecated: Use Tile_Layer.ProtoReflect.Descriptor.ExtensionRanges instead.
func (*Tile_Layer) ExtensionRangeArray() []protoiface.ExtensionRangeV1 {
	return extRange_Tile_Layer
}

func (x *Tile_Layer) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return Default_Tile_Layer_Version
}

func (x *Tile_Layer) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Tile_Layer) GetFeatures() []*Tile_Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *Tile_Layer) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Tile_Layer) GetValues() []*Tile_Value {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Tile_Layer) GetExtent() uint32 {
	if x != nil && x.Extent != nil {
		return *x.Extent
	}
	return Default_Tile_Layer_Extent
}

var File_vector_tile_proto protoreflect.FileDescriptor

var file_vector_tile_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x76, 0x74, 0x22, 0xc3, 0x05, 0x0a, 0x04, 0x54, 0x69, 0x6c,
	0x65, 0x12, 0x27, 0x0a, 0x06, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x76, 0x74, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x2e, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x06, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x1a, 0xf2, 0x01, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x69,
	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x6e, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x73, 0x69, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x08, 0x08, 0x08, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x1a,
	0x85, 0x01, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x11, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x3a, 0x01, 0x30, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x76, 0x74, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x2e,
	0x47, 0x65, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x08, 0x67,
	0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x1a, 0xcc, 0x01, 0x0a, 0x05, 0x4c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x02,
	0x28, 0x0d, 0x3a, 0x01, 0x31, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x76, 0x74, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x76, 0x74, 0x2e, 0x54, 0x69, 0x6c, 0x65,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1c,
	0x0a, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x04,
	0x34, 0x30, 0x39, 0x36, 0x52, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x08, 0x08, 0x10,
	0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x22, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x49,
	0x4e, 0x45, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x4f,
	0x4c, 0x59, 0x47, 0x4f, 0x4e, 0x10, 0x03, 0x2a, 0x05, 0x08, 0x10, 0x10, 0x80, 0x40, 0x42, 0x02,
	0x48, 0x03,
}

var (
	file_vector_tile_proto_rawDescOnce sync.Once
	file_vector_tile_proto_rawDescData = file_vector_tile_proto_rawDesc
)

func file_vector_tile_proto_rawDescGZIP() []byte {
	file_vector_tile_proto_rawDescOnce.Do(func() {
		file_vector_tile_proto_rawDescData = protoimpl.X.CompressGZIP(file_vector_tile_proto_rawDescData)
	})
	return file_vector_tile_proto_rawDescData
}

var file_vector_tile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vector_tile_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_vector_tile_proto_goTypes = []interface{}{
	(Tile_GeomType)(0),   // 0: mvt.Tile.GeomType
	(*Tile)(nil),         // 1: mvt.Tile
	(*Tile_Value)(nil),   // 2: mvt.Tile.Value
	(*Tile_Feature)(nil), // 3: mvt.Tile.Feature
	(*Tile_Layer)(nil),   // 4: mvt.Tile.Layer
}
var file_vector_tile_proto_depIdxs = []int32{
	4, // 0: mvt.Tile.layers:type_name -> mvt.Tile.Layer
	0, // 1: mvt.Tile.Feature.type:type_name -> mvt.Tile.GeomType
	3, // 2: mvt.Tile.Layer.features:type_name -> mvt.Tile.Feature
	2, // 3: mvt.Tile.Layer.values:type_name -> mvt.Tile.Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_vector_tile_proto_init() }
func file_vector_tile_proto_init() {
	if File_vector_tile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vector_tile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
		file_vector_tile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tile_Value); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
		file_vector_tile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tile_Feature); i {
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
		file_vector_tile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tile_Layer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vector_tile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vector_tile_proto_goTypes,
		DependencyIndexes: file_vector_tile_proto_depIdxs,
		EnumInfos:         file_vector_tile_proto_enumTypes,
		MessageInfos:      file_vector_tile_proto_msgTypes,
	}.Build()
	File_vector_tile_proto = out.File
	file_vector_tile_proto_rawDesc = nil
	file_vector_tile_proto_goTypes = nil
	file_vector_tile_proto_depIdxs = nil
}
