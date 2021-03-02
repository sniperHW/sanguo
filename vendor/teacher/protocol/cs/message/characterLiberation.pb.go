// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/characterLiberation.proto

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

type CharacterLiberationToS struct {
	CharacterID          *int32   `protobuf:"varint,1,opt,name=CharacterID" json:"CharacterID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CharacterLiberationToS) Reset()         { *m = CharacterLiberationToS{} }
func (m *CharacterLiberationToS) String() string { return proto.CompactTextString(m) }
func (*CharacterLiberationToS) ProtoMessage()    {}
func (*CharacterLiberationToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b27dbcecc48b29, []int{0}
}

func (m *CharacterLiberationToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharacterLiberationToS.Unmarshal(m, b)
}
func (m *CharacterLiberationToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharacterLiberationToS.Marshal(b, m, deterministic)
}
func (m *CharacterLiberationToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterLiberationToS.Merge(m, src)
}
func (m *CharacterLiberationToS) XXX_Size() int {
	return xxx_messageInfo_CharacterLiberationToS.Size(m)
}
func (m *CharacterLiberationToS) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterLiberationToS.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterLiberationToS proto.InternalMessageInfo

func (m *CharacterLiberationToS) GetCharacterID() int32 {
	if m != nil && m.CharacterID != nil {
		return *m.CharacterID
	}
	return 0
}

type CharacterLiberationToC struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CharacterLiberationToC) Reset()         { *m = CharacterLiberationToC{} }
func (m *CharacterLiberationToC) String() string { return proto.CompactTextString(m) }
func (*CharacterLiberationToC) ProtoMessage()    {}
func (*CharacterLiberationToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b27dbcecc48b29, []int{1}
}

func (m *CharacterLiberationToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharacterLiberationToC.Unmarshal(m, b)
}
func (m *CharacterLiberationToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharacterLiberationToC.Marshal(b, m, deterministic)
}
func (m *CharacterLiberationToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterLiberationToC.Merge(m, src)
}
func (m *CharacterLiberationToC) XXX_Size() int {
	return xxx_messageInfo_CharacterLiberationToC.Size(m)
}
func (m *CharacterLiberationToC) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterLiberationToC.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterLiberationToC proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CharacterLiberationToS)(nil), "message.characterLiberation_toS")
	proto.RegisterType((*CharacterLiberationToC)(nil), "message.characterLiberation_toC")
}

func init() {
	proto.RegisterFile("cs/proto/message/characterLiberation.proto", fileDescriptor_15b27dbcecc48b29)
}

var fileDescriptor_15b27dbcecc48b29 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x4f, 0xce, 0x48,
	0x2c, 0x4a, 0x4c, 0x2e, 0x49, 0x2d, 0xf2, 0xc9, 0x4c, 0x4a, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf,
	0xd3, 0x03, 0xab, 0x10, 0x62, 0x87, 0x2a, 0x51, 0xb2, 0xe6, 0x12, 0xc7, 0xa2, 0x2a, 0xbe, 0x24,
	0x3f, 0x58, 0x48, 0x81, 0x8b, 0xdb, 0x19, 0x26, 0xe5, 0xe9, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x1a, 0x84, 0x2c, 0xa4, 0x24, 0x89, 0x4b, 0xb3, 0xb3, 0x93, 0x6c, 0x94, 0x74, 0x49, 0x6a, 0x62,
	0x72, 0x46, 0x6a, 0x11, 0xc4, 0x4d, 0xc9, 0xf9, 0x39, 0xfa, 0xc9, 0xc5, 0x30, 0x97, 0x01, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xdb, 0x18, 0x0a, 0x58, 0xac, 0x00, 0x00, 0x00,
}