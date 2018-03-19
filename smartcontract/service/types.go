package service

import (
	. "nkn-core/common"
	"nkn-core/core/asset"
	"nkn-core/core/ledger"
	"nkn-core/core/transaction"
	. "nkn-core/net/httpjsonrpc"
	"nkn-core/net/httprestful/common"
	"nkn-core/smartcontract/states"
)

type AccountInfo struct {
	ProgramHash string
	IsFrozen    bool
	Balances    map[string]Fixed64
}

type AssetInfo struct {
	Name      string
	Precision byte
	AssetType byte
}

func GetHeaderInfo(header *ledger.BlockHeader) *BlockHead {
	h := header.Hash()
	return &BlockHead{
		Version:          header.Version,
		PrevBlockHash:    ToHexString(header.PrevBlockHash.ToArrayReverse()),
		TransactionsRoot: ToHexString(header.TransactionsRoot.ToArrayReverse()),
		Timestamp:        header.Timestamp,
		Height:           header.Height,
		Nonce:            header.Nonce,
		Program: ProgramInfo{
			Code:      ToHexString(header.Program.Code),
			Parameter: ToHexString(header.Program.Parameter),
		},
		Hash: ToHexString(h.ToArrayReverse()),
	}

}

func GetBlockInfo(block *ledger.Block) *BlockInfo {
	blockInfo := common.GetBlockInfo(block)
	return &blockInfo
}

func GetTransactionInfo(transaction *transaction.Transaction) *Transactions {
	return TransArryByteToHexString(transaction)
}

func GetTransactionInputs(inputs []*transaction.UTXOTxInput) []UTXOTxInputInfo {
	inputList := make([]UTXOTxInputInfo, len(inputs))
	for k, v := range inputs {
		inputList[k].ReferTxID = ToHexString(v.ReferTxID.ToArrayReverse())
		inputList[k].ReferTxOutputIndex = v.ReferTxOutputIndex
	}
	return inputList
}

func GetTransactionOutputs(outputs []*transaction.TxOutput) []TxoutputInfo {
	outputList := make([]TxoutputInfo, len(outputs))
	for k, v := range outputs {
		outputList[k].AssetID = ToHexString(v.AssetID.ToArrayReverse())
		outputList[k].Value = v.Value
		outputList[k].ProgramHash = ToHexString(v.ProgramHash.ToArrayReverse())
	}
	return outputList
}

func GetTransactionAttributes(attributes []*transaction.TxAttribute) []TxAttributeInfo {
	attributeList := make([]TxAttributeInfo, len(attributes))
	for k, v := range attributes {
		attributeList[k].Usage = v.Usage
		attributeList[k].Data = ToHexString(v.Data)
	}
	return attributeList
}

func GetAccountInfo(account *states.AccountState) *AccountInfo {
	balances := make(map[string]Fixed64)
	for k, v := range account.Balances {
		assetId := ToHexString(k.ToArrayReverse())
		balances[assetId] = v
	}
	return &AccountInfo{
		ProgramHash: ToHexString(account.ProgramHash.ToArrayReverse()),
		IsFrozen:    account.IsFrozen,
		Balances:    balances,
	}
}

func GetAssetInfo(asset *asset.Asset) *AssetInfo {
	return &AssetInfo{
		Name:      asset.Name,
		Precision: asset.Precision,
		AssetType: byte(asset.AssetType),
	}
}
