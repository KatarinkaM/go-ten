package ethereummock

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/obscuronet/obscuro-playground/go/ethclient/mgmtcontractlib"
	"github.com/obscuronet/obscuro-playground/go/obscurocommon"
)

// MockTxHandler implements mgmtcontractlib.TxHandler for the ethereummock package
// it never PackTx (because the mock eth takes care of it)
// it always UnPackTx given the expected direct conversion of types.Transaction.Data -> obscurocommon.L1TxData
type MockTxHandler struct{}

func (m MockTxHandler) PackTx(*obscurocommon.L1TxData, common.Address, uint64) (types.TxData, error) {
	panic("implement me")
}

func (m MockTxHandler) UnPackTx(tx *types.Transaction) *obscurocommon.L1TxData {
	t := obscurocommon.TxData(tx)
	return &t
}

func NewMockTxHandler() mgmtcontractlib.TxHandler {
	return &MockTxHandler{}
}