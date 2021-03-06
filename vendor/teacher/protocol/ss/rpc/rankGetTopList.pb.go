// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/rpc/rankGetTopList.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	message "github.com/sniperHW/sanguo/protocol/cs/message"
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

type RankGetTopListReq struct {
	Tos                  *message.RankGetTopListToS `protobuf:"bytes,1,req,name=tos" json:"tos,omitempty"`
	RoleID               *uint64                    `protobuf:"varint,2,req,name=roleID" json:"roleID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RankGetTopListReq) Reset()         { *m = RankGetTopListReq{} }
func (m *RankGetTopListReq) String() string { return proto.CompactTextString(m) }
func (*RankGetTopListReq) ProtoMessage()    {}
func (*RankGetTopListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b72873a7b127985, []int{0}
}

func (m *RankGetTopListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankGetTopListReq.Unmarshal(m, b)
}
func (m *RankGetTopListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankGetTopListReq.Marshal(b, m, deterministic)
}
func (m *RankGetTopListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankGetTopListReq.Merge(m, src)
}
func (m *RankGetTopListReq) XXX_Size() int {
	return xxx_messageInfo_RankGetTopListReq.Size(m)
}
func (m *RankGetTopListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankGetTopListReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankGetTopListReq proto.InternalMessageInfo

func (m *RankGetTopListReq) GetTos() *message.RankGetTopListToS {
	if m != nil {
		return m.Tos
	}
	return nil
}

func (m *RankGetTopListReq) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

type RankGetTopListResp struct {
	RedirectRankAddr     *string                    `protobuf:"bytes,1,opt,name=redirectRankAddr" json:"redirectRankAddr,omitempty"`
	CurRankID            *uint32                    `protobuf:"varint,2,opt,name=curRankID" json:"curRankID,omitempty"`
	Toc                  *message.RankGetTopListToC `protobuf:"bytes,3,opt,name=toc" json:"toc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RankGetTopListResp) Reset()         { *m = RankGetTopListResp{} }
func (m *RankGetTopListResp) String() string { return proto.CompactTextString(m) }
func (*RankGetTopListResp) ProtoMessage()    {}
func (*RankGetTopListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b72873a7b127985, []int{1}
}

func (m *RankGetTopListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankGetTopListResp.Unmarshal(m, b)
}
func (m *RankGetTopListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankGetTopListResp.Marshal(b, m, deterministic)
}
func (m *RankGetTopListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankGetTopListResp.Merge(m, src)
}
func (m *RankGetTopListResp) XXX_Size() int {
	return xxx_messageInfo_RankGetTopListResp.Size(m)
}
func (m *RankGetTopListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RankGetTopListResp.DiscardUnknown(m)
}

var xxx_messageInfo_RankGetTopListResp proto.InternalMessageInfo

func (m *RankGetTopListResp) GetRedirectRankAddr() string {
	if m != nil && m.RedirectRankAddr != nil {
		return *m.RedirectRankAddr
	}
	return ""
}

func (m *RankGetTopListResp) GetCurRankID() uint32 {
	if m != nil && m.CurRankID != nil {
		return *m.CurRankID
	}
	return 0
}

func (m *RankGetTopListResp) GetToc() *message.RankGetTopListToC {
	if m != nil {
		return m.Toc
	}
	return nil
}

func init() {
	proto.RegisterType((*RankGetTopListReq)(nil), "rpc.rankGetTopList_req")
	proto.RegisterType((*RankGetTopListResp)(nil), "rpc.rankGetTopList_resp")
}

func init() { proto.RegisterFile("ss/proto/rpc/rankGetTopList.proto", fileDescriptor_4b72873a7b127985) }

var fileDescriptor_4b72873a7b127985 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x2e, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x4a, 0xcc, 0xcb, 0x76, 0x4f, 0x2d,
	0x09, 0xc9, 0x2f, 0xf0, 0xc9, 0x2c, 0x2e, 0xd1, 0x03, 0x4b, 0x08, 0x31, 0x17, 0x15, 0x24, 0x4b,
	0xa9, 0x26, 0xc3, 0xd4, 0xe5, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x62, 0x55, 0xab, 0x14, 0xcd,
	0x25, 0x84, 0x2a, 0x1e, 0x5f, 0x94, 0x5a, 0x28, 0xa4, 0xcb, 0xc5, 0x5c, 0x92, 0x5f, 0x2c, 0xc1,
	0xa8, 0xc0, 0xa4, 0xc1, 0x6d, 0x24, 0xad, 0x07, 0x35, 0x41, 0x0f, 0x4d, 0x65, 0x49, 0x7e, 0x70,
	0x10, 0x48, 0x9d, 0x90, 0x18, 0x17, 0x5b, 0x51, 0x7e, 0x4e, 0xaa, 0xa7, 0x8b, 0x04, 0x93, 0x02,
	0x93, 0x06, 0x4b, 0x10, 0x94, 0xa7, 0xd4, 0xc7, 0xc8, 0x25, 0x8c, 0x61, 0x7a, 0x71, 0x81, 0x90,
	0x16, 0x97, 0x40, 0x51, 0x6a, 0x4a, 0x66, 0x51, 0x6a, 0x72, 0x49, 0x50, 0x62, 0x5e, 0xb6, 0x63,
	0x4a, 0x4a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x86, 0xb8, 0x90, 0x0c, 0x17, 0x67,
	0x72, 0x69, 0x11, 0x88, 0x0b, 0x36, 0x9e, 0x51, 0x83, 0x37, 0x08, 0x21, 0x00, 0x71, 0x68, 0xb2,
	0x04, 0xb3, 0x02, 0x23, 0x7e, 0x87, 0x3a, 0x83, 0x1c, 0x9a, 0xec, 0x24, 0x19, 0x25, 0x5e, 0x92,
	0x9a, 0x98, 0x9c, 0x91, 0x5a, 0x04, 0x09, 0x9b, 0xe4, 0xfc, 0x1c, 0xfd, 0xe2, 0x62, 0x50, 0x48,
	0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x9c, 0xa3, 0x37, 0x58, 0x01, 0x00, 0x00,
}
