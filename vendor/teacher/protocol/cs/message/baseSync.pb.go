// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/baseSync.proto

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

type BaseSyncToS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseSyncToS) Reset()         { *m = BaseSyncToS{} }
func (m *BaseSyncToS) String() string { return proto.CompactTextString(m) }
func (*BaseSyncToS) ProtoMessage()    {}
func (*BaseSyncToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8efb5a5252997dc, []int{0}
}

func (m *BaseSyncToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseSyncToS.Unmarshal(m, b)
}
func (m *BaseSyncToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseSyncToS.Marshal(b, m, deterministic)
}
func (m *BaseSyncToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseSyncToS.Merge(m, src)
}
func (m *BaseSyncToS) XXX_Size() int {
	return xxx_messageInfo_BaseSyncToS.Size(m)
}
func (m *BaseSyncToS) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseSyncToS.DiscardUnknown(m)
}

var xxx_messageInfo_BaseSyncToS proto.InternalMessageInfo

type BaseSyncToC struct {
	UserID               *string  `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	GameID               *uint64  `protobuf:"varint,2,opt,name=gameID" json:"gameID,omitempty"`
	Name                 *string  `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseSyncToC) Reset()         { *m = BaseSyncToC{} }
func (m *BaseSyncToC) String() string { return proto.CompactTextString(m) }
func (*BaseSyncToC) ProtoMessage()    {}
func (*BaseSyncToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8efb5a5252997dc, []int{1}
}

func (m *BaseSyncToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseSyncToC.Unmarshal(m, b)
}
func (m *BaseSyncToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseSyncToC.Marshal(b, m, deterministic)
}
func (m *BaseSyncToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseSyncToC.Merge(m, src)
}
func (m *BaseSyncToC) XXX_Size() int {
	return xxx_messageInfo_BaseSyncToC.Size(m)
}
func (m *BaseSyncToC) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseSyncToC.DiscardUnknown(m)
}

var xxx_messageInfo_BaseSyncToC proto.InternalMessageInfo

func (m *BaseSyncToC) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *BaseSyncToC) GetGameID() uint64 {
	if m != nil && m.GameID != nil {
		return *m.GameID
	}
	return 0
}

func (m *BaseSyncToC) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseSyncToS)(nil), "message.baseSync_toS")
	proto.RegisterType((*BaseSyncToC)(nil), "message.baseSync_toC")
}

func init() { proto.RegisterFile("cs/proto/message/baseSync.proto", fileDescriptor_b8efb5a5252997dc) }

var fileDescriptor_b8efb5a5252997dc = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x4f, 0x4a, 0x2c,
	0x4e, 0x0d, 0xae, 0xcc, 0x4b, 0xd6, 0x03, 0x0b, 0x0b, 0xb1, 0x43, 0xc5, 0x95, 0xf8, 0xb8, 0x78,
	0x60, 0x52, 0xf1, 0x25, 0xf9, 0xc1, 0x4a, 0x41, 0x28, 0x7c, 0x67, 0x21, 0x31, 0x2e, 0xb6, 0xd2,
	0xe2, 0xd4, 0x22, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x0f, 0x24, 0x9e,
	0x9e, 0x98, 0x9b, 0xea, 0xe9, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x12, 0x04, 0xe5, 0x09, 0x09,
	0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x83, 0x55, 0x83, 0xd9, 0x4e, 0xb2, 0x51, 0xd2,
	0x25, 0xa9, 0x89, 0xc9, 0x19, 0xa9, 0x45, 0x10, 0x47, 0x25, 0xe7, 0xe7, 0xe8, 0x27, 0x17, 0xc3,
	0x9c, 0x06, 0x08, 0x00, 0x00, 0xff, 0xff, 0x37, 0x7f, 0xb7, 0xfb, 0xad, 0x00, 0x00, 0x00,
}