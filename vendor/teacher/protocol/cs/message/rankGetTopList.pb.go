// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/rankGetTopList.proto

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

type RankGetTopListToS struct {
	Version              *uint32  `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	RankID               *uint32  `protobuf:"varint,2,req,name=rankID" json:"rankID,omitempty"`
	GetLast              *bool    `protobuf:"varint,3,opt,name=getLast" json:"getLast,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankGetTopListToS) Reset()         { *m = RankGetTopListToS{} }
func (m *RankGetTopListToS) String() string { return proto.CompactTextString(m) }
func (*RankGetTopListToS) ProtoMessage()    {}
func (*RankGetTopListToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0109f0479f32891, []int{0}
}

func (m *RankGetTopListToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankGetTopListToS.Unmarshal(m, b)
}
func (m *RankGetTopListToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankGetTopListToS.Marshal(b, m, deterministic)
}
func (m *RankGetTopListToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankGetTopListToS.Merge(m, src)
}
func (m *RankGetTopListToS) XXX_Size() int {
	return xxx_messageInfo_RankGetTopListToS.Size(m)
}
func (m *RankGetTopListToS) XXX_DiscardUnknown() {
	xxx_messageInfo_RankGetTopListToS.DiscardUnknown(m)
}

var xxx_messageInfo_RankGetTopListToS proto.InternalMessageInfo

func (m *RankGetTopListToS) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *RankGetTopListToS) GetRankID() uint32 {
	if m != nil && m.RankID != nil {
		return *m.RankID
	}
	return 0
}

func (m *RankGetTopListToS) GetGetLast() bool {
	if m != nil && m.GetLast != nil {
		return *m.GetLast
	}
	return false
}

type RankGetTopListToC struct {
	Version              *uint32         `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	RankInfo             *RankInfo       `protobuf:"bytes,2,opt,name=rankInfo" json:"rankInfo,omitempty"`
	Roles                []*RankRoleInfo `protobuf:"bytes,3,rep,name=roles" json:"roles,omitempty"`
	RankIdx              *int32          `protobuf:"varint,4,opt,name=rankIdx" json:"rankIdx,omitempty"`
	RankPercent          *int32          `protobuf:"varint,5,opt,name=rankPercent" json:"rankPercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RankGetTopListToC) Reset()         { *m = RankGetTopListToC{} }
func (m *RankGetTopListToC) String() string { return proto.CompactTextString(m) }
func (*RankGetTopListToC) ProtoMessage()    {}
func (*RankGetTopListToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0109f0479f32891, []int{1}
}

func (m *RankGetTopListToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankGetTopListToC.Unmarshal(m, b)
}
func (m *RankGetTopListToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankGetTopListToC.Marshal(b, m, deterministic)
}
func (m *RankGetTopListToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankGetTopListToC.Merge(m, src)
}
func (m *RankGetTopListToC) XXX_Size() int {
	return xxx_messageInfo_RankGetTopListToC.Size(m)
}
func (m *RankGetTopListToC) XXX_DiscardUnknown() {
	xxx_messageInfo_RankGetTopListToC.DiscardUnknown(m)
}

var xxx_messageInfo_RankGetTopListToC proto.InternalMessageInfo

func (m *RankGetTopListToC) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *RankGetTopListToC) GetRankInfo() *RankInfo {
	if m != nil {
		return m.RankInfo
	}
	return nil
}

func (m *RankGetTopListToC) GetRoles() []*RankRoleInfo {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *RankGetTopListToC) GetRankIdx() int32 {
	if m != nil && m.RankIdx != nil {
		return *m.RankIdx
	}
	return 0
}

func (m *RankGetTopListToC) GetRankPercent() int32 {
	if m != nil && m.RankPercent != nil {
		return *m.RankPercent
	}
	return 0
}

type RankInfo struct {
	RankID               *uint32  `protobuf:"varint,1,opt,name=RankID" json:"RankID,omitempty"`
	BeginTime            *int64   `protobuf:"varint,2,opt,name=BeginTime" json:"BeginTime,omitempty"`
	EndTime              *int64   `protobuf:"varint,3,opt,name=EndTime" json:"EndTime,omitempty"`
	ExpireTime           *int64   `protobuf:"varint,4,opt,name=ExpireTime" json:"ExpireTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankInfo) Reset()         { *m = RankInfo{} }
func (m *RankInfo) String() string { return proto.CompactTextString(m) }
func (*RankInfo) ProtoMessage()    {}
func (*RankInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0109f0479f32891, []int{2}
}

func (m *RankInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankInfo.Unmarshal(m, b)
}
func (m *RankInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankInfo.Marshal(b, m, deterministic)
}
func (m *RankInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankInfo.Merge(m, src)
}
func (m *RankInfo) XXX_Size() int {
	return xxx_messageInfo_RankInfo.Size(m)
}
func (m *RankInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RankInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RankInfo proto.InternalMessageInfo

func (m *RankInfo) GetRankID() uint32 {
	if m != nil && m.RankID != nil {
		return *m.RankID
	}
	return 0
}

func (m *RankInfo) GetBeginTime() int64 {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return 0
}

func (m *RankInfo) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *RankInfo) GetExpireTime() int64 {
	if m != nil && m.ExpireTime != nil {
		return *m.ExpireTime
	}
	return 0
}

type RankRoleInfo struct {
	ID                   *uint64  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name                 *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Level                *int32   `protobuf:"varint,3,opt,name=level" json:"level,omitempty"`
	Score                *int32   `protobuf:"varint,4,opt,name=score" json:"score,omitempty"`
	CharacterList        []int32  `protobuf:"varint,5,rep,name=characterList" json:"characterList,omitempty"`
	Avatar               *int32   `protobuf:"varint,6,opt,name=avatar" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankRoleInfo) Reset()         { *m = RankRoleInfo{} }
func (m *RankRoleInfo) String() string { return proto.CompactTextString(m) }
func (*RankRoleInfo) ProtoMessage()    {}
func (*RankRoleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0109f0479f32891, []int{3}
}

func (m *RankRoleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankRoleInfo.Unmarshal(m, b)
}
func (m *RankRoleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankRoleInfo.Marshal(b, m, deterministic)
}
func (m *RankRoleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankRoleInfo.Merge(m, src)
}
func (m *RankRoleInfo) XXX_Size() int {
	return xxx_messageInfo_RankRoleInfo.Size(m)
}
func (m *RankRoleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RankRoleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RankRoleInfo proto.InternalMessageInfo

func (m *RankRoleInfo) GetID() uint64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *RankRoleInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RankRoleInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *RankRoleInfo) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *RankRoleInfo) GetCharacterList() []int32 {
	if m != nil {
		return m.CharacterList
	}
	return nil
}

func (m *RankRoleInfo) GetAvatar() int32 {
	if m != nil && m.Avatar != nil {
		return *m.Avatar
	}
	return 0
}

func init() {
	proto.RegisterType((*RankGetTopListToS)(nil), "message.rankGetTopList_toS")
	proto.RegisterType((*RankGetTopListToC)(nil), "message.rankGetTopList_toC")
	proto.RegisterType((*RankInfo)(nil), "message.RankInfo")
	proto.RegisterType((*RankRoleInfo)(nil), "message.RankRoleInfo")
}

func init() {
	proto.RegisterFile("cs/proto/message/rankGetTopList.proto", fileDescriptor_d0109f0479f32891)
}

var fileDescriptor_d0109f0479f32891 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdb, 0xaa, 0x13, 0x31,
	0x18, 0x85, 0xc9, 0x1c, 0x7a, 0xf8, 0x6b, 0x05, 0x83, 0x4a, 0xc0, 0x03, 0x61, 0x50, 0x18, 0x10,
	0x3b, 0xd0, 0x47, 0xa8, 0x2d, 0x52, 0xe8, 0x85, 0xc4, 0x5e, 0x79, 0xa3, 0x21, 0xfd, 0x6d, 0x07,
	0xa7, 0x93, 0x92, 0x84, 0x52, 0x7c, 0x16, 0x9f, 0xc7, 0xe7, 0xda, 0x24, 0x93, 0xd9, 0xbb, 0x85,
	0xcd, 0xbe, 0xcb, 0x5a, 0xdf, 0xca, 0xac, 0xe4, 0x9f, 0xc0, 0x47, 0x65, 0xab, 0x93, 0xd1, 0x4e,
	0x57, 0x47, 0xb4, 0x56, 0xee, 0xb1, 0x32, 0xb2, 0xfd, 0xf3, 0x15, 0xdd, 0x56, 0x9f, 0x36, 0xb5,
	0x75, 0xb3, 0x00, 0xe9, 0x30, 0xd2, 0xe2, 0x17, 0xd0, 0xdb, 0xc0, 0x4f, 0xa7, 0xbf, 0x53, 0x06,
	0xc3, 0x33, 0x1a, 0x5b, 0xeb, 0x96, 0x11, 0x9e, 0x94, 0x53, 0xd1, 0x4b, 0xfa, 0x1a, 0x06, 0x3e,
	0xbf, 0x5e, 0xb2, 0x24, 0x80, 0xa8, 0xfc, 0x8e, 0x3d, 0xba, 0x8d, 0xb4, 0x8e, 0xa5, 0x9c, 0x94,
	0x23, 0xd1, 0xcb, 0xe2, 0x3f, 0x79, 0xa4, 0xe2, 0xcb, 0x13, 0x15, 0x9f, 0x61, 0x14, 0x3e, 0xda,
	0xfe, 0xd6, 0x2c, 0xe1, 0xa4, 0x9c, 0xcc, 0x5f, 0xcc, 0xe2, 0x71, 0x67, 0x22, 0x02, 0x71, 0x1f,
	0xa1, 0x9f, 0x20, 0x37, 0xba, 0x41, 0xcb, 0x52, 0x9e, 0x96, 0x93, 0xf9, 0xab, 0x9b, 0xac, 0xd0,
	0x0d, 0x86, 0x7c, 0x97, 0xf1, 0xad, 0x61, 0xe3, 0xee, 0xc2, 0x32, 0x4e, 0xca, 0x5c, 0xf4, 0x92,
	0x72, 0x98, 0xf8, 0xe5, 0x37, 0x34, 0x0a, 0x5b, 0xc7, 0xf2, 0x40, 0xaf, 0xad, 0xe2, 0x2f, 0x8c,
	0xfa, 0x7a, 0x3f, 0x06, 0xd1, 0x8d, 0x81, 0x70, 0xe2, 0xc7, 0xd0, 0x29, 0xfa, 0x16, 0xc6, 0x0b,
	0xdc, 0xd7, 0xed, 0xb6, 0x3e, 0x62, 0x38, 0x7c, 0x2a, 0x1e, 0x0c, 0xdf, 0xbe, 0x6a, 0x77, 0x81,
	0xa5, 0x81, 0xf5, 0x92, 0xbe, 0x07, 0x58, 0x5d, 0x4e, 0xb5, 0xc1, 0x00, 0xb3, 0x00, 0xaf, 0x9c,
	0xe2, 0x1f, 0x81, 0x67, 0xd7, 0xf7, 0xa1, 0xcf, 0x21, 0x89, 0xe5, 0x99, 0x48, 0xd6, 0x4b, 0x4a,
	0x21, 0x6b, 0x65, 0xec, 0x1c, 0x8b, 0xb0, 0xa6, 0x2f, 0x21, 0x6f, 0xf0, 0x8c, 0x4d, 0x28, 0xcb,
	0x45, 0x27, 0xbc, 0x6b, 0x95, 0x36, 0x18, 0x07, 0xd0, 0x09, 0xfa, 0x01, 0xa6, 0xea, 0x20, 0x8d,
	0x54, 0x0e, 0x8d, 0xff, 0x47, 0x2c, 0xe7, 0x69, 0x99, 0x8b, 0x5b, 0xd3, 0x5f, 0x5b, 0x9e, 0xa5,
	0x93, 0x86, 0x0d, 0xc2, 0xe6, 0xa8, 0x16, 0xef, 0x7e, 0xbc, 0x71, 0x28, 0xd5, 0x01, 0x4d, 0xf7,
	0xf8, 0x94, 0x6e, 0x2a, 0x65, 0xfb, 0x27, 0x78, 0x17, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xf7, 0x43,
	0xad, 0x95, 0x02, 0x00, 0x00,
}
