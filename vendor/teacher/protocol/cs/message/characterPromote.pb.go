// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/characterPromote.proto

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

type CharacterPromoteToS struct {
	CharacterID          *int32   `protobuf:"varint,1,opt,name=CharacterID" json:"CharacterID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CharacterPromoteToS) Reset()         { *m = CharacterPromoteToS{} }
func (m *CharacterPromoteToS) String() string { return proto.CompactTextString(m) }
func (*CharacterPromoteToS) ProtoMessage()    {}
func (*CharacterPromoteToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7373448f190b24, []int{0}
}

func (m *CharacterPromoteToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharacterPromoteToS.Unmarshal(m, b)
}
func (m *CharacterPromoteToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharacterPromoteToS.Marshal(b, m, deterministic)
}
func (m *CharacterPromoteToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterPromoteToS.Merge(m, src)
}
func (m *CharacterPromoteToS) XXX_Size() int {
	return xxx_messageInfo_CharacterPromoteToS.Size(m)
}
func (m *CharacterPromoteToS) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterPromoteToS.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterPromoteToS proto.InternalMessageInfo

func (m *CharacterPromoteToS) GetCharacterID() int32 {
	if m != nil && m.CharacterID != nil {
		return *m.CharacterID
	}
	return 0
}

type CharacterPromoteToC struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CharacterPromoteToC) Reset()         { *m = CharacterPromoteToC{} }
func (m *CharacterPromoteToC) String() string { return proto.CompactTextString(m) }
func (*CharacterPromoteToC) ProtoMessage()    {}
func (*CharacterPromoteToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7373448f190b24, []int{1}
}

func (m *CharacterPromoteToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharacterPromoteToC.Unmarshal(m, b)
}
func (m *CharacterPromoteToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharacterPromoteToC.Marshal(b, m, deterministic)
}
func (m *CharacterPromoteToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterPromoteToC.Merge(m, src)
}
func (m *CharacterPromoteToC) XXX_Size() int {
	return xxx_messageInfo_CharacterPromoteToC.Size(m)
}
func (m *CharacterPromoteToC) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterPromoteToC.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterPromoteToC proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CharacterPromoteToS)(nil), "message.characterPromote_toS")
	proto.RegisterType((*CharacterPromoteToC)(nil), "message.characterPromote_toC")
}

func init() {
	proto.RegisterFile("cs/proto/message/characterPromote.proto", fileDescriptor_dd7373448f190b24)
}

var fileDescriptor_dd7373448f190b24 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x4f, 0xce, 0x48,
	0x2c, 0x4a, 0x4c, 0x2e, 0x49, 0x2d, 0x0a, 0x28, 0xca, 0xcf, 0xcd, 0x2f, 0x49, 0xd5, 0x03, 0x4b,
	0x0b, 0xb1, 0x43, 0xe5, 0x95, 0x2c, 0xb8, 0x44, 0xd0, 0x95, 0xc4, 0x97, 0xe4, 0x07, 0x0b, 0x29,
	0x70, 0x71, 0x3b, 0xc3, 0xc4, 0x3d, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x90, 0x85,
	0x94, 0xc4, 0xb0, 0xea, 0x74, 0x76, 0x92, 0x8d, 0x92, 0x2e, 0x49, 0x4d, 0x4c, 0xce, 0x48, 0x2d,
	0x82, 0x38, 0x25, 0x39, 0x3f, 0x47, 0x3f, 0xb9, 0x18, 0xe6, 0x20, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xaf, 0x6d, 0x58, 0x36, 0xa3, 0x00, 0x00, 0x00,
}
