// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_UpdateTicker/Qot_UpdateTicker.proto

package Qot_UpdateTicker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "limgo/futupb/Common"
import Qot_Common "limgo/futupb/Qot_Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type S2C struct {
	Security             *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	TickerList           []*Qot_Common.Ticker `protobuf:"bytes,2,rep,name=tickerList" json:"tickerList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_UpdateTicker_232fb2cd3e89c536, []int{0}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *S2C) GetTickerList() []*Qot_Common.Ticker {
	if m != nil {
		return m.TickerList
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_UpdateTicker_232fb2cd3e89c536, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*S2C)(nil), "Qot_UpdateTicker.S2C")
	proto.RegisterType((*Response)(nil), "Qot_UpdateTicker.Response")
}

func init() {
	proto.RegisterFile("Qot_UpdateTicker/Qot_UpdateTicker.proto", fileDescriptor_Qot_UpdateTicker_232fb2cd3e89c536)
}

var fileDescriptor_Qot_UpdateTicker_232fb2cd3e89c536 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xe5, 0xa4, 0xa5, 0xe5, 0xb2, 0xa0, 0xe3, 0x8f, 0xac, 0x22, 0x21, 0xab, 0x4b, 0xbd,
	0x90, 0x46, 0x16, 0x13, 0x6b, 0x56, 0x18, 0x70, 0xca, 0x8c, 0x50, 0x7a, 0x42, 0x51, 0xd5, 0x3a,
	0xb2, 0x8f, 0xa1, 0x1f, 0x80, 0xef, 0x8d, 0x5a, 0xa7, 0x10, 0xa5, 0x93, 0xfd, 0xde, 0xfb, 0x9d,
	0xde, 0xe9, 0x60, 0xf1, 0xe6, 0xf8, 0xe3, 0xbd, 0x5d, 0x7f, 0x32, 0xad, 0x9a, 0x7a, 0x43, 0x7e,
	0x39, 0x34, 0xf2, 0xd6, 0x3b, 0x76, 0x78, 0x35, 0xf4, 0x67, 0xd7, 0xa5, 0xdb, 0x6e, 0xdd, 0x6e,
	0x19, 0x9f, 0x88, 0xcd, 0xee, 0x0f, 0x58, 0x17, 0xfc, 0x7f, 0x63, 0x38, 0xdf, 0x40, 0x5a, 0x99,
	0x12, 0x0b, 0x98, 0x06, 0xaa, 0xbf, 0x7d, 0xc3, 0x7b, 0x29, 0x54, 0xa2, 0x33, 0x73, 0x93, 0xf7,
	0xd8, 0xaa, 0xcb, 0xec, 0x1f, 0x85, 0x06, 0x80, 0x8f, 0xa5, 0x2f, 0x4d, 0x60, 0x99, 0xa8, 0x54,
	0x67, 0x06, 0xfb, 0x33, 0x71, 0x25, 0xdb, 0xa3, 0xe6, 0x3f, 0x02, 0xa6, 0x96, 0x42, 0xeb, 0x76,
	0x81, 0xf0, 0x01, 0x26, 0x9e, 0x78, 0xb5, 0x6f, 0xe9, 0xd8, 0x38, 0x7e, 0x1e, 0x3d, 0x3e, 0x15,
	0x85, 0x3d, 0x99, 0x78, 0x07, 0x17, 0x9e, 0xf8, 0x35, 0x7c, 0xc9, 0x44, 0x09, 0x7d, 0x69, 0x3b,
	0x85, 0x12, 0x26, 0xe4, 0x7d, 0xe9, 0xd6, 0x24, 0x53, 0x25, 0xf4, 0xd8, 0x9e, 0x24, 0x2e, 0x20,
	0x0d, 0xa6, 0x96, 0x23, 0x25, 0x74, 0x66, 0x6e, 0xf3, 0xb3, 0xab, 0x55, 0xa6, 0xb4, 0x07, 0xe2,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x35, 0x41, 0x04, 0x93, 0x62, 0x01, 0x00, 0x00,
}
