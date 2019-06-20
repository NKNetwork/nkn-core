// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/sigchain.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SigAlgo int32

const (
	SIGNATURE SigAlgo = 0
	VRF       SigAlgo = 1
)

var SigAlgo_name = map[int32]string{
	0: "SIGNATURE",
	1: "VRF",
}
var SigAlgo_value = map[string]int32{
	"SIGNATURE": 0,
	"VRF":       1,
}

func (SigAlgo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sigchain_a5e1681c49494756, []int{0}
}

type SigChainElem struct {
	Id         []byte  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NextPubkey []byte  `protobuf:"bytes,2,opt,name=next_pubkey,json=nextPubkey,proto3" json:"next_pubkey,omitempty"`
	Mining     bool    `protobuf:"varint,3,opt,name=mining,proto3" json:"mining,omitempty"`
	Signature  []byte  `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	SigAlgo    SigAlgo `protobuf:"varint,5,opt,name=sig_algo,json=sigAlgo,proto3,enum=pb.SigAlgo" json:"sig_algo,omitempty"`
	Vrf        []byte  `protobuf:"bytes,6,opt,name=vrf,proto3" json:"vrf,omitempty"`
	Proof      []byte  `protobuf:"bytes,7,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *SigChainElem) Reset()      { *m = SigChainElem{} }
func (*SigChainElem) ProtoMessage() {}
func (*SigChainElem) Descriptor() ([]byte, []int) {
	return fileDescriptor_sigchain_a5e1681c49494756, []int{0}
}
func (m *SigChainElem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigChainElem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigChainElem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SigChainElem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigChainElem.Merge(dst, src)
}
func (m *SigChainElem) XXX_Size() int {
	return m.Size()
}
func (m *SigChainElem) XXX_DiscardUnknown() {
	xxx_messageInfo_SigChainElem.DiscardUnknown(m)
}

var xxx_messageInfo_SigChainElem proto.InternalMessageInfo

func (m *SigChainElem) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SigChainElem) GetNextPubkey() []byte {
	if m != nil {
		return m.NextPubkey
	}
	return nil
}

func (m *SigChainElem) GetMining() bool {
	if m != nil {
		return m.Mining
	}
	return false
}

func (m *SigChainElem) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SigChainElem) GetSigAlgo() SigAlgo {
	if m != nil {
		return m.SigAlgo
	}
	return SIGNATURE
}

func (m *SigChainElem) GetVrf() []byte {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *SigChainElem) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type SigChain struct {
	Nonce      uint32          `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	DataSize   uint32          `protobuf:"varint,2,opt,name=data_size,json=dataSize,proto3" json:"data_size,omitempty"`
	BlockHash  []byte          `protobuf:"bytes,3,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	SrcId      []byte          `protobuf:"bytes,4,opt,name=src_id,json=srcId,proto3" json:"src_id,omitempty"`
	SrcPubkey  []byte          `protobuf:"bytes,5,opt,name=src_pubkey,json=srcPubkey,proto3" json:"src_pubkey,omitempty"`
	DestId     []byte          `protobuf:"bytes,6,opt,name=dest_id,json=destId,proto3" json:"dest_id,omitempty"`
	DestPubkey []byte          `protobuf:"bytes,7,opt,name=dest_pubkey,json=destPubkey,proto3" json:"dest_pubkey,omitempty"`
	Elems      []*SigChainElem `protobuf:"bytes,8,rep,name=elems" json:"elems,omitempty"`
}

func (m *SigChain) Reset()      { *m = SigChain{} }
func (*SigChain) ProtoMessage() {}
func (*SigChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_sigchain_a5e1681c49494756, []int{1}
}
func (m *SigChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SigChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigChain.Merge(dst, src)
}
func (m *SigChain) XXX_Size() int {
	return m.Size()
}
func (m *SigChain) XXX_DiscardUnknown() {
	xxx_messageInfo_SigChain.DiscardUnknown(m)
}

var xxx_messageInfo_SigChain proto.InternalMessageInfo

func (m *SigChain) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *SigChain) GetDataSize() uint32 {
	if m != nil {
		return m.DataSize
	}
	return 0
}

func (m *SigChain) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *SigChain) GetSrcId() []byte {
	if m != nil {
		return m.SrcId
	}
	return nil
}

func (m *SigChain) GetSrcPubkey() []byte {
	if m != nil {
		return m.SrcPubkey
	}
	return nil
}

func (m *SigChain) GetDestId() []byte {
	if m != nil {
		return m.DestId
	}
	return nil
}

func (m *SigChain) GetDestPubkey() []byte {
	if m != nil {
		return m.DestPubkey
	}
	return nil
}

func (m *SigChain) GetElems() []*SigChainElem {
	if m != nil {
		return m.Elems
	}
	return nil
}

func init() {
	proto.RegisterType((*SigChainElem)(nil), "pb.SigChainElem")
	proto.RegisterType((*SigChain)(nil), "pb.SigChain")
	proto.RegisterEnum("pb.SigAlgo", SigAlgo_name, SigAlgo_value)
}
func (x SigAlgo) String() string {
	s, ok := SigAlgo_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *SigChainElem) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SigChainElem)
	if !ok {
		that2, ok := that.(SigChainElem)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Id, that1.Id) {
		return false
	}
	if !bytes.Equal(this.NextPubkey, that1.NextPubkey) {
		return false
	}
	if this.Mining != that1.Mining {
		return false
	}
	if !bytes.Equal(this.Signature, that1.Signature) {
		return false
	}
	if this.SigAlgo != that1.SigAlgo {
		return false
	}
	if !bytes.Equal(this.Vrf, that1.Vrf) {
		return false
	}
	if !bytes.Equal(this.Proof, that1.Proof) {
		return false
	}
	return true
}
func (this *SigChain) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SigChain)
	if !ok {
		that2, ok := that.(SigChain)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	if this.DataSize != that1.DataSize {
		return false
	}
	if !bytes.Equal(this.BlockHash, that1.BlockHash) {
		return false
	}
	if !bytes.Equal(this.SrcId, that1.SrcId) {
		return false
	}
	if !bytes.Equal(this.SrcPubkey, that1.SrcPubkey) {
		return false
	}
	if !bytes.Equal(this.DestId, that1.DestId) {
		return false
	}
	if !bytes.Equal(this.DestPubkey, that1.DestPubkey) {
		return false
	}
	if len(this.Elems) != len(that1.Elems) {
		return false
	}
	for i := range this.Elems {
		if !this.Elems[i].Equal(that1.Elems[i]) {
			return false
		}
	}
	return true
}
func (this *SigChainElem) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&pb.SigChainElem{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "NextPubkey: "+fmt.Sprintf("%#v", this.NextPubkey)+",\n")
	s = append(s, "Mining: "+fmt.Sprintf("%#v", this.Mining)+",\n")
	s = append(s, "Signature: "+fmt.Sprintf("%#v", this.Signature)+",\n")
	s = append(s, "SigAlgo: "+fmt.Sprintf("%#v", this.SigAlgo)+",\n")
	s = append(s, "Vrf: "+fmt.Sprintf("%#v", this.Vrf)+",\n")
	s = append(s, "Proof: "+fmt.Sprintf("%#v", this.Proof)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SigChain) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&pb.SigChain{")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "DataSize: "+fmt.Sprintf("%#v", this.DataSize)+",\n")
	s = append(s, "BlockHash: "+fmt.Sprintf("%#v", this.BlockHash)+",\n")
	s = append(s, "SrcId: "+fmt.Sprintf("%#v", this.SrcId)+",\n")
	s = append(s, "SrcPubkey: "+fmt.Sprintf("%#v", this.SrcPubkey)+",\n")
	s = append(s, "DestId: "+fmt.Sprintf("%#v", this.DestId)+",\n")
	s = append(s, "DestPubkey: "+fmt.Sprintf("%#v", this.DestPubkey)+",\n")
	if this.Elems != nil {
		s = append(s, "Elems: "+fmt.Sprintf("%#v", this.Elems)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSigchain(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SigChainElem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigChainElem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.NextPubkey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.NextPubkey)))
		i += copy(dAtA[i:], m.NextPubkey)
	}
	if m.Mining {
		dAtA[i] = 0x18
		i++
		if m.Mining {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Signature) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	if m.SigAlgo != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(m.SigAlgo))
	}
	if len(m.Vrf) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.Vrf)))
		i += copy(dAtA[i:], m.Vrf)
	}
	if len(m.Proof) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.Proof)))
		i += copy(dAtA[i:], m.Proof)
	}
	return i, nil
}

func (m *SigChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigChain) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Nonce != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(m.Nonce))
	}
	if m.DataSize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(m.DataSize))
	}
	if len(m.BlockHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.BlockHash)))
		i += copy(dAtA[i:], m.BlockHash)
	}
	if len(m.SrcId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.SrcId)))
		i += copy(dAtA[i:], m.SrcId)
	}
	if len(m.SrcPubkey) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.SrcPubkey)))
		i += copy(dAtA[i:], m.SrcPubkey)
	}
	if len(m.DestId) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.DestId)))
		i += copy(dAtA[i:], m.DestId)
	}
	if len(m.DestPubkey) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.DestPubkey)))
		i += copy(dAtA[i:], m.DestPubkey)
	}
	if len(m.Elems) > 0 {
		for _, msg := range m.Elems {
			dAtA[i] = 0x42
			i++
			i = encodeVarintSigchain(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintSigchain(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedSigChainElem(r randySigchain, easy bool) *SigChainElem {
	this := &SigChainElem{}
	v1 := r.Intn(100)
	this.Id = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Id[i] = byte(r.Intn(256))
	}
	v2 := r.Intn(100)
	this.NextPubkey = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.NextPubkey[i] = byte(r.Intn(256))
	}
	this.Mining = bool(bool(r.Intn(2) == 0))
	v3 := r.Intn(100)
	this.Signature = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.Signature[i] = byte(r.Intn(256))
	}
	this.SigAlgo = SigAlgo([]int32{0, 1}[r.Intn(2)])
	v4 := r.Intn(100)
	this.Vrf = make([]byte, v4)
	for i := 0; i < v4; i++ {
		this.Vrf[i] = byte(r.Intn(256))
	}
	v5 := r.Intn(100)
	this.Proof = make([]byte, v5)
	for i := 0; i < v5; i++ {
		this.Proof[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSigChain(r randySigchain, easy bool) *SigChain {
	this := &SigChain{}
	this.Nonce = uint32(r.Uint32())
	this.DataSize = uint32(r.Uint32())
	v6 := r.Intn(100)
	this.BlockHash = make([]byte, v6)
	for i := 0; i < v6; i++ {
		this.BlockHash[i] = byte(r.Intn(256))
	}
	v7 := r.Intn(100)
	this.SrcId = make([]byte, v7)
	for i := 0; i < v7; i++ {
		this.SrcId[i] = byte(r.Intn(256))
	}
	v8 := r.Intn(100)
	this.SrcPubkey = make([]byte, v8)
	for i := 0; i < v8; i++ {
		this.SrcPubkey[i] = byte(r.Intn(256))
	}
	v9 := r.Intn(100)
	this.DestId = make([]byte, v9)
	for i := 0; i < v9; i++ {
		this.DestId[i] = byte(r.Intn(256))
	}
	v10 := r.Intn(100)
	this.DestPubkey = make([]byte, v10)
	for i := 0; i < v10; i++ {
		this.DestPubkey[i] = byte(r.Intn(256))
	}
	if r.Intn(10) != 0 {
		v11 := r.Intn(5)
		this.Elems = make([]*SigChainElem, v11)
		for i := 0; i < v11; i++ {
			this.Elems[i] = NewPopulatedSigChainElem(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randySigchain interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSigchain(r randySigchain) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSigchain(r randySigchain) string {
	v12 := r.Intn(100)
	tmps := make([]rune, v12)
	for i := 0; i < v12; i++ {
		tmps[i] = randUTF8RuneSigchain(r)
	}
	return string(tmps)
}
func randUnrecognizedSigchain(r randySigchain, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSigchain(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSigchain(dAtA []byte, r randySigchain, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSigchain(dAtA, uint64(key))
		v13 := r.Int63()
		if r.Intn(2) == 0 {
			v13 *= -1
		}
		dAtA = encodeVarintPopulateSigchain(dAtA, uint64(v13))
	case 1:
		dAtA = encodeVarintPopulateSigchain(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSigchain(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSigchain(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSigchain(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSigchain(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *SigChainElem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	l = len(m.NextPubkey)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	if m.Mining {
		n += 2
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	if m.SigAlgo != 0 {
		n += 1 + sovSigchain(uint64(m.SigAlgo))
	}
	l = len(m.Vrf)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	return n
}

func (m *SigChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovSigchain(uint64(m.Nonce))
	}
	if m.DataSize != 0 {
		n += 1 + sovSigchain(uint64(m.DataSize))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	l = len(m.SrcId)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	l = len(m.SrcPubkey)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	l = len(m.DestId)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	l = len(m.DestPubkey)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	if len(m.Elems) > 0 {
		for _, e := range m.Elems {
			l = e.Size()
			n += 1 + l + sovSigchain(uint64(l))
		}
	}
	return n
}

func sovSigchain(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSigchain(x uint64) (n int) {
	return sovSigchain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SigChainElem) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SigChainElem{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`NextPubkey:` + fmt.Sprintf("%v", this.NextPubkey) + `,`,
		`Mining:` + fmt.Sprintf("%v", this.Mining) + `,`,
		`Signature:` + fmt.Sprintf("%v", this.Signature) + `,`,
		`SigAlgo:` + fmt.Sprintf("%v", this.SigAlgo) + `,`,
		`Vrf:` + fmt.Sprintf("%v", this.Vrf) + `,`,
		`Proof:` + fmt.Sprintf("%v", this.Proof) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SigChain) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SigChain{`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`DataSize:` + fmt.Sprintf("%v", this.DataSize) + `,`,
		`BlockHash:` + fmt.Sprintf("%v", this.BlockHash) + `,`,
		`SrcId:` + fmt.Sprintf("%v", this.SrcId) + `,`,
		`SrcPubkey:` + fmt.Sprintf("%v", this.SrcPubkey) + `,`,
		`DestId:` + fmt.Sprintf("%v", this.DestId) + `,`,
		`DestPubkey:` + fmt.Sprintf("%v", this.DestPubkey) + `,`,
		`Elems:` + strings.Replace(fmt.Sprintf("%v", this.Elems), "SigChainElem", "SigChainElem", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSigchain(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SigChainElem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSigchain
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SigChainElem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigChainElem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextPubkey = append(m.NextPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.NextPubkey == nil {
				m.NextPubkey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mining", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Mining = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigAlgo", wireType)
			}
			m.SigAlgo = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigAlgo |= (SigAlgo(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vrf", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vrf = append(m.Vrf[:0], dAtA[iNdEx:postIndex]...)
			if m.Vrf == nil {
				m.Vrf = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = append(m.Proof[:0], dAtA[iNdEx:postIndex]...)
			if m.Proof == nil {
				m.Proof = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSigchain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSigchain
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SigChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSigchain
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SigChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataSize", wireType)
			}
			m.DataSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcId = append(m.SrcId[:0], dAtA[iNdEx:postIndex]...)
			if m.SrcId == nil {
				m.SrcId = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcPubkey = append(m.SrcPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.SrcPubkey == nil {
				m.SrcPubkey = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestId = append(m.DestId[:0], dAtA[iNdEx:postIndex]...)
			if m.DestId == nil {
				m.DestId = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestPubkey = append(m.DestPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.DestPubkey == nil {
				m.DestPubkey = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elems", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elems = append(m.Elems, &SigChainElem{})
			if err := m.Elems[len(m.Elems)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSigchain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSigchain
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSigchain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSigchain
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSigchain
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSigchain
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSigchain(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSigchain = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSigchain   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/sigchain.proto", fileDescriptor_sigchain_a5e1681c49494756) }

var fileDescriptor_sigchain_a5e1681c49494756 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x92, 0xbb, 0x8e, 0xd3, 0x40,
	0x18, 0x85, 0x3d, 0x09, 0xb1, 0x9d, 0x3f, 0xc9, 0x2a, 0x8c, 0xb8, 0x58, 0x5c, 0x86, 0xb0, 0xc5,
	0x2a, 0x42, 0x22, 0x91, 0x96, 0x96, 0x66, 0x41, 0x0b, 0xa4, 0x41, 0x68, 0x02, 0xb4, 0x96, 0x2f,
	0x93, 0xf1, 0x68, 0x1d, 0x8f, 0xe5, 0x71, 0x10, 0x6c, 0xc5, 0x23, 0xf0, 0x18, 0x3c, 0x02, 0x2f,
	0x80, 0x44, 0x99, 0x72, 0x4b, 0xe2, 0x34, 0x94, 0x5b, 0x50, 0x50, 0xa2, 0x19, 0x0f, 0x6c, 0xf7,
	0x9f, 0x73, 0x7c, 0x2c, 0x7d, 0xc7, 0x86, 0xeb, 0x65, 0x3c, 0x57, 0x82, 0x27, 0x59, 0x24, 0x8a,
	0x59, 0x59, 0xc9, 0x5a, 0xe2, 0x4e, 0x19, 0xdf, 0x79, 0xcc, 0x45, 0x9d, 0x6d, 0xe2, 0x59, 0x22,
	0xd7, 0x73, 0x2e, 0xb9, 0x9c, 0x9b, 0x28, 0xde, 0xac, 0x8c, 0x32, 0xc2, 0x5c, 0x6d, 0xe5, 0xf0,
	0x3b, 0x82, 0xe1, 0x52, 0xf0, 0xe7, 0xfa, 0x2d, 0xa7, 0x39, 0x5b, 0xe3, 0x03, 0xe8, 0x88, 0x34,
	0x40, 0x13, 0x34, 0x1d, 0xd2, 0x8e, 0x48, 0xf1, 0x03, 0x18, 0x14, 0xec, 0x63, 0x1d, 0x96, 0x9b,
	0xf8, 0x8c, 0x7d, 0x0a, 0x3a, 0x26, 0x00, 0x6d, 0xbd, 0x31, 0x0e, 0xbe, 0x05, 0xee, 0x5a, 0x14,
	0xa2, 0xe0, 0x41, 0x77, 0x82, 0xa6, 0x3e, 0xb5, 0x0a, 0xdf, 0x83, 0xbe, 0x12, 0xbc, 0x88, 0xea,
	0x4d, 0xc5, 0x82, 0x6b, 0xa6, 0x76, 0x65, 0xe0, 0x23, 0xf0, 0x95, 0xe0, 0x61, 0x94, 0x73, 0x19,
	0xf4, 0x26, 0x68, 0x7a, 0x70, 0x3c, 0x98, 0x95, 0xf1, 0x6c, 0x29, 0xf8, 0x49, 0xce, 0x25, 0xf5,
	0x54, 0x7b, 0xe0, 0x31, 0x74, 0x3f, 0x54, 0xab, 0xc0, 0x35, 0x7d, 0x7d, 0xe2, 0x1b, 0xd0, 0x2b,
	0x2b, 0x29, 0x57, 0x81, 0x67, 0xbc, 0x56, 0x1c, 0xfe, 0x46, 0xe0, 0xff, 0xe3, 0xd0, 0x8f, 0x14,
	0xb2, 0x48, 0x98, 0xc1, 0x18, 0xd1, 0x56, 0xe0, 0xbb, 0xd0, 0x4f, 0xa3, 0x3a, 0x0a, 0x95, 0x38,
	0x67, 0x86, 0x63, 0x44, 0x7d, 0x6d, 0x2c, 0xc5, 0x39, 0xc3, 0xf7, 0x01, 0xe2, 0x5c, 0x26, 0x67,
	0x61, 0x16, 0xa9, 0xcc, 0x90, 0x0c, 0x69, 0xdf, 0x38, 0xaf, 0x22, 0x95, 0xe1, 0x9b, 0xe0, 0xaa,
	0x2a, 0x09, 0x45, 0x6a, 0x49, 0x7a, 0xaa, 0x4a, 0x16, 0xa9, 0x6e, 0x69, 0xdb, 0x6e, 0xd3, 0xb3,
	0x90, 0x55, 0x62, 0xa7, 0xb9, 0x0d, 0x5e, 0xca, 0x54, 0xad, 0x6b, 0x2d, 0x80, 0xab, 0xe5, 0xc2,
	0x8c, 0x6a, 0x02, 0x5b, 0x6c, 0x49, 0x40, 0x5b, 0xb6, 0x79, 0x04, 0x3d, 0x96, 0xb3, 0xb5, 0x0a,
	0xfc, 0x49, 0x77, 0x3a, 0x38, 0x1e, 0xdb, 0x6d, 0xfe, 0x7f, 0x26, 0xda, 0xc6, 0x8f, 0x1e, 0x82,
	0x67, 0x27, 0xc3, 0x23, 0xe8, 0x2f, 0x17, 0x2f, 0x5f, 0x9f, 0xbc, 0x7d, 0x47, 0x4f, 0xc7, 0x0e,
	0xf6, 0xa0, 0xfb, 0x9e, 0xbe, 0x18, 0xa3, 0x67, 0x4f, 0xb7, 0x3b, 0xe2, 0x5c, 0xec, 0x88, 0x73,
	0xb9, 0x23, 0xe8, 0xcf, 0x8e, 0xa0, 0xcf, 0x0d, 0x41, 0x5f, 0x1b, 0x82, 0xbe, 0x35, 0x04, 0xfd,
	0x68, 0x08, 0xda, 0x36, 0x04, 0xfd, 0x6c, 0x08, 0xfa, 0xd5, 0x10, 0xe7, 0xb2, 0x21, 0xe8, 0xcb,
	0x9e, 0x38, 0xdb, 0x3d, 0x71, 0x2e, 0xf6, 0xc4, 0x89, 0x5d, 0xf3, 0x9b, 0x3c, 0xf9, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xe3, 0xb3, 0xac, 0xfe, 0x6e, 0x02, 0x00, 0x00,
}
