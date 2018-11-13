// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package request

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ethAddress common.Address, paymentMethod uint8, paymentNumber, chargeAmountInWei *big.Int
type SignRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Number               string   `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	ChargeAmount         string   `protobuf:"bytes,4,opt,name=chargeAmount,proto3" json:"chargeAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRequest) Reset()         { *m = SignRequest{} }
func (m *SignRequest) String() string { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()    {}
func (*SignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{0}
}

func (m *SignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRequest.Unmarshal(m, b)
}
func (m *SignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRequest.Marshal(b, m, deterministic)
}
func (m *SignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRequest.Merge(m, src)
}
func (m *SignRequest) XXX_Size() int {
	return xxx_messageInfo_SignRequest.Size(m)
}
func (m *SignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignRequest proto.InternalMessageInfo

func (m *SignRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SignRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *SignRequest) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *SignRequest) GetChargeAmount() string {
	if m != nil {
		return m.ChargeAmount
	}
	return ""
}

func init() {
	proto.RegisterType((*SignRequest)(nil), "request.SignRequest")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_7f73548e33e655fe) }

var fileDescriptor_7f73548e33e655fe = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xce, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0xc6, 0xf1, 0xa0, 0x06, 0x62, 0xd5, 0xa5, 0x83, 0xe9, 0x68, 0x98, 0x8c, 0x31, 0x74, 0xf0,
	0x09, 0x74, 0x76, 0x42, 0x27, 0xb7, 0x42, 0x2f, 0x85, 0xc4, 0x72, 0x78, 0x5c, 0x27, 0x5f, 0xde,
	0x08, 0x75, 0x70, 0xfc, 0xfd, 0xbf, 0xe5, 0x13, 0x1b, 0x82, 0x57, 0x80, 0x81, 0x8b, 0x9e, 0x90,
	0x51, 0x66, 0x91, 0xf9, 0x5b, 0xac, 0x6e, 0xad, 0xeb, 0xca, 0x89, 0x52, 0x89, 0xcc, 0x58, 0x4b,
	0x30, 0x0c, 0x2a, 0xd9, 0x25, 0xfb, 0x65, 0xf9, 0xa3, 0xdc, 0x8a, 0xd4, 0x03, 0x37, 0x68, 0xd5,
	0x6c, 0x1c, 0xa2, 0xbe, 0xbd, 0x0b, 0xbe, 0x02, 0x52, 0xf3, 0xa9, 0x4f, 0x92, 0xb9, 0x58, 0xd7,
	0x8d, 0x21, 0x07, 0x67, 0x8f, 0xa1, 0x63, 0xb5, 0x18, 0xd7, 0xbf, 0x76, 0x39, 0x3e, 0x0e, 0xae,
	0xe5, 0x26, 0x54, 0x45, 0x8d, 0x5e, 0x97, 0x77, 0x32, 0x16, 0xae, 0x6c, 0xb5, 0xa3, 0xbe, 0xd6,
	0x0c, 0xbe, 0x47, 0x32, 0x4f, 0x1d, 0xaf, 0x56, 0xe9, 0x78, 0xfd, 0xf4, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x42, 0xc7, 0x24, 0xe4, 0xcb, 0x00, 0x00, 0x00,
}
