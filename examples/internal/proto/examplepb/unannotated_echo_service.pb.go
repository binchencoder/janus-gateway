// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: examples/internal/proto/examplepb/unannotated_echo_service.proto

// Unannotated Echo Service
// Similar to echo_service.proto but without annotations. See
// unannotated_echo_service.yaml for the equivalent of the annotations in
// gRPC API configuration format.
//
// Echo Service API consists of a single service which returns
// a message.

package examplepb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Embedded represents a message embedded in SimpleMessage.
type UnannotatedEmbedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Mark:
	//	*UnannotatedEmbedded_Progress
	//	*UnannotatedEmbedded_Note
	Mark isUnannotatedEmbedded_Mark `protobuf_oneof:"mark"`
}

func (x *UnannotatedEmbedded) Reset() {
	*x = UnannotatedEmbedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_unannotated_echo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnannotatedEmbedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnannotatedEmbedded) ProtoMessage() {}

func (x *UnannotatedEmbedded) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_unannotated_echo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnannotatedEmbedded.ProtoReflect.Descriptor instead.
func (*UnannotatedEmbedded) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDescGZIP(), []int{0}
}

func (m *UnannotatedEmbedded) GetMark() isUnannotatedEmbedded_Mark {
	if m != nil {
		return m.Mark
	}
	return nil
}

func (x *UnannotatedEmbedded) GetProgress() int64 {
	if x, ok := x.GetMark().(*UnannotatedEmbedded_Progress); ok {
		return x.Progress
	}
	return 0
}

func (x *UnannotatedEmbedded) GetNote() string {
	if x, ok := x.GetMark().(*UnannotatedEmbedded_Note); ok {
		return x.Note
	}
	return ""
}

type isUnannotatedEmbedded_Mark interface {
	isUnannotatedEmbedded_Mark()
}

type UnannotatedEmbedded_Progress struct {
	Progress int64 `protobuf:"varint,1,opt,name=progress,proto3,oneof"`
}

type UnannotatedEmbedded_Note struct {
	Note string `protobuf:"bytes,2,opt,name=note,proto3,oneof"`
}

func (*UnannotatedEmbedded_Progress) isUnannotatedEmbedded_Mark() {}

func (*UnannotatedEmbedded_Note) isUnannotatedEmbedded_Mark() {}

// UnannotatedSimpleMessage represents a simple message sent to the unannotated Echo service.
type UnannotatedSimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represents the message identifier.
	Id       string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num      int64              `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Duration *duration.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	// Types that are assignable to Code:
	//	*UnannotatedSimpleMessage_LineNum
	//	*UnannotatedSimpleMessage_Lang
	Code   isUnannotatedSimpleMessage_Code `protobuf_oneof:"code"`
	Status *UnannotatedEmbedded            `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are assignable to Ext:
	//	*UnannotatedSimpleMessage_En
	//	*UnannotatedSimpleMessage_No
	Ext isUnannotatedSimpleMessage_Ext `protobuf_oneof:"ext"`
}

func (x *UnannotatedSimpleMessage) Reset() {
	*x = UnannotatedSimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_unannotated_echo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnannotatedSimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnannotatedSimpleMessage) ProtoMessage() {}

func (x *UnannotatedSimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_unannotated_echo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnannotatedSimpleMessage.ProtoReflect.Descriptor instead.
func (*UnannotatedSimpleMessage) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDescGZIP(), []int{1}
}

func (x *UnannotatedSimpleMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UnannotatedSimpleMessage) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *UnannotatedSimpleMessage) GetDuration() *duration.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (m *UnannotatedSimpleMessage) GetCode() isUnannotatedSimpleMessage_Code {
	if m != nil {
		return m.Code
	}
	return nil
}

func (x *UnannotatedSimpleMessage) GetLineNum() int64 {
	if x, ok := x.GetCode().(*UnannotatedSimpleMessage_LineNum); ok {
		return x.LineNum
	}
	return 0
}

func (x *UnannotatedSimpleMessage) GetLang() string {
	if x, ok := x.GetCode().(*UnannotatedSimpleMessage_Lang); ok {
		return x.Lang
	}
	return ""
}

func (x *UnannotatedSimpleMessage) GetStatus() *UnannotatedEmbedded {
	if x != nil {
		return x.Status
	}
	return nil
}

func (m *UnannotatedSimpleMessage) GetExt() isUnannotatedSimpleMessage_Ext {
	if m != nil {
		return m.Ext
	}
	return nil
}

func (x *UnannotatedSimpleMessage) GetEn() int64 {
	if x, ok := x.GetExt().(*UnannotatedSimpleMessage_En); ok {
		return x.En
	}
	return 0
}

func (x *UnannotatedSimpleMessage) GetNo() *UnannotatedEmbedded {
	if x, ok := x.GetExt().(*UnannotatedSimpleMessage_No); ok {
		return x.No
	}
	return nil
}

type isUnannotatedSimpleMessage_Code interface {
	isUnannotatedSimpleMessage_Code()
}

type UnannotatedSimpleMessage_LineNum struct {
	LineNum int64 `protobuf:"varint,4,opt,name=line_num,json=lineNum,proto3,oneof"`
}

type UnannotatedSimpleMessage_Lang struct {
	Lang string `protobuf:"bytes,5,opt,name=lang,proto3,oneof"`
}

func (*UnannotatedSimpleMessage_LineNum) isUnannotatedSimpleMessage_Code() {}

func (*UnannotatedSimpleMessage_Lang) isUnannotatedSimpleMessage_Code() {}

type isUnannotatedSimpleMessage_Ext interface {
	isUnannotatedSimpleMessage_Ext()
}

type UnannotatedSimpleMessage_En struct {
	En int64 `protobuf:"varint,7,opt,name=en,proto3,oneof"`
}

type UnannotatedSimpleMessage_No struct {
	No *UnannotatedEmbedded `protobuf:"bytes,8,opt,name=no,proto3,oneof"`
}

func (*UnannotatedSimpleMessage_En) isUnannotatedSimpleMessage_Ext() {}

func (*UnannotatedSimpleMessage_No) isUnannotatedSimpleMessage_Ext() {}

var File_examples_internal_proto_examplepb_unannotated_echo_service_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x65, 0x63, 0x68, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x51, 0x0a, 0x13, 0x55, 0x6e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xfb, 0x02, 0x0a, 0x18, 0x55, 0x6e, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6e, 0x75, 0x6d, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x08, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x07, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x5b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e,
	0x55, 0x6e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x02, 0x65,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x02, 0x65, 0x6e, 0x12, 0x55, 0x0a,
	0x02, 0x6e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x48, 0x01,
	0x52, 0x02, 0x6e, 0x6f, 0x42, 0x06, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x05, 0x0a, 0x03,
	0x65, 0x78, 0x74, 0x32, 0xf9, 0x03, 0x0a, 0x16, 0x55, 0x6e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x65, 0x64, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9a,
	0x01, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x48, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x48, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x2e, 0x55, 0x6e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x08,
	0x45, 0x63, 0x68, 0x6f, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x48, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x48, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0xa0, 0x01, 0x0a,
	0x0a, 0x45, 0x63, 0x68, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x48, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x48, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69,
	0x6e, 0x63, 0x68, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2f, 0x65, 0x61, 0x73, 0x65, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDescData = file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDesc
)

func file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDescData
}

var file_examples_internal_proto_examplepb_unannotated_echo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_examples_internal_proto_examplepb_unannotated_echo_service_proto_goTypes = []interface{}{
	(*UnannotatedEmbedded)(nil),      // 0: grpc.gateway.examples.internal.proto.examplepb.UnannotatedEmbedded
	(*UnannotatedSimpleMessage)(nil), // 1: grpc.gateway.examples.internal.proto.examplepb.UnannotatedSimpleMessage
	(*duration.Duration)(nil),        // 2: google.protobuf.Duration
}
var file_examples_internal_proto_examplepb_unannotated_echo_service_proto_depIdxs = []int32{
	2, // 0: grpc.gateway.examples.internal.proto.examplepb.UnannotatedSimpleMessage.duration:type_name -> google.protobuf.Duration
	0, // 1: grpc.gateway.examples.internal.proto.examplepb.UnannotatedSimpleMessage.status:type_name -> grpc.gateway.examples.internal.proto.examplepb.UnannotatedEmbedded
	0, // 2: grpc.gateway.examples.internal.proto.examplepb.UnannotatedSimpleMessage.no:type_name -> grpc.gateway.examples.internal.proto.examplepb.UnannotatedEmbedded
	1, // 3: grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService.Echo:input_type -> grpc.gateway.examples.internal.proto.examplepb.UnannotatedSimpleMessage
	1, // 4: grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService.EchoBody:input_type -> grpc.gateway.examples.internal.proto.examplepb.UnannotatedSimpleMessage
	1, // 5: grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService.EchoDelete:input_type -> grpc.gateway.examples.internal.proto.examplepb.UnannotatedSimpleMessage
	1, // 6: grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService.Echo:output_type -> grpc.gateway.examples.internal.proto.examplepb.UnannotatedSimpleMessage
	1, // 7: grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService.EchoBody:output_type -> grpc.gateway.examples.internal.proto.examplepb.UnannotatedSimpleMessage
	1, // 8: grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService.EchoDelete:output_type -> grpc.gateway.examples.internal.proto.examplepb.UnannotatedSimpleMessage
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_unannotated_echo_service_proto_init() }
func file_examples_internal_proto_examplepb_unannotated_echo_service_proto_init() {
	if File_examples_internal_proto_examplepb_unannotated_echo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_internal_proto_examplepb_unannotated_echo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnannotatedEmbedded); i {
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
		file_examples_internal_proto_examplepb_unannotated_echo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnannotatedSimpleMessage); i {
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
	file_examples_internal_proto_examplepb_unannotated_echo_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UnannotatedEmbedded_Progress)(nil),
		(*UnannotatedEmbedded_Note)(nil),
	}
	file_examples_internal_proto_examplepb_unannotated_echo_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UnannotatedSimpleMessage_LineNum)(nil),
		(*UnannotatedSimpleMessage_Lang)(nil),
		(*UnannotatedSimpleMessage_En)(nil),
		(*UnannotatedSimpleMessage_No)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_examples_internal_proto_examplepb_unannotated_echo_service_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_unannotated_echo_service_proto_depIdxs,
		MessageInfos:      file_examples_internal_proto_examplepb_unannotated_echo_service_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_unannotated_echo_service_proto = out.File
	file_examples_internal_proto_examplepb_unannotated_echo_service_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_unannotated_echo_service_proto_goTypes = nil
	file_examples_internal_proto_examplepb_unannotated_echo_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UnannotatedEchoServiceClient is the client API for UnannotatedEchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UnannotatedEchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error)
}

type unannotatedEchoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUnannotatedEchoServiceClient(cc grpc.ClientConnInterface) UnannotatedEchoServiceClient {
	return &unannotatedEchoServiceClient{cc}
}

func (c *unannotatedEchoServiceClient) Echo(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error) {
	out := new(UnannotatedSimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unannotatedEchoServiceClient) EchoBody(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error) {
	out := new(UnannotatedSimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService/EchoBody", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unannotatedEchoServiceClient) EchoDelete(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error) {
	out := new(UnannotatedSimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService/EchoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnannotatedEchoServiceServer is the server API for UnannotatedEchoService service.
type UnannotatedEchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error)
}

// UnimplementedUnannotatedEchoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUnannotatedEchoServiceServer struct {
}

func (*UnimplementedUnannotatedEchoServiceServer) Echo(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedUnannotatedEchoServiceServer) EchoBody(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoBody not implemented")
}
func (*UnimplementedUnannotatedEchoServiceServer) EchoDelete(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoDelete not implemented")
}

func RegisterUnannotatedEchoServiceServer(s *grpc.Server, srv UnannotatedEchoServiceServer) {
	s.RegisterService(&_UnannotatedEchoService_serviceDesc, srv)
}

func _UnannotatedEchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnannotatedSimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnannotatedEchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnannotatedEchoServiceServer).Echo(ctx, req.(*UnannotatedSimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnannotatedEchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnannotatedSimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnannotatedEchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService/EchoBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnannotatedEchoServiceServer).EchoBody(ctx, req.(*UnannotatedSimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnannotatedEchoService_EchoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnannotatedSimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnannotatedEchoServiceServer).EchoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService/EchoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnannotatedEchoServiceServer).EchoDelete(ctx, req.(*UnannotatedSimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _UnannotatedEchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService",
	HandlerType: (*UnannotatedEchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _UnannotatedEchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _UnannotatedEchoService_EchoBody_Handler,
		},
		{
			MethodName: "EchoDelete",
			Handler:    _UnannotatedEchoService_EchoDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/internal/proto/examplepb/unannotated_echo_service.proto",
}