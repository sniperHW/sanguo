// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/rpc/delMail.proto

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

type DelMail struct {
	Idx                  *int32   `protobuf:"varint,1,opt,name=idx" json:"idx,omitempty"`
	Version              *int64   `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelMail) Reset()         { *m = DelMail{} }
func (m *DelMail) String() string { return proto.CompactTextString(m) }
func (*DelMail) ProtoMessage()    {}
func (*DelMail) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebe74bd41565dbd, []int{0}
}

func (m *DelMail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelMail.Unmarshal(m, b)
}
func (m *DelMail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelMail.Marshal(b, m, deterministic)
}
func (m *DelMail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelMail.Merge(m, src)
}
func (m *DelMail) XXX_Size() int {
	return xxx_messageInfo_DelMail.Size(m)
}
func (m *DelMail) XXX_DiscardUnknown() {
	xxx_messageInfo_DelMail.DiscardUnknown(m)
}

var xxx_messageInfo_DelMail proto.InternalMessageInfo

func (m *DelMail) GetIdx() int32 {
	if m != nil && m.Idx != nil {
		return *m.Idx
	}
	return 0
}

func (m *DelMail) GetVersion() int64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type DelMailReq struct {
	BoxType              *int32     `protobuf:"varint,1,opt,name=boxType" json:"boxType,omitempty"`
	Owner                *string    `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
	Mails                []*DelMail `protobuf:"bytes,3,rep,name=mails" json:"mails,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DelMailReq) Reset()         { *m = DelMailReq{} }
func (m *DelMailReq) String() string { return proto.CompactTextString(m) }
func (*DelMailReq) ProtoMessage()    {}
func (*DelMailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebe74bd41565dbd, []int{1}
}

func (m *DelMailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelMailReq.Unmarshal(m, b)
}
func (m *DelMailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelMailReq.Marshal(b, m, deterministic)
}
func (m *DelMailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelMailReq.Merge(m, src)
}
func (m *DelMailReq) XXX_Size() int {
	return xxx_messageInfo_DelMailReq.Size(m)
}
func (m *DelMailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelMailReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelMailReq proto.InternalMessageInfo

func (m *DelMailReq) GetBoxType() int32 {
	if m != nil && m.BoxType != nil {
		return *m.BoxType
	}
	return 0
}

func (m *DelMailReq) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *DelMailReq) GetMails() []*DelMail {
	if m != nil {
		return m.Mails
	}
	return nil
}

type DelMailResp struct {
	ErrCode              *int32   `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	Version              *int64   `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelMailResp) Reset()         { *m = DelMailResp{} }
func (m *DelMailResp) String() string { return proto.CompactTextString(m) }
func (*DelMailResp) ProtoMessage()    {}
func (*DelMailResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebe74bd41565dbd, []int{2}
}

func (m *DelMailResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelMailResp.Unmarshal(m, b)
}
func (m *DelMailResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelMailResp.Marshal(b, m, deterministic)
}
func (m *DelMailResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelMailResp.Merge(m, src)
}
func (m *DelMailResp) XXX_Size() int {
	return xxx_messageInfo_DelMailResp.Size(m)
}
func (m *DelMailResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DelMailResp.DiscardUnknown(m)
}

var xxx_messageInfo_DelMailResp proto.InternalMessageInfo

func (m *DelMailResp) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *DelMailResp) GetVersion() int64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*DelMail)(nil), "rpc.delMail")
	proto.RegisterType((*DelMailReq)(nil), "rpc.delMail_req")
	proto.RegisterType((*DelMailResp)(nil), "rpc.delMail_resp")
}

func init() { proto.RegisterFile("ss/proto/rpc/delMail.proto", fileDescriptor_6ebe74bd41565dbd) }

var fileDescriptor_6ebe74bd41565dbd = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0xc7, 0x30,
	0x14, 0xc4, 0xa9, 0xa1, 0xfc, 0xf1, 0xb5, 0x83, 0x04, 0xc1, 0xe8, 0x54, 0x32, 0x75, 0x6a, 0x40,
	0xf0, 0x0b, 0xd4, 0xd9, 0x25, 0x38, 0xb9, 0x48, 0x4c, 0x1f, 0x18, 0x88, 0x4d, 0x7c, 0x29, 0x5a,
	0xbf, 0xbd, 0xa4, 0x4d, 0x75, 0xfa, 0x6f, 0xf9, 0x5d, 0x72, 0x77, 0xe1, 0xe0, 0x2e, 0x25, 0x15,
	0x29, 0x2c, 0x41, 0x51, 0xb4, 0x6a, 0x42, 0xff, 0x64, 0x9c, 0x1f, 0x36, 0x85, 0x33, 0x8a, 0x56,
	0x3e, 0xc0, 0xa9, 0xa8, 0xfc, 0x0a, 0x98, 0x9b, 0x56, 0x51, 0x75, 0x55, 0x5f, 0xeb, 0x7c, 0xe4,
	0x02, 0x4e, 0x5f, 0x48, 0xc9, 0x85, 0x59, 0x5c, 0x74, 0x55, 0xcf, 0xf4, 0x81, 0xd2, 0x40, 0x53,
	0x6c, 0xaf, 0x84, 0x9f, 0xf9, 0xe1, 0x5b, 0x58, 0x9f, 0x7f, 0x22, 0x16, 0xfb, 0x81, 0xfc, 0x1a,
	0xea, 0xf0, 0x3d, 0x23, 0x6d, 0x01, 0x97, 0x7a, 0x07, 0x2e, 0xa1, 0xfe, 0x30, 0xce, 0x27, 0xc1,
	0x3a, 0xd6, 0x37, 0xf7, 0xed, 0x40, 0xd1, 0x0e, 0x25, 0x50, 0xef, 0x57, 0x72, 0x84, 0xf6, 0xbf,
	0x22, 0xc5, 0xdc, 0x81, 0x44, 0x8f, 0x61, 0xfa, 0xeb, 0x28, 0x78, 0xfe, 0x9b, 0xe3, 0xed, 0xcb,
	0xcd, 0x82, 0xc6, 0xbe, 0x23, 0xed, 0x2b, 0xd8, 0xe0, 0x55, 0x4a, 0x79, 0x8b, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x87, 0x5e, 0xc5, 0x5e, 0x1a, 0x01, 0x00, 0x00,
}
