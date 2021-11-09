// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/tokens/proposal.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type ProposalUpsertTokenAlias struct {
	Symbol   string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon     string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Decimals uint32   `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Denoms   []string `protobuf:"bytes,5,rep,name=denoms,proto3" json:"denoms,omitempty"`
}

func (m *ProposalUpsertTokenAlias) Reset()         { *m = ProposalUpsertTokenAlias{} }
func (m *ProposalUpsertTokenAlias) String() string { return proto.CompactTextString(m) }
func (*ProposalUpsertTokenAlias) ProtoMessage()    {}
func (*ProposalUpsertTokenAlias) Descriptor() ([]byte, []int) {
	return fileDescriptor_68008d794328e180, []int{0}
}
func (m *ProposalUpsertTokenAlias) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalUpsertTokenAlias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalUpsertTokenAlias.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalUpsertTokenAlias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalUpsertTokenAlias.Merge(m, src)
}
func (m *ProposalUpsertTokenAlias) XXX_Size() int {
	return m.Size()
}
func (m *ProposalUpsertTokenAlias) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalUpsertTokenAlias.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalUpsertTokenAlias proto.InternalMessageInfo

func (m *ProposalUpsertTokenAlias) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ProposalUpsertTokenAlias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProposalUpsertTokenAlias) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *ProposalUpsertTokenAlias) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *ProposalUpsertTokenAlias) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

type ProposalUpsertTokenRates struct {
	Denom       string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Rate        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate" yaml:"rate"`
	FeePayments bool                                   `protobuf:"varint,3,opt,name=fee_payments,json=feePayments,proto3" json:"fee_payments,omitempty"`
}

func (m *ProposalUpsertTokenRates) Reset()         { *m = ProposalUpsertTokenRates{} }
func (m *ProposalUpsertTokenRates) String() string { return proto.CompactTextString(m) }
func (*ProposalUpsertTokenRates) ProtoMessage()    {}
func (*ProposalUpsertTokenRates) Descriptor() ([]byte, []int) {
	return fileDescriptor_68008d794328e180, []int{1}
}
func (m *ProposalUpsertTokenRates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalUpsertTokenRates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalUpsertTokenRates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalUpsertTokenRates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalUpsertTokenRates.Merge(m, src)
}
func (m *ProposalUpsertTokenRates) XXX_Size() int {
	return m.Size()
}
func (m *ProposalUpsertTokenRates) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalUpsertTokenRates.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalUpsertTokenRates proto.InternalMessageInfo

func (m *ProposalUpsertTokenRates) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *ProposalUpsertTokenRates) GetFeePayments() bool {
	if m != nil {
		return m.FeePayments
	}
	return false
}

func init() {
	proto.RegisterType((*ProposalUpsertTokenAlias)(nil), "kira.tokens.ProposalUpsertTokenAlias")
	proto.RegisterType((*ProposalUpsertTokenRates)(nil), "kira.tokens.ProposalUpsertTokenRates")
}

func init() { proto.RegisterFile("kira/tokens/proposal.proto", fileDescriptor_68008d794328e180) }

var fileDescriptor_68008d794328e180 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x6e, 0xdb, 0x30,
	0x18, 0xc5, 0xc5, 0xfa, 0x4f, 0x5d, 0xba, 0x45, 0x01, 0xc2, 0x28, 0x54, 0x0d, 0x92, 0xab, 0xa1,
	0xf0, 0x62, 0x69, 0xe8, 0x66, 0xa0, 0x43, 0xed, 0x6e, 0x59, 0x1c, 0x21, 0x59, 0xb2, 0x18, 0xb4,
	0xfc, 0xd9, 0x11, 0x24, 0x8a, 0x82, 0xc8, 0x00, 0xd1, 0x2d, 0x72, 0x83, 0xe4, 0x10, 0x01, 0x72,
	0x05, 0x8f, 0x46, 0xa6, 0x20, 0x83, 0x11, 0xd8, 0x4b, 0xe6, 0x9c, 0x20, 0x20, 0xa9, 0x04, 0x01,
	0x92, 0x49, 0xdf, 0x7b, 0x4f, 0x7c, 0xf8, 0x91, 0x1f, 0x76, 0xd2, 0xa4, 0xa4, 0xa1, 0xe4, 0x29,
	0xe4, 0x22, 0x2c, 0x4a, 0x5e, 0x70, 0x41, 0xb3, 0xa0, 0x28, 0xb9, 0xe4, 0xa4, 0xab, 0xb2, 0xc0,
	0x64, 0x4e, 0x6f, 0xc5, 0x57, 0x5c, 0xfb, 0xa1, 0x9a, 0xcc, 0x2f, 0xce, 0xcf, 0x98, 0x0b, 0xc6,
	0xc5, 0xcc, 0x04, 0x46, 0x98, 0xc8, 0xbf, 0x44, 0xd8, 0x9e, 0xd6, 0x85, 0xc7, 0x85, 0x80, 0x52,
	0x1e, 0xa9, 0xa6, 0x7f, 0x59, 0x42, 0x05, 0xf9, 0x81, 0xdb, 0xa2, 0x62, 0x73, 0x9e, 0xd9, 0xa8,
	0x8f, 0x06, 0x5f, 0xa2, 0x5a, 0x11, 0x82, 0x9b, 0x39, 0x65, 0x60, 0x7f, 0xd2, 0xae, 0x9e, 0x95,
	0x97, 0xc4, 0x3c, 0xb7, 0x1b, 0xc6, 0x53, 0x33, 0x71, 0x70, 0x67, 0x01, 0x71, 0xc2, 0x68, 0x26,
	0xec, 0x66, 0x1f, 0x0d, 0xbe, 0x45, 0xaf, 0x5a, 0x75, 0x2f, 0x20, 0xe7, 0x4c, 0xd8, 0xad, 0x7e,
	0x43, 0x75, 0x1b, 0x35, 0xfa, 0xfe, 0x78, 0xe5, 0xa1, 0xdb, 0xeb, 0xe1, 0xe7, 0x09, 0xcf, 0x25,
	0xe4, 0xd2, 0xbf, 0xf9, 0x98, 0x30, 0xa2, 0x12, 0x04, 0xe9, 0xe1, 0x96, 0x3e, 0x57, 0x03, 0x1a,
	0x41, 0x0e, 0x71, 0xb3, 0xa4, 0xb2, 0xe6, 0x1b, 0xff, 0x5d, 0x6f, 0x3d, 0xeb, 0x7e, 0xeb, 0xfd,
	0x5e, 0x25, 0xf2, 0xf4, 0x6c, 0x1e, 0xc4, 0x9c, 0xd5, 0x6f, 0x50, 0x7f, 0x86, 0x62, 0x91, 0x86,
	0xb2, 0x2a, 0x40, 0x04, 0xff, 0x21, 0x7e, 0xda, 0x7a, 0xdd, 0x8a, 0xb2, 0x6c, 0xe4, 0xab, 0x0e,
	0x3f, 0xd2, 0x55, 0xe4, 0x17, 0xfe, 0xba, 0x04, 0x98, 0x15, 0xb4, 0x62, 0x90, 0x4b, 0xa1, 0xaf,
	0xd9, 0x89, 0xba, 0x4b, 0x80, 0x69, 0x6d, 0xbd, 0x23, 0x1f, 0x8f, 0xd7, 0x3b, 0x17, 0x6d, 0x76,
	0x2e, 0x7a, 0xd8, 0xb9, 0xe8, 0x62, 0xef, 0x5a, 0x9b, 0xbd, 0x6b, 0xdd, 0xed, 0x5d, 0xeb, 0x64,
	0xf0, 0x06, 0xe5, 0x20, 0x29, 0xe9, 0x84, 0x97, 0x10, 0x0a, 0x48, 0x69, 0x12, 0x9e, 0xbf, 0xac,
	0x59, 0x03, 0xcd, 0xdb, 0x7a, 0x4d, 0x7f, 0x9e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x2a, 0xb4,
	0xc9, 0x02, 0x02, 0x00, 0x00,
}

func (this *ProposalUpsertTokenAlias) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProposalUpsertTokenAlias)
	if !ok {
		that2, ok := that.(ProposalUpsertTokenAlias)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Symbol != that1.Symbol {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Icon != that1.Icon {
		return false
	}
	if this.Decimals != that1.Decimals {
		return false
	}
	if len(this.Denoms) != len(that1.Denoms) {
		return false
	}
	for i := range this.Denoms {
		if this.Denoms[i] != that1.Denoms[i] {
			return false
		}
	}
	return true
}
func (this *ProposalUpsertTokenRates) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProposalUpsertTokenRates)
	if !ok {
		that2, ok := that.(ProposalUpsertTokenRates)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if !this.Rate.Equal(that1.Rate) {
		return false
	}
	if this.FeePayments != that1.FeePayments {
		return false
	}
	return true
}
func (m *ProposalUpsertTokenAlias) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalUpsertTokenAlias) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalUpsertTokenAlias) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Decimals != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Icon) > 0 {
		i -= len(m.Icon)
		copy(dAtA[i:], m.Icon)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Icon)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProposalUpsertTokenRates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalUpsertTokenRates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalUpsertTokenRates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeePayments {
		i--
		if m.FeePayments {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposalUpsertTokenAlias) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovProposal(uint64(m.Decimals))
	}
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *ProposalUpsertTokenRates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = m.Rate.Size()
	n += 1 + l + sovProposal(uint64(l))
	if m.FeePayments {
		n += 2
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposalUpsertTokenAlias) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: ProposalUpsertTokenAlias: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalUpsertTokenAlias: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Icon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Icon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *ProposalUpsertTokenRates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: ProposalUpsertTokenRates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalUpsertTokenRates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePayments", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FeePayments = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)