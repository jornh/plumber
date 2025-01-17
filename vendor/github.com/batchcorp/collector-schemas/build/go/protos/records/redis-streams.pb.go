// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/redis-streams.proto

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

type RedisStreamsRecord struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Stream               string   `protobuf:"bytes,4,opt,name=stream,proto3" json:"stream,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedisStreamsRecord) Reset()         { *m = RedisStreamsRecord{} }
func (m *RedisStreamsRecord) String() string { return proto.CompactTextString(m) }
func (*RedisStreamsRecord) ProtoMessage()    {}
func (*RedisStreamsRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b46bb4d4d13fdc84, []int{0}
}

func (m *RedisStreamsRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisStreamsRecord.Unmarshal(m, b)
}
func (m *RedisStreamsRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisStreamsRecord.Marshal(b, m, deterministic)
}
func (m *RedisStreamsRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisStreamsRecord.Merge(m, src)
}
func (m *RedisStreamsRecord) XXX_Size() int {
	return xxx_messageInfo_RedisStreamsRecord.Size(m)
}
func (m *RedisStreamsRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisStreamsRecord.DiscardUnknown(m)
}

var xxx_messageInfo_RedisStreamsRecord proto.InternalMessageInfo

func (m *RedisStreamsRecord) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RedisStreamsRecord) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RedisStreamsRecord) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RedisStreamsRecord) GetStream() string {
	if m != nil {
		return m.Stream
	}
	return ""
}

func (m *RedisStreamsRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*RedisStreamsRecord)(nil), "records.RedisStreamsRecord")
}

func init() { proto.RegisterFile("records/redis-streams.proto", fileDescriptor_b46bb4d4d13fdc84) }

var fileDescriptor_b46bb4d4d13fdc84 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x31, 0x4f, 0xc4, 0x30,
	0x0c, 0x85, 0xd5, 0x96, 0x3b, 0x74, 0x19, 0x10, 0x8a, 0x10, 0x44, 0x82, 0xe1, 0xc4, 0x74, 0xcb,
	0x25, 0x03, 0x3b, 0x03, 0x3f, 0xa1, 0x4c, 0xb0, 0xa5, 0x89, 0x75, 0x8d, 0x68, 0x70, 0x15, 0xbb,
	0x95, 0x58, 0xf9, 0xe5, 0xa8, 0x6e, 0x25, 0x36, 0xbf, 0xef, 0x3d, 0xcb, 0x7e, 0xea, 0xb1, 0x40,
	0xc0, 0x12, 0xc9, 0x15, 0x88, 0x89, 0xce, 0xc4, 0x05, 0x7c, 0x26, 0x3b, 0x16, 0x64, 0xd4, 0xd7,
	0x9b, 0xf9, 0xfc, 0x5b, 0x29, 0xdd, 0x2e, 0x81, 0xf7, 0xd5, 0x6f, 0x85, 0xeb, 0x1b, 0x55, 0xa7,
	0x68, 0xaa, 0x63, 0x75, 0x3a, 0xb4, 0x75, 0x8a, 0xfa, 0x56, 0x35, 0x5f, 0xf0, 0x63, 0x6a, 0x01,
	0xcb, 0xa8, 0xef, 0xd4, 0x6e, 0xf6, 0xc3, 0x04, 0xa6, 0x11, 0xb6, 0x0a, 0x7d, 0xaf, 0xf6, 0xeb,
	0x21, 0x73, 0x25, 0x78, 0x53, 0xfa, 0x49, 0x1d, 0x38, 0x65, 0x20, 0xf6, 0x79, 0x34, 0xbb, 0x63,
	0x75, 0x6a, 0xda, 0x7f, 0xf0, 0xf6, 0xa1, 0x1e, 0xa8, 0xb7, 0x9d, 0xe7, 0xd0, 0x5b, 0x98, 0xe1,
	0x9b, 0xc9, 0x6e, 0xff, 0x7d, 0xbe, 0x5e, 0x12, 0xf7, 0x53, 0x67, 0x03, 0x66, 0x27, 0x81, 0x80,
	0x65, 0x74, 0x01, 0x87, 0x01, 0x02, 0x63, 0x39, 0x53, 0xe8, 0x21, 0x7b, 0x72, 0xdd, 0x94, 0x86,
	0xe8, 0x2e, 0xe8, 0xa4, 0xdf, 0x52, 0x5a, 0xf6, 0xbb, 0xbd, 0xe8, 0x97, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9b, 0xae, 0x9d, 0x78, 0x0e, 0x01, 0x00, 0x00,
}
