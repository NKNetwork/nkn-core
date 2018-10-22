package net

import (
	"github.com/nknorg/nkn/crypto"
	"github.com/nknorg/nkn/net/node"
	"github.com/nknorg/nkn/net/protocol"
	"github.com/nknorg/nnet"
)

func StartProtocol(pubKey *crypto.PubKey, nn *nnet.NNet) (protocol.Noder, error) {
	net, err := node.InitNode(pubKey, nn)
	if err != nil {
		return nil, err
	}

	err = nn.Start()
	if err != nil {
		return nil, err
	}

	return net, nil
}
