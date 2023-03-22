// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/fee/v1/fee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types1 "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
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

// Fee defines the ICS29 receive, acknowledgement and timeout fees
type Fee struct {
	// the packet receive fee
	RecvFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=recv_fee,json=recvFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"recv_fee"`
	// the packet acknowledgement fee
	AckFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=ack_fee,json=ackFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"ack_fee"`
	// the packet timeout fee
	TimeoutFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=timeout_fee,json=timeoutFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"timeout_fee"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{0}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func (m *Fee) GetRecvFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.RecvFee
	}
	return nil
}

func (m *Fee) GetAckFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.AckFee
	}
	return nil
}

func (m *Fee) GetTimeoutFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TimeoutFee
	}
	return nil
}

// PacketFee contains ICS29 relayer fees, refund address and optional list of permitted relayers
type PacketFee struct {
	// fee encapsulates the recv, ack and timeout fees associated with an IBC packet
	Fee Fee `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`
	// the refund address for unspent fees
	RefundAddress string `protobuf:"bytes,2,opt,name=refund_address,json=refundAddress,proto3" json:"refund_address,omitempty"`
	// optional list of relayers permitted to receive fees
	Relayers []string `protobuf:"bytes,3,rep,name=relayers,proto3" json:"relayers,omitempty"`
}

func (m *PacketFee) Reset()         { *m = PacketFee{} }
func (m *PacketFee) String() string { return proto.CompactTextString(m) }
func (*PacketFee) ProtoMessage()    {}
func (*PacketFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{1}
}
func (m *PacketFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketFee.Merge(m, src)
}
func (m *PacketFee) XXX_Size() int {
	return m.Size()
}
func (m *PacketFee) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketFee.DiscardUnknown(m)
}

var xxx_messageInfo_PacketFee proto.InternalMessageInfo

func (m *PacketFee) GetFee() Fee {
	if m != nil {
		return m.Fee
	}
	return Fee{}
}

func (m *PacketFee) GetRefundAddress() string {
	if m != nil {
		return m.RefundAddress
	}
	return ""
}

func (m *PacketFee) GetRelayers() []string {
	if m != nil {
		return m.Relayers
	}
	return nil
}

// PacketFees contains a list of type PacketFee
type PacketFees struct {
	// list of packet fees
	PacketFees []PacketFee `protobuf:"bytes,1,rep,name=packet_fees,json=packetFees,proto3" json:"packet_fees"`
}

func (m *PacketFees) Reset()         { *m = PacketFees{} }
func (m *PacketFees) String() string { return proto.CompactTextString(m) }
func (*PacketFees) ProtoMessage()    {}
func (*PacketFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{2}
}
func (m *PacketFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketFees.Merge(m, src)
}
func (m *PacketFees) XXX_Size() int {
	return m.Size()
}
func (m *PacketFees) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketFees.DiscardUnknown(m)
}

var xxx_messageInfo_PacketFees proto.InternalMessageInfo

func (m *PacketFees) GetPacketFees() []PacketFee {
	if m != nil {
		return m.PacketFees
	}
	return nil
}

// IdentifiedPacketFees contains a list of type PacketFee and associated PacketId
type IdentifiedPacketFees struct {
	// unique packet identifier comprised of the channel ID, port ID and sequence
	PacketId types1.PacketId `protobuf:"bytes,1,opt,name=packet_id,json=packetId,proto3" json:"packet_id"`
	// list of packet fees
	PacketFees []PacketFee `protobuf:"bytes,2,rep,name=packet_fees,json=packetFees,proto3" json:"packet_fees"`
}

func (m *IdentifiedPacketFees) Reset()         { *m = IdentifiedPacketFees{} }
func (m *IdentifiedPacketFees) String() string { return proto.CompactTextString(m) }
func (*IdentifiedPacketFees) ProtoMessage()    {}
func (*IdentifiedPacketFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{3}
}
func (m *IdentifiedPacketFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentifiedPacketFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentifiedPacketFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentifiedPacketFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifiedPacketFees.Merge(m, src)
}
func (m *IdentifiedPacketFees) XXX_Size() int {
	return m.Size()
}
func (m *IdentifiedPacketFees) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifiedPacketFees.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifiedPacketFees proto.InternalMessageInfo

func (m *IdentifiedPacketFees) GetPacketId() types1.PacketId {
	if m != nil {
		return m.PacketId
	}
	return types1.PacketId{}
}

func (m *IdentifiedPacketFees) GetPacketFees() []PacketFee {
	if m != nil {
		return m.PacketFees
	}
	return nil
}

func init() {
	proto.RegisterType((*Fee)(nil), "ibc.applications.fee.v1.Fee")
	proto.RegisterType((*PacketFee)(nil), "ibc.applications.fee.v1.PacketFee")
	proto.RegisterType((*PacketFees)(nil), "ibc.applications.fee.v1.PacketFees")
	proto.RegisterType((*IdentifiedPacketFees)(nil), "ibc.applications.fee.v1.IdentifiedPacketFees")
}

func init() { proto.RegisterFile("ibc/applications/fee/v1/fee.proto", fileDescriptor_cb3319f1af2a53e5) }

var fileDescriptor_cb3319f1af2a53e5 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x8b, 0x13, 0x31,
	0x14, 0xc7, 0x9b, 0x56, 0x76, 0xb7, 0x29, 0x7a, 0x18, 0x16, 0x5c, 0x8b, 0xce, 0xae, 0x03, 0x42,
	0x2f, 0x4d, 0x6c, 0x55, 0xc4, 0x9b, 0x56, 0x28, 0xf4, 0xa4, 0xf4, 0x22, 0x78, 0x59, 0x32, 0xc9,
	0x9b, 0x6e, 0xe8, 0x74, 0x32, 0x4c, 0xd2, 0xc2, 0xde, 0xfd, 0x00, 0x7e, 0x07, 0x6f, 0x7e, 0x0b,
	0x6f, 0x7b, 0xdc, 0xa3, 0x27, 0x95, 0xf6, 0x8b, 0xc8, 0xcb, 0xc4, 0x61, 0x51, 0xf6, 0x22, 0x3d,
	0xcd, 0x7b, 0x6f, 0xde, 0x7b, 0xbf, 0x7f, 0xf2, 0xf2, 0xe8, 0x63, 0x9d, 0x4a, 0x2e, 0xca, 0x32,
	0xd7, 0x52, 0x38, 0x6d, 0x0a, 0xcb, 0x33, 0x00, 0xbe, 0x19, 0xe1, 0x87, 0x95, 0x95, 0x71, 0x26,
	0xba, 0xaf, 0x53, 0xc9, 0x6e, 0xa6, 0x30, 0xfc, 0xb7, 0x19, 0xf5, 0x63, 0x69, 0xec, 0xca, 0x58,
	0x9e, 0x0a, 0x8b, 0x25, 0x29, 0x38, 0x31, 0xe2, 0xd2, 0xe8, 0xa2, 0x2e, 0xec, 0x1f, 0x2f, 0xcc,
	0xc2, 0x78, 0x93, 0xa3, 0x15, 0xa2, 0x9e, 0x28, 0x4d, 0x05, 0x5c, 0x5e, 0x88, 0xa2, 0x80, 0x1c,
	0x69, 0xc1, 0xac, 0x53, 0x92, 0x6f, 0x6d, 0xda, 0x99, 0x02, 0x44, 0x19, 0x3d, 0xaa, 0x40, 0x6e,
	0xce, 0x33, 0x80, 0x13, 0x72, 0xd6, 0x19, 0xf4, 0xc6, 0x0f, 0x58, 0xcd, 0x64, 0xc8, 0x64, 0x81,
	0xc9, 0xde, 0x1a, 0x5d, 0x4c, 0x9e, 0x5e, 0xfd, 0x38, 0x6d, 0x7d, 0xfd, 0x79, 0x3a, 0x58, 0x68,
	0x77, 0xb1, 0x4e, 0x99, 0x34, 0x2b, 0x1e, 0x04, 0xd6, 0x9f, 0xa1, 0x55, 0x4b, 0xee, 0x2e, 0x4b,
	0xb0, 0xbe, 0xc0, 0xce, 0x0f, 0xb1, 0x39, 0x72, 0x14, 0x3d, 0x14, 0x72, 0xe9, 0x31, 0xed, 0xfd,
	0x63, 0x0e, 0x84, 0x5c, 0x22, 0x25, 0xa7, 0x3d, 0xa7, 0x57, 0x60, 0xd6, 0xce, 0x93, 0x3a, 0xfb,
	0x27, 0xd1, 0xd0, 0x7f, 0x0a, 0x90, 0x7c, 0x22, 0xb4, 0xfb, 0x5e, 0xc8, 0x25, 0xa0, 0x17, 0x3d,
	0xa7, 0x9d, 0xfa, 0x12, 0xc9, 0xa0, 0x37, 0x7e, 0xc8, 0x6e, 0x99, 0x28, 0x9b, 0x02, 0x4c, 0xee,
	0x20, 0x76, 0x8e, 0xe9, 0xd1, 0x13, 0x7a, 0xaf, 0x82, 0x6c, 0x5d, 0xa8, 0x73, 0xa1, 0x54, 0x05,
	0xd6, 0x9e, 0xb4, 0xcf, 0xc8, 0xa0, 0x3b, 0xbf, 0x5b, 0x47, 0xdf, 0xd4, 0xc1, 0xa8, 0x8f, 0x63,
	0xca, 0xc5, 0x25, 0x54, 0xd6, 0x9f, 0xaa, 0x3b, 0x6f, 0xfc, 0xe4, 0x03, 0xa5, 0x8d, 0x0a, 0x1b,
	0xcd, 0x68, 0xaf, 0xf4, 0x1e, 0xde, 0x80, 0x0d, 0x33, 0x4d, 0x6e, 0x95, 0xd3, 0x54, 0x06, 0x51,
	0xb4, 0x6c, 0x5a, 0x25, 0x5f, 0x08, 0x3d, 0x9e, 0x29, 0x28, 0x9c, 0xce, 0x34, 0xa8, 0x1b, 0x8c,
	0xd7, 0xb4, 0x1b, 0x18, 0x5a, 0x85, 0x03, 0x3f, 0xf2, 0x04, 0x7c, 0x73, 0xec, 0xcf, 0x43, 0x6b,
	0xba, 0xcf, 0x54, 0x68, 0x7e, 0x54, 0x06, 0xff, 0x6f, 0x95, 0xed, 0xff, 0x57, 0x39, 0x79, 0x77,
	0xb5, 0x8d, 0xc9, 0xf5, 0x36, 0x26, 0xbf, 0xb6, 0x31, 0xf9, 0xbc, 0x8b, 0x5b, 0xd7, 0xbb, 0xb8,
	0xf5, 0x7d, 0x17, 0xb7, 0x3e, 0xbe, 0xf8, 0x77, 0xaa, 0x3a, 0x95, 0xc3, 0x85, 0xe1, 0x9b, 0x97,
	0x7c, 0x65, 0xd4, 0x3a, 0x07, 0x8b, 0x8b, 0x69, 0xf9, 0xf8, 0xd5, 0x10, 0x77, 0xd2, 0x0f, 0x3a,
	0x3d, 0xf0, 0x1b, 0xf2, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xec, 0x7b, 0x52, 0xb8,
	0x03, 0x00, 0x00,
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TimeoutFee) > 0 {
		for iNdEx := len(m.TimeoutFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimeoutFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AckFee) > 0 {
		for iNdEx := len(m.AckFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AckFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RecvFee) > 0 {
		for iNdEx := len(m.RecvFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecvFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PacketFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relayers) > 0 {
		for iNdEx := len(m.Relayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Relayers[iNdEx])
			copy(dAtA[i:], m.Relayers[iNdEx])
			i = encodeVarintFee(dAtA, i, uint64(len(m.Relayers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RefundAddress) > 0 {
		i -= len(m.RefundAddress)
		copy(dAtA[i:], m.RefundAddress)
		i = encodeVarintFee(dAtA, i, uint64(len(m.RefundAddress)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PacketFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PacketFees) > 0 {
		for iNdEx := len(m.PacketFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PacketFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *IdentifiedPacketFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentifiedPacketFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentifiedPacketFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PacketFees) > 0 {
		for iNdEx := len(m.PacketFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PacketFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.PacketId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RecvFee) > 0 {
		for _, e := range m.RecvFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.AckFee) > 0 {
		for _, e := range m.AckFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.TimeoutFee) > 0 {
		for _, e := range m.TimeoutFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *PacketFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Fee.Size()
	n += 1 + l + sovFee(uint64(l))
	l = len(m.RefundAddress)
	if l > 0 {
		n += 1 + l + sovFee(uint64(l))
	}
	if len(m.Relayers) > 0 {
		for _, s := range m.Relayers {
			l = len(s)
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *PacketFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PacketFees) > 0 {
		for _, e := range m.PacketFees {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *IdentifiedPacketFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PacketId.Size()
	n += 1 + l + sovFee(uint64(l))
	if len(m.PacketFees) > 0 {
		for _, e := range m.PacketFees {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func sovFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFee(x uint64) (n int) {
	return sovFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecvFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecvFee = append(m.RecvFee, types.Coin{})
			if err := m.RecvFee[len(m.RecvFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckFee = append(m.AckFee, types.Coin{})
			if err := m.AckFee[len(m.AckFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeoutFee = append(m.TimeoutFee, types.Coin{})
			if err := m.TimeoutFee[len(m.TimeoutFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func (m *PacketFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: PacketFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefundAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayers = append(m.Relayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func (m *PacketFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: PacketFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketFees = append(m.PacketFees, PacketFee{})
			if err := m.PacketFees[len(m.PacketFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func (m *IdentifiedPacketFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: IdentifiedPacketFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentifiedPacketFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PacketId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketFees = append(m.PacketFees, PacketFee{})
			if err := m.PacketFees[len(m.PacketFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func skipFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
				return 0, ErrInvalidLengthFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFee = fmt.Errorf("proto: unexpected end of group")
)
