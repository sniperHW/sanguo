// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/rpc/worldObjDo.proto

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

type WorldObjDoReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorldObjDoReq) Reset()         { *m = WorldObjDoReq{} }
func (m *WorldObjDoReq) String() string { return proto.CompactTextString(m) }
func (*WorldObjDoReq) ProtoMessage()    {}
func (*WorldObjDoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b50e0a97f48fc6a, []int{0}
}

func (m *WorldObjDoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorldObjDoReq.Unmarshal(m, b)
}
func (m *WorldObjDoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorldObjDoReq.Marshal(b, m, deterministic)
}
func (m *WorldObjDoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorldObjDoReq.Merge(m, src)
}
func (m *WorldObjDoReq) XXX_Size() int {
	return xxx_messageInfo_WorldObjDoReq.Size(m)
}
func (m *WorldObjDoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_WorldObjDoReq.DiscardUnknown(m)
}

var xxx_messageInfo_WorldObjDoReq proto.InternalMessageInfo

type WorldObjDoResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorldObjDoResp) Reset()         { *m = WorldObjDoResp{} }
func (m *WorldObjDoResp) String() string { return proto.CompactTextString(m) }
func (*WorldObjDoResp) ProtoMessage()    {}
func (*WorldObjDoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b50e0a97f48fc6a, []int{1}
}

func (m *WorldObjDoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorldObjDoResp.Unmarshal(m, b)
}
func (m *WorldObjDoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorldObjDoResp.Marshal(b, m, deterministic)
}
func (m *WorldObjDoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorldObjDoResp.Merge(m, src)
}
func (m *WorldObjDoResp) XXX_Size() int {
	return xxx_messageInfo_WorldObjDoResp.Size(m)
}
func (m *WorldObjDoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_WorldObjDoResp.DiscardUnknown(m)
}

var xxx_messageInfo_WorldObjDoResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WorldObjDoReq)(nil), "rpc.worldObjDo_req")
	proto.RegisterType((*WorldObjDoResp)(nil), "rpc.worldObjDo_resp")
}

func init() { proto.RegisterFile("ss/proto/rpc/worldObjDo.proto", fileDescriptor_9b50e0a97f48fc6a) }

var fileDescriptor_9b50e0a97f48fc6a = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xf1, 0x4f,
	0xca, 0x72, 0xc9, 0xd7, 0x03, 0x0b, 0x0a, 0x31, 0x17, 0x15, 0x24, 0x2b, 0x09, 0x70, 0xf1, 0x21,
	0x24, 0xe2, 0x8b, 0x52, 0x0b, 0x95, 0x04, 0xb9, 0xf8, 0x51, 0x44, 0x8a, 0x0b, 0x9c, 0x24, 0xa3,
	0xc4, 0x4b, 0x52, 0x13, 0x93, 0x33, 0x52, 0x8b, 0x20, 0xe6, 0x25, 0xe7, 0xe7, 0xe8, 0x17, 0x17,
	0x83, 0x4c, 0x05, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x17, 0xd2, 0x0a, 0x64, 0x00, 0x00, 0x00,
}
