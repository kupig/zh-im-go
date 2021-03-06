// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg_test.proto

package pb

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

type MsgTestRep struct {
	MsgType              int32    `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgTestRep) Reset()         { *m = MsgTestRep{} }
func (m *MsgTestRep) String() string { return proto.CompactTextString(m) }
func (*MsgTestRep) ProtoMessage()    {}
func (*MsgTestRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_393017e1f95de7c6, []int{0}
}

func (m *MsgTestRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgTestRep.Unmarshal(m, b)
}
func (m *MsgTestRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgTestRep.Marshal(b, m, deterministic)
}
func (m *MsgTestRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTestRep.Merge(m, src)
}
func (m *MsgTestRep) XXX_Size() int {
	return xxx_messageInfo_MsgTestRep.Size(m)
}
func (m *MsgTestRep) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTestRep.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTestRep proto.InternalMessageInfo

func (m *MsgTestRep) GetMsgType() int32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *MsgTestRep) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MsgTestRep) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type MsgTypeResq struct {
	MsgType              int32    `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgTypeResq) Reset()         { *m = MsgTypeResq{} }
func (m *MsgTypeResq) String() string { return proto.CompactTextString(m) }
func (*MsgTypeResq) ProtoMessage()    {}
func (*MsgTypeResq) Descriptor() ([]byte, []int) {
	return fileDescriptor_393017e1f95de7c6, []int{1}
}

func (m *MsgTypeResq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgTypeResq.Unmarshal(m, b)
}
func (m *MsgTypeResq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgTypeResq.Marshal(b, m, deterministic)
}
func (m *MsgTypeResq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTypeResq.Merge(m, src)
}
func (m *MsgTypeResq) XXX_Size() int {
	return xxx_messageInfo_MsgTypeResq.Size(m)
}
func (m *MsgTypeResq) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTypeResq.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTypeResq proto.InternalMessageInfo

func (m *MsgTypeResq) GetMsgType() int32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *MsgTypeResq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgTestRep)(nil), "pb.MsgTestRep")
	proto.RegisterType((*MsgTypeResq)(nil), "pb.MsgTypeResq")
}

func init() { proto.RegisterFile("msg_test.proto", fileDescriptor_393017e1f95de7c6) }

var fileDescriptor_393017e1f95de7c6 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2d, 0x4e, 0x8f,
	0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x0a,
	0xe5, 0xe2, 0xf2, 0x2d, 0x4e, 0x0f, 0x49, 0x2d, 0x2e, 0x09, 0x4a, 0x2d, 0x10, 0x92, 0xe4, 0xe2,
	0x00, 0xab, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xcf, 0x2d, 0x4e,
	0x0f, 0xa9, 0x2c, 0x48, 0x15, 0x92, 0xe2, 0xe2, 0x28, 0x2d, 0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x04, 0xb8, 0x98, 0x13, 0xd3, 0x53,
	0x25, 0x98, 0xc1, 0x3a, 0x40, 0x4c, 0x25, 0x27, 0x2e, 0x6e, 0x5f, 0x88, 0xc6, 0xa0, 0xd4, 0xe2,
	0x42, 0x7c, 0xe6, 0x4a, 0x70, 0xb1, 0x27, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x40, 0x8d, 0x85,
	0x71, 0x93, 0xd8, 0xc0, 0xae, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xa8, 0x67, 0xa5,
	0xb7, 0x00, 0x00, 0x00,
}
