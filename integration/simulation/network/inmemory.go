package network

import (
	"time"

	"github.com/obscuronet/obscuro-playground/go/obscuronode/obscuroclient"

	"github.com/obscuronet/obscuro-playground/go/ethclient"

	"github.com/obscuronet/obscuro-playground/integration/simulation/p2p"

	"github.com/obscuronet/obscuro-playground/integration/simulation/params"

	"github.com/obscuronet/obscuro-playground/integration/simulation/stats"

	"github.com/obscuronet/obscuro-playground/go/obscuronode/host"
	ethereum_mock "github.com/obscuronet/obscuro-playground/integration/ethereummock"
)

type basicNetworkOfInMemoryNodes struct {
	ethNodes       []*ethereum_mock.Node
	obscuroNodes   []*host.Node
	obscuroClients []*obscuroclient.Client
}

func NewBasicNetworkOfInMemoryNodes() Network {
	return &basicNetworkOfInMemoryNodes{}
}

// Create inits and starts the nodes, wires them up, and populates the network objects
func (n *basicNetworkOfInMemoryNodes) Create(params *params.SimParams, stats *stats.Stats) ([]ethclient.EthClient, []*obscuroclient.Client, []string) {
	l1Clients := make([]ethclient.EthClient, params.NumberOfNodes)
	n.ethNodes = make([]*ethereum_mock.Node, params.NumberOfNodes)
	n.obscuroNodes = make([]*host.Node, params.NumberOfNodes)
	n.obscuroClients = make([]*obscuroclient.Client, params.NumberOfNodes)

	for i := 0; i < params.NumberOfNodes; i++ {
		isGenesis := i == 0

		// create the in memory l1 and l2 node
		miner := createMockEthNode(int64(i), params.NumberOfNodes, params.AvgBlockDuration, params.AvgNetworkLatency, stats)
		agg := createInMemObscuroNode(int64(i), isGenesis, params.TxHandler, params.AvgGossipPeriod, params.AvgBlockDuration, params.AvgNetworkLatency, stats, false, nil)
		obscuroClient := host.NewInMemObscuroClient(int64(i), &agg.P2p, agg.DB(), &agg.EnclaveClient)

		// and connect them to each other
		agg.ConnectToEthNode(miner)
		miner.AddClient(agg)

		n.ethNodes[i] = miner
		n.obscuroNodes[i] = agg
		n.obscuroClients[i] = &obscuroClient
		l1Clients[i] = miner
	}

	// populate the nodes field of each network
	for i := 0; i < params.NumberOfNodes; i++ {
		n.ethNodes[i].Network.(*ethereum_mock.MockEthNetwork).AllNodes = n.ethNodes
		n.obscuroNodes[i].P2p.(*p2p.MockP2P).Nodes = n.obscuroNodes
	}

	// The sequence of starting the nodes is important to catch various edge cases.
	// Here we first start the mock layer 1 nodes, with a pause between them of a fraction of a block duration.
	// The reason is to make sure that they catch up correctly.
	// Then we pause for a while, to give the L1 network enough time to create a number of blocks, which will have to be ingested by the Obscuro nodes
	// Then, we begin the starting sequence of the Obscuro nodes, again with a delay between them, to test that they are able to cach up correctly.
	// Note: Other simulations might test variations of this pattern.
	for _, m := range n.ethNodes {
		t := m
		go t.Start()
		time.Sleep(params.AvgBlockDuration / 8)
	}

	time.Sleep(params.AvgBlockDuration * 20)
	for _, m := range n.obscuroNodes {
		t := m
		go t.Start()
		time.Sleep(params.AvgBlockDuration / 3)
	}

	return l1Clients, n.obscuroClients, nil
}

func (n *basicNetworkOfInMemoryNodes) TearDown() {
	for _, client := range n.obscuroClients {
		temp := client
		go (*temp).Stop()
	}

	for _, node := range n.obscuroNodes {
		temp := node
		go temp.Stop()
	}

	for _, node := range n.ethNodes {
		temp := node
		go temp.Stop()
	}
}
