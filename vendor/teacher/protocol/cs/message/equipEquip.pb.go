// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/equipEquip.proto

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

type EquipEquipToS struct {
	CharacterID          *int32   `protobuf:"varint,1,req,name=characterID" json:"characterID,omitempty"`
	EquipID              *uint32  `protobuf:"varint,2,req,name=equipID" json:"equipID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EquipEquipToS) Reset()         { *m = EquipEquipToS{} }
func (m *EquipEquipToS) String() string { return proto.CompactTextString(m) }
func (*EquipEquipToS) ProtoMessage()    {}
func (*EquipEquipToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9438df244caa202c, []int{0}
}

func (m *EquipEquipToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquipEquipToS.Unmarshal(m, b)
}
func (m *EquipEquipToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquipEquipToS.Marshal(b, m, deterministic)
}
func (m *EquipEquipToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipEquipToS.Merge(m, src)
}
func (m *EquipEquipToS) XXX_Size() int {
	return xxx_messageInfo_EquipEquipToS.Size(m)
}
func (m *EquipEquipToS) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipEquipToS.DiscardUnknown(m)
}

var xxx_messageInfo_EquipEquipToS proto.InternalMessageInfo

func (m *EquipEquipToS) GetCharacterID() int32 {
	if m != nil && m.CharacterID != nil {
		return *m.CharacterID
	}
	return 0
}

func (m *EquipEquipToS) GetEquipID() uint32 {
	if m != nil && m.EquipID != nil {
		return *m.EquipID
	}
	return 0
}

type EquipEquipToC struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EquipEquipToC) Reset()         { *m = EquipEquipToC{} }
func (m *EquipEquipToC) String() string { return proto.CompactTextString(m) }
func (*EquipEquipToC) ProtoMessage()    {}
func (*EquipEquipToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_9438df244caa202c, []int{1}
}

func (m *EquipEquipToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquipEquipToC.Unmarshal(m, b)
}
func (m *EquipEquipToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquipEquipToC.Marshal(b, m, deterministic)
}
func (m *EquipEquipToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipEquipToC.Merge(m, src)
}
func (m *EquipEquipToC) XXX_Size() int {
	return xxx_messageInfo_EquipEquipToC.Size(m)
}
func (m *EquipEquipToC) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipEquipToC.DiscardUnknown(m)
}

var xxx_messageInfo_EquipEquipToC proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EquipEquipToS)(nil), "message.equipEquip_toS")
	proto.RegisterType((*EquipEquipToC)(nil), "message.equipEquip_toC")
}

func init() { proto.RegisterFile("cs/proto/message/equipEquip.proto", fileDescriptor_9438df244caa202c) }

var fileDescriptor_9438df244caa202c = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x4f, 0x2d, 0x2c,
	0xcd, 0x2c, 0x70, 0x05, 0x11, 0x7a, 0x60, 0x09, 0x21, 0x76, 0xa8, 0x8c, 0x92, 0x0f, 0x17, 0x1f,
	0x42, 0x32, 0xbe, 0x24, 0x3f, 0x58, 0x48, 0x81, 0x8b, 0x3b, 0x39, 0x23, 0xb1, 0x28, 0x31, 0xb9,
	0x24, 0xb5, 0xc8, 0xd3, 0x45, 0x82, 0x51, 0x81, 0x49, 0x83, 0x35, 0x08, 0x59, 0x48, 0x48, 0x82,
	0x8b, 0x1d, 0xac, 0xc7, 0xd3, 0x45, 0x82, 0x49, 0x81, 0x49, 0x83, 0x37, 0x08, 0xc6, 0x55, 0x12,
	0x40, 0x33, 0xcd, 0xd9, 0x49, 0x36, 0x4a, 0xba, 0x24, 0x35, 0x31, 0x39, 0x23, 0xb5, 0x08, 0xe2,
	0xa4, 0xe4, 0xfc, 0x1c, 0xfd, 0xe4, 0x62, 0x98, 0xc3, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0x70,
	0xdb, 0xb3, 0x09, 0xab, 0x00, 0x00, 0x00,
}