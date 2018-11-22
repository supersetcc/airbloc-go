// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/data.proto

package api // import "github.com/airbloc/airbloc-go/data/datamanager/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import api "github.com/airbloc/airbloc-go/common/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type BatchRequest struct {
	BatchId              string   `protobuf:"bytes,1,opt,name=batchId,proto3" json:"batchId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchRequest) Reset()         { *m = BatchRequest{} }
func (m *BatchRequest) String() string { return proto.CompactTextString(m) }
func (*BatchRequest) ProtoMessage()    {}
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_7ff67542866a67ca, []int{0}
}
func (m *BatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchRequest.Unmarshal(m, b)
}
func (m *BatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchRequest.Marshal(b, m, deterministic)
}
func (dst *BatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchRequest.Merge(dst, src)
}
func (m *BatchRequest) XXX_Size() int {
	return xxx_messageInfo_BatchRequest.Size(m)
}
func (m *BatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchRequest proto.InternalMessageInfo

func (m *BatchRequest) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

type DataId struct {
	DataId               string   `protobuf:"bytes,1,opt,name=dataId,proto3" json:"dataId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataId) Reset()         { *m = DataId{} }
func (m *DataId) String() string { return proto.CompactTextString(m) }
func (*DataId) ProtoMessage()    {}
func (*DataId) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_7ff67542866a67ca, []int{1}
}
func (m *DataId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataId.Unmarshal(m, b)
}
func (m *DataId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataId.Marshal(b, m, deterministic)
}
func (dst *DataId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataId.Merge(dst, src)
}
func (m *DataId) XXX_Size() int {
	return xxx_messageInfo_DataId.Size(m)
}
func (m *DataId) XXX_DiscardUnknown() {
	xxx_messageInfo_DataId.DiscardUnknown(m)
}

var xxx_messageInfo_DataId proto.InternalMessageInfo

func (m *DataId) GetDataId() string {
	if m != nil {
		return m.DataId
	}
	return ""
}

type DataResult struct {
	DataId               string                   `protobuf:"bytes,1,opt,name=dataId,proto3" json:"dataId,omitempty"`
	CollectionId         string                   `protobuf:"bytes,2,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	OwnerUserAid         string                   `protobuf:"bytes,3,opt,name=ownerUserAid,proto3" json:"ownerUserAid,omitempty"`
	Uri                  string                   `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
	RegisteredAt         uint64                   `protobuf:"varint,5,opt,name=registeredAt,proto3" json:"registeredAt,omitempty"`
	Permissions          []*DataResult_Permission `protobuf:"bytes,6,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DataResult) Reset()         { *m = DataResult{} }
func (m *DataResult) String() string { return proto.CompactTextString(m) }
func (*DataResult) ProtoMessage()    {}
func (*DataResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_7ff67542866a67ca, []int{2}
}
func (m *DataResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResult.Unmarshal(m, b)
}
func (m *DataResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResult.Marshal(b, m, deterministic)
}
func (dst *DataResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResult.Merge(dst, src)
}
func (m *DataResult) XXX_Size() int {
	return xxx_messageInfo_DataResult.Size(m)
}
func (m *DataResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResult.DiscardUnknown(m)
}

var xxx_messageInfo_DataResult proto.InternalMessageInfo

func (m *DataResult) GetDataId() string {
	if m != nil {
		return m.DataId
	}
	return ""
}

func (m *DataResult) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *DataResult) GetOwnerUserAid() string {
	if m != nil {
		return m.OwnerUserAid
	}
	return ""
}

func (m *DataResult) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *DataResult) GetRegisteredAt() uint64 {
	if m != nil {
		return m.RegisteredAt
	}
	return 0
}

func (m *DataResult) GetPermissions() []*DataResult_Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type DataResult_Permission struct {
	ConsumerId           string   `protobuf:"bytes,1,opt,name=consumerId,proto3" json:"consumerId,omitempty"`
	AllowAccess          bool     `protobuf:"varint,2,opt,name=allowAccess,proto3" json:"allowAccess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataResult_Permission) Reset()         { *m = DataResult_Permission{} }
func (m *DataResult_Permission) String() string { return proto.CompactTextString(m) }
func (*DataResult_Permission) ProtoMessage()    {}
func (*DataResult_Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_7ff67542866a67ca, []int{2, 0}
}
func (m *DataResult_Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResult_Permission.Unmarshal(m, b)
}
func (m *DataResult_Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResult_Permission.Marshal(b, m, deterministic)
}
func (dst *DataResult_Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResult_Permission.Merge(dst, src)
}
func (m *DataResult_Permission) XXX_Size() int {
	return xxx_messageInfo_DataResult_Permission.Size(m)
}
func (m *DataResult_Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResult_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_DataResult_Permission proto.InternalMessageInfo

func (m *DataResult_Permission) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *DataResult_Permission) GetAllowAccess() bool {
	if m != nil {
		return m.AllowAccess
	}
	return false
}

type BatchGetResult struct {
	Data                 []*DataResult `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BatchGetResult) Reset()         { *m = BatchGetResult{} }
func (m *BatchGetResult) String() string { return proto.CompactTextString(m) }
func (*BatchGetResult) ProtoMessage()    {}
func (*BatchGetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_7ff67542866a67ca, []int{3}
}
func (m *BatchGetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetResult.Unmarshal(m, b)
}
func (m *BatchGetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetResult.Marshal(b, m, deterministic)
}
func (dst *BatchGetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetResult.Merge(dst, src)
}
func (m *BatchGetResult) XXX_Size() int {
	return xxx_messageInfo_BatchGetResult.Size(m)
}
func (m *BatchGetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetResult.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetResult proto.InternalMessageInfo

func (m *BatchGetResult) GetData() []*DataResult {
	if m != nil {
		return m.Data
	}
	return nil
}

type SetDataPermissionRequest struct {
	DataId               string   `protobuf:"bytes,1,opt,name=dataId,proto3" json:"dataId,omitempty"`
	ConsumerId           string   `protobuf:"bytes,2,opt,name=consumerId,proto3" json:"consumerId,omitempty"`
	AllowAccess          bool     `protobuf:"varint,3,opt,name=allowAccess,proto3" json:"allowAccess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDataPermissionRequest) Reset()         { *m = SetDataPermissionRequest{} }
func (m *SetDataPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*SetDataPermissionRequest) ProtoMessage()    {}
func (*SetDataPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_7ff67542866a67ca, []int{4}
}
func (m *SetDataPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDataPermissionRequest.Unmarshal(m, b)
}
func (m *SetDataPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDataPermissionRequest.Marshal(b, m, deterministic)
}
func (dst *SetDataPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDataPermissionRequest.Merge(dst, src)
}
func (m *SetDataPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_SetDataPermissionRequest.Size(m)
}
func (m *SetDataPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDataPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDataPermissionRequest proto.InternalMessageInfo

func (m *SetDataPermissionRequest) GetDataId() string {
	if m != nil {
		return m.DataId
	}
	return ""
}

func (m *SetDataPermissionRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *SetDataPermissionRequest) GetAllowAccess() bool {
	if m != nil {
		return m.AllowAccess
	}
	return false
}

type SetBatchDataPermissionRequest struct {
	BatchId              string   `protobuf:"bytes,1,opt,name=batchId,proto3" json:"batchId,omitempty"`
	ConsumerId           string   `protobuf:"bytes,2,opt,name=consumerId,proto3" json:"consumerId,omitempty"`
	AllowAccess          bool     `protobuf:"varint,3,opt,name=allowAccess,proto3" json:"allowAccess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetBatchDataPermissionRequest) Reset()         { *m = SetBatchDataPermissionRequest{} }
func (m *SetBatchDataPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*SetBatchDataPermissionRequest) ProtoMessage()    {}
func (*SetBatchDataPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_7ff67542866a67ca, []int{5}
}
func (m *SetBatchDataPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetBatchDataPermissionRequest.Unmarshal(m, b)
}
func (m *SetBatchDataPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetBatchDataPermissionRequest.Marshal(b, m, deterministic)
}
func (dst *SetBatchDataPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBatchDataPermissionRequest.Merge(dst, src)
}
func (m *SetBatchDataPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_SetBatchDataPermissionRequest.Size(m)
}
func (m *SetBatchDataPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBatchDataPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetBatchDataPermissionRequest proto.InternalMessageInfo

func (m *SetBatchDataPermissionRequest) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

func (m *SetBatchDataPermissionRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *SetBatchDataPermissionRequest) GetAllowAccess() bool {
	if m != nil {
		return m.AllowAccess
	}
	return false
}

func init() {
	proto.RegisterType((*BatchRequest)(nil), "airbloc.data.BatchRequest")
	proto.RegisterType((*DataId)(nil), "airbloc.data.DataId")
	proto.RegisterType((*DataResult)(nil), "airbloc.data.DataResult")
	proto.RegisterType((*DataResult_Permission)(nil), "airbloc.data.DataResult.Permission")
	proto.RegisterType((*BatchGetResult)(nil), "airbloc.data.BatchGetResult")
	proto.RegisterType((*SetDataPermissionRequest)(nil), "airbloc.data.SetDataPermissionRequest")
	proto.RegisterType((*SetBatchDataPermissionRequest)(nil), "airbloc.data.SetBatchDataPermissionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataClient interface {
	Get(ctx context.Context, in *DataId, opts ...grpc.CallOption) (*DataResult, error)
	BatchGet(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchGetResult, error)
	SetPermission(ctx context.Context, in *SetDataPermissionRequest, opts ...grpc.CallOption) (*api.Result, error)
	SetPermissionBatch(ctx context.Context, in *SetBatchDataPermissionRequest, opts ...grpc.CallOption) (*api.Results, error)
	Delete(ctx context.Context, in *DataId, opts ...grpc.CallOption) (*api.Result, error)
	DeleteBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*api.Results, error)
	Select(ctx context.Context, opts ...grpc.CallOption) (Data_SelectClient, error)
	Release(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*api.Result, error)
}

type dataClient struct {
	cc *grpc.ClientConn
}

func NewDataClient(cc *grpc.ClientConn) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) Get(ctx context.Context, in *DataId, opts ...grpc.CallOption) (*DataResult, error) {
	out := new(DataResult)
	err := c.cc.Invoke(ctx, "/airbloc.data.Data/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) BatchGet(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchGetResult, error) {
	out := new(BatchGetResult)
	err := c.cc.Invoke(ctx, "/airbloc.data.Data/BatchGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) SetPermission(ctx context.Context, in *SetDataPermissionRequest, opts ...grpc.CallOption) (*api.Result, error) {
	out := new(api.Result)
	err := c.cc.Invoke(ctx, "/airbloc.data.Data/SetPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) SetPermissionBatch(ctx context.Context, in *SetBatchDataPermissionRequest, opts ...grpc.CallOption) (*api.Results, error) {
	out := new(api.Results)
	err := c.cc.Invoke(ctx, "/airbloc.data.Data/SetPermissionBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Delete(ctx context.Context, in *DataId, opts ...grpc.CallOption) (*api.Result, error) {
	out := new(api.Result)
	err := c.cc.Invoke(ctx, "/airbloc.data.Data/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) DeleteBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*api.Results, error) {
	out := new(api.Results)
	err := c.cc.Invoke(ctx, "/airbloc.data.Data/DeleteBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Select(ctx context.Context, opts ...grpc.CallOption) (Data_SelectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Data_serviceDesc.Streams[0], "/airbloc.data.Data/Select", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataSelectClient{stream}
	return x, nil
}

type Data_SelectClient interface {
	Send(*DataId) error
	CloseAndRecv() (*BatchRequest, error)
	grpc.ClientStream
}

type dataSelectClient struct {
	grpc.ClientStream
}

func (x *dataSelectClient) Send(m *DataId) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataSelectClient) CloseAndRecv() (*BatchRequest, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BatchRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataClient) Release(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*api.Result, error) {
	out := new(api.Result)
	err := c.cc.Invoke(ctx, "/airbloc.data.Data/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServer is the server API for Data service.
type DataServer interface {
	Get(context.Context, *DataId) (*DataResult, error)
	BatchGet(context.Context, *BatchRequest) (*BatchGetResult, error)
	SetPermission(context.Context, *SetDataPermissionRequest) (*api.Result, error)
	SetPermissionBatch(context.Context, *SetBatchDataPermissionRequest) (*api.Results, error)
	Delete(context.Context, *DataId) (*api.Result, error)
	DeleteBatch(context.Context, *BatchRequest) (*api.Results, error)
	Select(Data_SelectServer) error
	Release(context.Context, *BatchRequest) (*api.Result, error)
}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.data.Data/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Get(ctx, req.(*DataId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_BatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).BatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.data.Data/BatchGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).BatchGet(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_SetPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDataPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).SetPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.data.Data/SetPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).SetPermission(ctx, req.(*SetDataPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_SetPermissionBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBatchDataPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).SetPermissionBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.data.Data/SetPermissionBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).SetPermissionBatch(ctx, req.(*SetBatchDataPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.data.Data/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Delete(ctx, req.(*DataId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_DeleteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).DeleteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.data.Data/DeleteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).DeleteBatch(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Select_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServer).Select(&dataSelectServer{stream})
}

type Data_SelectServer interface {
	SendAndClose(*BatchRequest) error
	Recv() (*DataId, error)
	grpc.ServerStream
}

type dataSelectServer struct {
	grpc.ServerStream
}

func (x *dataSelectServer) SendAndClose(m *BatchRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataSelectServer) Recv() (*DataId, error) {
	m := new(DataId)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Data_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.data.Data/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Release(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.data.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Data_Get_Handler,
		},
		{
			MethodName: "BatchGet",
			Handler:    _Data_BatchGet_Handler,
		},
		{
			MethodName: "SetPermission",
			Handler:    _Data_SetPermission_Handler,
		},
		{
			MethodName: "SetPermissionBatch",
			Handler:    _Data_SetPermissionBatch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Data_Delete_Handler,
		},
		{
			MethodName: "DeleteBatch",
			Handler:    _Data_DeleteBatch_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _Data_Release_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Select",
			Handler:       _Data_Select_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/data.proto",
}

func init() { proto.RegisterFile("proto/data.proto", fileDescriptor_data_7ff67542866a67ca) }

var fileDescriptor_data_7ff67542866a67ca = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xa5, 0x6d, 0xcc, 0xae, 0xb7, 0x55, 0xd6, 0x41, 0xd6, 0x21, 0xa8, 0x84, 0x08, 0x12, 0x50,
	0x53, 0xa8, 0x1f, 0x4f, 0x2a, 0xb4, 0x54, 0x96, 0xbe, 0xa8, 0xa4, 0xf8, 0xb2, 0x6f, 0xd3, 0xe4,
	0xd2, 0x0d, 0x24, 0x99, 0x3a, 0x33, 0x61, 0x11, 0x7f, 0x91, 0x3f, 0xc1, 0x7f, 0x27, 0x33, 0x49,
	0xdb, 0x84, 0x76, 0xea, 0xc3, 0xbe, 0xb4, 0x33, 0xe7, 0x9e, 0xb9, 0xf7, 0xdc, 0x73, 0x33, 0x03,
	0x17, 0x1b, 0xc1, 0x15, 0x1f, 0xa7, 0x4c, 0xb1, 0xc8, 0x2c, 0xc9, 0x88, 0x65, 0x62, 0x95, 0xf3,
	0x24, 0xd2, 0x98, 0xf7, 0xa8, 0x8e, 0xab, 0x5f, 0x1b, 0x94, 0x35, 0x21, 0x08, 0x61, 0x34, 0x63,
	0x2a, 0xb9, 0x89, 0xf1, 0x67, 0x85, 0x52, 0x11, 0x0a, 0x67, 0x2b, 0xbd, 0x5f, 0xa4, 0xb4, 0xe7,
	0xf7, 0xc2, 0xfb, 0xf1, 0x76, 0x1b, 0xf8, 0xe0, 0xce, 0x99, 0x62, 0x8b, 0x94, 0x5c, 0x82, 0x9b,
	0x9a, 0x55, 0x43, 0x69, 0x76, 0xc1, 0xdf, 0x3e, 0x80, 0xa6, 0xc4, 0x28, 0xab, 0x5c, 0xd9, 0x68,
	0x24, 0x80, 0x51, 0xc2, 0xf3, 0x1c, 0x13, 0x95, 0xf1, 0x72, 0x91, 0xd2, 0xbe, 0x89, 0x76, 0x30,
	0xcd, 0xe1, 0xb7, 0x25, 0x8a, 0x1f, 0x12, 0xc5, 0x34, 0x4b, 0xe9, 0xa0, 0xe6, 0xb4, 0x31, 0x72,
	0x01, 0x83, 0x4a, 0x64, 0xd4, 0x31, 0x21, 0xbd, 0xd4, 0xa7, 0x04, 0xae, 0x33, 0xa9, 0x50, 0x60,
	0x3a, 0x55, 0xf4, 0x9e, 0xdf, 0x0b, 0x9d, 0xb8, 0x83, 0x91, 0x2f, 0x30, 0xdc, 0xa0, 0x28, 0x32,
	0x29, 0x33, 0x5e, 0x4a, 0xea, 0xfa, 0x83, 0x70, 0x38, 0x79, 0x11, 0xb5, 0x7d, 0x8a, 0xf6, 0x4d,
	0x44, 0xdf, 0x77, 0xdc, 0xb8, 0x7d, 0xce, 0xfb, 0x0a, 0xb0, 0x0f, 0x91, 0xe7, 0x00, 0x09, 0x2f,
	0x65, 0x55, 0xa0, 0xd8, 0xb5, 0xdb, 0x42, 0x88, 0x0f, 0x43, 0x96, 0xe7, 0xfc, 0x76, 0x9a, 0x24,
	0x28, 0xa5, 0xe9, 0xf8, 0x3c, 0x6e, 0x43, 0xc1, 0x67, 0x78, 0x68, 0xe6, 0x70, 0x85, 0xaa, 0xb1,
	0xef, 0x35, 0x38, 0x5a, 0x0c, 0xed, 0x19, 0x85, 0xd4, 0xa6, 0x30, 0x36, 0xac, 0x40, 0x01, 0x5d,
	0xa2, 0xd2, 0x70, 0x4b, 0x71, 0x33, 0x53, 0xdb, 0x20, 0xba, 0xaa, 0xfb, 0xff, 0x53, 0x3d, 0x38,
	0x54, 0xfd, 0x1b, 0x9e, 0x2d, 0x51, 0x19, 0xe1, 0xc7, 0x4b, 0x5b, 0x3f, 0xa7, 0xbb, 0x17, 0x9f,
	0xfc, 0x71, 0xc0, 0xd1, 0x55, 0xc9, 0x7b, 0x18, 0x5c, 0xa1, 0x22, 0x8f, 0x0f, 0x2d, 0x5a, 0xa4,
	0x9e, 0xd5, 0x38, 0x32, 0x87, 0xf3, 0xad, 0xe5, 0xc4, 0xeb, 0xb2, 0xda, 0x57, 0xc2, 0x7b, 0x7a,
	0x24, 0xb6, 0x1f, 0xd3, 0x37, 0x78, 0xb0, 0x44, 0xd5, 0xfa, 0x16, 0x5e, 0x76, 0xe9, 0xb6, 0xa9,
	0x78, 0x97, 0x3b, 0x5e, 0xc2, 0x8b, 0x82, 0x97, 0x51, 0x93, 0xf0, 0x1a, 0x48, 0x27, 0xa1, 0xa9,
	0x47, 0x5e, 0x1d, 0x64, 0xb5, 0xbb, 0xee, 0x3d, 0x39, 0x9e, 0x5a, 0x92, 0x0f, 0xe0, 0xce, 0x31,
	0x47, 0x85, 0x16, 0xb3, 0x6c, 0x9a, 0x66, 0x30, 0xac, 0xcf, 0xd5, 0x62, 0x4e, 0xb9, 0x65, 0xad,
	0xfd, 0x11, 0xdc, 0x25, 0xea, 0x1b, 0x6e, 0xa9, 0x7d, 0x22, 0x69, 0xd8, 0x23, 0x9f, 0xe0, 0x2c,
	0xc6, 0x1c, 0x99, 0xc4, 0x93, 0xd5, 0x2d, 0x0d, 0xcc, 0xde, 0x5d, 0x4f, 0xd6, 0x99, 0xba, 0xa9,
	0x56, 0x1a, 0x1f, 0x37, 0x9c, 0xed, 0xff, 0x9b, 0x75, 0xfd, 0x66, 0x9a, 0x9f, 0x82, 0x95, 0x6c,
	0x8d, 0x62, 0xcc, 0x36, 0xd9, 0xca, 0x35, 0x6f, 0xe4, 0xdb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0xd2, 0xf2, 0xd3, 0x58, 0x05, 0x00, 0x00,
}
