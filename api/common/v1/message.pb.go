// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: api/common/v1/message.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CmdType int32

const (
	CmdType_CMD_TYPE_UNSPECIFIED CmdType = 0
	CmdType_LOGIN                CmdType = 1
	CmdType_HEARTBEAT            CmdType = 2
	CmdType_RE_CONN              CmdType = 3
	CmdType_ACK                  CmdType = 4
	CmdType_UP                   CmdType = 5 // 上行消息
	CmdType_PUSH                 CmdType = 6 // 下行 推送消息
)

// Enum value maps for CmdType.
var (
	CmdType_name = map[int32]string{
		0: "CMD_TYPE_UNSPECIFIED",
		1: "LOGIN",
		2: "HEARTBEAT",
		3: "RE_CONN",
		4: "ACK",
		5: "UP",
		6: "PUSH",
	}
	CmdType_value = map[string]int32{
		"CMD_TYPE_UNSPECIFIED": 0,
		"LOGIN":                1,
		"HEARTBEAT":            2,
		"RE_CONN":              3,
		"ACK":                  4,
		"UP":                   5,
		"PUSH":                 6,
	}
)

func (x CmdType) Enum() *CmdType {
	p := new(CmdType)
	*p = x
	return p
}

func (x CmdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_common_v1_message_proto_enumTypes[0].Descriptor()
}

func (CmdType) Type() protoreflect.EnumType {
	return &file_api_common_v1_message_proto_enumTypes[0]
}

func (x CmdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdType.Descriptor instead.
func (CmdType) EnumDescriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{0}
}

// 顶层消息实体 端->State->业务层 实际上使用的消息实体 Gateway则略过这个实体，只负责透传
type MajorMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd     CmdType `protobuf:"varint,1,opt,name=cmd,proto3,enum=api.common.v1.CmdType" json:"cmd,omitempty"`
	Payload []byte  `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *MajorMsg) Reset() {
	*x = MajorMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MajorMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MajorMsg) ProtoMessage() {}

func (x *MajorMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MajorMsg.ProtoReflect.Descriptor instead.
func (*MajorMsg) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *MajorMsg) GetCmd() CmdType {
	if x != nil {
		return x.Cmd
	}
	return CmdType_CMD_TYPE_UNSPECIFIED
}

func (x *MajorMsg) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// 上行消息 pb结构
type UpMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *UpMsgHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Body []byte     `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *UpMsg) Reset() {
	*x = UpMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpMsg) ProtoMessage() {}

func (x *UpMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpMsg.ProtoReflect.Descriptor instead.
func (*UpMsg) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *UpMsg) GetHead() *UpMsgHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *UpMsg) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// 上行消息头 pb结构
type UpMsgHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  uint64 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ConnId    uint64 `protobuf:"varint,2,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
	SessionId string `protobuf:"bytes,3,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
}

func (x *UpMsgHead) Reset() {
	*x = UpMsgHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpMsgHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpMsgHead) ProtoMessage() {}

func (x *UpMsgHead) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpMsgHead.ProtoReflect.Descriptor instead.
func (*UpMsgHead) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *UpMsgHead) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *UpMsgHead) GetConnId() uint64 {
	if x != nil {
		return x.ConnId
	}
	return 0
}

func (x *UpMsgHead) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

// 下行消息
type PushMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId     uint64 `protobuf:"varint,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	SessionId uint64 `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Content   []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PushMsg) Reset() {
	*x = PushMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsg) ProtoMessage() {}

func (x *PushMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsg.ProtoReflect.Descriptor instead.
func (*PushMsg) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *PushMsg) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *PushMsg) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *PushMsg) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// Ack 消息
type AckMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      uint32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Cmd       CmdType `protobuf:"varint,3,opt,name=cmd,proto3,enum=api.common.v1.CmdType" json:"cmd,omitempty"`
	ConnId    uint64  `protobuf:"varint,4,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
	ClientId  uint64  `protobuf:"varint,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	SessionId uint64  `protobuf:"varint,6,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	MsgId     uint64  `protobuf:"varint,7,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *AckMsg) Reset() {
	*x = AckMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMsg) ProtoMessage() {}

func (x *AckMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMsg.ProtoReflect.Descriptor instead.
func (*AckMsg) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *AckMsg) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AckMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AckMsg) GetCmd() CmdType {
	if x != nil {
		return x.Cmd
	}
	return CmdType_CMD_TYPE_UNSPECIFIED
}

func (x *AckMsg) GetConnId() uint64 {
	if x != nil {
		return x.ConnId
	}
	return 0
}

func (x *AckMsg) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *AckMsg) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *AckMsg) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

// 登陆消息
type LoginMsgHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId uint64 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *LoginMsgHead) Reset() {
	*x = LoginMsgHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMsgHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMsgHead) ProtoMessage() {}

func (x *LoginMsgHead) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMsgHead.ProtoReflect.Descriptor instead.
func (*LoginMsgHead) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *LoginMsgHead) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

type LoginMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *LoginMsgHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Body []byte        `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *LoginMsg) Reset() {
	*x = LoginMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMsg) ProtoMessage() {}

func (x *LoginMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMsg.ProtoReflect.Descriptor instead.
func (*LoginMsg) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{6}
}

func (x *LoginMsg) GetHead() *LoginMsgHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *LoginMsg) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// 心跳消息
type HeartbeatMsgHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatMsgHead) Reset() {
	*x = HeartbeatMsgHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatMsgHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatMsgHead) ProtoMessage() {}

func (x *HeartbeatMsgHead) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatMsgHead.ProtoReflect.Descriptor instead.
func (*HeartbeatMsgHead) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{7}
}

type HeartbeatMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *HeartbeatMsgHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Body []byte            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *HeartbeatMsg) Reset() {
	*x = HeartbeatMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatMsg) ProtoMessage() {}

func (x *HeartbeatMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatMsg.ProtoReflect.Descriptor instead.
func (*HeartbeatMsg) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{8}
}

func (x *HeartbeatMsg) GetHead() *HeartbeatMsgHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *HeartbeatMsg) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// 重连消息
type ReConnMsgHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnId uint64 `protobuf:"varint,1,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
}

func (x *ReConnMsgHead) Reset() {
	*x = ReConnMsgHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReConnMsgHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReConnMsgHead) ProtoMessage() {}

func (x *ReConnMsgHead) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReConnMsgHead.ProtoReflect.Descriptor instead.
func (*ReConnMsgHead) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{9}
}

func (x *ReConnMsgHead) GetConnId() uint64 {
	if x != nil {
		return x.ConnId
	}
	return 0
}

type ReConnMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *ReConnMsgHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Body []byte         `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ReConnMsg) Reset() {
	*x = ReConnMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_message_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReConnMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReConnMsg) ProtoMessage() {}

func (x *ReConnMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_message_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReConnMsg.ProtoReflect.Descriptor instead.
func (*ReConnMsg) Descriptor() ([]byte, []int) {
	return file_api_common_v1_message_proto_rawDescGZIP(), []int{10}
}

func (x *ReConnMsg) GetHead() *ReConnMsgHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *ReConnMsg) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_api_common_v1_message_proto protoreflect.FileDescriptor

var file_api_common_v1_message_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x4e, 0x0a, 0x08,
	0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x63,
	0x6d, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x49, 0x0a, 0x05,
	0x55, 0x70, 0x4d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68,
	0x65, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x5f, 0x0a, 0x09, 0x55, 0x70, 0x4d, 0x73, 0x67,
	0x48, 0x65, 0x61, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x07, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0xc4, 0x01, 0x0a, 0x06, 0x41, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0c, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x4d, 0x73, 0x67, 0x12, 0x2f, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04,
	0x68, 0x65, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x22, 0x57, 0x0a, 0x0c,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x33, 0x0a, 0x04,
	0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x28, 0x0a, 0x0d, 0x52, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x4d,
	0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x22,
	0x51, 0x0a, 0x09, 0x52, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x30, 0x0a, 0x04,
	0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x2a, 0x65, 0x0a, 0x07, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x14, 0x43, 0x4d, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x49, 0x4e,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x10, 0x03, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x43, 0x4b, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x05, 0x12,
	0x08, 0x0a, 0x04, 0x50, 0x55, 0x53, 0x48, 0x10, 0x06, 0x42, 0x3f, 0x0a, 0x0d, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x79, 0x73, 0x69, 0x63, 0x73,
	0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2d, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_common_v1_message_proto_rawDescOnce sync.Once
	file_api_common_v1_message_proto_rawDescData = file_api_common_v1_message_proto_rawDesc
)

func file_api_common_v1_message_proto_rawDescGZIP() []byte {
	file_api_common_v1_message_proto_rawDescOnce.Do(func() {
		file_api_common_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_v1_message_proto_rawDescData)
	})
	return file_api_common_v1_message_proto_rawDescData
}

var file_api_common_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_common_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_common_v1_message_proto_goTypes = []any{
	(CmdType)(0),             // 0: api.common.v1.CmdType
	(*MajorMsg)(nil),         // 1: api.common.v1.MajorMsg
	(*UpMsg)(nil),            // 2: api.common.v1.UpMsg
	(*UpMsgHead)(nil),        // 3: api.common.v1.UpMsgHead
	(*PushMsg)(nil),          // 4: api.common.v1.PushMsg
	(*AckMsg)(nil),           // 5: api.common.v1.AckMsg
	(*LoginMsgHead)(nil),     // 6: api.common.v1.LoginMsgHead
	(*LoginMsg)(nil),         // 7: api.common.v1.LoginMsg
	(*HeartbeatMsgHead)(nil), // 8: api.common.v1.HeartbeatMsgHead
	(*HeartbeatMsg)(nil),     // 9: api.common.v1.HeartbeatMsg
	(*ReConnMsgHead)(nil),    // 10: api.common.v1.ReConnMsgHead
	(*ReConnMsg)(nil),        // 11: api.common.v1.ReConnMsg
}
var file_api_common_v1_message_proto_depIdxs = []int32{
	0,  // 0: api.common.v1.MajorMsg.cmd:type_name -> api.common.v1.CmdType
	3,  // 1: api.common.v1.UpMsg.head:type_name -> api.common.v1.UpMsgHead
	0,  // 2: api.common.v1.AckMsg.cmd:type_name -> api.common.v1.CmdType
	6,  // 3: api.common.v1.LoginMsg.head:type_name -> api.common.v1.LoginMsgHead
	8,  // 4: api.common.v1.HeartbeatMsg.head:type_name -> api.common.v1.HeartbeatMsgHead
	10, // 5: api.common.v1.ReConnMsg.head:type_name -> api.common.v1.ReConnMsgHead
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_common_v1_message_proto_init() }
func file_api_common_v1_message_proto_init() {
	if File_api_common_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_common_v1_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MajorMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_common_v1_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_common_v1_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpMsgHead); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_common_v1_message_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PushMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_common_v1_message_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AckMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_common_v1_message_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LoginMsgHead); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_common_v1_message_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*LoginMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_common_v1_message_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*HeartbeatMsgHead); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_common_v1_message_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*HeartbeatMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_common_v1_message_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ReConnMsgHead); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_common_v1_message_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ReConnMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_common_v1_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_v1_message_proto_goTypes,
		DependencyIndexes: file_api_common_v1_message_proto_depIdxs,
		EnumInfos:         file_api_common_v1_message_proto_enumTypes,
		MessageInfos:      file_api_common_v1_message_proto_msgTypes,
	}.Build()
	File_api_common_v1_message_proto = out.File
	file_api_common_v1_message_proto_rawDesc = nil
	file_api_common_v1_message_proto_goTypes = nil
	file_api_common_v1_message_proto_depIdxs = nil
}
