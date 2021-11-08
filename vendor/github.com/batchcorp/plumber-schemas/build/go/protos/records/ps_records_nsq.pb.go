// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/ps_records_nsq.proto

package records

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

type NSQ struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Channel              string   `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Attempts             int32    `protobuf:"varint,4,opt,name=attempts,proto3" json:"attempts,omitempty"`
	NsqdAddress          string   `protobuf:"bytes,5,opt,name=nsqd_address,json=nsqdAddress,proto3" json:"nsqd_address,omitempty"`
	Value                []byte   `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            int64    `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NSQ) Reset()         { *m = NSQ{} }
func (m *NSQ) String() string { return proto.CompactTextString(m) }
func (*NSQ) ProtoMessage()    {}
func (*NSQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_0733f38f551f19cf, []int{0}
}

func (m *NSQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NSQ.Unmarshal(m, b)
}
func (m *NSQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NSQ.Marshal(b, m, deterministic)
}
func (m *NSQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NSQ.Merge(m, src)
}
func (m *NSQ) XXX_Size() int {
	return xxx_messageInfo_NSQ.Size(m)
}
func (m *NSQ) XXX_DiscardUnknown() {
	xxx_messageInfo_NSQ.DiscardUnknown(m)
}

var xxx_messageInfo_NSQ proto.InternalMessageInfo

func (m *NSQ) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NSQ) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *NSQ) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *NSQ) GetAttempts() int32 {
	if m != nil {
		return m.Attempts
	}
	return 0
}

func (m *NSQ) GetNsqdAddress() string {
	if m != nil {
		return m.NsqdAddress
	}
	return ""
}

func (m *NSQ) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *NSQ) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*NSQ)(nil), "protos.records.NSQ")
}

func init() { proto.RegisterFile("records/ps_records_nsq.proto", fileDescriptor_0733f38f551f19cf) }

var fileDescriptor_0733f38f551f19cf = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xbd, 0x4e, 0xec, 0x30,
	0x10, 0x85, 0xe5, 0xe4, 0x66, 0xf7, 0xae, 0x59, 0x6d, 0x61, 0x51, 0x58, 0x68, 0x8b, 0x40, 0x95,
	0x86, 0xb8, 0xa0, 0x45, 0x48, 0xf0, 0x00, 0x48, 0x84, 0x8e, 0x26, 0xf2, 0x9f, 0x36, 0x96, 0xe2,
	0xd8, 0xeb, 0x99, 0xf0, 0x6c, 0x3c, 0x1e, 0x5a, 0x27, 0x40, 0x35, 0xf3, 0x7d, 0xa3, 0xd1, 0xd1,
	0xa1, 0xc7, 0x64, 0x75, 0x48, 0x06, 0x44, 0x84, 0x7e, 0x5d, 0xfb, 0x09, 0xce, 0x6d, 0x4c, 0x01,
	0x03, 0x3b, 0xe4, 0x01, 0xed, 0x7a, 0xb9, 0xfb, 0x22, 0xb4, 0x7c, 0x7d, 0x7f, 0x63, 0x07, 0x5a,
	0x38, 0xc3, 0x49, 0x4d, 0x9a, 0x5d, 0x57, 0x38, 0xc3, 0xae, 0x69, 0x85, 0x21, 0x3a, 0xcd, 0x8b,
	0xac, 0x16, 0x60, 0x9c, 0x6e, 0xf5, 0x20, 0xa7, 0xc9, 0x8e, 0xbc, 0xcc, 0xfe, 0x07, 0xd9, 0x0d,
	0xfd, 0x2f, 0x11, 0xad, 0x8f, 0x08, 0xfc, 0x5f, 0x4d, 0x9a, 0xaa, 0xfb, 0x65, 0x76, 0x4b, 0xf7,
	0x13, 0x9c, 0x4d, 0x2f, 0x8d, 0x49, 0x16, 0x80, 0x57, 0xf9, 0xf5, 0xea, 0xe2, 0x9e, 0x17, 0x75,
	0x89, 0xfb, 0x94, 0xe3, 0x6c, 0xf9, 0xa6, 0x26, 0xcd, 0xbe, 0x5b, 0x80, 0x1d, 0xe9, 0x0e, 0x9d,
	0xb7, 0x80, 0xd2, 0x47, 0xbe, 0xad, 0x49, 0x53, 0x76, 0x7f, 0xe2, 0xe5, 0xe9, 0xe3, 0xf1, 0xe4,
	0x70, 0x98, 0x55, 0xab, 0x83, 0x17, 0x4a, 0xa2, 0x1e, 0x74, 0x48, 0x51, 0xc4, 0x71, 0xf6, 0xca,
	0xa6, 0x7b, 0xd0, 0x83, 0xf5, 0x12, 0x84, 0x9a, 0xdd, 0x68, 0xc4, 0x29, 0x88, 0xa5, 0xba, 0x58,
	0xab, 0xab, 0x4d, 0xe6, 0x87, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xce, 0x6c, 0x86, 0x31,
	0x01, 0x00, 0x00,
}