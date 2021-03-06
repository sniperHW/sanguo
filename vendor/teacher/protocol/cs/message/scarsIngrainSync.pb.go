// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs/proto/message/scarsIngrainSync.proto

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

type ScarsIngrainSyncToS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScarsIngrainSyncToS) Reset()         { *m = ScarsIngrainSyncToS{} }
func (m *ScarsIngrainSyncToS) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainSyncToS) ProtoMessage()    {}
func (*ScarsIngrainSyncToS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4d0a64d81223676, []int{0}
}

func (m *ScarsIngrainSyncToS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainSyncToS.Unmarshal(m, b)
}
func (m *ScarsIngrainSyncToS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainSyncToS.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainSyncToS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainSyncToS.Merge(m, src)
}
func (m *ScarsIngrainSyncToS) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainSyncToS.Size(m)
}
func (m *ScarsIngrainSyncToS) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainSyncToS.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainSyncToS proto.InternalMessageInfo

type ScarsIngrainSyncToC struct {
	ScarsIngrainGroup    []*ScarsIngrain          `protobuf:"bytes,1,rep,name=scarsIngrainGroup" json:"scarsIngrainGroup,omitempty"`
	CurrentID            *int32                   `protobuf:"varint,2,req,name=currentID" json:"currentID,omitempty"`
	TotalScore           *int32                   `protobuf:"varint,3,opt,name=totalScore" json:"totalScore,omitempty"`
	BossGroup            []*ScarsIngrainBossGroup `protobuf:"bytes,4,rep,name=bossGroup" json:"bossGroup,omitempty"`
	CurrentRankId        *uint32                  `protobuf:"varint,5,opt,name=currentRankId" json:"currentRankId,omitempty"`
	DailyTimesUsed       *int32                   `protobuf:"varint,6,opt,name=dailyTimesUsed" json:"dailyTimesUsed,omitempty"`
	RoleUsed             []*RoleChallengeTimes    `protobuf:"bytes,7,rep,name=roleUsed" json:"roleUsed,omitempty"`
	ScoreRewardGet       []int32                  `protobuf:"varint,8,rep,name=scoreRewardGet" json:"scoreRewardGet,omitempty"`
	AutoFightTimes       *int32                   `protobuf:"varint,9,opt,name=autoFightTimes" json:"autoFightTimes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ScarsIngrainSyncToC) Reset()         { *m = ScarsIngrainSyncToC{} }
func (m *ScarsIngrainSyncToC) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainSyncToC) ProtoMessage()    {}
func (*ScarsIngrainSyncToC) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4d0a64d81223676, []int{1}
}

func (m *ScarsIngrainSyncToC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainSyncToC.Unmarshal(m, b)
}
func (m *ScarsIngrainSyncToC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainSyncToC.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainSyncToC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainSyncToC.Merge(m, src)
}
func (m *ScarsIngrainSyncToC) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainSyncToC.Size(m)
}
func (m *ScarsIngrainSyncToC) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainSyncToC.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainSyncToC proto.InternalMessageInfo

func (m *ScarsIngrainSyncToC) GetScarsIngrainGroup() []*ScarsIngrain {
	if m != nil {
		return m.ScarsIngrainGroup
	}
	return nil
}

func (m *ScarsIngrainSyncToC) GetCurrentID() int32 {
	if m != nil && m.CurrentID != nil {
		return *m.CurrentID
	}
	return 0
}

func (m *ScarsIngrainSyncToC) GetTotalScore() int32 {
	if m != nil && m.TotalScore != nil {
		return *m.TotalScore
	}
	return 0
}

func (m *ScarsIngrainSyncToC) GetBossGroup() []*ScarsIngrainBossGroup {
	if m != nil {
		return m.BossGroup
	}
	return nil
}

func (m *ScarsIngrainSyncToC) GetCurrentRankId() uint32 {
	if m != nil && m.CurrentRankId != nil {
		return *m.CurrentRankId
	}
	return 0
}

func (m *ScarsIngrainSyncToC) GetDailyTimesUsed() int32 {
	if m != nil && m.DailyTimesUsed != nil {
		return *m.DailyTimesUsed
	}
	return 0
}

func (m *ScarsIngrainSyncToC) GetRoleUsed() []*RoleChallengeTimes {
	if m != nil {
		return m.RoleUsed
	}
	return nil
}

func (m *ScarsIngrainSyncToC) GetScoreRewardGet() []int32 {
	if m != nil {
		return m.ScoreRewardGet
	}
	return nil
}

func (m *ScarsIngrainSyncToC) GetAutoFightTimes() int32 {
	if m != nil && m.AutoFightTimes != nil {
		return *m.AutoFightTimes
	}
	return 0
}

type ScarsIngrain struct {
	ScarsIngrainID       *int32                   `protobuf:"varint,1,req,name=scarsIngrainID" json:"scarsIngrainID,omitempty"`
	RankID               *uint32                  `protobuf:"varint,2,req,name=rankID" json:"rankID,omitempty"`
	BossGroup            []*ScarsIngrainBossGroup `protobuf:"bytes,3,rep,name=bossGroup" json:"bossGroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ScarsIngrain) Reset()         { *m = ScarsIngrain{} }
func (m *ScarsIngrain) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrain) ProtoMessage()    {}
func (*ScarsIngrain) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4d0a64d81223676, []int{2}
}

func (m *ScarsIngrain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrain.Unmarshal(m, b)
}
func (m *ScarsIngrain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrain.Marshal(b, m, deterministic)
}
func (m *ScarsIngrain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrain.Merge(m, src)
}
func (m *ScarsIngrain) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrain.Size(m)
}
func (m *ScarsIngrain) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrain.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrain proto.InternalMessageInfo

func (m *ScarsIngrain) GetScarsIngrainID() int32 {
	if m != nil && m.ScarsIngrainID != nil {
		return *m.ScarsIngrainID
	}
	return 0
}

func (m *ScarsIngrain) GetRankID() uint32 {
	if m != nil && m.RankID != nil {
		return *m.RankID
	}
	return 0
}

func (m *ScarsIngrain) GetBossGroup() []*ScarsIngrainBossGroup {
	if m != nil {
		return m.BossGroup
	}
	return nil
}

type RoleChallengeTimes struct {
	RoleID               *int32   `protobuf:"varint,1,req,name=roleID" json:"roleID,omitempty"`
	UseTimes             *int32   `protobuf:"varint,2,req,name=useTimes" json:"useTimes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleChallengeTimes) Reset()         { *m = RoleChallengeTimes{} }
func (m *RoleChallengeTimes) String() string { return proto.CompactTextString(m) }
func (*RoleChallengeTimes) ProtoMessage()    {}
func (*RoleChallengeTimes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4d0a64d81223676, []int{3}
}

func (m *RoleChallengeTimes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleChallengeTimes.Unmarshal(m, b)
}
func (m *RoleChallengeTimes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleChallengeTimes.Marshal(b, m, deterministic)
}
func (m *RoleChallengeTimes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleChallengeTimes.Merge(m, src)
}
func (m *RoleChallengeTimes) XXX_Size() int {
	return xxx_messageInfo_RoleChallengeTimes.Size(m)
}
func (m *RoleChallengeTimes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleChallengeTimes.DiscardUnknown(m)
}

var xxx_messageInfo_RoleChallengeTimes proto.InternalMessageInfo

func (m *RoleChallengeTimes) GetRoleID() int32 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

func (m *RoleChallengeTimes) GetUseTimes() int32 {
	if m != nil && m.UseTimes != nil {
		return *m.UseTimes
	}
	return 0
}

type ScarsIngrainBossGroup struct {
	GroupConfigID        *int32                       `protobuf:"varint,1,req,name=groupConfigID" json:"groupConfigID,omitempty"`
	BossID               *int32                       `protobuf:"varint,2,req,name=bossID" json:"bossID,omitempty"`
	BossDifficult        []*ScarsIngrainBossDifficult `protobuf:"bytes,3,rep,name=bossDifficult" json:"bossDifficult,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ScarsIngrainBossGroup) Reset()         { *m = ScarsIngrainBossGroup{} }
func (m *ScarsIngrainBossGroup) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainBossGroup) ProtoMessage()    {}
func (*ScarsIngrainBossGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4d0a64d81223676, []int{4}
}

func (m *ScarsIngrainBossGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainBossGroup.Unmarshal(m, b)
}
func (m *ScarsIngrainBossGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainBossGroup.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainBossGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainBossGroup.Merge(m, src)
}
func (m *ScarsIngrainBossGroup) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainBossGroup.Size(m)
}
func (m *ScarsIngrainBossGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainBossGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainBossGroup proto.InternalMessageInfo

func (m *ScarsIngrainBossGroup) GetGroupConfigID() int32 {
	if m != nil && m.GroupConfigID != nil {
		return *m.GroupConfigID
	}
	return 0
}

func (m *ScarsIngrainBossGroup) GetBossID() int32 {
	if m != nil && m.BossID != nil {
		return *m.BossID
	}
	return 0
}

func (m *ScarsIngrainBossGroup) GetBossDifficult() []*ScarsIngrainBossDifficult {
	if m != nil {
		return m.BossDifficult
	}
	return nil
}

type ScarsIngrainBossDifficult struct {
	Difficult            *int32         `protobuf:"varint,1,req,name=difficult" json:"difficult,omitempty"`
	WeeklyHighScore      *int32         `protobuf:"varint,2,req,name=weeklyHighScore" json:"weeklyHighScore,omitempty"`
	HighScore            *int32         `protobuf:"varint,3,opt,name=highScore" json:"highScore,omitempty"`
	HighTeam             *CharacterTeam `protobuf:"bytes,4,opt,name=highTeam" json:"highTeam,omitempty"`
	CanAuto              *bool          `protobuf:"varint,5,opt,name=canAuto" json:"canAuto,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ScarsIngrainBossDifficult) Reset()         { *m = ScarsIngrainBossDifficult{} }
func (m *ScarsIngrainBossDifficult) String() string { return proto.CompactTextString(m) }
func (*ScarsIngrainBossDifficult) ProtoMessage()    {}
func (*ScarsIngrainBossDifficult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4d0a64d81223676, []int{5}
}

func (m *ScarsIngrainBossDifficult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScarsIngrainBossDifficult.Unmarshal(m, b)
}
func (m *ScarsIngrainBossDifficult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScarsIngrainBossDifficult.Marshal(b, m, deterministic)
}
func (m *ScarsIngrainBossDifficult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScarsIngrainBossDifficult.Merge(m, src)
}
func (m *ScarsIngrainBossDifficult) XXX_Size() int {
	return xxx_messageInfo_ScarsIngrainBossDifficult.Size(m)
}
func (m *ScarsIngrainBossDifficult) XXX_DiscardUnknown() {
	xxx_messageInfo_ScarsIngrainBossDifficult.DiscardUnknown(m)
}

var xxx_messageInfo_ScarsIngrainBossDifficult proto.InternalMessageInfo

func (m *ScarsIngrainBossDifficult) GetDifficult() int32 {
	if m != nil && m.Difficult != nil {
		return *m.Difficult
	}
	return 0
}

func (m *ScarsIngrainBossDifficult) GetWeeklyHighScore() int32 {
	if m != nil && m.WeeklyHighScore != nil {
		return *m.WeeklyHighScore
	}
	return 0
}

func (m *ScarsIngrainBossDifficult) GetHighScore() int32 {
	if m != nil && m.HighScore != nil {
		return *m.HighScore
	}
	return 0
}

func (m *ScarsIngrainBossDifficult) GetHighTeam() *CharacterTeam {
	if m != nil {
		return m.HighTeam
	}
	return nil
}

func (m *ScarsIngrainBossDifficult) GetCanAuto() bool {
	if m != nil && m.CanAuto != nil {
		return *m.CanAuto
	}
	return false
}

func init() {
	proto.RegisterType((*ScarsIngrainSyncToS)(nil), "message.scarsIngrainSync_toS")
	proto.RegisterType((*ScarsIngrainSyncToC)(nil), "message.scarsIngrainSync_toC")
	proto.RegisterType((*ScarsIngrain)(nil), "message.scarsIngrain")
	proto.RegisterType((*RoleChallengeTimes)(nil), "message.roleChallengeTimes")
	proto.RegisterType((*ScarsIngrainBossGroup)(nil), "message.scarsIngrainBossGroup")
	proto.RegisterType((*ScarsIngrainBossDifficult)(nil), "message.scarsIngrainBossDifficult")
}

func init() {
	proto.RegisterFile("cs/proto/message/scarsIngrainSync.proto", fileDescriptor_d4d0a64d81223676)
}

var fileDescriptor_d4d0a64d81223676 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0xd5, 0x36, 0x4d, 0x93, 0x4c, 0x09, 0x08, 0x8b, 0x46, 0xa6, 0x85, 0x6a, 0xb5, 0x42, 0xb0,
	0xa7, 0x44, 0xca, 0x85, 0x0b, 0x17, 0x9a, 0x88, 0x26, 0x57, 0xa7, 0x5c, 0xb8, 0x20, 0xe3, 0x38,
	0xbb, 0xab, 0x6e, 0xd6, 0x95, 0xed, 0x55, 0x95, 0x7f, 0xe0, 0x1f, 0xf8, 0x03, 0x7e, 0x85, 0x5f,
	0x42, 0xb6, 0x37, 0xde, 0x6c, 0x92, 0x1e, 0x7a, 0xdb, 0x79, 0x7e, 0x33, 0xef, 0x8d, 0x77, 0xc6,
	0xf0, 0x89, 0xa9, 0xd1, 0x83, 0x14, 0x5a, 0x8c, 0xd6, 0x5c, 0x29, 0x9a, 0xf0, 0x91, 0x62, 0x54,
	0xaa, 0x79, 0x91, 0x48, 0x9a, 0x15, 0x8b, 0x4d, 0xc1, 0x86, 0xf6, 0x18, 0x75, 0xaa, 0xf3, 0xcb,
	0xf8, 0x20, 0x83, 0xa5, 0x54, 0x52, 0xa6, 0xb9, 0xbc, 0xe3, 0x74, 0x5d, 0xa7, 0x44, 0x03, 0x78,
	0xb3, 0x5f, 0xec, 0xa7, 0x16, 0x8b, 0xe8, 0x6f, 0xeb, 0xe8, 0xc1, 0x04, 0x4d, 0xe0, 0xf5, 0x2e,
	0x7e, 0x2b, 0x45, 0xf9, 0x80, 0x83, 0xb0, 0x15, 0x9f, 0x8f, 0x2f, 0x86, 0x95, 0xda, 0x70, 0x97,
	0x41, 0x0e, 0xf9, 0xe8, 0x1d, 0xf4, 0x58, 0x29, 0x25, 0x2f, 0xf4, 0x7c, 0x8a, 0x4f, 0xc2, 0x93,
	0xb8, 0x4d, 0x6a, 0x00, 0x5d, 0x03, 0x68, 0xa1, 0x69, 0xbe, 0x60, 0x42, 0x72, 0xdc, 0x0a, 0x83,
	0xb8, 0x4d, 0x76, 0x10, 0xf4, 0x05, 0x7a, 0xbf, 0x84, 0x52, 0x4e, 0xfa, 0xd4, 0x4a, 0x5f, 0x1f,
	0x95, 0xbe, 0xd9, 0xb2, 0x48, 0x9d, 0x80, 0x3e, 0x40, 0xbf, 0x92, 0x22, 0xb4, 0xb8, 0x9f, 0x2f,
	0x71, 0x3b, 0x0c, 0xe2, 0x3e, 0x69, 0x82, 0xe8, 0x23, 0xbc, 0x5c, 0xd2, 0x2c, 0xdf, 0xdc, 0x65,
	0x6b, 0xae, 0xbe, 0x2b, 0xbe, 0xc4, 0x67, 0xd6, 0xc7, 0x1e, 0x8a, 0x3e, 0x43, 0x57, 0x8a, 0x9c,
	0x5b, 0x46, 0xc7, 0x5a, 0xb9, 0xf2, 0x56, 0xcc, 0xc1, 0x24, 0xa5, 0x79, 0xce, 0x8b, 0x84, 0xdb,
	0x14, 0xe2, 0xc9, 0x46, 0x40, 0x99, 0x6e, 0x08, 0x7f, 0xa4, 0x72, 0x79, 0xcb, 0x35, 0xee, 0x86,
	0x2d, 0x23, 0xd0, 0x44, 0x0d, 0x8f, 0x96, 0x5a, 0x7c, 0xcb, 0x92, 0x54, 0xdb, 0x1a, 0xb8, 0xe7,
	0x8c, 0x34, 0xd1, 0xe8, 0x77, 0x00, 0x2f, 0x76, 0x7b, 0x77, 0x02, 0x75, 0x3c, 0x9f, 0xe2, 0xc0,
	0x5e, 0xf4, 0x1e, 0x8a, 0x06, 0x70, 0x26, 0x4d, 0xcf, 0xee, 0x47, 0xf4, 0x49, 0x15, 0x35, 0x6f,
	0xb9, 0xf5, 0xcc, 0x5b, 0x8e, 0x66, 0x80, 0x0e, 0xdb, 0xb7, 0x5a, 0x22, 0xe7, 0xde, 0x4b, 0x15,
	0xa1, 0x4b, 0xe8, 0x96, 0xca, 0x71, 0xaa, 0x71, 0xf0, 0x71, 0xf4, 0x27, 0x80, 0x8b, 0xa3, 0x72,
	0xe6, 0x4f, 0x26, 0xe6, 0x63, 0x22, 0x8a, 0x55, 0x96, 0xf8, 0xa2, 0x4d, 0xd0, 0x68, 0x1a, 0x5b,
	0x7e, 0xd0, 0xaa, 0x08, 0xcd, 0xa0, 0x6f, 0xbe, 0xa6, 0xd9, 0x6a, 0x95, 0xb1, 0x32, 0xd7, 0x55,
	0x8f, 0xd1, 0x93, 0x3d, 0x7a, 0x26, 0x69, 0x26, 0x46, 0xff, 0x02, 0x78, 0xfb, 0x24, 0xd9, 0xcc,
	0xfa, 0xd2, 0x6b, 0x38, 0x87, 0x35, 0x80, 0x62, 0x78, 0xf5, 0xc8, 0xf9, 0x7d, 0xbe, 0x99, 0x65,
	0x49, 0xea, 0x06, 0xde, 0xd9, 0xdc, 0x87, 0x4d, 0x9d, 0xd4, 0x73, 0xdc, 0x52, 0xd4, 0x00, 0x1a,
	0x43, 0xd7, 0x04, 0x66, 0xbb, 0xf1, 0x69, 0x18, 0xc4, 0xe7, 0xe3, 0x81, 0x6f, 0xa4, 0xb1, 0xfb,
	0xc4, 0xf3, 0x10, 0x86, 0x0e, 0xa3, 0xc5, 0xd7, 0x52, 0x0b, 0xbb, 0x03, 0x5d, 0xb2, 0x0d, 0x6f,
	0xde, 0xff, 0xb8, 0xd2, 0x9c, 0xb2, 0x94, 0x4b, 0xf7, 0x8c, 0x30, 0x91, 0x8f, 0x98, 0xda, 0x3e,
	0x26, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xf6, 0xcc, 0xb3, 0x91, 0x04, 0x00, 0x00,
}
