// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/createRole.proto

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

type CreateRoleToS struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoleToS) Reset()         { *m = CreateRoleToS{} }
func (m *CreateRoleToS) String() string { return proto.CompactTextString(m) }
func (*CreateRoleToS) ProtoMessage()    {}
func (*CreateRoleToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad4611ac9df85327, []int{0}
}

func (m *CreateRoleToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleToS.Unmarshal(m, b)
}
func (m *CreateRoleToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleToS.Marshal(b, m, deterministic)
}
func (m *CreateRoleToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleToS.Merge(m, src)
}
func (m *CreateRoleToS) XXX_Size() int {
	return xxx_messageInfo_CreateRoleToS.Size(m)
}
func (m *CreateRoleToS) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleToS.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleToS proto.InternalMessageInfo

func (m *CreateRoleToS) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type CreateRoleToC struct {
	NewRole              *Role    `protobuf:"bytes,2,opt,name=newRole" json:"newRole,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoleToC) Reset()         { *m = CreateRoleToC{} }
func (m *CreateRoleToC) String() string { return proto.CompactTextString(m) }
func (*CreateRoleToC) ProtoMessage()    {}
func (*CreateRoleToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad4611ac9df85327, []int{1}
}

func (m *CreateRoleToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleToC.Unmarshal(m, b)
}
func (m *CreateRoleToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleToC.Marshal(b, m, deterministic)
}
func (m *CreateRoleToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleToC.Merge(m, src)
}
func (m *CreateRoleToC) XXX_Size() int {
	return xxx_messageInfo_CreateRoleToC.Size(m)
}
func (m *CreateRoleToC) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleToC.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleToC proto.InternalMessageInfo

func (m *CreateRoleToC) GetNewRole() *Role {
	if m != nil {
		return m.NewRole
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRoleToS)(nil), "message.createRole_toS")
	proto.RegisterType((*CreateRoleToC)(nil), "message.createRole_toC")
}

func init() { proto.RegisterFile("cs/proto/message/createRole.proto", fileDescriptor_ad4611ac9df85327) }

var fileDescriptor_ad4611ac9df85327 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x4f, 0x2e, 0x4a,
	0x4d, 0x2c, 0x49, 0x0d, 0xca, 0xcf, 0x49, 0xd5, 0x03, 0x4b, 0x08, 0xb1, 0x43, 0x65, 0xa4, 0x64,
	0x31, 0xd5, 0xe6, 0xe7, 0xe6, 0xe6, 0xe7, 0x41, 0xd4, 0x29, 0xa9, 0x70, 0xf1, 0x21, 0xf4, 0xc6,
	0x97, 0xe4, 0x07, 0x0b, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0xd9, 0x4a, 0x96, 0x68, 0xaa, 0x9c, 0x85, 0xd4, 0xb9, 0xd8, 0xf3, 0x52, 0xcb,
	0x41, 0x5c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x5e, 0x3d, 0xa8, 0xf9, 0x7a, 0x45, 0xf9,
	0x39, 0xa9, 0x41, 0x30, 0x59, 0x27, 0xd9, 0x28, 0xe9, 0x92, 0xd4, 0xc4, 0xe4, 0x8c, 0xd4, 0x22,
	0x88, 0x33, 0x92, 0xf3, 0x73, 0xf4, 0x93, 0x8b, 0x61, 0x8e, 0x01, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x50, 0xd7, 0xdb, 0x9c, 0xcb, 0x00, 0x00, 0x00,
}
