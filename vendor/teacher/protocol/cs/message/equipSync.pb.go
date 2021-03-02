// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/equipSync.proto

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

type EquipSyncToS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EquipSyncToS) Reset()         { *m = EquipSyncToS{} }
func (m *EquipSyncToS) String() string { return proto.CompactTextString(m) }
func (*EquipSyncToS) ProtoMessage()    {}
func (*EquipSyncToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3793271f2ab7fe8f, []int{0}
}

func (m *EquipSyncToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquipSyncToS.Unmarshal(m, b)
}
func (m *EquipSyncToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquipSyncToS.Marshal(b, m, deterministic)
}
func (m *EquipSyncToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipSyncToS.Merge(m, src)
}
func (m *EquipSyncToS) XXX_Size() int {
	return xxx_messageInfo_EquipSyncToS.Size(m)
}
func (m *EquipSyncToS) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipSyncToS.DiscardUnknown(m)
}

var xxx_messageInfo_EquipSyncToS proto.InternalMessageInfo

type EquipSyncToC struct {
	IsAll                *bool    `protobuf:"varint,1,req,name=isAll" json:"isAll,omitempty"`
	Equips               []*Equip `protobuf:"bytes,2,rep,name=equips" json:"equips,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EquipSyncToC) Reset()         { *m = EquipSyncToC{} }
func (m *EquipSyncToC) String() string { return proto.CompactTextString(m) }
func (*EquipSyncToC) ProtoMessage()    {}
func (*EquipSyncToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_3793271f2ab7fe8f, []int{1}
}

func (m *EquipSyncToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquipSyncToC.Unmarshal(m, b)
}
func (m *EquipSyncToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquipSyncToC.Marshal(b, m, deterministic)
}
func (m *EquipSyncToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipSyncToC.Merge(m, src)
}
func (m *EquipSyncToC) XXX_Size() int {
	return xxx_messageInfo_EquipSyncToC.Size(m)
}
func (m *EquipSyncToC) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipSyncToC.DiscardUnknown(m)
}

var xxx_messageInfo_EquipSyncToC proto.InternalMessageInfo

func (m *EquipSyncToC) GetIsAll() bool {
	if m != nil && m.IsAll != nil {
		return *m.IsAll
	}
	return false
}

func (m *EquipSyncToC) GetEquips() []*Equip {
	if m != nil {
		return m.Equips
	}
	return nil
}

type Equip struct {
	ID                   *uint32      `protobuf:"varint,1,req,name=ID" json:"ID,omitempty"`
	ConfigID             *int32       `protobuf:"varint,2,opt,name=ConfigID" json:"ConfigID,omitempty"`
	Level                *int32       `protobuf:"varint,3,opt,name=level" json:"level,omitempty"`
	Exp                  *int32       `protobuf:"varint,4,opt,name=exp" json:"exp,omitempty"`
	BreakLevel           *int32       `protobuf:"varint,5,opt,name=breakLevel" json:"breakLevel,omitempty"`
	AllResonance         []*Resonance `protobuf:"bytes,6,rep,name=allResonance" json:"allResonance,omitempty"`
	EquipCharacterID     *int32       `protobuf:"varint,7,opt,name=EquipCharacterID" json:"EquipCharacterID,omitempty"`
	IsLock               *bool        `protobuf:"varint,8,opt,name=isLock" json:"isLock,omitempty"`
	IsRemove             *bool        `protobuf:"varint,9,opt,name=isRemove" json:"isRemove,omitempty"`
	GetTime              *int64       `protobuf:"varint,10,opt,name=getTime" json:"getTime,omitempty"`
	Overclock            []bool       `protobuf:"varint,11,rep,name=overclock" json:"overclock,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Equip) Reset()         { *m = Equip{} }
func (m *Equip) String() string { return proto.CompactTextString(m) }
func (*Equip) ProtoMessage()    {}
func (*Equip) Descriptor() ([]byte, []int) {
	return fileDescriptor_3793271f2ab7fe8f, []int{2}
}

func (m *Equip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Equip.Unmarshal(m, b)
}
func (m *Equip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Equip.Marshal(b, m, deterministic)
}
func (m *Equip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Equip.Merge(m, src)
}
func (m *Equip) XXX_Size() int {
	return xxx_messageInfo_Equip.Size(m)
}
func (m *Equip) XXX_DiscardUnknown() {
	xxx_messageInfo_Equip.DiscardUnknown(m)
}

var xxx_messageInfo_Equip proto.InternalMessageInfo

func (m *Equip) GetID() uint32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Equip) GetConfigID() int32 {
	if m != nil && m.ConfigID != nil {
		return *m.ConfigID
	}
	return 0
}

func (m *Equip) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Equip) GetExp() int32 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *Equip) GetBreakLevel() int32 {
	if m != nil && m.BreakLevel != nil {
		return *m.BreakLevel
	}
	return 0
}

func (m *Equip) GetAllResonance() []*Resonance {
	if m != nil {
		return m.AllResonance
	}
	return nil
}

func (m *Equip) GetEquipCharacterID() int32 {
	if m != nil && m.EquipCharacterID != nil {
		return *m.EquipCharacterID
	}
	return 0
}

func (m *Equip) GetIsLock() bool {
	if m != nil && m.IsLock != nil {
		return *m.IsLock
	}
	return false
}

func (m *Equip) GetIsRemove() bool {
	if m != nil && m.IsRemove != nil {
		return *m.IsRemove
	}
	return false
}

func (m *Equip) GetGetTime() int64 {
	if m != nil && m.GetTime != nil {
		return *m.GetTime
	}
	return 0
}

func (m *Equip) GetOverclock() []bool {
	if m != nil {
		return m.Overclock
	}
	return nil
}

type Resonance struct {
	SkillCfgID           *int32   `protobuf:"varint,1,req,name=SkillCfgID" json:"SkillCfgID,omitempty"`
	Order                *int32   `protobuf:"varint,2,req,name=order" json:"order,omitempty"`
	BindCharacterID      *int32   `protobuf:"varint,3,opt,name=BindCharacterID" json:"BindCharacterID,omitempty"`
	IsWaitReplace        *bool    `protobuf:"varint,4,opt,name=isWaitReplace" json:"isWaitReplace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resonance) Reset()         { *m = Resonance{} }
func (m *Resonance) String() string { return proto.CompactTextString(m) }
func (*Resonance) ProtoMessage()    {}
func (*Resonance) Descriptor() ([]byte, []int) {
	return fileDescriptor_3793271f2ab7fe8f, []int{3}
}

func (m *Resonance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resonance.Unmarshal(m, b)
}
func (m *Resonance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resonance.Marshal(b, m, deterministic)
}
func (m *Resonance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resonance.Merge(m, src)
}
func (m *Resonance) XXX_Size() int {
	return xxx_messageInfo_Resonance.Size(m)
}
func (m *Resonance) XXX_DiscardUnknown() {
	xxx_messageInfo_Resonance.DiscardUnknown(m)
}

var xxx_messageInfo_Resonance proto.InternalMessageInfo

func (m *Resonance) GetSkillCfgID() int32 {
	if m != nil && m.SkillCfgID != nil {
		return *m.SkillCfgID
	}
	return 0
}

func (m *Resonance) GetOrder() int32 {
	if m != nil && m.Order != nil {
		return *m.Order
	}
	return 0
}

func (m *Resonance) GetBindCharacterID() int32 {
	if m != nil && m.BindCharacterID != nil {
		return *m.BindCharacterID
	}
	return 0
}

func (m *Resonance) GetIsWaitReplace() bool {
	if m != nil && m.IsWaitReplace != nil {
		return *m.IsWaitReplace
	}
	return false
}

func init() {
	proto.RegisterType((*EquipSyncToS)(nil), "message.equipSync_toS")
	proto.RegisterType((*EquipSyncToC)(nil), "message.equipSync_toC")
	proto.RegisterType((*Equip)(nil), "message.equip")
	proto.RegisterType((*Resonance)(nil), "message.Resonance")
}

func init() { proto.RegisterFile("cs/proto/message/equipSync.proto", fileDescriptor_3793271f2ab7fe8f) }

var fileDescriptor_3793271f2ab7fe8f = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0xab, 0xd3, 0x30,
	0x14, 0xc7, 0x69, 0x6b, 0xb7, 0xee, 0x5c, 0x77, 0xef, 0x25, 0x88, 0x04, 0x7f, 0x5c, 0x4a, 0x11,
	0x29, 0x3e, 0xac, 0xe0, 0x83, 0xef, 0xae, 0xf5, 0xa1, 0x30, 0x5f, 0x32, 0x41, 0xf0, 0x45, 0x62,
	0x76, 0xd6, 0x85, 0x65, 0x4d, 0x4d, 0xea, 0xd0, 0x3f, 0xc3, 0xbf, 0xcb, 0x7f, 0x4a, 0x9a, 0x76,
	0xdd, 0xa6, 0x6f, 0xf9, 0x7e, 0x3e, 0x87, 0x93, 0x93, 0x43, 0x20, 0x16, 0x36, 0x6b, 0x8c, 0x6e,
	0x75, 0x76, 0x40, 0x6b, 0x79, 0x85, 0x19, 0x7e, 0xff, 0x21, 0x9b, 0xf5, 0xaf, 0x5a, 0x2c, 0x1c,
	0x27, 0xd3, 0x41, 0x24, 0x77, 0x30, 0x1f, 0xdd, 0xd7, 0x56, 0xaf, 0x93, 0x8f, 0xd7, 0x20, 0x27,
	0x4f, 0x20, 0x94, 0xf6, 0xbd, 0x52, 0xd4, 0x8b, 0xfd, 0x34, 0x62, 0x7d, 0x20, 0xaf, 0x61, 0xe2,
	0xca, 0x2c, 0xf5, 0xe3, 0x20, 0xbd, 0x79, 0x7b, 0xbb, 0x18, 0x3a, 0x2e, 0x1c, 0x66, 0x83, 0x4d,
	0xfe, 0xf8, 0x10, 0xba, 0x23, 0xb9, 0x05, 0xbf, 0x2c, 0x5c, 0x93, 0x39, 0xf3, 0xcb, 0x82, 0x3c,
	0x83, 0x28, 0xd7, 0xf5, 0x56, 0x56, 0x65, 0x41, 0xfd, 0xd8, 0x4b, 0x43, 0x36, 0xe6, 0xee, 0x4e,
	0x85, 0x47, 0x54, 0x34, 0x70, 0xa2, 0x0f, 0xe4, 0x1e, 0x02, 0xfc, 0xd9, 0xd0, 0x47, 0x8e, 0x75,
	0x47, 0xf2, 0x00, 0xf0, 0xcd, 0x20, 0xdf, 0xaf, 0x5c, 0x71, 0xe8, 0xc4, 0x05, 0x21, 0xef, 0xe0,
	0x31, 0x57, 0x8a, 0xa1, 0xd5, 0x35, 0xaf, 0x05, 0xd2, 0x89, 0x9b, 0x95, 0x8c, 0xb3, 0x8e, 0x86,
	0x5d, 0xd5, 0x91, 0x37, 0x70, 0xff, 0xa1, 0x1b, 0x3a, 0xdf, 0x71, 0xc3, 0x45, 0x8b, 0xa6, 0x2c,
	0xe8, 0xd4, 0x75, 0xff, 0x8f, 0x93, 0xa7, 0x30, 0x91, 0x76, 0xa5, 0xc5, 0x9e, 0x46, 0xb1, 0x97,
	0x46, 0x6c, 0x48, 0xdd, 0xfb, 0xa4, 0x65, 0x78, 0xd0, 0x47, 0xa4, 0x33, 0x67, 0xc6, 0x4c, 0x28,
	0x4c, 0x2b, 0x6c, 0x3f, 0xc9, 0x03, 0x52, 0x88, 0xbd, 0x34, 0x60, 0xa7, 0x48, 0x5e, 0xc0, 0x4c,
	0x1f, 0xd1, 0x08, 0xd5, 0x35, 0xbc, 0x89, 0x83, 0x34, 0x62, 0x67, 0x90, 0xfc, 0xf6, 0x60, 0x76,
	0x9e, 0xf2, 0x01, 0x60, 0xbd, 0x97, 0x4a, 0xe5, 0xdb, 0x6a, 0xd8, 0x6c, 0xc8, 0x2e, 0x48, 0xb7,
	0x45, 0x6d, 0x36, 0x68, 0xa8, 0xef, 0x54, 0x1f, 0x48, 0x0a, 0x77, 0x4b, 0x59, 0x6f, 0x2e, 0x9f,
	0xd6, 0x6f, 0xf9, 0x5f, 0x4c, 0x5e, 0xc1, 0x5c, 0xda, 0xcf, 0x5c, 0xb6, 0x0c, 0x1b, 0xc5, 0x05,
	0xba, 0xcd, 0x47, 0xec, 0x1a, 0x2e, 0x5f, 0x7e, 0x79, 0xde, 0x22, 0x17, 0x3b, 0x34, 0xfd, 0x9f,
	0x13, 0x5a, 0x65, 0xc2, 0x9e, 0x7e, 0xde, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xe2, 0x36,
	0xc5, 0x8c, 0x02, 0x00, 0x00,
}