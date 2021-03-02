// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/login.proto

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

type LoginToS struct {
	UserID               *string  `protobuf:"bytes,1,req,name=userID" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginToS) Reset()         { *m = LoginToS{} }
func (m *LoginToS) String() string { return proto.CompactTextString(m) }
func (*LoginToS) ProtoMessage()    {}
func (*LoginToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a1aa0076dca4aca, []int{0}
}

func (m *LoginToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginToS.Unmarshal(m, b)
}
func (m *LoginToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginToS.Marshal(b, m, deterministic)
}
func (m *LoginToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginToS.Merge(m, src)
}
func (m *LoginToS) XXX_Size() int {
	return xxx_messageInfo_LoginToS.Size(m)
}
func (m *LoginToS) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginToS.DiscardUnknown(m)
}

var xxx_messageInfo_LoginToS proto.InternalMessageInfo

func (m *LoginToS) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

type LoginToC struct {
	Game                 *string  `protobuf:"bytes,1,opt,name=game" json:"game,omitempty"`
	Token                *string  `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	UserID               *string  `protobuf:"bytes,3,opt,name=userID" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginToC) Reset()         { *m = LoginToC{} }
func (m *LoginToC) String() string { return proto.CompactTextString(m) }
func (*LoginToC) ProtoMessage()    {}
func (*LoginToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a1aa0076dca4aca, []int{1}
}

func (m *LoginToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginToC.Unmarshal(m, b)
}
func (m *LoginToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginToC.Marshal(b, m, deterministic)
}
func (m *LoginToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginToC.Merge(m, src)
}
func (m *LoginToC) XXX_Size() int {
	return xxx_messageInfo_LoginToC.Size(m)
}
func (m *LoginToC) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginToC.DiscardUnknown(m)
}

var xxx_messageInfo_LoginToC proto.InternalMessageInfo

func (m *LoginToC) GetGame() string {
	if m != nil && m.Game != nil {
		return *m.Game
	}
	return ""
}

func (m *LoginToC) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *LoginToC) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginToS)(nil), "message.login_toS")
	proto.RegisterType((*LoginToC)(nil), "message.login_toC")
}

func init() { proto.RegisterFile("cs/proto/message/login.proto", fileDescriptor_0a1aa0076dca4aca) }

var fileDescriptor_0a1aa0076dca4aca = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0xcf, 0xc9, 0x4f,
	0xcf, 0xcc, 0xd3, 0x03, 0x8b, 0x09, 0xb1, 0x43, 0x05, 0x95, 0x94, 0xb9, 0x38, 0xc1, 0xe2, 0xf1,
	0x25, 0xf9, 0xc1, 0x42, 0x62, 0x5c, 0x6c, 0xa5, 0xc5, 0xa9, 0x45, 0x9e, 0x2e, 0x12, 0x8c, 0x0a,
	0x4c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x92, 0x2f, 0x42, 0x91, 0xb3, 0x90, 0x10, 0x17, 0x4b, 0x7a,
	0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc2, 0xc5, 0x5a,
	0x92, 0x9f, 0x9d, 0x9a, 0x27, 0xc1, 0x04, 0x16, 0x84, 0x70, 0x90, 0x8c, 0x63, 0x06, 0x0b, 0x43,
	0x79, 0x4e, 0xb2, 0x51, 0xd2, 0x25, 0xa9, 0x89, 0xc9, 0x19, 0xa9, 0x45, 0x10, 0x17, 0x26, 0xe7,
	0xe7, 0xe8, 0x27, 0x17, 0xc3, 0xdc, 0x09, 0x08, 0x00, 0x00, 0xff, 0xff, 0x83, 0xf0, 0x8b, 0xe0,
	0xba, 0x00, 0x00, 0x00,
}
