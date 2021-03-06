// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/staking/pagination.proto

package staking

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// PageRequest is to be embedded in gRPC request messages for efficient
// pagination. Ex:
//
//  message SomeRequest {
//          Foo some_parameter = 1;
//          PageRequest pagination = 2;
//  }
type PageRequest struct {
	// key is a value returned in PageResponse.next_key to begin
	// querying the next page most efficiently. Only one of offset or key
	// should be set.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// offset is a numeric offset that can be used when key is unavailable.
	// It is less efficient than using key. Only one of offset or key should
	// be set.
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// limit is the total number of results to be returned in the result page.
	// If left empty it will default to a value to be set by each app.
	Limit uint64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// count_total is set to true  to indicate that the result set should include
	// a count of the total number of items available for pagination in UIs. count_total
	// is only respected when offset is used. It is ignored when key is set.
	CountTotal           bool     `protobuf:"varint,4,opt,name=count_total,json=countTotal,proto3" json:"count_total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageRequest) Reset()         { *m = PageRequest{} }
func (m *PageRequest) String() string { return proto.CompactTextString(m) }
func (*PageRequest) ProtoMessage()    {}
func (*PageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_553a06477479208d, []int{0}
}

func (m *PageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageRequest.Unmarshal(m, b)
}
func (m *PageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageRequest.Marshal(b, m, deterministic)
}
func (m *PageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageRequest.Merge(m, src)
}
func (m *PageRequest) XXX_Size() int {
	return xxx_messageInfo_PageRequest.Size(m)
}
func (m *PageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PageRequest proto.InternalMessageInfo

func (m *PageRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PageRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *PageRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PageRequest) GetCountTotal() bool {
	if m != nil {
		return m.CountTotal
	}
	return false
}

// PageResponse is to be embedded in gRPC response messages where the corresponding
// request message has used PageRequest.
//
//  message SomeResponse {
//          repeated Bar results = 1;
//          PageResponse page = 2;
//  }
type PageResponse struct {
	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	NextKey []byte `protobuf:"bytes,1,opt,name=next_key,json=nextKey,proto3" json:"next_key,omitempty"`
	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total                uint64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageResponse) Reset()         { *m = PageResponse{} }
func (m *PageResponse) String() string { return proto.CompactTextString(m) }
func (*PageResponse) ProtoMessage()    {}
func (*PageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_553a06477479208d, []int{1}
}

func (m *PageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageResponse.Unmarshal(m, b)
}
func (m *PageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageResponse.Marshal(b, m, deterministic)
}
func (m *PageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageResponse.Merge(m, src)
}
func (m *PageResponse) XXX_Size() int {
	return xxx_messageInfo_PageResponse.Size(m)
}
func (m *PageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PageResponse proto.InternalMessageInfo

func (m *PageResponse) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

func (m *PageResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*PageRequest)(nil), "kira.staking.PageRequest")
	proto.RegisterType((*PageResponse)(nil), "kira.staking.PageResponse")
}

func init() {
	proto.RegisterFile("kira/staking/pagination.proto", fileDescriptor_553a06477479208d)
}

var fileDescriptor_553a06477479208d = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4b, 0x03, 0x31,
	0x14, 0x87, 0x39, 0x5b, 0x6b, 0x79, 0xbd, 0x41, 0x82, 0xc8, 0x39, 0x88, 0x47, 0xa7, 0x5b, 0xbc,
	0x0c, 0x0e, 0xe2, 0x24, 0x28, 0x0e, 0x52, 0x10, 0x09, 0x1d, 0xc4, 0xa5, 0xa4, 0xe5, 0x35, 0x3e,
	0xae, 0xcd, 0x3b, 0x2f, 0xef, 0xc0, 0xfe, 0xf7, 0x92, 0x4b, 0x41, 0xb7, 0x7c, 0xdf, 0x0f, 0xf2,
	0x91, 0xc0, 0x75, 0x43, 0x9d, 0xd5, 0x41, 0x6c, 0x43, 0xde, 0xe9, 0xd6, 0x3a, 0xf2, 0x56, 0x88,
	0x7d, 0xdd, 0x76, 0x2c, 0xac, 0xf2, 0x38, 0xd7, 0xc7, 0x79, 0xee, 0x61, 0xf6, 0x6e, 0x1d, 0x1a,
	0xfc, 0xee, 0x31, 0x88, 0x3a, 0x87, 0x51, 0x83, 0x87, 0x22, 0x2b, 0xb3, 0x2a, 0x37, 0xf1, 0xa8,
	0x2e, 0x61, 0xc2, 0xdb, 0x6d, 0x40, 0x29, 0x4e, 0xca, 0xac, 0x1a, 0x9b, 0x23, 0xa9, 0x0b, 0x38,
	0xdd, 0xd1, 0x9e, 0xa4, 0x18, 0x0d, 0x3a, 0x81, 0xba, 0x81, 0xd9, 0x86, 0x7b, 0x2f, 0x2b, 0x61,
	0xb1, 0xbb, 0x62, 0x5c, 0x66, 0xd5, 0xd4, 0xc0, 0xa0, 0x96, 0xd1, 0xcc, 0x1f, 0x21, 0x4f, 0xbd,
	0xd0, 0xb2, 0x0f, 0xa8, 0xae, 0x60, 0xea, 0xf1, 0x47, 0x56, 0x7f, 0xd5, 0xb3, 0xc8, 0x0b, 0x3c,
	0xc4, 0x42, 0xba, 0x25, 0x85, 0x13, 0x3c, 0x3d, 0x7c, 0xde, 0x3b, 0x92, 0xaf, 0x7e, 0x5d, 0x6f,
	0x78, 0xaf, 0x17, 0xd4, 0xd9, 0x67, 0xee, 0x50, 0x07, 0x6c, 0x2c, 0xe9, 0xd7, 0xb7, 0xe5, 0x8b,
	0xf9, 0xd0, 0xc3, 0x43, 0x6f, 0x1d, 0x7a, 0xfd, 0xff, 0x2b, 0xd6, 0x93, 0xc1, 0xdf, 0xfd, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x89, 0x9c, 0x0a, 0xd3, 0x21, 0x01, 0x00, 0x00,
}
