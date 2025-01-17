// Code generated by protoc-gen-go. DO NOT EDIT.
// source: outbound.proto

package events

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

// Emitted by the reader to HSB which is then consumed by the replayer and dProxy
type Outbound struct {
	ReplayId             string   `protobuf:"bytes,1,opt,name=replay_id,json=replayId,proto3" json:"replay_id,omitempty"`
	Blob                 []byte   `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
	Last                 bool     `protobuf:"varint,3,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Outbound) Reset()         { *m = Outbound{} }
func (m *Outbound) String() string { return proto.CompactTextString(m) }
func (*Outbound) ProtoMessage()    {}
func (*Outbound) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dbaa15aa01abbc0, []int{0}
}

func (m *Outbound) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Outbound.Unmarshal(m, b)
}
func (m *Outbound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Outbound.Marshal(b, m, deterministic)
}
func (m *Outbound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outbound.Merge(m, src)
}
func (m *Outbound) XXX_Size() int {
	return xxx_messageInfo_Outbound.Size(m)
}
func (m *Outbound) XXX_DiscardUnknown() {
	xxx_messageInfo_Outbound.DiscardUnknown(m)
}

var xxx_messageInfo_Outbound proto.InternalMessageInfo

func (m *Outbound) GetReplayId() string {
	if m != nil {
		return m.ReplayId
	}
	return ""
}

func (m *Outbound) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

func (m *Outbound) GetLast() bool {
	if m != nil {
		return m.Last
	}
	return false
}

func init() {
	proto.RegisterType((*Outbound)(nil), "events.Outbound")
}

func init() { proto.RegisterFile("outbound.proto", fileDescriptor_5dbaa15aa01abbc0) }

var fileDescriptor_5dbaa15aa01abbc0 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x46, 0x89, 0x4a, 0x69, 0x83, 0x38, 0x64, 0x2a, 0xb8, 0x14, 0xa7, 0x2e, 0x36, 0x83, 0xb3,
	0x08, 0x6e, 0x4e, 0x85, 0x8e, 0x2e, 0x92, 0x4b, 0x42, 0x5b, 0x48, 0x7b, 0x25, 0xb9, 0x08, 0xfe,
	0x7b, 0x31, 0x75, 0x7b, 0xdf, 0xfb, 0x96, 0xc7, 0x0f, 0x18, 0x09, 0x30, 0xce, 0xa6, 0x59, 0x3c,
	0x12, 0x8a, 0xcc, 0xbe, 0xed, 0x4c, 0xe1, 0xd4, 0xf2, 0xbc, 0xfd, 0x3f, 0xe2, 0xc8, 0x0b, 0x6f,
	0x17, 0xa7, 0x3e, 0xaf, 0xd1, 0x94, 0xac, 0x62, 0x75, 0xd1, 0xe5, 0xab, 0x78, 0x18, 0x21, 0xf8,
	0x0e, 0x1c, 0x42, 0xb9, 0xa9, 0x58, 0xbd, 0xef, 0x12, 0xff, 0x9c, 0x53, 0x81, 0xca, 0x6d, 0xc5,
	0xea, 0xbc, 0x4b, 0x7c, 0xbf, 0x3d, 0xaf, 0xfd, 0x48, 0x43, 0x84, 0x46, 0xe3, 0x24, 0x41, 0x91,
	0x1e, 0x34, 0xfa, 0x45, 0x6a, 0x74, 0xce, 0x6a, 0x42, 0x7f, 0x0e, 0x7a, 0xb0, 0x93, 0x0a, 0x12,
	0xe2, 0xe8, 0x8c, 0xec, 0x51, 0xa6, 0xa2, 0x20, 0xd7, 0x22, 0xc8, 0xd2, 0xbc, 0x7c, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x4e, 0x7b, 0xb9, 0x4b, 0xb2, 0x00, 0x00, 0x00,
}
