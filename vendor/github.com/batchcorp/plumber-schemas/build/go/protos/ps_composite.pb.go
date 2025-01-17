// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_composite.proto

package protos

import (
	fmt "fmt"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
	opts "github.com/batchcorp/plumber-schemas/build/go/protos/opts"
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

type GetAllCompositesRequest struct {
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetAllCompositesRequest) Reset()         { *m = GetAllCompositesRequest{} }
func (m *GetAllCompositesRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllCompositesRequest) ProtoMessage()    {}
func (*GetAllCompositesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fc8df8568dab43f, []int{0}
}

func (m *GetAllCompositesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllCompositesRequest.Unmarshal(m, b)
}
func (m *GetAllCompositesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllCompositesRequest.Marshal(b, m, deterministic)
}
func (m *GetAllCompositesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCompositesRequest.Merge(m, src)
}
func (m *GetAllCompositesRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllCompositesRequest.Size(m)
}
func (m *GetAllCompositesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCompositesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCompositesRequest proto.InternalMessageInfo

func (m *GetAllCompositesRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type GetAllCompositesResponse struct {
	Composites           []*opts.Composite `protobuf:"bytes,1,rep,name=composites,proto3" json:"composites,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetAllCompositesResponse) Reset()         { *m = GetAllCompositesResponse{} }
func (m *GetAllCompositesResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllCompositesResponse) ProtoMessage()    {}
func (*GetAllCompositesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fc8df8568dab43f, []int{1}
}

func (m *GetAllCompositesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllCompositesResponse.Unmarshal(m, b)
}
func (m *GetAllCompositesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllCompositesResponse.Marshal(b, m, deterministic)
}
func (m *GetAllCompositesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCompositesResponse.Merge(m, src)
}
func (m *GetAllCompositesResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllCompositesResponse.Size(m)
}
func (m *GetAllCompositesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCompositesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCompositesResponse proto.InternalMessageInfo

func (m *GetAllCompositesResponse) GetComposites() []*opts.Composite {
	if m != nil {
		return m.Composites
	}
	return nil
}

type GetCompositeRequest struct {
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetCompositeRequest) Reset()         { *m = GetCompositeRequest{} }
func (m *GetCompositeRequest) String() string { return proto.CompactTextString(m) }
func (*GetCompositeRequest) ProtoMessage()    {}
func (*GetCompositeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fc8df8568dab43f, []int{2}
}

func (m *GetCompositeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCompositeRequest.Unmarshal(m, b)
}
func (m *GetCompositeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCompositeRequest.Marshal(b, m, deterministic)
}
func (m *GetCompositeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCompositeRequest.Merge(m, src)
}
func (m *GetCompositeRequest) XXX_Size() int {
	return xxx_messageInfo_GetCompositeRequest.Size(m)
}
func (m *GetCompositeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCompositeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCompositeRequest proto.InternalMessageInfo

func (m *GetCompositeRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *GetCompositeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetCompositeResponse struct {
	Composite            *opts.Composite `protobuf:"bytes,1,opt,name=composite,proto3" json:"composite,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetCompositeResponse) Reset()         { *m = GetCompositeResponse{} }
func (m *GetCompositeResponse) String() string { return proto.CompactTextString(m) }
func (*GetCompositeResponse) ProtoMessage()    {}
func (*GetCompositeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fc8df8568dab43f, []int{3}
}

func (m *GetCompositeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCompositeResponse.Unmarshal(m, b)
}
func (m *GetCompositeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCompositeResponse.Marshal(b, m, deterministic)
}
func (m *GetCompositeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCompositeResponse.Merge(m, src)
}
func (m *GetCompositeResponse) XXX_Size() int {
	return xxx_messageInfo_GetCompositeResponse.Size(m)
}
func (m *GetCompositeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCompositeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCompositeResponse proto.InternalMessageInfo

func (m *GetCompositeResponse) GetComposite() *opts.Composite {
	if m != nil {
		return m.Composite
	}
	return nil
}

type CreateCompositeRequest struct {
	Auth                 *common.Auth    `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Composite            *opts.Composite `protobuf:"bytes,2,opt,name=composite,proto3" json:"composite,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateCompositeRequest) Reset()         { *m = CreateCompositeRequest{} }
func (m *CreateCompositeRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCompositeRequest) ProtoMessage()    {}
func (*CreateCompositeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fc8df8568dab43f, []int{4}
}

func (m *CreateCompositeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCompositeRequest.Unmarshal(m, b)
}
func (m *CreateCompositeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCompositeRequest.Marshal(b, m, deterministic)
}
func (m *CreateCompositeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCompositeRequest.Merge(m, src)
}
func (m *CreateCompositeRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCompositeRequest.Size(m)
}
func (m *CreateCompositeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCompositeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCompositeRequest proto.InternalMessageInfo

func (m *CreateCompositeRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *CreateCompositeRequest) GetComposite() *opts.Composite {
	if m != nil {
		return m.Composite
	}
	return nil
}

type CreateCompositeResponse struct {
	Composite            *opts.Composite `protobuf:"bytes,1,opt,name=composite,proto3" json:"composite,omitempty"`
	Status               *common.Status  `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateCompositeResponse) Reset()         { *m = CreateCompositeResponse{} }
func (m *CreateCompositeResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCompositeResponse) ProtoMessage()    {}
func (*CreateCompositeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fc8df8568dab43f, []int{5}
}

func (m *CreateCompositeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCompositeResponse.Unmarshal(m, b)
}
func (m *CreateCompositeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCompositeResponse.Marshal(b, m, deterministic)
}
func (m *CreateCompositeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCompositeResponse.Merge(m, src)
}
func (m *CreateCompositeResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCompositeResponse.Size(m)
}
func (m *CreateCompositeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCompositeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCompositeResponse proto.InternalMessageInfo

func (m *CreateCompositeResponse) GetComposite() *opts.Composite {
	if m != nil {
		return m.Composite
	}
	return nil
}

func (m *CreateCompositeResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type UpdateCompositeRequest struct {
	Auth                 *common.Auth    `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Composite            *opts.Composite `protobuf:"bytes,2,opt,name=composite,proto3" json:"composite,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateCompositeRequest) Reset()         { *m = UpdateCompositeRequest{} }
func (m *UpdateCompositeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCompositeRequest) ProtoMessage()    {}
func (*UpdateCompositeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fc8df8568dab43f, []int{6}
}

func (m *UpdateCompositeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCompositeRequest.Unmarshal(m, b)
}
func (m *UpdateCompositeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCompositeRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCompositeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCompositeRequest.Merge(m, src)
}
func (m *UpdateCompositeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCompositeRequest.Size(m)
}
func (m *UpdateCompositeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCompositeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCompositeRequest proto.InternalMessageInfo

func (m *UpdateCompositeRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UpdateCompositeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateCompositeRequest) GetComposite() *opts.Composite {
	if m != nil {
		return m.Composite
	}
	return nil
}

type UpdateCompositeResponse struct {
	Composite            *opts.Composite `protobuf:"bytes,1,opt,name=composite,proto3" json:"composite,omitempty"`
	Status               *common.Status  `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateCompositeResponse) Reset()         { *m = UpdateCompositeResponse{} }
func (m *UpdateCompositeResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCompositeResponse) ProtoMessage()    {}
func (*UpdateCompositeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fc8df8568dab43f, []int{7}
}

func (m *UpdateCompositeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCompositeResponse.Unmarshal(m, b)
}
func (m *UpdateCompositeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCompositeResponse.Marshal(b, m, deterministic)
}
func (m *UpdateCompositeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCompositeResponse.Merge(m, src)
}
func (m *UpdateCompositeResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCompositeResponse.Size(m)
}
func (m *UpdateCompositeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCompositeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCompositeResponse proto.InternalMessageInfo

func (m *UpdateCompositeResponse) GetComposite() *opts.Composite {
	if m != nil {
		return m.Composite
	}
	return nil
}

func (m *UpdateCompositeResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type DeleteCompositeRequest struct {
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteCompositeRequest) Reset()         { *m = DeleteCompositeRequest{} }
func (m *DeleteCompositeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCompositeRequest) ProtoMessage()    {}
func (*DeleteCompositeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fc8df8568dab43f, []int{8}
}

func (m *DeleteCompositeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCompositeRequest.Unmarshal(m, b)
}
func (m *DeleteCompositeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCompositeRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCompositeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCompositeRequest.Merge(m, src)
}
func (m *DeleteCompositeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCompositeRequest.Size(m)
}
func (m *DeleteCompositeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCompositeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCompositeRequest proto.InternalMessageInfo

func (m *DeleteCompositeRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *DeleteCompositeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteCompositeResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteCompositeResponse) Reset()         { *m = DeleteCompositeResponse{} }
func (m *DeleteCompositeResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCompositeResponse) ProtoMessage()    {}
func (*DeleteCompositeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fc8df8568dab43f, []int{9}
}

func (m *DeleteCompositeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCompositeResponse.Unmarshal(m, b)
}
func (m *DeleteCompositeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCompositeResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCompositeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCompositeResponse.Merge(m, src)
}
func (m *DeleteCompositeResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCompositeResponse.Size(m)
}
func (m *DeleteCompositeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCompositeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCompositeResponse proto.InternalMessageInfo

func (m *DeleteCompositeResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAllCompositesRequest)(nil), "protos.GetAllCompositesRequest")
	proto.RegisterType((*GetAllCompositesResponse)(nil), "protos.GetAllCompositesResponse")
	proto.RegisterType((*GetCompositeRequest)(nil), "protos.GetCompositeRequest")
	proto.RegisterType((*GetCompositeResponse)(nil), "protos.GetCompositeResponse")
	proto.RegisterType((*CreateCompositeRequest)(nil), "protos.CreateCompositeRequest")
	proto.RegisterType((*CreateCompositeResponse)(nil), "protos.CreateCompositeResponse")
	proto.RegisterType((*UpdateCompositeRequest)(nil), "protos.UpdateCompositeRequest")
	proto.RegisterType((*UpdateCompositeResponse)(nil), "protos.UpdateCompositeResponse")
	proto.RegisterType((*DeleteCompositeRequest)(nil), "protos.DeleteCompositeRequest")
	proto.RegisterType((*DeleteCompositeResponse)(nil), "protos.DeleteCompositeResponse")
}

func init() { proto.RegisterFile("ps_composite.proto", fileDescriptor_0fc8df8568dab43f) }

var fileDescriptor_0fc8df8568dab43f = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0xbf, 0x8f, 0x4a, 0x6f, 0xc1, 0x45, 0xaa, 0x69, 0xa8, 0x0a, 0x65, 0x56, 0xd9,
	0x98, 0x40, 0x2d, 0xdd, 0xd7, 0x0a, 0x45, 0x10, 0x85, 0x11, 0x37, 0x6e, 0x4a, 0xfe, 0x0c, 0x4d,
	0x20, 0xe9, 0x8c, 0xb9, 0x33, 0xe0, 0xce, 0xad, 0x3b, 0x5f, 0xd3, 0xc7, 0x90, 0x64, 0xd2, 0xd6,
	0x36, 0x08, 0xd6, 0x8a, 0xab, 0x59, 0x9c, 0x73, 0xcf, 0xf9, 0xdd, 0xdb, 0x12, 0x30, 0x05, 0xce,
	0x42, 0x9e, 0x09, 0x8e, 0x89, 0x64, 0xae, 0xc8, 0xb9, 0xe4, 0x66, 0xb3, 0x7c, 0xb0, 0x77, 0x12,
	0xf2, 0x2c, 0xe3, 0x0b, 0x4f, 0x5b, 0x32, 0xbe, 0x98, 0xf9, 0x4a, 0xc6, 0xda, 0xd4, 0x3b, 0xab,
	0x89, 0x28, 0x7d, 0xa9, 0xb0, 0x92, 0x4f, 0xb9, 0x90, 0x58, 0x88, 0xc5, 0xbb, 0xdd, 0x40, 0x26,
	0xd0, 0x9d, 0x32, 0x39, 0x4e, 0xd3, 0xc9, 0x52, 0x40, 0xca, 0x9e, 0x14, 0x43, 0x69, 0x3a, 0xf0,
	0xbf, 0x68, 0xb1, 0xdf, 0x6e, 0xfb, 0x86, 0xd3, 0x1e, 0x74, 0xf4, 0x04, 0xba, 0xba, 0xc4, 0x1d,
	0x2b, 0x19, 0xd3, 0xd2, 0x41, 0x28, 0xd8, 0xf5, 0x10, 0x14, 0x7c, 0x81, 0xcc, 0x1c, 0x01, 0xac,
	0x3a, 0xd1, 0x36, 0xfa, 0xff, 0x9c, 0xf6, 0xc0, 0x5a, 0x46, 0x15, 0x48, 0xee, 0x6a, 0x88, 0x7e,
	0x72, 0x92, 0x3b, 0xe8, 0x4c, 0x99, 0x5c, 0x6b, 0xbb, 0x42, 0x99, 0x87, 0xd0, 0x48, 0x22, 0xdb,
	0xe8, 0x1b, 0x4e, 0x8b, 0x36, 0x92, 0x88, 0xdc, 0xc0, 0xd1, 0x66, 0x60, 0x05, 0x38, 0x84, 0xd6,
	0xaa, 0xb6, 0xb4, 0x7f, 0xcd, 0xb7, 0x36, 0x92, 0x67, 0xb0, 0x26, 0x39, 0xf3, 0x25, 0xdb, 0x83,
	0x70, 0xa3, 0xb9, 0xf1, 0xdd, 0xe6, 0x17, 0xe8, 0xd6, 0x9a, 0xf7, 0x59, 0xc5, 0x74, 0xa1, 0xa9,
	0xff, 0x30, 0xf6, 0xfb, 0x41, 0x39, 0x73, 0xbc, 0x85, 0x7c, 0x5f, 0xaa, 0xb4, 0x72, 0x91, 0x57,
	0x03, 0xac, 0x07, 0x11, 0xed, 0xb7, 0xfb, 0xd6, 0xaf, 0xf3, 0xf3, 0x5b, 0xd4, 0x48, 0xfe, 0xf4,
	0x16, 0x14, 0xac, 0x2b, 0x96, 0xb2, 0xdf, 0x3c, 0x05, 0xb9, 0x86, 0x6e, 0x2d, 0xb3, 0x5a, 0x6a,
	0x47, 0xbc, 0xcb, 0xd1, 0xe3, 0x70, 0x9e, 0xc8, 0x58, 0x05, 0x85, 0xee, 0x05, 0xbe, 0x0c, 0xe3,
	0x90, 0xe7, 0xc2, 0x13, 0xa9, 0xca, 0x02, 0x96, 0x9f, 0x63, 0x18, 0xb3, 0xcc, 0x47, 0x2f, 0x50,
	0x49, 0x1a, 0x79, 0x73, 0xee, 0xe9, 0xb4, 0x40, 0x7f, 0x77, 0x2e, 0x3e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xba, 0xa3, 0x39, 0xce, 0x94, 0x04, 0x00, 0x00,
}
