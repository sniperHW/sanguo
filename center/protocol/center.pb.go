// Code generated by protoc-gen-go. DO NOT EDIT.
// source: center.proto

package protocol

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

type Login struct {
	LogicAddr            *uint32  `protobuf:"varint,1,req,name=logicAddr" json:"logicAddr,omitempty"`
	NetAddr              *string  `protobuf:"bytes,2,req,name=netAddr" json:"netAddr,omitempty"`
	ExportService        *uint32  `protobuf:"varint,3,opt,name=exportService" json:"exportService,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}
func (*Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{0}
}

func (m *Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login.Unmarshal(m, b)
}
func (m *Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login.Marshal(b, m, deterministic)
}
func (m *Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login.Merge(m, src)
}
func (m *Login) XXX_Size() int {
	return xxx_messageInfo_Login.Size(m)
}
func (m *Login) XXX_DiscardUnknown() {
	xxx_messageInfo_Login.DiscardUnknown(m)
}

var xxx_messageInfo_Login proto.InternalMessageInfo

func (m *Login) GetLogicAddr() uint32 {
	if m != nil && m.LogicAddr != nil {
		return *m.LogicAddr
	}
	return 0
}

func (m *Login) GetNetAddr() string {
	if m != nil && m.NetAddr != nil {
		return *m.NetAddr
	}
	return ""
}

func (m *Login) GetExportService() uint32 {
	if m != nil && m.ExportService != nil {
		return *m.ExportService
	}
	return 0
}

type LoginRet struct {
	ErrCode              *int32   `protobuf:"varint,1,req,name=errCode" json:"errCode,omitempty"`
	Msg                  *string  `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRet) Reset()         { *m = LoginRet{} }
func (m *LoginRet) String() string { return proto.CompactTextString(m) }
func (*LoginRet) ProtoMessage()    {}
func (*LoginRet) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{1}
}

func (m *LoginRet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRet.Unmarshal(m, b)
}
func (m *LoginRet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRet.Marshal(b, m, deterministic)
}
func (m *LoginRet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRet.Merge(m, src)
}
func (m *LoginRet) XXX_Size() int {
	return xxx_messageInfo_LoginRet.Size(m)
}
func (m *LoginRet) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRet.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRet proto.InternalMessageInfo

func (m *LoginRet) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *LoginRet) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type HeartbeatToCenter struct {
	Timestamp            *int64   `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatToCenter) Reset()         { *m = HeartbeatToCenter{} }
func (m *HeartbeatToCenter) String() string { return proto.CompactTextString(m) }
func (*HeartbeatToCenter) ProtoMessage()    {}
func (*HeartbeatToCenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{2}
}

func (m *HeartbeatToCenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatToCenter.Unmarshal(m, b)
}
func (m *HeartbeatToCenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatToCenter.Marshal(b, m, deterministic)
}
func (m *HeartbeatToCenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatToCenter.Merge(m, src)
}
func (m *HeartbeatToCenter) XXX_Size() int {
	return xxx_messageInfo_HeartbeatToCenter.Size(m)
}
func (m *HeartbeatToCenter) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatToCenter.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatToCenter proto.InternalMessageInfo

func (m *HeartbeatToCenter) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type HeartbeatToNode struct {
	Timestamp            *int64   `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"`
	TimestampBack        *int64   `protobuf:"varint,2,req,name=timestamp_back,json=timestampBack" json:"timestamp_back,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatToNode) Reset()         { *m = HeartbeatToNode{} }
func (m *HeartbeatToNode) String() string { return proto.CompactTextString(m) }
func (*HeartbeatToNode) ProtoMessage()    {}
func (*HeartbeatToNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{3}
}

func (m *HeartbeatToNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatToNode.Unmarshal(m, b)
}
func (m *HeartbeatToNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatToNode.Marshal(b, m, deterministic)
}
func (m *HeartbeatToNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatToNode.Merge(m, src)
}
func (m *HeartbeatToNode) XXX_Size() int {
	return xxx_messageInfo_HeartbeatToNode.Size(m)
}
func (m *HeartbeatToNode) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatToNode.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatToNode proto.InternalMessageInfo

func (m *HeartbeatToNode) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *HeartbeatToNode) GetTimestampBack() int64 {
	if m != nil && m.TimestampBack != nil {
		return *m.TimestampBack
	}
	return 0
}

type NodeInfo struct {
	LogicAddr            *uint32  `protobuf:"varint,1,req,name=logicAddr" json:"logicAddr,omitempty"`
	NetAddr              *string  `protobuf:"bytes,2,req,name=netAddr" json:"netAddr,omitempty"`
	ExportService        *uint32  `protobuf:"varint,3,opt,name=exportService" json:"exportService,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{4}
}

func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetLogicAddr() uint32 {
	if m != nil && m.LogicAddr != nil {
		return *m.LogicAddr
	}
	return 0
}

func (m *NodeInfo) GetNetAddr() string {
	if m != nil && m.NetAddr != nil {
		return *m.NetAddr
	}
	return ""
}

func (m *NodeInfo) GetExportService() uint32 {
	if m != nil && m.ExportService != nil {
		return *m.ExportService
	}
	return 0
}

type NodeAdd struct {
	Nodes                []*NodeInfo `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NodeAdd) Reset()         { *m = NodeAdd{} }
func (m *NodeAdd) String() string { return proto.CompactTextString(m) }
func (*NodeAdd) ProtoMessage()    {}
func (*NodeAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{5}
}

func (m *NodeAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAdd.Unmarshal(m, b)
}
func (m *NodeAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAdd.Marshal(b, m, deterministic)
}
func (m *NodeAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAdd.Merge(m, src)
}
func (m *NodeAdd) XXX_Size() int {
	return xxx_messageInfo_NodeAdd.Size(m)
}
func (m *NodeAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAdd.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAdd proto.InternalMessageInfo

func (m *NodeAdd) GetNodes() []*NodeInfo {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type NotifyNodeInfo struct {
	Nodes                []*NodeInfo `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NotifyNodeInfo) Reset()         { *m = NotifyNodeInfo{} }
func (m *NotifyNodeInfo) String() string { return proto.CompactTextString(m) }
func (*NotifyNodeInfo) ProtoMessage()    {}
func (*NotifyNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{6}
}

func (m *NotifyNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyNodeInfo.Unmarshal(m, b)
}
func (m *NotifyNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyNodeInfo.Marshal(b, m, deterministic)
}
func (m *NotifyNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyNodeInfo.Merge(m, src)
}
func (m *NotifyNodeInfo) XXX_Size() int {
	return xxx_messageInfo_NotifyNodeInfo.Size(m)
}
func (m *NotifyNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyNodeInfo proto.InternalMessageInfo

func (m *NotifyNodeInfo) GetNodes() []*NodeInfo {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type NodeLeave struct {
	Nodes                []uint32 `protobuf:"varint,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeLeave) Reset()         { *m = NodeLeave{} }
func (m *NodeLeave) String() string { return proto.CompactTextString(m) }
func (*NodeLeave) ProtoMessage()    {}
func (*NodeLeave) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{7}
}

func (m *NodeLeave) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeLeave.Unmarshal(m, b)
}
func (m *NodeLeave) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeLeave.Marshal(b, m, deterministic)
}
func (m *NodeLeave) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeLeave.Merge(m, src)
}
func (m *NodeLeave) XXX_Size() int {
	return xxx_messageInfo_NodeLeave.Size(m)
}
func (m *NodeLeave) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeLeave.DiscardUnknown(m)
}

var xxx_messageInfo_NodeLeave proto.InternalMessageInfo

func (m *NodeLeave) GetNodes() []uint32 {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*Login)(nil), "protocol.login")
	proto.RegisterType((*LoginRet)(nil), "protocol.loginRet")
	proto.RegisterType((*HeartbeatToCenter)(nil), "protocol.heartbeatToCenter")
	proto.RegisterType((*HeartbeatToNode)(nil), "protocol.heartbeatToNode")
	proto.RegisterType((*NodeInfo)(nil), "protocol.nodeInfo")
	proto.RegisterType((*NodeAdd)(nil), "protocol.nodeAdd")
	proto.RegisterType((*NotifyNodeInfo)(nil), "protocol.notifyNodeInfo")
	proto.RegisterType((*NodeLeave)(nil), "protocol.nodeLeave")
}

func init() { proto.RegisterFile("center.proto", fileDescriptor_1de517c49d537f4b) }

var fileDescriptor_1de517c49d537f4b = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x51, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x55, 0x12, 0x45, 0x4d, 0xee, 0xfb, 0x52, 0xc0, 0x62, 0xc8, 0xc0, 0x10, 0x2c, 0x90, 0x3c,
	0x45, 0x02, 0x24, 0x06, 0xb6, 0xd2, 0x09, 0x09, 0x75, 0x30, 0x88, 0x15, 0xb9, 0xf6, 0xb5, 0x8d,
	0xda, 0xd8, 0x91, 0x63, 0x55, 0xf0, 0xef, 0x91, 0x0d, 0x69, 0xcb, 0x82, 0x58, 0x98, 0xfc, 0xee,
	0xf9, 0xde, 0xbd, 0x77, 0x36, 0xfc, 0x97, 0xa8, 0x1d, 0xda, 0xba, 0xb3, 0xc6, 0x19, 0x92, 0x85,
	0x43, 0x9a, 0x0d, 0x45, 0x48, 0x37, 0x66, 0xd9, 0x68, 0x72, 0x06, 0xb9, 0x07, 0x72, 0xa2, 0x94,
	0x2d, 0xa3, 0x2a, 0x66, 0x05, 0xdf, 0x13, 0xa4, 0x84, 0x91, 0x46, 0x17, 0xee, 0xe2, 0x2a, 0x66,
	0x39, 0x1f, 0x4a, 0x72, 0x01, 0x05, 0xbe, 0x75, 0xc6, 0xba, 0x27, 0xb4, 0xdb, 0x46, 0x62, 0x99,
	0x54, 0x11, 0x2b, 0xf8, 0x77, 0x92, 0xde, 0x42, 0x16, 0x6c, 0x38, 0x3a, 0x3f, 0x0b, 0xad, 0x9d,
	0x1a, 0x85, 0xc1, 0x27, 0xe5, 0x43, 0x49, 0x8e, 0x21, 0x69, 0xfb, 0x65, 0x19, 0x57, 0x11, 0xcb,
	0xb9, 0x87, 0xf4, 0x0a, 0x4e, 0x56, 0x28, 0xac, 0x9b, 0xa3, 0x70, 0xcf, 0x66, 0x1a, 0x76, 0xf0,
	0x51, 0x5d, 0xd3, 0x62, 0xef, 0x44, 0xdb, 0x85, 0x11, 0x09, 0xdf, 0x13, 0xf4, 0x05, 0x8e, 0x0e,
	0x24, 0x33, 0x3f, 0xf7, 0x47, 0x01, 0xb9, 0x84, 0xf1, 0xae, 0x78, 0x9d, 0x0b, 0xb9, 0x0e, 0x2b,
	0x26, 0xbc, 0xd8, 0xb1, 0xf7, 0x42, 0xae, 0xe9, 0x0a, 0x32, 0x6d, 0x14, 0x3e, 0xe8, 0x85, 0xf9,
	0xe3, 0xc7, 0xba, 0x81, 0x91, 0x77, 0x9a, 0x28, 0x45, 0x18, 0xa4, 0x1e, 0xf6, 0x65, 0x54, 0x25,
	0xec, 0xdf, 0x35, 0xa9, 0x87, 0x8f, 0xab, 0x87, 0x2c, 0xfc, 0xb3, 0x81, 0xde, 0xc1, 0x58, 0x1b,
	0xd7, 0x2c, 0xde, 0x67, 0x43, 0xc8, 0xdf, 0x6b, 0xcf, 0x21, 0xf7, 0xe0, 0x11, 0xc5, 0x16, 0xc9,
	0xe9, 0xa1, 0xac, 0xf8, 0x6a, 0xf9, 0x08, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xb0, 0xd5, 0x81, 0x40,
	0x02, 0x00, 0x00,
}
