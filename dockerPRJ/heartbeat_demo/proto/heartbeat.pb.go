// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat.proto

package heartbeat

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

type Request struct {
	ClientId             uint64   `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	UserId               uint64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MsgId                uint64   `protobuf:"varint,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetClientId() uint64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Request) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Request) GetMsgId() uint64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *Request) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Response struct {
	ClientId             uint64   `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	UserId               uint64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MsgId                uint64   `protobuf:"varint,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetClientId() uint64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Response) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Response) GetMsgId() uint64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "heartbeat.Request")
	proto.RegisterType((*Response)(nil), "heartbeat.Response")
}

func init() { proto.RegisterFile("heartbeat.proto", fileDescriptor_3c667767fb9826a9) }

var fileDescriptor_3c667767fb9826a9 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0x4d, 0x2c,
	0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x65, 0x71, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x73, 0x71, 0x26, 0xe7,
	0x64, 0xa6, 0xe6, 0x95, 0xc4, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x71, 0x40,
	0x04, 0x3c, 0x53, 0x84, 0xc4, 0xb9, 0xd8, 0x4b, 0x8b, 0x53, 0x8b, 0x40, 0x52, 0x4c, 0x60, 0x29,
	0x36, 0x10, 0xd7, 0x33, 0x45, 0x48, 0x94, 0x8b, 0x2d, 0xb7, 0x38, 0x1d, 0x24, 0xce, 0x0c, 0x16,
	0x67, 0xcd, 0x2d, 0x4e, 0xf7, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60,
	0x51, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xb2, 0xb9, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0x69, 0x6e, 0x59, 0x12, 0x1b, 0xd8, 0xab, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x25, 0xad, 0xf6, 0xfd, 0xfd, 0x00, 0x00, 0x00,
}
