// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/ssmessage/ssToGateError.proto

package ssmessage

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

type SsToGateError struct {
	Guid                 *uint64  `protobuf:"varint,1,opt,name=guid" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SsToGateError) Reset()         { *m = SsToGateError{} }
func (m *SsToGateError) String() string { return proto.CompactTextString(m) }
func (*SsToGateError) ProtoMessage()    {}
func (*SsToGateError) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab73afe740ff3b01, []int{0}
}

func (m *SsToGateError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SsToGateError.Unmarshal(m, b)
}
func (m *SsToGateError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SsToGateError.Marshal(b, m, deterministic)
}
func (m *SsToGateError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SsToGateError.Merge(m, src)
}
func (m *SsToGateError) XXX_Size() int {
	return xxx_messageInfo_SsToGateError.Size(m)
}
func (m *SsToGateError) XXX_DiscardUnknown() {
	xxx_messageInfo_SsToGateError.DiscardUnknown(m)
}

var xxx_messageInfo_SsToGateError proto.InternalMessageInfo

func (m *SsToGateError) GetGuid() uint64 {
	if m != nil && m.Guid != nil {
		return *m.Guid
	}
	return 0
}

func init() {
	proto.RegisterType((*SsToGateError)(nil), "ssmessage.ssToGateError")
}

func init() {
	proto.RegisterFile("ss/proto/ssmessage/ssToGateError.proto", fileDescriptor_ab73afe740ff3b01)
}

var fileDescriptor_ab73afe740ff3b01 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0xce, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2f,
	0x2e, 0x0e, 0xc9, 0x77, 0x4f, 0x2c, 0x49, 0x75, 0x2d, 0x2a, 0xca, 0x2f, 0xd2, 0x03, 0xcb, 0x0a,
	0x71, 0xc2, 0xa5, 0x95, 0x94, 0xb9, 0x78, 0x51, 0x54, 0x08, 0x09, 0x71, 0xb1, 0xa4, 0x97, 0x66,
	0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x81, 0xd9, 0x4e, 0xf2, 0x51, 0xb2, 0x25, 0xa9,
	0x89, 0xc9, 0x19, 0xa9, 0x45, 0x10, 0xe3, 0x93, 0xf3, 0x73, 0xf4, 0x8b, 0x8b, 0x11, 0x96, 0x00,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x22, 0x51, 0x9f, 0x79, 0x00, 0x00, 0x00,
}