package por

import (
	"bytes"
	"errors"
	"io"

	"github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/common/serialization"
	"github.com/nknorg/nkn/core/ledger"
	"github.com/nknorg/nkn/core/transaction"
	"github.com/nknorg/nkn/core/transaction/payload"
	"github.com/nknorg/nkn/util/log"
)

type porPackage struct {
	owner        []byte
	height       uint32
	blockHash    common.Uint256
	txHash       common.Uint256
	sigchainHash common.Uint256
	sigchain     *SigChain
}

type porPackages []*porPackage

func (c porPackages) Len() int {
	return len(c)
}
func (c porPackages) Swap(i, j int) {
	if i >= 0 && i < len(c) && j >= 0 && j < len(c) { // Unit Test modify
		c[i], c[j] = c[j], c[i]
	}
}
func (c porPackages) Less(i, j int) bool {
	if i >= 0 && i < len(c) && j >= 0 && j < len(c) { // Unit Test modify
		return c[i].CompareTo(c[j]) < 0
	}

	return false
}

func NewPorPackage(txn *transaction.Transaction) (*porPackage, error) {
	if txn.TxType != transaction.Commit {
		return nil, errors.New("Transaction type mismatch")
	}
	rs := txn.Payload.(*payload.Commit)
	buf := bytes.NewBuffer(rs.SigChain)
	var sigchain SigChain
	sigchain.Deserialize(buf)

	//TODO threshold

	blockHeader, err := ledger.DefaultLedger.Store.GetHeader(*sigchain.blockHash)
	if err != nil {
		log.Error("Get block header error:", err)
		return nil, err
	}

	owner, err := sigchain.GetOwner()
	if err != nil {
		log.Error("Get owner error:", err)
		return nil, err
	}

	pp := &porPackage{
		owner:        owner,
		height:       blockHeader.Height,
		blockHash:    *sigchain.blockHash,
		txHash:       txn.Hash(),
		sigchainHash: sigchain.Hash(),
		sigchain:     &sigchain,
	}
	return pp, nil
}

func (pp *porPackage) Hash() common.Uint256 {
	return pp.sigchainHash
}

func (pp *porPackage) GetHeight() uint32 {
	return pp.height
}

func (pp *porPackage) GetblockHash() *common.Uint256 {
	return &pp.blockHash
}

func (pp *porPackage) GetTxHash() *common.Uint256 {
	return &pp.txHash
}

func (pp *porPackage) GetSigChain() *SigChain {
	return pp.sigchain
}

func (pp *porPackage) CompareTo(o *porPackage) int {
	return (&pp.sigchainHash).CompareTo(o.sigchainHash)
}

func (pp *porPackage) Serialize(w io.Writer) error {
	var err error

	err = serialization.WriteVarBytes(w, pp.owner)
	if err != nil {
		return err
	}

	err = serialization.WriteUint32(w, pp.height)
	if err != nil {
		return err
	}

	_, err = pp.txHash.Serialize(w)
	if err != nil {
		return err
	}

	_, err = pp.sigchainHash.Serialize(w)
	if err != nil {
		return err
	}

	err = pp.sigchain.Serialize(w)

	return err
}

func (pp *porPackage) Deserialize(r io.Reader) error {
	var err error

	pp.owner, err = serialization.ReadVarBytes(r)
	if err != nil {
		return err
	}

	pp.height, err = serialization.ReadUint32(r)
	if err != nil {
		return err
	}

	err = pp.txHash.Deserialize(r)
	if err != nil {
		return err
	}

	err = pp.sigchainHash.Deserialize(r)
	if err != nil {
		return err
	}

	sigchain := new(SigChain)
	err = sigchain.Deserialize(r)
	if err != nil {
		return nil
	}
	pp.sigchain = sigchain

	return nil
}

func (pp *porPackage) DumpInfo() {
	log.Info("owner: ", common.BytesToHexString(pp.owner))
	log.Info("txHash: ", pp.txHash)
	log.Info("sigchainHash: ", pp.sigchainHash)
	sc := pp.sigchain
	sc.DumpInfo()
}
