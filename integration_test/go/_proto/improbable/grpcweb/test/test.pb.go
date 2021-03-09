// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/grpcweb/test/test.proto

package test

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

type PingRequest_FailureType int32

const (
	PingRequest_NONE         PingRequest_FailureType = 0
	PingRequest_CODE         PingRequest_FailureType = 1
	PingRequest_CODE_UNICODE PingRequest_FailureType = 3
)

var PingRequest_FailureType_name = map[int32]string{
	0: "NONE",
	1: "CODE",
	3: "CODE_UNICODE",
}

var PingRequest_FailureType_value = map[string]int32{
	"NONE":         0,
	"CODE":         1,
	"CODE_UNICODE": 3,
}

func (x PingRequest_FailureType) String() string {
	return proto.EnumName(PingRequest_FailureType_name, int32(x))
}

func (PingRequest_FailureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_18e8d9116297b46c, []int{0, 0}
}

type PingRequest struct {
	Value                string                  `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ResponseCount        int32                   `protobuf:"varint,2,opt,name=response_count,json=responseCount,proto3" json:"response_count,omitempty"`
	ErrorCodeReturned    uint32                  `protobuf:"varint,3,opt,name=error_code_returned,json=errorCodeReturned,proto3" json:"error_code_returned,omitempty"`
	FailureType          PingRequest_FailureType `protobuf:"varint,4,opt,name=failure_type,json=failureType,proto3,enum=improbable.grpcweb.test.PingRequest_FailureType" json:"failure_type,omitempty"`
	CheckMetadata        bool                    `protobuf:"varint,5,opt,name=check_metadata,json=checkMetadata,proto3" json:"check_metadata,omitempty"`
	SendHeaders          bool                    `protobuf:"varint,6,opt,name=send_headers,json=sendHeaders,proto3" json:"send_headers,omitempty"`
	SendTrailers         bool                    `protobuf:"varint,7,opt,name=send_trailers,json=sendTrailers,proto3" json:"send_trailers,omitempty"`
	StreamIdentifier     string                  `protobuf:"bytes,8,opt,name=stream_identifier,json=streamIdentifier,proto3" json:"stream_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e8d9116297b46c, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingRequest) GetResponseCount() int32 {
	if m != nil {
		return m.ResponseCount
	}
	return 0
}

func (m *PingRequest) GetErrorCodeReturned() uint32 {
	if m != nil {
		return m.ErrorCodeReturned
	}
	return 0
}

func (m *PingRequest) GetFailureType() PingRequest_FailureType {
	if m != nil {
		return m.FailureType
	}
	return PingRequest_NONE
}

func (m *PingRequest) GetCheckMetadata() bool {
	if m != nil {
		return m.CheckMetadata
	}
	return false
}

func (m *PingRequest) GetSendHeaders() bool {
	if m != nil {
		return m.SendHeaders
	}
	return false
}

func (m *PingRequest) GetSendTrailers() bool {
	if m != nil {
		return m.SendTrailers
	}
	return false
}

func (m *PingRequest) GetStreamIdentifier() string {
	if m != nil {
		return m.StreamIdentifier
	}
	return ""
}

type PingResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Counter              int32    `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e8d9116297b46c, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingResponse) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

type TextMessage struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	SendHeaders          bool     `protobuf:"varint,2,opt,name=send_headers,json=sendHeaders,proto3" json:"send_headers,omitempty"`
	SendTrailers         bool     `protobuf:"varint,3,opt,name=send_trailers,json=sendTrailers,proto3" json:"send_trailers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextMessage) Reset()         { *m = TextMessage{} }
func (m *TextMessage) String() string { return proto.CompactTextString(m) }
func (*TextMessage) ProtoMessage()    {}
func (*TextMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e8d9116297b46c, []int{2}
}

func (m *TextMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextMessage.Unmarshal(m, b)
}
func (m *TextMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextMessage.Marshal(b, m, deterministic)
}
func (m *TextMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextMessage.Merge(m, src)
}
func (m *TextMessage) XXX_Size() int {
	return xxx_messageInfo_TextMessage.Size(m)
}
func (m *TextMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TextMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TextMessage proto.InternalMessageInfo

func (m *TextMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *TextMessage) GetSendHeaders() bool {
	if m != nil {
		return m.SendHeaders
	}
	return false
}

func (m *TextMessage) GetSendTrailers() bool {
	if m != nil {
		return m.SendTrailers
	}
	return false
}

type ContinueStreamRequest struct {
	StreamIdentifier     string   `protobuf:"bytes,1,opt,name=stream_identifier,json=streamIdentifier,proto3" json:"stream_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContinueStreamRequest) Reset()         { *m = ContinueStreamRequest{} }
func (m *ContinueStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ContinueStreamRequest) ProtoMessage()    {}
func (*ContinueStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e8d9116297b46c, []int{3}
}

func (m *ContinueStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContinueStreamRequest.Unmarshal(m, b)
}
func (m *ContinueStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContinueStreamRequest.Marshal(b, m, deterministic)
}
func (m *ContinueStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinueStreamRequest.Merge(m, src)
}
func (m *ContinueStreamRequest) XXX_Size() int {
	return xxx_messageInfo_ContinueStreamRequest.Size(m)
}
func (m *ContinueStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinueStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContinueStreamRequest proto.InternalMessageInfo

func (m *ContinueStreamRequest) GetStreamIdentifier() string {
	if m != nil {
		return m.StreamIdentifier
	}
	return ""
}

type CheckStreamClosedRequest struct {
	StreamIdentifier     string   `protobuf:"bytes,1,opt,name=stream_identifier,json=streamIdentifier,proto3" json:"stream_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckStreamClosedRequest) Reset()         { *m = CheckStreamClosedRequest{} }
func (m *CheckStreamClosedRequest) String() string { return proto.CompactTextString(m) }
func (*CheckStreamClosedRequest) ProtoMessage()    {}
func (*CheckStreamClosedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e8d9116297b46c, []int{4}
}

func (m *CheckStreamClosedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckStreamClosedRequest.Unmarshal(m, b)
}
func (m *CheckStreamClosedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckStreamClosedRequest.Marshal(b, m, deterministic)
}
func (m *CheckStreamClosedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckStreamClosedRequest.Merge(m, src)
}
func (m *CheckStreamClosedRequest) XXX_Size() int {
	return xxx_messageInfo_CheckStreamClosedRequest.Size(m)
}
func (m *CheckStreamClosedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckStreamClosedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckStreamClosedRequest proto.InternalMessageInfo

func (m *CheckStreamClosedRequest) GetStreamIdentifier() string {
	if m != nil {
		return m.StreamIdentifier
	}
	return ""
}

type CheckStreamClosedResponse struct {
	Closed               bool     `protobuf:"varint,1,opt,name=closed,proto3" json:"closed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckStreamClosedResponse) Reset()         { *m = CheckStreamClosedResponse{} }
func (m *CheckStreamClosedResponse) String() string { return proto.CompactTextString(m) }
func (*CheckStreamClosedResponse) ProtoMessage()    {}
func (*CheckStreamClosedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e8d9116297b46c, []int{5}
}

func (m *CheckStreamClosedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckStreamClosedResponse.Unmarshal(m, b)
}
func (m *CheckStreamClosedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckStreamClosedResponse.Marshal(b, m, deterministic)
}
func (m *CheckStreamClosedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckStreamClosedResponse.Merge(m, src)
}
func (m *CheckStreamClosedResponse) XXX_Size() int {
	return xxx_messageInfo_CheckStreamClosedResponse.Size(m)
}
func (m *CheckStreamClosedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckStreamClosedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckStreamClosedResponse proto.InternalMessageInfo

func (m *CheckStreamClosedResponse) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
}

func init() {
	proto.RegisterEnum("improbable.grpcweb.test.PingRequest_FailureType", PingRequest_FailureType_name, PingRequest_FailureType_value)
	proto.RegisterType((*PingRequest)(nil), "improbable.grpcweb.test.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "improbable.grpcweb.test.PingResponse")
	proto.RegisterType((*TextMessage)(nil), "improbable.grpcweb.test.TextMessage")
	proto.RegisterType((*ContinueStreamRequest)(nil), "improbable.grpcweb.test.ContinueStreamRequest")
	proto.RegisterType((*CheckStreamClosedRequest)(nil), "improbable.grpcweb.test.CheckStreamClosedRequest")
	proto.RegisterType((*CheckStreamClosedResponse)(nil), "improbable.grpcweb.test.CheckStreamClosedResponse")
}

func init() {
	proto.RegisterFile("improbable/grpcweb/test/test.proto", fileDescriptor_18e8d9116297b46c)
}

var fileDescriptor_18e8d9116297b46c = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x41, 0x4f, 0xdb, 0x4a,
	0x10, 0x8e, 0x93, 0x00, 0x61, 0x92, 0xf0, 0xc2, 0xbe, 0xf7, 0xc0, 0xa4, 0x97, 0xd4, 0x2d, 0x52,
	0xa4, 0x4a, 0x0e, 0x0d, 0xa7, 0x5e, 0x7a, 0x20, 0xa4, 0x2d, 0x2a, 0x04, 0x64, 0x42, 0x0f, 0xb4,
	0x95, 0xe5, 0xd8, 0x93, 0xb0, 0xc2, 0xf1, 0xba, 0xeb, 0x35, 0x05, 0xa9, 0xbf, 0xa7, 0x7f, 0xae,
	0xd7, 0xfe, 0x80, 0x6a, 0xd7, 0x76, 0x49, 0x05, 0xa6, 0x69, 0x95, 0x4b, 0xb4, 0xfb, 0xed, 0x37,
	0xb3, 0xb3, 0xdf, 0xcc, 0x17, 0x83, 0x41, 0xa7, 0x21, 0x67, 0x23, 0x67, 0xe4, 0x63, 0x67, 0xc2,
	0x43, 0xf7, 0x33, 0x8e, 0x3a, 0x02, 0x23, 0xa1, 0x7e, 0xcc, 0x90, 0x33, 0xc1, 0xc8, 0xe6, 0x2d,
	0xc7, 0x4c, 0x39, 0xa6, 0x3c, 0x6e, 0x3e, 0x9a, 0x30, 0x36, 0xf1, 0xb1, 0xa3, 0x68, 0xa3, 0x78,
	0xdc, 0xc1, 0x69, 0x28, 0x6e, 0x92, 0x28, 0xe3, 0x6b, 0x09, 0xaa, 0x27, 0x34, 0x98, 0x58, 0xf8,
	0x29, 0xc6, 0x48, 0x90, 0xff, 0x60, 0xe9, 0xca, 0xf1, 0x63, 0xd4, 0xb5, 0x96, 0xd6, 0x5e, 0xb5,
	0x92, 0x0d, 0xd9, 0x86, 0x35, 0x8e, 0x51, 0xc8, 0x82, 0x08, 0x6d, 0x97, 0xc5, 0x81, 0xd0, 0x8b,
	0x2d, 0xad, 0xbd, 0x64, 0xd5, 0x33, 0xb4, 0x27, 0x41, 0x62, 0xc2, 0xbf, 0xc8, 0x39, 0xe3, 0xb6,
	0xcb, 0x3c, 0xb4, 0x39, 0x8a, 0x98, 0x07, 0xe8, 0xe9, 0xa5, 0x96, 0xd6, 0xae, 0x5b, 0xeb, 0xea,
	0xa8, 0xc7, 0x3c, 0xb4, 0xd2, 0x03, 0x72, 0x0a, 0xb5, 0xb1, 0x43, 0xfd, 0x98, 0xa3, 0x2d, 0x6e,
	0x42, 0xd4, 0xcb, 0x2d, 0xad, 0xbd, 0xd6, 0xdd, 0x31, 0x73, 0x5e, 0x62, 0xce, 0x14, 0x6a, 0xbe,
	0x4a, 0x02, 0x87, 0x37, 0x21, 0x5a, 0xd5, 0xf1, 0xed, 0x46, 0xd6, 0xea, 0x5e, 0xa0, 0x7b, 0x69,
	0x4f, 0x51, 0x38, 0x9e, 0x23, 0x1c, 0x7d, 0xa9, 0xa5, 0xb5, 0x2b, 0x56, 0x5d, 0xa1, 0x47, 0x29,
	0x48, 0x1e, 0x43, 0x2d, 0xc2, 0xc0, 0xb3, 0x2f, 0xd0, 0xf1, 0x90, 0x47, 0xfa, 0xb2, 0x22, 0x55,
	0x25, 0xf6, 0x26, 0x81, 0xc8, 0x13, 0xa8, 0x2b, 0x8a, 0xe0, 0x0e, 0xf5, 0x25, 0x67, 0x45, 0x71,
	0x54, 0xdc, 0x30, 0xc5, 0xc8, 0x33, 0x58, 0x8f, 0x04, 0x47, 0x67, 0x6a, 0x53, 0x0f, 0x03, 0x41,
	0xc7, 0x14, 0xb9, 0x5e, 0x51, 0xe2, 0x35, 0x92, 0x83, 0x83, 0x9f, 0xb8, 0xf1, 0x02, 0xaa, 0x33,
	0x75, 0x93, 0x0a, 0x94, 0x07, 0xc7, 0x83, 0x7e, 0xa3, 0x20, 0x57, 0xbd, 0xe3, 0xfd, 0x7e, 0x43,
	0x23, 0x0d, 0xa8, 0xc9, 0x95, 0x7d, 0x36, 0x38, 0x50, 0x48, 0xc9, 0x28, 0x57, 0x8a, 0x8d, 0xa2,
	0xf1, 0x12, 0x6a, 0xc9, 0xf3, 0x13, 0xc1, 0x65, 0xa3, 0xde, 0xcd, 0x36, 0x4a, 0x6d, 0x88, 0x0e,
	0x2b, 0xaa, 0x3f, 0xc8, 0xd3, 0x0e, 0x65, 0x5b, 0x83, 0x42, 0x75, 0x88, 0xd7, 0xe2, 0x08, 0xa3,
	0xc8, 0x99, 0x20, 0x21, 0x50, 0x16, 0x78, 0x2d, 0xd2, 0x68, 0xb5, 0xbe, 0x23, 0x49, 0x71, 0x0e,
	0x49, 0x4a, 0x77, 0x25, 0x31, 0xf6, 0xe1, 0xff, 0x1e, 0x0b, 0x04, 0x0d, 0x62, 0x3c, 0x55, 0x0a,
	0x64, 0xc3, 0x75, 0xaf, 0x56, 0x5a, 0x8e, 0x56, 0xaf, 0x41, 0xef, 0xc9, 0x8e, 0x25, 0x29, 0x7a,
	0x3e, 0x8b, 0xd0, 0xfb, 0xab, 0x44, 0xbb, 0xb0, 0x75, 0x4f, 0xa2, 0x54, 0xc6, 0x0d, 0x58, 0x76,
	0x15, 0xa2, 0xc2, 0x2b, 0x56, 0xba, 0xeb, 0x7e, 0x2f, 0x4b, 0xbd, 0x22, 0x71, 0x8a, 0xfc, 0x8a,
	0xba, 0x48, 0x0e, 0x61, 0x55, 0xca, 0xdf, 0x97, 0xd6, 0x21, 0x1b, 0x66, 0x62, 0x29, 0x33, 0xb3,
	0x94, 0xa9, 0xf0, 0xe6, 0xf6, 0x6f, 0x26, 0x37, 0xb9, 0xd3, 0x28, 0x90, 0x33, 0x28, 0x4b, 0x84,
	0x3c, 0x9d, 0x67, 0xd4, 0xe7, 0x4f, 0xfb, 0x36, 0x2d, 0x52, 0x1a, 0x6d, 0xce, 0xdc, 0x39, 0x4f,
	0x31, 0x0a, 0xe4, 0x3d, 0x54, 0x24, 0xf1, 0x90, 0x46, 0x62, 0xc1, 0x75, 0xee, 0x68, 0xc4, 0x49,
	0xa6, 0xf9, 0x84, 0x05, 0x93, 0x3d, 0xea, 0xd1, 0x05, 0x5f, 0xd0, 0xd6, 0x76, 0x34, 0xf2, 0x11,
	0x40, 0xa2, 0x49, 0xd7, 0x17, 0x7e, 0x01, 0x19, 0x42, 0xb9, 0xef, 0x5e, 0xb0, 0x07, 0x12, 0xcf,
	0xd8, 0xad, 0x39, 0x17, 0xcb, 0x28, 0x74, 0xbf, 0x69, 0xf0, 0x8f, 0x1c, 0xbb, 0x33, 0x41, 0xfd,
	0x6c, 0xf4, 0xce, 0x61, 0xed, 0x57, 0x3b, 0x11, 0x33, 0x37, 0xdb, 0xbd, 0xbe, 0x7b, 0xa0, 0xc9,
	0x5f, 0x60, 0xfd, 0x8e, 0x37, 0xc8, 0xf3, 0xfc, 0xf4, 0x39, 0x86, 0x6c, 0x76, 0xff, 0x24, 0x24,
	0x53, 0xb1, 0x7b, 0x99, 0xfc, 0x1d, 0x66, 0x0f, 0xfd, 0x00, 0xd5, 0x01, 0x0b, 0xfa, 0xd7, 0x34,
	0x12, 0x4e, 0xb0, 0xe8, 0xa1, 0xdb, 0xdb, 0x3a, 0xdf, 0xcc, 0xf9, 0x8a, 0x8e, 0x96, 0x95, 0x2e,
	0xbb, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x82, 0xcb, 0x10, 0xd2, 0x67, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	PingEmpty(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error)
	PingPongBidi(ctx context.Context, opts ...grpc.CallOption) (TestService_PingPongBidiClient, error)
	PingStream(ctx context.Context, opts ...grpc.CallOption) (TestService_PingStreamClient, error)
	Echo(ctx context.Context, in *TextMessage, opts ...grpc.CallOption) (*TextMessage, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) PingEmpty(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/improbable.grpcweb.test.TestService/PingEmpty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/improbable.grpcweb.test.TestService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/improbable.grpcweb.test.TestService/PingError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/improbable.grpcweb.test.TestService/PingList", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_PingListClient interface {
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingListClient struct {
	grpc.ClientStream
}

func (x *testServicePingListClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) PingPongBidi(ctx context.Context, opts ...grpc.CallOption) (TestService_PingPongBidiClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[1], "/improbable.grpcweb.test.TestService/PingPongBidi", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingPongBidiClient{stream}
	return x, nil
}

type TestService_PingPongBidiClient interface {
	Send(*PingRequest) error
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingPongBidiClient struct {
	grpc.ClientStream
}

func (x *testServicePingPongBidiClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServicePingPongBidiClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) PingStream(ctx context.Context, opts ...grpc.CallOption) (TestService_PingStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[2], "/improbable.grpcweb.test.TestService/PingStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingStreamClient{stream}
	return x, nil
}

type TestService_PingStreamClient interface {
	Send(*PingRequest) error
	CloseAndRecv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingStreamClient struct {
	grpc.ClientStream
}

func (x *testServicePingStreamClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServicePingStreamClient) CloseAndRecv() (*PingResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) Echo(ctx context.Context, in *TextMessage, opts ...grpc.CallOption) (*TextMessage, error) {
	out := new(TextMessage)
	err := c.cc.Invoke(ctx, "/improbable.grpcweb.test.TestService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	PingEmpty(context.Context, *empty.Empty) (*PingResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	PingError(context.Context, *PingRequest) (*empty.Empty, error)
	PingList(*PingRequest, TestService_PingListServer) error
	PingPongBidi(TestService_PingPongBidiServer) error
	PingStream(TestService_PingStreamServer) error
	Echo(context.Context, *TextMessage) (*TextMessage, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_PingEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingEmpty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.TestService/PingEmpty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingEmpty(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.TestService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.TestService/PingError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingError(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).PingList(m, &testServicePingListServer{stream})
}

type TestService_PingListServer interface {
	Send(*PingResponse) error
	grpc.ServerStream
}

type testServicePingListServer struct {
	grpc.ServerStream
}

func (x *testServicePingListServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_PingPongBidi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).PingPongBidi(&testServicePingPongBidiServer{stream})
}

type TestService_PingPongBidiServer interface {
	Send(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type testServicePingPongBidiServer struct {
	grpc.ServerStream
}

func (x *testServicePingPongBidiServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServicePingPongBidiServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_PingStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).PingStream(&testServicePingStreamServer{stream})
}

type TestService_PingStreamServer interface {
	SendAndClose(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type testServicePingStreamServer struct {
	grpc.ServerStream
}

func (x *testServicePingStreamServer) SendAndClose(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServicePingStreamServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.TestService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Echo(ctx, req.(*TextMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.grpcweb.test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingEmpty",
			Handler:    _TestService_PingEmpty_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TestService_Ping_Handler,
		},
		{
			MethodName: "PingError",
			Handler:    _TestService_PingError_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _TestService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingList",
			Handler:       _TestService_PingList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PingPongBidi",
			Handler:       _TestService_PingPongBidi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PingStream",
			Handler:       _TestService_PingStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "improbable/grpcweb/test/test.proto",
}

// TestUtilServiceClient is the client API for TestUtilService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestUtilServiceClient interface {
	ContinueStream(ctx context.Context, in *ContinueStreamRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CheckStreamClosed(ctx context.Context, in *CheckStreamClosedRequest, opts ...grpc.CallOption) (*CheckStreamClosedResponse, error)
}

type testUtilServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestUtilServiceClient(cc *grpc.ClientConn) TestUtilServiceClient {
	return &testUtilServiceClient{cc}
}

func (c *testUtilServiceClient) ContinueStream(ctx context.Context, in *ContinueStreamRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/improbable.grpcweb.test.TestUtilService/ContinueStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testUtilServiceClient) CheckStreamClosed(ctx context.Context, in *CheckStreamClosedRequest, opts ...grpc.CallOption) (*CheckStreamClosedResponse, error) {
	out := new(CheckStreamClosedResponse)
	err := c.cc.Invoke(ctx, "/improbable.grpcweb.test.TestUtilService/CheckStreamClosed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestUtilServiceServer is the server API for TestUtilService service.
type TestUtilServiceServer interface {
	ContinueStream(context.Context, *ContinueStreamRequest) (*empty.Empty, error)
	CheckStreamClosed(context.Context, *CheckStreamClosedRequest) (*CheckStreamClosedResponse, error)
}

func RegisterTestUtilServiceServer(s *grpc.Server, srv TestUtilServiceServer) {
	s.RegisterService(&_TestUtilService_serviceDesc, srv)
}

func _TestUtilService_ContinueStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContinueStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUtilServiceServer).ContinueStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.TestUtilService/ContinueStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUtilServiceServer).ContinueStream(ctx, req.(*ContinueStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestUtilService_CheckStreamClosed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStreamClosedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUtilServiceServer).CheckStreamClosed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.TestUtilService/CheckStreamClosed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUtilServiceServer).CheckStreamClosed(ctx, req.(*CheckStreamClosedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestUtilService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.grpcweb.test.TestUtilService",
	HandlerType: (*TestUtilServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContinueStream",
			Handler:    _TestUtilService_ContinueStream_Handler,
		},
		{
			MethodName: "CheckStreamClosed",
			Handler:    _TestUtilService_CheckStreamClosed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/grpcweb/test/test.proto",
}

// FailServiceClient is the client API for FailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FailServiceClient interface {
	NonExistant(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type failServiceClient struct {
	cc *grpc.ClientConn
}

func NewFailServiceClient(cc *grpc.ClientConn) FailServiceClient {
	return &failServiceClient{cc}
}

func (c *failServiceClient) NonExistant(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/improbable.grpcweb.test.FailService/NonExistant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FailServiceServer is the server API for FailService service.
type FailServiceServer interface {
	NonExistant(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterFailServiceServer(s *grpc.Server, srv FailServiceServer) {
	s.RegisterService(&_FailService_serviceDesc, srv)
}

func _FailService_NonExistant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailServiceServer).NonExistant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.FailService/NonExistant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailServiceServer).NonExistant(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.grpcweb.test.FailService",
	HandlerType: (*FailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NonExistant",
			Handler:    _FailService_NonExistant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/grpcweb/test/test.proto",
}
