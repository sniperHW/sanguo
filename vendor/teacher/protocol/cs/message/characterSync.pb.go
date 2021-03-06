// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/characterSync.proto

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

type CharacterSyncToS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CharacterSyncToS) Reset()         { *m = CharacterSyncToS{} }
func (m *CharacterSyncToS) String() string { return proto.CompactTextString(m) }
func (*CharacterSyncToS) ProtoMessage()    {}
func (*CharacterSyncToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5896d908b09f87b9, []int{0}
}

func (m *CharacterSyncToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharacterSyncToS.Unmarshal(m, b)
}
func (m *CharacterSyncToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharacterSyncToS.Marshal(b, m, deterministic)
}
func (m *CharacterSyncToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterSyncToS.Merge(m, src)
}
func (m *CharacterSyncToS) XXX_Size() int {
	return xxx_messageInfo_CharacterSyncToS.Size(m)
}
func (m *CharacterSyncToS) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterSyncToS.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterSyncToS proto.InternalMessageInfo

type CharacterSyncToC struct {
	IsAll                *bool        `protobuf:"varint,1,req,name=isAll" json:"isAll,omitempty"`
	Characters           []*Character `protobuf:"bytes,2,rep,name=characters" json:"characters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CharacterSyncToC) Reset()         { *m = CharacterSyncToC{} }
func (m *CharacterSyncToC) String() string { return proto.CompactTextString(m) }
func (*CharacterSyncToC) ProtoMessage()    {}
func (*CharacterSyncToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_5896d908b09f87b9, []int{1}
}

func (m *CharacterSyncToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharacterSyncToC.Unmarshal(m, b)
}
func (m *CharacterSyncToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharacterSyncToC.Marshal(b, m, deterministic)
}
func (m *CharacterSyncToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterSyncToC.Merge(m, src)
}
func (m *CharacterSyncToC) XXX_Size() int {
	return xxx_messageInfo_CharacterSyncToC.Size(m)
}
func (m *CharacterSyncToC) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterSyncToC.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterSyncToC proto.InternalMessageInfo

func (m *CharacterSyncToC) GetIsAll() bool {
	if m != nil && m.IsAll != nil {
		return *m.IsAll
	}
	return false
}

func (m *CharacterSyncToC) GetCharacters() []*Character {
	if m != nil {
		return m.Characters
	}
	return nil
}

type Character struct {
	CharacterID          *int32   `protobuf:"varint,1,opt,name=CharacterID" json:"CharacterID,omitempty"`
	Level                *int32   `protobuf:"varint,2,opt,name=Level" json:"Level,omitempty"`
	CurrentExp           *int32   `protobuf:"varint,3,opt,name=CurrentExp" json:"CurrentExp,omitempty"`
	Status               *int32   `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	BreakLevel           *int32   `protobuf:"varint,5,opt,name=BreakLevel" json:"BreakLevel,omitempty"`
	Promote              *int32   `protobuf:"varint,6,opt,name=Promote" json:"Promote,omitempty"`
	Liberation           *int32   `protobuf:"varint,7,opt,name=Liberation" json:"Liberation,omitempty"`
	Evolution            *int32   `protobuf:"varint,8,opt,name=Evolution" json:"Evolution,omitempty"`
	Rarity               *int32   `protobuf:"varint,9,opt,name=Rarity" json:"Rarity,omitempty"`
	EquipIDs             []uint32 `protobuf:"varint,10,rep,name=EquipIDs" json:"EquipIDs,omitempty"`
	GetTime              *int64   `protobuf:"varint,11,opt,name=getTime" json:"getTime,omitempty"`
	FavorLevel           *int32   `protobuf:"varint,12,opt,name=favorLevel" json:"favorLevel,omitempty"`
	FavorExp             *int32   `protobuf:"varint,13,opt,name=favorExp" json:"favorExp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Character) Reset()         { *m = Character{} }
func (m *Character) String() string { return proto.CompactTextString(m) }
func (*Character) ProtoMessage()    {}
func (*Character) Descriptor() ([]byte, []int) {
	return fileDescriptor_5896d908b09f87b9, []int{2}
}

func (m *Character) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Character.Unmarshal(m, b)
}
func (m *Character) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Character.Marshal(b, m, deterministic)
}
func (m *Character) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Character.Merge(m, src)
}
func (m *Character) XXX_Size() int {
	return xxx_messageInfo_Character.Size(m)
}
func (m *Character) XXX_DiscardUnknown() {
	xxx_messageInfo_Character.DiscardUnknown(m)
}

var xxx_messageInfo_Character proto.InternalMessageInfo

func (m *Character) GetCharacterID() int32 {
	if m != nil && m.CharacterID != nil {
		return *m.CharacterID
	}
	return 0
}

func (m *Character) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Character) GetCurrentExp() int32 {
	if m != nil && m.CurrentExp != nil {
		return *m.CurrentExp
	}
	return 0
}

func (m *Character) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *Character) GetBreakLevel() int32 {
	if m != nil && m.BreakLevel != nil {
		return *m.BreakLevel
	}
	return 0
}

func (m *Character) GetPromote() int32 {
	if m != nil && m.Promote != nil {
		return *m.Promote
	}
	return 0
}

func (m *Character) GetLiberation() int32 {
	if m != nil && m.Liberation != nil {
		return *m.Liberation
	}
	return 0
}

func (m *Character) GetEvolution() int32 {
	if m != nil && m.Evolution != nil {
		return *m.Evolution
	}
	return 0
}

func (m *Character) GetRarity() int32 {
	if m != nil && m.Rarity != nil {
		return *m.Rarity
	}
	return 0
}

func (m *Character) GetEquipIDs() []uint32 {
	if m != nil {
		return m.EquipIDs
	}
	return nil
}

func (m *Character) GetGetTime() int64 {
	if m != nil && m.GetTime != nil {
		return *m.GetTime
	}
	return 0
}

func (m *Character) GetFavorLevel() int32 {
	if m != nil && m.FavorLevel != nil {
		return *m.FavorLevel
	}
	return 0
}

func (m *Character) GetFavorExp() int32 {
	if m != nil && m.FavorExp != nil {
		return *m.FavorExp
	}
	return 0
}

func init() {
	proto.RegisterType((*CharacterSyncToS)(nil), "message.characterSync_toS")
	proto.RegisterType((*CharacterSyncToC)(nil), "message.characterSync_toC")
	proto.RegisterType((*Character)(nil), "message.character")
}

func init() {
	proto.RegisterFile("cs/proto/message/characterSync.proto", fileDescriptor_5896d908b09f87b9)
}

var fileDescriptor_5896d908b09f87b9 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4b, 0x4b, 0x33, 0x31,
	0x14, 0x86, 0x99, 0x99, 0xaf, 0xb7, 0xd3, 0xaf, 0x0b, 0xa3, 0x48, 0xf0, 0xc6, 0x50, 0x5c, 0xcc,
	0xaa, 0x03, 0xfd, 0x07, 0xf6, 0xb2, 0x28, 0x74, 0x21, 0x53, 0x57, 0x82, 0x48, 0x0c, 0xc7, 0x36,
	0x38, 0x6d, 0x6a, 0x92, 0x29, 0xf6, 0xc7, 0x0b, 0x92, 0xa4, 0x73, 0x11, 0x97, 0xcf, 0xfb, 0xe4,
	0x9c, 0x79, 0x27, 0x04, 0xee, 0xb9, 0x4e, 0xf7, 0x4a, 0x1a, 0x99, 0x6e, 0x51, 0x6b, 0xb6, 0xc6,
	0x94, 0x6f, 0x98, 0x62, 0xdc, 0xa0, 0x5a, 0x1d, 0x77, 0x7c, 0xe4, 0x1c, 0xe9, 0x9c, 0xe4, 0xf0,
	0x1c, 0xce, 0x7e, 0xf9, 0x57, 0x23, 0x57, 0xc3, 0x97, 0xbf, 0xe1, 0x94, 0x5c, 0x40, 0x4b, 0xe8,
	0x87, 0x3c, 0xa7, 0x41, 0x1c, 0x26, 0xdd, 0xcc, 0x03, 0x19, 0x03, 0x54, 0x47, 0x35, 0x0d, 0xe3,
	0x28, 0xe9, 0x8f, 0xc9, 0xe8, 0xb4, 0x7d, 0x54, 0xa9, 0xac, 0x71, 0x6a, 0xf8, 0x1d, 0x42, 0xaf,
	0x42, 0x12, 0x43, 0x7f, 0x5a, 0xc2, 0x62, 0x46, 0x83, 0x38, 0x48, 0x5a, 0x59, 0x33, 0xb2, 0x5f,
	0x5e, 0xe2, 0x01, 0x73, 0x1a, 0x3a, 0xe7, 0x81, 0xdc, 0x01, 0x4c, 0x0b, 0xa5, 0x70, 0x67, 0xe6,
	0x5f, 0x7b, 0x1a, 0x39, 0xd5, 0x48, 0xc8, 0x25, 0xb4, 0x57, 0x86, 0x99, 0x42, 0xd3, 0x7f, 0xce,
	0x9d, 0xc8, 0xce, 0x4d, 0x14, 0xb2, 0x0f, 0xbf, 0xb2, 0xe5, 0xe7, 0xea, 0x84, 0x50, 0xe8, 0x3c,
	0x2a, 0xb9, 0x95, 0x06, 0x69, 0xdb, 0xc9, 0x12, 0xed, 0xe4, 0x52, 0xbc, 0xa1, 0x62, 0x46, 0xc8,
	0x1d, 0xed, 0xf8, 0xc9, 0x3a, 0x21, 0x37, 0xd0, 0x9b, 0x1f, 0x64, 0x5e, 0x38, 0xdd, 0x75, 0xba,
	0x0e, 0x6c, 0x9f, 0x8c, 0x29, 0x61, 0x8e, 0xb4, 0xe7, 0xfb, 0x78, 0x22, 0x57, 0xd0, 0x9d, 0x7f,
	0x16, 0x62, 0xbf, 0x98, 0x69, 0x0a, 0x71, 0x94, 0x0c, 0xb2, 0x8a, 0x6d, 0x97, 0x35, 0x9a, 0x27,
	0xb1, 0x45, 0xda, 0x8f, 0x83, 0x24, 0xca, 0x4a, 0xb4, 0x5d, 0xde, 0xd9, 0x41, 0x2a, 0xff, 0x17,
	0xff, 0x7d, 0x97, 0x3a, 0xb1, 0x5b, 0x1d, 0xd9, 0xbb, 0x19, 0x38, 0x5b, 0xf1, 0xe4, 0xf6, 0xf9,
	0xda, 0x20, 0xe3, 0x1b, 0x54, 0xfe, 0xa5, 0x70, 0x99, 0xa7, 0x5c, 0x97, 0xef, 0xe5, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xbf, 0x38, 0xbb, 0x0c, 0x42, 0x02, 0x00, 0x00,
}
