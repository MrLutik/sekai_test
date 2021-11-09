// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/tokens/proposal.proto

package tokens

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type ProposalUpsertTokenAlias struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon                 string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Decimals             uint32   `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Denoms               []string `protobuf:"bytes,5,rep,name=denoms,proto3" json:"denoms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProposalUpsertTokenAlias) Reset()         { *m = ProposalUpsertTokenAlias{} }
func (m *ProposalUpsertTokenAlias) String() string { return proto.CompactTextString(m) }
func (*ProposalUpsertTokenAlias) ProtoMessage()    {}
func (*ProposalUpsertTokenAlias) Descriptor() ([]byte, []int) {
	return fileDescriptor_68008d794328e180, []int{0}
}

func (m *ProposalUpsertTokenAlias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalUpsertTokenAlias.Unmarshal(m, b)
}
func (m *ProposalUpsertTokenAlias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalUpsertTokenAlias.Marshal(b, m, deterministic)
}
func (m *ProposalUpsertTokenAlias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalUpsertTokenAlias.Merge(m, src)
}
func (m *ProposalUpsertTokenAlias) XXX_Size() int {
	return xxx_messageInfo_ProposalUpsertTokenAlias.Size(m)
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
	Denom                string   `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Rate                 string   `protobuf:"bytes,2,opt,name=rate,proto3" json:"rate,omitempty"`
	FeePayments          bool     `protobuf:"varint,3,opt,name=fee_payments,json=feePayments,proto3" json:"fee_payments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProposalUpsertTokenRates) Reset()         { *m = ProposalUpsertTokenRates{} }
func (m *ProposalUpsertTokenRates) String() string { return proto.CompactTextString(m) }
func (*ProposalUpsertTokenRates) ProtoMessage()    {}
func (*ProposalUpsertTokenRates) Descriptor() ([]byte, []int) {
	return fileDescriptor_68008d794328e180, []int{1}
}

func (m *ProposalUpsertTokenRates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalUpsertTokenRates.Unmarshal(m, b)
}
func (m *ProposalUpsertTokenRates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalUpsertTokenRates.Marshal(b, m, deterministic)
}
func (m *ProposalUpsertTokenRates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalUpsertTokenRates.Merge(m, src)
}
func (m *ProposalUpsertTokenRates) XXX_Size() int {
	return xxx_messageInfo_ProposalUpsertTokenRates.Size(m)
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

func (m *ProposalUpsertTokenRates) GetRate() string {
	if m != nil {
		return m.Rate
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

func init() {
	proto.RegisterFile("kira/tokens/proposal.proto", fileDescriptor_68008d794328e180)
}

var fileDescriptor_68008d794328e180 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4f, 0x4b, 0xf3, 0x30,
	0x1c, 0x7e, 0xfb, 0xee, 0xcf, 0xbb, 0x37, 0x53, 0x84, 0x30, 0xa4, 0xf6, 0xd2, 0xd9, 0x83, 0xec,
	0xb2, 0xe6, 0x20, 0x88, 0x0c, 0x3c, 0xb8, 0xe9, 0x41, 0x04, 0x99, 0x65, 0x82, 0x78, 0x19, 0x59,
	0xf7, 0x5b, 0x2d, 0x6d, 0x92, 0xd2, 0xc4, 0x43, 0x3f, 0x8d, 0x7e, 0x08, 0xc1, 0xaf, 0xe0, 0xd9,
	0xa3, 0x87, 0x9d, 0x3d, 0xfb, 0x09, 0x24, 0x49, 0x91, 0x81, 0x9e, 0xf2, 0xfc, 0x49, 0x9e, 0x3c,
	0xc9, 0x0f, 0x79, 0x59, 0x5a, 0x52, 0xa2, 0x44, 0x06, 0x5c, 0x92, 0xa2, 0x14, 0x85, 0x90, 0x34,
	0x0f, 0x8b, 0x52, 0x28, 0x81, 0xbb, 0xda, 0x0b, 0xad, 0xe7, 0xf5, 0x12, 0x91, 0x08, 0xa3, 0x13,
	0x8d, 0xec, 0x16, 0x6f, 0x2f, 0x16, 0x92, 0x09, 0x39, 0xb7, 0x86, 0x25, 0xd6, 0x0a, 0x1e, 0x1d,
	0xe4, 0x4e, 0xeb, 0xc0, 0x9b, 0x42, 0x42, 0xa9, 0x66, 0x3a, 0xe9, 0x34, 0x4f, 0xa9, 0xc4, 0xbb,
	0xa8, 0x2d, 0x2b, 0xb6, 0x10, 0xb9, 0xeb, 0xf4, 0x9d, 0xc1, 0xff, 0xa8, 0x66, 0x18, 0xa3, 0x26,
	0xa7, 0x0c, 0xdc, 0xbf, 0x46, 0x35, 0x58, 0x6b, 0x69, 0x2c, 0xb8, 0xdb, 0xb0, 0x9a, 0xc6, 0xd8,
	0x43, 0x9d, 0x25, 0xc4, 0x29, 0xa3, 0xb9, 0x74, 0x9b, 0x7d, 0x67, 0xb0, 0x1d, 0x7d, 0x73, 0x9d,
	0xbd, 0x04, 0x2e, 0x98, 0x74, 0x5b, 0xfd, 0x86, 0xce, 0xb6, 0x6c, 0xb4, 0xf3, 0xf1, 0xe4, 0x3b,
	0x6f, 0xcf, 0xc3, 0x7f, 0x13, 0xc1, 0x15, 0x70, 0x15, 0xbc, 0xfc, 0xde, 0x30, 0xa2, 0x0a, 0x24,
	0xee, 0xa1, 0x96, 0x39, 0x57, 0x17, 0xb4, 0x04, 0x5f, 0xa3, 0x66, 0x49, 0x55, 0xdd, 0x6f, 0x7c,
	0xf2, 0xba, 0xf6, 0xff, 0xbc, 0xaf, 0xfd, 0x83, 0x24, 0x55, 0xf7, 0x0f, 0x8b, 0x30, 0x16, 0xac,
	0xfe, 0x83, 0x7a, 0x19, 0xca, 0x65, 0x46, 0x54, 0x55, 0x80, 0x0c, 0xcf, 0x20, 0xfe, 0x5c, 0xfb,
	0xdd, 0x8a, 0xb2, 0x7c, 0x14, 0xe8, 0x8c, 0x20, 0x32, 0x51, 0x78, 0x1f, 0x6d, 0xad, 0x00, 0xe6,
	0x05, 0xad, 0x18, 0x70, 0x25, 0xcd, 0x33, 0x3b, 0x51, 0x77, 0x05, 0x30, 0xad, 0xa5, 0x1f, 0xcd,
	0xc7, 0xc7, 0x77, 0x47, 0x1b, 0x57, 0x5e, 0xa6, 0x25, 0x9d, 0x88, 0x12, 0x88, 0x84, 0x8c, 0xa6,
	0xe4, 0xe2, 0x6a, 0x76, 0x1e, 0xdd, 0x12, 0x33, 0x85, 0x61, 0x02, 0x9c, 0x6c, 0x8c, 0x78, 0xd1,
	0x36, 0xf2, 0xe1, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0xb0, 0x30, 0x08, 0xf8, 0x01, 0x00,
	0x00,
}