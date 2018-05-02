package por

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"

	"github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/common/serialization"
	"github.com/nknorg/nkn/crypto"
	"github.com/nknorg/nkn/wallet"
)

type sigchainer interface {
	Sign(nextPubkey []byte, signer wallet.Account) error
	Verify() error
	Path() [][]byte
	Length() int
	IsFinal() bool
	GetSignerIndex(pubkey []byte) int
	GetDataHash() *common.Uint256
}

// for the first relay node
// 1. NewSigChain : create a new Signature Chain and sign
//
// for the next relay node
// 1. Sign: sign the element created in Sign

type SigChain struct {
	nonce      [4]byte
	dataSize   uint32          // payload size
	dataHash   *common.Uint256 // payload hash
	srcPubkey  []byte          // source pubkey
	destPubkey []byte          // destination pubkey
	elems      []*SigChainElem
}

type SigChainElem struct {
	pubkey     []byte // current signer
	nextPubkey []byte // next signer
	signature  []byte // signature for signature chain element
}

// first relay node starts a new signature chain which consists of meta data and the first element.
func NewSigChain(owner *wallet.Account, dataSize uint32, dataHash *common.Uint256, destPubkey []byte, nextPubkey []byte) (*SigChain, error) {
	ownPk := owner.PubKey()
	srcPubkey, err := ownPk.EncodePoint(true)
	if err != nil {
		return nil, errors.New("EncodePoint ownpk error")
	}

	sc := &SigChain{
		dataSize:   dataSize,
		dataHash:   dataHash,
		srcPubkey:  srcPubkey,
		destPubkey: destPubkey,
		elems: []*SigChainElem{
			&SigChainElem{
				pubkey:     srcPubkey,
				nextPubkey: nextPubkey,
			},
		},
	}

	rand.Read(sc.nonce[:])
	buff := bytes.NewBuffer(nil)
	if err := sc.SerializationMetadata(buff); err != nil {
		return nil, err
	}
	elem := sc.elems[0]
	if err := elem.SerializationUnsigned(buff); err != nil {
		return nil, err
	}

	hash := sha256.Sum256(buff.Bytes())
	signature, err := crypto.Sign(owner.PrivKey(), hash[:])
	if err != nil {
		return nil, err
	}
	sc.elems[0].signature = signature

	return sc, nil
}

func NewSigChainElem(pubKey []byte, nextPubkey []byte) *SigChainElem {
	return &SigChainElem{
		pubkey:     pubKey,
		nextPubkey: nextPubkey,
		signature:  nil,
	}
}

// Sign new created signature chain with local wallet.
func (p *SigChain) Sign(nextPubkey []byte, signer *wallet.Account) error {
	sigNum := p.Length()
	if sigNum < 1 {
		return errors.New("there are not enough signatures")
	}

	if err := p.Verify(); err != nil {
		return err
	}

	pk, err := p.nextSigner()
	if err != nil {
		return err
	}

	//TODO decode nextpk or encode signer pubkey
	nxPk, err := crypto.DecodePoint(pk)
	if err != nil {
		return errors.New("the next pubkey is wrong")
	}

	if !crypto.Equal(signer.PubKey(), nxPk) {
		return errors.New("signer is not the one")
	}

	lastElem, err := p.lastSigElem()
	//buff := bytes.NewBuffer(nil)
	//err = serialization.WriteVarBytes(buff, lastElem.signature)
	//if err != nil {
	//	return err
	//}
	buff := bytes.NewBuffer(lastElem.signature)
	elem := NewSigChainElem(pk, nextPubkey)
	err = elem.SerializationUnsigned(buff)
	if err != nil {
		return err
	}
	hash := sha256.Sum256(buff.Bytes())
	signature, err := crypto.Sign(signer.PrivKey(), hash[:])
	if err != nil {
		return err
	}
	elem.signature = signature
	p.elems = append(p.elems, elem)

	return nil
}

// Verify returns result of signature chain verification.
func (p *SigChain) Verify() error {

	prevNextPubkey := p.srcPubkey
	buff := bytes.NewBuffer(nil)
	p.SerializationMetadata(buff)
	prevSig := buff.Bytes()
	for _, e := range p.elems {
		// verify each element public key is correct
		if !common.IsEqualBytes(prevNextPubkey, e.pubkey) {
			return errors.New("unmatch public key in signature chain")
		}

		ePk, err := crypto.DecodePoint(e.pubkey)
		if err != nil {
			return errors.New("the pubkey of e is wrong")
		}

		// verify each element signature
		buff := bytes.NewBuffer(prevSig)
		//	serialization.WriteVarBytes(buff, prevSig)
		e.SerializationUnsigned(buff)
		currHash := sha256.Sum256(buff.Bytes())
		err = crypto.Verify(*ePk, currHash[:], e.signature)
		if err != nil {
			return err
		}

		prevNextPubkey = e.nextPubkey
		prevSig = e.signature
	}

	return nil
}

// Path returns signer path in signature chain.
func (p *SigChain) Path() [][]byte {
	var publicKeys [][]byte
	for _, e := range p.elems {
		publicKeys = append(publicKeys, e.pubkey)
	}

	return publicKeys
}

// Length returns element num in current signature chain
func (p *SigChain) Length() int {
	return len(p.elems)
}

func (p *SigChain) GetDataHash() *common.Uint256 {
	return p.dataHash
}

// firstSigElem returns the first element in signature chain.
func (p *SigChain) firstSigElem() (*SigChainElem, error) {
	if p == nil || len(p.elems) == 0 {
		return nil, errors.New("nil signature chain")
	}

	return p.elems[0], nil
}

// lastSigElem returns the last element in signature chain.
func (p *SigChain) lastSigElem() (*SigChainElem, error) {
	if p == nil || len(p.elems) == 0 {
		return nil, errors.New("nil signature chain")
	}
	num := len(p.elems)

	return p.elems[num-1], nil
}

func (p *SigChain) finalSigElem() (*SigChainElem, error) {
	if !common.IsEqualBytes(p.destPubkey, p.elems[len(p.elems)-1].pubkey) {
		return nil, errors.New("unfinal")
	}

	return p.elems[len(p.elems)-1], nil
}

func (p *SigChain) IsFinal() bool {
	if !common.IsEqualBytes(p.destPubkey, p.elems[len(p.elems)-1].pubkey) {
		return false
	}
	return true
}

func (p *SigChain) getElemByPubkey(pubkey []byte) (*SigChainElem, int, error) {
	if p == nil || len(p.elems) == 0 {
		return nil, 0, errors.New("nil signature chain")
	}
	for i, elem := range p.elems {
		if common.IsEqualBytes(elem.pubkey, pubkey) {
			return elem, i, nil
		}
	}

	return nil, 0, errors.New("not in sigchain")
}

func (p *SigChain) getElemByIndex(idx int) (*SigChainElem, error) {
	if p == nil || len(p.elems) < idx {
		return nil, errors.New("nil signature chain")
	}

	return p.elems[idx], nil
}

func (p *SigChain) GetSignerIndex(pubkey []byte) (int, error) {
	_, idx, err := p.getElemByPubkey(pubkey)
	return idx, err
}

func (p *SigChain) nextSigner() ([]byte, error) {
	e, err := p.lastSigElem()
	if err != nil {
		return nil, errors.New("no elem")
	}
	return e.nextPubkey, nil
}

func (p *SigChain) Serialize(w io.Writer) error {
	var err error
	err = p.SerializationMetadata(w)
	if err != nil {
		return err
	}

	num := len(p.elems)
	err = serialization.WriteUint32(w, uint32(num))
	if err != nil {
		return err
	}
	for i := 0; i < num; i++ {
		err = p.elems[i].Serialize(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *SigChain) SerializationMetadata(w io.Writer) error {
	var err error

	err = serialization.WriteVarBytes(w, p.nonce[:])
	if err != nil {
		return err
	}

	err = serialization.WriteUint32(w, p.dataSize)
	if err != nil {
		return err
	}

	_, err = p.dataHash.Serialize(w)
	if err != nil {
		return err
	}

	err = serialization.WriteVarBytes(w, p.srcPubkey)
	if err != nil {
		return err
	}

	err = serialization.WriteVarBytes(w, p.destPubkey)
	if err != nil {
		return err
	}

	return nil
}

func (p *SigChain) Deserialize(r io.Reader) error {
	var err error
	err = p.DeserializationMetadata(r)
	if err != nil {
		return err
	}

	num, err := serialization.ReadUint32(r)
	if err != nil {
		return err
	}
	for i := 0; i < int(num); i++ {
		err = p.elems[i].Deserialize(r)
		return err
	}

	return nil
}

func (p *SigChain) DeserializationMetadata(r io.Reader) error {
	var err error

	nonce, err := serialization.ReadVarBytes(r)
	if err != nil {
		return err
	}
	copy(p.nonce[:], nonce)

	dataSize, err := serialization.ReadUint32(r)
	if err != nil {
		return err
	}
	p.dataSize = dataSize
	err = p.dataHash.Deserialize(r)
	if err != nil {
		return err
	}

	p.srcPubkey, err = serialization.ReadVarBytes(r)
	if err != nil {
		return err
	}

	p.destPubkey, err = serialization.ReadVarBytes(r)
	if err != nil {
		return err
	}

	return nil
}

func (p *SigChainElem) Serialize(w io.Writer) error {
	var err error
	err = p.SerializationUnsigned(w)
	if err != nil {
		return err
	}
	err = serialization.WriteVarBytes(w, p.signature[:])
	if err != nil {
		return err
	}

	return nil
}

func (p *SigChainElem) SerializationUnsigned(w io.Writer) error {
	var err error

	err = serialization.WriteVarBytes(w, p.pubkey)
	if err != nil {
		return err
	}

	err = serialization.WriteVarBytes(w, p.nextPubkey)
	if err != nil {
		return err
	}

	return nil
}

func (p *SigChainElem) Deserialize(r io.Reader) error {
	var err error
	err = p.DeserializationUnsigned(r)
	if err != nil {
		return err
	}
	signature, err := serialization.ReadVarBytes(r)
	if err != nil {
		return err
	}
	p.signature = signature

	return nil
}

func (p *SigChainElem) DeserializationUnsigned(r io.Reader) error {
	var err error

	p.pubkey, err = serialization.ReadVarBytes(r)
	if err != nil {
		return err
	}

	p.nextPubkey, err = serialization.ReadVarBytes(r)
	if err != nil {
		return err
	}

	return nil
}

func (p *SigChain) dump() {
	fmt.Println("dataSize: ", p.dataSize)
	fmt.Println("dataHash: ", p.dataHash)
	fmt.Println("srcPubkey: ", common.BytesToHexString(p.srcPubkey))
	fmt.Println("dstPubkey: ", common.BytesToHexString(p.destPubkey))

	for i, e := range p.elems {
		fmt.Printf("curPubkey[%d]: %s\n", i, common.BytesToHexString(e.pubkey))
		fmt.Printf("nextPubkey[%d]: %s\n", i, common.BytesToHexString(e.nextPubkey))
		fmt.Printf("signature[%d]: %s\n", i, common.BytesToHexString(e.signature))
	}
}
