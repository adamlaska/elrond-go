package mock

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
)

// GasHandlerMock -
type GasHandlerMock struct {
	InitCalled                           func()
	SetGasConsumedCalled                 func(gasConsumed uint64, hash []byte)
	SetGasRefundedCalled                 func(gasRefunded uint64, hash []byte)
	GasConsumedCalled                    func(hash []byte) uint64
	GasRefundedCalled                    func(hash []byte) uint64
	TotalGasConsumedCalled               func() uint64
	TotalGasRefundedCalled               func() uint64
	RemoveGasConsumedCalled              func(hashes [][]byte)
	RemoveGasRefundedCalled              func(hashes [][]byte)
	ComputeGasConsumedByMiniBlockCalled  func(miniBlock *block.MiniBlock, mapHashTx map[string]data.TransactionHandler) (uint64, uint64, error)
	ComputeGasConsumedByTxCalled         func(txSenderShardId uint32, txReceiverSharedId uint32, txHandler data.TransactionHandler) (uint64, uint64, error)
	AddTotalGasConsumedInSelfShardCalled func(gasConsumed uint64)
	GetTotalGasConsumedInSelfShardCalled func() uint64
}

// AddTotalGasConsumedInSelfShard -
func (ghm *GasHandlerMock) AddTotalGasConsumedInSelfShard(gasConsumed uint64) {
	if ghm.AddTotalGasConsumedInSelfShardCalled != nil {
		ghm.AddTotalGasConsumedInSelfShardCalled(gasConsumed)
	}
}

// GetTotalGasConsumedInSelfShard -
func (ghm *GasHandlerMock) GetTotalGasConsumedInSelfShard() uint64 {
	if ghm.GetTotalGasConsumedInSelfShardCalled != nil {
		return ghm.GetTotalGasConsumedInSelfShardCalled()
	}

	return 0
}

// Init -
func (ghm *GasHandlerMock) Init() {
	if ghm.InitCalled != nil {
		ghm.InitCalled()
	}
}

// SetGasConsumed -
func (ghm *GasHandlerMock) SetGasConsumed(gasConsumed uint64, hash []byte) {
	if ghm.SetGasConsumedCalled != nil {
		ghm.SetGasConsumedCalled(gasConsumed, hash)
	}
}

// SetGasRefunded -
func (ghm *GasHandlerMock) SetGasRefunded(gasRefunded uint64, hash []byte) {
	if ghm.SetGasRefundedCalled != nil {
		ghm.SetGasRefundedCalled(gasRefunded, hash)
	}
}

// GasConsumed -
func (ghm *GasHandlerMock) GasConsumed(hash []byte) uint64 {
	if ghm.GasConsumedCalled != nil {
		return ghm.GasConsumedCalled(hash)
	}
	return 0
}

// GasRefunded -
func (ghm *GasHandlerMock) GasRefunded(hash []byte) uint64 {
	if ghm.GasRefundedCalled != nil {
		return ghm.GasRefundedCalled(hash)
	}
	return 0
}

// TotalGasConsumed -
func (ghm *GasHandlerMock) TotalGasConsumed() uint64 {
	if ghm.TotalGasConsumedCalled != nil {
		return ghm.TotalGasConsumedCalled()
	}
	return 0
}

// TotalGasRefunded -
func (ghm *GasHandlerMock) TotalGasRefunded() uint64 {
	if ghm.TotalGasRefundedCalled != nil {
		return ghm.TotalGasRefundedCalled()
	}
	return 0
}

// RemoveGasConsumed -
func (ghm *GasHandlerMock) RemoveGasConsumed(hashes [][]byte) {
	if ghm.RemoveGasConsumedCalled != nil {
		ghm.RemoveGasConsumedCalled(hashes)
	}
}

// RemoveGasRefunded -
func (ghm *GasHandlerMock) RemoveGasRefunded(hashes [][]byte) {
	if ghm.RemoveGasRefundedCalled != nil {
		ghm.RemoveGasRefundedCalled(hashes)
	}
}

// ComputeGasConsumedByMiniBlock -
func (ghm *GasHandlerMock) ComputeGasConsumedByMiniBlock(miniBlock *block.MiniBlock, mapHashTx map[string]data.TransactionHandler) (uint64, uint64, error) {
	if ghm.ComputeGasConsumedByMiniBlockCalled != nil {
		return ghm.ComputeGasConsumedByMiniBlockCalled(miniBlock, mapHashTx)
	}
	return 0, 0, nil
}

// ComputeGasConsumedByTx -
func (ghm *GasHandlerMock) ComputeGasConsumedByTx(txSenderShardId uint32, txReceiverShardId uint32, txHandler data.TransactionHandler) (uint64, uint64, error) {
	if ghm.ComputeGasConsumedByTxCalled != nil {
		return ghm.ComputeGasConsumedByTxCalled(txSenderShardId, txReceiverShardId, txHandler)
	}
	return 0, 0, nil
}

// IsInterfaceNil -
func (ghm *GasHandlerMock) IsInterfaceNil() bool {
	return ghm == nil
}
