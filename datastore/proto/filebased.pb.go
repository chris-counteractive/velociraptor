// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filebased.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Attribute struct {
	Attribute            string   `protobuf:"bytes,1,opt,name=attribute,proto3" json:"attribute,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attribute) Reset()         { *m = Attribute{} }
func (m *Attribute) String() string { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()    {}
func (*Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_filebased_142fcd64f0d5dae0, []int{0}
}
func (m *Attribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attribute.Unmarshal(m, b)
}
func (m *Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attribute.Marshal(b, m, deterministic)
}
func (dst *Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribute.Merge(dst, src)
}
func (m *Attribute) XXX_Size() int {
	return xxx_messageInfo_Attribute.Size(m)
}
func (m *Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Attribute proto.InternalMessageInfo

func (m *Attribute) GetAttribute() string {
	if m != nil {
		return m.Attribute
	}
	return ""
}

func (m *Attribute) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type FileContent struct {
	Attr                 []*Attribute `protobuf:"bytes,1,rep,name=attr,proto3" json:"attr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FileContent) Reset()         { *m = FileContent{} }
func (m *FileContent) String() string { return proto.CompactTextString(m) }
func (*FileContent) ProtoMessage()    {}
func (*FileContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_filebased_142fcd64f0d5dae0, []int{1}
}
func (m *FileContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileContent.Unmarshal(m, b)
}
func (m *FileContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileContent.Marshal(b, m, deterministic)
}
func (dst *FileContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileContent.Merge(dst, src)
}
func (m *FileContent) XXX_Size() int {
	return xxx_messageInfo_FileContent.Size(m)
}
func (m *FileContent) XXX_DiscardUnknown() {
	xxx_messageInfo_FileContent.DiscardUnknown(m)
}

var xxx_messageInfo_FileContent proto.InternalMessageInfo

func (m *FileContent) GetAttr() []*Attribute {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*Attribute)(nil), "proto.Attribute")
	proto.RegisterType((*FileContent)(nil), "proto.FileContent")
}

func init() { proto.RegisterFile("filebased.proto", fileDescriptor_filebased_142fcd64f0d5dae0) }

var fileDescriptor_filebased_142fcd64f0d5dae0 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcb, 0xcc, 0x49,
	0x4d, 0x4a, 0x2c, 0x4e, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a,
	0xb6, 0x5c, 0x9c, 0x8e, 0x25, 0x25, 0x45, 0x99, 0x49, 0xa5, 0x25, 0xa9, 0x42, 0x32, 0x5c, 0x9c,
	0x89, 0x30, 0x8e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x42, 0x40, 0x48, 0x88, 0x8b, 0x25,
	0x25, 0xb1, 0x24, 0x51, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x56, 0x32, 0xe6, 0xe2,
	0x76, 0xcb, 0xcc, 0x49, 0x75, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x11, 0x52, 0xe1, 0x62, 0x01,
	0xa9, 0x97, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x12, 0x80, 0x58, 0xa5, 0x07, 0xb7, 0x20, 0x08,
	0x2c, 0x9b, 0xc4, 0x06, 0x16, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xef, 0x7b, 0x46,
	0x94, 0x00, 0x00, 0x00,
}