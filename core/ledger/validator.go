package ledger

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/core/signature"
	tx "github.com/nknorg/nkn/core/transaction"
	"github.com/nknorg/nkn/core/transaction/payload"
	"github.com/nknorg/nkn/crypto"
	. "github.com/nknorg/nkn/errors"
	"github.com/nknorg/nkn/por"
	"github.com/nknorg/nkn/util/config"
	"github.com/nknorg/nkn/util/log"
)

const (
	GenesisBlockProposedHeight = 4
)

type VBlock struct {
	Block       *Block
	ReceiveTime int64
}

func NewVBlock(b *Block, t int64) *VBlock {
	return &VBlock{
		Block:       b,
		ReceiveTime: t,
	}
}

func TransactionCheck(block *Block) error {
	if block.Transactions == nil {
		return errors.New("empty block")
	}
	if block.Transactions[0].TxType != tx.Coinbase {
		return errors.New("first transaction in block is not Coinbase")
	}
	for i, txn := range block.Transactions {
		if i != 0 && txn.TxType == tx.Coinbase {
			return errors.New("Coinbase transaction order is incorrect")
		}
		if errCode := tx.VerifyTransaction(txn); errCode != ErrNoError {
			return errors.New("transaction sanity check failed")
		}
		if errCode := tx.VerifyTransactionWithLedger(txn); errCode != ErrNoError {
			return errors.New("transaction history check failed")
		}
	}

	return nil
}

func HeaderCheck(header *Header, receiveTime int64, ledger *Ledger) error {
	height := header.Height
	if height == 0 {
		return nil
	}
	prevHeader, err := ledger.Blockchain.GetHeader(header.PrevBlockHash)
	if err != nil {
		return errors.New("prev header doesn't exist")
	}
	if prevHeader == nil {
		return errors.New("invalid prev header")
	}
	if prevHeader.Height+1 != height {
		return errors.New("invalid header height")
	}
	if prevHeader.Timestamp >= header.Timestamp {
		return errors.New("invalid header timestamp")
	}
	if header.WinningHashType == GenesisHash && header.Height >= GenesisBlockProposedHeight {
		return errors.New("invalid winning hash type")
	}

	// calculate time difference
	var timeDiff int64
	genesisBlockHash, err := DefaultLedger.Store.GetBlockHash(0)
	if err != nil {
		return err
	}
	genesisBlock, err := DefaultLedger.Store.GetBlock(genesisBlockHash)
	if err != nil {
		return err
	}
	prevTimestamp, err := DefaultLedger.Blockchain.GetBlockTime(header.PrevBlockHash)
	if err != nil {
		return err
	}
	if prevTimestamp == genesisBlock.Header.Timestamp {
		timeDiff = 0
	} else {
		timeDiff = receiveTime - prevTimestamp
	}

	// get miner who will sign next block
	var miner []byte
	timeSlot := int64(config.ProposerChangeTime / time.Second)
	if timeDiff > timeSlot {
		index := timeDiff / timeSlot
		proposerBlockHeight := int64(DefaultLedger.Store.GetHeight()) - index
		if proposerBlockHeight < 0 {
			proposerBlockHeight = 0
		}
		proposerBlockHash, err := DefaultLedger.Store.GetBlockHash(uint32(proposerBlockHeight))
		if err != nil {
			return err
		}
		proposerBlock, err := DefaultLedger.Store.GetBlock(proposerBlockHash)
		if err != nil {
			return err
		}
		miner, err = proposerBlock.GetSigner()
		log.Infof("timeouts, singer: %s is expected, which is signer of height %d",
			common.BytesToHexString(miner), proposerBlockHeight)
		if err != nil {
			return err
		}
	} else {
		winningHash := prevHeader.WinningHash
		winningHashType := prevHeader.WinningHashType
		switch winningHashType {
		case GenesisHash:
			miner, err = genesisBlock.GetSigner()
			if err != nil {
				return err
			}
		case WinningTxnHash:
			txn, err := DefaultLedger.Store.GetTransaction(winningHash)
			if err != nil {
				return err
			}
			payload, ok := txn.Payload.(*payload.Commit)
			if !ok {
				return errors.New("invalid transaction type")
			}
			sigchain := &por.SigChain{}
			proto.Unmarshal(payload.SigChain, sigchain)
			miner, err = sigchain.GetMiner()
			if err != nil {
				return err
			}
			txnHash := txn.Hash()
			log.Infof("WinningTxnHash, singer: %s is expected, txn hash: %s",
				common.BytesToHexString(miner), common.BytesToHexString(txnHash.ToArrayReverse()))
		case WinningNilHash:
			miner = prevHeader.Signer
			log.Infof("WinningNilHash, singer: %s is expected, which is signer of prev height %d",
				common.BytesToHexString(miner), prevHeader.Height)
		}
	}

	// verify header signature
	if bytes.Compare(miner, header.Signer) != 0 {
		return fmt.Errorf("invalid Block signer, expected: %s", common.BytesToHexString(miner))
	}
	rawPubKey, err := crypto.DecodePoint(miner)
	if err != nil {
		return err
	}
	err = crypto.Verify(*rawPubKey, signature.GetHashForSigning(header), header.Signature)
	if err != nil {
		log.Error("block header verification error: ", err)
		return err
	}

	return nil
}

func BlockCheck(v *VBlock, ledger *Ledger) error {
	err := HeaderCheck(v.Block.Header, v.ReceiveTime, ledger)
	if err != nil {
		return err
	}
	if err := TransactionCheck(v.Block); err != nil {
		return errors.New("transaction check error when verify block")
	}

	return nil
}
