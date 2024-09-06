// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: proto/raft.proto

package raftpb

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

type MessageType int32

const (
	MessageType_VOTE                  MessageType = 0
	MessageType_VOTE_RESP             MessageType = 1
	MessageType_HEARTBEAT             MessageType = 2
	MessageType_HEARTBEAT_RESP        MessageType = 3
	MessageType_APPEND_ENTRY          MessageType = 4
	MessageType_APPEND_ENTRY_RESP     MessageType = 5
	MessageType_PROPOSE               MessageType = 6
	MessageType_PROPOSE_RESP          MessageType = 7
	MessageType_INSTALL_SNAPSHOT      MessageType = 8
	MessageType_INSTALL_SNAPSHOT_RESP MessageType = 9
	MessageType_READINDEX             MessageType = 10
	MessageType_READINDEX_RESP        MessageType = 11
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0:  "VOTE",
		1:  "VOTE_RESP",
		2:  "HEARTBEAT",
		3:  "HEARTBEAT_RESP",
		4:  "APPEND_ENTRY",
		5:  "APPEND_ENTRY_RESP",
		6:  "PROPOSE",
		7:  "PROPOSE_RESP",
		8:  "INSTALL_SNAPSHOT",
		9:  "INSTALL_SNAPSHOT_RESP",
		10: "READINDEX",
		11: "READINDEX_RESP",
	}
	MessageType_value = map[string]int32{
		"VOTE":                  0,
		"VOTE_RESP":             1,
		"HEARTBEAT":             2,
		"HEARTBEAT_RESP":        3,
		"APPEND_ENTRY":          4,
		"APPEND_ENTRY_RESP":     5,
		"PROPOSE":               6,
		"PROPOSE_RESP":          7,
		"INSTALL_SNAPSHOT":      8,
		"INSTALL_SNAPSHOT_RESP": 9,
		"READINDEX":             10,
		"READINDEX_RESP":        11,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_raft_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_proto_raft_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{0}
}

type EntryType int32

const (
	EntryType_NORMAL        EntryType = 0
	EntryType_MEMBER_CHNAGE EntryType = 1
)

// Enum value maps for EntryType.
var (
	EntryType_name = map[int32]string{
		0: "NORMAL",
		1: "MEMBER_CHNAGE",
	}
	EntryType_value = map[string]int32{
		"NORMAL":        0,
		"MEMBER_CHNAGE": 1,
	}
)

func (x EntryType) Enum() *EntryType {
	p := new(EntryType)
	*p = x
	return p
}

func (x EntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_raft_proto_enumTypes[1].Descriptor()
}

func (EntryType) Type() protoreflect.EnumType {
	return &file_proto_raft_proto_enumTypes[1]
}

func (x EntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntryType.Descriptor instead.
func (EntryType) EnumDescriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{1}
}

type MemberChangeType int32

const (
	MemberChangeType_ADD_NODE    MemberChangeType = 0
	MemberChangeType_REMOVE_NODE MemberChangeType = 1
)

// Enum value maps for MemberChangeType.
var (
	MemberChangeType_name = map[int32]string{
		0: "ADD_NODE",
		1: "REMOVE_NODE",
	}
	MemberChangeType_value = map[string]int32{
		"ADD_NODE":    0,
		"REMOVE_NODE": 1,
	}
)

func (x MemberChangeType) Enum() *MemberChangeType {
	p := new(MemberChangeType)
	*p = x
	return p
}

func (x MemberChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_raft_proto_enumTypes[2].Descriptor()
}

func (MemberChangeType) Type() protoreflect.EnumType {
	return &file_proto_raft_proto_enumTypes[2]
}

func (x MemberChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberChangeType.Descriptor instead.
func (MemberChangeType) EnumDescriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{2}
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  EntryType `protobuf:"varint,1,opt,name=type,proto3,enum=raft.EntryType" json:"type,omitempty"`
	Term  uint64    `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Index uint64    `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Data  []byte    `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{0}
}

func (x *LogEntry) GetType() EntryType {
	if x != nil {
		return x.Type
	}
	return EntryType_NORMAL
}

func (x *LogEntry) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *LogEntry) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *LogEntry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type RaftMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType      MessageType `protobuf:"varint,1,opt,name=msgType,proto3,enum=raft.MessageType" json:"msgType,omitempty"`
	Term         uint64      `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	From         uint64      `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	To           uint64      `protobuf:"varint,4,opt,name=to,proto3" json:"to,omitempty"`
	LastLogIndex uint64      `protobuf:"varint,5,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	LastLogTerm  uint64      `protobuf:"varint,6,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
	LastCommit   uint64      `protobuf:"varint,7,opt,name=lastCommit,proto3" json:"lastCommit,omitempty"`
	Entry        []*LogEntry `protobuf:"bytes,8,rep,name=entry,proto3" json:"entry,omitempty"`
	Success      bool        `protobuf:"varint,9,opt,name=success,proto3" json:"success,omitempty"`
	Context      []byte      `protobuf:"bytes,10,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *RaftMessage) Reset() {
	*x = RaftMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftMessage) ProtoMessage() {}

func (x *RaftMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftMessage.ProtoReflect.Descriptor instead.
func (*RaftMessage) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{1}
}

func (x *RaftMessage) GetMsgType() MessageType {
	if x != nil {
		return x.MsgType
	}
	return MessageType_VOTE
}

func (x *RaftMessage) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RaftMessage) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *RaftMessage) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *RaftMessage) GetLastLogIndex() uint64 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *RaftMessage) GetLastLogTerm() uint64 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

func (x *RaftMessage) GetLastCommit() uint64 {
	if x != nil {
		return x.LastCommit
	}
	return 0
}

func (x *RaftMessage) GetEntry() []*LogEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *RaftMessage) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RaftMessage) GetContext() []byte {
	if x != nil {
		return x.Context
	}
	return nil
}

var File_proto_raft_proto protoreflect.FileDescriptor

var file_proto_raft_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x66, 0x74, 0x22, 0x6d, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb2, 0x02, 0x0a, 0x0b, 0x52, 0x61, 0x66, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x22, 0x0a, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65,
	0x72, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2a, 0xe5, 0x01, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x56, 0x4f, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x56, 0x4f, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45,
	0x41, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41,
	0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x50, 0x50, 0x45,
	0x4e, 0x44, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x50,
	0x50, 0x45, 0x4e, 0x44, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10,
	0x05, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x45, 0x10, 0x06, 0x12, 0x10,
	0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x07,
	0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x4e, 0x41, 0x50,
	0x53, 0x48, 0x4f, 0x54, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c,
	0x4c, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10,
	0x09, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x41, 0x44, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x10, 0x0a,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x41, 0x44, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x52, 0x45,
	0x53, 0x50, 0x10, 0x0b, 0x2a, 0x2a, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x43, 0x48, 0x4e, 0x41, 0x47, 0x45, 0x10, 0x01,
	0x2a, 0x31, 0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x44, 0x44, 0x5f, 0x4e, 0x4f, 0x44, 0x45,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x4e, 0x4f, 0x44,
	0x45, 0x10, 0x01, 0x32, 0x3f, 0x0a, 0x04, 0x52, 0x61, 0x66, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x12, 0x11, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e,
	0x52, 0x61, 0x66, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x72, 0x61,
	0x66, 0x74, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x1e, 0x5a, 0x1c, 0x65, 0x61, 0x73, 0x79, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x3b, 0x72, 0x61,
	0x66, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_raft_proto_rawDescOnce sync.Once
	file_proto_raft_proto_rawDescData = file_proto_raft_proto_rawDesc
)

func file_proto_raft_proto_rawDescGZIP() []byte {
	file_proto_raft_proto_rawDescOnce.Do(func() {
		file_proto_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_raft_proto_rawDescData)
	})
	return file_proto_raft_proto_rawDescData
}

var file_proto_raft_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_raft_proto_goTypes = []interface{}{
	(MessageType)(0),      // 0: raft.MessageType
	(EntryType)(0),        // 1: raft.EntryType
	(MemberChangeType)(0), // 2: raft.MemberChangeType
	(*LogEntry)(nil),      // 3: raft.LogEntry
	(*RaftMessage)(nil),   // 4: raft.RaftMessage
}
var file_proto_raft_proto_depIdxs = []int32{
	1, // 0: raft.LogEntry.type:type_name -> raft.EntryType
	0, // 1: raft.RaftMessage.msgType:type_name -> raft.MessageType
	3, // 2: raft.RaftMessage.entry:type_name -> raft.LogEntry
	4, // 3: raft.Raft.consensus:input_type -> raft.RaftMessage
	4, // 4: raft.Raft.consensus:output_type -> raft.RaftMessage
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_raft_proto_init() }
func file_proto_raft_proto_init() {
	if File_proto_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
		file_proto_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftMessage); i {
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
			RawDescriptor: file_proto_raft_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_raft_proto_goTypes,
		DependencyIndexes: file_proto_raft_proto_depIdxs,
		EnumInfos:         file_proto_raft_proto_enumTypes,
		MessageInfos:      file_proto_raft_proto_msgTypes,
	}.Build()
	File_proto_raft_proto = out.File
	file_proto_raft_proto_rawDesc = nil
	file_proto_raft_proto_goTypes = nil
	file_proto_raft_proto_depIdxs = nil
}
