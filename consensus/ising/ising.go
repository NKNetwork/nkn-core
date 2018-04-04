package ising

import (
	"fmt"
	. "nkn-core/common"
	"nkn-core/common/serialization"
	"nkn-core/core/ledger"
	"nkn-core/core/transaction"
	"nkn-core/crypto"
	"nkn-core/events"
	"nkn-core/net"
	"nkn-core/net/message"
	"nkn-core/wallet"
)

const (
	TxnAmountToBePackaged = 1024
)

type Ising struct {
	wallet               wallet.Wallet             // local account
	role                 Role                      // proposer or voter
	localNode            net.Neter                 // local node
	txnCollector         *transaction.TxnCollector // collect transaction from where
	blockCache           *BlockCache               // blocks waiting for voting
	consensusMsgReceived events.Subscriber         // consensus events listening
}

func New(wallet wallet.Wallet, node net.Neter) *Ising {
	var role Role
	account, err := wallet.GetDefaultAccount()
	if err != nil {
		return nil
	}
	encPubKey, err := account.PublicKey.EncodePoint(true)
	if err != nil {
		return nil
	}
	confPubKey, err := ledger.StandbyBookKeepers[0].EncodePoint(true)
	if IsEqualBytes(encPubKey, confPubKey) {
		role = BlockProposer
	} else {
		role = BlockVoter
	}
	ising := &Ising{
		role:         role,
		wallet:       wallet,
		localNode:    node,
		txnCollector: transaction.NewTxnCollector(node.GetTxnPool(), TxnAmountToBePackaged),
		blockCache:   NewCache(),
	}

	return ising
}

func (p *Ising) Start() error {
	p.consensusMsgReceived = p.localNode.GetEvent("consensus").Subscribe(events.EventConsensusMsgReceived, p.ConsensusMsgReceived)
	switch p.role {
	case BlockProposer:
		fmt.Println("I am block proposer")
		block, err := p.BuildBlock()
		if err != nil {
			return err
		}
		err = p.blockCache.AddBlockToCache(block)
		if err != nil {
			return err
		}
		blockFlooding := &BlockFlooding{
			block: block,
		}
		err = p.localNode.Xmit(BuildIsingConsensusPayload(blockFlooding))
		if err != nil {
			return err
		}
	case BlockVoter:
		fmt.Println("I am block voter")
	}

	return nil
}

func (p *Ising) BuildBlock() (*ledger.Block, error) {
	var txnList []*transaction.Transaction
	var txnHashList []Uint256
	account, _ := p.wallet.GetDefaultAccount()
	coinbase, _ := transaction.NewBookKeeperTransaction(account.PublicKey, true, nil, account.PublicKey)
	txnList = append(txnList, coinbase)
	txnHashList = append(txnHashList, coinbase.Hash())
	txns := p.txnCollector.Collect()
	for txnHash, txn := range txns {
		txnList = append(txnList, txn)
		txnHashList = append(txnHashList, txnHash)
	}
	txnRoot, err := crypto.ComputeRoot(txnHashList)
	if err != nil {
		return nil, err
	}
	header := &ledger.Header{
		TransactionsRoot: txnRoot,
	}
	block := &ledger.Block{
		Header:       header,
		Transactions: txnList,
	}

	return block, nil
}

func (p *Ising) ConsensusMsgReceived(v interface{}) {
	if payload, ok := v.(*message.IsingPayload); ok {
		switch payload.PayloadData.(type) {
		case *BlockFlooding:
			fmt.Println("receive block flooding")
		case *BlockProposal:
			fmt.Println("receive block proposal")
		case *BlockRequest:
			fmt.Println("receive block request")
		case *BlockVote:
			fmt.Println("receive block vote")
		}
	}
}

func BuildIsingConsensusPayload(s serialization.SerializableData) *message.IsingPayload {
	return &message.IsingPayload{s}
}
