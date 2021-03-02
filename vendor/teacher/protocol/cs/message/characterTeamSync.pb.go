// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/characterTeamSync.proto

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

type CharacterTeamSyncToS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CharacterTeamSyncToS) Reset()         { *m = CharacterTeamSyncToS{} }
func (m *CharacterTeamSyncToS) String() string { return proto.CompactTextString(m) }
func (*CharacterTeamSyncToS) ProtoMessage()    {}
func (*CharacterTeamSyncToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fd0fbccb5c97e3c, []int{0}
}

func (m *CharacterTeamSyncToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharacterTeamSyncToS.Unmarshal(m, b)
}
func (m *CharacterTeamSyncToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharacterTeamSyncToS.Marshal(b, m, deterministic)
}
func (m *CharacterTeamSyncToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterTeamSyncToS.Merge(m, src)
}
func (m *CharacterTeamSyncToS) XXX_Size() int {
	return xxx_messageInfo_CharacterTeamSyncToS.Size(m)
}
func (m *CharacterTeamSyncToS) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterTeamSyncToS.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterTeamSyncToS proto.InternalMessageInfo

type CharacterTeamSyncToC struct {
	CharacterTeam        *CharacterTeam `protobuf:"bytes,1,req,name=characterTeam" json:"characterTeam,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CharacterTeamSyncToC) Reset()         { *m = CharacterTeamSyncToC{} }
func (m *CharacterTeamSyncToC) String() string { return proto.CompactTextString(m) }
func (*CharacterTeamSyncToC) ProtoMessage()    {}
func (*CharacterTeamSyncToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fd0fbccb5c97e3c, []int{1}
}

func (m *CharacterTeamSyncToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharacterTeamSyncToC.Unmarshal(m, b)
}
func (m *CharacterTeamSyncToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharacterTeamSyncToC.Marshal(b, m, deterministic)
}
func (m *CharacterTeamSyncToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterTeamSyncToC.Merge(m, src)
}
func (m *CharacterTeamSyncToC) XXX_Size() int {
	return xxx_messageInfo_CharacterTeamSyncToC.Size(m)
}
func (m *CharacterTeamSyncToC) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterTeamSyncToC.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterTeamSyncToC proto.InternalMessageInfo

func (m *CharacterTeamSyncToC) GetCharacterTeam() *CharacterTeam {
	if m != nil {
		return m.CharacterTeam
	}
	return nil
}

type CharacterTeam struct {
	CharacterList        []int32  `protobuf:"varint,1,rep,name=characterList" json:"characterList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CharacterTeam) Reset()         { *m = CharacterTeam{} }
func (m *CharacterTeam) String() string { return proto.CompactTextString(m) }
func (*CharacterTeam) ProtoMessage()    {}
func (*CharacterTeam) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fd0fbccb5c97e3c, []int{2}
}

func (m *CharacterTeam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharacterTeam.Unmarshal(m, b)
}
func (m *CharacterTeam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharacterTeam.Marshal(b, m, deterministic)
}
func (m *CharacterTeam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterTeam.Merge(m, src)
}
func (m *CharacterTeam) XXX_Size() int {
	return xxx_messageInfo_CharacterTeam.Size(m)
}
func (m *CharacterTeam) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterTeam.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterTeam proto.InternalMessageInfo

func (m *CharacterTeam) GetCharacterList() []int32 {
	if m != nil {
		return m.CharacterList
	}
	return nil
}

func init() {
	proto.RegisterType((*CharacterTeamSyncToS)(nil), "message.characterTeamSync_toS")
	proto.RegisterType((*CharacterTeamSyncToC)(nil), "message.characterTeamSync_toC")
	proto.RegisterType((*CharacterTeam)(nil), "message.characterTeam")
}

func init() {
	proto.RegisterFile("cs/proto/message/characterTeamSync.proto", fileDescriptor_9fd0fbccb5c97e3c)
}

var fileDescriptor_9fd0fbccb5c97e3c = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x4f, 0xce, 0x48,
	0x2c, 0x4a, 0x4c, 0x2e, 0x49, 0x2d, 0x0a, 0x49, 0x4d, 0xcc, 0x0d, 0xae, 0xcc, 0x4b, 0xd6, 0x03,
	0xcb, 0x0b, 0xb1, 0x43, 0x15, 0x28, 0x89, 0x73, 0x89, 0x62, 0xa8, 0x89, 0x2f, 0xc9, 0x0f, 0x56,
	0x0a, 0xc5, 0x2e, 0xe1, 0x2c, 0x64, 0xc3, 0xc5, 0x8b, 0x22, 0x21, 0xc1, 0xa8, 0xc0, 0xa4, 0xc1,
	0x6d, 0x24, 0xa6, 0x07, 0x35, 0x52, 0x0f, 0x45, 0x36, 0x08, 0x55, 0xb1, 0x92, 0x29, 0x9a, 0x6e,
	0x21, 0x15, 0x24, 0x01, 0x9f, 0xcc, 0xe2, 0x12, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xd6, 0x20, 0x54,
	0x41, 0x27, 0xd9, 0x28, 0xe9, 0x92, 0xd4, 0xc4, 0xe4, 0x8c, 0xd4, 0x22, 0x88, 0x07, 0x93, 0xf3,
	0x73, 0xf4, 0x93, 0x8b, 0x61, 0xde, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x9d, 0xaa, 0x5c,
	0xf9, 0x00, 0x00, 0x00,
}