// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/ssmessage/kickGameUser.proto

package ssmessage

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

type KickGameUser struct {
	UserID               *string  `protobuf:"bytes,1,req,name=userID" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KickGameUser) Reset()         { *m = KickGameUser{} }
func (m *KickGameUser) String() string { return proto.CompactTextString(m) }
func (*KickGameUser) ProtoMessage()    {}
func (*KickGameUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_0debfd6321def788, []int{0}
}

func (m *KickGameUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickGameUser.Unmarshal(m, b)
}
func (m *KickGameUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickGameUser.Marshal(b, m, deterministic)
}
func (m *KickGameUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickGameUser.Merge(m, src)
}
func (m *KickGameUser) XXX_Size() int {
	return xxx_messageInfo_KickGameUser.Size(m)
}
func (m *KickGameUser) XXX_DiscardUnknown() {
	xxx_messageInfo_KickGameUser.DiscardUnknown(m)
}

var xxx_messageInfo_KickGameUser proto.InternalMessageInfo

func (m *KickGameUser) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func init() {
	proto.RegisterType((*KickGameUser)(nil), "ssmessage.kickGameUser")
}

func init() {
	proto.RegisterFile("ss/proto/ssmessage/kickGameUser.proto", fileDescriptor_0debfd6321def788)
}

var fileDescriptor_0debfd6321def788 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0xce, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0xcf,
	0xce, 0x4c, 0xce, 0x76, 0x4f, 0xcc, 0x4d, 0x0d, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0x4b, 0x0a, 0x71,
	0xc2, 0x65, 0x95, 0xd4, 0xb8, 0x78, 0x90, 0x15, 0x08, 0x89, 0x71, 0xb1, 0x95, 0x16, 0xa7, 0x16,
	0x79, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x69, 0x70, 0x06, 0x41, 0x79, 0x4e, 0xf2, 0x51, 0xb2, 0x25,
	0xa9, 0x89, 0xc9, 0x19, 0xa9, 0x45, 0x10, 0x0b, 0x92, 0xf3, 0x73, 0xf4, 0x8b, 0x8b, 0x11, 0xd6,
	0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xfc, 0x77, 0xed, 0x7b, 0x00, 0x00, 0x00,
}
