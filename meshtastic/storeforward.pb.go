// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: meshtastic/storeforward.proto

package generated

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

// 001 - 063 = From Router
// 064 - 127 = From Client
type StoreAndForward_RequestResponse int32

const (
	// Unset/unused
	StoreAndForward_UNSET StoreAndForward_RequestResponse = 0
	// Router is an in error state.
	StoreAndForward_ROUTER_ERROR StoreAndForward_RequestResponse = 1
	// Router heartbeat
	StoreAndForward_ROUTER_HEARTBEAT StoreAndForward_RequestResponse = 2
	// Router has requested the client respond. This can work as a
	// "are you there" message.
	StoreAndForward_ROUTER_PING StoreAndForward_RequestResponse = 3
	// The response to a "Ping"
	StoreAndForward_ROUTER_PONG StoreAndForward_RequestResponse = 4
	// Router is currently busy. Please try again later.
	StoreAndForward_ROUTER_BUSY StoreAndForward_RequestResponse = 5
	// Router is responding to a request for history.
	StoreAndForward_ROUTER_HISTORY StoreAndForward_RequestResponse = 6
	// Router is responding to a request for stats.
	StoreAndForward_ROUTER_STATS StoreAndForward_RequestResponse = 7
	// Router sends a text message from its history that was a direct message.
	StoreAndForward_ROUTER_TEXT_DIRECT StoreAndForward_RequestResponse = 8
	// Router sends a text message from its history that was a broadcast.
	StoreAndForward_ROUTER_TEXT_BROADCAST StoreAndForward_RequestResponse = 9
	// Client is an in error state.
	StoreAndForward_CLIENT_ERROR StoreAndForward_RequestResponse = 64
	// Client has requested a replay from the router.
	StoreAndForward_CLIENT_HISTORY StoreAndForward_RequestResponse = 65
	// Client has requested stats from the router.
	StoreAndForward_CLIENT_STATS StoreAndForward_RequestResponse = 66
	// Client has requested the router respond. This can work as a
	// "are you there" message.
	StoreAndForward_CLIENT_PING StoreAndForward_RequestResponse = 67
	// The response to a "Ping"
	StoreAndForward_CLIENT_PONG StoreAndForward_RequestResponse = 68
	// Client has requested that the router abort processing the client's request
	StoreAndForward_CLIENT_ABORT StoreAndForward_RequestResponse = 106
)

// Enum value maps for StoreAndForward_RequestResponse.
var (
	StoreAndForward_RequestResponse_name = map[int32]string{
		0:   "UNSET",
		1:   "ROUTER_ERROR",
		2:   "ROUTER_HEARTBEAT",
		3:   "ROUTER_PING",
		4:   "ROUTER_PONG",
		5:   "ROUTER_BUSY",
		6:   "ROUTER_HISTORY",
		7:   "ROUTER_STATS",
		8:   "ROUTER_TEXT_DIRECT",
		9:   "ROUTER_TEXT_BROADCAST",
		64:  "CLIENT_ERROR",
		65:  "CLIENT_HISTORY",
		66:  "CLIENT_STATS",
		67:  "CLIENT_PING",
		68:  "CLIENT_PONG",
		106: "CLIENT_ABORT",
	}
	StoreAndForward_RequestResponse_value = map[string]int32{
		"UNSET":                 0,
		"ROUTER_ERROR":          1,
		"ROUTER_HEARTBEAT":      2,
		"ROUTER_PING":           3,
		"ROUTER_PONG":           4,
		"ROUTER_BUSY":           5,
		"ROUTER_HISTORY":        6,
		"ROUTER_STATS":          7,
		"ROUTER_TEXT_DIRECT":    8,
		"ROUTER_TEXT_BROADCAST": 9,
		"CLIENT_ERROR":          64,
		"CLIENT_HISTORY":        65,
		"CLIENT_STATS":          66,
		"CLIENT_PING":           67,
		"CLIENT_PONG":           68,
		"CLIENT_ABORT":          106,
	}
)

func (x StoreAndForward_RequestResponse) Enum() *StoreAndForward_RequestResponse {
	p := new(StoreAndForward_RequestResponse)
	*p = x
	return p
}

func (x StoreAndForward_RequestResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StoreAndForward_RequestResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_meshtastic_storeforward_proto_enumTypes[0].Descriptor()
}

func (StoreAndForward_RequestResponse) Type() protoreflect.EnumType {
	return &file_meshtastic_storeforward_proto_enumTypes[0]
}

func (x StoreAndForward_RequestResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StoreAndForward_RequestResponse.Descriptor instead.
func (StoreAndForward_RequestResponse) EnumDescriptor() ([]byte, []int) {
	return file_meshtastic_storeforward_proto_rawDescGZIP(), []int{0, 0}
}

// TODO: REPLACE
type StoreAndForward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: REPLACE
	Rr StoreAndForward_RequestResponse `protobuf:"varint,1,opt,name=rr,proto3,enum=meshtastic.StoreAndForward_RequestResponse" json:"rr,omitempty"`
	// TODO: REPLACE
	//
	// Types that are assignable to Variant:
	//
	//	*StoreAndForward_Stats
	//	*StoreAndForward_History_
	//	*StoreAndForward_Heartbeat_
	//	*StoreAndForward_Text
	Variant isStoreAndForward_Variant `protobuf_oneof:"variant"`
}

func (x *StoreAndForward) Reset() {
	*x = StoreAndForward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_storeforward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreAndForward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreAndForward) ProtoMessage() {}

func (x *StoreAndForward) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_storeforward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreAndForward.ProtoReflect.Descriptor instead.
func (*StoreAndForward) Descriptor() ([]byte, []int) {
	return file_meshtastic_storeforward_proto_rawDescGZIP(), []int{0}
}

func (x *StoreAndForward) GetRr() StoreAndForward_RequestResponse {
	if x != nil {
		return x.Rr
	}
	return StoreAndForward_UNSET
}

func (m *StoreAndForward) GetVariant() isStoreAndForward_Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (x *StoreAndForward) GetStats() *StoreAndForward_Statistics {
	if x, ok := x.GetVariant().(*StoreAndForward_Stats); ok {
		return x.Stats
	}
	return nil
}

func (x *StoreAndForward) GetHistory() *StoreAndForward_History {
	if x, ok := x.GetVariant().(*StoreAndForward_History_); ok {
		return x.History
	}
	return nil
}

func (x *StoreAndForward) GetHeartbeat() *StoreAndForward_Heartbeat {
	if x, ok := x.GetVariant().(*StoreAndForward_Heartbeat_); ok {
		return x.Heartbeat
	}
	return nil
}

func (x *StoreAndForward) GetText() []byte {
	if x, ok := x.GetVariant().(*StoreAndForward_Text); ok {
		return x.Text
	}
	return nil
}

type isStoreAndForward_Variant interface {
	isStoreAndForward_Variant()
}

type StoreAndForward_Stats struct {
	// TODO: REPLACE
	Stats *StoreAndForward_Statistics `protobuf:"bytes,2,opt,name=stats,proto3,oneof"`
}

type StoreAndForward_History_ struct {
	// TODO: REPLACE
	History *StoreAndForward_History `protobuf:"bytes,3,opt,name=history,proto3,oneof"`
}

type StoreAndForward_Heartbeat_ struct {
	// TODO: REPLACE
	Heartbeat *StoreAndForward_Heartbeat `protobuf:"bytes,4,opt,name=heartbeat,proto3,oneof"`
}

type StoreAndForward_Text struct {
	// Text from history message.
	Text []byte `protobuf:"bytes,5,opt,name=text,proto3,oneof"`
}

func (*StoreAndForward_Stats) isStoreAndForward_Variant() {}

func (*StoreAndForward_History_) isStoreAndForward_Variant() {}

func (*StoreAndForward_Heartbeat_) isStoreAndForward_Variant() {}

func (*StoreAndForward_Text) isStoreAndForward_Variant() {}

// TODO: REPLACE
type StoreAndForward_Statistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of messages we have ever seen
	MessagesTotal uint32 `protobuf:"varint,1,opt,name=messages_total,json=messagesTotal,proto3" json:"messages_total,omitempty"`
	// Number of messages we have currently saved our history.
	MessagesSaved uint32 `protobuf:"varint,2,opt,name=messages_saved,json=messagesSaved,proto3" json:"messages_saved,omitempty"`
	// Maximum number of messages we will save
	MessagesMax uint32 `protobuf:"varint,3,opt,name=messages_max,json=messagesMax,proto3" json:"messages_max,omitempty"`
	// Router uptime in seconds
	UpTime uint32 `protobuf:"varint,4,opt,name=up_time,json=upTime,proto3" json:"up_time,omitempty"`
	// Number of times any client sent a request to the S&F.
	Requests uint32 `protobuf:"varint,5,opt,name=requests,proto3" json:"requests,omitempty"`
	// Number of times the history was requested.
	RequestsHistory uint32 `protobuf:"varint,6,opt,name=requests_history,json=requestsHistory,proto3" json:"requests_history,omitempty"`
	// Is the heartbeat enabled on the server?
	Heartbeat bool `protobuf:"varint,7,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
	// Maximum number of messages the server will return.
	ReturnMax uint32 `protobuf:"varint,8,opt,name=return_max,json=returnMax,proto3" json:"return_max,omitempty"`
	// Maximum history window in minutes the server will return messages from.
	ReturnWindow uint32 `protobuf:"varint,9,opt,name=return_window,json=returnWindow,proto3" json:"return_window,omitempty"`
}

func (x *StoreAndForward_Statistics) Reset() {
	*x = StoreAndForward_Statistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_storeforward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreAndForward_Statistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreAndForward_Statistics) ProtoMessage() {}

func (x *StoreAndForward_Statistics) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_storeforward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreAndForward_Statistics.ProtoReflect.Descriptor instead.
func (*StoreAndForward_Statistics) Descriptor() ([]byte, []int) {
	return file_meshtastic_storeforward_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StoreAndForward_Statistics) GetMessagesTotal() uint32 {
	if x != nil {
		return x.MessagesTotal
	}
	return 0
}

func (x *StoreAndForward_Statistics) GetMessagesSaved() uint32 {
	if x != nil {
		return x.MessagesSaved
	}
	return 0
}

func (x *StoreAndForward_Statistics) GetMessagesMax() uint32 {
	if x != nil {
		return x.MessagesMax
	}
	return 0
}

func (x *StoreAndForward_Statistics) GetUpTime() uint32 {
	if x != nil {
		return x.UpTime
	}
	return 0
}

func (x *StoreAndForward_Statistics) GetRequests() uint32 {
	if x != nil {
		return x.Requests
	}
	return 0
}

func (x *StoreAndForward_Statistics) GetRequestsHistory() uint32 {
	if x != nil {
		return x.RequestsHistory
	}
	return 0
}

func (x *StoreAndForward_Statistics) GetHeartbeat() bool {
	if x != nil {
		return x.Heartbeat
	}
	return false
}

func (x *StoreAndForward_Statistics) GetReturnMax() uint32 {
	if x != nil {
		return x.ReturnMax
	}
	return 0
}

func (x *StoreAndForward_Statistics) GetReturnWindow() uint32 {
	if x != nil {
		return x.ReturnWindow
	}
	return 0
}

// TODO: REPLACE
type StoreAndForward_History struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of that will be sent to the client
	HistoryMessages uint32 `protobuf:"varint,1,opt,name=history_messages,json=historyMessages,proto3" json:"history_messages,omitempty"`
	// The window of messages that was used to filter the history client requested
	Window uint32 `protobuf:"varint,2,opt,name=window,proto3" json:"window,omitempty"`
	// Index in the packet history of the last message sent in a previous request to the server.
	// Will be sent to the client before sending the history and can be set in a subsequent request to avoid getting packets the server already sent to the client.
	LastRequest uint32 `protobuf:"varint,3,opt,name=last_request,json=lastRequest,proto3" json:"last_request,omitempty"`
}

func (x *StoreAndForward_History) Reset() {
	*x = StoreAndForward_History{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_storeforward_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreAndForward_History) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreAndForward_History) ProtoMessage() {}

func (x *StoreAndForward_History) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_storeforward_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreAndForward_History.ProtoReflect.Descriptor instead.
func (*StoreAndForward_History) Descriptor() ([]byte, []int) {
	return file_meshtastic_storeforward_proto_rawDescGZIP(), []int{0, 1}
}

func (x *StoreAndForward_History) GetHistoryMessages() uint32 {
	if x != nil {
		return x.HistoryMessages
	}
	return 0
}

func (x *StoreAndForward_History) GetWindow() uint32 {
	if x != nil {
		return x.Window
	}
	return 0
}

func (x *StoreAndForward_History) GetLastRequest() uint32 {
	if x != nil {
		return x.LastRequest
	}
	return 0
}

// TODO: REPLACE
type StoreAndForward_Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Period in seconds that the heartbeat is sent out that will be sent to the client
	Period uint32 `protobuf:"varint,1,opt,name=period,proto3" json:"period,omitempty"`
	// If set, this is not the primary Store & Forward router on the mesh
	Secondary uint32 `protobuf:"varint,2,opt,name=secondary,proto3" json:"secondary,omitempty"`
}

func (x *StoreAndForward_Heartbeat) Reset() {
	*x = StoreAndForward_Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_storeforward_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreAndForward_Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreAndForward_Heartbeat) ProtoMessage() {}

func (x *StoreAndForward_Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_storeforward_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreAndForward_Heartbeat.ProtoReflect.Descriptor instead.
func (*StoreAndForward_Heartbeat) Descriptor() ([]byte, []int) {
	return file_meshtastic_storeforward_proto_rawDescGZIP(), []int{0, 2}
}

func (x *StoreAndForward_Heartbeat) GetPeriod() uint32 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *StoreAndForward_Heartbeat) GetSecondary() uint32 {
	if x != nil {
		return x.Secondary
	}
	return 0
}

var File_meshtastic_storeforward_proto protoreflect.FileDescriptor

var file_meshtastic_storeforward_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x22, 0xec, 0x08, 0x0a, 0x0f,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12,
	0x3b, 0x0a, 0x02, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e,
	0x64, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x02, 0x72, 0x72, 0x12, 0x3e, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e,
	0x64, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x07,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x41, 0x6e, 0x64, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x48, 0x00, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x45, 0x0a,
	0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x48, 0x00, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x1a, 0xbf, 0x02, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x61, 0x76,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x4d, 0x61, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x70,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x70, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12,
	0x29, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x4d, 0x61, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x1a, 0x6f, 0x0a, 0x07,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x0a,
	0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x22, 0xbc, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x48, 0x45, 0x41, 0x52,
	0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x55, 0x54, 0x45,
	0x52, 0x5f, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x55, 0x54,
	0x45, 0x52, 0x5f, 0x50, 0x4f, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x55,
	0x54, 0x45, 0x52, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f,
	0x55, 0x54, 0x45, 0x52, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x06, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x10, 0x07,
	0x12, 0x16, 0x0a, 0x12, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x5f,
	0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x4f, 0x55, 0x54,
	0x45, 0x52, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53,
	0x54, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x40, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f,
	0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x41, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x10, 0x42, 0x12, 0x0f, 0x0a, 0x0b, 0x43,
	0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x43, 0x12, 0x0f, 0x0a, 0x0b,
	0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x4f, 0x4e, 0x47, 0x10, 0x44, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x10, 0x6a, 0x42,
	0x09, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x42, 0x6a, 0x0a, 0x13, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x73, 0x76, 0x69, 0x6c, 0x6c, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x42, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f,
	0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0xaa, 0x02, 0x14, 0x4d,
	0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x73, 0xba, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meshtastic_storeforward_proto_rawDescOnce sync.Once
	file_meshtastic_storeforward_proto_rawDescData = file_meshtastic_storeforward_proto_rawDesc
)

func file_meshtastic_storeforward_proto_rawDescGZIP() []byte {
	file_meshtastic_storeforward_proto_rawDescOnce.Do(func() {
		file_meshtastic_storeforward_proto_rawDescData = protoimpl.X.CompressGZIP(file_meshtastic_storeforward_proto_rawDescData)
	})
	return file_meshtastic_storeforward_proto_rawDescData
}

var file_meshtastic_storeforward_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meshtastic_storeforward_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_meshtastic_storeforward_proto_goTypes = []interface{}{
	(StoreAndForward_RequestResponse)(0), // 0: meshtastic.StoreAndForward.RequestResponse
	(*StoreAndForward)(nil),              // 1: meshtastic.StoreAndForward
	(*StoreAndForward_Statistics)(nil),   // 2: meshtastic.StoreAndForward.Statistics
	(*StoreAndForward_History)(nil),      // 3: meshtastic.StoreAndForward.History
	(*StoreAndForward_Heartbeat)(nil),    // 4: meshtastic.StoreAndForward.Heartbeat
}
var file_meshtastic_storeforward_proto_depIdxs = []int32{
	0, // 0: meshtastic.StoreAndForward.rr:type_name -> meshtastic.StoreAndForward.RequestResponse
	2, // 1: meshtastic.StoreAndForward.stats:type_name -> meshtastic.StoreAndForward.Statistics
	3, // 2: meshtastic.StoreAndForward.history:type_name -> meshtastic.StoreAndForward.History
	4, // 3: meshtastic.StoreAndForward.heartbeat:type_name -> meshtastic.StoreAndForward.Heartbeat
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_meshtastic_storeforward_proto_init() }
func file_meshtastic_storeforward_proto_init() {
	if File_meshtastic_storeforward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meshtastic_storeforward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreAndForward); i {
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
		file_meshtastic_storeforward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreAndForward_Statistics); i {
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
		file_meshtastic_storeforward_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreAndForward_History); i {
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
		file_meshtastic_storeforward_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreAndForward_Heartbeat); i {
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
	file_meshtastic_storeforward_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StoreAndForward_Stats)(nil),
		(*StoreAndForward_History_)(nil),
		(*StoreAndForward_Heartbeat_)(nil),
		(*StoreAndForward_Text)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meshtastic_storeforward_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meshtastic_storeforward_proto_goTypes,
		DependencyIndexes: file_meshtastic_storeforward_proto_depIdxs,
		EnumInfos:         file_meshtastic_storeforward_proto_enumTypes,
		MessageInfos:      file_meshtastic_storeforward_proto_msgTypes,
	}.Build()
	File_meshtastic_storeforward_proto = out.File
	file_meshtastic_storeforward_proto_rawDesc = nil
	file_meshtastic_storeforward_proto_goTypes = nil
	file_meshtastic_storeforward_proto_depIdxs = nil
}
