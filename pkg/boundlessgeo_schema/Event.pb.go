// Code generated by protoc-gen-go. DO NOT EDIT.
// source: boundlessgeo_schema/Event.proto

package schema

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

// An event emitted
type Event struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Data                 []byte    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	ParentId             string    `protobuf:"bytes,4,opt,name=parentId,proto3" json:"parentId,omitempty"`
	Metadata             *Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Created              int64     `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_Event_079970722a233327, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *Event) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Event) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
}

func init() {
	proto.RegisterFile("boundlessgeo_schema/Event.proto", fileDescriptor_Event_079970722a233327)
}

var fileDescriptor_Event_079970722a233327 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xca, 0x2f, 0xcd,
	0x4b, 0xc9, 0x49, 0x2d, 0x2e, 0x4e, 0x4f, 0xcd, 0x8f, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4,
	0x77, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x52, 0xc2, 0xa6, 0xc0,
	0x37, 0xb5, 0x24, 0x31, 0x25, 0xb1, 0x24, 0x11, 0xa2, 0x46, 0x69, 0x0e, 0x23, 0x17, 0x2b, 0x58,
	0x8f, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x53, 0x66,
	0x8a, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x65, 0x41, 0xaa, 0x04, 0x13, 0x58, 0x04, 0xcc, 0x06, 0x89,
	0x81, 0xf4, 0x4a, 0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x42, 0x52, 0x5c, 0x1c, 0x05,
	0x89, 0x45, 0xa9, 0x79, 0x25, 0x9e, 0x29, 0x12, 0x2c, 0x60, 0xb5, 0x70, 0xbe, 0x90, 0x2a, 0x17,
	0x47, 0x2e, 0xd4, 0x3e, 0x09, 0x56, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x4e, 0x3d, 0x98, 0x03, 0x82,
	0xe0, 0x52, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x29, 0x12, 0x6c, 0x0a,
	0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x93, 0x36, 0x97, 0x78, 0x72, 0x7e, 0xae, 0x1e, 0xb2, 0x47,
	0xf4, 0x20, 0x1e, 0x71, 0xe2, 0x00, 0x3b, 0x3b, 0x20, 0x29, 0x2d, 0x8a, 0x0d, 0x22, 0x92, 0xc4,
	0x06, 0xf6, 0x92, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x75, 0xb1, 0x5a, 0xd5, 0x19, 0x01, 0x00,
	0x00,
}
