// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package request

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status struct {
	Callback             string   `protobuf:"bytes,1,opt,name=callback,proto3" json:"callback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCallback() string {
	if m != nil {
		return m.Callback
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "request.Status")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_7f73548e33e655fe) }

var fileDescriptor_7f73548e33e655fe = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x54, 0xb8,
	0xd8, 0x82, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x85, 0xa4, 0xb8, 0x38, 0x92, 0x13, 0x73, 0x72, 0x92,
	0x12, 0x93, 0xb3, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x27, 0xfd, 0x28, 0xdd,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xd2, 0xa4, 0xe4, 0x9c, 0xc4,
	0xd2, 0xbc, 0xe4, 0x8c, 0x82, 0xc4, 0x14, 0xfd, 0x82, 0xcc, 0xbc, 0x82, 0xfc, 0xcc, 0xbc, 0x12,
	0xfd, 0xf4, 0xa2, 0x82, 0x64, 0x7d, 0xa8, 0xb1, 0x49, 0x6c, 0x60, 0x6b, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x2a, 0x55, 0x91, 0xc8, 0x77, 0x00, 0x00, 0x00,
}
