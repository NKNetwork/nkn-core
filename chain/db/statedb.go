package db

import (
	"sync"

	"github.com/nknorg/nkn/common"
)

var (
	AccountPrefix        = []byte{0x00}
	NanoPayPrefix        = []byte{0x01}
	NanoPayCleanupPrefix = []byte{0x02}
	NamePrefix           = []byte{0x03}
	NameRegistrantPrefix = []byte{0x04}
)

type StateDB struct {
	db              *cachingDB
	trie            ITrie
	accounts        sync.Map
	nanoPay         map[string]*nanoPay
	nanoPayCleanup  map[uint32]map[string]struct{}
	names           map[string]string
	nameRegistrants map[string][]byte
}

func NewStateDB(root common.Uint256, db *cachingDB) (*StateDB, error) {
	trie, err := db.OpenTrie(root)
	if err != nil {
		return nil, err
	}
	return &StateDB{
		db:             db,
		trie:           trie,
		nanoPay:        make(map[string]*nanoPay, 0),
		nanoPayCleanup: make(map[uint32]map[string]struct{}, 0),
	}, nil
}

func (sdb *StateDB) Finalize(commit bool) (root common.Uint256, err error) {
	sdb.FinalizeAccounts(commit)

	sdb.FinalizeNanoPay(commit)

	sdb.FinalizeNames(commit)

	if commit {
		root, err = sdb.trie.CommitTo()

		return root, err
	}

	return common.EmptyUint256, nil
}

func (sdb *StateDB) IntermediateRoot() common.Uint256 {
	sdb.Finalize(false)
	return sdb.trie.Hash()
}

