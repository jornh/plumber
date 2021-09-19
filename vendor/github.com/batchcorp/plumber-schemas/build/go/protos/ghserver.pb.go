// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ghserver.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GithubEvent_Type int32

const (
	GithubEvent_UNSET           GithubEvent_Type = 0
	GithubEvent_AUTH_RESPONSE   GithubEvent_Type = 1
	GithubEvent_INSTALL_CREATED GithubEvent_Type = 2
	GithubEvent_INSTALL_UPDATED GithubEvent_Type = 3
	GithubEvent_INSTALL_DELETED GithubEvent_Type = 4
	GithubEvent_PULL_CREATED    GithubEvent_Type = 5
	GithubEvent_PULL_MERGED     GithubEvent_Type = 6
	// The plumber server has issued a new JWT token that needs to be stored in etcd
	GithubEvent_NEW_JWT_TOKEN GithubEvent_Type = 7
)

var GithubEvent_Type_name = map[int32]string{
	0: "UNSET",
	1: "AUTH_RESPONSE",
	2: "INSTALL_CREATED",
	3: "INSTALL_UPDATED",
	4: "INSTALL_DELETED",
	5: "PULL_CREATED",
	6: "PULL_MERGED",
	7: "NEW_JWT_TOKEN",
}

var GithubEvent_Type_value = map[string]int32{
	"UNSET":           0,
	"AUTH_RESPONSE":   1,
	"INSTALL_CREATED": 2,
	"INSTALL_UPDATED": 3,
	"INSTALL_DELETED": 4,
	"PULL_CREATED":    5,
	"PULL_MERGED":     6,
	"NEW_JWT_TOKEN":   7,
}

func (x GithubEvent_Type) String() string {
	return proto.EnumName(GithubEvent_Type_name, int32(x))
}

func (GithubEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_667a085d04fc3ea5, []int{0, 0}
}

// GithubEvent is sent by batchcorp/github-service and received by Plumber instances to be acted upon
// It is also sent to the frontend to be acted upon
// See the following URL for reference to events we are receiving from github
// https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#
type GithubEvent struct {
	Type GithubEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protos.GithubEvent_Type" json:"type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*GithubEvent_AuthResponse
	//	*GithubEvent_InstallCreated
	//	*GithubEvent_InstallUpdated
	//	*GithubEvent_InstallDeleted
	//	*GithubEvent_PrCreated
	//	*GithubEvent_PrMerged
	//	*GithubEvent_NewJwtToken
	Payload              isGithubEvent_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GithubEvent) Reset()         { *m = GithubEvent{} }
func (m *GithubEvent) String() string { return proto.CompactTextString(m) }
func (*GithubEvent) ProtoMessage()    {}
func (*GithubEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_667a085d04fc3ea5, []int{0}
}

func (m *GithubEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GithubEvent.Unmarshal(m, b)
}
func (m *GithubEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GithubEvent.Marshal(b, m, deterministic)
}
func (m *GithubEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GithubEvent.Merge(m, src)
}
func (m *GithubEvent) XXX_Size() int {
	return xxx_messageInfo_GithubEvent.Size(m)
}
func (m *GithubEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_GithubEvent.DiscardUnknown(m)
}

var xxx_messageInfo_GithubEvent proto.InternalMessageInfo

func (m *GithubEvent) GetType() GithubEvent_Type {
	if m != nil {
		return m.Type
	}
	return GithubEvent_UNSET
}

type isGithubEvent_Payload interface {
	isGithubEvent_Payload()
}

type GithubEvent_AuthResponse struct {
	AuthResponse *AuthResponse `protobuf:"bytes,100,opt,name=auth_response,json=authResponse,proto3,oneof"`
}

type GithubEvent_InstallCreated struct {
	InstallCreated *InstallCreated `protobuf:"bytes,101,opt,name=install_created,json=installCreated,proto3,oneof"`
}

type GithubEvent_InstallUpdated struct {
	InstallUpdated *InstallUpdated `protobuf:"bytes,102,opt,name=install_updated,json=installUpdated,proto3,oneof"`
}

type GithubEvent_InstallDeleted struct {
	InstallDeleted *InstallDeleted `protobuf:"bytes,103,opt,name=install_deleted,json=installDeleted,proto3,oneof"`
}

type GithubEvent_PrCreated struct {
	PrCreated *PullRequestCreated `protobuf:"bytes,104,opt,name=pr_created,json=prCreated,proto3,oneof"`
}

type GithubEvent_PrMerged struct {
	PrMerged *PullRequestMerged `protobuf:"bytes,105,opt,name=pr_merged,json=prMerged,proto3,oneof"`
}

type GithubEvent_NewJwtToken struct {
	NewJwtToken *NewJwtToken `protobuf:"bytes,106,opt,name=new_jwt_token,json=newJwtToken,proto3,oneof"`
}

func (*GithubEvent_AuthResponse) isGithubEvent_Payload() {}

func (*GithubEvent_InstallCreated) isGithubEvent_Payload() {}

func (*GithubEvent_InstallUpdated) isGithubEvent_Payload() {}

func (*GithubEvent_InstallDeleted) isGithubEvent_Payload() {}

func (*GithubEvent_PrCreated) isGithubEvent_Payload() {}

func (*GithubEvent_PrMerged) isGithubEvent_Payload() {}

func (*GithubEvent_NewJwtToken) isGithubEvent_Payload() {}

func (m *GithubEvent) GetPayload() isGithubEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *GithubEvent) GetAuthResponse() *AuthResponse {
	if x, ok := m.GetPayload().(*GithubEvent_AuthResponse); ok {
		return x.AuthResponse
	}
	return nil
}

func (m *GithubEvent) GetInstallCreated() *InstallCreated {
	if x, ok := m.GetPayload().(*GithubEvent_InstallCreated); ok {
		return x.InstallCreated
	}
	return nil
}

func (m *GithubEvent) GetInstallUpdated() *InstallUpdated {
	if x, ok := m.GetPayload().(*GithubEvent_InstallUpdated); ok {
		return x.InstallUpdated
	}
	return nil
}

func (m *GithubEvent) GetInstallDeleted() *InstallDeleted {
	if x, ok := m.GetPayload().(*GithubEvent_InstallDeleted); ok {
		return x.InstallDeleted
	}
	return nil
}

func (m *GithubEvent) GetPrCreated() *PullRequestCreated {
	if x, ok := m.GetPayload().(*GithubEvent_PrCreated); ok {
		return x.PrCreated
	}
	return nil
}

func (m *GithubEvent) GetPrMerged() *PullRequestMerged {
	if x, ok := m.GetPayload().(*GithubEvent_PrMerged); ok {
		return x.PrMerged
	}
	return nil
}

func (m *GithubEvent) GetNewJwtToken() *NewJwtToken {
	if x, ok := m.GetPayload().(*GithubEvent_NewJwtToken); ok {
		return x.NewJwtToken
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GithubEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GithubEvent_AuthResponse)(nil),
		(*GithubEvent_InstallCreated)(nil),
		(*GithubEvent_InstallUpdated)(nil),
		(*GithubEvent_InstallDeleted)(nil),
		(*GithubEvent_PrCreated)(nil),
		(*GithubEvent_PrMerged)(nil),
		(*GithubEvent_NewJwtToken)(nil),
	}
}

// Sent by plumber, received by github-service
type ConnectAuthRequest struct {
	// JWT token returned during install process
	ApiToken             string   `protobuf:"bytes,1,opt,name=api_token,json=apiToken,proto3" json:"api_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectAuthRequest) Reset()         { *m = ConnectAuthRequest{} }
func (m *ConnectAuthRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectAuthRequest) ProtoMessage()    {}
func (*ConnectAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_667a085d04fc3ea5, []int{1}
}

func (m *ConnectAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectAuthRequest.Unmarshal(m, b)
}
func (m *ConnectAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectAuthRequest.Marshal(b, m, deterministic)
}
func (m *ConnectAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectAuthRequest.Merge(m, src)
}
func (m *ConnectAuthRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectAuthRequest.Size(m)
}
func (m *ConnectAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectAuthRequest proto.InternalMessageInfo

func (m *ConnectAuthRequest) GetApiToken() string {
	if m != nil {
		return m.ApiToken
	}
	return ""
}

// AuthResponse is sent by github-service and received by plumber GithubServer.Connect()
type AuthResponse struct {
	Authorized           bool     `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_667a085d04fc3ea5, []int{2}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func (m *AuthResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Sent by github-service and received by plumber
// Sent by plumber, received by UI
type PullRequestCreated struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Repo                 string   `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Number               int32    `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullRequestCreated) Reset()         { *m = PullRequestCreated{} }
func (m *PullRequestCreated) String() string { return proto.CompactTextString(m) }
func (*PullRequestCreated) ProtoMessage()    {}
func (*PullRequestCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_667a085d04fc3ea5, []int{3}
}

func (m *PullRequestCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullRequestCreated.Unmarshal(m, b)
}
func (m *PullRequestCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullRequestCreated.Marshal(b, m, deterministic)
}
func (m *PullRequestCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullRequestCreated.Merge(m, src)
}
func (m *PullRequestCreated) XXX_Size() int {
	return xxx_messageInfo_PullRequestCreated.Size(m)
}
func (m *PullRequestCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_PullRequestCreated.DiscardUnknown(m)
}

var xxx_messageInfo_PullRequestCreated proto.InternalMessageInfo

func (m *PullRequestCreated) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *PullRequestCreated) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *PullRequestCreated) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *PullRequestCreated) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PullRequestCreated) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Sent by github-service and received by plumber
// Sent by plumber, received by UI
type PullRequestMerged struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Repo                 string   `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Number               int32    `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullRequestMerged) Reset()         { *m = PullRequestMerged{} }
func (m *PullRequestMerged) String() string { return proto.CompactTextString(m) }
func (*PullRequestMerged) ProtoMessage()    {}
func (*PullRequestMerged) Descriptor() ([]byte, []int) {
	return fileDescriptor_667a085d04fc3ea5, []int{4}
}

func (m *PullRequestMerged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullRequestMerged.Unmarshal(m, b)
}
func (m *PullRequestMerged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullRequestMerged.Marshal(b, m, deterministic)
}
func (m *PullRequestMerged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullRequestMerged.Merge(m, src)
}
func (m *PullRequestMerged) XXX_Size() int {
	return xxx_messageInfo_PullRequestMerged.Size(m)
}
func (m *PullRequestMerged) XXX_DiscardUnknown() {
	xxx_messageInfo_PullRequestMerged.DiscardUnknown(m)
}

var xxx_messageInfo_PullRequestMerged proto.InternalMessageInfo

func (m *PullRequestMerged) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *PullRequestMerged) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *PullRequestMerged) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *PullRequestMerged) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// Sent by github-service and received by plumber
// Sent by plumber, received by UI
type InstallCreated struct {
	InstallId            int64    `protobuf:"varint,1,opt,name=install_id,json=installId,proto3" json:"install_id,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallCreated) Reset()         { *m = InstallCreated{} }
func (m *InstallCreated) String() string { return proto.CompactTextString(m) }
func (*InstallCreated) ProtoMessage()    {}
func (*InstallCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_667a085d04fc3ea5, []int{5}
}

func (m *InstallCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallCreated.Unmarshal(m, b)
}
func (m *InstallCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallCreated.Marshal(b, m, deterministic)
}
func (m *InstallCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallCreated.Merge(m, src)
}
func (m *InstallCreated) XXX_Size() int {
	return xxx_messageInfo_InstallCreated.Size(m)
}
func (m *InstallCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallCreated.DiscardUnknown(m)
}

var xxx_messageInfo_InstallCreated proto.InternalMessageInfo

func (m *InstallCreated) GetInstallId() int64 {
	if m != nil {
		return m.InstallId
	}
	return 0
}

func (m *InstallCreated) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

// Sent by github-service and received by plumber
// Sent by plumber, received by UI
type InstallUpdated struct {
	InstallId            int64    `protobuf:"varint,1,opt,name=install_id,json=installId,proto3" json:"install_id,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallUpdated) Reset()         { *m = InstallUpdated{} }
func (m *InstallUpdated) String() string { return proto.CompactTextString(m) }
func (*InstallUpdated) ProtoMessage()    {}
func (*InstallUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_667a085d04fc3ea5, []int{6}
}

func (m *InstallUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallUpdated.Unmarshal(m, b)
}
func (m *InstallUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallUpdated.Marshal(b, m, deterministic)
}
func (m *InstallUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallUpdated.Merge(m, src)
}
func (m *InstallUpdated) XXX_Size() int {
	return xxx_messageInfo_InstallUpdated.Size(m)
}
func (m *InstallUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_InstallUpdated proto.InternalMessageInfo

func (m *InstallUpdated) GetInstallId() int64 {
	if m != nil {
		return m.InstallId
	}
	return 0
}

func (m *InstallUpdated) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

// Sent by github-service and received by plumber
// Sent by plumber, received by UI
type InstallDeleted struct {
	InstallId            int64    `protobuf:"varint,1,opt,name=install_id,json=installId,proto3" json:"install_id,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallDeleted) Reset()         { *m = InstallDeleted{} }
func (m *InstallDeleted) String() string { return proto.CompactTextString(m) }
func (*InstallDeleted) ProtoMessage()    {}
func (*InstallDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_667a085d04fc3ea5, []int{7}
}

func (m *InstallDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallDeleted.Unmarshal(m, b)
}
func (m *InstallDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallDeleted.Marshal(b, m, deterministic)
}
func (m *InstallDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallDeleted.Merge(m, src)
}
func (m *InstallDeleted) XXX_Size() int {
	return xxx_messageInfo_InstallDeleted.Size(m)
}
func (m *InstallDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_InstallDeleted proto.InternalMessageInfo

func (m *InstallDeleted) GetInstallId() int64 {
	if m != nil {
		return m.InstallId
	}
	return 0
}

func (m *InstallDeleted) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

type NewJwtToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewJwtToken) Reset()         { *m = NewJwtToken{} }
func (m *NewJwtToken) String() string { return proto.CompactTextString(m) }
func (*NewJwtToken) ProtoMessage()    {}
func (*NewJwtToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_667a085d04fc3ea5, []int{8}
}

func (m *NewJwtToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewJwtToken.Unmarshal(m, b)
}
func (m *NewJwtToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewJwtToken.Marshal(b, m, deterministic)
}
func (m *NewJwtToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewJwtToken.Merge(m, src)
}
func (m *NewJwtToken) XXX_Size() int {
	return xxx_messageInfo_NewJwtToken.Size(m)
}
func (m *NewJwtToken) XXX_DiscardUnknown() {
	xxx_messageInfo_NewJwtToken.DiscardUnknown(m)
}

var xxx_messageInfo_NewJwtToken proto.InternalMessageInfo

func (m *NewJwtToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterEnum("protos.GithubEvent_Type", GithubEvent_Type_name, GithubEvent_Type_value)
	proto.RegisterType((*GithubEvent)(nil), "protos.GithubEvent")
	proto.RegisterType((*ConnectAuthRequest)(nil), "protos.ConnectAuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "protos.AuthResponse")
	proto.RegisterType((*PullRequestCreated)(nil), "protos.PullRequestCreated")
	proto.RegisterType((*PullRequestMerged)(nil), "protos.PullRequestMerged")
	proto.RegisterType((*InstallCreated)(nil), "protos.InstallCreated")
	proto.RegisterType((*InstallUpdated)(nil), "protos.InstallUpdated")
	proto.RegisterType((*InstallDeleted)(nil), "protos.InstallDeleted")
	proto.RegisterType((*NewJwtToken)(nil), "protos.NewJwtToken")
}

func init() { proto.RegisterFile("ghserver.proto", fileDescriptor_667a085d04fc3ea5) }

var fileDescriptor_667a085d04fc3ea5 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x6e, 0xda, 0x4a,
	0x10, 0xc5, 0x01, 0x42, 0x18, 0x08, 0x21, 0x9b, 0x28, 0xf2, 0xcd, 0xd5, 0xbd, 0x42, 0xf4, 0x25,
	0x0f, 0x2d, 0xb4, 0x69, 0x55, 0xb5, 0x4a, 0x5f, 0x48, 0xb0, 0x02, 0x29, 0x21, 0xc8, 0x18, 0x45,
	0xea, 0x8b, 0x65, 0xec, 0x2d, 0x76, 0x6a, 0xbc, 0xdb, 0xf5, 0x3a, 0x28, 0xfd, 0x84, 0xfe, 0x41,
	0xfb, 0xb5, 0x95, 0x77, 0xed, 0x64, 0xa3, 0xf0, 0xd6, 0xf6, 0x89, 0x9d, 0x73, 0xe6, 0x1c, 0x33,
	0xbb, 0x33, 0x03, 0x8d, 0x85, 0x1f, 0x63, 0x76, 0x8b, 0x59, 0x87, 0x32, 0xc2, 0x09, 0xda, 0x14,
	0x3f, 0x71, 0xfb, 0x67, 0x19, 0x6a, 0xe7, 0x01, 0xf7, 0x93, 0xb9, 0x71, 0x8b, 0x23, 0x8e, 0x9e,
	0x43, 0x89, 0xdf, 0x51, 0xac, 0x6b, 0x2d, 0xed, 0xa8, 0x71, 0xac, 0xcb, 0xec, 0xb8, 0xa3, 0xa4,
	0x74, 0xac, 0x3b, 0x8a, 0x4d, 0x91, 0x85, 0x4e, 0x60, 0xdb, 0x49, 0xb8, 0x6f, 0x33, 0x1c, 0x53,
	0x12, 0xc5, 0x58, 0xf7, 0x5a, 0xda, 0x51, 0xed, 0x78, 0x3f, 0x97, 0xf5, 0x12, 0xee, 0x9b, 0x19,
	0x37, 0x28, 0x98, 0x75, 0x47, 0x89, 0x51, 0x0f, 0x76, 0x82, 0x28, 0xe6, 0x4e, 0x18, 0xda, 0x2e,
	0xc3, 0x0e, 0xc7, 0x9e, 0x8e, 0x85, 0xfc, 0x20, 0x97, 0x0f, 0x25, 0x7d, 0x26, 0xd9, 0x41, 0xc1,
	0x6c, 0x04, 0x8f, 0x10, 0xd5, 0x22, 0xa1, 0x9e, 0xb0, 0xf8, 0xbc, 0xd6, 0x62, 0x26, 0x59, 0xc5,
	0x22, 0x43, 0x54, 0x0b, 0x0f, 0x87, 0x38, 0xb5, 0x58, 0xac, 0xb5, 0xe8, 0x4b, 0x56, 0xb1, 0xc8,
	0x10, 0x74, 0x02, 0x40, 0xd9, 0x7d, 0x0d, 0xbe, 0x50, 0x1f, 0xe6, 0xea, 0x49, 0x12, 0x86, 0x26,
	0xfe, 0x9a, 0xe0, 0x98, 0x3f, 0xd4, 0x51, 0xa5, 0x2c, 0x2f, 0xe1, 0x1d, 0x54, 0x29, 0xb3, 0x97,
	0x98, 0x2d, 0xb0, 0xa7, 0x07, 0x42, 0xfb, 0xcf, 0x1a, 0xed, 0xa5, 0x48, 0x18, 0x14, 0xcc, 0x2d,
	0xca, 0xe4, 0x19, 0xbd, 0x87, 0xed, 0x08, 0xaf, 0xec, 0x9b, 0x15, 0xb7, 0x39, 0xf9, 0x82, 0x23,
	0xfd, 0x46, 0xa8, 0xf7, 0x72, 0xf5, 0x18, 0xaf, 0x2e, 0x56, 0xdc, 0x4a, 0xa9, 0x41, 0xc1, 0xac,
	0x45, 0x0f, 0x61, 0xfb, 0x87, 0x06, 0xa5, 0xf4, 0x19, 0x51, 0x15, 0xca, 0xb3, 0xf1, 0xd4, 0xb0,
	0x9a, 0x05, 0xb4, 0x0b, 0xdb, 0xbd, 0x99, 0x35, 0xb0, 0x4d, 0x63, 0x3a, 0xb9, 0x1a, 0x4f, 0x8d,
	0xa6, 0x86, 0xf6, 0x60, 0x67, 0x38, 0x9e, 0x5a, 0xbd, 0xd1, 0xc8, 0x3e, 0x33, 0x8d, 0x9e, 0x65,
	0xf4, 0x9b, 0x1b, 0x2a, 0x38, 0x9b, 0xf4, 0x05, 0x58, 0x54, 0xc1, 0xbe, 0x31, 0x32, 0x52, 0xb0,
	0x84, 0x9a, 0x50, 0x9f, 0xcc, 0x14, 0x6d, 0x19, 0xed, 0x40, 0x4d, 0x20, 0x97, 0x86, 0x79, 0x6e,
	0xf4, 0x9b, 0x9b, 0xe9, 0x47, 0xc7, 0xc6, 0xb5, 0x7d, 0x71, 0x6d, 0xd9, 0xd6, 0xd5, 0x47, 0x63,
	0xdc, 0xac, 0x9c, 0x56, 0xa1, 0x42, 0x9d, 0xbb, 0x90, 0x38, 0x5e, 0xfb, 0x15, 0xa0, 0x33, 0x12,
	0x45, 0xd8, 0xe5, 0xb2, 0x91, 0xc4, 0x4d, 0xa0, 0x7f, 0xa1, 0xea, 0xd0, 0x20, 0xab, 0x39, 0xed,
	0xd3, 0xaa, 0xb9, 0xe5, 0xd0, 0x40, 0x56, 0x36, 0x80, 0xba, 0xda, 0x74, 0xe8, 0x7f, 0x80, 0xb4,
	0xe9, 0x08, 0x0b, 0xbe, 0x61, 0x4f, 0x64, 0x6f, 0x99, 0x0a, 0x82, 0x74, 0xa8, 0x2c, 0x71, 0x1c,
	0x3b, 0x0b, 0xac, 0x6f, 0x08, 0xab, 0x3c, 0x6c, 0x7f, 0xd7, 0x00, 0x3d, 0x7d, 0x3c, 0xb4, 0x0f,
	0x65, 0xb2, 0x8a, 0x30, 0xcb, 0xbe, 0x2c, 0x03, 0x84, 0xa0, 0xc4, 0x30, 0x25, 0x99, 0x87, 0x38,
	0xa3, 0x03, 0xd8, 0x8c, 0x92, 0xe5, 0x1c, 0x33, 0xbd, 0xd8, 0xd2, 0x8e, 0xca, 0x66, 0x16, 0xa1,
	0x26, 0x14, 0x13, 0x16, 0xea, 0x25, 0x91, 0x9a, 0x1e, 0x51, 0x0b, 0x6a, 0x1e, 0x8e, 0x5d, 0x16,
	0x50, 0x1e, 0x90, 0x48, 0x2f, 0x0b, 0x46, 0x85, 0xda, 0x0b, 0xd8, 0x7d, 0xd2, 0x0c, 0x7f, 0xe3,
	0xaf, 0xb4, 0xc7, 0xd0, 0x78, 0x3c, 0x75, 0xe8, 0x3f, 0x80, 0x7c, 0x40, 0x02, 0x79, 0x83, 0x45,
	0xb3, 0x9a, 0x21, 0x43, 0x41, 0x3b, 0xae, 0x4b, 0x92, 0x88, 0xa7, 0xf4, 0x86, 0xa4, 0x33, 0x64,
	0xe8, 0x29, 0x7e, 0xf9, 0xc0, 0xfd, 0x29, 0xbf, 0x7c, 0xfa, 0x7e, 0xcf, 0xef, 0x19, 0xd4, 0x94,
	0x39, 0x49, 0xaf, 0x54, 0xed, 0x2b, 0x19, 0x1c, 0x8f, 0xa0, 0x2e, 0x17, 0xe0, 0x54, 0xac, 0x50,
	0xf4, 0x01, 0x2a, 0x59, 0x5f, 0xa2, 0xfb, 0x39, 0x7f, 0xda, 0xa8, 0x87, 0x7b, 0x6b, 0xb6, 0xe7,
	0x4b, 0xed, 0xf4, 0xed, 0xa7, 0x37, 0x0b, 0x01, 0x74, 0x5c, 0xb2, 0xec, 0xce, 0x1d, 0xee, 0xfa,
	0x2e, 0x61, 0xb4, 0x4b, 0x43, 0xf1, 0x28, 0x2f, 0x62, 0xd7, 0xc7, 0x4b, 0x27, 0xee, 0xce, 0x93,
	0x20, 0xf4, 0xba, 0x0b, 0xd2, 0x95, 0x2e, 0x73, 0xb9, 0xb2, 0x5f, 0xff, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x5d, 0x05, 0x59, 0x7b, 0xcb, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GithubServerClient is the client API for GithubServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GithubServerClient interface {
	Connect(ctx context.Context, in *ConnectAuthRequest, opts ...grpc.CallOption) (GithubServer_ConnectClient, error)
}

type githubServerClient struct {
	cc *grpc.ClientConn
}

func NewGithubServerClient(cc *grpc.ClientConn) GithubServerClient {
	return &githubServerClient{cc}
}

func (c *githubServerClient) Connect(ctx context.Context, in *ConnectAuthRequest, opts ...grpc.CallOption) (GithubServer_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GithubServer_serviceDesc.Streams[0], "/protos.GithubServer/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &githubServerConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GithubServer_ConnectClient interface {
	Recv() (*GithubEvent, error)
	grpc.ClientStream
}

type githubServerConnectClient struct {
	grpc.ClientStream
}

func (x *githubServerConnectClient) Recv() (*GithubEvent, error) {
	m := new(GithubEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GithubServerServer is the server API for GithubServer service.
type GithubServerServer interface {
	Connect(*ConnectAuthRequest, GithubServer_ConnectServer) error
}

// UnimplementedGithubServerServer can be embedded to have forward compatible implementations.
type UnimplementedGithubServerServer struct {
}

func (*UnimplementedGithubServerServer) Connect(req *ConnectAuthRequest, srv GithubServer_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterGithubServerServer(s *grpc.Server, srv GithubServerServer) {
	s.RegisterService(&_GithubServer_serviceDesc, srv)
}

func _GithubServer_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectAuthRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GithubServerServer).Connect(m, &githubServerConnectServer{stream})
}

type GithubServer_ConnectServer interface {
	Send(*GithubEvent) error
	grpc.ServerStream
}

type githubServerConnectServer struct {
	grpc.ServerStream
}

func (x *githubServerConnectServer) Send(m *GithubEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _GithubServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.GithubServer",
	HandlerType: (*GithubServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _GithubServer_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ghserver.proto",
}
