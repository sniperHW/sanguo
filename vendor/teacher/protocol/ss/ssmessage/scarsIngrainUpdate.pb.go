// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/ssmessage/scarsIngrainUpdate.proto

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

type ScarsIngrainUpdate struct {
	Version              *int64   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Id                   *int32   `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScarsIngrainUpdate) Reset()         { *m = ScarsIngrainUpdate{} }
func (m *ScarsIngrainUpdate) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainUpdate) ProtoMessage()    {}
func (*ScarsIngrainUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_08f22c2f07377120, []int{0}
}

func (m *ScarsIngrainUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainUpdate.Unmarshal(m, b)
}
func (m *ScarsIngrainUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainUpdate.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainUpdate.Merge(m, src)
}
func (m *ScarsIngrainUpdate) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainUpdate.Size(m)
}
func (m *ScarsIngrainUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainUpdate proto.InternalMessageInfo

func (m *ScarsIngrainUpdate) GetVersion() int64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *ScarsIngrainUpdate) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*ScarsIngrainUpdate)(nil), "ssmessage.scarsIngrainUpdate")
}

func init() {
	proto.RegisterFile("ss/proto/ssmessage/scarsIngrainUpdate.proto", fileDescriptor_08f22c2f07377120)
}

var fileDescriptor_08f22c2f07377120 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0xce, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2f,
	0x4e, 0x4e, 0x2c, 0x2a, 0xf6, 0xcc, 0x4b, 0x2f, 0x4a, 0xcc, 0xcc, 0x0b, 0x2d, 0x48, 0x49, 0x2c,
	0x49, 0xd5, 0x03, 0x2b, 0x11, 0xe2, 0x84, 0xab, 0x51, 0xb2, 0xe3, 0x12, 0xc2, 0x54, 0x26, 0x24,
	0xc1, 0xc5, 0x5e, 0x96, 0x5a, 0x54, 0x9c, 0x99, 0x9f, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1c,
	0x04, 0xe3, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x06, 0x31,
	0x65, 0xa6, 0x38, 0xc9, 0x47, 0xc9, 0x96, 0xa4, 0x26, 0x26, 0x67, 0xa4, 0x16, 0x41, 0xac, 0x4f,
	0xce, 0xcf, 0xd1, 0x2f, 0x2e, 0x46, 0x38, 0x02, 0x10, 0x00, 0x00, 0xff, 0xff, 0x27, 0x9b, 0xf2,
	0x85, 0x99, 0x00, 0x00, 0x00,
}
