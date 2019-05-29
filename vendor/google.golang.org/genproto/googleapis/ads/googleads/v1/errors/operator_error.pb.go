// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/operator_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible operator errors.
type OperatorErrorEnum_OperatorError int32

const (
	// Enum unspecified.
	OperatorErrorEnum_UNSPECIFIED OperatorErrorEnum_OperatorError = 0
	// The received error code is not known in this version.
	OperatorErrorEnum_UNKNOWN OperatorErrorEnum_OperatorError = 1
	// Operator not supported.
	OperatorErrorEnum_OPERATOR_NOT_SUPPORTED OperatorErrorEnum_OperatorError = 2
)

var OperatorErrorEnum_OperatorError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPERATOR_NOT_SUPPORTED",
}
var OperatorErrorEnum_OperatorError_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"OPERATOR_NOT_SUPPORTED": 2,
}

func (x OperatorErrorEnum_OperatorError) String() string {
	return proto.EnumName(OperatorErrorEnum_OperatorError_name, int32(x))
}
func (OperatorErrorEnum_OperatorError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_operator_error_87b1e31e98dffa4f, []int{0, 0}
}

// Container for enum describing possible operator errors.
type OperatorErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorErrorEnum) Reset()         { *m = OperatorErrorEnum{} }
func (m *OperatorErrorEnum) String() string { return proto.CompactTextString(m) }
func (*OperatorErrorEnum) ProtoMessage()    {}
func (*OperatorErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_operator_error_87b1e31e98dffa4f, []int{0}
}
func (m *OperatorErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorErrorEnum.Unmarshal(m, b)
}
func (m *OperatorErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorErrorEnum.Marshal(b, m, deterministic)
}
func (dst *OperatorErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorErrorEnum.Merge(dst, src)
}
func (m *OperatorErrorEnum) XXX_Size() int {
	return xxx_messageInfo_OperatorErrorEnum.Size(m)
}
func (m *OperatorErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OperatorErrorEnum)(nil), "google.ads.googleads.v1.errors.OperatorErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.OperatorErrorEnum_OperatorError", OperatorErrorEnum_OperatorError_name, OperatorErrorEnum_OperatorError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/operator_error.proto", fileDescriptor_operator_error_87b1e31e98dffa4f)
}

var fileDescriptor_operator_error_87b1e31e98dffa4f = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0xbf, 0xce, 0x07, 0x0a, 0x29, 0x62, 0x9d, 0x85, 0x8b, 0x22, 0x5d, 0xcc, 0x03, 0x24,
	0x0c, 0xdd, 0xc5, 0x55, 0x6a, 0x63, 0x29, 0xc2, 0x24, 0xf4, 0x9f, 0x20, 0x83, 0x25, 0x3a, 0x43,
	0x18, 0x68, 0x73, 0x87, 0x64, 0xec, 0x03, 0xb9, 0xf4, 0x51, 0x7c, 0x14, 0xc1, 0x77, 0x90, 0x99,
	0x74, 0x06, 0xba, 0xd0, 0x55, 0x4e, 0x2e, 0xbf, 0x73, 0xee, 0xe1, 0xa2, 0xb1, 0x06, 0xd0, 0xbb,
	0x9c, 0xa8, 0xcc, 0x11, 0x2f, 0x6b, 0x75, 0x88, 0x49, 0x6e, 0x2d, 0x58, 0x47, 0xa0, 0xcc, 0xad,
	0xaa, 0xc0, 0x6e, 0x9b, 0x3f, 0x2e, 0x2d, 0x54, 0x10, 0x8e, 0x3c, 0x89, 0x55, 0xe6, 0x70, 0x67,
	0xc2, 0x87, 0x18, 0x7b, 0xd3, 0xf0, 0xa6, 0x0d, 0x2d, 0x0b, 0xa2, 0x8c, 0x81, 0x4a, 0x55, 0x05,
	0x18, 0xe7, 0xdd, 0xd1, 0x33, 0xba, 0x12, 0xc7, 0x54, 0x5e, 0xf3, 0xdc, 0xbc, 0xed, 0xa3, 0x39,
	0xba, 0x38, 0x19, 0x86, 0x97, 0xa8, 0xbf, 0x4e, 0x96, 0x92, 0xdf, 0xcd, 0xef, 0xe7, 0x7c, 0x3a,
	0xf8, 0x17, 0xf6, 0xd1, 0xf9, 0x3a, 0x79, 0x48, 0xc4, 0x63, 0x32, 0xe8, 0x85, 0x43, 0x74, 0x2d,
	0x24, 0x5f, 0xb0, 0x95, 0x58, 0x6c, 0x13, 0xb1, 0xda, 0x2e, 0xd7, 0x52, 0x8a, 0xc5, 0x8a, 0x4f,
	0x07, 0xc1, 0xe4, 0xbb, 0x87, 0xa2, 0x57, 0xd8, 0xe3, 0xbf, 0x4b, 0x4e, 0xc2, 0x93, 0x7d, 0xb2,
	0xae, 0x26, 0x7b, 0x4f, 0xd3, 0xa3, 0x4b, 0xc3, 0x4e, 0x19, 0x8d, 0xc1, 0x6a, 0xa2, 0x73, 0xd3,
	0x14, 0x6f, 0xef, 0x53, 0x16, 0xee, 0xb7, 0x73, 0xdd, 0xfa, 0xe7, 0x3d, 0xf8, 0x3f, 0x63, 0xec,
	0x23, 0x18, 0xcd, 0x7c, 0x18, 0xcb, 0x1c, 0xf6, 0xb2, 0x56, 0x9b, 0x18, 0x37, 0x2b, 0xdd, 0x67,
	0x0b, 0xa4, 0x2c, 0x73, 0x69, 0x07, 0xa4, 0x9b, 0x38, 0xf5, 0xc0, 0x57, 0x10, 0xf9, 0x29, 0xa5,
	0x2c, 0x73, 0x94, 0x76, 0x08, 0xa5, 0x9b, 0x98, 0x52, 0x0f, 0xbd, 0x9c, 0x35, 0xed, 0xc6, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x09, 0x5e, 0xde, 0xc0, 0xcb, 0x01, 0x00, 0x00,
}