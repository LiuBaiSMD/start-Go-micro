// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/metaPT/api.proto

package api

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

type CallRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallRequest) Reset()         { *m = CallRequest{} }
func (m *CallRequest) String() string { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()    {}
func (*CallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_799d05e6110794ff, []int{0}
}

func (m *CallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallRequest.Unmarshal(m, b)
}
func (m *CallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallRequest.Marshal(b, m, deterministic)
}
func (m *CallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallRequest.Merge(m, src)
}
func (m *CallRequest) XXX_Size() int {
	return xxx_messageInfo_CallRequest.Size(m)
}
func (m *CallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallRequest proto.InternalMessageInfo

func (m *CallRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CallResponse struct {
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallResponse) Reset()         { *m = CallResponse{} }
func (m *CallResponse) String() string { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()    {}
func (*CallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_799d05e6110794ff, []int{1}
}

func (m *CallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallResponse.Unmarshal(m, b)
}
func (m *CallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallResponse.Marshal(b, m, deterministic)
}
func (m *CallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallResponse.Merge(m, src)
}
func (m *CallResponse) XXX_Size() int {
	return xxx_messageInfo_CallResponse.Size(m)
}
func (m *CallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallResponse proto.InternalMessageInfo

func (m *CallResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_799d05e6110794ff, []int{2}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_799d05e6110794ff, []int{3}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CallRequest)(nil), "CallRequest")
	proto.RegisterType((*CallResponse)(nil), "CallResponse")
	proto.RegisterType((*EmptyRequest)(nil), "EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
}

func init() { proto.RegisterFile("proto/metaPT/api.proto", fileDescriptor_799d05e6110794ff) }

var fileDescriptor_799d05e6110794ff = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x49, 0x0c, 0x08, 0xd1, 0x4f, 0x2c, 0xc8, 0xd4, 0x03, 0x0b, 0x28,
	0x29, 0x72, 0x71, 0x3b, 0x27, 0xe6, 0xe4, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09,
	0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a,
	0x1a, 0x5c, 0x3c, 0x10, 0x25, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c, 0xec, 0xb9,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x60, 0x65, 0x30, 0xae, 0x12, 0x1f, 0x17, 0x8f,
	0x6b, 0x6e, 0x41, 0x49, 0x25, 0xd4, 0x34, 0x25, 0x7e, 0x2e, 0x5e, 0x28, 0x1f, 0xa2, 0xd5, 0xc8,
	0x80, 0x8b, 0xdd, 0xb5, 0x22, 0x31, 0xb7, 0x20, 0x27, 0x55, 0x48, 0x95, 0x8b, 0x05, 0x64, 0xaa,
	0x10, 0x8f, 0x1e, 0x92, 0xfd, 0x52, 0xbc, 0x7a, 0xc8, 0x56, 0x29, 0x31, 0x18, 0xe9, 0x72, 0x31,
	0xbb, 0xe5, 0xe7, 0x0b, 0xa9, 0x71, 0x31, 0x3b, 0x25, 0x16, 0x09, 0xf1, 0xea, 0x21, 0x9b, 0x2f,
	0xc5, 0xa7, 0x87, 0x62, 0xbc, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x57, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5c, 0xc8, 0xbf, 0x76, 0xef, 0x00, 0x00, 0x00,
}
