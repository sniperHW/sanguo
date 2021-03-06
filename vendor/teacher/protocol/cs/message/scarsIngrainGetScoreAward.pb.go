// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/scarsIngrainGetScoreAward.proto

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

type ScarsIngrainGetScoreAwardToS struct {
	Score                []int32  `protobuf:"varint,1,rep,name=score" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScarsIngrainGetScoreAwardToS) Reset()         { *m = ScarsIngrainGetScoreAwardToS{} }
func (m *ScarsIngrainGetScoreAwardToS) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainGetScoreAwardToS) ProtoMessage()    {}
func (*ScarsIngrainGetScoreAwardToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dcacb693c4da021, []int{0}
}

func (m *ScarsIngrainGetScoreAwardToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainGetScoreAwardToS.Unmarshal(m, b)
}
func (m *ScarsIngrainGetScoreAwardToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainGetScoreAwardToS.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainGetScoreAwardToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainGetScoreAwardToS.Merge(m, src)
}
func (m *ScarsIngrainGetScoreAwardToS) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainGetScoreAwardToS.Size(m)
}
func (m *ScarsIngrainGetScoreAwardToS) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainGetScoreAwardToS.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainGetScoreAwardToS proto.InternalMessageInfo

func (m *ScarsIngrainGetScoreAwardToS) GetScore() []int32 {
	if m != nil {
		return m.Score
	}
	return nil
}

type ScarsIngrainGetScoreAwardToC struct {
	DropAwards           []*Award `protobuf:"bytes,1,rep,name=dropAwards" json:"dropAwards,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScarsIngrainGetScoreAwardToC) Reset()         { *m = ScarsIngrainGetScoreAwardToC{} }
func (m *ScarsIngrainGetScoreAwardToC) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainGetScoreAwardToC) ProtoMessage()    {}
func (*ScarsIngrainGetScoreAwardToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dcacb693c4da021, []int{1}
}

func (m *ScarsIngrainGetScoreAwardToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainGetScoreAwardToC.Unmarshal(m, b)
}
func (m *ScarsIngrainGetScoreAwardToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainGetScoreAwardToC.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainGetScoreAwardToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainGetScoreAwardToC.Merge(m, src)
}
func (m *ScarsIngrainGetScoreAwardToC) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainGetScoreAwardToC.Size(m)
}
func (m *ScarsIngrainGetScoreAwardToC) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainGetScoreAwardToC.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainGetScoreAwardToC proto.InternalMessageInfo

func (m *ScarsIngrainGetScoreAwardToC) GetDropAwards() []*Award {
	if m != nil {
		return m.DropAwards
	}
	return nil
}

func init() {
	proto.RegisterType((*ScarsIngrainGetScoreAwardToS)(nil), "message.scarsIngrainGetScoreAward_toS")
	proto.RegisterType((*ScarsIngrainGetScoreAwardToC)(nil), "message.scarsIngrainGetScoreAward_toC")
}

func init() {
	proto.RegisterFile("cs/proto/message/scarsIngrainGetScoreAward.proto", fileDescriptor_3dcacb693c4da021)
}

var fileDescriptor_3dcacb693c4da021 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2f, 0x4e, 0x4e,
	0x2c, 0x2a, 0xf6, 0xcc, 0x4b, 0x2f, 0x4a, 0xcc, 0xcc, 0x73, 0x4f, 0x2d, 0x09, 0x4e, 0xce, 0x2f,
	0x4a, 0x75, 0x2c, 0x4f, 0x2c, 0x4a, 0xd1, 0x03, 0xab, 0x13, 0x62, 0x87, 0x2a, 0x94, 0x92, 0xc5,
	0xd0, 0x9a, 0x9c, 0x9f, 0x9b, 0x9b, 0x9f, 0x07, 0x51, 0xa7, 0x64, 0xca, 0x25, 0x8b, 0xd3, 0xa8,
	0xf8, 0x92, 0xfc, 0x60, 0x21, 0x11, 0x2e, 0xd6, 0x62, 0x90, 0x88, 0x04, 0xa3, 0x02, 0xb3, 0x06,
	0x6b, 0x10, 0x84, 0xa3, 0xe4, 0x8f, 0x5f, 0x9b, 0xb3, 0x90, 0x1e, 0x17, 0x57, 0x4a, 0x51, 0x7e,
	0x01, 0x58, 0xa0, 0x18, 0xac, 0x97, 0xdb, 0x88, 0x4f, 0x0f, 0xea, 0x04, 0x3d, 0xb0, 0x70, 0x10,
	0x92, 0x0a, 0x27, 0xd9, 0x28, 0xe9, 0x92, 0xd4, 0xc4, 0xe4, 0x8c, 0xd4, 0x22, 0x88, 0x6b, 0x93,
	0xf3, 0x73, 0xf4, 0x93, 0x8b, 0x61, 0x6e, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x38, 0x55, 0x54,
	0xb1, 0x01, 0x01, 0x00, 0x00,
}
