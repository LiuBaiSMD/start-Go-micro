// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package base

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

// Example message
type Event struct {
	// unique id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// unix timestamp
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// message
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
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

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xf2, 0xe7, 0x62, 0x75, 0x2d, 0x4b, 0xcd, 0x2b, 0x11,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11,
	0x92, 0xe1, 0xe2, 0x2c, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x52, 0x60,
	0xd4, 0x60, 0x0e, 0x42, 0x08, 0x08, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7,
	0x4a, 0x30, 0x83, 0xb5, 0xc0, 0xb8, 0x49, 0x6c, 0x60, 0x73, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xf7, 0x25, 0x49, 0x7e, 0x65, 0x00, 0x00, 0x00,
}
