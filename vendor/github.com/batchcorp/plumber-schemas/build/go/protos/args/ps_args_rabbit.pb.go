// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_rabbit.proto

package args

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

type RabbitConn struct {
	// @gotags: kong:"help='Destination host address (full DSN)',env='PLUMBER_RELAY_RABBIT_ADDRESS',default='amqp://localhost',required"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" kong:"help='Destination host address (full DSN)',env='PLUMBER_RELAY_RABBIT_ADDRESS',default='amqp://localhost',required"`
	// @gotags: kong:"help='Force TLS usage (regardless of DSN)',env='PLUMBER_RELAY_RABBIT_USE_TLS'"
	UseTls bool `protobuf:"varint,2,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty" kong:"help='Force TLS usage (regardless of DSN)',env='PLUMBER_RELAY_RABBIT_USE_TLS'"`
	// @gotags: kong:"help='Whether to verify server TLS certificate',env='PLUMBER_RELAY_RABBIT_SKIP_VERIFY_TLS'"
	TlsSkipVerify        bool     `protobuf:"varint,3,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty" kong:"help='Whether to verify server TLS certificate',env='PLUMBER_RELAY_RABBIT_SKIP_VERIFY_TLS'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RabbitConn) Reset()         { *m = RabbitConn{} }
func (m *RabbitConn) String() string { return proto.CompactTextString(m) }
func (*RabbitConn) ProtoMessage()    {}
func (*RabbitConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_01d1eee3dc8ebf97, []int{0}
}

func (m *RabbitConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RabbitConn.Unmarshal(m, b)
}
func (m *RabbitConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RabbitConn.Marshal(b, m, deterministic)
}
func (m *RabbitConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RabbitConn.Merge(m, src)
}
func (m *RabbitConn) XXX_Size() int {
	return xxx_messageInfo_RabbitConn.Size(m)
}
func (m *RabbitConn) XXX_DiscardUnknown() {
	xxx_messageInfo_RabbitConn.DiscardUnknown(m)
}

var xxx_messageInfo_RabbitConn proto.InternalMessageInfo

func (m *RabbitConn) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RabbitConn) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *RabbitConn) GetTlsSkipVerify() bool {
	if m != nil {
		return m.TlsSkipVerify
	}
	return false
}

type RabbitReadArgs struct {
	// @gotags: kong:"help='Name of the exchange',env='PLUMBER_RELAY_RABBIT_EXCHANGE',required"
	ExchangeName string `protobuf:"bytes,1,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty" kong:"help='Name of the exchange',env='PLUMBER_RELAY_RABBIT_EXCHANGE',required"`
	// @gotags: kong:"help='Name of the queue where messages will be routed to',env='PLUMBER_RELAY_RABBIT_QUEUE',required"
	QueueName string `protobuf:"bytes,2,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty" kong:"help='Name of the queue where messages will be routed to',env='PLUMBER_RELAY_RABBIT_QUEUE',required"`
	// @gotags: kong:"help='Binding key for topic based exchanges',env='PLUMBER_RELAY_RABBIT_ROUTING_KEY',required"
	BindingKey string `protobuf:"bytes,3,opt,name=binding_key,json=bindingKey,proto3" json:"binding_key,omitempty" kong:"help='Binding key for topic based exchanges',env='PLUMBER_RELAY_RABBIT_ROUTING_KEY',required"`
	// @gotags: kong:"help='Whether plumber should be the only one using the queue',env='PLUMBER_RELAY_RABBIT_QUEUE_EXCLUSIVE'"
	QueueExclusive bool `protobuf:"varint,4,opt,name=queue_exclusive,json=queueExclusive,proto3" json:"queue_exclusive,omitempty" kong:"help='Whether plumber should be the only one using the queue',env='PLUMBER_RELAY_RABBIT_QUEUE_EXCLUSIVE'"`
	// @gotags: kong:"help='Whether to create/declare the queue (if it does not exist)',env='PLUMBER_RELAY_RABBIT_QUEUE_DECLARE',default=true"
	QueueDeclare bool `protobuf:"varint,5,opt,name=queue_declare,json=queueDeclare,proto3" json:"queue_declare,omitempty" kong:"help='Whether to create/declare the queue (if it does not exist)',env='PLUMBER_RELAY_RABBIT_QUEUE_DECLARE',default=true"`
	// @gotags: kong:"help='Whether the queue should survive after disconnect',env='PLUMBER_RELAY_RABBIT_QUEUE_DURABLE'"
	QueueDurable bool `protobuf:"varint,6,opt,name=queue_durable,json=queueDurable,proto3" json:"queue_durable,omitempty" kong:"help='Whether the queue should survive after disconnect',env='PLUMBER_RELAY_RABBIT_QUEUE_DURABLE'"`
	// @gotags: kong:"help='Automatically acknowledge receipt of read/received messages',env='PLUMBER_RELAY_RABBIT_AUTOACK',default=true"
	AutoAck bool `protobuf:"varint,7,opt,name=auto_ack,json=autoAck,proto3" json:"auto_ack,omitempty" kong:"help='Automatically acknowledge receipt of read/received messages',env='PLUMBER_RELAY_RABBIT_AUTOACK',default=true"`
	// @gotags: kong:"help='How to identify the consumer to RabbitMQ',env='PLUMBER_RELAY_CONSUMER_TAG',default=plumber"
	ConsumerTag string `protobuf:"bytes,8,opt,name=consumer_tag,json=consumerTag,proto3" json:"consumer_tag,omitempty" kong:"help='How to identify the consumer to RabbitMQ',env='PLUMBER_RELAY_CONSUMER_TAG',default=plumber"`
	// @gotags: kong:"help='Whether to auto-delete the queue after plumber has disconnected',env='PLUMBER_RELAY_RABBIT_QUEUE_AUTO_DELETE',default=true"
	QueueDelete          bool     `protobuf:"varint,9,opt,name=queue_delete,json=queueDelete,proto3" json:"queue_delete,omitempty" kong:"help='Whether to auto-delete the queue after plumber has disconnected',env='PLUMBER_RELAY_RABBIT_QUEUE_AUTO_DELETE',default=true"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RabbitReadArgs) Reset()         { *m = RabbitReadArgs{} }
func (m *RabbitReadArgs) String() string { return proto.CompactTextString(m) }
func (*RabbitReadArgs) ProtoMessage()    {}
func (*RabbitReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_01d1eee3dc8ebf97, []int{1}
}

func (m *RabbitReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RabbitReadArgs.Unmarshal(m, b)
}
func (m *RabbitReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RabbitReadArgs.Marshal(b, m, deterministic)
}
func (m *RabbitReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RabbitReadArgs.Merge(m, src)
}
func (m *RabbitReadArgs) XXX_Size() int {
	return xxx_messageInfo_RabbitReadArgs.Size(m)
}
func (m *RabbitReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RabbitReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RabbitReadArgs proto.InternalMessageInfo

func (m *RabbitReadArgs) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *RabbitReadArgs) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *RabbitReadArgs) GetBindingKey() string {
	if m != nil {
		return m.BindingKey
	}
	return ""
}

func (m *RabbitReadArgs) GetQueueExclusive() bool {
	if m != nil {
		return m.QueueExclusive
	}
	return false
}

func (m *RabbitReadArgs) GetQueueDeclare() bool {
	if m != nil {
		return m.QueueDeclare
	}
	return false
}

func (m *RabbitReadArgs) GetQueueDurable() bool {
	if m != nil {
		return m.QueueDurable
	}
	return false
}

func (m *RabbitReadArgs) GetAutoAck() bool {
	if m != nil {
		return m.AutoAck
	}
	return false
}

func (m *RabbitReadArgs) GetConsumerTag() string {
	if m != nil {
		return m.ConsumerTag
	}
	return ""
}

func (m *RabbitReadArgs) GetQueueDelete() bool {
	if m != nil {
		return m.QueueDelete
	}
	return false
}

type RabbitWriteArgs struct {
	// @gotags: kong:"help='Exchange to write message(s) to',required"
	ExchangeName string `protobuf:"bytes,1,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty" kong:"help='Exchange to write message(s) to',required"`
	// @gotags: kong:"help='Routing key to write message(s) to',required"
	RoutingKey string `protobuf:"bytes,2,opt,name=routing_key,json=routingKey,proto3" json:"routing_key,omitempty" kong:"help='Routing key to write message(s) to',required"`
	// @gotags: kong:"help='Fills message properties $app_id with this value',default=plumber"
	AppId string `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" kong:"help='Fills message properties  with this value',default=plumber"`
	// @gotags: kong:"help='The type of exchange we are working with',enum='direct,topic,headers,fanout',default=topic,group=exchange"
	ExchangeType string `protobuf:"bytes,4,opt,name=exchange_type,json=exchangeType,proto3" json:"exchange_type,omitempty" kong:"help='The type of exchange we are working with',enum='direct,topic,headers,fanout',default=topic,group=exchange"`
	// @gotags: kong:"help='Whether to declare an exchange (if it does not exist)',group=exchange"
	ExchangeDeclare bool `protobuf:"varint,5,opt,name=exchange_declare,json=exchangeDeclare,proto3" json:"exchange_declare,omitempty" kong:"help='Whether to declare an exchange (if it does not exist)',group=exchange"`
	// @gotags: kong:"help='Whether to make a declared exchange durable',group=exchange"
	ExchangeDurable bool `protobuf:"varint,6,opt,name=exchange_durable,json=exchangeDurable,proto3" json:"exchange_durable,omitempty" kong:"help='Whether to make a declared exchange durable',group=exchange"`
	// @gotags: kong:"help='Whether to auto-delete the exchange (after writes)',group=exchange"
	ExchangeAutoDelete   bool     `protobuf:"varint,7,opt,name=exchange_auto_delete,json=exchangeAutoDelete,proto3" json:"exchange_auto_delete,omitempty" kong:"help='Whether to auto-delete the exchange (after writes)',group=exchange"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RabbitWriteArgs) Reset()         { *m = RabbitWriteArgs{} }
func (m *RabbitWriteArgs) String() string { return proto.CompactTextString(m) }
func (*RabbitWriteArgs) ProtoMessage()    {}
func (*RabbitWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_01d1eee3dc8ebf97, []int{2}
}

func (m *RabbitWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RabbitWriteArgs.Unmarshal(m, b)
}
func (m *RabbitWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RabbitWriteArgs.Marshal(b, m, deterministic)
}
func (m *RabbitWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RabbitWriteArgs.Merge(m, src)
}
func (m *RabbitWriteArgs) XXX_Size() int {
	return xxx_messageInfo_RabbitWriteArgs.Size(m)
}
func (m *RabbitWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RabbitWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RabbitWriteArgs proto.InternalMessageInfo

func (m *RabbitWriteArgs) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *RabbitWriteArgs) GetRoutingKey() string {
	if m != nil {
		return m.RoutingKey
	}
	return ""
}

func (m *RabbitWriteArgs) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *RabbitWriteArgs) GetExchangeType() string {
	if m != nil {
		return m.ExchangeType
	}
	return ""
}

func (m *RabbitWriteArgs) GetExchangeDeclare() bool {
	if m != nil {
		return m.ExchangeDeclare
	}
	return false
}

func (m *RabbitWriteArgs) GetExchangeDurable() bool {
	if m != nil {
		return m.ExchangeDurable
	}
	return false
}

func (m *RabbitWriteArgs) GetExchangeAutoDelete() bool {
	if m != nil {
		return m.ExchangeAutoDelete
	}
	return false
}

func init() {
	proto.RegisterType((*RabbitConn)(nil), "protos.args.RabbitConn")
	proto.RegisterType((*RabbitReadArgs)(nil), "protos.args.RabbitReadArgs")
	proto.RegisterType((*RabbitWriteArgs)(nil), "protos.args.RabbitWriteArgs")
}

func init() { proto.RegisterFile("ps_args_rabbit.proto", fileDescriptor_01d1eee3dc8ebf97) }

var fileDescriptor_01d1eee3dc8ebf97 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0xd5, 0xc2, 0xf6, 0x63, 0xda, 0xdd, 0x22, 0x6b, 0x11, 0xe6, 0x80, 0x76, 0x29, 0x12,
	0x2c, 0x07, 0x1a, 0x24, 0x4e, 0x88, 0x53, 0xf9, 0x38, 0x20, 0x24, 0x0e, 0xa5, 0x02, 0x89, 0x8b,
	0xe5, 0x38, 0x43, 0x6a, 0xc5, 0x89, 0x8d, 0x3f, 0x56, 0xdb, 0x67, 0xe1, 0x89, 0x78, 0x2b, 0x14,
	0x3b, 0x59, 0xda, 0xdb, 0x9e, 0xa2, 0xf9, 0xcd, 0x3f, 0x33, 0xa3, 0x9f, 0x64, 0x38, 0x37, 0x8e,
	0x71, 0x5b, 0x3a, 0x66, 0x79, 0x9e, 0x4b, 0xbf, 0x32, 0x56, 0x7b, 0x4d, 0x66, 0xf1, 0xe3, 0x56,
	0x6d, 0x67, 0x59, 0x02, 0x6c, 0x62, 0xf3, 0x83, 0x6e, 0x1a, 0x42, 0x61, 0xcc, 0x8b, 0xc2, 0xa2,
	0x73, 0x74, 0x70, 0x39, 0xb8, 0x9a, 0x6e, 0xfa, 0x92, 0x3c, 0x82, 0x71, 0x70, 0xc8, 0xbc, 0x72,
	0x74, 0x78, 0x39, 0xb8, 0x9a, 0x6c, 0x46, 0xc1, 0xe1, 0x56, 0x39, 0xf2, 0x1c, 0x16, 0x5e, 0x39,
	0xe6, 0x2a, 0x69, 0xd8, 0x35, 0x5a, 0xf9, 0x6b, 0x4f, 0xef, 0xc5, 0xc0, 0xa9, 0x57, 0xee, 0x5b,
	0x25, 0xcd, 0xf7, 0x08, 0x97, 0x7f, 0x87, 0x70, 0x96, 0x36, 0x6d, 0x90, 0x17, 0x6b, 0x5b, 0x3a,
	0xf2, 0x0c, 0x4e, 0xf1, 0x46, 0xec, 0x78, 0x53, 0x22, 0x6b, 0x78, 0x8d, 0xdd, 0xce, 0x79, 0x0f,
	0xbf, 0xf2, 0x1a, 0xc9, 0x13, 0x80, 0xdf, 0x01, 0x43, 0x97, 0x18, 0xc6, 0xc4, 0x34, 0x92, 0xd8,
	0xbe, 0x80, 0x59, 0x2e, 0x9b, 0x42, 0x36, 0x25, 0xab, 0x30, 0xad, 0x9e, 0x6e, 0xa0, 0x43, 0x5f,
	0x70, 0x4f, 0x5e, 0xc0, 0x22, 0xfd, 0x8f, 0x37, 0x42, 0x05, 0x27, 0xaf, 0x91, 0xde, 0x8f, 0xf7,
	0x9d, 0x45, 0xfc, 0xa9, 0xa7, 0xed, 0x35, 0x29, 0x58, 0xa0, 0x50, 0xdc, 0x22, 0x3d, 0x89, 0xb1,
	0x79, 0x84, 0x1f, 0x13, 0x3b, 0x08, 0x05, 0xcb, 0x73, 0x85, 0x74, 0x74, 0x18, 0x4a, 0x8c, 0x3c,
	0x86, 0x09, 0x0f, 0x5e, 0x33, 0x2e, 0x2a, 0x3a, 0x8e, 0xfd, 0x71, 0x5b, 0xaf, 0x45, 0x45, 0x9e,
	0xc2, 0x5c, 0xe8, 0xc6, 0x85, 0x1a, 0x2d, 0xf3, 0xbc, 0xa4, 0x93, 0x78, 0xef, 0xac, 0x67, 0x5b,
	0x5e, 0xb6, 0x91, 0xfe, 0x0e, 0x85, 0x1e, 0xe9, 0x34, 0x4e, 0x98, 0x75, 0x67, 0xb4, 0x68, 0xf9,
	0x67, 0x08, 0x8b, 0xe4, 0xf2, 0x87, 0x95, 0x1e, 0xef, 0x2e, 0xf3, 0x02, 0x66, 0x56, 0x07, 0xdf,
	0xdb, 0x4a, 0x36, 0xa1, 0x43, 0xad, 0xad, 0x87, 0x30, 0xe2, 0xc6, 0x30, 0x59, 0x74, 0x26, 0x4f,
	0xb8, 0x31, 0x9f, 0x8b, 0xa3, 0xe1, 0x7e, 0x6f, 0x92, 0xc2, 0x83, 0xe1, 0xdb, 0xbd, 0x41, 0xf2,
	0x12, 0x1e, 0xdc, 0x86, 0x8e, 0x1d, 0x2e, 0x7a, 0xde, 0x6b, 0x3c, 0x8a, 0x1e, 0x99, 0xfc, 0x1f,
	0xed, 0x64, 0xbe, 0x86, 0xf3, 0xdb, 0x68, 0xb4, 0xda, 0x69, 0x49, 0x62, 0x49, 0xdf, 0x5b, 0x07,
	0xaf, 0x93, 0x9d, 0xf7, 0xef, 0x7e, 0xbe, 0x2d, 0xa5, 0xdf, 0x85, 0x7c, 0x25, 0x74, 0x9d, 0xe5,
	0xdc, 0x8b, 0x9d, 0xd0, 0xd6, 0x64, 0x46, 0x85, 0x3a, 0x47, 0xfb, 0xca, 0x89, 0x1d, 0xd6, 0xdc,
	0x65, 0x79, 0x90, 0xaa, 0xc8, 0x4a, 0x9d, 0xa5, 0xf7, 0x90, 0xb5, 0xef, 0x21, 0x1f, 0xc5, 0xe2,
	0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x55, 0x3c, 0x51, 0x3b, 0x03, 0x00, 0x00,
}
