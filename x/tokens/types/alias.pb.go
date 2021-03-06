// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/tokens/alias.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type TokenAlias struct {
	Symbol   string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon     string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Decimals uint32   `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Denoms   []string `protobuf:"bytes,5,rep,name=denoms,proto3" json:"denoms,omitempty"`
}

func (m *TokenAlias) Reset()         { *m = TokenAlias{} }
func (m *TokenAlias) String() string { return proto.CompactTextString(m) }
func (*TokenAlias) ProtoMessage()    {}
func (*TokenAlias) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5f7b26b3e48e5a6, []int{0}
}
func (m *TokenAlias) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenAlias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenAlias.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenAlias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenAlias.Merge(m, src)
}
func (m *TokenAlias) XXX_Size() int {
	return m.Size()
}
func (m *TokenAlias) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenAlias.DiscardUnknown(m)
}

var xxx_messageInfo_TokenAlias proto.InternalMessageInfo

func (m *TokenAlias) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TokenAlias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TokenAlias) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *TokenAlias) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *TokenAlias) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

type MsgUpsertTokenAlias struct {
	Symbol   string                                        `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name     string                                        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon     string                                        `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Decimals uint32                                        `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Denoms   []string                                      `protobuf:"bytes,5,rep,name=denoms,proto3" json:"denoms,omitempty"`
	Proposer github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty" yaml:"proposer"`
}

func (m *MsgUpsertTokenAlias) Reset()         { *m = MsgUpsertTokenAlias{} }
func (m *MsgUpsertTokenAlias) String() string { return proto.CompactTextString(m) }
func (*MsgUpsertTokenAlias) ProtoMessage()    {}
func (*MsgUpsertTokenAlias) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5f7b26b3e48e5a6, []int{1}
}
func (m *MsgUpsertTokenAlias) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpsertTokenAlias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpsertTokenAlias.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpsertTokenAlias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpsertTokenAlias.Merge(m, src)
}
func (m *MsgUpsertTokenAlias) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpsertTokenAlias) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpsertTokenAlias.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpsertTokenAlias proto.InternalMessageInfo

func (m *MsgUpsertTokenAlias) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MsgUpsertTokenAlias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgUpsertTokenAlias) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *MsgUpsertTokenAlias) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *MsgUpsertTokenAlias) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

func (m *MsgUpsertTokenAlias) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenAlias)(nil), "kira.tokens.TokenAlias")
	proto.RegisterType((*MsgUpsertTokenAlias)(nil), "kira.tokens.MsgUpsertTokenAlias")
}

func init() { proto.RegisterFile("kira/tokens/alias.proto", fileDescriptor_e5f7b26b3e48e5a6) }

var fileDescriptor_e5f7b26b3e48e5a6 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x91, 0xbf, 0x4e, 0xfb, 0x30,
	0x10, 0xc7, 0xeb, 0x5f, 0xfb, 0xab, 0x8a, 0x01, 0x21, 0x05, 0x04, 0x51, 0x07, 0x37, 0xca, 0x94,
	0xa5, 0xf1, 0xc0, 0xc6, 0xd6, 0xc2, 0x86, 0x58, 0x2a, 0x58, 0x98, 0x70, 0x1d, 0x2b, 0x58, 0x89,
	0x73, 0x91, 0x2f, 0x48, 0x74, 0xe3, 0x11, 0x78, 0x2c, 0xc6, 0x8e, 0x4c, 0x15, 0x6a, 0x46, 0x36,
	0x46, 0x26, 0xe4, 0xa4, 0xad, 0x78, 0x03, 0x26, 0x7f, 0xff, 0x9c, 0x3f, 0x3a, 0xe9, 0xe8, 0x59,
	0xa6, 0xad, 0xe0, 0x15, 0x64, 0xaa, 0x40, 0x2e, 0x72, 0x2d, 0x30, 0x2e, 0x2d, 0x54, 0xe0, 0xed,
	0xbb, 0x22, 0x6e, 0x8b, 0xe1, 0x49, 0x0a, 0x29, 0x34, 0x39, 0x77, 0xaa, 0x1d, 0x09, 0x5f, 0x08,
	0xa5, 0xb7, 0x6e, 0x60, 0xe2, 0xfe, 0x79, 0xa7, 0xb4, 0x8f, 0x0b, 0x33, 0x87, 0xdc, 0x27, 0x01,
	0x89, 0xf6, 0x66, 0x1b, 0xe7, 0x79, 0xb4, 0x57, 0x08, 0xa3, 0xfc, 0x7f, 0x4d, 0xda, 0x68, 0x97,
	0x69, 0x09, 0x85, 0xdf, 0x6d, 0x33, 0xa7, 0xbd, 0x21, 0x1d, 0x24, 0x4a, 0x6a, 0x23, 0x72, 0xf4,
	0x7b, 0x01, 0x89, 0x0e, 0x67, 0x3b, 0xef, 0xd8, 0x89, 0x2a, 0xc0, 0xa0, 0xff, 0x3f, 0xe8, 0x3a,
	0x76, 0xeb, 0xc2, 0x4f, 0x42, 0x8f, 0x6f, 0x30, 0xbd, 0x2b, 0x51, 0xd9, 0xea, 0x6f, 0x77, 0xf1,
	0x1e, 0xe8, 0xa0, 0xb4, 0x50, 0x02, 0x2a, 0xeb, 0xf7, 0x03, 0x12, 0x1d, 0x4c, 0xaf, 0xbe, 0x56,
	0xa3, 0xa3, 0x85, 0x30, 0xf9, 0x45, 0xb8, 0x6d, 0xc2, 0xef, 0xd5, 0x68, 0x9c, 0xea, 0xea, 0xf1,
	0x69, 0x1e, 0x4b, 0x30, 0x5c, 0x02, 0x1a, 0xc0, 0xcd, 0x33, 0xc6, 0x24, 0xe3, 0xd5, 0xa2, 0x54,
	0x18, 0x4f, 0xa4, 0x9c, 0x24, 0x89, 0x55, 0x88, 0xb3, 0x1d, 0x75, 0x3a, 0x7d, 0x5b, 0x33, 0xb2,
	0x5c, 0x33, 0xf2, 0xb1, 0x66, 0xe4, 0xb5, 0x66, 0x9d, 0x65, 0xcd, 0x3a, 0xef, 0x35, 0xeb, 0xdc,
	0x47, 0xbf, 0x90, 0xd7, 0xda, 0x8a, 0x4b, 0xb0, 0x8a, 0xa3, 0xca, 0x84, 0xe6, 0xcf, 0xdb, 0xeb,
	0x36, 0xe0, 0x79, 0xbf, 0xb9, 0xdd, 0xf9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0x38, 0xf5,
	0x56, 0xf9, 0x01, 0x00, 0x00,
}

func (m *TokenAlias) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenAlias) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenAlias) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintAlias(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Decimals != 0 {
		i = encodeVarintAlias(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Icon) > 0 {
		i -= len(m.Icon)
		copy(dAtA[i:], m.Icon)
		i = encodeVarintAlias(dAtA, i, uint64(len(m.Icon)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAlias(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintAlias(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpsertTokenAlias) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpsertTokenAlias) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpsertTokenAlias) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintAlias(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintAlias(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Decimals != 0 {
		i = encodeVarintAlias(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Icon) > 0 {
		i -= len(m.Icon)
		copy(dAtA[i:], m.Icon)
		i = encodeVarintAlias(dAtA, i, uint64(len(m.Icon)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAlias(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintAlias(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAlias(dAtA []byte, offset int, v uint64) int {
	offset -= sovAlias(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenAlias) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovAlias(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAlias(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovAlias(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovAlias(uint64(m.Decimals))
	}
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovAlias(uint64(l))
		}
	}
	return n
}

func (m *MsgUpsertTokenAlias) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovAlias(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAlias(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovAlias(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovAlias(uint64(m.Decimals))
	}
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovAlias(uint64(l))
		}
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovAlias(uint64(l))
	}
	return n
}

func sovAlias(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAlias(x uint64) (n int) {
	return sovAlias(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenAlias) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlias
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
			return fmt.Errorf("proto: TokenAlias: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenAlias: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlias
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
				return ErrInvalidLengthAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlias
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
					return ErrIntOverflowAlias
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
				return ErrInvalidLengthAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlias
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
					return ErrIntOverflowAlias
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
				return ErrInvalidLengthAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlias
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
					return ErrIntOverflowAlias
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
					return ErrIntOverflowAlias
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
				return ErrInvalidLengthAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlias
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlias(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlias
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
func (m *MsgUpsertTokenAlias) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlias
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
			return fmt.Errorf("proto: MsgUpsertTokenAlias: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpsertTokenAlias: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlias
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
				return ErrInvalidLengthAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlias
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
					return ErrIntOverflowAlias
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
				return ErrInvalidLengthAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlias
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
					return ErrIntOverflowAlias
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
				return ErrInvalidLengthAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlias
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
					return ErrIntOverflowAlias
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
					return ErrIntOverflowAlias
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
				return ErrInvalidLengthAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlias
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlias
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAlias
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAlias
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlias(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlias
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
func skipAlias(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAlias
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
					return 0, ErrIntOverflowAlias
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
					return 0, ErrIntOverflowAlias
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
				return 0, ErrInvalidLengthAlias
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAlias
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAlias
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAlias        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAlias          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAlias = fmt.Errorf("proto: unexpected end of group")
)
