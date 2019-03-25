// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infrpc.proto

package infrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type WritePointsRequest struct {
	ShardID              uint64   `protobuf:"varint,1,opt,name=shardID,proto3" json:"shardID,omitempty"`
	Points               [][]byte `protobuf:"bytes,2,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WritePointsRequest) Reset()         { *m = WritePointsRequest{} }
func (m *WritePointsRequest) String() string { return proto.CompactTextString(m) }
func (*WritePointsRequest) ProtoMessage()    {}
func (*WritePointsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92fc5e3694af4d57, []int{0}
}

func (m *WritePointsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WritePointsRequest.Unmarshal(m, b)
}
func (m *WritePointsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WritePointsRequest.Marshal(b, m, deterministic)
}
func (m *WritePointsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WritePointsRequest.Merge(m, src)
}
func (m *WritePointsRequest) XXX_Size() int {
	return xxx_messageInfo_WritePointsRequest.Size(m)
}
func (m *WritePointsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WritePointsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WritePointsRequest proto.InternalMessageInfo

func (m *WritePointsRequest) GetShardID() uint64 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *WritePointsRequest) GetPoints() [][]byte {
	if m != nil {
		return m.Points
	}
	return nil
}

type WritePointsReply struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WritePointsReply) Reset()         { *m = WritePointsReply{} }
func (m *WritePointsReply) String() string { return proto.CompactTextString(m) }
func (*WritePointsReply) ProtoMessage()    {}
func (*WritePointsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_92fc5e3694af4d57, []int{1}
}

func (m *WritePointsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WritePointsReply.Unmarshal(m, b)
}
func (m *WritePointsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WritePointsReply.Marshal(b, m, deterministic)
}
func (m *WritePointsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WritePointsReply.Merge(m, src)
}
func (m *WritePointsReply) XXX_Size() int {
	return xxx_messageInfo_WritePointsReply.Size(m)
}
func (m *WritePointsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WritePointsReply.DiscardUnknown(m)
}

var xxx_messageInfo_WritePointsReply proto.InternalMessageInfo

func (m *WritePointsReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*WritePointsRequest)(nil), "infrpc.WritePointsRequest")
	proto.RegisterType((*WritePointsReply)(nil), "infrpc.WritePointsReply")
}

func init() { proto.RegisterFile("infrpc.proto", fileDescriptor_92fc5e3694af4d57) }

var fileDescriptor_92fc5e3694af4d57 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xcc, 0x4b, 0x2b,
	0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xdc, 0xb8, 0x84,
	0xc2, 0x8b, 0x32, 0x4b, 0x52, 0x03, 0xf2, 0x33, 0xf3, 0x4a, 0x8a, 0x83, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x8b, 0x33, 0x12, 0x8b, 0x52, 0x3c, 0x5d, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x58, 0x82, 0x60, 0x5c, 0x21, 0x31, 0x2e, 0xb6, 0x02, 0xb0, 0x52, 0x09, 0x26, 0x05,
	0x66, 0x0d, 0x9e, 0x20, 0x28, 0x4f, 0x49, 0x85, 0x4b, 0x00, 0xc5, 0x9c, 0x82, 0x9c, 0x4a, 0x21,
	0x01, 0x2e, 0xe6, 0xd4, 0xa2, 0x22, 0xb0, 0x09, 0x9c, 0x41, 0x20, 0xa6, 0x91, 0x3f, 0x17, 0x9b,
	0x67, 0x5e, 0x5a, 0x50, 0x80, 0xb3, 0x90, 0x2b, 0x17, 0x37, 0x92, 0x7a, 0x21, 0x29, 0x3d, 0xa8,
	0xeb, 0x30, 0x1d, 0x23, 0x25, 0x81, 0x55, 0xae, 0x20, 0xa7, 0x52, 0x89, 0x21, 0x89, 0x0d, 0xec,
	0x1b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x71, 0x15, 0x82, 0xdd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InfRPCClient is the client API for InfRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfRPCClient interface {
	WritePoints(ctx context.Context, in *WritePointsRequest, opts ...grpc.CallOption) (*WritePointsReply, error)
}

type infRPCClient struct {
	cc *grpc.ClientConn
}

func NewInfRPCClient(cc *grpc.ClientConn) InfRPCClient {
	return &infRPCClient{cc}
}

func (c *infRPCClient) WritePoints(ctx context.Context, in *WritePointsRequest, opts ...grpc.CallOption) (*WritePointsReply, error) {
	out := new(WritePointsReply)
	err := c.cc.Invoke(ctx, "/infrpc.InfRPC/WritePoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfRPCServer is the server API for InfRPC service.
type InfRPCServer interface {
	WritePoints(context.Context, *WritePointsRequest) (*WritePointsReply, error)
}

func RegisterInfRPCServer(s *grpc.Server, srv InfRPCServer) {
	s.RegisterService(&_InfRPC_serviceDesc, srv)
}

func _InfRPC_WritePoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WritePointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfRPCServer).WritePoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infrpc.InfRPC/WritePoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfRPCServer).WritePoints(ctx, req.(*WritePointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infrpc.InfRPC",
	HandlerType: (*InfRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WritePoints",
			Handler:    _InfRPC_WritePoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infrpc.proto",
}