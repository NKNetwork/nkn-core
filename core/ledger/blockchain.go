package ledger

import (
	"sync"

	. "nkn-core/common"
	"nkn-core/common/log"
	. "nkn-core/errors"
	"nkn-core/events"
)

type Blockchain struct {
	BlockHeight uint32
	BCEvents    *events.Event
	mutex       sync.Mutex
}

func NewBlockchain(height uint32) *Blockchain {
	return &Blockchain{
		BlockHeight: height,
		BCEvents:    events.NewEvent(),
	}
}

func NewBlockchainWithGenesisBlock() (*Blockchain, error) {
	genesisBlock, err := GenesisBlockInit()
	if err != nil {
		return nil, NewDetailErr(err, ErrNoCode, "[Blockchain], NewBlockchainWithGenesisBlock failed.")
	}
	genesisBlock.RebuildMerkleRoot()
	hashx := genesisBlock.Hash()
	genesisBlock.hash = &hashx

	height, err := DefaultLedger.Store.InitLedgerStoreWithGenesisBlock(genesisBlock)
	if err != nil {
		return nil, NewDetailErr(err, ErrNoCode, "[Blockchain], InitLevelDBStoreWithGenesisBlock failed.")
	}
	blockchain := NewBlockchain(height)
	return blockchain, nil
}

func (bc *Blockchain) AddBlock(block *Block) error {
	log.Debug()
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	err := bc.SaveBlock(block)
	if err != nil {
		return err
	}

	return nil
}

func (bc *Blockchain) GetHeader(hash Uint256) (*BlockHeader, error) {
	header, err := DefaultLedger.Store.GetHeader(hash)
	if err != nil {
		return nil, NewDetailErr(err, ErrNoCode, "[Blockchain], GetHeader failed.")
	}
	return header, nil
}

func (bc *Blockchain) SaveBlock(block *Block) error {
	log.Debugf("Save block, block hash %x", block.Hash())
	err := DefaultLedger.Store.SaveBlock(block)
	if err != nil {
		log.Warn("Save block failure , ", err)
		return err
	}

	return nil
}

func (bc *Blockchain) ContainsTransaction(hash Uint256) bool {
	//TODO: implement error catch
	_, err := DefaultLedger.Store.GetTransaction(hash)
	if err != nil {
		return false
	}
	return true
}

func (bc *Blockchain) CurrentBlockHash() Uint256 {
	return DefaultLedger.Store.GetCurrentBlockHash()
}
