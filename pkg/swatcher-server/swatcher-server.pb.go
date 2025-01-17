// Code generated by protoc-gen-go. DO NOT EDIT.
// source: swatcher-server.proto

package swatcher_server

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

type ServerUpdate_UpdateType int32

const (
	ServerUpdate_ADD      ServerUpdate_UpdateType = 0
	ServerUpdate_DOWNLOAD ServerUpdate_UpdateType = 1
	ServerUpdate_DELETE   ServerUpdate_UpdateType = 2
)

var ServerUpdate_UpdateType_name = map[int32]string{
	0: "ADD",
	1: "DOWNLOAD",
	2: "DELETE",
}

var ServerUpdate_UpdateType_value = map[string]int32{
	"ADD":      0,
	"DOWNLOAD": 1,
	"DELETE":   2,
}

func (x ServerUpdate_UpdateType) String() string {
	return proto.EnumName(ServerUpdate_UpdateType_name, int32(x))
}

func (ServerUpdate_UpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67ed6a4e241b5717, []int{1, 0}
}

type SendUpdatesRequest_UpdateType int32

const (
	SendUpdatesRequest_PROGRESS SendUpdatesRequest_UpdateType = 0
	SendUpdatesRequest_STATUS   SendUpdatesRequest_UpdateType = 1
)

var SendUpdatesRequest_UpdateType_name = map[int32]string{
	0: "PROGRESS",
	1: "STATUS",
}

var SendUpdatesRequest_UpdateType_value = map[string]int32{
	"PROGRESS": 0,
	"STATUS":   1,
}

func (x SendUpdatesRequest_UpdateType) String() string {
	return proto.EnumName(SendUpdatesRequest_UpdateType_name, int32(x))
}

func (SendUpdatesRequest_UpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67ed6a4e241b5717, []int{3, 0}
}

type GetUpdatesRequest struct {
	ClientId             int64    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Rev                  int64    `protobuf:"varint,2,opt,name=rev,proto3" json:"rev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUpdatesRequest) Reset()         { *m = GetUpdatesRequest{} }
func (m *GetUpdatesRequest) String() string { return proto.CompactTextString(m) }
func (*GetUpdatesRequest) ProtoMessage()    {}
func (*GetUpdatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67ed6a4e241b5717, []int{0}
}

func (m *GetUpdatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUpdatesRequest.Unmarshal(m, b)
}
func (m *GetUpdatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUpdatesRequest.Marshal(b, m, deterministic)
}
func (m *GetUpdatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpdatesRequest.Merge(m, src)
}
func (m *GetUpdatesRequest) XXX_Size() int {
	return xxx_messageInfo_GetUpdatesRequest.Size(m)
}
func (m *GetUpdatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpdatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpdatesRequest proto.InternalMessageInfo

func (m *GetUpdatesRequest) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *GetUpdatesRequest) GetRev() int64 {
	if m != nil {
		return m.Rev
	}
	return 0
}

type ServerUpdate struct {
	Type                 ServerUpdate_UpdateType `protobuf:"varint,1,opt,name=type,proto3,enum=ServerUpdate_UpdateType" json:"type,omitempty"`
	Data                 string                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ServerUpdate) Reset()         { *m = ServerUpdate{} }
func (m *ServerUpdate) String() string { return proto.CompactTextString(m) }
func (*ServerUpdate) ProtoMessage()    {}
func (*ServerUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_67ed6a4e241b5717, []int{1}
}

func (m *ServerUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerUpdate.Unmarshal(m, b)
}
func (m *ServerUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerUpdate.Marshal(b, m, deterministic)
}
func (m *ServerUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerUpdate.Merge(m, src)
}
func (m *ServerUpdate) XXX_Size() int {
	return xxx_messageInfo_ServerUpdate.Size(m)
}
func (m *ServerUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ServerUpdate proto.InternalMessageInfo

func (m *ServerUpdate) GetType() ServerUpdate_UpdateType {
	if m != nil {
		return m.Type
	}
	return ServerUpdate_ADD
}

func (m *ServerUpdate) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type GetUpdatesResponse struct {
	Updates              []*ServerUpdate `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetUpdatesResponse) Reset()         { *m = GetUpdatesResponse{} }
func (m *GetUpdatesResponse) String() string { return proto.CompactTextString(m) }
func (*GetUpdatesResponse) ProtoMessage()    {}
func (*GetUpdatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67ed6a4e241b5717, []int{2}
}

func (m *GetUpdatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUpdatesResponse.Unmarshal(m, b)
}
func (m *GetUpdatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUpdatesResponse.Marshal(b, m, deterministic)
}
func (m *GetUpdatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpdatesResponse.Merge(m, src)
}
func (m *GetUpdatesResponse) XXX_Size() int {
	return xxx_messageInfo_GetUpdatesResponse.Size(m)
}
func (m *GetUpdatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpdatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpdatesResponse proto.InternalMessageInfo

func (m *GetUpdatesResponse) GetUpdates() []*ServerUpdate {
	if m != nil {
		return m.Updates
	}
	return nil
}

type SendUpdatesRequest struct {
	ClientId             int64                         `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Type                 SendUpdatesRequest_UpdateType `protobuf:"varint,2,opt,name=type,proto3,enum=SendUpdatesRequest_UpdateType" json:"type,omitempty"`
	Data                 string                        `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SendUpdatesRequest) Reset()         { *m = SendUpdatesRequest{} }
func (m *SendUpdatesRequest) String() string { return proto.CompactTextString(m) }
func (*SendUpdatesRequest) ProtoMessage()    {}
func (*SendUpdatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67ed6a4e241b5717, []int{3}
}

func (m *SendUpdatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendUpdatesRequest.Unmarshal(m, b)
}
func (m *SendUpdatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendUpdatesRequest.Marshal(b, m, deterministic)
}
func (m *SendUpdatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendUpdatesRequest.Merge(m, src)
}
func (m *SendUpdatesRequest) XXX_Size() int {
	return xxx_messageInfo_SendUpdatesRequest.Size(m)
}
func (m *SendUpdatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendUpdatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendUpdatesRequest proto.InternalMessageInfo

func (m *SendUpdatesRequest) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *SendUpdatesRequest) GetType() SendUpdatesRequest_UpdateType {
	if m != nil {
		return m.Type
	}
	return SendUpdatesRequest_PROGRESS
}

func (m *SendUpdatesRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// TODO: if some updates applied and some not - which answer is it?
type SendUpdatesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendUpdatesResponse) Reset()         { *m = SendUpdatesResponse{} }
func (m *SendUpdatesResponse) String() string { return proto.CompactTextString(m) }
func (*SendUpdatesResponse) ProtoMessage()    {}
func (*SendUpdatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67ed6a4e241b5717, []int{4}
}

func (m *SendUpdatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendUpdatesResponse.Unmarshal(m, b)
}
func (m *SendUpdatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendUpdatesResponse.Marshal(b, m, deterministic)
}
func (m *SendUpdatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendUpdatesResponse.Merge(m, src)
}
func (m *SendUpdatesResponse) XXX_Size() int {
	return xxx_messageInfo_SendUpdatesResponse.Size(m)
}
func (m *SendUpdatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendUpdatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendUpdatesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ServerUpdate_UpdateType", ServerUpdate_UpdateType_name, ServerUpdate_UpdateType_value)
	proto.RegisterEnum("SendUpdatesRequest_UpdateType", SendUpdatesRequest_UpdateType_name, SendUpdatesRequest_UpdateType_value)
	proto.RegisterType((*GetUpdatesRequest)(nil), "GetUpdatesRequest")
	proto.RegisterType((*ServerUpdate)(nil), "ServerUpdate")
	proto.RegisterType((*GetUpdatesResponse)(nil), "GetUpdatesResponse")
	proto.RegisterType((*SendUpdatesRequest)(nil), "SendUpdatesRequest")
	proto.RegisterType((*SendUpdatesResponse)(nil), "SendUpdatesResponse")
}

func init() { proto.RegisterFile("swatcher-server.proto", fileDescriptor_67ed6a4e241b5717) }

var fileDescriptor_67ed6a4e241b5717 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x77, 0x5c, 0xf1, 0xc7, 0xd3, 0x64, 0x7b, 0x26, 0x2c, 0x06, 0x21, 0x73, 0x28, 0x0f,
	0xb5, 0xc1, 0x76, 0xe8, 0x52, 0x07, 0x63, 0x17, 0x09, 0x24, 0x63, 0x66, 0xa5, 0x63, 0x6c, 0xee,
	0x83, 0x84, 0xd0, 0x6d, 0x77, 0x34, 0x3c, 0x76, 0xea, 0xef, 0xe8, 0x3f, 0x0d, 0x67, 0x0d, 0x57,
	0xb4, 0x43, 0xa7, 0x79, 0xbc, 0xdf, 0x9f, 0xef, 0x1b, 0x68, 0xa5, 0x1f, 0xa1, 0x1a, 0xbf, 0x52,
	0x72, 0x91, 0x52, 0xb2, 0xa0, 0xc4, 0x89, 0x93, 0x99, 0x9a, 0xf1, 0x3b, 0x38, 0xec, 0x93, 0x1a,
	0xc5, 0x51, 0xa8, 0x28, 0x15, 0xf4, 0x3e, 0xa7, 0x54, 0xe1, 0x31, 0x54, 0xc7, 0x6f, 0x13, 0x9a,
	0xaa, 0xe7, 0x49, 0x64, 0xb3, 0x0e, 0xeb, 0x9a, 0xa2, 0x92, 0x39, 0xee, 0x23, 0xb4, 0xc0, 0x4c,
	0x68, 0x61, 0x17, 0xb4, 0x7b, 0x65, 0xf2, 0x4f, 0x06, 0x75, 0xa9, 0x9b, 0x66, 0x7d, 0xf0, 0x1c,
	0x8a, 0x6a, 0x19, 0x93, 0x2e, 0x6d, 0xb8, 0xb6, 0x93, 0x0f, 0x3a, 0xd9, 0x13, 0x2c, 0x63, 0x12,
	0x3a, 0x0b, 0x11, 0x8a, 0x51, 0xa8, 0x42, 0xdd, 0xb1, 0x2a, 0xb4, 0xcd, 0x2f, 0x01, 0x36, 0x79,
	0x58, 0x06, 0xb3, 0xe7, 0x79, 0x96, 0x81, 0x75, 0xa8, 0x78, 0xc3, 0xa7, 0x87, 0xc1, 0xb0, 0xe7,
	0x59, 0x0c, 0x01, 0x4a, 0x9e, 0x3f, 0xf0, 0x03, 0xdf, 0x2a, 0xf0, 0x5b, 0xc0, 0x3c, 0x47, 0x1a,
	0xcf, 0xa6, 0x29, 0xe1, 0x19, 0x94, 0xe7, 0x99, 0xcb, 0x66, 0x1d, 0xb3, 0x5b, 0x73, 0x0f, 0xb6,
	0x76, 0x11, 0xbf, 0x51, 0xfe, 0xcd, 0x00, 0x25, 0x4d, 0xa3, 0xff, 0x08, 0xe1, 0xae, 0x29, 0x0b,
	0x9a, 0xf2, 0xc4, 0xd9, 0xad, 0xff, 0x9b, 0xd5, 0xcc, 0xb1, 0x9e, 0x6e, 0xb1, 0xd6, 0xa1, 0xf2,
	0x28, 0x86, 0x7d, 0xe1, 0x4b, 0x69, 0x19, 0x2b, 0x44, 0x19, 0xf4, 0x82, 0x91, 0xb4, 0x18, 0x6f,
	0x41, 0x73, 0x6b, 0x44, 0xc6, 0xe8, 0x7e, 0x31, 0x68, 0xc8, 0xf5, 0x6d, 0x33, 0x38, 0xbc, 0x06,
	0xd8, 0x88, 0x81, 0xe8, 0xec, 0x5c, 0xb8, 0xdd, 0x74, 0x76, 0xd5, 0xe2, 0x06, 0xde, 0x40, 0x2d,
	0x37, 0x02, 0x9b, 0x7b, 0x98, 0xda, 0x47, 0xce, 0x9e, 0x2d, 0xb8, 0xd1, 0x65, 0x2f, 0x25, 0xfd,
	0xa5, 0xae, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xc2, 0x74, 0x91, 0x6b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SwatcherServerClient is the client API for SwatcherServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SwatcherServerClient interface {
	GetUpdates(ctx context.Context, in *GetUpdatesRequest, opts ...grpc.CallOption) (*GetUpdatesResponse, error)
	SendUpdates(ctx context.Context, opts ...grpc.CallOption) (SwatcherServer_SendUpdatesClient, error)
}

type swatcherServerClient struct {
	cc *grpc.ClientConn
}

func NewSwatcherServerClient(cc *grpc.ClientConn) SwatcherServerClient {
	return &swatcherServerClient{cc}
}

func (c *swatcherServerClient) GetUpdates(ctx context.Context, in *GetUpdatesRequest, opts ...grpc.CallOption) (*GetUpdatesResponse, error) {
	out := new(GetUpdatesResponse)
	err := c.cc.Invoke(ctx, "/SwatcherServer/GetUpdates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swatcherServerClient) SendUpdates(ctx context.Context, opts ...grpc.CallOption) (SwatcherServer_SendUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SwatcherServer_serviceDesc.Streams[0], "/SwatcherServer/SendUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &swatcherServerSendUpdatesClient{stream}
	return x, nil
}

type SwatcherServer_SendUpdatesClient interface {
	Send(*SendUpdatesRequest) error
	CloseAndRecv() (*SendUpdatesResponse, error)
	grpc.ClientStream
}

type swatcherServerSendUpdatesClient struct {
	grpc.ClientStream
}

func (x *swatcherServerSendUpdatesClient) Send(m *SendUpdatesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *swatcherServerSendUpdatesClient) CloseAndRecv() (*SendUpdatesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendUpdatesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SwatcherServerServer is the server API for SwatcherServer service.
type SwatcherServerServer interface {
	GetUpdates(context.Context, *GetUpdatesRequest) (*GetUpdatesResponse, error)
	SendUpdates(SwatcherServer_SendUpdatesServer) error
}

// UnimplementedSwatcherServerServer can be embedded to have forward compatible implementations.
type UnimplementedSwatcherServerServer struct {
}

func (*UnimplementedSwatcherServerServer) GetUpdates(ctx context.Context, req *GetUpdatesRequest) (*GetUpdatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpdates not implemented")
}
func (*UnimplementedSwatcherServerServer) SendUpdates(srv SwatcherServer_SendUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method SendUpdates not implemented")
}

func RegisterSwatcherServerServer(s *grpc.Server, srv SwatcherServerServer) {
	s.RegisterService(&_SwatcherServer_serviceDesc, srv)
}

func _SwatcherServer_GetUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwatcherServerServer).GetUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SwatcherServer/GetUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwatcherServerServer).GetUpdates(ctx, req.(*GetUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwatcherServer_SendUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SwatcherServerServer).SendUpdates(&swatcherServerSendUpdatesServer{stream})
}

type SwatcherServer_SendUpdatesServer interface {
	SendAndClose(*SendUpdatesResponse) error
	Recv() (*SendUpdatesRequest, error)
	grpc.ServerStream
}

type swatcherServerSendUpdatesServer struct {
	grpc.ServerStream
}

func (x *swatcherServerSendUpdatesServer) SendAndClose(m *SendUpdatesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *swatcherServerSendUpdatesServer) Recv() (*SendUpdatesRequest, error) {
	m := new(SendUpdatesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SwatcherServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SwatcherServer",
	HandlerType: (*SwatcherServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpdates",
			Handler:    _SwatcherServer_GetUpdates_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendUpdates",
			Handler:       _SwatcherServer_SendUpdates_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "swatcher-server.proto",
}
