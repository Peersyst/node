// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: packages/blockchain/poa/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	types "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgAddValidator defines a message for adding a new validator
type MsgAddValidator struct {
	Authority        string            `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	ValidatorAddress string            `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Description      types.Description `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	Pubkey           *types1.Any       `protobuf:"bytes,4,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (m *MsgAddValidator) Reset()         { *m = MsgAddValidator{} }
func (m *MsgAddValidator) String() string { return proto.CompactTextString(m) }
func (*MsgAddValidator) ProtoMessage()    {}
func (*MsgAddValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_22361659a4d452f4, []int{0}
}
func (m *MsgAddValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddValidator.Merge(m, src)
}
func (m *MsgAddValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddValidator proto.InternalMessageInfo

func (m *MsgAddValidator) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgAddValidator) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *MsgAddValidator) GetDescription() types.Description {
	if m != nil {
		return m.Description
	}
	return types.Description{}
}

func (m *MsgAddValidator) GetPubkey() *types1.Any {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

// MsgAddValidatorResponse defines the response for adding a new validator
type MsgAddValidatorResponse struct {
}

func (m *MsgAddValidatorResponse) Reset()         { *m = MsgAddValidatorResponse{} }
func (m *MsgAddValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddValidatorResponse) ProtoMessage()    {}
func (*MsgAddValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22361659a4d452f4, []int{1}
}
func (m *MsgAddValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddValidatorResponse.Merge(m, src)
}
func (m *MsgAddValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddValidatorResponse proto.InternalMessageInfo

// MsgRemoveValidator defines a message for removing an existing validator
type MsgRemoveValidator struct {
	Authority        string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *MsgRemoveValidator) Reset()         { *m = MsgRemoveValidator{} }
func (m *MsgRemoveValidator) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveValidator) ProtoMessage()    {}
func (*MsgRemoveValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_22361659a4d452f4, []int{2}
}
func (m *MsgRemoveValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveValidator.Merge(m, src)
}
func (m *MsgRemoveValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveValidator proto.InternalMessageInfo

func (m *MsgRemoveValidator) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgRemoveValidator) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

// MsgRemoveValidatorResponse defines the response for removing an existing
// validator
type MsgRemoveValidatorResponse struct {
}

func (m *MsgRemoveValidatorResponse) Reset()         { *m = MsgRemoveValidatorResponse{} }
func (m *MsgRemoveValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveValidatorResponse) ProtoMessage()    {}
func (*MsgRemoveValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22361659a4d452f4, []int{3}
}
func (m *MsgRemoveValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveValidatorResponse.Merge(m, src)
}
func (m *MsgRemoveValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveValidatorResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddValidator)(nil), "packages.blockchain.poa.MsgAddValidator")
	proto.RegisterType((*MsgAddValidatorResponse)(nil), "packages.blockchain.poa.MsgAddValidatorResponse")
	proto.RegisterType((*MsgRemoveValidator)(nil), "packages.blockchain.poa.MsgRemoveValidator")
	proto.RegisterType((*MsgRemoveValidatorResponse)(nil), "packages.blockchain.poa.MsgRemoveValidatorResponse")
}

func init() { proto.RegisterFile("packages/blockchain/poa/tx.proto", fileDescriptor_22361659a4d452f4) }

var fileDescriptor_22361659a4d452f4 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x4f, 0x76, 0x65, 0xa1, 0xb3, 0xe2, 0xba, 0xa1, 0xd0, 0x6c, 0xd0, 0x58, 0xaa, 0x48, 0x59,
	0x71, 0xc6, 0xdd, 0x82, 0x07, 0x4f, 0xb6, 0xa8, 0x07, 0xa5, 0xb0, 0x44, 0xf0, 0xe0, 0x65, 0x99,
	0x24, 0xe3, 0x34, 0xb6, 0xc9, 0x1b, 0x32, 0x93, 0xd0, 0x5c, 0xfd, 0x04, 0x7e, 0x07, 0x2f, 0x1e,
	0xf7, 0xe0, 0xdd, 0xeb, 0xe2, 0x69, 0xf1, 0xe4, 0x49, 0xa4, 0x05, 0xfb, 0x35, 0xa4, 0xf9, 0x63,
	0xb4, 0xa5, 0xa2, 0x78, 0x49, 0xf2, 0xf2, 0xfb, 0x33, 0xef, 0xfd, 0x78, 0x0c, 0x6a, 0x0b, 0xea,
	0x8d, 0x29, 0x67, 0x92, 0xb8, 0x13, 0xf0, 0xc6, 0xde, 0x88, 0x06, 0x11, 0x11, 0x40, 0x89, 0x9a,
	0x62, 0x11, 0x83, 0x02, 0xa3, 0x55, 0x31, 0x70, 0xcd, 0xc0, 0x02, 0xa8, 0xd5, 0xe4, 0xc0, 0x21,
	0xe7, 0x90, 0xe5, 0x57, 0x41, 0xb7, 0x0e, 0x3c, 0x90, 0x21, 0xc8, 0xd3, 0x02, 0x28, 0x8a, 0x12,
	0x6a, 0x15, 0x15, 0x09, 0x25, 0x27, 0xe9, 0xd1, 0xf2, 0x55, 0x02, 0xb7, 0x4a, 0x40, 0x2a, 0x3a,
	0x0e, 0xa2, 0x25, 0xe8, 0x32, 0x45, 0x8f, 0xaa, 0xba, 0x72, 0xe6, 0x00, 0x7c, 0xc2, 0x48, 0x5e,
	0xb9, 0xc9, 0x2b, 0x42, 0xa3, 0xac, 0x84, 0xf6, 0x69, 0x18, 0x44, 0x40, 0xf2, 0x67, 0xf1, 0xab,
	0xf3, 0x71, 0x0b, 0xed, 0x0d, 0x25, 0xef, 0xfb, 0xfe, 0x0b, 0x3a, 0x09, 0x7c, 0xaa, 0x20, 0x36,
	0xee, 0xa3, 0x06, 0x4d, 0xd4, 0x08, 0xe2, 0x40, 0x65, 0xa6, 0xde, 0xd6, 0xbb, 0x8d, 0x81, 0xf9,
	0xf9, 0xc3, 0xdd, 0x66, 0xd9, 0x65, 0xdf, 0xf7, 0x63, 0x26, 0xe5, 0x73, 0x15, 0x07, 0x11, 0x77,
	0x6a, 0xaa, 0xf1, 0x14, 0xed, 0xa7, 0x95, 0xc9, 0x29, 0x2d, 0x58, 0xe6, 0x56, 0xae, 0xbf, 0xbe,
	0x49, 0xff, 0x7e, 0x71, 0x76, 0xa8, 0x39, 0x57, 0x7f, 0xea, 0x4a, 0xd0, 0x38, 0x41, 0xbb, 0x3e,
	0x93, 0x5e, 0x1c, 0x08, 0x15, 0x40, 0x64, 0x6e, 0xb7, 0xf5, 0xee, 0xee, 0xf1, 0x4d, 0x5c, 0x5a,
	0x54, 0x13, 0x97, 0x09, 0xe0, 0x47, 0x35, 0x75, 0xd0, 0x38, 0xff, 0x7a, 0x43, 0x5b, 0xda, 0xea,
	0xce, 0xaf, 0x16, 0xc6, 0x13, 0xb4, 0x23, 0x12, 0x77, 0xcc, 0x32, 0xf3, 0x52, 0x6e, 0xd6, 0xc4,
	0x45, 0x50, 0xb8, 0x0a, 0x0a, 0xf7, 0xa3, 0x6c, 0x60, 0x7e, 0xaa, 0x1b, 0xf5, 0xe2, 0x4c, 0x28,
	0xc0, 0x27, 0x89, 0xfb, 0x8c, 0x65, 0x4e, 0xa9, 0x7e, 0x70, 0xe5, 0xcd, 0xe2, 0xec, 0xb0, 0x9e,
	0xba, 0x73, 0x80, 0x5a, 0x2b, 0x01, 0x3a, 0x4c, 0x0a, 0x88, 0x24, 0xeb, 0xbc, 0xd3, 0x91, 0x31,
	0x94, 0xdc, 0x61, 0x21, 0xa4, 0xec, 0xff, 0xf3, 0x7d, 0xbc, 0x39, 0xdf, 0xcd, 0xfa, 0xb5, 0x68,
	0xd7, 0x06, 0xb8, 0x86, 0xac, 0xf5, 0x26, 0xab, 0x19, 0x8e, 0xbf, 0xeb, 0x68, 0x7b, 0x28, 0xb9,
	0xf1, 0x1a, 0x5d, 0xfe, 0x6d, 0x49, 0xba, 0x78, 0xc3, 0xc2, 0xe3, 0x95, 0x34, 0xac, 0x7b, 0x7f,
	0xcb, 0xac, 0xce, 0x34, 0x24, 0xda, 0x5b, 0xcd, 0xec, 0xce, 0x9f, 0x4c, 0x56, 0xc8, 0x56, 0xef,
	0x1f, 0xc8, 0xd5, 0xa1, 0x83, 0x87, 0xe7, 0x33, 0x5b, 0xbf, 0x98, 0xd9, 0xfa, 0xb7, 0x99, 0xad,
	0xbf, 0x9d, 0xdb, 0xda, 0xc5, 0xdc, 0xd6, 0xbe, 0xcc, 0x6d, 0xed, 0xe5, 0x6d, 0x1e, 0xa8, 0x51,
	0xe2, 0x62, 0x0f, 0x42, 0x32, 0x8d, 0xc5, 0x84, 0xa5, 0x21, 0x89, 0xc0, 0x67, 0x24, 0xed, 0x91,
	0x69, 0x71, 0x0b, 0x64, 0x82, 0x49, 0x77, 0x27, 0xdf, 0xa4, 0xde, 0x8f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x59, 0x74, 0x0f, 0x97, 0x2d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Adds a new validator into the authority
	AddValidator(ctx context.Context, in *MsgAddValidator, opts ...grpc.CallOption) (*MsgAddValidatorResponse, error)
	// Removes an existing validator from the authority
	RemoveValidator(ctx context.Context, in *MsgRemoveValidator, opts ...grpc.CallOption) (*MsgRemoveValidatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddValidator(ctx context.Context, in *MsgAddValidator, opts ...grpc.CallOption) (*MsgAddValidatorResponse, error) {
	out := new(MsgAddValidatorResponse)
	err := c.cc.Invoke(ctx, "/packages.blockchain.poa.Msg/AddValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveValidator(ctx context.Context, in *MsgRemoveValidator, opts ...grpc.CallOption) (*MsgRemoveValidatorResponse, error) {
	out := new(MsgRemoveValidatorResponse)
	err := c.cc.Invoke(ctx, "/packages.blockchain.poa.Msg/RemoveValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Adds a new validator into the authority
	AddValidator(context.Context, *MsgAddValidator) (*MsgAddValidatorResponse, error)
	// Removes an existing validator from the authority
	RemoveValidator(context.Context, *MsgRemoveValidator) (*MsgRemoveValidatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddValidator(ctx context.Context, req *MsgAddValidator) (*MsgAddValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddValidator not implemented")
}
func (*UnimplementedMsgServer) RemoveValidator(ctx context.Context, req *MsgRemoveValidator) (*MsgRemoveValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveValidator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packages.blockchain.poa.Msg/AddValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddValidator(ctx, req.(*MsgAddValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packages.blockchain.poa.Msg/RemoveValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveValidator(ctx, req.(*MsgRemoveValidator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "packages.blockchain.poa.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddValidator",
			Handler:    _Msg_AddValidator_Handler,
		},
		{
			MethodName: "RemoveValidator",
			Handler:    _Msg_RemoveValidator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "packages/blockchain/poa/tx.proto",
}

func (m *MsgAddValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pubkey != nil {
		{
			size, err := m.Pubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRemoveValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRemoveValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Description.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.Pubkey != nil {
		l = m.Pubkey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRemoveValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRemoveValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAddValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pubkey == nil {
				m.Pubkey = &types1.Any{}
			}
			if err := m.Pubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAddValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAddValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRemoveValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRemoveValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRemoveValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRemoveValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
