// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/rpc/sendMail.proto

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

type SendMailReq struct {
	BoxType              *int32   `protobuf:"varint,1,opt,name=boxType" json:"boxType,omitempty"`
	Owner                *string  `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
	Id                   *int64   `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	Content              []byte   `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMailReq) Reset()         { *m = SendMailReq{} }
func (m *SendMailReq) String() string { return proto.CompactTextString(m) }
func (*SendMailReq) ProtoMessage()    {}
func (*SendMailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_226eba25479a7413, []int{0}
}

func (m *SendMailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMailReq.Unmarshal(m, b)
}
func (m *SendMailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMailReq.Marshal(b, m, deterministic)
}
func (m *SendMailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMailReq.Merge(m, src)
}
func (m *SendMailReq) XXX_Size() int {
	return xxx_messageInfo_SendMailReq.Size(m)
}
func (m *SendMailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMailReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMailReq proto.InternalMessageInfo

func (m *SendMailReq) GetBoxType() int32 {
	if m != nil && m.BoxType != nil {
		return *m.BoxType
	}
	return 0
}

func (m *SendMailReq) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *SendMailReq) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SendMailReq) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type SendMailResp struct {
	ErrCode              *int32   `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMailResp) Reset()         { *m = SendMailResp{} }
func (m *SendMailResp) String() string { return proto.CompactTextString(m) }
func (*SendMailResp) ProtoMessage()    {}
func (*SendMailResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_226eba25479a7413, []int{1}
}

func (m *SendMailResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMailResp.Unmarshal(m, b)
}
func (m *SendMailResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMailResp.Marshal(b, m, deterministic)
}
func (m *SendMailResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMailResp.Merge(m, src)
}
func (m *SendMailResp) XXX_Size() int {
	return xxx_messageInfo_SendMailResp.Size(m)
}
func (m *SendMailResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMailResp.DiscardUnknown(m)
}

var xxx_messageInfo_SendMailResp proto.InternalMessageInfo

func (m *SendMailResp) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func init() {
	proto.RegisterType((*SendMailReq)(nil), "rpc.sendMail_req")
	proto.RegisterType((*SendMailResp)(nil), "rpc.sendMail_resp")
}

func init() { proto.RegisterFile("ss/proto/rpc/sendMail.proto", fileDescriptor_226eba25479a7413) }

var fileDescriptor_226eba25479a7413 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x4e, 0xcd, 0x4b, 0xf1, 0x4d, 0xcc,
	0xcc, 0xd1, 0x03, 0x0b, 0x09, 0x31, 0x17, 0x15, 0x24, 0x2b, 0x65, 0x70, 0xf1, 0xc0, 0x84, 0xe3,
	0x8b, 0x52, 0x0b, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0xf2, 0x2b, 0x42, 0x2a, 0x0b, 0x52, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x58, 0x83, 0x60, 0x5c, 0x21, 0x11, 0x2e, 0xd6, 0xfc, 0xf2, 0xbc, 0xd4, 0x22,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82,
	0x59, 0x81, 0x51, 0x83, 0x39, 0x88, 0x29, 0x33, 0x05, 0xa4, 0x3f, 0x39, 0x3f, 0xaf, 0x24, 0x35,
	0xaf, 0x44, 0x82, 0x45, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc6, 0x55, 0xd2, 0xe4, 0xe2, 0x45, 0xb2,
	0xa9, 0xb8, 0x00, 0xa4, 0x34, 0xb5, 0xa8, 0xc8, 0x39, 0x3f, 0x05, 0x6e, 0x15, 0x94, 0xeb, 0x24,
	0x19, 0x25, 0x5e, 0x92, 0x9a, 0x98, 0x9c, 0x91, 0x5a, 0x04, 0x71, 0x7d, 0x72, 0x7e, 0x8e, 0x7e,
	0x71, 0x31, 0xc8, 0x0f, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x17, 0x28, 0x71, 0xd2, 0x00,
	0x00, 0x00,
}