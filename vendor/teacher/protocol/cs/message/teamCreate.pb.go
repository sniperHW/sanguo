// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/teamCreate.proto

package message

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

type TeamCreateToS struct {
	Target               *TeamTarget `protobuf:"bytes,1,req,name=target" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TeamCreateToS) Reset()         { *m = TeamCreateToS{} }
func (m *TeamCreateToS) String() string { return proto.CompactTextString(m) }
func (*TeamCreateToS) ProtoMessage()    {}
func (*TeamCreateToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fda5f53cab150eb6, []int{0}
}

func (m *TeamCreateToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamCreateToS.Unmarshal(m, b)
}
func (m *TeamCreateToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamCreateToS.Marshal(b, m, deterministic)
}
func (m *TeamCreateToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamCreateToS.Merge(m, src)
}
func (m *TeamCreateToS) XXX_Size() int {
	return xxx_messageInfo_TeamCreateToS.Size(m)
}
func (m *TeamCreateToS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamCreateToS.DiscardUnknown(m)
}

var xxx_messageInfo_TeamCreateToS proto.InternalMessageInfo

func (m *TeamCreateToS) GetTarget() *TeamTarget {
	if m != nil {
		return m.Target
	}
	return nil
}

type TeamCreateToC struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamCreateToC) Reset()         { *m = TeamCreateToC{} }
func (m *TeamCreateToC) String() string { return proto.CompactTextString(m) }
func (*TeamCreateToC) ProtoMessage()    {}
func (*TeamCreateToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_fda5f53cab150eb6, []int{1}
}

func (m *TeamCreateToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamCreateToC.Unmarshal(m, b)
}
func (m *TeamCreateToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamCreateToC.Marshal(b, m, deterministic)
}
func (m *TeamCreateToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamCreateToC.Merge(m, src)
}
func (m *TeamCreateToC) XXX_Size() int {
	return xxx_messageInfo_TeamCreateToC.Size(m)
}
func (m *TeamCreateToC) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamCreateToC.DiscardUnknown(m)
}

var xxx_messageInfo_TeamCreateToC proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TeamCreateToS)(nil), "message.teamCreate_toS")
	proto.RegisterType((*TeamCreateToC)(nil), "message.teamCreate_toC")
}

func init() { proto.RegisterFile("cs/proto/message/teamCreate.proto", fileDescriptor_fda5f53cab150eb6) }

var fileDescriptor_fda5f53cab150eb6 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2f, 0x49, 0x4d,
	0xcc, 0x75, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0xd5, 0x03, 0x4b, 0x08, 0xb1, 0x43, 0x65, 0xa4, 0xa4,
	0xb1, 0xaa, 0x85, 0xa8, 0x52, 0xb2, 0xe5, 0xe2, 0x43, 0xe8, 0x8c, 0x2f, 0xc9, 0x0f, 0x16, 0xd2,
	0xe6, 0x62, 0x2b, 0x49, 0x2c, 0x4a, 0x4f, 0x2d, 0x91, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x36, 0x12,
	0xd6, 0x83, 0x6a, 0xd3, 0x0b, 0x49, 0x4d, 0xcc, 0x0d, 0x01, 0x4b, 0x05, 0x41, 0x95, 0x28, 0x09,
	0xa0, 0x69, 0x77, 0x76, 0x92, 0x8d, 0x92, 0x2e, 0x49, 0x4d, 0x4c, 0xce, 0x48, 0x2d, 0x82, 0x58,
	0x9a, 0x9c, 0x9f, 0xa3, 0x9f, 0x5c, 0x0c, 0xb3, 0x1a, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x9c,
	0x51, 0x51, 0xb9, 0x00, 0x00, 0x00,
}
