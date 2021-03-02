// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/scarsIngrainAutoFight.proto

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

type ScarsIngrainAutoFightToS struct {
	DungeonId            *int32   `protobuf:"varint,1,req,name=dungeonId" json:"dungeonId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScarsIngrainAutoFightToS) Reset()         { *m = ScarsIngrainAutoFightToS{} }
func (m *ScarsIngrainAutoFightToS) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainAutoFightToS) ProtoMessage()    {}
func (*ScarsIngrainAutoFightToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_14c731d72dc806f2, []int{0}
}

func (m *ScarsIngrainAutoFightToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainAutoFightToS.Unmarshal(m, b)
}
func (m *ScarsIngrainAutoFightToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainAutoFightToS.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainAutoFightToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainAutoFightToS.Merge(m, src)
}
func (m *ScarsIngrainAutoFightToS) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainAutoFightToS.Size(m)
}
func (m *ScarsIngrainAutoFightToS) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainAutoFightToS.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainAutoFightToS proto.InternalMessageInfo

func (m *ScarsIngrainAutoFightToS) GetDungeonId() int32 {
	if m != nil && m.DungeonId != nil {
		return *m.DungeonId
	}
	return 0
}

type ScarsIngrainAutoFightToC struct {
	Score                *int32   `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
	DropAwards           []*Award `protobuf:"bytes,2,rep,name=dropAwards" json:"dropAwards,omitempty"`
	RankIdx              *int32   `protobuf:"varint,3,opt,name=rankIdx" json:"rankIdx,omitempty"`
	RankPercent          *int32   `protobuf:"varint,4,opt,name=rankPercent" json:"rankPercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScarsIngrainAutoFightToC) Reset()         { *m = ScarsIngrainAutoFightToC{} }
func (m *ScarsIngrainAutoFightToC) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainAutoFightToC) ProtoMessage()    {}
func (*ScarsIngrainAutoFightToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_14c731d72dc806f2, []int{1}
}

func (m *ScarsIngrainAutoFightToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainAutoFightToC.Unmarshal(m, b)
}
func (m *ScarsIngrainAutoFightToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainAutoFightToC.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainAutoFightToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainAutoFightToC.Merge(m, src)
}
func (m *ScarsIngrainAutoFightToC) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainAutoFightToC.Size(m)
}
func (m *ScarsIngrainAutoFightToC) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainAutoFightToC.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainAutoFightToC proto.InternalMessageInfo

func (m *ScarsIngrainAutoFightToC) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *ScarsIngrainAutoFightToC) GetDropAwards() []*Award {
	if m != nil {
		return m.DropAwards
	}
	return nil
}

func (m *ScarsIngrainAutoFightToC) GetRankIdx() int32 {
	if m != nil && m.RankIdx != nil {
		return *m.RankIdx
	}
	return 0
}

func (m *ScarsIngrainAutoFightToC) GetRankPercent() int32 {
	if m != nil && m.RankPercent != nil {
		return *m.RankPercent
	}
	return 0
}

func init() {
	proto.RegisterType((*ScarsIngrainAutoFightToS)(nil), "message.scarsIngrainAutoFight_toS")
	proto.RegisterType((*ScarsIngrainAutoFightToC)(nil), "message.scarsIngrainAutoFight_toC")
}

func init() {
	proto.RegisterFile("cs/proto/message/scarsIngrainAutoFight.proto", fileDescriptor_14c731d72dc806f2)
}

var fileDescriptor_14c731d72dc806f2 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2f, 0x4e, 0x4e,
	0x2c, 0x2a, 0xf6, 0xcc, 0x4b, 0x2f, 0x4a, 0xcc, 0xcc, 0x73, 0x2c, 0x2d, 0xc9, 0x77, 0xcb, 0x4c,
	0xcf, 0x28, 0xd1, 0x03, 0xab, 0x11, 0x62, 0x87, 0x2a, 0x92, 0x92, 0xc5, 0xd0, 0x96, 0x9c, 0x9f,
	0x9b, 0x9b, 0x9f, 0x07, 0x51, 0xa7, 0x64, 0xc9, 0x25, 0x89, 0xd5, 0x98, 0xf8, 0x92, 0xfc, 0x60,
	0x21, 0x19, 0x2e, 0xce, 0x94, 0xd2, 0xbc, 0xf4, 0xd4, 0xfc, 0x3c, 0xcf, 0x14, 0x09, 0x46, 0x05,
	0x26, 0x0d, 0xd6, 0x20, 0x84, 0x80, 0xd2, 0x5c, 0x46, 0xdc, 0x7a, 0x9d, 0x85, 0x44, 0xb8, 0x58,
	0x8b, 0x93, 0xf3, 0x8b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c, 0x21, 0x3d,
	0x2e, 0xae, 0x94, 0xa2, 0xfc, 0x02, 0xc7, 0xf2, 0xc4, 0xa2, 0x94, 0x62, 0x09, 0x26, 0x05, 0x66,
	0x0d, 0x6e, 0x23, 0x3e, 0x3d, 0xa8, 0xcb, 0xf4, 0xc0, 0xc2, 0x41, 0x48, 0x2a, 0x84, 0x24, 0xb8,
	0xd8, 0x8b, 0x12, 0xf3, 0xb2, 0x3d, 0x53, 0x2a, 0x24, 0x98, 0xc1, 0xe6, 0xc0, 0xb8, 0x42, 0x0a,
	0x5c, 0xdc, 0x20, 0x66, 0x40, 0x6a, 0x51, 0x72, 0x6a, 0x5e, 0x89, 0x04, 0x0b, 0x58, 0x16, 0x59,
	0xc8, 0x49, 0x36, 0x4a, 0xba, 0x24, 0x35, 0x31, 0x39, 0x23, 0xb5, 0x08, 0x12, 0x00, 0xc9, 0xf9,
	0x39, 0xfa, 0xc9, 0xc5, 0xb0, 0x60, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0x71, 0x59, 0x83, 0x90,
	0x50, 0x01, 0x00, 0x00,
}