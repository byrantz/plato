// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

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

// cd common/idl; protoc -I message  --go_out=plugins=grpc:message  message/message.proto
type CmdType int32

const (
	CmdType_Login     CmdType = 0
	CmdType_Heartbeat CmdType = 1
	CmdType_ReConn    CmdType = 2
	CmdType_ACK       CmdType = 3
	CmdType_UP        CmdType = 4
	CmdType_Push      CmdType = 5
)

var CmdType_name = map[int32]string{
	0: "Login",
	1: "Heartbeat",
	2: "ReConn",
	3: "ACK",
	4: "UP",
	5: "Push",
}

var CmdType_value = map[string]int32{
	"Login":     0,
	"Heartbeat": 1,
	"ReConn":    2,
	"ACK":       3,
	"UP":        4,
	"Push":      5,
}

func (x CmdType) String() string {
	return proto.EnumName(CmdType_name, int32(x))
}

func (CmdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

// 顶层cmd pb结构
type MsgCmd struct {
	Type                 CmdType  `protobuf:"varint,1,opt,name=Type,proto3,enum=message.CmdType" json:"Type,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgCmd) Reset()         { *m = MsgCmd{} }
func (m *MsgCmd) String() string { return proto.CompactTextString(m) }
func (*MsgCmd) ProtoMessage()    {}
func (*MsgCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *MsgCmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgCmd.Unmarshal(m, b)
}
func (m *MsgCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgCmd.Marshal(b, m, deterministic)
}
func (m *MsgCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCmd.Merge(m, src)
}
func (m *MsgCmd) XXX_Size() int {
	return xxx_messageInfo_MsgCmd.Size(m)
}
func (m *MsgCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCmd.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCmd proto.InternalMessageInfo

func (m *MsgCmd) GetType() CmdType {
	if m != nil {
		return m.Type
	}
	return CmdType_Login
}

func (m *MsgCmd) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// 上行消息 pb结构
type UPMsg struct {
	Head                 *UPMsgHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	UPMsgBody            []byte     `protobuf:"bytes,2,opt,name=UPMsgBody,proto3" json:"UPMsgBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UPMsg) Reset()         { *m = UPMsg{} }
func (m *UPMsg) String() string { return proto.CompactTextString(m) }
func (*UPMsg) ProtoMessage()    {}
func (*UPMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *UPMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UPMsg.Unmarshal(m, b)
}
func (m *UPMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UPMsg.Marshal(b, m, deterministic)
}
func (m *UPMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UPMsg.Merge(m, src)
}
func (m *UPMsg) XXX_Size() int {
	return xxx_messageInfo_UPMsg.Size(m)
}
func (m *UPMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UPMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UPMsg proto.InternalMessageInfo

func (m *UPMsg) GetHead() *UPMsgHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *UPMsg) GetUPMsgBody() []byte {
	if m != nil {
		return m.UPMsgBody
	}
	return nil
}

// 上行消息头 pb结构
type UPMsgHead struct {
	ClientID             uint64   `protobuf:"varint,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	ConnID               uint64   `protobuf:"varint,2,opt,name=ConnID,proto3" json:"ConnID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UPMsgHead) Reset()         { *m = UPMsgHead{} }
func (m *UPMsgHead) String() string { return proto.CompactTextString(m) }
func (*UPMsgHead) ProtoMessage()    {}
func (*UPMsgHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *UPMsgHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UPMsgHead.Unmarshal(m, b)
}
func (m *UPMsgHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UPMsgHead.Marshal(b, m, deterministic)
}
func (m *UPMsgHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UPMsgHead.Merge(m, src)
}
func (m *UPMsgHead) XXX_Size() int {
	return xxx_messageInfo_UPMsgHead.Size(m)
}
func (m *UPMsgHead) XXX_DiscardUnknown() {
	xxx_messageInfo_UPMsgHead.DiscardUnknown(m)
}

var xxx_messageInfo_UPMsgHead proto.InternalMessageInfo

func (m *UPMsgHead) GetClientID() uint64 {
	if m != nil {
		return m.ClientID
	}
	return 0
}

func (m *UPMsgHead) GetConnID() uint64 {
	if m != nil {
		return m.ConnID
	}
	return 0
}

type PushMsg struct {
	MsgID                uint64   `protobuf:"varint,1,opt,name=MsgID,proto3" json:"MsgID,omitempty"`
	SessionID            uint64   `protobuf:"varint,2,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Content              []byte   `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushMsg) Reset()         { *m = PushMsg{} }
func (m *PushMsg) String() string { return proto.CompactTextString(m) }
func (*PushMsg) ProtoMessage()    {}
func (*PushMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *PushMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushMsg.Unmarshal(m, b)
}
func (m *PushMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushMsg.Marshal(b, m, deterministic)
}
func (m *PushMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMsg.Merge(m, src)
}
func (m *PushMsg) XXX_Size() int {
	return xxx_messageInfo_PushMsg.Size(m)
}
func (m *PushMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PushMsg proto.InternalMessageInfo

func (m *PushMsg) GetMsgID() uint64 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

func (m *PushMsg) GetSessionID() uint64 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

func (m *PushMsg) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// ACK 消息
type ACKMsg struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Type                 CmdType  `protobuf:"varint,3,opt,name=Type,proto3,enum=message.CmdType" json:"Type,omitempty"`
	ConnID               uint64   `protobuf:"varint,4,opt,name=ConnID,proto3" json:"ConnID,omitempty"`
	ClientID             uint64   `protobuf:"varint,5,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	SessionID            uint64   `protobuf:"varint,6,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	MsgID                uint64   `protobuf:"varint,7,opt,name=MsgID,proto3" json:"MsgID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACKMsg) Reset()         { *m = ACKMsg{} }
func (m *ACKMsg) String() string { return proto.CompactTextString(m) }
func (*ACKMsg) ProtoMessage()    {}
func (*ACKMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *ACKMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACKMsg.Unmarshal(m, b)
}
func (m *ACKMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACKMsg.Marshal(b, m, deterministic)
}
func (m *ACKMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACKMsg.Merge(m, src)
}
func (m *ACKMsg) XXX_Size() int {
	return xxx_messageInfo_ACKMsg.Size(m)
}
func (m *ACKMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ACKMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ACKMsg proto.InternalMessageInfo

func (m *ACKMsg) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ACKMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ACKMsg) GetType() CmdType {
	if m != nil {
		return m.Type
	}
	return CmdType_Login
}

func (m *ACKMsg) GetConnID() uint64 {
	if m != nil {
		return m.ConnID
	}
	return 0
}

func (m *ACKMsg) GetClientID() uint64 {
	if m != nil {
		return m.ClientID
	}
	return 0
}

func (m *ACKMsg) GetSessionID() uint64 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

func (m *ACKMsg) GetMsgID() uint64 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

// 登陆消息
type LoginMsgHead struct {
	DeviceID             uint64   `protobuf:"varint,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginMsgHead) Reset()         { *m = LoginMsgHead{} }
func (m *LoginMsgHead) String() string { return proto.CompactTextString(m) }
func (*LoginMsgHead) ProtoMessage()    {}
func (*LoginMsgHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *LoginMsgHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginMsgHead.Unmarshal(m, b)
}
func (m *LoginMsgHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginMsgHead.Marshal(b, m, deterministic)
}
func (m *LoginMsgHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginMsgHead.Merge(m, src)
}
func (m *LoginMsgHead) XXX_Size() int {
	return xxx_messageInfo_LoginMsgHead.Size(m)
}
func (m *LoginMsgHead) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginMsgHead.DiscardUnknown(m)
}

var xxx_messageInfo_LoginMsgHead proto.InternalMessageInfo

func (m *LoginMsgHead) GetDeviceID() uint64 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

type LoginMsg struct {
	Head                 *LoginMsgHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	LoginMsgBody         []byte        `protobuf:"bytes,2,opt,name=LoginMsgBody,proto3" json:"LoginMsgBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LoginMsg) Reset()         { *m = LoginMsg{} }
func (m *LoginMsg) String() string { return proto.CompactTextString(m) }
func (*LoginMsg) ProtoMessage()    {}
func (*LoginMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *LoginMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginMsg.Unmarshal(m, b)
}
func (m *LoginMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginMsg.Marshal(b, m, deterministic)
}
func (m *LoginMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginMsg.Merge(m, src)
}
func (m *LoginMsg) XXX_Size() int {
	return xxx_messageInfo_LoginMsg.Size(m)
}
func (m *LoginMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginMsg.DiscardUnknown(m)
}

var xxx_messageInfo_LoginMsg proto.InternalMessageInfo

func (m *LoginMsg) GetHead() *LoginMsgHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *LoginMsg) GetLoginMsgBody() []byte {
	if m != nil {
		return m.LoginMsgBody
	}
	return nil
}

// 心跳消息
type HeartbeatMsgHead struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatMsgHead) Reset()         { *m = HeartbeatMsgHead{} }
func (m *HeartbeatMsgHead) String() string { return proto.CompactTextString(m) }
func (*HeartbeatMsgHead) ProtoMessage()    {}
func (*HeartbeatMsgHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{7}
}

func (m *HeartbeatMsgHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatMsgHead.Unmarshal(m, b)
}
func (m *HeartbeatMsgHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatMsgHead.Marshal(b, m, deterministic)
}
func (m *HeartbeatMsgHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatMsgHead.Merge(m, src)
}
func (m *HeartbeatMsgHead) XXX_Size() int {
	return xxx_messageInfo_HeartbeatMsgHead.Size(m)
}
func (m *HeartbeatMsgHead) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatMsgHead.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatMsgHead proto.InternalMessageInfo

type HeartbeatMsg struct {
	Head                 *HeartbeatMsgHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	HeartbeatMsgBody     []byte            `protobuf:"bytes,2,opt,name=HeartbeatMsgBody,proto3" json:"HeartbeatMsgBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HeartbeatMsg) Reset()         { *m = HeartbeatMsg{} }
func (m *HeartbeatMsg) String() string { return proto.CompactTextString(m) }
func (*HeartbeatMsg) ProtoMessage()    {}
func (*HeartbeatMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{8}
}

func (m *HeartbeatMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatMsg.Unmarshal(m, b)
}
func (m *HeartbeatMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatMsg.Marshal(b, m, deterministic)
}
func (m *HeartbeatMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatMsg.Merge(m, src)
}
func (m *HeartbeatMsg) XXX_Size() int {
	return xxx_messageInfo_HeartbeatMsg.Size(m)
}
func (m *HeartbeatMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatMsg proto.InternalMessageInfo

func (m *HeartbeatMsg) GetHead() *HeartbeatMsgHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *HeartbeatMsg) GetHeartbeatMsgBody() []byte {
	if m != nil {
		return m.HeartbeatMsgBody
	}
	return nil
}

// 重连消息
type ReConnMsgHead struct {
	ConnID               uint64   `protobuf:"varint,1,opt,name=ConnID,proto3" json:"ConnID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReConnMsgHead) Reset()         { *m = ReConnMsgHead{} }
func (m *ReConnMsgHead) String() string { return proto.CompactTextString(m) }
func (*ReConnMsgHead) ProtoMessage()    {}
func (*ReConnMsgHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{9}
}

func (m *ReConnMsgHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReConnMsgHead.Unmarshal(m, b)
}
func (m *ReConnMsgHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReConnMsgHead.Marshal(b, m, deterministic)
}
func (m *ReConnMsgHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReConnMsgHead.Merge(m, src)
}
func (m *ReConnMsgHead) XXX_Size() int {
	return xxx_messageInfo_ReConnMsgHead.Size(m)
}
func (m *ReConnMsgHead) XXX_DiscardUnknown() {
	xxx_messageInfo_ReConnMsgHead.DiscardUnknown(m)
}

var xxx_messageInfo_ReConnMsgHead proto.InternalMessageInfo

func (m *ReConnMsgHead) GetConnID() uint64 {
	if m != nil {
		return m.ConnID
	}
	return 0
}

type ReConnMsg struct {
	Head                 *ReConnMsgHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	ReConnMsgBody        []byte         `protobuf:"bytes,2,opt,name=ReConnMsgBody,proto3" json:"ReConnMsgBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReConnMsg) Reset()         { *m = ReConnMsg{} }
func (m *ReConnMsg) String() string { return proto.CompactTextString(m) }
func (*ReConnMsg) ProtoMessage()    {}
func (*ReConnMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{10}
}

func (m *ReConnMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReConnMsg.Unmarshal(m, b)
}
func (m *ReConnMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReConnMsg.Marshal(b, m, deterministic)
}
func (m *ReConnMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReConnMsg.Merge(m, src)
}
func (m *ReConnMsg) XXX_Size() int {
	return xxx_messageInfo_ReConnMsg.Size(m)
}
func (m *ReConnMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ReConnMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ReConnMsg proto.InternalMessageInfo

func (m *ReConnMsg) GetHead() *ReConnMsgHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *ReConnMsg) GetReConnMsgBody() []byte {
	if m != nil {
		return m.ReConnMsgBody
	}
	return nil
}

func init() {
	proto.RegisterEnum("message.CmdType", CmdType_name, CmdType_value)
	proto.RegisterType((*MsgCmd)(nil), "message.MsgCmd")
	proto.RegisterType((*UPMsg)(nil), "message.UPMsg")
	proto.RegisterType((*UPMsgHead)(nil), "message.UPMsgHead")
	proto.RegisterType((*PushMsg)(nil), "message.PushMsg")
	proto.RegisterType((*ACKMsg)(nil), "message.ACKMsg")
	proto.RegisterType((*LoginMsgHead)(nil), "message.LoginMsgHead")
	proto.RegisterType((*LoginMsg)(nil), "message.LoginMsg")
	proto.RegisterType((*HeartbeatMsgHead)(nil), "message.HeartbeatMsgHead")
	proto.RegisterType((*HeartbeatMsg)(nil), "message.HeartbeatMsg")
	proto.RegisterType((*ReConnMsgHead)(nil), "message.ReConnMsgHead")
	proto.RegisterType((*ReConnMsg)(nil), "message.ReConnMsg")
}

func init() {
	proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd)
}

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x8b, 0xd4, 0x30,
	0x14, 0xb5, 0xd3, 0xaf, 0xe9, 0x75, 0x46, 0xc2, 0x45, 0x97, 0x51, 0x7c, 0x58, 0xc2, 0xa2, 0xeb,
	0x80, 0x23, 0xac, 0x8f, 0x3e, 0xc8, 0x6e, 0xe7, 0x61, 0xd6, 0xb5, 0x30, 0x44, 0x17, 0x51, 0xf0,
	0xa1, 0x6b, 0x43, 0x2d, 0xec, 0x34, 0xcb, 0xa6, 0x0a, 0xf3, 0xdf, 0xfc, 0x71, 0x92, 0x34, 0x4d,
	0x93, 0x11, 0x7c, 0xeb, 0x3d, 0x37, 0xf7, 0xdc, 0x73, 0x4e, 0x52, 0x98, 0xef, 0xb8, 0x94, 0x65,
	0xcd, 0x57, 0x77, 0xf7, 0xa2, 0x13, 0x98, 0x9a, 0x92, 0x6e, 0x20, 0x29, 0x64, 0x9d, 0xef, 0x2a,
	0x3c, 0x81, 0xe8, 0xf3, 0xfe, 0x8e, 0x2f, 0x82, 0xe3, 0xe0, 0xf4, 0xd1, 0x19, 0x59, 0x0d, 0x03,
	0xf9, 0xae, 0x52, 0x38, 0xd3, 0x5d, 0x5c, 0x40, 0xba, 0x2d, 0xf7, 0xb7, 0xa2, 0xac, 0x16, 0x93,
	0xe3, 0xe0, 0x74, 0xc6, 0x86, 0x92, 0x16, 0x10, 0x5f, 0x6f, 0x0b, 0x59, 0xe3, 0x0b, 0x88, 0x36,
	0xbc, 0xac, 0x34, 0xd1, 0xc3, 0x33, 0xb4, 0x44, 0xba, 0xab, 0x3a, 0x4c, 0xf7, 0xf1, 0x39, 0x64,
	0x1a, 0xba, 0x10, 0xd5, 0xde, 0x90, 0x8d, 0x00, 0x7d, 0x6f, 0xba, 0xfa, 0xe8, 0x33, 0x98, 0xe6,
	0xb7, 0x0d, 0x6f, 0xbb, 0xcb, 0xb5, 0xa6, 0x8d, 0x98, 0xad, 0xf1, 0x08, 0x92, 0x5c, 0xb4, 0xed,
	0xe5, 0x5a, 0x73, 0x44, 0xcc, 0x54, 0xf4, 0x0b, 0xa4, 0xdb, 0x5f, 0xf2, 0xa7, 0x52, 0xf4, 0x18,
	0xe2, 0x42, 0xd6, 0x76, 0xb6, 0x2f, 0xd4, 0xfe, 0x4f, 0x5c, 0xca, 0x46, 0x8c, 0xb3, 0x23, 0xa0,
	0x8c, 0xe6, 0xa2, 0xed, 0x78, 0xdb, 0x2d, 0xc2, 0xde, 0xa8, 0x29, 0xe9, 0x9f, 0x00, 0x92, 0xf3,
	0xfc, 0x4a, 0x11, 0x23, 0x44, 0xb9, 0xa8, 0xfa, 0xcc, 0xe6, 0x4c, 0x7f, 0x23, 0x81, 0xb0, 0x90,
	0xb5, 0x26, 0xcc, 0x98, 0xfa, 0xb4, 0xc9, 0x86, 0xff, 0x4d, 0x76, 0xf4, 0x11, 0xb9, 0x3e, 0x3c,
	0xef, 0xf1, 0x81, 0x77, 0xcf, 0x42, 0x72, 0x68, 0xc1, 0xda, 0x4e, 0x1d, 0xdb, 0x74, 0x09, 0xb3,
	0x8f, 0xa2, 0x6e, 0x5a, 0x27, 0xdb, 0x35, 0xff, 0xdd, 0xfc, 0xe0, 0x63, 0xb6, 0x43, 0x4d, 0xbf,
	0xc2, 0x74, 0x38, 0x8b, 0xaf, 0xbc, 0x6b, 0x7d, 0x62, 0x5d, 0xb8, 0x64, 0xe6, 0x66, 0xe9, 0xb8,
	0xc2, 0xb9, 0x5c, 0x0f, 0xa3, 0x08, 0x64, 0xc3, 0xcb, 0xfb, 0xee, 0x86, 0x97, 0x9d, 0x99, 0xa6,
	0x0d, 0xcc, 0x5c, 0x0c, 0x5f, 0x7b, 0x2b, 0x9f, 0xda, 0x95, 0x87, 0x83, 0x66, 0xed, 0xd2, 0xa7,
	0x74, 0x56, 0xff, 0x83, 0xd3, 0x97, 0x30, 0x67, 0x5c, 0x25, 0x3c, 0xc4, 0x30, 0xc6, 0x1f, 0x78,
	0xcf, 0xe8, 0x3b, 0x64, 0xf6, 0x20, 0x2e, 0x3d, 0x41, 0x47, 0x56, 0x90, 0x47, 0x65, 0xd4, 0x9c,
	0x38, 0x1b, 0x1c, 0x29, 0x3e, 0xb8, 0xfc, 0x00, 0xa9, 0x79, 0x06, 0x98, 0x41, 0xac, 0x13, 0x22,
	0x0f, 0x70, 0x0e, 0x99, 0x55, 0x4c, 0x02, 0x04, 0x48, 0xfa, 0x29, 0x32, 0xc1, 0x14, 0xc2, 0xf3,
	0xfc, 0x8a, 0x84, 0x98, 0xc0, 0xe4, 0x7a, 0x4b, 0x22, 0x9c, 0x42, 0xa4, 0xde, 0x39, 0x89, 0x2f,
	0x66, 0xdf, 0x60, 0xf5, 0xe6, 0x9d, 0xd1, 0x74, 0x93, 0xe8, 0x3f, 0xfd, 0xed, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x17, 0xa6, 0xc9, 0x47, 0xfa, 0x03, 0x00, 0x00,
}
