// Code generated by protoc-gen-go. DO NOT EDIT.
// source: porpackage.proto

package por

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

type PorPackage struct {
	Height               uint32    `protobuf:"varint,1,opt,name=Height" json:"Height,omitempty"`
	Owner                []byte    `protobuf:"bytes,2,opt,name=Owner,proto3" json:"Owner,omitempty"`
	BlockHash            []byte    `protobuf:"bytes,3,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	TxHash               []byte    `protobuf:"bytes,4,opt,name=TxHash,proto3" json:"TxHash,omitempty"`
	SigChainHash         []byte    `protobuf:"bytes,5,opt,name=SigChainHash,proto3" json:"SigChainHash,omitempty"`
	SigChain             *SigChain `protobuf:"bytes,6,opt,name=SigChain" json:"SigChain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PorPackage) Reset()         { *m = PorPackage{} }
func (m *PorPackage) String() string { return proto.CompactTextString(m) }
func (*PorPackage) ProtoMessage()    {}
func (*PorPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_porpackage_e3da189f1b5e31b2, []int{0}
}
func (m *PorPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PorPackage.Unmarshal(m, b)
}
func (m *PorPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PorPackage.Marshal(b, m, deterministic)
}
func (dst *PorPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PorPackage.Merge(dst, src)
}
func (m *PorPackage) XXX_Size() int {
	return xxx_messageInfo_PorPackage.Size(m)
}
func (m *PorPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_PorPackage.DiscardUnknown(m)
}

var xxx_messageInfo_PorPackage proto.InternalMessageInfo

func (m *PorPackage) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *PorPackage) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PorPackage) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *PorPackage) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *PorPackage) GetSigChainHash() []byte {
	if m != nil {
		return m.SigChainHash
	}
	return nil
}

func (m *PorPackage) GetSigChain() *SigChain {
	if m != nil {
		return m.SigChain
	}
	return nil
}

func init() {
	proto.RegisterType((*PorPackage)(nil), "por.PorPackage")
}

func init() { proto.RegisterFile("porpackage.proto", fileDescriptor_porpackage_e3da189f1b5e31b2) }

var fileDescriptor_porpackage_e3da189f1b5e31b2 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xc8, 0x2f, 0x2a,
	0x48, 0x4c, 0xce, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0xc8,
	0x2f, 0x92, 0xe2, 0x2b, 0xce, 0x4c, 0x4f, 0xce, 0x48, 0xcc, 0xcc, 0x83, 0x08, 0x2a, 0xed, 0x67,
	0xe4, 0xe2, 0x0a, 0xc8, 0x2f, 0x0a, 0x80, 0xa8, 0x14, 0x12, 0xe3, 0x62, 0xf3, 0x48, 0xcd, 0x4c,
	0xcf, 0x28, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x82, 0xf2, 0x84, 0x44, 0xb8, 0x58, 0xfd,
	0xcb, 0xf3, 0x52, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0x20, 0x1c, 0x21, 0x19, 0x2e,
	0x4e, 0xa7, 0x9c, 0xfc, 0xe4, 0x6c, 0x8f, 0xc4, 0xe2, 0x0c, 0x09, 0x66, 0xb0, 0x0c, 0x42, 0x00,
	0x64, 0x56, 0x48, 0x05, 0x58, 0x8a, 0x05, 0x2c, 0x05, 0xe5, 0x09, 0x29, 0x71, 0xf1, 0x04, 0x67,
	0xa6, 0x3b, 0x83, 0x1c, 0x01, 0x96, 0x65, 0x05, 0xcb, 0xa2, 0x88, 0x09, 0x69, 0x72, 0x71, 0xc0,
	0xf8, 0x12, 0x6c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xbc, 0x7a, 0x05, 0xf9, 0x45, 0x7a, 0x30, 0xc1,
	0x20, 0xb8, 0x74, 0x12, 0x1b, 0xd8, 0x23, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x5e,
	0x24, 0x6b, 0xf1, 0x00, 0x00, 0x00,
}
