// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_handle.proto

package tensorflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandleProto struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode,proto3" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName        string   `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName,proto3" json:"maybe_type_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceHandleProto) Reset()         { *m = ResourceHandleProto{} }
func (m *ResourceHandleProto) String() string { return proto.CompactTextString(m) }
func (*ResourceHandleProto) ProtoMessage()    {}
func (*ResourceHandleProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ca6d7983246ca36, []int{0}
}

func (m *ResourceHandleProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceHandleProto.Unmarshal(m, b)
}
func (m *ResourceHandleProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceHandleProto.Marshal(b, m, deterministic)
}
func (m *ResourceHandleProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceHandleProto.Merge(m, src)
}
func (m *ResourceHandleProto) XXX_Size() int {
	return xxx_messageInfo_ResourceHandleProto.Size(m)
}
func (m *ResourceHandleProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceHandleProto.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceHandleProto proto.InternalMessageInfo

func (m *ResourceHandleProto) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandleProto) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandleProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandleProto) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandleProto) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceHandleProto)(nil), "tensorflow.ResourceHandleProto")
}

func init() { proto.RegisterFile("resource_handle.proto", fileDescriptor_2ca6d7983246ca36) }

var fileDescriptor_2ca6d7983246ca36 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xc1, 0xaa, 0xc2, 0x30,
	0x10, 0x45, 0xc9, 0x7b, 0x7d, 0xe5, 0x75, 0x40, 0x84, 0x88, 0x12, 0xd0, 0x45, 0x71, 0x21, 0x5d,
	0xb9, 0xf1, 0x13, 0xdc, 0xb8, 0x12, 0x29, 0xee, 0x43, 0x9a, 0x8c, 0xb4, 0xd0, 0x66, 0x4a, 0x1a,
	0x95, 0xfe, 0x8f, 0x1f, 0x2a, 0x0e, 0x05, 0x77, 0x33, 0xe7, 0x9e, 0x0b, 0x17, 0x96, 0x01, 0x07,
	0xba, 0x07, 0x8b, 0xba, 0x36, 0xde, 0xb5, 0xb8, 0xef, 0x03, 0x45, 0x92, 0x10, 0xd1, 0x0f, 0x14,
	0x6e, 0x2d, 0x3d, 0xb7, 0x2f, 0x01, 0x8b, 0x72, 0xb2, 0x4e, 0x2c, 0x5d, 0xd8, 0x59, 0x41, 0xea,
	0xf0, 0xd1, 0x58, 0x54, 0x22, 0x17, 0x45, 0x56, 0x4e, 0x9f, 0xdc, 0x40, 0x66, 0xc9, 0x47, 0xd3,
	0x78, 0x0c, 0xea, 0x87, 0xa3, 0x2f, 0x90, 0x12, 0x12, 0x6f, 0x3a, 0x54, 0xbf, 0x1c, 0xf0, 0x2d,
	0xd7, 0x90, 0xd5, 0x66, 0xa8, 0xb5, 0x25, 0x87, 0x2a, 0xc9, 0x45, 0x91, 0x94, 0xff, 0x1f, 0x70,
	0x24, 0x87, 0x72, 0x07, 0xf3, 0xce, 0x8c, 0x15, 0xea, 0x38, 0xf6, 0xa8, 0xb9, 0xfb, 0xc7, 0xdd,
	0x19, 0xe3, 0xeb, 0xd8, 0xe3, 0xd9, 0x74, 0x58, 0xa5, 0xbc, 0xfc, 0xf0, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x0a, 0x22, 0x89, 0x70, 0xd2, 0x00, 0x00, 0x00,
}
