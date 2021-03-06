// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/gov/permission.proto

package gov

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type PermValue int32

const (
	// PERMISSION_ZERO is a no-op permission.
	PermValue_PERMISSION_ZERO PermValue = 0
	// PERMISSION_SET_PERMISSIONS defines the permission that allows to Set Permissions to other actors.
	PermValue_PERMISSION_SET_PERMISSIONS PermValue = 1
	// PERMISSION_CLAIM_VALIDATOR defines the permission that allows to Claim a validator Seat.
	PermValue_PERMISSION_CLAIM_VALIDATOR PermValue = 2
	// PERMISSION_CLAIM_COUNCILOR defines the permission that allows to Claim a Councilor Seat.
	PermValue_PERMISSION_CLAIM_COUNCILOR PermValue = 3
	// PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL defines the permission needed to create proposals for setting permissions.
	PermValue_PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL PermValue = 4
	// PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set permissions.
	PermValue_PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL PermValue = 5
	//  PERMISSION_UPSERT_TOKEN_ALIAS
	PermValue_PERMISSION_UPSERT_TOKEN_ALIAS PermValue = 6
	// PERMISSION_CHANGE_TX_FEE
	PermValue_PERMISSION_CHANGE_TX_FEE PermValue = 7
	// PERMISSION_UPSERT_TOKEN_RATE
	PermValue_PERMISSION_UPSERT_TOKEN_RATE PermValue = 8
	// PERMISSION_UPSERT_ROLE makes possible to add, modify and assign roles.
	PermValue_PERMISSION_UPSERT_ROLE PermValue = 9
	// PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermValue_PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL PermValue = 10
	// PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermValue_PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL PermValue = 11
	// PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission needed to create proposals for setting network property.
	PermValue_PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL PermValue = 12
	// PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set network property.
	PermValue_PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL PermValue = 13
	// PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to create proposals for upsert token Alias.
	PermValue_PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL PermValue = 14
	// PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to vote proposals for upsert token.
	PermValue_PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL PermValue = 15
	// PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES defines the permission needed to create proposals for setting poor network messages
	PermValue_PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES PermValue = 16
	// PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL defines the permission needed to vote proposals to set poor network messages
	PermValue_PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL PermValue = 17
	// PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to create proposals for upsert token rate.
	PermValue_PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL PermValue = 18
	// PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to vote proposals for upsert token rate.
	PermValue_PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL PermValue = 19
	// PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to create a proposal to unjail a validator.
	PermValue_PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL PermValue = 20
	// PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to vote a proposal to unjail a validator.
	PermValue_PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL PermValue = 21
	// PERMISSION_CREATE_CREATE_ROLE_PROPOSAL defines the permission needed to create a proposal to create a role.
	PermValue_PERMISSION_CREATE_CREATE_ROLE_PROPOSAL PermValue = 22
	// PERMISSION_VOTE_CREATE_ROLE_PROPOSAL defines the permission needed to vote a proposal to create a role.
	PermValue_PERMISSION_VOTE_CREATE_ROLE_PROPOSAL PermValue = 23
	// PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL defines the permission needed to create a proposal to blacklist/whitelisted tokens
	PermValue_PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL PermValue = 24
	// PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL defines the permission needed to vote on blacklist/whitelisted tokens proposal
	PermValue_PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL PermValue = 25
	// PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL defines the permission needed to create a proposal to reset whole validator rank
	PermValue_PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL PermValue = 26
	// PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL defines the permission needed to vote on reset whole validator rank proposal
	PermValue_PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL PermValue = 27
	// PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL defines the permission needed to create a proposal for software upgrade
	PermValue_PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL PermValue = 28
	// PERMISSION_SOFTWARE_UPGRADE_PROPOSAL defines the permission needed to vote on software upgrade proposal
	PermValue_PERMISSION_SOFTWARE_UPGRADE_PROPOSAL PermValue = 29
	// PERMISSION_SET_PERMISSIONS defines the permission that allows to Set ClaimValidatorPermission to other actors.
	PermValue_PERMISSION_SET_CLAIM_VALIDATOR_PERMISSION PermValue = 30
)

var PermValue_name = map[int32]string{
	0:  "PERMISSION_ZERO",
	1:  "PERMISSION_SET_PERMISSIONS",
	2:  "PERMISSION_CLAIM_VALIDATOR",
	3:  "PERMISSION_CLAIM_COUNCILOR",
	4:  "PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL",
	5:  "PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL",
	6:  "PERMISSION_UPSERT_TOKEN_ALIAS",
	7:  "PERMISSION_CHANGE_TX_FEE",
	8:  "PERMISSION_UPSERT_TOKEN_RATE",
	9:  "PERMISSION_UPSERT_ROLE",
	10: "PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL",
	11: "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL",
	12: "PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL",
	13: "PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL",
	14: "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL",
	15: "PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL",
	16: "PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES",
	17: "PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL",
	18: "PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL",
	19: "PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL",
	20: "PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL",
	21: "PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL",
	22: "PERMISSION_CREATE_CREATE_ROLE_PROPOSAL",
	23: "PERMISSION_VOTE_CREATE_ROLE_PROPOSAL",
	24: "PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL",
	25: "PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL",
	26: "PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL",
	27: "PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL",
	28: "PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL",
	29: "PERMISSION_SOFTWARE_UPGRADE_PROPOSAL",
	30: "PERMISSION_SET_CLAIM_VALIDATOR_PERMISSION",
}

var PermValue_value = map[string]int32{
	"PERMISSION_ZERO":                                       0,
	"PERMISSION_SET_PERMISSIONS":                            1,
	"PERMISSION_CLAIM_VALIDATOR":                            2,
	"PERMISSION_CLAIM_COUNCILOR":                            3,
	"PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL":            4,
	"PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL":              5,
	"PERMISSION_UPSERT_TOKEN_ALIAS":                         6,
	"PERMISSION_CHANGE_TX_FEE":                              7,
	"PERMISSION_UPSERT_TOKEN_RATE":                          8,
	"PERMISSION_UPSERT_ROLE":                                9,
	"PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL":       10,
	"PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL":         11,
	"PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL":       12,
	"PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL":         13,
	"PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL":         14,
	"PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL":           15,
	"PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES":           16,
	"PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL":    17,
	"PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL":          18,
	"PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL":            19,
	"PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL":           20,
	"PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL":             21,
	"PERMISSION_CREATE_CREATE_ROLE_PROPOSAL":                22,
	"PERMISSION_VOTE_CREATE_ROLE_PROPOSAL":                  23,
	"PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL":  24,
	"PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL":    25,
	"PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL": 26,
	"PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL":   27,
	"PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL":           28,
	"PERMISSION_SOFTWARE_UPGRADE_PROPOSAL":                  29,
	"PERMISSION_SET_CLAIM_VALIDATOR_PERMISSION":             30,
}

func (x PermValue) String() string {
	return proto.EnumName(PermValue_name, int32(x))
}

func (PermValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_214168f8815c1062, []int{0}
}

func init() {
	proto.RegisterEnum("kira.gov.PermValue", PermValue_name, PermValue_value)
}

func init() {
	proto.RegisterFile("kira/gov/permission.proto", fileDescriptor_214168f8815c1062)
}

var fileDescriptor_214168f8815c1062 = []byte{
	// 930 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdf, 0x52, 0xdb, 0x46,
	0x14, 0xc6, 0xd3, 0x36, 0x4d, 0xc9, 0x36, 0x0d, 0xaa, 0xa0, 0x84, 0x6c, 0x03, 0x5d, 0x5a, 0x4a,
	0x09, 0x09, 0x78, 0x5a, 0x9a, 0xce, 0x74, 0x7a, 0x25, 0xcc, 0x02, 0x2a, 0xb6, 0xe5, 0x59, 0xc9,
	0x38, 0x61, 0x26, 0xa3, 0x59, 0x60, 0x23, 0xab, 0x36, 0x5e, 0xcf, 0x4a, 0x24, 0xed, 0x1b, 0x74,
	0xf6, 0x1d, 0xf6, 0xaa, 0x6f, 0xd3, 0xa7, 0xca, 0x48, 0x60, 0xad, 0x24, 0xcb, 0x86, 0x2b, 0xb0,
	0xa4, 0xf3, 0x3b, 0xdf, 0x7e, 0xe7, 0xcf, 0x2c, 0x78, 0xda, 0x0f, 0x05, 0xad, 0x05, 0xfc, 0x7d,
	0x6d, 0xc4, 0xc4, 0x65, 0x18, 0x45, 0x21, 0x1f, 0xee, 0x8c, 0x04, 0x8f, 0xb9, 0x39, 0x97, 0xbc,
	0xda, 0x09, 0xf8, 0x7b, 0xb8, 0x18, 0xf0, 0x80, 0xa7, 0x0f, 0x6b, 0xc9, 0x7f, 0xd7, 0xef, 0xb7,
	0xfe, 0x5f, 0x00, 0x0f, 0xdb, 0x4c, 0x5c, 0x9e, 0xd0, 0xc1, 0x15, 0x33, 0xd7, 0xc0, 0x7c, 0x1b,
	0x93, 0xa6, 0xed, 0xba, 0xb6, 0xd3, 0xf2, 0x4f, 0x31, 0x71, 0x8c, 0x7b, 0xf0, 0x91, 0x54, 0x68,
	0x2e, 0xf9, 0xe6, 0x94, 0x09, 0x6e, 0xfe, 0x06, 0x60, 0xee, 0x13, 0x17, 0x7b, 0xbe, 0xfe, 0xe9,
	0x1a, 0x9f, 0xc0, 0x25, 0xa9, 0x90, 0x99, 0x7c, 0xed, 0xb2, 0xb8, 0x9d, 0xa9, 0x89, 0x4a, 0x71,
	0xf5, 0x86, 0x65, 0x37, 0xfd, 0x13, 0xab, 0x61, 0xef, 0x5b, 0x9e, 0x43, 0x8c, 0x4f, 0x75, 0x5c,
	0x7d, 0x40, 0xc3, 0x44, 0x4e, 0x78, 0x41, 0x63, 0x2e, 0x2a, 0xe3, 0xea, 0x4e, 0xa7, 0x55, 0xb7,
	0x1b, 0x0e, 0x31, 0x3e, 0x2b, 0xc5, 0xd5, 0xf9, 0xd5, 0xf0, 0x3c, 0x1c, 0x70, 0x61, 0x7a, 0x60,
	0x2b, 0x1f, 0x47, 0xb0, 0xe5, 0xe1, 0xb2, 0x5c, 0xbf, 0x4d, 0x9c, 0xb6, 0xe3, 0x5a, 0x0d, 0xe3,
	0x3e, 0x5c, 0x97, 0x0a, 0xa1, 0x94, 0x23, 0x18, 0x8d, 0x59, 0x51, 0x7d, 0x5b, 0xf0, 0x11, 0x8f,
	0xe8, 0xc0, 0x74, 0xc0, 0x66, 0x8e, 0x7a, 0xe2, 0xcc, 0x62, 0x7e, 0x0e, 0xd7, 0xa4, 0x42, 0x2b,
	0xa9, 0xbb, 0xbc, 0x44, 0xcc, 0x80, 0x7f, 0x80, 0x95, 0x1c, 0xb0, 0xd3, 0x76, 0x31, 0xf1, 0x7c,
	0xcf, 0x39, 0xc6, 0x2d, 0xdf, 0x6a, 0xd8, 0x96, 0x6b, 0x3c, 0x80, 0xcb, 0x52, 0xa1, 0xc5, 0x24,
	0xb4, 0x33, 0x8a, 0x98, 0x88, 0x3d, 0xde, 0x67, 0x43, 0x6b, 0x10, 0xd2, 0xc8, 0xfc, 0x19, 0x2c,
	0xe7, 0xcf, 0x78, 0x64, 0xb5, 0x0e, 0xb1, 0xef, 0xbd, 0xf6, 0x0f, 0x30, 0x36, 0xbe, 0x80, 0x0b,
	0x52, 0xa1, 0xf9, 0xf4, 0x44, 0x3d, 0x3a, 0x0c, 0x98, 0xf7, 0xf7, 0x01, 0x63, 0xe6, 0xef, 0xe0,
	0xd9, 0xb4, 0x7c, 0xc4, 0xf2, 0xb0, 0x31, 0x07, 0x9f, 0x48, 0x85, 0x16, 0x4a, 0xe9, 0x08, 0x8d,
	0x99, 0xb9, 0x03, 0x96, 0x26, 0x43, 0x89, 0xd3, 0xc0, 0xc6, 0x43, 0x68, 0x4a, 0x85, 0x1e, 0xeb,
	0x20, 0xc2, 0x07, 0xcc, 0x7c, 0x0b, 0x6a, 0x93, 0x15, 0xb8, 0x09, 0xdb, 0xb7, 0x3c, 0xcb, 0x27,
	0xf8, 0xd0, 0x76, 0x3d, 0xf2, 0x46, 0x5b, 0x06, 0xe0, 0xa6, 0x54, 0x68, 0x5d, 0x97, 0xe1, 0x1a,
	0xb7, 0x4f, 0x63, 0x4a, 0x58, 0x10, 0x46, 0xb1, 0xf8, 0x27, 0x73, 0xee, 0x0d, 0xd8, 0x2e, 0x97,
	0x62, 0x36, 0xfc, 0x4b, 0xb8, 0x21, 0x15, 0xfa, 0x7e, 0x5c, 0x8f, 0x19, 0xe8, 0x4a, 0xe5, 0x49,
	0x9d, 0x5b, 0xd8, 0xeb, 0x3a, 0xe4, 0x38, 0x65, 0x62, 0xe2, 0xe5, 0xe0, 0x8f, 0xca, 0xca, 0x5d,
	0x16, 0xb7, 0x58, 0xfc, 0x81, 0x8b, 0x7e, 0x82, 0x65, 0x22, 0x9e, 0xa9, 0x7c, 0x36, 0xfc, 0xab,
	0xa2, 0xf2, 0x3b, 0xa3, 0x8b, 0x9e, 0xe7, 0xba, 0x4a, 0xa3, 0x1f, 0x6b, 0x74, 0xde, 0x71, 0xdd,
	0x64, 0x19, 0xba, 0x03, 0x5e, 0x4c, 0xf1, 0xbb, 0x12, 0x3c, 0xaf, 0x27, 0x4a, 0xbb, 0x5d, 0x81,
	0x7d, 0x5b, 0xc0, 0xe6, 0xe7, 0xd4, 0x71, 0x48, 0xe6, 0x49, 0x13, 0xbb, 0xae, 0x75, 0x88, 0x5d,
	0xc3, 0x80, 0x2f, 0xa5, 0x42, 0x9b, 0xc5, 0x41, 0xe5, 0x5c, 0xdc, 0x18, 0xd2, 0x64, 0x51, 0x44,
	0x03, 0xa6, 0xf1, 0x67, 0xe0, 0x97, 0xca, 0x81, 0xad, 0x82, 0x6b, 0xf1, 0x5f, 0xc3, 0x2d, 0xa9,
	0xd0, 0x46, 0x7e, 0x74, 0x67, 0xe4, 0xe8, 0x82, 0x97, 0xb7, 0x98, 0x9e, 0x8c, 0x96, 0xa6, 0x9b,
	0xf0, 0x47, 0xa9, 0xd0, 0x5a, 0xa5, 0xe7, 0xc9, 0xa4, 0x65, 0x60, 0xb7, 0xb0, 0xc3, 0x26, 0x2d,
	0x2f, 0x62, 0x17, 0xe0, 0x0f, 0x52, 0xa1, 0xef, 0x2a, 0x1c, 0x2f, 0x40, 0x4f, 0xaa, 0x0c, 0xef,
	0xb4, 0xfe, 0xb4, 0xec, 0x86, 0x5e, 0xc8, 0x9a, 0xba, 0x38, 0x21, 0x76, 0xf8, 0x17, 0x0d, 0x07,
	0xd9, 0x82, 0xce, 0xb8, 0x04, 0x3c, 0x9f, 0x10, 0x3b, 0x95, 0xfa, 0x4d, 0x49, 0xeb, 0x14, 0xe6,
	0x01, 0xd8, 0x98, 0xd4, 0x7a, 0xf3, 0x27, 0xd9, 0x3c, 0x1a, 0xb8, 0x04, 0xa1, 0x54, 0x68, 0x49,
	0xcb, 0x4c, 0x56, 0x50, 0xc6, 0x39, 0x02, 0xeb, 0x65, 0x6d, 0x95, 0x94, 0x27, 0x70, 0x55, 0x2a,
	0x04, 0xc7, 0xb2, 0x2a, 0x48, 0xef, 0xc0, 0xaf, 0x93, 0x8a, 0xd2, 0x6a, 0xb8, 0x7e, 0xf7, 0xc8,
	0xf6, 0xb0, 0xbf, 0xd7, 0xb0, 0xea, 0xc7, 0xe3, 0x65, 0x9c, 0x91, 0x97, 0xcb, 0x7d, 0x9b, 0x16,
	0x26, 0xea, 0xf6, 0xc2, 0x98, 0xed, 0x0d, 0xe8, 0x79, 0xff, 0x7a, 0x49, 0xcf, 0xea, 0xdb, 0x3b,
	0x64, 0x79, 0x5a, 0xec, 0xdb, 0x5b, 0x72, 0xf4, 0xc0, 0xab, 0xc9, 0xb3, 0x10, 0x9c, 0xcc, 0x47,
	0xf7, 0x28, 0xf1, 0x45, 0x17, 0x8e, 0x58, 0xad, 0x63, 0x9d, 0x06, 0xc2, 0x6d, 0xa9, 0xd0, 0xf3,
	0x9c, 0xd9, 0x2c, 0x62, 0x71, 0xb7, 0xc7, 0x07, 0x2c, 0xab, 0x21, 0xa1, 0xc3, 0x7e, 0x96, 0xe9,
	0x02, 0xec, 0x96, 0x4f, 0x73, 0x97, 0x3c, 0xdf, 0xc2, 0x17, 0x52, 0xa1, 0x9f, 0xc6, 0xc7, 0xb9,
	0x2d, 0x4b, 0x65, 0x67, 0xbb, 0xce, 0x81, 0xd7, 0xb5, 0x48, 0x32, 0x39, 0x87, 0xc4, 0xda, 0xcf,
	0x99, 0xf5, 0xac, 0xdc, 0xd9, 0x2e, 0x7f, 0x17, 0x7f, 0xa0, 0x82, 0x75, 0x46, 0x81, 0xa0, 0x17,
	0xda, 0xa7, 0x66, 0xa1, 0x7b, 0xa6, 0x03, 0x57, 0x8a, 0x4d, 0x3d, 0x0d, 0x57, 0x1c, 0x94, 0xc4,
	0x85, 0xd2, 0x6d, 0x28, 0x77, 0x9d, 0x30, 0x56, 0x35, 0xd3, 0x65, 0x71, 0xf1, 0x6e, 0xa4, 0xaf,
	0x13, 0xf0, 0xfe, 0xbf, 0xff, 0xad, 0xde, 0xdb, 0x7b, 0x75, 0xba, 0x1b, 0x84, 0x71, 0xef, 0xea,
	0x6c, 0xe7, 0x9c, 0x5f, 0xd6, 0x8e, 0x43, 0x41, 0xeb, 0x5c, 0xb0, 0x5a, 0xc4, 0xfa, 0x34, 0xac,
	0xd9, 0x2d, 0x0f, 0x93, 0xd7, 0xb5, 0xf4, 0xda, 0xb7, 0x1d, 0xb0, 0x61, 0x6d, 0x7c, 0x69, 0x3c,
	0x7b, 0x90, 0x3e, 0xdb, 0xfd, 0x18, 0x00, 0x00, 0xff, 0xff, 0x02, 0xf9, 0xd6, 0x59, 0x47, 0x0a,
	0x00, 0x00,
}
