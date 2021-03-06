// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/rewardQuestSync.proto

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

type RewardQuestSyncToS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RewardQuestSyncToS) Reset()         { *m = RewardQuestSyncToS{} }
func (m *RewardQuestSyncToS) String() string { return proto.CompactTextString(m) }
func (*RewardQuestSyncToS) ProtoMessage()    {}
func (*RewardQuestSyncToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d2af3604d8696e3, []int{0}
}

func (m *RewardQuestSyncToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewardQuestSyncToS.Unmarshal(m, b)
}
func (m *RewardQuestSyncToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewardQuestSyncToS.Marshal(b, m, deterministic)
}
func (m *RewardQuestSyncToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardQuestSyncToS.Merge(m, src)
}
func (m *RewardQuestSyncToS) XXX_Size() int {
	return xxx_messageInfo_RewardQuestSyncToS.Size(m)
}
func (m *RewardQuestSyncToS) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardQuestSyncToS.DiscardUnknown(m)
}

var xxx_messageInfo_RewardQuestSyncToS proto.InternalMessageInfo

type RewardQuestSyncToC struct {
	IsAll                *bool          `protobuf:"varint,1,req,name=isAll" json:"isAll,omitempty"`
	Quests               []*RewardQuest `protobuf:"bytes,2,rep,name=Quests" json:"Quests,omitempty"`
	RefreshTimes         *int32         `protobuf:"varint,3,req,name=refreshTimes" json:"refreshTimes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RewardQuestSyncToC) Reset()         { *m = RewardQuestSyncToC{} }
func (m *RewardQuestSyncToC) String() string { return proto.CompactTextString(m) }
func (*RewardQuestSyncToC) ProtoMessage()    {}
func (*RewardQuestSyncToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d2af3604d8696e3, []int{1}
}

func (m *RewardQuestSyncToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewardQuestSyncToC.Unmarshal(m, b)
}
func (m *RewardQuestSyncToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewardQuestSyncToC.Marshal(b, m, deterministic)
}
func (m *RewardQuestSyncToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardQuestSyncToC.Merge(m, src)
}
func (m *RewardQuestSyncToC) XXX_Size() int {
	return xxx_messageInfo_RewardQuestSyncToC.Size(m)
}
func (m *RewardQuestSyncToC) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardQuestSyncToC.DiscardUnknown(m)
}

var xxx_messageInfo_RewardQuestSyncToC proto.InternalMessageInfo

func (m *RewardQuestSyncToC) GetIsAll() bool {
	if m != nil && m.IsAll != nil {
		return *m.IsAll
	}
	return false
}

func (m *RewardQuestSyncToC) GetQuests() []*RewardQuest {
	if m != nil {
		return m.Quests
	}
	return nil
}

func (m *RewardQuestSyncToC) GetRefreshTimes() int32 {
	if m != nil && m.RefreshTimes != nil {
		return *m.RefreshTimes
	}
	return 0
}

type RewardQuest struct {
	QuestID              *int32      `protobuf:"varint,1,req,name=questID" json:"questID,omitempty"`
	State                *QuestState `protobuf:"varint,2,opt,name=State,enum=message.QuestState" json:"State,omitempty"`
	Characters           []int32     `protobuf:"varint,3,rep,name=characters" json:"characters,omitempty"`
	AcceptTimestamp      *int64      `protobuf:"varint,4,opt,name=acceptTimestamp" json:"acceptTimestamp,omitempty"`
	IsRemoved            *bool       `protobuf:"varint,5,opt,name=isRemoved" json:"isRemoved,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RewardQuest) Reset()         { *m = RewardQuest{} }
func (m *RewardQuest) String() string { return proto.CompactTextString(m) }
func (*RewardQuest) ProtoMessage()    {}
func (*RewardQuest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d2af3604d8696e3, []int{2}
}

func (m *RewardQuest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewardQuest.Unmarshal(m, b)
}
func (m *RewardQuest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewardQuest.Marshal(b, m, deterministic)
}
func (m *RewardQuest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardQuest.Merge(m, src)
}
func (m *RewardQuest) XXX_Size() int {
	return xxx_messageInfo_RewardQuest.Size(m)
}
func (m *RewardQuest) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardQuest.DiscardUnknown(m)
}

var xxx_messageInfo_RewardQuest proto.InternalMessageInfo

func (m *RewardQuest) GetQuestID() int32 {
	if m != nil && m.QuestID != nil {
		return *m.QuestID
	}
	return 0
}

func (m *RewardQuest) GetState() QuestState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return QuestState_UnAcceptable
}

func (m *RewardQuest) GetCharacters() []int32 {
	if m != nil {
		return m.Characters
	}
	return nil
}

func (m *RewardQuest) GetAcceptTimestamp() int64 {
	if m != nil && m.AcceptTimestamp != nil {
		return *m.AcceptTimestamp
	}
	return 0
}

func (m *RewardQuest) GetIsRemoved() bool {
	if m != nil && m.IsRemoved != nil {
		return *m.IsRemoved
	}
	return false
}

func init() {
	proto.RegisterType((*RewardQuestSyncToS)(nil), "message.rewardQuestSync_toS")
	proto.RegisterType((*RewardQuestSyncToC)(nil), "message.rewardQuestSync_toC")
	proto.RegisterType((*RewardQuest)(nil), "message.rewardQuest")
}

func init() {
	proto.RegisterFile("cs/proto/message/rewardQuestSync.proto", fileDescriptor_3d2af3604d8696e3)
}

var fileDescriptor_3d2af3604d8696e3 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4a, 0xfc, 0x30,
	0x14, 0xc5, 0x69, 0xfb, 0xcf, 0x7f, 0xc6, 0x3b, 0xa2, 0x90, 0x19, 0x21, 0xf8, 0x45, 0xe8, 0x42,
	0x22, 0x48, 0x0b, 0xf3, 0x06, 0x7e, 0x6c, 0x5c, 0x9a, 0xba, 0x72, 0x23, 0x21, 0x73, 0xb5, 0x85,
	0xd6, 0xd6, 0xdc, 0xa8, 0xb8, 0xf0, 0xc5, 0x7c, 0x3a, 0x99, 0xb6, 0xa3, 0x9d, 0xd1, 0xe5, 0xfd,
	0x9d, 0x73, 0xef, 0x39, 0x24, 0x70, 0x62, 0x29, 0x6d, 0x5c, 0xed, 0xeb, 0xb4, 0x42, 0x22, 0xf3,
	0x88, 0xa9, 0xc3, 0x37, 0xe3, 0x16, 0x37, 0x2f, 0x48, 0x3e, 0x7b, 0x7f, 0xb2, 0x49, 0xab, 0xf2,
	0x51, 0x2f, 0xef, 0xcb, 0x5f, 0x0b, 0xcf, 0xeb, 0xd6, 0x78, 0x0f, 0xa6, 0x1b, 0x37, 0xee, 0x7d,
	0x9d, 0xc5, 0x1f, 0x7f, 0xe1, 0x4b, 0x3e, 0x03, 0x56, 0xd0, 0x79, 0x59, 0x8a, 0x40, 0x86, 0x6a,
	0xac, 0xbb, 0x81, 0x9f, 0xc1, 0xff, 0xd6, 0x46, 0x22, 0x94, 0x91, 0x9a, 0xcc, 0x67, 0x49, 0x9f,
	0x96, 0x0c, 0x6e, 0xe8, 0xde, 0xc3, 0x63, 0xd8, 0x76, 0xf8, 0xe0, 0x90, 0xf2, 0xdb, 0xa2, 0x42,
	0x12, 0x91, 0x0c, 0x15, 0xd3, 0x6b, 0x2c, 0xfe, 0x0c, 0x60, 0x32, 0xd8, 0xe5, 0x02, 0x46, 0x6d,
	0xf1, 0xeb, 0xab, 0x36, 0x99, 0xe9, 0xd5, 0xc8, 0x4f, 0x81, 0x65, 0xde, 0x78, 0x14, 0xa1, 0x0c,
	0xd4, 0xce, 0x7c, 0xfa, 0x1d, 0xdd, 0x15, 0x5f, 0x4a, 0xba, 0x73, 0xf0, 0x63, 0x00, 0x9b, 0x1b,
	0x67, 0xac, 0x47, 0xb7, 0x8c, 0x8d, 0x14, 0xd3, 0x03, 0xc2, 0x15, 0xec, 0x1a, 0x6b, 0xb1, 0xf1,
	0x6d, 0x07, 0x6f, 0xaa, 0x46, 0xfc, 0x93, 0x81, 0x8a, 0xf4, 0x26, 0xe6, 0x87, 0xb0, 0x55, 0x90,
	0xc6, 0xaa, 0x7e, 0xc5, 0x85, 0x60, 0x32, 0x50, 0x63, 0xfd, 0x03, 0x2e, 0x8e, 0xee, 0x0e, 0x3c,
	0x1a, 0x9b, 0xa3, 0xeb, 0xde, 0xde, 0xd6, 0x65, 0x6a, 0x69, 0xf5, 0x03, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xd5, 0xa2, 0x58, 0xc7, 0xc5, 0x01, 0x00, 0x00,
}
