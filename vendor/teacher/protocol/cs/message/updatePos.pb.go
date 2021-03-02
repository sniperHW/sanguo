// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/updatePos.proto

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

type UpdatePosToS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePosToS) Reset()         { *m = UpdatePosToS{} }
func (m *UpdatePosToS) String() string { return proto.CompactTextString(m) }
func (*UpdatePosToS) ProtoMessage()    {}
func (*UpdatePosToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bec299f3bd3f6396, []int{0}
}

func (m *UpdatePosToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePosToS.Unmarshal(m, b)
}
func (m *UpdatePosToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePosToS.Marshal(b, m, deterministic)
}
func (m *UpdatePosToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePosToS.Merge(m, src)
}
func (m *UpdatePosToS) XXX_Size() int {
	return xxx_messageInfo_UpdatePosToS.Size(m)
}
func (m *UpdatePosToS) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePosToS.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePosToS proto.InternalMessageInfo

type UpdateObj struct {
	RoleID               *uint64   `protobuf:"varint,1,opt,name=roleID" json:"roleID,omitempty"`
	Pos                  *Position `protobuf:"bytes,2,opt,name=pos" json:"pos,omitempty"`
	Angle                *int32    `protobuf:"varint,3,opt,name=angle" json:"angle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateObj) Reset()         { *m = UpdateObj{} }
func (m *UpdateObj) String() string { return proto.CompactTextString(m) }
func (*UpdateObj) ProtoMessage()    {}
func (*UpdateObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_bec299f3bd3f6396, []int{1}
}

func (m *UpdateObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateObj.Unmarshal(m, b)
}
func (m *UpdateObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateObj.Marshal(b, m, deterministic)
}
func (m *UpdateObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateObj.Merge(m, src)
}
func (m *UpdateObj) XXX_Size() int {
	return xxx_messageInfo_UpdateObj.Size(m)
}
func (m *UpdateObj) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateObj.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateObj proto.InternalMessageInfo

func (m *UpdateObj) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

func (m *UpdateObj) GetPos() *Position {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *UpdateObj) GetAngle() int32 {
	if m != nil && m.Angle != nil {
		return *m.Angle
	}
	return 0
}

type UpdatePosToC struct {
	Objs                 []*UpdateObj `protobuf:"bytes,1,rep,name=objs" json:"objs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdatePosToC) Reset()         { *m = UpdatePosToC{} }
func (m *UpdatePosToC) String() string { return proto.CompactTextString(m) }
func (*UpdatePosToC) ProtoMessage()    {}
func (*UpdatePosToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_bec299f3bd3f6396, []int{2}
}

func (m *UpdatePosToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePosToC.Unmarshal(m, b)
}
func (m *UpdatePosToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePosToC.Marshal(b, m, deterministic)
}
func (m *UpdatePosToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePosToC.Merge(m, src)
}
func (m *UpdatePosToC) XXX_Size() int {
	return xxx_messageInfo_UpdatePosToC.Size(m)
}
func (m *UpdatePosToC) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePosToC.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePosToC proto.InternalMessageInfo

func (m *UpdatePosToC) GetObjs() []*UpdateObj {
	if m != nil {
		return m.Objs
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdatePosToS)(nil), "message.updatePos_toS")
	proto.RegisterType((*UpdateObj)(nil), "message.updateObj")
	proto.RegisterType((*UpdatePosToC)(nil), "message.updatePos_toC")
}

func init() { proto.RegisterFile("cs/proto/message/updatePos.proto", fileDescriptor_bec299f3bd3f6396) }

var fileDescriptor_bec299f3bd3f6396 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2f, 0x2d, 0x48,
	0x49, 0x2c, 0x49, 0x0d, 0xc8, 0x2f, 0xd6, 0x03, 0x8b, 0x0b, 0xb1, 0x43, 0x25, 0xa4, 0x64, 0x31,
	0x94, 0x26, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0x41, 0xd4, 0x29, 0xf1, 0x73, 0xf1, 0xc2, 0xb5, 0xc6,
	0x97, 0xe4, 0x07, 0x2b, 0xc5, 0x71, 0x71, 0x42, 0x04, 0xfc, 0x93, 0xb2, 0x84, 0xc4, 0xb8, 0xd8,
	0x8a, 0xf2, 0x73, 0x52, 0x3d, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0xa0, 0x3c, 0x21,
	0x65, 0x2e, 0xe6, 0x82, 0xfc, 0x62, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x41, 0x3d, 0xa8,
	0xc9, 0x7a, 0x01, 0xf9, 0xc5, 0x99, 0x25, 0x99, 0xf9, 0x79, 0x41, 0x20, 0x59, 0x21, 0x11, 0x2e,
	0xd6, 0xc4, 0xbc, 0xf4, 0x9c, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0xc9,
	0x1c, 0xd5, 0x42, 0x67, 0x21, 0x35, 0x2e, 0x96, 0xfc, 0xa4, 0xac, 0x62, 0x09, 0x46, 0x05, 0x66,
	0x0d, 0x6e, 0x23, 0x21, 0xb8, 0x61, 0x70, 0x57, 0x04, 0x81, 0xe5, 0x9d, 0x64, 0xa3, 0xa4, 0x4b,
	0x52, 0x13, 0x93, 0x33, 0x52, 0x8b, 0x20, 0xfe, 0x49, 0xce, 0xcf, 0xd1, 0x4f, 0x2e, 0x86, 0xf9,
	0x0a, 0x10, 0x00, 0x00, 0xff, 0xff, 0xec, 0x4f, 0x56, 0x83, 0x13, 0x01, 0x00, 0x00,
}