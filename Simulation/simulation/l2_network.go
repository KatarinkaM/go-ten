package simulation

import (
	"simulation/common"
	"simulation/obscuro"
	"time"
)

// NetworkCfg - models a full network including artificial random latencies
type L2NetworkCfg struct {
	nodes []*obscuro.Node
	delay common.Latency // the latency
}

// Broadcasts the rollup to all L2 peers
func (c *L2NetworkCfg) BroadcastRollup(r common.Rollup) {
	for _, a := range c.nodes {
		if a.Id != r.Agg {
			t := a
			common.Schedule(c.delay(), func() { t.P2PGossipRollup(&r) })
		}
	}
}

func (c *L2NetworkCfg) BroadcastTx(tx common.L2Tx) {
	for _, a := range c.nodes {
		t := a
		common.Schedule(c.delay()/2, func() { t.P2PReceiveTx(tx) })
	}
}

func (n *L2NetworkCfg) Start(delay time.Duration) {
	// Start l1 nodes
	for _, m := range n.nodes {
		t := m
		go t.Start()
		// don't start everything at once
		time.Sleep(delay)
	}
}

func (n *L2NetworkCfg) Stop() {
	for _, m := range n.nodes {
		m.Stop()
		//fmt.Printf("Stopped L2 node: %d\n", m.Id)
	}
}