// Code generated by protoc-gen-go. DO NOT EDIT.
// source: harbor.proto

package cluster

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

type NotifyForginServicesH2SReq struct {
	Nodes                []uint32 `protobuf:"varint,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyForginServicesH2SReq) Reset()         { *m = NotifyForginServicesH2SReq{} }
func (m *NotifyForginServicesH2SReq) String() string { return proto.CompactTextString(m) }
func (*NotifyForginServicesH2SReq) ProtoMessage()    {}
func (*NotifyForginServicesH2SReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{0}
}

func (m *NotifyForginServicesH2SReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyForginServicesH2SReq.Unmarshal(m, b)
}
func (m *NotifyForginServicesH2SReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyForginServicesH2SReq.Marshal(b, m, deterministic)
}
func (m *NotifyForginServicesH2SReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyForginServicesH2SReq.Merge(m, src)
}
func (m *NotifyForginServicesH2SReq) XXX_Size() int {
	return xxx_messageInfo_NotifyForginServicesH2SReq.Size(m)
}
func (m *NotifyForginServicesH2SReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyForginServicesH2SReq.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyForginServicesH2SReq proto.InternalMessageInfo

func (m *NotifyForginServicesH2SReq) GetNodes() []uint32 {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type NotifyForginServicesH2SResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyForginServicesH2SResp) Reset()         { *m = NotifyForginServicesH2SResp{} }
func (m *NotifyForginServicesH2SResp) String() string { return proto.CompactTextString(m) }
func (*NotifyForginServicesH2SResp) ProtoMessage()    {}
func (*NotifyForginServicesH2SResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{1}
}

func (m *NotifyForginServicesH2SResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyForginServicesH2SResp.Unmarshal(m, b)
}
func (m *NotifyForginServicesH2SResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyForginServicesH2SResp.Marshal(b, m, deterministic)
}
func (m *NotifyForginServicesH2SResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyForginServicesH2SResp.Merge(m, src)
}
func (m *NotifyForginServicesH2SResp) XXX_Size() int {
	return xxx_messageInfo_NotifyForginServicesH2SResp.Size(m)
}
func (m *NotifyForginServicesH2SResp) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyForginServicesH2SResp.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyForginServicesH2SResp proto.InternalMessageInfo

type AddForginServicesH2SReq struct {
	Nodes                []uint32 `protobuf:"varint,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddForginServicesH2SReq) Reset()         { *m = AddForginServicesH2SReq{} }
func (m *AddForginServicesH2SReq) String() string { return proto.CompactTextString(m) }
func (*AddForginServicesH2SReq) ProtoMessage()    {}
func (*AddForginServicesH2SReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{2}
}

func (m *AddForginServicesH2SReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddForginServicesH2SReq.Unmarshal(m, b)
}
func (m *AddForginServicesH2SReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddForginServicesH2SReq.Marshal(b, m, deterministic)
}
func (m *AddForginServicesH2SReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddForginServicesH2SReq.Merge(m, src)
}
func (m *AddForginServicesH2SReq) XXX_Size() int {
	return xxx_messageInfo_AddForginServicesH2SReq.Size(m)
}
func (m *AddForginServicesH2SReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddForginServicesH2SReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddForginServicesH2SReq proto.InternalMessageInfo

func (m *AddForginServicesH2SReq) GetNodes() []uint32 {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type AddForginServicesH2SResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddForginServicesH2SResp) Reset()         { *m = AddForginServicesH2SResp{} }
func (m *AddForginServicesH2SResp) String() string { return proto.CompactTextString(m) }
func (*AddForginServicesH2SResp) ProtoMessage()    {}
func (*AddForginServicesH2SResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{3}
}

func (m *AddForginServicesH2SResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddForginServicesH2SResp.Unmarshal(m, b)
}
func (m *AddForginServicesH2SResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddForginServicesH2SResp.Marshal(b, m, deterministic)
}
func (m *AddForginServicesH2SResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddForginServicesH2SResp.Merge(m, src)
}
func (m *AddForginServicesH2SResp) XXX_Size() int {
	return xxx_messageInfo_AddForginServicesH2SResp.Size(m)
}
func (m *AddForginServicesH2SResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddForginServicesH2SResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddForginServicesH2SResp proto.InternalMessageInfo

type RemForginServicesH2SReq struct {
	Nodes                []uint32 `protobuf:"varint,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemForginServicesH2SReq) Reset()         { *m = RemForginServicesH2SReq{} }
func (m *RemForginServicesH2SReq) String() string { return proto.CompactTextString(m) }
func (*RemForginServicesH2SReq) ProtoMessage()    {}
func (*RemForginServicesH2SReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{4}
}

func (m *RemForginServicesH2SReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemForginServicesH2SReq.Unmarshal(m, b)
}
func (m *RemForginServicesH2SReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemForginServicesH2SReq.Marshal(b, m, deterministic)
}
func (m *RemForginServicesH2SReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemForginServicesH2SReq.Merge(m, src)
}
func (m *RemForginServicesH2SReq) XXX_Size() int {
	return xxx_messageInfo_RemForginServicesH2SReq.Size(m)
}
func (m *RemForginServicesH2SReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemForginServicesH2SReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemForginServicesH2SReq proto.InternalMessageInfo

func (m *RemForginServicesH2SReq) GetNodes() []uint32 {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type RemForginServicesH2SResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemForginServicesH2SResp) Reset()         { *m = RemForginServicesH2SResp{} }
func (m *RemForginServicesH2SResp) String() string { return proto.CompactTextString(m) }
func (*RemForginServicesH2SResp) ProtoMessage()    {}
func (*RemForginServicesH2SResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{5}
}

func (m *RemForginServicesH2SResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemForginServicesH2SResp.Unmarshal(m, b)
}
func (m *RemForginServicesH2SResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemForginServicesH2SResp.Marshal(b, m, deterministic)
}
func (m *RemForginServicesH2SResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemForginServicesH2SResp.Merge(m, src)
}
func (m *RemForginServicesH2SResp) XXX_Size() int {
	return xxx_messageInfo_RemForginServicesH2SResp.Size(m)
}
func (m *RemForginServicesH2SResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemForginServicesH2SResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemForginServicesH2SResp proto.InternalMessageInfo

//harbor之间发送
type NotifyForginServicesH2HReq struct {
	Nodes                []uint32 `protobuf:"varint,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyForginServicesH2HReq) Reset()         { *m = NotifyForginServicesH2HReq{} }
func (m *NotifyForginServicesH2HReq) String() string { return proto.CompactTextString(m) }
func (*NotifyForginServicesH2HReq) ProtoMessage()    {}
func (*NotifyForginServicesH2HReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{6}
}

func (m *NotifyForginServicesH2HReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyForginServicesH2HReq.Unmarshal(m, b)
}
func (m *NotifyForginServicesH2HReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyForginServicesH2HReq.Marshal(b, m, deterministic)
}
func (m *NotifyForginServicesH2HReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyForginServicesH2HReq.Merge(m, src)
}
func (m *NotifyForginServicesH2HReq) XXX_Size() int {
	return xxx_messageInfo_NotifyForginServicesH2HReq.Size(m)
}
func (m *NotifyForginServicesH2HReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyForginServicesH2HReq.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyForginServicesH2HReq proto.InternalMessageInfo

func (m *NotifyForginServicesH2HReq) GetNodes() []uint32 {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type NotifyForginServicesH2HResp struct {
	Nodes                []uint32 `protobuf:"varint,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyForginServicesH2HResp) Reset()         { *m = NotifyForginServicesH2HResp{} }
func (m *NotifyForginServicesH2HResp) String() string { return proto.CompactTextString(m) }
func (*NotifyForginServicesH2HResp) ProtoMessage()    {}
func (*NotifyForginServicesH2HResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{7}
}

func (m *NotifyForginServicesH2HResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyForginServicesH2HResp.Unmarshal(m, b)
}
func (m *NotifyForginServicesH2HResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyForginServicesH2HResp.Marshal(b, m, deterministic)
}
func (m *NotifyForginServicesH2HResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyForginServicesH2HResp.Merge(m, src)
}
func (m *NotifyForginServicesH2HResp) XXX_Size() int {
	return xxx_messageInfo_NotifyForginServicesH2HResp.Size(m)
}
func (m *NotifyForginServicesH2HResp) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyForginServicesH2HResp.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyForginServicesH2HResp proto.InternalMessageInfo

func (m *NotifyForginServicesH2HResp) GetNodes() []uint32 {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type AddForginServicesH2HReq struct {
	Nodes                []uint32 `protobuf:"varint,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddForginServicesH2HReq) Reset()         { *m = AddForginServicesH2HReq{} }
func (m *AddForginServicesH2HReq) String() string { return proto.CompactTextString(m) }
func (*AddForginServicesH2HReq) ProtoMessage()    {}
func (*AddForginServicesH2HReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{8}
}

func (m *AddForginServicesH2HReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddForginServicesH2HReq.Unmarshal(m, b)
}
func (m *AddForginServicesH2HReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddForginServicesH2HReq.Marshal(b, m, deterministic)
}
func (m *AddForginServicesH2HReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddForginServicesH2HReq.Merge(m, src)
}
func (m *AddForginServicesH2HReq) XXX_Size() int {
	return xxx_messageInfo_AddForginServicesH2HReq.Size(m)
}
func (m *AddForginServicesH2HReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddForginServicesH2HReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddForginServicesH2HReq proto.InternalMessageInfo

func (m *AddForginServicesH2HReq) GetNodes() []uint32 {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type AddForginServicesH2HResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddForginServicesH2HResp) Reset()         { *m = AddForginServicesH2HResp{} }
func (m *AddForginServicesH2HResp) String() string { return proto.CompactTextString(m) }
func (*AddForginServicesH2HResp) ProtoMessage()    {}
func (*AddForginServicesH2HResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{9}
}

func (m *AddForginServicesH2HResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddForginServicesH2HResp.Unmarshal(m, b)
}
func (m *AddForginServicesH2HResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddForginServicesH2HResp.Marshal(b, m, deterministic)
}
func (m *AddForginServicesH2HResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddForginServicesH2HResp.Merge(m, src)
}
func (m *AddForginServicesH2HResp) XXX_Size() int {
	return xxx_messageInfo_AddForginServicesH2HResp.Size(m)
}
func (m *AddForginServicesH2HResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddForginServicesH2HResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddForginServicesH2HResp proto.InternalMessageInfo

type RemForginServicesH2HReq struct {
	Nodes                []uint32 `protobuf:"varint,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemForginServicesH2HReq) Reset()         { *m = RemForginServicesH2HReq{} }
func (m *RemForginServicesH2HReq) String() string { return proto.CompactTextString(m) }
func (*RemForginServicesH2HReq) ProtoMessage()    {}
func (*RemForginServicesH2HReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{10}
}

func (m *RemForginServicesH2HReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemForginServicesH2HReq.Unmarshal(m, b)
}
func (m *RemForginServicesH2HReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemForginServicesH2HReq.Marshal(b, m, deterministic)
}
func (m *RemForginServicesH2HReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemForginServicesH2HReq.Merge(m, src)
}
func (m *RemForginServicesH2HReq) XXX_Size() int {
	return xxx_messageInfo_RemForginServicesH2HReq.Size(m)
}
func (m *RemForginServicesH2HReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemForginServicesH2HReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemForginServicesH2HReq proto.InternalMessageInfo

func (m *RemForginServicesH2HReq) GetNodes() []uint32 {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type RemForginServicesH2HResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemForginServicesH2HResp) Reset()         { *m = RemForginServicesH2HResp{} }
func (m *RemForginServicesH2HResp) String() string { return proto.CompactTextString(m) }
func (*RemForginServicesH2HResp) ProtoMessage()    {}
func (*RemForginServicesH2HResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40fcde6cab63c6, []int{11}
}

func (m *RemForginServicesH2HResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemForginServicesH2HResp.Unmarshal(m, b)
}
func (m *RemForginServicesH2HResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemForginServicesH2HResp.Marshal(b, m, deterministic)
}
func (m *RemForginServicesH2HResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemForginServicesH2HResp.Merge(m, src)
}
func (m *RemForginServicesH2HResp) XXX_Size() int {
	return xxx_messageInfo_RemForginServicesH2HResp.Size(m)
}
func (m *RemForginServicesH2HResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemForginServicesH2HResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemForginServicesH2HResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NotifyForginServicesH2SReq)(nil), "cluster.notifyForginServicesH2S_req")
	proto.RegisterType((*NotifyForginServicesH2SResp)(nil), "cluster.notifyForginServicesH2S_resp")
	proto.RegisterType((*AddForginServicesH2SReq)(nil), "cluster.addForginServicesH2S_req")
	proto.RegisterType((*AddForginServicesH2SResp)(nil), "cluster.addForginServicesH2S_resp")
	proto.RegisterType((*RemForginServicesH2SReq)(nil), "cluster.remForginServicesH2S_req")
	proto.RegisterType((*RemForginServicesH2SResp)(nil), "cluster.remForginServicesH2S_resp")
	proto.RegisterType((*NotifyForginServicesH2HReq)(nil), "cluster.notifyForginServicesH2H_req")
	proto.RegisterType((*NotifyForginServicesH2HResp)(nil), "cluster.notifyForginServicesH2H_resp")
	proto.RegisterType((*AddForginServicesH2HReq)(nil), "cluster.addForginServicesH2H_req")
	proto.RegisterType((*AddForginServicesH2HResp)(nil), "cluster.addForginServicesH2H_resp")
	proto.RegisterType((*RemForginServicesH2HReq)(nil), "cluster.remForginServicesH2H_req")
	proto.RegisterType((*RemForginServicesH2HResp)(nil), "cluster.remForginServicesH2H_resp")
}

func init() { proto.RegisterFile("harbor.proto", fileDescriptor_1b40fcde6cab63c6) }

var fileDescriptor_1b40fcde6cab63c6 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x2c, 0x4a,
	0xca, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49,
	0x2d, 0x52, 0x32, 0xe6, 0x92, 0xce, 0xcb, 0x2f, 0xc9, 0x4c, 0xab, 0x74, 0xcb, 0x2f, 0x4a, 0xcf,
	0xcc, 0x0b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xf6, 0x30, 0x0a, 0x8e, 0x2f, 0x4a, 0x2d,
	0x14, 0x12, 0xe1, 0x62, 0xcd, 0xcb, 0x4f, 0x49, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0d,
	0x82, 0x70, 0x94, 0xe4, 0xb8, 0x64, 0x70, 0x6b, 0x2a, 0x2e, 0x50, 0x32, 0xe0, 0x92, 0x48, 0x4c,
	0x49, 0x21, 0xc5, 0x44, 0x69, 0x2e, 0x49, 0x1c, 0x3a, 0x20, 0xc6, 0x15, 0xa5, 0xe6, 0x92, 0x68,
	0x1c, 0x0e, 0x1d, 0xc5, 0x05, 0xb8, 0xbd, 0xec, 0x81, 0xc7, 0x44, 0x13, 0x5c, 0x5e, 0xf6, 0x00,
	0x1b, 0x8a, 0x43, 0x17, 0xf6, 0x80, 0xf0, 0x20, 0x39, 0x20, 0x3c, 0xf0, 0x05, 0x84, 0x07, 0xc9,
	0x01, 0x01, 0x31, 0x0e, 0x10, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x0e, 0xea, 0x65, 0x13, 0x02, 0x00,
	0x00,
}
