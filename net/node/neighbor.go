package node

import (
	"fmt"
	"sync"

	"github.com/nknorg/nkn/net/protocol"
	nnetnode "github.com/nknorg/nnet/node"
)

// The neighbor node list
type nbrNodes struct {
	List sync.Map
}

func (nm *nbrNodes) GetNbrNode(uid uint64) protocol.Noder {
	v, ok := nm.List.Load(uid)
	if !ok {
		return nil
	}
	n, ok := v.(*Node)
	if ok {
		return n
	}
	return nil
}

func (nm *nbrNodes) AddNbrNode(n protocol.Noder) {
	node, ok := n.(*Node)
	if !ok {
		fmt.Println("Convert the noder error when add node")
		return
	}
	nm.List.LoadOrStore(n.GetID(), node)
}

func (nm *nbrNodes) DelNbrNode(id uint64) (protocol.Noder, bool) {
	v, ok := nm.List.Load(id)
	if !ok {
		return nil, false
	}
	n, _ := v.(*Node)
	nm.List.Delete(id)
	return n, true
}

func (nm *nbrNodes) GetConnectionCnt() uint {
	return uint(len(nm.GetNeighborNoder(nil)))
}

func (nm *nbrNodes) GetNeighborAddrs() ([]protocol.NodeAddr, uint) {
	var addrs []protocol.NodeAddr
	for _, n := range nm.GetNeighborNoder(nil) {
		ip, _ := n.GetAddr16()
		addrs = append(addrs, protocol.NodeAddr{
			IpAddr:  ip,
			IpStr:   n.GetAddr(),
			InOut:   n.GetConnDirection(),
			Time:    n.GetTime(),
			Port:    n.GetPort(),
			ID:      n.GetID(),
			NKNaddr: fmt.Sprintf("%x", n.GetChordAddr()),
		})
	}
	return addrs, uint(len(addrs))
}

func (nm *nbrNodes) GetNeighborHeights() ([]uint32, uint) {
	heights := []uint32{}
	for _, n := range nm.GetNeighborNoder(nil) {
		heights = append(heights, n.GetHeight())
	}
	return heights, uint(len(heights))
}

func (nm *nbrNodes) GetNeighborNoder(filter func(protocol.Noder) bool) []protocol.Noder {
	nodes := []protocol.Noder{}
	nm.List.Range(func(key, value interface{}) bool {
		if n, ok := value.(*Node); ok {
			if filter == nil || filter(n) {
				nodes = append(nodes, n)
			}
		}
		return true
	})
	return nodes
}

func (n *Node) getNbrByNNetNode(remoteNode *nnetnode.RemoteNode) protocol.Noder {
	if remoteNode == nil {
		return nil
	}

	nodeID, err := chordIDToNodeID(remoteNode.Id)
	if err != nil {
		return nil
	}

	nbr := n.GetNbrNode(nodeID)
	return nbr
}
