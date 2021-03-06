// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow.serving/proto/classification.proto

package tensorflow_serving

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// A single class.
type Class struct {
	// Label or name of the class.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// Score for this class (e.g., the probability the item belongs to this
	// class). As per the proto3 default-value semantics, if the score is missing,
	// it should be treated as 0.
	Score                float32  `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Class) Reset()         { *m = Class{} }
func (m *Class) String() string { return proto.CompactTextString(m) }
func (*Class) ProtoMessage()    {}
func (*Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_325571a8b7f54d57, []int{0}
}

func (m *Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Class.Unmarshal(m, b)
}
func (m *Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Class.Marshal(b, m, deterministic)
}
func (m *Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Class.Merge(m, src)
}
func (m *Class) XXX_Size() int {
	return xxx_messageInfo_Class.Size(m)
}
func (m *Class) XXX_DiscardUnknown() {
	xxx_messageInfo_Class.DiscardUnknown(m)
}

var xxx_messageInfo_Class proto.InternalMessageInfo

func (m *Class) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Class) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// List of classes for a single item (tensorflow.Example).
type Classifications struct {
	Classes              []*Class `protobuf:"bytes,1,rep,name=classes,proto3" json:"classes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Classifications) Reset()         { *m = Classifications{} }
func (m *Classifications) String() string { return proto.CompactTextString(m) }
func (*Classifications) ProtoMessage()    {}
func (*Classifications) Descriptor() ([]byte, []int) {
	return fileDescriptor_325571a8b7f54d57, []int{1}
}

func (m *Classifications) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Classifications.Unmarshal(m, b)
}
func (m *Classifications) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Classifications.Marshal(b, m, deterministic)
}
func (m *Classifications) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Classifications.Merge(m, src)
}
func (m *Classifications) XXX_Size() int {
	return xxx_messageInfo_Classifications.Size(m)
}
func (m *Classifications) XXX_DiscardUnknown() {
	xxx_messageInfo_Classifications.DiscardUnknown(m)
}

var xxx_messageInfo_Classifications proto.InternalMessageInfo

func (m *Classifications) GetClasses() []*Class {
	if m != nil {
		return m.Classes
	}
	return nil
}

// Contains one result per input example, in the same order as the input in
// ClassificationRequest.
type ClassificationResult struct {
	Classifications      []*Classifications `protobuf:"bytes,1,rep,name=classifications,proto3" json:"classifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClassificationResult) Reset()         { *m = ClassificationResult{} }
func (m *ClassificationResult) String() string { return proto.CompactTextString(m) }
func (*ClassificationResult) ProtoMessage()    {}
func (*ClassificationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_325571a8b7f54d57, []int{2}
}

func (m *ClassificationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationResult.Unmarshal(m, b)
}
func (m *ClassificationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationResult.Marshal(b, m, deterministic)
}
func (m *ClassificationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationResult.Merge(m, src)
}
func (m *ClassificationResult) XXX_Size() int {
	return xxx_messageInfo_ClassificationResult.Size(m)
}
func (m *ClassificationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationResult.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationResult proto.InternalMessageInfo

func (m *ClassificationResult) GetClassifications() []*Classifications {
	if m != nil {
		return m.Classifications
	}
	return nil
}

type ClassificationRequest struct {
	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Input data.
	Input                *Input   `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassificationRequest) Reset()         { *m = ClassificationRequest{} }
func (m *ClassificationRequest) String() string { return proto.CompactTextString(m) }
func (*ClassificationRequest) ProtoMessage()    {}
func (*ClassificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_325571a8b7f54d57, []int{3}
}

func (m *ClassificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationRequest.Unmarshal(m, b)
}
func (m *ClassificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationRequest.Marshal(b, m, deterministic)
}
func (m *ClassificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationRequest.Merge(m, src)
}
func (m *ClassificationRequest) XXX_Size() int {
	return xxx_messageInfo_ClassificationRequest.Size(m)
}
func (m *ClassificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationRequest proto.InternalMessageInfo

func (m *ClassificationRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *ClassificationRequest) GetInput() *Input {
	if m != nil {
		return m.Input
	}
	return nil
}

type ClassificationResponse struct {
	// Effective Model Specification used for classification.
	ModelSpec *ModelSpec `protobuf:"bytes,2,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Result of the classification.
	Result               *ClassificationResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClassificationResponse) Reset()         { *m = ClassificationResponse{} }
func (m *ClassificationResponse) String() string { return proto.CompactTextString(m) }
func (*ClassificationResponse) ProtoMessage()    {}
func (*ClassificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_325571a8b7f54d57, []int{4}
}

func (m *ClassificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationResponse.Unmarshal(m, b)
}
func (m *ClassificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationResponse.Marshal(b, m, deterministic)
}
func (m *ClassificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationResponse.Merge(m, src)
}
func (m *ClassificationResponse) XXX_Size() int {
	return xxx_messageInfo_ClassificationResponse.Size(m)
}
func (m *ClassificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationResponse proto.InternalMessageInfo

func (m *ClassificationResponse) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *ClassificationResponse) GetResult() *ClassificationResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Class)(nil), "tensorflow.serving.Class")
	proto.RegisterType((*Classifications)(nil), "tensorflow.serving.Classifications")
	proto.RegisterType((*ClassificationResult)(nil), "tensorflow.serving.ClassificationResult")
	proto.RegisterType((*ClassificationRequest)(nil), "tensorflow.serving.ClassificationRequest")
	proto.RegisterType((*ClassificationResponse)(nil), "tensorflow.serving.ClassificationResponse")
}

func init() {
	proto.RegisterFile("tensorflow.serving/proto/classification.proto", fileDescriptor_325571a8b7f54d57)
}

var fileDescriptor_325571a8b7f54d57 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0xd9, 0x94, 0x54, 0x3a, 0x3d, 0x14, 0x96, 0x2a, 0x51, 0x10, 0x42, 0xf4, 0x90, 0x8b,
	0x09, 0x34, 0x57, 0x0f, 0x62, 0x41, 0xf0, 0xd0, 0xcb, 0xfa, 0x00, 0x92, 0xc6, 0xa9, 0x04, 0xb6,
	0xd9, 0x98, 0xd9, 0xe8, 0x1b, 0xf8, 0x0c, 0x3e, 0xaa, 0x47, 0xc9, 0x6e, 0x83, 0x24, 0x31, 0x8a,
	0xb7, 0xec, 0xf0, 0x9b, 0xf9, 0x7f, 0x10, 0xb8, 0xd2, 0x58, 0x90, 0xaa, 0x76, 0x52, 0xbd, 0x45,
	0x84, 0xd5, 0x6b, 0x5e, 0x3c, 0xc7, 0x65, 0xa5, 0xb4, 0x8a, 0x33, 0x99, 0x12, 0xe5, 0xbb, 0x3c,
	0x4b, 0x75, 0xae, 0x8a, 0xc8, 0x0c, 0x39, 0x1f, 0xe2, 0x67, 0x97, 0xa3, 0x27, 0xf2, 0xa2, 0xac,
	0xb5, 0xdd, 0xfc, 0x85, 0xda, 0xab, 0x27, 0x94, 0x96, 0x0a, 0x12, 0x70, 0xd7, 0x8d, 0x2e, 0x5f,
	0x82, 0x2b, 0xd3, 0x2d, 0x4a, 0x8f, 0xf9, 0x2c, 0x9c, 0x09, 0xfb, 0x68, 0xa6, 0x94, 0xa9, 0x0a,
	0x3d, 0xc7, 0x67, 0xa1, 0x23, 0xec, 0x23, 0xb8, 0x83, 0xc5, 0xba, 0x63, 0x96, 0x78, 0x02, 0x47,
	0xc6, 0x3f, 0x92, 0xc7, 0xfc, 0x49, 0x38, 0x5f, 0x9d, 0x46, 0x43, 0xfd, 0xc8, 0x6c, 0x89, 0x96,
	0x0c, 0x10, 0x96, 0xdd, 0x3b, 0x02, 0xa9, 0x96, 0x9a, 0x6f, 0x60, 0xd1, 0x2d, 0xa3, 0x3d, 0x7a,
	0x31, 0x7a, 0xf4, 0x1b, 0x15, 0xfd, 0xdd, 0xe0, 0x9d, 0xc1, 0x71, 0x5f, 0xe7, 0xa5, 0x46, 0xd2,
	0xfc, 0x1a, 0xc0, 0x94, 0xf1, 0x48, 0x25, 0x66, 0x26, 0xf9, 0x7c, 0x75, 0xfe, 0x93, 0xc6, 0xa6,
	0xa1, 0x1e, 0x4a, 0xcc, 0xc4, 0x6c, 0xdf, 0x7e, 0xf2, 0x18, 0x5c, 0x53, 0xb8, 0x29, 0x67, 0x24,
	0xf1, 0x7d, 0x03, 0x08, 0xcb, 0x05, 0x1f, 0x0c, 0x4e, 0x06, 0x81, 0x4b, 0x55, 0x10, 0xf6, 0x9c,
	0x38, 0xff, 0x74, 0x72, 0x03, 0xd3, 0xca, 0x54, 0x77, 0xc8, 0x10, 0xfe, 0xdd, 0x93, 0xad, 0x5a,
	0x1c, 0xf6, 0x6e, 0x27, 0x9f, 0x8c, 0x6d, 0xa7, 0xe6, 0x9f, 0x48, 0xbe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x41, 0xf8, 0x17, 0x6f, 0xa4, 0x02, 0x00, 0x00,
}
