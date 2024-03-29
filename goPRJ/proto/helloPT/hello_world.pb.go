// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloPT/hello_world.proto

package hello_world

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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f0e144b618a37f5, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting             string   `protobuf:"bytes,2,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f0e144b618a37f5, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
}

func init() { proto.RegisterFile("helloPT/hello_world.proto", fileDescriptor_9f0e144b618a37f5) }

var fileDescriptor_9f0e144b618a37f5 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x48, 0xcd, 0xc9,
	0xc9, 0x0f, 0x08, 0xd1, 0x07, 0xd3, 0xf1, 0xe5, 0xf9, 0x45, 0x39, 0x29, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e, 0x20, 0xc1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30,
	0x5b, 0x49, 0x9b, 0x8b, 0x17, 0xaa, 0xa6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8a, 0x8b,
	0x23, 0xbd, 0x28, 0x35, 0xb5, 0x24, 0x33, 0x2f, 0x5d, 0x82, 0x09, 0xac, 0x10, 0xce, 0x37, 0x32,
	0xe3, 0xe2, 0x02, 0x2b, 0x0e, 0x07, 0x59, 0x22, 0xa4, 0xc1, 0xc5, 0x0a, 0xe6, 0x09, 0xf1, 0xea,
	0x21, 0x5b, 0x23, 0xc5, 0xa7, 0x87, 0x62, 0xa2, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x3d, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xf1, 0x11, 0xc3, 0xac, 0x00, 0x00, 0x00,
}
