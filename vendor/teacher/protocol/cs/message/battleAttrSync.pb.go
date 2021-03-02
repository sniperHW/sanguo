// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/battleAttrSync.proto

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

type BattleAttrSyncToS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BattleAttrSyncToS) Reset()         { *m = BattleAttrSyncToS{} }
func (m *BattleAttrSyncToS) String() string { return proto.CompactTextString(m) }
func (*BattleAttrSyncToS) ProtoMessage()    {}
func (*BattleAttrSyncToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_234a4a74487a8eda, []int{0}
}

func (m *BattleAttrSyncToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleAttrSyncToS.Unmarshal(m, b)
}
func (m *BattleAttrSyncToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleAttrSyncToS.Marshal(b, m, deterministic)
}
func (m *BattleAttrSyncToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleAttrSyncToS.Merge(m, src)
}
func (m *BattleAttrSyncToS) XXX_Size() int {
	return xxx_messageInfo_BattleAttrSyncToS.Size(m)
}
func (m *BattleAttrSyncToS) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleAttrSyncToS.DiscardUnknown(m)
}

var xxx_messageInfo_BattleAttrSyncToS proto.InternalMessageInfo

type BattleAttrSyncToC struct {
	IsAll                *bool              `protobuf:"varint,1,req,name=isAll" json:"isAll,omitempty"`
	BattleAttrSlice      []*BattleAttrSlice `protobuf:"bytes,2,rep,name=battleAttrSlice" json:"battleAttrSlice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BattleAttrSyncToC) Reset()         { *m = BattleAttrSyncToC{} }
func (m *BattleAttrSyncToC) String() string { return proto.CompactTextString(m) }
func (*BattleAttrSyncToC) ProtoMessage()    {}
func (*BattleAttrSyncToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_234a4a74487a8eda, []int{1}
}

func (m *BattleAttrSyncToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleAttrSyncToC.Unmarshal(m, b)
}
func (m *BattleAttrSyncToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleAttrSyncToC.Marshal(b, m, deterministic)
}
func (m *BattleAttrSyncToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleAttrSyncToC.Merge(m, src)
}
func (m *BattleAttrSyncToC) XXX_Size() int {
	return xxx_messageInfo_BattleAttrSyncToC.Size(m)
}
func (m *BattleAttrSyncToC) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleAttrSyncToC.DiscardUnknown(m)
}

var xxx_messageInfo_BattleAttrSyncToC proto.InternalMessageInfo

func (m *BattleAttrSyncToC) GetIsAll() bool {
	if m != nil && m.IsAll != nil {
		return *m.IsAll
	}
	return false
}

func (m *BattleAttrSyncToC) GetBattleAttrSlice() []*BattleAttrSlice {
	if m != nil {
		return m.BattleAttrSlice
	}
	return nil
}

type BattleAttrSlice struct {
	BattleAttrs          []*BattleAttr `protobuf:"bytes,1,rep,name=battleAttrs" json:"battleAttrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BattleAttrSlice) Reset()         { *m = BattleAttrSlice{} }
func (m *BattleAttrSlice) String() string { return proto.CompactTextString(m) }
func (*BattleAttrSlice) ProtoMessage()    {}
func (*BattleAttrSlice) Descriptor() ([]byte, []int) {
	return fileDescriptor_234a4a74487a8eda, []int{2}
}

func (m *BattleAttrSlice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleAttrSlice.Unmarshal(m, b)
}
func (m *BattleAttrSlice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleAttrSlice.Marshal(b, m, deterministic)
}
func (m *BattleAttrSlice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleAttrSlice.Merge(m, src)
}
func (m *BattleAttrSlice) XXX_Size() int {
	return xxx_messageInfo_BattleAttrSlice.Size(m)
}
func (m *BattleAttrSlice) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleAttrSlice.DiscardUnknown(m)
}

var xxx_messageInfo_BattleAttrSlice proto.InternalMessageInfo

func (m *BattleAttrSlice) GetBattleAttrs() []*BattleAttr {
	if m != nil {
		return m.BattleAttrs
	}
	return nil
}

func init() {
	proto.RegisterType((*BattleAttrSyncToS)(nil), "message.battleAttrSync_toS")
	proto.RegisterType((*BattleAttrSyncToC)(nil), "message.battleAttrSync_toC")
	proto.RegisterType((*BattleAttrSlice)(nil), "message.BattleAttrSlice")
}

func init() {
	proto.RegisterFile("cs/proto/message/battleAttrSync.proto", fileDescriptor_234a4a74487a8eda)
}

var fileDescriptor_234a4a74487a8eda = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x4f, 0x4a, 0x2c,
	0x29, 0xc9, 0x49, 0x75, 0x2c, 0x29, 0x29, 0x0a, 0xae, 0xcc, 0x4b, 0xd6, 0x03, 0x4b, 0x0a, 0xb1,
	0x43, 0x65, 0xa5, 0x64, 0x31, 0xd4, 0x27, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0x41, 0xd4, 0x29, 0x89,
	0x70, 0x09, 0xa1, 0xea, 0x8f, 0x2f, 0xc9, 0x0f, 0x56, 0xca, 0xc3, 0x22, 0xea, 0x2c, 0x24, 0xc2,
	0xc5, 0x9a, 0x59, 0xec, 0x98, 0x93, 0x23, 0xc1, 0xa8, 0xc0, 0xa4, 0xc1, 0x11, 0x04, 0xe1, 0x08,
	0x39, 0x71, 0xf1, 0x23, 0xa9, 0xcd, 0xc9, 0x4c, 0x4e, 0x95, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36,
	0x92, 0xd0, 0x83, 0xda, 0xa8, 0xe7, 0x84, 0x2a, 0x1f, 0x84, 0xae, 0x41, 0xc9, 0x83, 0x8b, 0x1f,
	0x4d, 0x8d, 0x90, 0x29, 0x17, 0x37, 0x42, 0x55, 0xb1, 0x04, 0x23, 0xd8, 0x48, 0x61, 0x2c, 0x46,
	0x06, 0x21, 0xab, 0x73, 0x92, 0x8d, 0x92, 0x2e, 0x49, 0x4d, 0x4c, 0xce, 0x48, 0x2d, 0x82, 0xf8,
	0x3a, 0x39, 0x3f, 0x47, 0x3f, 0xb9, 0x18, 0xe6, 0x77, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60,
	0xfc, 0xc6, 0xcc, 0x3e, 0x01, 0x00, 0x00,
}