// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: rtc.proto

package livekit

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SignalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*SignalRequest_Offer
	//	*SignalRequest_Negotiate
	//	*SignalRequest_Trickle
	//	*SignalRequest_Control
	Message isSignalRequest_Message `protobuf_oneof:"message"`
}

func (x *SignalRequest) Reset() {
	*x = SignalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalRequest) ProtoMessage() {}

func (x *SignalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalRequest.ProtoReflect.Descriptor instead.
func (*SignalRequest) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{0}
}

func (m *SignalRequest) GetMessage() isSignalRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *SignalRequest) GetOffer() *SessionDescription {
	if x, ok := x.GetMessage().(*SignalRequest_Offer); ok {
		return x.Offer
	}
	return nil
}

func (x *SignalRequest) GetNegotiate() *SessionDescription {
	if x, ok := x.GetMessage().(*SignalRequest_Negotiate); ok {
		return x.Negotiate
	}
	return nil
}

func (x *SignalRequest) GetTrickle() *Trickle {
	if x, ok := x.GetMessage().(*SignalRequest_Trickle); ok {
		return x.Trickle
	}
	return nil
}

func (x *SignalRequest) GetControl() *MediaControl {
	if x, ok := x.GetMessage().(*SignalRequest_Control); ok {
		return x.Control
	}
	return nil
}

type isSignalRequest_Message interface {
	isSignalRequest_Message()
}

type SignalRequest_Offer struct {
	Offer *SessionDescription `protobuf:"bytes,1,opt,name=offer,proto3,oneof"`
}

type SignalRequest_Negotiate struct {
	Negotiate *SessionDescription `protobuf:"bytes,2,opt,name=negotiate,proto3,oneof"`
}

type SignalRequest_Trickle struct {
	Trickle *Trickle `protobuf:"bytes,3,opt,name=trickle,proto3,oneof"`
}

type SignalRequest_Control struct {
	Control *MediaControl `protobuf:"bytes,4,opt,name=control,proto3,oneof"`
}

func (*SignalRequest_Offer) isSignalRequest_Message() {}

func (*SignalRequest_Negotiate) isSignalRequest_Message() {}

func (*SignalRequest_Trickle) isSignalRequest_Message() {}

func (*SignalRequest_Control) isSignalRequest_Message() {}

type SignalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*SignalResponse_Join
	//	*SignalResponse_Answer
	//	*SignalResponse_Negotiate
	//	*SignalResponse_Trickle
	//	*SignalResponse_Update
	//	*SignalResponse_TrackPublished
	Message isSignalResponse_Message `protobuf_oneof:"message"`
}

func (x *SignalResponse) Reset() {
	*x = SignalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalResponse) ProtoMessage() {}

func (x *SignalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalResponse.ProtoReflect.Descriptor instead.
func (*SignalResponse) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{1}
}

func (m *SignalResponse) GetMessage() isSignalResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *SignalResponse) GetJoin() *JoinResponse {
	if x, ok := x.GetMessage().(*SignalResponse_Join); ok {
		return x.Join
	}
	return nil
}

func (x *SignalResponse) GetAnswer() *SessionDescription {
	if x, ok := x.GetMessage().(*SignalResponse_Answer); ok {
		return x.Answer
	}
	return nil
}

func (x *SignalResponse) GetNegotiate() *SessionDescription {
	if x, ok := x.GetMessage().(*SignalResponse_Negotiate); ok {
		return x.Negotiate
	}
	return nil
}

func (x *SignalResponse) GetTrickle() *Trickle {
	if x, ok := x.GetMessage().(*SignalResponse_Trickle); ok {
		return x.Trickle
	}
	return nil
}

func (x *SignalResponse) GetUpdate() *ParticipantUpdate {
	if x, ok := x.GetMessage().(*SignalResponse_Update); ok {
		return x.Update
	}
	return nil
}

func (x *SignalResponse) GetTrackPublished() *TrackInfo {
	if x, ok := x.GetMessage().(*SignalResponse_TrackPublished); ok {
		return x.TrackPublished
	}
	return nil
}

type isSignalResponse_Message interface {
	isSignalResponse_Message()
}

type SignalResponse_Join struct {
	// sent when join is accepted
	Join *JoinResponse `protobuf:"bytes,1,opt,name=join,proto3,oneof"`
}

type SignalResponse_Answer struct {
	// sent when offer is answered
	Answer *SessionDescription `protobuf:"bytes,2,opt,name=answer,proto3,oneof"`
}

type SignalResponse_Negotiate struct {
	// sent when a negotiated sd is available (could be either offer or answer)
	Negotiate *SessionDescription `protobuf:"bytes,3,opt,name=negotiate,proto3,oneof"`
}

type SignalResponse_Trickle struct {
	// sent when an ICE candidate is available
	Trickle *Trickle `protobuf:"bytes,4,opt,name=trickle,proto3,oneof"`
}

type SignalResponse_Update struct {
	// sent when participants in the room has changed
	Update *ParticipantUpdate `protobuf:"bytes,5,opt,name=update,proto3,oneof"`
}

type SignalResponse_TrackPublished struct {
	// sent to the participant when their track has been published
	TrackPublished *TrackInfo `protobuf:"bytes,6,opt,name=trackPublished,proto3,oneof"`
}

func (*SignalResponse_Join) isSignalResponse_Message() {}

func (*SignalResponse_Answer) isSignalResponse_Message() {}

func (*SignalResponse_Negotiate) isSignalResponse_Message() {}

func (*SignalResponse_Trickle) isSignalResponse_Message() {}

func (*SignalResponse_Update) isSignalResponse_Message() {}

func (*SignalResponse_TrackPublished) isSignalResponse_Message() {}

type Trickle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CandidateInit string `protobuf:"bytes,1,opt,name=candidateInit,proto3" json:"candidateInit,omitempty"`
}

func (x *Trickle) Reset() {
	*x = Trickle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trickle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trickle) ProtoMessage() {}

func (x *Trickle) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trickle.ProtoReflect.Descriptor instead.
func (*Trickle) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{2}
}

func (x *Trickle) GetCandidateInit() string {
	if x != nil {
		return x.CandidateInit
	}
	return ""
}

type SessionDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"` // "answer" | "offer" | "pranswer" | "rollback"
	Sdp  string `protobuf:"bytes,2,opt,name=sdp,proto3" json:"sdp,omitempty"`
}

func (x *SessionDescription) Reset() {
	*x = SessionDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionDescription) ProtoMessage() {}

func (x *SessionDescription) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionDescription.ProtoReflect.Descriptor instead.
func (*SessionDescription) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{3}
}

func (x *SessionDescription) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SessionDescription) GetSdp() string {
	if x != nil {
		return x.Sdp
	}
	return ""
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room              *RoomInfo          `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	Participant       *ParticipantInfo   `protobuf:"bytes,2,opt,name=participant,proto3" json:"participant,omitempty"`
	OtherParticipants []*ParticipantInfo `protobuf:"bytes,3,rep,name=other_participants,json=otherParticipants,proto3" json:"other_participants,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{4}
}

func (x *JoinResponse) GetRoom() *RoomInfo {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *JoinResponse) GetParticipant() *ParticipantInfo {
	if x != nil {
		return x.Participant
	}
	return nil
}

func (x *JoinResponse) GetOtherParticipants() []*ParticipantInfo {
	if x != nil {
		return x.OtherParticipants
	}
	return nil
}

type MediaControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MediaControl) Reset() {
	*x = MediaControl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaControl) ProtoMessage() {}

func (x *MediaControl) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaControl.ProtoReflect.Descriptor instead.
func (*MediaControl) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{5}
}

type ParticipantUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Participants []*ParticipantInfo `protobuf:"bytes,1,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (x *ParticipantUpdate) Reset() {
	*x = ParticipantUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipantUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantUpdate) ProtoMessage() {}

func (x *ParticipantUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantUpdate.ProtoReflect.Descriptor instead.
func (*ParticipantUpdate) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{6}
}

func (x *ParticipantUpdate) GetParticipants() []*ParticipantInfo {
	if x != nil {
		return x.Participants
	}
	return nil
}

var File_rtc_proto protoreflect.FileDescriptor

var file_rtc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x74, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xed, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x09, 0x6e, 0x65, 0x67, 0x6f,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x65, 0x67, 0x6f,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x63, 0x6b, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x54, 0x72, 0x69, 0x63, 0x6b, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x07, 0x74, 0x72, 0x69, 0x63,
	0x6b, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xde, 0x02, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x69,
	0x6e, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x09, 0x6e, 0x65, 0x67, 0x6f,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x65, 0x67, 0x6f,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x63, 0x6b, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x54, 0x72, 0x69, 0x63, 0x6b, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x07, 0x74, 0x72, 0x69, 0x63,
	0x6b, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2f, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x63, 0x6b, 0x6c, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x69, 0x74, 0x22, 0x3a, 0x0a, 0x12, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x64, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x64, 0x70, 0x22,
	0xba, 0x01, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x3a, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x12, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x0e, 0x0a, 0x0c,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x22, 0x51, 0x0a, 0x11,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rtc_proto_rawDescOnce sync.Once
	file_rtc_proto_rawDescData = file_rtc_proto_rawDesc
)

func file_rtc_proto_rawDescGZIP() []byte {
	file_rtc_proto_rawDescOnce.Do(func() {
		file_rtc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rtc_proto_rawDescData)
	})
	return file_rtc_proto_rawDescData
}

var file_rtc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rtc_proto_goTypes = []interface{}{
	(*SignalRequest)(nil),      // 0: livekit.SignalRequest
	(*SignalResponse)(nil),     // 1: livekit.SignalResponse
	(*Trickle)(nil),            // 2: livekit.Trickle
	(*SessionDescription)(nil), // 3: livekit.SessionDescription
	(*JoinResponse)(nil),       // 4: livekit.JoinResponse
	(*MediaControl)(nil),       // 5: livekit.MediaControl
	(*ParticipantUpdate)(nil),  // 6: livekit.ParticipantUpdate
	(*TrackInfo)(nil),          // 7: livekit.TrackInfo
	(*RoomInfo)(nil),           // 8: livekit.RoomInfo
	(*ParticipantInfo)(nil),    // 9: livekit.ParticipantInfo
}
var file_rtc_proto_depIdxs = []int32{
	3,  // 0: livekit.SignalRequest.offer:type_name -> livekit.SessionDescription
	3,  // 1: livekit.SignalRequest.negotiate:type_name -> livekit.SessionDescription
	2,  // 2: livekit.SignalRequest.trickle:type_name -> livekit.Trickle
	5,  // 3: livekit.SignalRequest.control:type_name -> livekit.MediaControl
	4,  // 4: livekit.SignalResponse.join:type_name -> livekit.JoinResponse
	3,  // 5: livekit.SignalResponse.answer:type_name -> livekit.SessionDescription
	3,  // 6: livekit.SignalResponse.negotiate:type_name -> livekit.SessionDescription
	2,  // 7: livekit.SignalResponse.trickle:type_name -> livekit.Trickle
	6,  // 8: livekit.SignalResponse.update:type_name -> livekit.ParticipantUpdate
	7,  // 9: livekit.SignalResponse.trackPublished:type_name -> livekit.TrackInfo
	8,  // 10: livekit.JoinResponse.room:type_name -> livekit.RoomInfo
	9,  // 11: livekit.JoinResponse.participant:type_name -> livekit.ParticipantInfo
	9,  // 12: livekit.JoinResponse.other_participants:type_name -> livekit.ParticipantInfo
	9,  // 13: livekit.ParticipantUpdate.participants:type_name -> livekit.ParticipantInfo
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_rtc_proto_init() }
func file_rtc_proto_init() {
	if File_rtc_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rtc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalRequest); i {
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
		file_rtc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalResponse); i {
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
		file_rtc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trickle); i {
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
		file_rtc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionDescription); i {
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
		file_rtc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_rtc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaControl); i {
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
		file_rtc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipantUpdate); i {
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
	file_rtc_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SignalRequest_Offer)(nil),
		(*SignalRequest_Negotiate)(nil),
		(*SignalRequest_Trickle)(nil),
		(*SignalRequest_Control)(nil),
	}
	file_rtc_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SignalResponse_Join)(nil),
		(*SignalResponse_Answer)(nil),
		(*SignalResponse_Negotiate)(nil),
		(*SignalResponse_Trickle)(nil),
		(*SignalResponse_Update)(nil),
		(*SignalResponse_TrackPublished)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rtc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rtc_proto_goTypes,
		DependencyIndexes: file_rtc_proto_depIdxs,
		MessageInfos:      file_rtc_proto_msgTypes,
	}.Build()
	File_rtc_proto = out.File
	file_rtc_proto_rawDesc = nil
	file_rtc_proto_goTypes = nil
	file_rtc_proto_depIdxs = nil
}
