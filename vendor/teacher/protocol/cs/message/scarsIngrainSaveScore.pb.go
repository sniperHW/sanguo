// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/scarsIngrainSaveScore.proto

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

type ScarsIngrainSaveScoreToS struct {
	ScoreID              *int32   `protobuf:"varint,1,req,name=scoreID" json:"scoreID,omitempty"`
	IsSave               *bool    `protobuf:"varint,2,req,name=isSave" json:"isSave,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScarsIngrainSaveScoreToS) Reset()         { *m = ScarsIngrainSaveScoreToS{} }
func (m *ScarsIngrainSaveScoreToS) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainSaveScoreToS) ProtoMessage()    {}
func (*ScarsIngrainSaveScoreToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7df275104c1cac5d, []int{0}
}

func (m *ScarsIngrainSaveScoreToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainSaveScoreToS.Unmarshal(m, b)
}
func (m *ScarsIngrainSaveScoreToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainSaveScoreToS.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainSaveScoreToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainSaveScoreToS.Merge(m, src)
}
func (m *ScarsIngrainSaveScoreToS) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainSaveScoreToS.Size(m)
}
func (m *ScarsIngrainSaveScoreToS) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainSaveScoreToS.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainSaveScoreToS proto.InternalMessageInfo

func (m *ScarsIngrainSaveScoreToS) GetScoreID() int32 {
	if m != nil && m.ScoreID != nil {
		return *m.ScoreID
	}
	return 0
}

func (m *ScarsIngrainSaveScoreToS) GetIsSave() bool {
	if m != nil && m.IsSave != nil {
		return *m.IsSave
	}
	return false
}

type ScarsIngrainSaveScoreToC struct {
	Score                *int32   `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
	DropAwards           []*Award `protobuf:"bytes,2,rep,name=dropAwards" json:"dropAwards,omitempty"`
	RankIdx              *int32   `protobuf:"varint,3,opt,name=rankIdx" json:"rankIdx,omitempty"`
	RankPercent          *int32   `protobuf:"varint,4,opt,name=rankPercent" json:"rankPercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScarsIngrainSaveScoreToC) Reset()         { *m = ScarsIngrainSaveScoreToC{} }
func (m *ScarsIngrainSaveScoreToC) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainSaveScoreToC) ProtoMessage()    {}
func (*ScarsIngrainSaveScoreToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_7df275104c1cac5d, []int{1}
}

func (m *ScarsIngrainSaveScoreToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainSaveScoreToC.Unmarshal(m, b)
}
func (m *ScarsIngrainSaveScoreToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainSaveScoreToC.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainSaveScoreToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainSaveScoreToC.Merge(m, src)
}
func (m *ScarsIngrainSaveScoreToC) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainSaveScoreToC.Size(m)
}
func (m *ScarsIngrainSaveScoreToC) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainSaveScoreToC.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainSaveScoreToC proto.InternalMessageInfo

func (m *ScarsIngrainSaveScoreToC) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *ScarsIngrainSaveScoreToC) GetDropAwards() []*Award {
	if m != nil {
		return m.DropAwards
	}
	return nil
}

func (m *ScarsIngrainSaveScoreToC) GetRankIdx() int32 {
	if m != nil && m.RankIdx != nil {
		return *m.RankIdx
	}
	return 0
}

func (m *ScarsIngrainSaveScoreToC) GetRankPercent() int32 {
	if m != nil && m.RankPercent != nil {
		return *m.RankPercent
	}
	return 0
}

func init() {
	proto.RegisterType((*ScarsIngrainSaveScoreToS)(nil), "message.scarsIngrainSaveScore_toS")
	proto.RegisterType((*ScarsIngrainSaveScoreToC)(nil), "message.scarsIngrainSaveScore_toC")
}

func init() {
	proto.RegisterFile("cs/proto/message/scarsIngrainSaveScore.proto", fileDescriptor_7df275104c1cac5d)
}

var fileDescriptor_7df275104c1cac5d = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x6a, 0xac, 0x4c, 0xc1, 0xc3, 0x22, 0xb2, 0x2a, 0x85, 0xa5, 0xa7, 0x1c, 0x24,
	0x01, 0xdf, 0xc0, 0x3f, 0x97, 0x1c, 0x04, 0x49, 0x6e, 0x5e, 0x64, 0x99, 0x0c, 0xb5, 0x68, 0x76,
	0xca, 0x4c, 0x50, 0x5f, 0xc6, 0x77, 0x95, 0x6c, 0x12, 0x28, 0xa8, 0xb7, 0xfd, 0xcd, 0xfe, 0xe6,
	0x63, 0xf8, 0xe0, 0x1a, 0xb5, 0xdc, 0x0b, 0xf7, 0x5c, 0x76, 0xa4, 0xea, 0xb7, 0x54, 0x2a, 0x7a,
	0xd1, 0x2a, 0x6c, 0xc5, 0xef, 0x42, 0xe3, 0x3f, 0xa8, 0x41, 0x16, 0x2a, 0xa2, 0x63, 0x96, 0x93,
	0x74, 0xb9, 0xfe, 0xb5, 0x86, 0xdc, 0x75, 0x1c, 0x46, 0x6f, 0xf3, 0x08, 0x17, 0x7f, 0xc6, 0xbc,
	0xf4, 0xdc, 0x18, 0x0b, 0x4b, 0x1d, 0xa0, 0x7a, 0xb0, 0x89, 0x4b, 0xf3, 0xac, 0x9e, 0xd1, 0x9c,
	0xc3, 0xf1, 0x4e, 0x07, 0xd9, 0xa6, 0x2e, 0xcd, 0x4f, 0xea, 0x89, 0x36, 0xdf, 0xc9, 0xff, 0x79,
	0xf7, 0xe6, 0x0c, 0xb2, 0x18, 0x60, 0x13, 0x97, 0xe4, 0x59, 0x3d, 0x82, 0x29, 0x00, 0x5a, 0xe1,
	0xfd, 0xed, 0xa7, 0x97, 0x56, 0x6d, 0xea, 0x16, 0xf9, 0xea, 0xe6, 0xb4, 0x98, 0xae, 0x2d, 0xe2,
	0xb8, 0x3e, 0x30, 0x86, 0xab, 0xc4, 0x87, 0xb7, 0xaa, 0xfd, 0xb2, 0x8b, 0x98, 0x33, 0xa3, 0x71,
	0xb0, 0x1a, 0x9e, 0x4f, 0x24, 0x48, 0xa1, 0xb7, 0x47, 0xf1, 0xf7, 0x70, 0x74, 0xb7, 0x7e, 0xbe,
	0xea, 0xc9, 0xe3, 0x2b, 0xc9, 0x58, 0x0a, 0xf2, 0x7b, 0x89, 0x3a, 0x57, 0xf3, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0x50, 0x0b, 0x7a, 0x64, 0x01, 0x00, 0x00,
}
