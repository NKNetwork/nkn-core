package ledger

import (
	"errors"
	"fmt"
	. "nkn-core/errors"
)

func VerifyBlock(block *Block, completely bool) error {
	if block.Header.Height == 0 {
		return nil
	}
	err := VerifyHeader(block.Header)
	if err != nil {
		return err
	}

	flag, err := VerifySignableData(block)
	if flag == false || err != nil {
		return err
	}

	if block.Transactions == nil {
		return errors.New(fmt.Sprintf("No Transactions Exist in Block."))
	}
	if block.Transactions[0].TxType != BookKeeping {
		return errors.New(fmt.Sprintf("BlockHeader Verify failed first Transacion in block is not BookKeeping type."))
	}
	for index, v := range block.Transactions {
		if v.TxType == BookKeeping && index != 0 {
			return errors.New(fmt.Sprintf("This Block Has BookKeeping transaction after first transaction in block."))
		}
	}

	//verfiy block's transactions
	if completely {
		for _, txVerify := range block.Transactions {
			if errCode := VerifyTransaction(txVerify); errCode != ErrNoError {
				return errors.New(fmt.Sprintf("VerifyTransaction failed when verifiy block"))
			}
			if errCode := VerifyTransactionWithLedger(txVerify, DefaultLedger); errCode != ErrNoError {
				return errors.New(fmt.Sprintf("VerifyTransactionWithLedger failed when verifiy block"))
			}
		}
	}

	return nil
}

func VerifyHeader(bd *BlockHeader) error {

	prevHeader, err := DefaultLedger.Blockchain.GetHeader(bd.PrevBlockHash)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[BlockValidator], Cannnot find prevHeader..")
	}
	if prevHeader == nil {
		return NewDetailErr(errors.New("[BlockValidator] error"), ErrNoCode, "[BlockValidator], Cannnot find previous block.")
	}

	if prevHeader.Height+1 != bd.Height {
		return NewDetailErr(errors.New("[BlockValidator] error"), ErrNoCode, "[BlockValidator], block height is incorrect.")
	}

	if prevHeader.Timestamp >= bd.Timestamp {
		return NewDetailErr(errors.New("[BlockValidator] error"), ErrNoCode, "[BlockValidator], block timestamp is incorrect.")
	}

	return nil
}
