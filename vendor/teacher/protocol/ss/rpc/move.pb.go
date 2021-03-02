// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss/proto/rpc/move.proto

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

type MoveReq struct {
	MapID                *int32    `protobuf:"varint,1,opt,name=mapID" json:"mapID,omitempty"`
	SceneIdx             *int32    `protobuf:"varint,2,opt,name=sceneIdx" json:"sceneIdx,omitempty"`
	UserID               *string   `protobuf:"bytes,3,opt,name=userID" json:"userID,omitempty"`
	RoleID               *uint64   `protobuf:"varint,4,opt,name=roleID" json:"roleID,omitempty"`
	Pos                  *Position `protobuf:"bytes,5,opt,name=pos" json:"pos,omitempty"`
	Angle                *int32    `protobuf:"varint,6,opt,name=angle" json:"angle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MoveReq) Reset()         { *m = MoveReq{} }
func (m *MoveReq) String() string { return proto.CompactTextString(m) }
func (*MoveReq) ProtoMessage()    {}
func (*MoveReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_60fc40db01a3e07d, []int{0}
}

func (m *MoveReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveReq.Unmarshal(m, b)
}
func (m *MoveReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveReq.Marshal(b, m, deterministic)
}
func (m *MoveReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveReq.Merge(m, src)
}
func (m *MoveReq) XXX_Size() int {
	return xxx_messageInfo_MoveReq.Size(m)
}
func (m *MoveReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveReq.DiscardUnknown(m)
}

var xxx_messageInfo_MoveReq proto.InternalMessageInfo

func (m *MoveReq) GetMapID() int32 {
	if m != nil && m.MapID != nil {
		return *m.MapID
	}
	return 0
}

func (m *MoveReq) GetSceneIdx() int32 {
	if m != nil && m.SceneIdx != nil {
		return *m.SceneIdx
	}
	return 0
}

func (m *MoveReq) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *MoveReq) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

func (m *MoveReq) GetPos() *Position {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *MoveReq) GetAngle() int32 {
	if m != nil && m.Angle != nil {
		return *m.Angle
	}
	return 0
}

type MoveResp struct {
	Ok                   *bool     `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Pos                  *Position `protobuf:"bytes,2,opt,name=pos" json:"pos,omitempty"`
	Angle                *int32    `protobuf:"varint,3,opt,name=angle" json:"angle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MoveResp) Reset()         { *m = MoveResp{} }
func (m *MoveResp) String() string { return proto.CompactTextString(m) }
func (*MoveResp) ProtoMessage()    {}
func (*MoveResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_60fc40db01a3e07d, []int{1}
}

func (m *MoveResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveResp.Unmarshal(m, b)
}
func (m *MoveResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveResp.Marshal(b, m, deterministic)
}
func (m *MoveResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveResp.Merge(m, src)
}
func (m *MoveResp) XXX_Size() int {
	return xxx_messageInfo_MoveResp.Size(m)
}
func (m *MoveResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveResp.DiscardUnknown(m)
}

var xxx_messageInfo_MoveResp proto.InternalMessageInfo

func (m *MoveResp) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *MoveResp) GetPos() *Position {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *MoveResp) GetAngle() int32 {
	if m != nil && m.Angle != nil {
		return *m.Angle
	}
	return 0
}

func init() {
	proto.RegisterType((*MoveReq)(nil), "rpc.move_req")
	proto.RegisterType((*MoveResp)(nil), "rpc.move_resp")
}

func init() { proto.RegisterFile("ss/proto/rpc/move.proto", fileDescriptor_60fc40db01a3e07d) }

var fileDescriptor_60fc40db01a3e07d = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0x87, 0xd9, 0xe4, 0x72, 0xe4, 0x56, 0xb4, 0x58, 0xc4, 0xdb, 0xbb, 0xc6, 0x70, 0x55, 0xaa,
	0x04, 0x7c, 0x04, 0x49, 0x93, 0x76, 0x4b, 0x1b, 0x09, 0xeb, 0xa0, 0x47, 0xfe, 0xcc, 0x38, 0x13,
	0xc5, 0xf7, 0xf1, 0x45, 0x65, 0xb3, 0x51, 0xb4, 0xb2, 0xfc, 0xf6, 0x5b, 0x7e, 0x7c, 0x8c, 0xde,
	0x8b, 0xd4, 0xc4, 0x38, 0x63, 0xcd, 0xe4, 0xeb, 0x11, 0xdf, 0xa1, 0x5a, 0xd0, 0xa4, 0x4c, 0xfe,
	0x78, 0xf8, 0x63, 0x3d, 0x8e, 0x23, 0x4e, 0xd1, 0x9f, 0x3e, 0x95, 0xce, 0xc3, 0xf7, 0x47, 0x86,
	0x57, 0x73, 0xad, 0xb3, 0xb1, 0xa3, 0xb6, 0xb1, 0xaa, 0x50, 0x65, 0xe6, 0x22, 0x98, 0xa3, 0xce,
	0xc5, 0xc3, 0x04, 0xed, 0xd3, 0x87, 0x4d, 0x16, 0xf1, 0xc3, 0xe6, 0x46, 0x6f, 0xdf, 0x04, 0xb8,
	0x6d, 0x6c, 0x5a, 0xa8, 0x72, 0xe7, 0x56, 0x0a, 0xef, 0x8c, 0x03, 0xb4, 0x8d, 0xdd, 0x14, 0xaa,
	0xdc, 0xb8, 0x95, 0xcc, 0xad, 0x4e, 0x09, 0xc5, 0x66, 0x85, 0x2a, 0x2f, 0xee, 0x2e, 0x2b, 0x26,
	0x5f, 0x11, 0xca, 0x79, 0x3e, 0xe3, 0xe4, 0x82, 0x09, 0x09, 0xdd, 0xf4, 0x3c, 0x80, 0xdd, 0xc6,
	0x84, 0x05, 0x4e, 0x4e, 0xef, 0xd6, 0x48, 0x21, 0x73, 0xa5, 0x13, 0xec, 0x97, 0xc4, 0xdc, 0x25,
	0xd8, 0x7f, 0x6f, 0x26, 0xff, 0x6f, 0xa6, 0xbf, 0x36, 0xef, 0x0f, 0x0f, 0xfb, 0x19, 0x3a, 0xff,
	0x02, 0x1c, 0x6f, 0xe3, 0x71, 0xa8, 0x45, 0xc2, 0x85, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf6,
	0x6a, 0x66, 0x8e, 0x4e, 0x01, 0x00, 0x00,
}