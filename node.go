package main

import (
	"net"
	"sync"
	"time"
)

//Node is represnter of physical servers(nodes)
type Node struct {
	IPADDR   string
	Conn     net.Conn
	NodeId   string
	lastSeen time.Time
}

func NewNode(nodeId string, c net.Conn) (Node, bool) {
	node := Node{
		NodeId:   nodeId,
		IPADDR:   c.RemoteAddr().String(),
		Conn:     c,
		lastSeen: time.Now(),
	}
	return node, true
}

type NodeList struct {
	nodes map[string]Node
	mu    sync.RWMutex
}

//Add Node to list if its not present
func (n *NodeList) Add(nodeId string, c net.Conn) (Node, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if node, ok := n.nodes[nodeId]; ok {
		// println("node found: ", node)
		return node, true
	}
	node, ok := NewNode(nodeId, c)
	if ok {
		n.nodes[nodeId] = node
		return node, true
	}
	return node, false
}
