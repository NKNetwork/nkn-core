package voting

import (
	. "github.com/nknorg/nkn/common"
)

type VotingContentType byte

const (
	SigChainVote VotingContentType = 0
	BlockVote    VotingContentType = 1
)

type VotingContent interface {
	Hash() Uint256
}

type Voting interface {
	// prepare to do voting
	Preparing(content VotingContent) error
	// get voting content type
	VotingType() VotingContentType
	// set hash in process
	SetConfirmingHash(hash Uint256)
	// get hash in process
	GetConfirmingHash() Uint256
	// set proposer state
	SetProposerState(hash Uint256, s State)
	// check proposer state
	HasProposerState(hash Uint256, s State) bool
	// set voter state
	SetVoterState(nid uint64, hash Uint256, s State)
	// check voter state
	HasVoterState(nid uint64, hash Uint256, s State) bool
	// get prefer voting content
	GetCurrentVotingContent() (VotingContent, error)
	// get voting content from local by hash
	GetVotingContent(hash Uint256) (VotingContent, error)
	// check if exist in local memory
	Exist(hash Uint256) bool
	// dump consensus state for testing
	DumpState(hash Uint256, desc string, verbose bool)
}
