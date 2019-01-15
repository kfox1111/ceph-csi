// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/feed_status.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible statuses of a feed.
type FeedStatusEnum_FeedStatus int32

const (
	// Not specified.
	FeedStatusEnum_UNSPECIFIED FeedStatusEnum_FeedStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedStatusEnum_UNKNOWN FeedStatusEnum_FeedStatus = 1
	// Feed is enabled.
	FeedStatusEnum_ENABLED FeedStatusEnum_FeedStatus = 2
	// Feed has been removed.
	FeedStatusEnum_REMOVED FeedStatusEnum_FeedStatus = 3
)

var FeedStatusEnum_FeedStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}
var FeedStatusEnum_FeedStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x FeedStatusEnum_FeedStatus) String() string {
	return proto.EnumName(FeedStatusEnum_FeedStatus_name, int32(x))
}
func (FeedStatusEnum_FeedStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_feed_status_dd159d67991f6ccc, []int{0, 0}
}

// Container for enum describing possible statuses of a feed.
type FeedStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedStatusEnum) Reset()         { *m = FeedStatusEnum{} }
func (m *FeedStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedStatusEnum) ProtoMessage()    {}
func (*FeedStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_status_dd159d67991f6ccc, []int{0}
}
func (m *FeedStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedStatusEnum.Unmarshal(m, b)
}
func (m *FeedStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedStatusEnum.Marshal(b, m, deterministic)
}
func (dst *FeedStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedStatusEnum.Merge(dst, src)
}
func (m *FeedStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedStatusEnum.Size(m)
}
func (m *FeedStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FeedStatusEnum)(nil), "google.ads.googleads.v0.enums.FeedStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.FeedStatusEnum_FeedStatus", FeedStatusEnum_FeedStatus_name, FeedStatusEnum_FeedStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/feed_status.proto", fileDescriptor_feed_status_dd159d67991f6ccc)
}

var fileDescriptor_feed_status_dd159d67991f6ccc = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0x86, 0x32, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc, 0xd2,
	0xdc, 0x62, 0xfd, 0xb4, 0xd4, 0xd4, 0x94, 0xf8, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0xbd, 0x82,
	0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x2a, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x06, 0xbd,
	0x32, 0x03, 0x3d, 0xb0, 0x06, 0xa5, 0x30, 0x2e, 0x3e, 0xb7, 0xd4, 0xd4, 0x94, 0x60, 0xb0, 0x16,
	0xd7, 0xbc, 0xd2, 0x5c, 0x25, 0x17, 0x2e, 0x2e, 0x84, 0x88, 0x10, 0x3f, 0x17, 0x77, 0xa8, 0x5f,
	0x70, 0x80, 0xab, 0xb3, 0xa7, 0x9b, 0xa7, 0xab, 0x8b, 0x00, 0x83, 0x10, 0x37, 0x17, 0x7b, 0xa8,
	0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0x23, 0x88, 0xe3, 0xea, 0xe7, 0xe8, 0xe4, 0xe3, 0xea,
	0x22, 0xc0, 0x04, 0xe2, 0x04, 0xb9, 0xfa, 0xfa, 0x87, 0xb9, 0xba, 0x08, 0x30, 0x3b, 0x1d, 0x60,
	0xe4, 0x52, 0x4c, 0xce, 0xcf, 0xd5, 0xc3, 0x6b, 0xbb, 0x13, 0x3f, 0xc2, 0xa6, 0x00, 0x90, 0x6b,
	0x03, 0x18, 0xa3, 0x9c, 0xa0, 0x3a, 0xd2, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2,
	0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0x7e, 0x81, 0x79, 0xb8, 0x20, 0xb3, 0x18, 0x87, 0xff, 0xad, 0xc1,
	0xe4, 0x22, 0x26, 0x66, 0x77, 0x47, 0xc7, 0x55, 0x4c, 0xb2, 0xee, 0x10, 0xa3, 0x1c, 0x53, 0x8a,
	0xf5, 0x20, 0x4c, 0x10, 0x2b, 0xcc, 0x40, 0x0f, 0xe4, 0xcf, 0xe2, 0x53, 0x30, 0xf9, 0x18, 0xc7,
	0x94, 0xe2, 0x18, 0xb8, 0x7c, 0x4c, 0x98, 0x41, 0x0c, 0x58, 0x3e, 0x89, 0x0d, 0x6c, 0xa9, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x60, 0x36, 0x0a, 0x4e, 0x73, 0x01, 0x00, 0x00,
}