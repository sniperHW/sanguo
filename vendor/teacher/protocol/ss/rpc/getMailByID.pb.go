// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/rpc/getMailByID.proto

package rpc

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

type GetMailByIDReq struct {
	BoxType              *int32   `protobuf:"varint,1,opt,name=boxType" json:"boxType,omitempty"`
	Owner                *string  `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
	MailID               *int64   `protobuf:"varint,3,opt,name=mailID" json:"mailID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMailByIDReq) Reset()         { *m = GetMailByIDReq{} }
func (m *GetMailByIDReq) String() string { return proto.CompactTextString(m) }
func (*GetMailByIDReq) ProtoMessage()    {}
func (*GetMailByIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_146729e2eb506229, []int{0}
}

func (m *GetMailByIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMailByIDReq.Unmarshal(m, b)
}
func (m *GetMailByIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMailByIDReq.Marshal(b, m, deterministic)
}
func (m *GetMailByIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMailByIDReq.Merge(m, src)
}
func (m *GetMailByIDReq) XXX_Size() int {
	return xxx_messageInfo_GetMailByIDReq.Size(m)
}
func (m *GetMailByIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMailByIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetMailByIDReq proto.InternalMessageInfo

func (m *GetMailByIDReq) GetBoxType() int32 {
	if m != nil && m.BoxType != nil {
		return *m.BoxType
	}
	return 0
}

func (m *GetMailByIDReq) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *GetMailByIDReq) GetMailID() int64 {
	if m != nil && m.MailID != nil {
		return *m.MailID
	}
	return 0
}

type GetMailByIDResp struct {
	ErrCode              *int32   `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	Mail                 *Mail    `protobuf:"bytes,2,opt,name=mail" json:"mail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMailByIDResp) Reset()         { *m = GetMailByIDResp{} }
func (m *GetMailByIDResp) String() string { return proto.CompactTextString(m) }
func (*GetMailByIDResp) ProtoMessage()    {}
func (*GetMailByIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_146729e2eb506229, []int{1}
}

func (m *GetMailByIDResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMailByIDResp.Unmarshal(m, b)
}
func (m *GetMailByIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMailByIDResp.Marshal(b, m, deterministic)
}
func (m *GetMailByIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMailByIDResp.Merge(m, src)
}
func (m *GetMailByIDResp) XXX_Size() int {
	return xxx_messageInfo_GetMailByIDResp.Size(m)
}
func (m *GetMailByIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMailByIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetMailByIDResp proto.InternalMessageInfo

func (m *GetMailByIDResp) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *GetMailByIDResp) GetMail() *Mail {
	if m != nil {
		return m.Mail
	}
	return nil
}

func init() {
	proto.RegisterType((*GetMailByIDReq)(nil), "rpc.getMailByID_req")
	proto.RegisterType((*GetMailByIDResp)(nil), "rpc.getMailByID_resp")
}

func init() { proto.RegisterFile("ss/proto/rpc/getMailByID.proto", fileDescriptor_146729e2eb506229) }

var fileDescriptor_146729e2eb506229 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0xd6, 0x4f, 0x4f, 0x2d, 0xf1, 0x4d, 0xcc, 0xcc,
	0x71, 0xaa, 0xf4, 0x74, 0xd1, 0x03, 0x8b, 0x0a, 0x31, 0x17, 0x15, 0x24, 0x4b, 0x61, 0x55, 0xe4,
	0x93, 0x59, 0x5c, 0x02, 0x51, 0xa4, 0x14, 0xc9, 0xc5, 0x8f, 0xa4, 0x33, 0xbe, 0x28, 0xb5, 0x50,
	0x48, 0x82, 0x8b, 0x3d, 0x29, 0xbf, 0x22, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x35, 0x08, 0xc6, 0x15, 0x12, 0xe1, 0x62, 0xcd, 0x2f, 0xcf, 0x4b, 0x2d, 0x92, 0x60, 0x52, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xc4, 0xb8, 0xd8, 0x72, 0x13, 0x33, 0x73, 0x3c, 0x5d, 0x24,
	0x98, 0x15, 0x18, 0x35, 0x98, 0x83, 0xa0, 0x3c, 0x25, 0x6f, 0x2e, 0x01, 0x54, 0xa3, 0x8b, 0x0b,
	0x40, 0x66, 0xa7, 0x16, 0x15, 0x39, 0xe7, 0xa7, 0xc0, 0xcd, 0x86, 0x72, 0x85, 0x64, 0xb9, 0x58,
	0x40, 0xfa, 0xc0, 0x46, 0x73, 0x1b, 0x71, 0xea, 0x15, 0x15, 0x24, 0xeb, 0x81, 0x04, 0x82, 0xc0,
	0xc2, 0x4e, 0x92, 0x51, 0xe2, 0x25, 0xa9, 0x89, 0xc9, 0x19, 0xa9, 0x45, 0x10, 0xef, 0x24, 0xe7,
	0xe7, 0xe8, 0x17, 0x17, 0x83, 0x3c, 0x05, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xf7, 0xfe, 0x98,
	0x08, 0x01, 0x00, 0x00,
}
