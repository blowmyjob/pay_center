// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Pay.proto

// 定义包名

package rpc_proto

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

type PayRecordCreateReq struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	ProductId            int64    `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	PayAmount            float32  `protobuf:"fixed32,4,opt,name=payAmount,proto3" json:"payAmount,omitempty"`
	PayMethod            string   `protobuf:"bytes,5,opt,name=payMethod,proto3" json:"payMethod,omitempty"`
	AppPayId             string   `protobuf:"bytes,6,opt,name=appPayId,proto3" json:"appPayId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayRecordCreateReq) Reset()         { *m = PayRecordCreateReq{} }
func (m *PayRecordCreateReq) String() string { return proto.CompactTextString(m) }
func (*PayRecordCreateReq) ProtoMessage()    {}
func (*PayRecordCreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9069421e92eda4b8, []int{0}
}

func (m *PayRecordCreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayRecordCreateReq.Unmarshal(m, b)
}
func (m *PayRecordCreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayRecordCreateReq.Marshal(b, m, deterministic)
}
func (m *PayRecordCreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayRecordCreateReq.Merge(m, src)
}
func (m *PayRecordCreateReq) XXX_Size() int {
	return xxx_messageInfo_PayRecordCreateReq.Size(m)
}
func (m *PayRecordCreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PayRecordCreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_PayRecordCreateReq proto.InternalMessageInfo

func (m *PayRecordCreateReq) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *PayRecordCreateReq) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *PayRecordCreateReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PayRecordCreateReq) GetPayAmount() float32 {
	if m != nil {
		return m.PayAmount
	}
	return 0
}

func (m *PayRecordCreateReq) GetPayMethod() string {
	if m != nil {
		return m.PayMethod
	}
	return ""
}

func (m *PayRecordCreateReq) GetAppPayId() string {
	if m != nil {
		return m.AppPayId
	}
	return ""
}

type PayRecordCreateResp struct {
	SeqNo                int64    `protobuf:"varint,1,opt,name=seqNo,proto3" json:"seqNo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayRecordCreateResp) Reset()         { *m = PayRecordCreateResp{} }
func (m *PayRecordCreateResp) String() string { return proto.CompactTextString(m) }
func (*PayRecordCreateResp) ProtoMessage()    {}
func (*PayRecordCreateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9069421e92eda4b8, []int{1}
}

func (m *PayRecordCreateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayRecordCreateResp.Unmarshal(m, b)
}
func (m *PayRecordCreateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayRecordCreateResp.Marshal(b, m, deterministic)
}
func (m *PayRecordCreateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayRecordCreateResp.Merge(m, src)
}
func (m *PayRecordCreateResp) XXX_Size() int {
	return xxx_messageInfo_PayRecordCreateResp.Size(m)
}
func (m *PayRecordCreateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PayRecordCreateResp.DiscardUnknown(m)
}

var xxx_messageInfo_PayRecordCreateResp proto.InternalMessageInfo

func (m *PayRecordCreateResp) GetSeqNo() int64 {
	if m != nil {
		return m.SeqNo
	}
	return 0
}

type PayRecordSelectReq struct {
	SeqNo                int64    `protobuf:"varint,1,opt,name=seqNo,proto3" json:"seqNo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayRecordSelectReq) Reset()         { *m = PayRecordSelectReq{} }
func (m *PayRecordSelectReq) String() string { return proto.CompactTextString(m) }
func (*PayRecordSelectReq) ProtoMessage()    {}
func (*PayRecordSelectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9069421e92eda4b8, []int{2}
}

func (m *PayRecordSelectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayRecordSelectReq.Unmarshal(m, b)
}
func (m *PayRecordSelectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayRecordSelectReq.Marshal(b, m, deterministic)
}
func (m *PayRecordSelectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayRecordSelectReq.Merge(m, src)
}
func (m *PayRecordSelectReq) XXX_Size() int {
	return xxx_messageInfo_PayRecordSelectReq.Size(m)
}
func (m *PayRecordSelectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PayRecordSelectReq.DiscardUnknown(m)
}

var xxx_messageInfo_PayRecordSelectReq proto.InternalMessageInfo

func (m *PayRecordSelectReq) GetSeqNo() int64 {
	if m != nil {
		return m.SeqNo
	}
	return 0
}

type PayRecordSelectResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayRecordSelectResp) Reset()         { *m = PayRecordSelectResp{} }
func (m *PayRecordSelectResp) String() string { return proto.CompactTextString(m) }
func (*PayRecordSelectResp) ProtoMessage()    {}
func (*PayRecordSelectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9069421e92eda4b8, []int{3}
}

func (m *PayRecordSelectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayRecordSelectResp.Unmarshal(m, b)
}
func (m *PayRecordSelectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayRecordSelectResp.Marshal(b, m, deterministic)
}
func (m *PayRecordSelectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayRecordSelectResp.Merge(m, src)
}
func (m *PayRecordSelectResp) XXX_Size() int {
	return xxx_messageInfo_PayRecordSelectResp.Size(m)
}
func (m *PayRecordSelectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PayRecordSelectResp.DiscardUnknown(m)
}

var xxx_messageInfo_PayRecordSelectResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PayRecordCreateReq)(nil), "rpc.proto.PayRecordCreateReq")
	proto.RegisterType((*PayRecordCreateResp)(nil), "rpc.proto.PayRecordCreateResp")
	proto.RegisterType((*PayRecordSelectReq)(nil), "rpc.proto.PayRecordSelectReq")
	proto.RegisterType((*PayRecordSelectResp)(nil), "rpc.proto.PayRecordSelectResp")
}

func init() { proto.RegisterFile("Pay.proto", fileDescriptor_9069421e92eda4b8) }

var fileDescriptor_9069421e92eda4b8 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xde, 0x6c, 0xdd, 0x62, 0xe7, 0x22, 0xc4, 0x1f, 0xc2, 0xa2, 0x52, 0x7a, 0x2a, 0x0a, 0x3d,
	0xe8, 0x13, 0x88, 0xa7, 0x1e, 0x94, 0x92, 0x7d, 0x82, 0xd8, 0x0c, 0x28, 0xa8, 0x99, 0x4d, 0x53,
	0x21, 0xef, 0xe5, 0xc1, 0xc7, 0x93, 0x4d, 0x6a, 0xbb, 0xb0, 0x8b, 0x7b, 0xcb, 0x37, 0xdf, 0xe4,
	0xfb, 0x19, 0xc8, 0x1a, 0xe5, 0x2b, 0xb2, 0xc6, 0x19, 0x9e, 0x59, 0x6a, 0xe3, 0xb3, 0xf8, 0x61,
	0xc0, 0x1b, 0xe5, 0x25, 0xb6, 0xc6, 0xea, 0x47, 0x8b, 0xca, 0xa1, 0xc4, 0x35, 0x3f, 0x83, 0x85,
	0x22, 0xaa, 0xb5, 0x60, 0x39, 0x2b, 0x33, 0x19, 0x01, 0xbf, 0x84, 0x8c, 0xac, 0xd1, 0x7d, 0xeb,
	0x6a, 0x2d, 0xe6, 0x39, 0x2b, 0x13, 0x39, 0x0d, 0xf8, 0x05, 0xa4, 0x7d, 0x87, 0xb6, 0xd6, 0x22,
	0x09, 0xd4, 0x80, 0xc2, 0x2f, 0xe5, 0x1f, 0x3e, 0x4c, 0xff, 0xe9, 0xc4, 0x51, 0xce, 0xca, 0xb9,
	0x9c, 0x06, 0x03, 0xfb, 0x84, 0xee, 0xd5, 0x68, 0xb1, 0x08, 0x6e, 0xd3, 0x80, 0x2f, 0xe1, 0x58,
	0x11, 0x35, 0xca, 0xd7, 0x5a, 0xa4, 0x81, 0x1c, 0x71, 0x71, 0x0b, 0xa7, 0x3b, 0xc9, 0x3b, 0xda,
	0x44, 0xef, 0x70, 0xfd, 0x6c, 0x42, 0xf4, 0x44, 0x46, 0x50, 0xdc, 0x6c, 0xd5, 0x5c, 0xe1, 0x3b,
	0xb6, 0x6e, 0xa8, 0xb9, 0x67, 0xf7, 0x7c, 0x4b, 0xf8, 0x6f, 0xb7, 0xa3, 0xbb, 0x6f, 0x06, 0xd0,
	0x28, 0xbf, 0x42, 0xfb, 0xf5, 0xd6, 0x22, 0x97, 0x70, 0x12, 0x5d, 0xc7, 0x5d, 0x7e, 0x55, 0x8d,
	0x87, 0xad, 0x76, 0x8f, 0xba, 0xbc, 0xfe, 0x8f, 0xee, 0xa8, 0x98, 0x6d, 0x34, 0xa3, 0xe1, 0x01,
	0xcd, 0xb1, 0xc1, 0x7e, 0xcd, 0x29, 0x74, 0x31, 0x7b, 0x49, 0x03, 0x79, 0xff, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x5c, 0x93, 0x93, 0xe8, 0x00, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PayServiceClient is the client API for PayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PayServiceClient interface {
	CreatePayRecord(ctx context.Context, in *PayRecordCreateReq, opts ...grpc.CallOption) (*PayRecordCreateResp, error)
	SelectPayRecord(ctx context.Context, in *PayRecordSelectReq, opts ...grpc.CallOption) (*PayRecordSelectResp, error)
}

type payServiceClient struct {
	cc *grpc.ClientConn
}

func NewPayServiceClient(cc *grpc.ClientConn) PayServiceClient {
	return &payServiceClient{cc}
}

func (c *payServiceClient) CreatePayRecord(ctx context.Context, in *PayRecordCreateReq, opts ...grpc.CallOption) (*PayRecordCreateResp, error) {
	out := new(PayRecordCreateResp)
	err := c.cc.Invoke(ctx, "/rpc.proto.PayService/CreatePayRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServiceClient) SelectPayRecord(ctx context.Context, in *PayRecordSelectReq, opts ...grpc.CallOption) (*PayRecordSelectResp, error) {
	out := new(PayRecordSelectResp)
	err := c.cc.Invoke(ctx, "/rpc.proto.PayService/SelectPayRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayServiceServer is the server API for PayService service.
type PayServiceServer interface {
	CreatePayRecord(context.Context, *PayRecordCreateReq) (*PayRecordCreateResp, error)
	SelectPayRecord(context.Context, *PayRecordSelectReq) (*PayRecordSelectResp, error)
}

// UnimplementedPayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPayServiceServer struct {
}

func (*UnimplementedPayServiceServer) CreatePayRecord(ctx context.Context, req *PayRecordCreateReq) (*PayRecordCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayRecord not implemented")
}
func (*UnimplementedPayServiceServer) SelectPayRecord(ctx context.Context, req *PayRecordSelectReq) (*PayRecordSelectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectPayRecord not implemented")
}

func RegisterPayServiceServer(s *grpc.Server, srv PayServiceServer) {
	s.RegisterService(&_PayService_serviceDesc, srv)
}

func _PayService_CreatePayRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRecordCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServiceServer).CreatePayRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.proto.PayService/CreatePayRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServiceServer).CreatePayRecord(ctx, req.(*PayRecordCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayService_SelectPayRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRecordSelectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServiceServer).SelectPayRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.proto.PayService/SelectPayRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServiceServer).SelectPayRecord(ctx, req.(*PayRecordSelectReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.proto.PayService",
	HandlerType: (*PayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayRecord",
			Handler:    _PayService_CreatePayRecord_Handler,
		},
		{
			MethodName: "SelectPayRecord",
			Handler:    _PayService_SelectPayRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Pay.proto",
}
