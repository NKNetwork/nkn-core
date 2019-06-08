package moca

import (
	"time"

	"github.com/nknorg/nkn/pb"
	"github.com/nknorg/nkn/util/config"
)

const (
	electionStartDelay          = config.ConsensusDuration / 2
	electionDuration            = config.ConsensusDuration / 2
	proposalVerificationTimeout = electionStartDelay * 4 / 5
	minVotingInterval           = 500 * time.Millisecond
	proposingInterval           = 500 * time.Millisecond
	proposingTimeout            = config.ConsensusDuration / 10
	cacheExpiration             = 3600 * time.Second
	cacheCleanupInterval        = 600 * time.Second
	proposingStartDelay         = config.ConsensusTimeout + time.Second
	proposalPropagationDelay    = time.Second
	getConsensusStateInterval   = config.ConsensusDuration
	getConsensusStateRetries    = 3
	getConsensusStateRetryDelay = 3 * time.Second
	proposalChanLen             = 100
	requestProposalChanLen      = 10000
	changeVoteMinRelativeWeight = 0.5
	consensusMinRelativeWeight  = 2.0 / 3.0
	syncMinRelativeWeight       = 1.0 / 2.0
	requestTransactionType      = pb.REQUEST_TRANSACTION_SHORT_HASH
)
