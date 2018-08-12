// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Common/Common.proto

package Common

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

// 返回结果
type RetType int32

const (
	RetType_RetType_Succeed RetType = 0
	RetType_RetType_Failed  RetType = -1
	RetType_RetType_TimeOut RetType = -100
	RetType_RetType_Unknown RetType = -400
)

var RetType_name = map[int32]string{
	0:    "RetType_Succeed",
	-1:   "RetType_Failed",
	-100: "RetType_TimeOut",
	-400: "RetType_Unknown",
}
var RetType_value = map[string]int32{
	"RetType_Succeed": 0,
	"RetType_Failed":  -1,
	"RetType_TimeOut": -100,
	"RetType_Unknown": -400,
}

func (x RetType) Enum() *RetType {
	p := new(RetType)
	*p = x
	return p
}
func (x RetType) String() string {
	return proto.EnumName(RetType_name, int32(x))
}
func (x *RetType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RetType_value, data, "RetType")
	if err != nil {
		return err
	}
	*x = RetType(value)
	return nil
}
func (RetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Common_ce453695c72a9286, []int{0}
}

// 包的唯一标识，用于回放攻击的识别和保护
type PacketID struct {
	ConnID               *uint64  `protobuf:"varint,1,req,name=connID" json:"connID,omitempty"`
	SerialNo             *uint32  `protobuf:"varint,2,req,name=serialNo" json:"serialNo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketID) Reset()         { *m = PacketID{} }
func (m *PacketID) String() string { return proto.CompactTextString(m) }
func (*PacketID) ProtoMessage()    {}
func (*PacketID) Descriptor() ([]byte, []int) {
	return fileDescriptor_Common_ce453695c72a9286, []int{0}
}
func (m *PacketID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketID.Unmarshal(m, b)
}
func (m *PacketID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketID.Marshal(b, m, deterministic)
}
func (dst *PacketID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketID.Merge(dst, src)
}
func (m *PacketID) XXX_Size() int {
	return xxx_messageInfo_PacketID.Size(m)
}
func (m *PacketID) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketID.DiscardUnknown(m)
}

var xxx_messageInfo_PacketID proto.InternalMessageInfo

func (m *PacketID) GetConnID() uint64 {
	if m != nil && m.ConnID != nil {
		return *m.ConnID
	}
	return 0
}

func (m *PacketID) GetSerialNo() uint32 {
	if m != nil && m.SerialNo != nil {
		return *m.SerialNo
	}
	return 0
}

func init() {
	proto.RegisterType((*PacketID)(nil), "Common.PacketID")
	proto.RegisterEnum("Common.RetType", RetType_name, RetType_value)
}

func init() { proto.RegisterFile("Common/Common.proto", fileDescriptor_Common_ce453695c72a9286) }

var fileDescriptor_Common_ce453695c72a9286 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x76, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92,
	0x1d, 0x17, 0x47, 0x40, 0x62, 0x72, 0x76, 0x6a, 0x89, 0xa7, 0x8b, 0x90, 0x18, 0x17, 0x5b, 0x72,
	0x7e, 0x5e, 0x9e, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0x93, 0x06, 0x4b, 0x10, 0x94, 0x27, 0x24, 0xc5,
	0xc5, 0x51, 0x9c, 0x5a, 0x94, 0x99, 0x98, 0xe3, 0x97, 0x2f, 0xc1, 0xa4, 0xc0, 0xa4, 0xc1, 0x1b,
	0x04, 0xe7, 0x6b, 0x95, 0x73, 0xb1, 0x07, 0xa5, 0x96, 0x84, 0x54, 0x16, 0xa4, 0x0a, 0x09, 0x73,
	0xf1, 0x43, 0x99, 0xf1, 0xc1, 0xa5, 0xc9, 0xc9, 0xa9, 0xa9, 0x29, 0x02, 0x0c, 0x42, 0xd2, 0x5c,
	0x7c, 0x30, 0x41, 0xb7, 0xc4, 0xcc, 0x9c, 0xd4, 0x14, 0x81, 0xff, 0x30, 0xc0, 0x28, 0x24, 0x83,
	0xd0, 0x11, 0x92, 0x99, 0x9b, 0xea, 0x5f, 0x5a, 0x22, 0x30, 0x07, 0xab, 0x6c, 0x68, 0x5e, 0x76,
	0x5e, 0x7e, 0x79, 0x9e, 0xc0, 0x87, 0x3f, 0x30, 0x59, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48,
	0x99, 0x96, 0xe6, 0xd6, 0x00, 0x00, 0x00,
}
