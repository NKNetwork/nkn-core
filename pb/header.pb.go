// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/header.proto

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

type WinnerType int32

const (
	GENESIS_SIGNER WinnerType = 0
	TXN_SIGNER     WinnerType = 1
	BLOCK_SIGNER   WinnerType = 2
)

var WinnerType_name = map[int32]string{
	0: "GENESIS_SIGNER",
	1: "TXN_SIGNER",
	2: "BLOCK_SIGNER",
}
var WinnerType_value = map[string]int32{
	"GENESIS_SIGNER": 0,
	"TXN_SIGNER":     1,
	"BLOCK_SIGNER":   2,
}

func (WinnerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_header_c1b05b6bee120f74, []int{0}
}

type UnsignedHeader struct {
	Version          uint32     `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	PrevBlockHash    []byte     `protobuf:"bytes,2,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash,omitempty"`
	TransactionsRoot []byte     `protobuf:"bytes,3,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"`
	StateRoot        []byte     `protobuf:"bytes,4,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	Timestamp        int64      `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Height           uint32     `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	ConsensusData    uint64     `protobuf:"varint,7,opt,name=consensus_data,json=consensusData,proto3" json:"consensus_data,omitempty"`
	WinnerHash       []byte     `protobuf:"bytes,8,opt,name=winner_hash,json=winnerHash,proto3" json:"winner_hash,omitempty"`
	WinnerType       WinnerType `protobuf:"varint,9,opt,name=winner_type,json=winnerType,proto3,enum=pb.WinnerType" json:"winner_type,omitempty"`
	Signer           []byte     `protobuf:"bytes,10,opt,name=signer,proto3" json:"signer,omitempty"`
	ChordId          []byte     `protobuf:"bytes,11,opt,name=chord_id,json=chordId,proto3" json:"chord_id,omitempty"`
}

func (m *UnsignedHeader) Reset()      { *m = UnsignedHeader{} }
func (*UnsignedHeader) ProtoMessage() {}
func (*UnsignedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_header_c1b05b6bee120f74, []int{0}
}
func (m *UnsignedHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnsignedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnsignedHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UnsignedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsignedHeader.Merge(dst, src)
}
func (m *UnsignedHeader) XXX_Size() int {
	return m.Size()
}
func (m *UnsignedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsignedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_UnsignedHeader proto.InternalMessageInfo

func (m *UnsignedHeader) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *UnsignedHeader) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *UnsignedHeader) GetTransactionsRoot() []byte {
	if m != nil {
		return m.TransactionsRoot
	}
	return nil
}

func (m *UnsignedHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *UnsignedHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *UnsignedHeader) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *UnsignedHeader) GetConsensusData() uint64 {
	if m != nil {
		return m.ConsensusData
	}
	return 0
}

func (m *UnsignedHeader) GetWinnerHash() []byte {
	if m != nil {
		return m.WinnerHash
	}
	return nil
}

func (m *UnsignedHeader) GetWinnerType() WinnerType {
	if m != nil {
		return m.WinnerType
	}
	return GENESIS_SIGNER
}

func (m *UnsignedHeader) GetSigner() []byte {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *UnsignedHeader) GetChordId() []byte {
	if m != nil {
		return m.ChordId
	}
	return nil
}

type BlockHeader struct {
	UnsignedHeader *UnsignedHeader `protobuf:"bytes,1,opt,name=unsigned_header,json=unsignedHeader" json:"unsigned_header,omitempty"`
	Signature      []byte          `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *BlockHeader) Reset()      { *m = BlockHeader{} }
func (*BlockHeader) ProtoMessage() {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_header_c1b05b6bee120f74, []int{1}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(dst, src)
}
func (m *BlockHeader) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetUnsignedHeader() *UnsignedHeader {
	if m != nil {
		return m.UnsignedHeader
	}
	return nil
}

func (m *BlockHeader) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*UnsignedHeader)(nil), "pb.UnsignedHeader")
	proto.RegisterType((*BlockHeader)(nil), "pb.BlockHeader")
	proto.RegisterEnum("pb.WinnerType", WinnerType_name, WinnerType_value)
}
func (x WinnerType) String() string {
	s, ok := WinnerType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *UnsignedHeader) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UnsignedHeader)
	if !ok {
		that2, ok := that.(UnsignedHeader)
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
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.PrevBlockHash, that1.PrevBlockHash) {
		return false
	}
	if !bytes.Equal(this.TransactionsRoot, that1.TransactionsRoot) {
		return false
	}
	if !bytes.Equal(this.StateRoot, that1.StateRoot) {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if this.ConsensusData != that1.ConsensusData {
		return false
	}
	if !bytes.Equal(this.WinnerHash, that1.WinnerHash) {
		return false
	}
	if this.WinnerType != that1.WinnerType {
		return false
	}
	if !bytes.Equal(this.Signer, that1.Signer) {
		return false
	}
	if !bytes.Equal(this.ChordId, that1.ChordId) {
		return false
	}
	return true
}
func (this *BlockHeader) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BlockHeader)
	if !ok {
		that2, ok := that.(BlockHeader)
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
	if !this.UnsignedHeader.Equal(that1.UnsignedHeader) {
		return false
	}
	if !bytes.Equal(this.Signature, that1.Signature) {
		return false
	}
	return true
}
func (this *UnsignedHeader) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 15)
	s = append(s, "&pb.UnsignedHeader{")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "PrevBlockHash: "+fmt.Sprintf("%#v", this.PrevBlockHash)+",\n")
	s = append(s, "TransactionsRoot: "+fmt.Sprintf("%#v", this.TransactionsRoot)+",\n")
	s = append(s, "StateRoot: "+fmt.Sprintf("%#v", this.StateRoot)+",\n")
	s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	s = append(s, "Height: "+fmt.Sprintf("%#v", this.Height)+",\n")
	s = append(s, "ConsensusData: "+fmt.Sprintf("%#v", this.ConsensusData)+",\n")
	s = append(s, "WinnerHash: "+fmt.Sprintf("%#v", this.WinnerHash)+",\n")
	s = append(s, "WinnerType: "+fmt.Sprintf("%#v", this.WinnerType)+",\n")
	s = append(s, "Signer: "+fmt.Sprintf("%#v", this.Signer)+",\n")
	s = append(s, "ChordId: "+fmt.Sprintf("%#v", this.ChordId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *BlockHeader) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.BlockHeader{")
	if this.UnsignedHeader != nil {
		s = append(s, "UnsignedHeader: "+fmt.Sprintf("%#v", this.UnsignedHeader)+",\n")
	}
	s = append(s, "Signature: "+fmt.Sprintf("%#v", this.Signature)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringHeader(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *UnsignedHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnsignedHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.Version))
	}
	if len(m.PrevBlockHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.PrevBlockHash)))
		i += copy(dAtA[i:], m.PrevBlockHash)
	}
	if len(m.TransactionsRoot) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.TransactionsRoot)))
		i += copy(dAtA[i:], m.TransactionsRoot)
	}
	if len(m.StateRoot) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.StateRoot)))
		i += copy(dAtA[i:], m.StateRoot)
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.Timestamp))
	}
	if m.Height != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.Height))
	}
	if m.ConsensusData != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.ConsensusData))
	}
	if len(m.WinnerHash) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.WinnerHash)))
		i += copy(dAtA[i:], m.WinnerHash)
	}
	if m.WinnerType != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.WinnerType))
	}
	if len(m.Signer) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.Signer)))
		i += copy(dAtA[i:], m.Signer)
	}
	if len(m.ChordId) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.ChordId)))
		i += copy(dAtA[i:], m.ChordId)
	}
	return i, nil
}

func (m *BlockHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UnsignedHeader != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.UnsignedHeader.Size()))
		n1, err := m.UnsignedHeader.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Signature) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	return i, nil
}

func encodeVarintHeader(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedUnsignedHeader(r randyHeader, easy bool) *UnsignedHeader {
	this := &UnsignedHeader{}
	this.Version = uint32(r.Uint32())
	v1 := r.Intn(100)
	this.PrevBlockHash = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.PrevBlockHash[i] = byte(r.Intn(256))
	}
	v2 := r.Intn(100)
	this.TransactionsRoot = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.TransactionsRoot[i] = byte(r.Intn(256))
	}
	v3 := r.Intn(100)
	this.StateRoot = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.StateRoot[i] = byte(r.Intn(256))
	}
	this.Timestamp = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Timestamp *= -1
	}
	this.Height = uint32(r.Uint32())
	this.ConsensusData = uint64(uint64(r.Uint32()))
	v4 := r.Intn(100)
	this.WinnerHash = make([]byte, v4)
	for i := 0; i < v4; i++ {
		this.WinnerHash[i] = byte(r.Intn(256))
	}
	this.WinnerType = WinnerType([]int32{0, 1, 2}[r.Intn(3)])
	v5 := r.Intn(100)
	this.Signer = make([]byte, v5)
	for i := 0; i < v5; i++ {
		this.Signer[i] = byte(r.Intn(256))
	}
	v6 := r.Intn(100)
	this.ChordId = make([]byte, v6)
	for i := 0; i < v6; i++ {
		this.ChordId[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedBlockHeader(r randyHeader, easy bool) *BlockHeader {
	this := &BlockHeader{}
	if r.Intn(10) != 0 {
		this.UnsignedHeader = NewPopulatedUnsignedHeader(r, easy)
	}
	v7 := r.Intn(100)
	this.Signature = make([]byte, v7)
	for i := 0; i < v7; i++ {
		this.Signature[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyHeader interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneHeader(r randyHeader) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringHeader(r randyHeader) string {
	v8 := r.Intn(100)
	tmps := make([]rune, v8)
	for i := 0; i < v8; i++ {
		tmps[i] = randUTF8RuneHeader(r)
	}
	return string(tmps)
}
func randUnrecognizedHeader(r randyHeader, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldHeader(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldHeader(dAtA []byte, r randyHeader, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateHeader(dAtA, uint64(key))
		v9 := r.Int63()
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		dAtA = encodeVarintPopulateHeader(dAtA, uint64(v9))
	case 1:
		dAtA = encodeVarintPopulateHeader(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateHeader(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateHeader(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateHeader(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateHeader(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *UnsignedHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovHeader(uint64(m.Version))
	}
	l = len(m.PrevBlockHash)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	l = len(m.TransactionsRoot)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	l = len(m.StateRoot)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovHeader(uint64(m.Timestamp))
	}
	if m.Height != 0 {
		n += 1 + sovHeader(uint64(m.Height))
	}
	if m.ConsensusData != 0 {
		n += 1 + sovHeader(uint64(m.ConsensusData))
	}
	l = len(m.WinnerHash)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	if m.WinnerType != 0 {
		n += 1 + sovHeader(uint64(m.WinnerType))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	l = len(m.ChordId)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	return n
}

func (m *BlockHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UnsignedHeader != nil {
		l = m.UnsignedHeader.Size()
		n += 1 + l + sovHeader(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	return n
}

func sovHeader(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHeader(x uint64) (n int) {
	return sovHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UnsignedHeader) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnsignedHeader{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`PrevBlockHash:` + fmt.Sprintf("%v", this.PrevBlockHash) + `,`,
		`TransactionsRoot:` + fmt.Sprintf("%v", this.TransactionsRoot) + `,`,
		`StateRoot:` + fmt.Sprintf("%v", this.StateRoot) + `,`,
		`Timestamp:` + fmt.Sprintf("%v", this.Timestamp) + `,`,
		`Height:` + fmt.Sprintf("%v", this.Height) + `,`,
		`ConsensusData:` + fmt.Sprintf("%v", this.ConsensusData) + `,`,
		`WinnerHash:` + fmt.Sprintf("%v", this.WinnerHash) + `,`,
		`WinnerType:` + fmt.Sprintf("%v", this.WinnerType) + `,`,
		`Signer:` + fmt.Sprintf("%v", this.Signer) + `,`,
		`ChordId:` + fmt.Sprintf("%v", this.ChordId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *BlockHeader) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BlockHeader{`,
		`UnsignedHeader:` + strings.Replace(fmt.Sprintf("%v", this.UnsignedHeader), "UnsignedHeader", "UnsignedHeader", 1) + `,`,
		`Signature:` + fmt.Sprintf("%v", this.Signature) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHeader(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UnsignedHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeader
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
			return fmt.Errorf("proto: UnsignedHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnsignedHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevBlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
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
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevBlockHash = append(m.PrevBlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PrevBlockHash == nil {
				m.PrevBlockHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionsRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
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
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactionsRoot = append(m.TransactionsRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.TransactionsRoot == nil {
				m.TransactionsRoot = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
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
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateRoot = append(m.StateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.StateRoot == nil {
				m.StateRoot = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusData", wireType)
			}
			m.ConsensusData = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConsensusData |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
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
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WinnerHash = append(m.WinnerHash[:0], dAtA[iNdEx:postIndex]...)
			if m.WinnerHash == nil {
				m.WinnerHash = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerType", wireType)
			}
			m.WinnerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WinnerType |= (WinnerType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
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
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = append(m.Signer[:0], dAtA[iNdEx:postIndex]...)
			if m.Signer == nil {
				m.Signer = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChordId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
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
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChordId = append(m.ChordId[:0], dAtA[iNdEx:postIndex]...)
			if m.ChordId == nil {
				m.ChordId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeader
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
func (m *BlockHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeader
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
			return fmt.Errorf("proto: BlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnsignedHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
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
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UnsignedHeader == nil {
				m.UnsignedHeader = &UnsignedHeader{}
			}
			if err := m.UnsignedHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
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
				return ErrInvalidLengthHeader
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
		default:
			iNdEx = preIndex
			skippy, err := skipHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeader
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
func skipHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeader
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
					return 0, ErrIntOverflowHeader
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
					return 0, ErrIntOverflowHeader
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
				return 0, ErrInvalidLengthHeader
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHeader
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
				next, err := skipHeader(dAtA[start:])
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
	ErrInvalidLengthHeader = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeader   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/header.proto", fileDescriptor_header_c1b05b6bee120f74) }

var fileDescriptor_header_c1b05b6bee120f74 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0xe3, 0x76, 0xff, 0x76, 0x7d, 0xbb, 0xa6, 0xfd, 0xfb, 0x80, 0xcc, 0x34, 0x4c, 0x34,
	0x09, 0x14, 0x81, 0x68, 0xa5, 0x71, 0x84, 0x53, 0xa1, 0xda, 0x2a, 0x50, 0x91, 0xd2, 0x21, 0xb8,
	0x45, 0x4e, 0x62, 0x92, 0x08, 0x1a, 0x47, 0xb6, 0x33, 0xb4, 0xdb, 0x3e, 0x02, 0x1f, 0x83, 0x8f,
	0xc0, 0x47, 0xe0, 0xd8, 0xe3, 0x8e, 0x34, 0xbd, 0x70, 0xdc, 0x91, 0x23, 0x8a, 0xd3, 0xae, 0xec,
	0xe6, 0xe7, 0xf7, 0x3e, 0xb6, 0x5f, 0xbf, 0x8f, 0xa1, 0x9f, 0x07, 0xa3, 0x84, 0xb3, 0x88, 0xcb,
	0x61, 0x2e, 0x85, 0x16, 0xb8, 0x91, 0x07, 0x87, 0xcf, 0xe2, 0x54, 0x27, 0x45, 0x30, 0x0c, 0xc5,
	0x62, 0x14, 0x8b, 0x58, 0x8c, 0x4c, 0x29, 0x28, 0x3e, 0x19, 0x65, 0x84, 0x59, 0xd5, 0x5b, 0x0e,
	0x07, 0x79, 0x50, 0x39, 0x62, 0xc9, 0x16, 0x35, 0x39, 0xbe, 0x6a, 0x82, 0xfd, 0x3e, 0x53, 0x69,
	0x9c, 0xf1, 0xe8, 0xcc, 0x9c, 0x8e, 0x09, 0xb4, 0x2f, 0xb8, 0x54, 0xa9, 0xc8, 0x08, 0x72, 0x90,
	0xdb, 0xf3, 0xb6, 0x12, 0x3f, 0x86, 0x7e, 0x2e, 0xf9, 0x85, 0x1f, 0x7c, 0x11, 0xe1, 0x67, 0x3f,
	0x61, 0x2a, 0x21, 0x0d, 0x07, 0xb9, 0x07, 0x5e, 0xaf, 0xc2, 0xe3, 0x8a, 0x9e, 0x31, 0x95, 0xe0,
	0xa7, 0xf0, 0xbf, 0x96, 0x2c, 0x53, 0x2c, 0xd4, 0xa9, 0xc8, 0x94, 0x2f, 0x85, 0xd0, 0xa4, 0x69,
	0x9c, 0x83, 0x7f, 0x0b, 0x9e, 0x10, 0x1a, 0x3f, 0x00, 0x50, 0x9a, 0x69, 0x5e, 0xbb, 0xf6, 0x8c,
	0xab, 0x63, 0x88, 0x29, 0x1f, 0x41, 0x47, 0xa7, 0x0b, 0xae, 0x34, 0x5b, 0xe4, 0xe4, 0x3f, 0x07,
	0xb9, 0x4d, 0x6f, 0x07, 0xf0, 0x3d, 0x68, 0x25, 0x3c, 0x8d, 0x13, 0x4d, 0x5a, 0xa6, 0xd5, 0x8d,
	0xc2, 0x8f, 0xc0, 0x0e, 0x45, 0xa6, 0x78, 0xa6, 0x0a, 0xe5, 0x47, 0x4c, 0x33, 0xd2, 0x76, 0x90,
	0xbb, 0xe7, 0xf5, 0x6e, 0xe9, 0x6b, 0xa6, 0x19, 0x7e, 0x08, 0xdd, 0xaf, 0x69, 0x96, 0x71, 0x59,
	0x3f, 0x66, 0xdf, 0x5c, 0x0e, 0x35, 0x32, 0x2f, 0x19, 0xdd, 0x1a, 0xf4, 0x65, 0xce, 0x49, 0xc7,
	0x41, 0xae, 0x7d, 0x62, 0x0f, 0xf3, 0x60, 0xf8, 0xc1, 0xe0, 0xf3, 0xcb, 0x9c, 0x6f, 0x37, 0x54,
	0xeb, 0xaa, 0x21, 0x33, 0x4c, 0x49, 0xc0, 0x1c, 0xb6, 0x51, 0xf8, 0x3e, 0xec, 0x87, 0x89, 0x90,
	0x91, 0x9f, 0x46, 0xa4, 0x6b, 0x2a, 0x6d, 0xa3, 0xa7, 0xd1, 0x71, 0x02, 0xdd, 0x7a, 0x74, 0xf5,
	0xf8, 0x5f, 0x40, 0xbf, 0xd8, 0x04, 0xe2, 0xd7, 0x79, 0x9b, 0x18, 0xba, 0x27, 0xb8, 0xba, 0xf6,
	0x6e, 0x56, 0x9e, 0x5d, 0xdc, 0xcd, 0xee, 0x08, 0x3a, 0x95, 0x66, 0xba, 0x90, 0x7c, 0x93, 0xcd,
	0x0e, 0x3c, 0x19, 0x03, 0xec, 0xda, 0xc6, 0x18, 0xec, 0xd3, 0xc9, 0x6c, 0x32, 0x9f, 0xce, 0xfd,
	0xf9, 0xf4, 0x74, 0x36, 0xf1, 0x06, 0x16, 0xb6, 0x01, 0xce, 0x3f, 0xce, 0xb6, 0x1a, 0xe1, 0x01,
	0x1c, 0x8c, 0xdf, 0xbe, 0x7b, 0xf5, 0x66, 0x4b, 0x1a, 0xe3, 0x97, 0xcb, 0x15, 0xb5, 0xae, 0x57,
	0xd4, 0xba, 0x59, 0x51, 0xf4, 0x67, 0x45, 0xd1, 0x55, 0x49, 0xd1, 0xf7, 0x92, 0xa2, 0x1f, 0x25,
	0x45, 0x3f, 0x4b, 0x8a, 0x96, 0x25, 0x45, 0xbf, 0x4a, 0x8a, 0x7e, 0x97, 0xd4, 0xba, 0x29, 0x29,
	0xfa, 0xb6, 0xa6, 0xd6, 0x72, 0x4d, 0xad, 0xeb, 0x35, 0xb5, 0x82, 0x96, 0xf9, 0x75, 0xcf, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x27, 0x6c, 0x02, 0xcd, 0x02, 0x00, 0x00,
}
