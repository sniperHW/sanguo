// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/equipGroupEquip.proto

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

type EquipGroupEquipToS struct {
	GroupID              *int32   `protobuf:"varint,1,req,name=groupID" json:"groupID,omitempty"`
	CharacterID          *int32   `protobuf:"varint,2,req,name=characterID" json:"characterID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EquipGroupEquipToS) Reset()         { *m = EquipGroupEquipToS{} }
func (m *EquipGroupEquipToS) String() string { return proto.CompactTextString(m) }
func (*EquipGroupEquipToS) ProtoMessage()    {}
func (*EquipGroupEquipToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f382f41d2c143f, []int{0}
}

func (m *EquipGroupEquipToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquipGroupEquipToS.Unmarshal(m, b)
}
func (m *EquipGroupEquipToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquipGroupEquipToS.Marshal(b, m, deterministic)
}
func (m *EquipGroupEquipToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipGroupEquipToS.Merge(m, src)
}
func (m *EquipGroupEquipToS) XXX_Size() int {
	return xxx_messageInfo_EquipGroupEquipToS.Size(m)
}
func (m *EquipGroupEquipToS) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipGroupEquipToS.DiscardUnknown(m)
}

var xxx_messageInfo_EquipGroupEquipToS proto.InternalMessageInfo

func (m *EquipGroupEquipToS) GetGroupID() int32 {
	if m != nil && m.GroupID != nil {
		return *m.GroupID
	}
	return 0
}

func (m *EquipGroupEquipToS) GetCharacterID() int32 {
	if m != nil && m.CharacterID != nil {
		return *m.CharacterID
	}
	return 0
}

type EquipGroupEquipToC struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EquipGroupEquipToC) Reset()         { *m = EquipGroupEquipToC{} }
func (m *EquipGroupEquipToC) String() string { return proto.CompactTextString(m) }
func (*EquipGroupEquipToC) ProtoMessage()    {}
func (*EquipGroupEquipToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f382f41d2c143f, []int{1}
}

func (m *EquipGroupEquipToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquipGroupEquipToC.Unmarshal(m, b)
}
func (m *EquipGroupEquipToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquipGroupEquipToC.Marshal(b, m, deterministic)
}
func (m *EquipGroupEquipToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipGroupEquipToC.Merge(m, src)
}
func (m *EquipGroupEquipToC) XXX_Size() int {
	return xxx_messageInfo_EquipGroupEquipToC.Size(m)
}
func (m *EquipGroupEquipToC) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipGroupEquipToC.DiscardUnknown(m)
}

var xxx_messageInfo_EquipGroupEquipToC proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EquipGroupEquipToS)(nil), "message.equipGroupEquip_toS")
	proto.RegisterType((*EquipGroupEquipToC)(nil), "message.equipGroupEquip_toC")
}

func init() {
	proto.RegisterFile("cs/proto/message/equipGroupEquip.proto", fileDescriptor_83f382f41d2c143f)
}

var fileDescriptor_83f382f41d2c143f = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x4f, 0x2d, 0x2c,
	0xcd, 0x2c, 0x70, 0x2f, 0xca, 0x2f, 0x2d, 0x70, 0x05, 0xb1, 0xf4, 0xc0, 0xb2, 0x42, 0xec, 0x50,
	0x69, 0xa5, 0x40, 0x2e, 0x61, 0x34, 0x15, 0xf1, 0x25, 0xf9, 0xc1, 0x42, 0x12, 0x5c, 0xec, 0xe9,
	0x20, 0x11, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x26, 0x0d, 0xd6, 0x20, 0x18, 0x57, 0x48, 0x81, 0x8b,
	0x3b, 0x39, 0x23, 0xb1, 0x28, 0x31, 0xb9, 0x24, 0xb5, 0xc8, 0xd3, 0x45, 0x82, 0x09, 0x2c, 0x8b,
	0x2c, 0xa4, 0x24, 0x8a, 0xcd, 0x48, 0x67, 0x27, 0xd9, 0x28, 0xe9, 0x92, 0xd4, 0xc4, 0xe4, 0x8c,
	0xd4, 0x22, 0x88, 0x0b, 0x93, 0xf3, 0x73, 0xf4, 0x93, 0x8b, 0x61, 0xee, 0x04, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x17, 0x64, 0x97, 0x46, 0xba, 0x00, 0x00, 0x00,
}
