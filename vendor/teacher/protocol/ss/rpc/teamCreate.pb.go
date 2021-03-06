// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/rpc/teamCreate.proto

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

type TeamCreateReq struct {
	Player               *message.TeamPlayer `protobuf:"bytes,1,req,name=player" json:"player,omitempty"`
	Target               *message.TeamTarget `protobuf:"bytes,2,req,name=target" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TeamCreateReq) Reset()         { *m = TeamCreateReq{} }
func (m *TeamCreateReq) String() string { return proto.CompactTextString(m) }
func (*TeamCreateReq) ProtoMessage()    {}
func (*TeamCreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec8f687580c083b4, []int{0}
}

func (m *TeamCreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamCreateReq.Unmarshal(m, b)
}
func (m *TeamCreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamCreateReq.Marshal(b, m, deterministic)
}
func (m *TeamCreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamCreateReq.Merge(m, src)
}
func (m *TeamCreateReq) XXX_Size() int {
	return xxx_messageInfo_TeamCreateReq.Size(m)
}
func (m *TeamCreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamCreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_TeamCreateReq proto.InternalMessageInfo

func (m *TeamCreateReq) GetPlayer() *message.TeamPlayer {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *TeamCreateReq) GetTarget() *message.TeamTarget {
	if m != nil {
		return m.Target
	}
	return nil
}

type TeamCreateResp struct {
	ErrCode              *message.ErrCode `protobuf:"varint,1,req,name=errCode,enum=message.ErrCode" json:"errCode,omitempty"`
	TeamID               *uint32          `protobuf:"varint,2,req,name=teamID" json:"teamID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TeamCreateResp) Reset()         { *m = TeamCreateResp{} }
func (m *TeamCreateResp) String() string { return proto.CompactTextString(m) }
func (*TeamCreateResp) ProtoMessage()    {}
func (*TeamCreateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec8f687580c083b4, []int{1}
}

func (m *TeamCreateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamCreateResp.Unmarshal(m, b)
}
func (m *TeamCreateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamCreateResp.Marshal(b, m, deterministic)
}
func (m *TeamCreateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamCreateResp.Merge(m, src)
}
func (m *TeamCreateResp) XXX_Size() int {
	return xxx_messageInfo_TeamCreateResp.Size(m)
}
func (m *TeamCreateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamCreateResp.DiscardUnknown(m)
}

var xxx_messageInfo_TeamCreateResp proto.InternalMessageInfo

func (m *TeamCreateResp) GetErrCode() message.ErrCode {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return message.ErrCode_OK
}

func (m *TeamCreateResp) GetTeamID() uint32 {
	if m != nil && m.TeamID != nil {
		return *m.TeamID
	}
	return 0
}

func init() {
	proto.RegisterType((*TeamCreateReq)(nil), "rpc.teamCreate_req")
	proto.RegisterType((*TeamCreateResp)(nil), "rpc.teamCreate_resp")
}

func init() { proto.RegisterFile("ss/proto/rpc/teamCreate.proto", fileDescriptor_ec8f687580c083b4) }

var fileDescriptor_ec8f687580c083b4 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0x41, 0x6b, 0xc2, 0x30,
	0x14, 0x07, 0x70, 0xd6, 0x41, 0x07, 0x19, 0xeb, 0x46, 0x06, 0x5b, 0xd7, 0x31, 0x26, 0x3d, 0x89,
	0x42, 0x03, 0xfd, 0x08, 0x56, 0x0f, 0xde, 0x24, 0xd4, 0x8b, 0x17, 0x09, 0xf1, 0x51, 0x91, 0x96,
	0xc4, 0x97, 0x5c, 0xfc, 0xf6, 0xd2, 0x24, 0x55, 0x44, 0x8f, 0x7d, 0xff, 0xdf, 0x7b, 0x7f, 0x1a,
	0xf2, 0x67, 0x0c, 0xd3, 0xa8, 0xac, 0x62, 0xa8, 0x25, 0xb3, 0x20, 0xba, 0x0a, 0x41, 0x58, 0x28,
	0xdc, 0x90, 0x3e, 0xa3, 0x96, 0xd9, 0xaf, 0x1c, 0x4c, 0x07, 0xc6, 0x88, 0x06, 0x9c, 0xf3, 0x22,
	0xfb, 0xbf, 0x0b, 0x01, 0x11, 0x10, 0x15, 0x7a, 0x90, 0x1f, 0x48, 0x72, 0x3d, 0xbb, 0x45, 0x38,
	0xd2, 0x29, 0x89, 0x75, 0x2b, 0x4e, 0x80, 0xe9, 0xd3, 0x28, 0x1a, 0xbf, 0x96, 0x9f, 0x45, 0x58,
	0x2d, 0x6a, 0x10, 0xdd, 0xca, 0x45, 0x3c, 0x90, 0x1e, 0x5b, 0x81, 0x0d, 0xd8, 0x34, 0x7a, 0x80,
	0x6b, 0x17, 0xf1, 0x40, 0xf2, 0x35, 0x79, 0xbf, 0xe9, 0x32, 0x9a, 0x4e, 0xc8, 0x0b, 0x20, 0x56,
	0x6a, 0x07, 0xae, 0x2d, 0x29, 0x3f, 0x2e, 0x07, 0x16, 0x7e, 0xce, 0x07, 0x40, 0xbf, 0x48, 0xdc,
	0xaf, 0x2f, 0xe7, 0xae, 0xeb, 0x8d, 0x87, 0xaf, 0xd9, 0xcf, 0xe6, 0xdb, 0x82, 0x90, 0x7b, 0x40,
	0xff, 0xab, 0x52, 0xb5, 0xcc, 0x98, 0xfe, 0xc5, 0xce, 0x01, 0x00, 0x00, 0xff, 0xff, 0x73, 0x07,
	0x7b, 0xba, 0x40, 0x01, 0x00, 0x00,
}
