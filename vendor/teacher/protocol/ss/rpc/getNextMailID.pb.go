// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/rpc/getNextMailID.proto

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

type GetNextMailIDReq struct {
	BoxType              *int32   `protobuf:"varint,1,opt,name=boxType" json:"boxType,omitempty"`
	Owner                *string  `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNextMailIDReq) Reset()         { *m = GetNextMailIDReq{} }
func (m *GetNextMailIDReq) String() string { return proto.CompactTextString(m) }
func (*GetNextMailIDReq) ProtoMessage()    {}
func (*GetNextMailIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4a95177f6a76e52, []int{0}
}

func (m *GetNextMailIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNextMailIDReq.Unmarshal(m, b)
}
func (m *GetNextMailIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNextMailIDReq.Marshal(b, m, deterministic)
}
func (m *GetNextMailIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNextMailIDReq.Merge(m, src)
}
func (m *GetNextMailIDReq) XXX_Size() int {
	return xxx_messageInfo_GetNextMailIDReq.Size(m)
}
func (m *GetNextMailIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNextMailIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetNextMailIDReq proto.InternalMessageInfo

func (m *GetNextMailIDReq) GetBoxType() int32 {
	if m != nil && m.BoxType != nil {
		return *m.BoxType
	}
	return 0
}

func (m *GetNextMailIDReq) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

type GetNextMailIDResp struct {
	ErrCode              *int32   `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	Id                   *int64   `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNextMailIDResp) Reset()         { *m = GetNextMailIDResp{} }
func (m *GetNextMailIDResp) String() string { return proto.CompactTextString(m) }
func (*GetNextMailIDResp) ProtoMessage()    {}
func (*GetNextMailIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4a95177f6a76e52, []int{1}
}

func (m *GetNextMailIDResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNextMailIDResp.Unmarshal(m, b)
}
func (m *GetNextMailIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNextMailIDResp.Marshal(b, m, deterministic)
}
func (m *GetNextMailIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNextMailIDResp.Merge(m, src)
}
func (m *GetNextMailIDResp) XXX_Size() int {
	return xxx_messageInfo_GetNextMailIDResp.Size(m)
}
func (m *GetNextMailIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNextMailIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetNextMailIDResp proto.InternalMessageInfo

func (m *GetNextMailIDResp) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *GetNextMailIDResp) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*GetNextMailIDReq)(nil), "rpc.getNextMailID_req")
	proto.RegisterType((*GetNextMailIDResp)(nil), "rpc.getNextMailID_resp")
}

func init() { proto.RegisterFile("ss/proto/rpc/getNextMailID.proto", fileDescriptor_b4a95177f6a76e52) }

var fileDescriptor_b4a95177f6a76e52 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0xd6, 0x4f, 0x4f, 0x2d, 0xf1, 0x4b, 0xad, 0x28,
	0xf1, 0x4d, 0xcc, 0xcc, 0xf1, 0x74, 0xd1, 0x03, 0x8b, 0x0b, 0x31, 0x17, 0x15, 0x24, 0x2b, 0x39,
	0x73, 0x09, 0xa2, 0xc8, 0xc5, 0x17, 0xa5, 0x16, 0x0a, 0x49, 0x70, 0xb1, 0x27, 0xe5, 0x57, 0x84,
	0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac,
	0xf9, 0xe5, 0x79, 0xa9, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x92, 0x1d,
	0x97, 0x10, 0xba, 0x21, 0xc5, 0x05, 0x20, 0x53, 0x52, 0x8b, 0x8a, 0x9c, 0xf3, 0x53, 0xe0, 0xa6,
	0x40, 0xb9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x60, 0x23, 0x98, 0x83, 0x98, 0x32, 0x53, 0x9c,
	0x24, 0xa3, 0xc4, 0x4b, 0x52, 0x13, 0x93, 0x33, 0x52, 0x8b, 0x20, 0x4e, 0x4e, 0xce, 0xcf, 0xd1,
	0x2f, 0x2e, 0x06, 0x39, 0x1c, 0x10, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x8c, 0xa4, 0x24, 0xc7, 0x00,
	0x00, 0x00,
}
