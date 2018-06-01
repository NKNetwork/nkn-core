package por

import (
	"bytes"
	"testing"

	"github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/core/transaction"
	"github.com/nknorg/nkn/crypto"
	"github.com/nknorg/nkn/util/log"
	"github.com/nknorg/nkn/wallet"
)

func TestPorPackage(t *testing.T) {
	crypto.SetAlg("P256R1")

	from, _ := wallet.NewAccount()
	rel, _ := wallet.NewAccount()
	to, _ := wallet.NewAccount()
	toPk, _ := to.PubKey().EncodePoint(true)
	relPk, _ := rel.PubKey().EncodePoint(true)

	sc, err := NewSigChain(from, 1, &common.Uint256{}, &common.Uint256{}, toPk, relPk)
	if err != nil {
		t.Error("sigchain created failed")
	}

	err = sc.Sign(toPk, rel)
	if err != nil || sc.Verify() != nil {
		t.Error("'rel' sign in error")
	}

	err = sc.Sign(toPk, to)
	if err != nil || sc.Verify() != nil {
		t.Error("'to' sign in error")
	}

	buff := bytes.NewBuffer(nil)
	sc.Serialize(buff)
	txn, err := transaction.NewCommitTransaction(buff.Bytes(), from.ProgramHash)
	if err != nil {
		log.Error("txn wrong", txn)
	}

	ppkg, err := NewPorPackage(txn)
	if err != nil {
		log.Error("Create por package error", err)
	}

	//test Hash
	ppkgHash := ppkg.Hash()
	if (&ppkgHash).CompareTo(sc.Hash()) != 0 {
		t.Error("[TestPorPackage] Hash test failed")
	}

	//GetblockHash
	if sc.GetblockHash() != ppkg.GetblockHash() {
		t.Error("[TestPorPackage] GetBlockHeight test failed")
	}

	//GetTxHash
	if ppkg.GetTxHash().CompareTo(txn.Hash()) != 0 {
		t.Error("[TestPorPackage] GetTxHash test failed")
	}

	//GetSigChain
	if (&ppkgHash).CompareTo(ppkg.GetSigChain().Hash()) != 0 {
		t.Error("[TestPorPackage] GetSigChain test failed")
	}

	//Serialize & Deserialize
	ppkg2 := new(porPackage)
	buff2 := bytes.NewBuffer(nil)
	ppkg.Serialize(buff2)
	ppkg2.Deserialize(buff2)

	if ppkg.CompareTo(ppkg2) != 0 {
		t.Error("[TestPorPackage] Serialize test failed")
	}
}
