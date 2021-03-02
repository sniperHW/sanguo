// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/teamDismiss.proto

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

type TeamDismissToS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamDismissToS) Reset()         { *m = TeamDismissToS{} }
func (m *TeamDismissToS) String() string { return proto.CompactTextString(m) }
func (*TeamDismissToS) ProtoMessage()    {}
func (*TeamDismissToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae589c307b8aaf1a, []int{0}
}

func (m *TeamDismissToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamDismissToS.Unmarshal(m, b)
}
func (m *TeamDismissToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamDismissToS.Marshal(b, m, deterministic)
}
func (m *TeamDismissToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamDismissToS.Merge(m, src)
}
func (m *TeamDismissToS) XXX_Size() int {
	return xxx_messageInfo_TeamDismissToS.Size(m)
}
func (m *TeamDismissToS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamDismissToS.DiscardUnknown(m)
}

var xxx_messageInfo_TeamDismissToS proto.InternalMessageInfo

type TeamDismissToC struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamDismissToC) Reset()         { *m = TeamDismissToC{} }
func (m *TeamDismissToC) String() string { return proto.CompactTextString(m) }
func (*TeamDismissToC) ProtoMessage()    {}
func (*TeamDismissToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae589c307b8aaf1a, []int{1}
}

func (m *TeamDismissToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamDismissToC.Unmarshal(m, b)
}
func (m *TeamDismissToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamDismissToC.Marshal(b, m, deterministic)
}
func (m *TeamDismissToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamDismissToC.Merge(m, src)
}
func (m *TeamDismissToC) XXX_Size() int {
	return xxx_messageInfo_TeamDismissToC.Size(m)
}
func (m *TeamDismissToC) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamDismissToC.DiscardUnknown(m)
}

var xxx_messageInfo_TeamDismissToC proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TeamDismissToS)(nil), "message.teamDismiss_toS")
	proto.RegisterType((*TeamDismissToC)(nil), "message.teamDismiss_toC")
}

func init() { proto.RegisterFile("cs/proto/message/teamDismiss.proto", fileDescriptor_ae589c307b8aaf1a) }

var fileDescriptor_ae589c307b8aaf1a = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2f, 0x49, 0x4d,
	0xcc, 0x75, 0xc9, 0x2c, 0xce, 0xcd, 0x2c, 0x2e, 0xd6, 0x03, 0xcb, 0x08, 0xb1, 0x43, 0xa5, 0x94,
	0x04, 0xb9, 0xf8, 0x91, 0x64, 0xe3, 0x4b, 0xf2, 0x83, 0x31, 0x85, 0x9c, 0x9d, 0x64, 0xa3, 0xa4,
	0x4b, 0x52, 0x13, 0x93, 0x33, 0x52, 0x8b, 0x20, 0x26, 0x27, 0xe7, 0xe7, 0xe8, 0x27, 0x17, 0xc3,
	0xcc, 0x07, 0x04, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xb4, 0xa8, 0xce, 0x72, 0x00, 0x00, 0x00,
}
