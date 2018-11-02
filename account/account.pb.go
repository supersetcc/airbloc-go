// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account.proto

package account

import (
	fmt "fmt"
	common "github.com/airbloc/airbloc-go/common"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AccountCreateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountCreateRequest) Reset()         { *m = AccountCreateRequest{} }
func (m *AccountCreateRequest) String() string { return proto.CompactTextString(m) }
func (*AccountCreateRequest) ProtoMessage()    {}
func (*AccountCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{0}
}

func (m *AccountCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCreateRequest.Unmarshal(m, b)
}
func (m *AccountCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCreateRequest.Marshal(b, m, deterministic)
}
func (m *AccountCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCreateRequest.Merge(m, src)
}
func (m *AccountCreateRequest) XXX_Size() int {
	return xxx_messageInfo_AccountCreateRequest.Size(m)
}
func (m *AccountCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCreateRequest proto.InternalMessageInfo

type AccountCreateResponse struct {
	AccountId            *common.Hash `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AccountCreateResponse) Reset()         { *m = AccountCreateResponse{} }
func (m *AccountCreateResponse) String() string { return proto.CompactTextString(m) }
func (*AccountCreateResponse) ProtoMessage()    {}
func (*AccountCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{1}
}

func (m *AccountCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCreateResponse.Unmarshal(m, b)
}
func (m *AccountCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCreateResponse.Marshal(b, m, deterministic)
}
func (m *AccountCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCreateResponse.Merge(m, src)
}
func (m *AccountCreateResponse) XXX_Size() int {
	return xxx_messageInfo_AccountCreateResponse.Size(m)
}
func (m *AccountCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCreateResponse proto.InternalMessageInfo

func (m *AccountCreateResponse) GetAccountId() *common.Hash {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountCreateRequest)(nil), "airbloc.collections.AccountCreateRequest")
	proto.RegisterType((*AccountCreateResponse)(nil), "airbloc.collections.AccountCreateResponse")
}

func init() { proto.RegisterFile("proto/account.proto", fileDescriptor_477cbf5ae5b46edf) }

var fileDescriptor_477cbf5ae5b46edf = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x03, 0xf3, 0x84, 0x84, 0x13, 0x33,
	0x8b, 0x92, 0x72, 0xf2, 0x93, 0xf5, 0x92, 0xf3, 0x73, 0x72, 0x52, 0x93, 0x4b, 0x32, 0xf3, 0xf3,
	0x8a, 0xa5, 0x04, 0x21, 0x2a, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x21, 0xea, 0x94, 0xc4, 0xb8, 0x44,
	0x1c, 0x21, 0x1a, 0x9d, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x94, 0xbc, 0xb9, 0x44, 0xd1, 0xc4, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x8c, 0xb8, 0x38,
	0xa1, 0x36, 0x79, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xe8, 0x21, 0x2c, 0xcb,
	0xcd, 0xcd, 0xcf, 0xd3, 0xf3, 0x48, 0x2c, 0xce, 0x08, 0x42, 0x28, 0x33, 0xca, 0xe2, 0x62, 0x87,
	0x1a, 0x26, 0x14, 0xcf, 0xc5, 0x06, 0x31, 0x50, 0x48, 0x53, 0x0f, 0x8b, 0x13, 0xf5, 0xb0, 0x39,
	0x46, 0x4a, 0x8b, 0x18, 0xa5, 0x10, 0xf7, 0x39, 0xa9, 0x47, 0xa9, 0xa6, 0x67, 0x96, 0x64, 0x94,
	0x26, 0x81, 0x1c, 0xa3, 0x0f, 0xd5, 0x07, 0xa3, 0x75, 0xd3, 0xe1, 0xe1, 0x94, 0xc4, 0x06, 0x0e,
	0x00, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xd0, 0x67, 0x72, 0x3f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	Create(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*AccountCreateResponse, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Create(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*AccountCreateResponse, error) {
	out := new(AccountCreateResponse)
	err := c.cc.Invoke(ctx, "/airbloc.collections.Account/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	Create(context.Context, *AccountCreateRequest) (*AccountCreateResponse, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.collections.Account/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Create(ctx, req.(*AccountCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.collections.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Account_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account.proto",
}
