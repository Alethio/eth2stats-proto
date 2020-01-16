// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry.proto

package proto

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

type PeersRequest struct {
	Peers                int64    `protobuf:"varint,1,opt,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeersRequest) Reset()         { *m = PeersRequest{} }
func (m *PeersRequest) String() string { return proto.CompactTextString(m) }
func (*PeersRequest) ProtoMessage()    {}
func (*PeersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbfcf76559f568d, []int{0}
}

func (m *PeersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeersRequest.Unmarshal(m, b)
}
func (m *PeersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeersRequest.Marshal(b, m, deterministic)
}
func (m *PeersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeersRequest.Merge(m, src)
}
func (m *PeersRequest) XXX_Size() int {
	return xxx_messageInfo_PeersRequest.Size(m)
}
func (m *PeersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeersRequest proto.InternalMessageInfo

func (m *PeersRequest) GetPeers() int64 {
	if m != nil {
		return m.Peers
	}
	return 0
}

type SyncingRequest struct {
	Syncing              bool     `protobuf:"varint,1,opt,name=syncing,proto3" json:"syncing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncingRequest) Reset()         { *m = SyncingRequest{} }
func (m *SyncingRequest) String() string { return proto.CompactTextString(m) }
func (*SyncingRequest) ProtoMessage()    {}
func (*SyncingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbfcf76559f568d, []int{1}
}

func (m *SyncingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncingRequest.Unmarshal(m, b)
}
func (m *SyncingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncingRequest.Marshal(b, m, deterministic)
}
func (m *SyncingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncingRequest.Merge(m, src)
}
func (m *SyncingRequest) XXX_Size() int {
	return xxx_messageInfo_SyncingRequest.Size(m)
}
func (m *SyncingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncingRequest proto.InternalMessageInfo

func (m *SyncingRequest) GetSyncing() bool {
	if m != nil {
		return m.Syncing
	}
	return false
}

type AttestationsRequest struct {
	AttestationsInPool   int64    `protobuf:"varint,1,opt,name=attestationsInPool,proto3" json:"attestationsInPool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestationsRequest) Reset()         { *m = AttestationsRequest{} }
func (m *AttestationsRequest) String() string { return proto.CompactTextString(m) }
func (*AttestationsRequest) ProtoMessage()    {}
func (*AttestationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbfcf76559f568d, []int{2}
}

func (m *AttestationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationsRequest.Unmarshal(m, b)
}
func (m *AttestationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationsRequest.Marshal(b, m, deterministic)
}
func (m *AttestationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationsRequest.Merge(m, src)
}
func (m *AttestationsRequest) XXX_Size() int {
	return xxx_messageInfo_AttestationsRequest.Size(m)
}
func (m *AttestationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationsRequest proto.InternalMessageInfo

func (m *AttestationsRequest) GetAttestationsInPool() int64 {
	if m != nil {
		return m.AttestationsInPool
	}
	return 0
}

type MemoryUsageRequest struct {
	MemoryUsage          int64    `protobuf:"varint,1,opt,name=memoryUsage,proto3" json:"memoryUsage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemoryUsageRequest) Reset()         { *m = MemoryUsageRequest{} }
func (m *MemoryUsageRequest) String() string { return proto.CompactTextString(m) }
func (*MemoryUsageRequest) ProtoMessage()    {}
func (*MemoryUsageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbfcf76559f568d, []int{3}
}

func (m *MemoryUsageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryUsageRequest.Unmarshal(m, b)
}
func (m *MemoryUsageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryUsageRequest.Marshal(b, m, deterministic)
}
func (m *MemoryUsageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryUsageRequest.Merge(m, src)
}
func (m *MemoryUsageRequest) XXX_Size() int {
	return xxx_messageInfo_MemoryUsageRequest.Size(m)
}
func (m *MemoryUsageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryUsageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryUsageRequest proto.InternalMessageInfo

func (m *MemoryUsageRequest) GetMemoryUsage() int64 {
	if m != nil {
		return m.MemoryUsage
	}
	return 0
}

type DefaultResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefaultResponse) Reset()         { *m = DefaultResponse{} }
func (m *DefaultResponse) String() string { return proto.CompactTextString(m) }
func (*DefaultResponse) ProtoMessage()    {}
func (*DefaultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbfcf76559f568d, []int{4}
}

func (m *DefaultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultResponse.Unmarshal(m, b)
}
func (m *DefaultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultResponse.Marshal(b, m, deterministic)
}
func (m *DefaultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultResponse.Merge(m, src)
}
func (m *DefaultResponse) XXX_Size() int {
	return xxx_messageInfo_DefaultResponse.Size(m)
}
func (m *DefaultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultResponse proto.InternalMessageInfo

func (m *DefaultResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *DefaultResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PeersRequest)(nil), "proto.PeersRequest")
	proto.RegisterType((*SyncingRequest)(nil), "proto.SyncingRequest")
	proto.RegisterType((*AttestationsRequest)(nil), "proto.AttestationsRequest")
	proto.RegisterType((*MemoryUsageRequest)(nil), "proto.MemoryUsageRequest")
	proto.RegisterType((*DefaultResponse)(nil), "proto.DefaultResponse")
}

func init() { proto.RegisterFile("telemetry.proto", fileDescriptor_edbfcf76559f568d) }

var fileDescriptor_edbfcf76559f568d = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0x95, 0xb4, 0x66, 0x5a, 0x5b, 0x98, 0xaa, 0xd4, 0x9c, 0xc2, 0xa2, 0x50, 0x3c, 0xe4,
	0x50, 0xa1, 0x07, 0x6f, 0x4a, 0x3d, 0x78, 0x10, 0xca, 0x56, 0x3f, 0x20, 0xca, 0x58, 0x02, 0x49,
	0x36, 0x66, 0x37, 0x87, 0x7c, 0x82, 0x7f, 0x2d, 0xd9, 0xec, 0xea, 0xaa, 0xb5, 0xa7, 0xf0, 0xde,
	0xbc, 0x79, 0xd9, 0xf7, 0x06, 0x26, 0x8a, 0x32, 0xca, 0x49, 0x55, 0x4d, 0x5c, 0x56, 0x42, 0x09,
	0xf4, 0xf5, 0x27, 0x0c, 0x92, 0x32, 0xed, 0x18, 0x76, 0x01, 0xa3, 0x35, 0x51, 0x25, 0x39, 0xbd,
	0xd7, 0x24, 0x15, 0x9e, 0x80, 0x5f, 0xb6, 0x78, 0xe6, 0x45, 0xde, 0xfc, 0x90, 0x77, 0x80, 0x5d,
	0xc1, 0x78, 0xd3, 0x14, 0xaf, 0x69, 0xb1, 0xb5, 0xba, 0x19, 0x0c, 0x64, 0xc7, 0x68, 0xe5, 0x11,
	0xb7, 0x90, 0xdd, 0xc3, 0xf4, 0x56, 0x29, 0x92, 0x2a, 0x51, 0xa9, 0x28, 0xbe, 0x8c, 0x63, 0xc0,
	0xc4, 0xa1, 0x1f, 0x8a, 0xb5, 0x10, 0x99, 0xf9, 0xcb, 0x8e, 0x09, 0x5b, 0x02, 0x3e, 0x52, 0x2e,
	0xaa, 0xe6, 0x59, 0x26, 0x5b, 0xb2, 0x2e, 0x11, 0x0c, 0xf3, 0x6f, 0xd6, 0xac, 0xbb, 0x14, 0xe3,
	0x30, 0x59, 0xd1, 0x5b, 0x52, 0x67, 0x8a, 0x93, 0x2c, 0x45, 0x21, 0x09, 0x2f, 0xa1, 0xdf, 0x9a,
	0xd7, 0x5d, 0xa8, 0xf1, 0xe2, 0xb8, 0xcb, 0x1e, 0x6f, 0x34, 0xc9, 0xcd, 0xb0, 0x8d, 0x94, 0x93,
	0xd4, 0xbe, 0xbd, 0xc8, 0x9b, 0x07, 0xdc, 0xc2, 0xc5, 0x47, 0x0f, 0x82, 0x27, 0x5b, 0x25, 0x2e,
	0xc1, 0xd7, 0x95, 0xe1, 0xd4, 0xf8, 0xb8, 0x05, 0x86, 0x67, 0x86, 0xfc, 0xf5, 0x08, 0x76, 0x80,
	0x2b, 0x18, 0xb9, 0xc5, 0x60, 0x68, 0x94, 0x3b, 0xda, 0xda, 0xe3, 0x72, 0x03, 0x03, 0x73, 0x0a,
	0x3c, 0xb5, 0x39, 0x7e, 0x9c, 0x66, 0xcf, 0xee, 0x1d, 0x0c, 0x9d, 0x4e, 0xf1, 0xdc, 0x08, 0xff,
	0xf6, 0xfc, 0xbf, 0xc7, 0x4b, 0x5f, 0x0f, 0xae, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xd7,
	0x66, 0xf1, 0x5c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TelemetryClient is the client API for Telemetry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TelemetryClient interface {
	Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
	Attestations(ctx context.Context, in *AttestationsRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
	Syncing(ctx context.Context, in *SyncingRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
	MemoryUsage(ctx context.Context, in *MemoryUsageRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
}

type telemetryClient struct {
	cc *grpc.ClientConn
}

func NewTelemetryClient(cc *grpc.ClientConn) TelemetryClient {
	return &telemetryClient{cc}
}

func (c *telemetryClient) Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/proto.Telemetry/Peers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) Attestations(ctx context.Context, in *AttestationsRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/proto.Telemetry/Attestations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) Syncing(ctx context.Context, in *SyncingRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/proto.Telemetry/Syncing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) MemoryUsage(ctx context.Context, in *MemoryUsageRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/proto.Telemetry/MemoryUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelemetryServer is the server API for Telemetry service.
type TelemetryServer interface {
	Peers(context.Context, *PeersRequest) (*DefaultResponse, error)
	Attestations(context.Context, *AttestationsRequest) (*DefaultResponse, error)
	Syncing(context.Context, *SyncingRequest) (*DefaultResponse, error)
	MemoryUsage(context.Context, *MemoryUsageRequest) (*DefaultResponse, error)
}

// UnimplementedTelemetryServer can be embedded to have forward compatible implementations.
type UnimplementedTelemetryServer struct {
}

func (*UnimplementedTelemetryServer) Peers(ctx context.Context, req *PeersRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peers not implemented")
}
func (*UnimplementedTelemetryServer) Attestations(ctx context.Context, req *AttestationsRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Attestations not implemented")
}
func (*UnimplementedTelemetryServer) Syncing(ctx context.Context, req *SyncingRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Syncing not implemented")
}
func (*UnimplementedTelemetryServer) MemoryUsage(ctx context.Context, req *MemoryUsageRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemoryUsage not implemented")
}

func RegisterTelemetryServer(s *grpc.Server, srv TelemetryServer) {
	s.RegisterService(&_Telemetry_serviceDesc, srv)
}

func _Telemetry_Peers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).Peers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Telemetry/Peers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).Peers(ctx, req.(*PeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_Attestations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).Attestations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Telemetry/Attestations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).Attestations(ctx, req.(*AttestationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_Syncing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).Syncing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Telemetry/Syncing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).Syncing(ctx, req.(*SyncingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_MemoryUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemoryUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).MemoryUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Telemetry/MemoryUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).MemoryUsage(ctx, req.(*MemoryUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Telemetry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Telemetry",
	HandlerType: (*TelemetryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Peers",
			Handler:    _Telemetry_Peers_Handler,
		},
		{
			MethodName: "Attestations",
			Handler:    _Telemetry_Attestations_Handler,
		},
		{
			MethodName: "Syncing",
			Handler:    _Telemetry_Syncing_Handler,
		},
		{
			MethodName: "MemoryUsage",
			Handler:    _Telemetry_MemoryUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "telemetry.proto",
}