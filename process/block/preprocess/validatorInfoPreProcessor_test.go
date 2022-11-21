package preprocess

import (
	"bytes"
	"errors"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/state"
	"github.com/ElrondNetwork/elrond-go/testscommon"
	"github.com/ElrondNetwork/elrond-go/testscommon/enableEpochsHandlerMock"
	"github.com/ElrondNetwork/elrond-go/testscommon/genericMocks"
	"github.com/ElrondNetwork/elrond-go/testscommon/hashingMocks"
	"github.com/ElrondNetwork/elrond-go/testscommon/marshallerMock"
	"github.com/ElrondNetwork/elrond-go/testscommon/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewValidatorInfoPreprocessor_NilHasherShouldErr(t *testing.T) {
	t.Parallel()

	tdp := initDataPool()
	rtp, err := NewValidatorInfoPreprocessor(
		nil,
		&marshallerMock.MarshalizerMock{},
		&testscommon.BlockSizeComputationStub{},
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	assert.Nil(t, rtp)
	assert.Equal(t, process.ErrNilHasher, err)
}

func TestNewValidatorInfoPreprocessor_NilMarshalizerShouldErr(t *testing.T) {
	t.Parallel()

	tdp := initDataPool()
	rtp, err := NewValidatorInfoPreprocessor(
		&hashingMocks.HasherMock{},
		nil,
		&testscommon.BlockSizeComputationStub{},
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	assert.Nil(t, rtp)
	assert.Equal(t, process.ErrNilMarshalizer, err)
}

func TestNewValidatorInfoPreprocessor_NilBlockSizeComputationHandlerShouldErr(t *testing.T) {
	t.Parallel()

	tdp := initDataPool()
	rtp, err := NewValidatorInfoPreprocessor(
		&hashingMocks.HasherMock{},
		&marshallerMock.MarshalizerMock{},
		nil,
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	assert.Nil(t, rtp)
	assert.Equal(t, process.ErrNilBlockSizeComputationHandler, err)
}

func TestNewValidatorInfoPreprocessor_NilValidatorInfoPoolShouldErr(t *testing.T) {
	t.Parallel()

	rtp, err := NewValidatorInfoPreprocessor(
		&hashingMocks.HasherMock{},
		&marshallerMock.MarshalizerMock{},
		&testscommon.BlockSizeComputationStub{},
		nil,
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	assert.Nil(t, rtp)
	assert.Equal(t, process.ErrNilValidatorInfoPool, err)
}

func TestNewValidatorInfoPreprocessor_NilStoreShouldErr(t *testing.T) {
	t.Parallel()

	tdp := initDataPool()
	rtp, err := NewValidatorInfoPreprocessor(
		&hashingMocks.HasherMock{},
		&marshallerMock.MarshalizerMock{},
		&testscommon.BlockSizeComputationStub{},
		tdp.ValidatorsInfo(),
		nil,
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	assert.Nil(t, rtp)
	assert.Equal(t, process.ErrNilStorage, err)
}

func TestNewValidatorInfoPreprocessor_NilEnableEpochHandlerShouldErr(t *testing.T) {
	t.Parallel()

	tdp := initDataPool()
	rtp, err := NewValidatorInfoPreprocessor(
		&hashingMocks.HasherMock{},
		&marshallerMock.MarshalizerMock{},
		&testscommon.BlockSizeComputationStub{},
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		nil,
	)

	assert.Nil(t, rtp)
	assert.Equal(t, process.ErrNilEnableEpochsHandler, err)
}

func TestNewValidatorInfoPreprocessor_OkValsShouldWork(t *testing.T) {
	t.Parallel()

	tdp := initDataPool()
	rtp, err := NewValidatorInfoPreprocessor(
		&hashingMocks.HasherMock{},
		&marshallerMock.MarshalizerMock{},
		&testscommon.BlockSizeComputationStub{},
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)
	assert.Nil(t, err)
	assert.NotNil(t, rtp)
}

func TestNewValidatorInfoPreprocessor_CreateMarshalizedDataShouldWork(t *testing.T) {
	t.Parallel()

	tdp := initDataPool()
	rtp, _ := NewValidatorInfoPreprocessor(
		&hashingMocks.HasherMock{},
		&marshallerMock.MarshalizerMock{},
		&testscommon.BlockSizeComputationStub{},
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	hash := make([][]byte, 0)
	res, err := rtp.CreateMarshalledData(hash)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(res))
}

func TestNewValidatorInfoPreprocessor_ProcessMiniBlockInvalidMiniBlockTypeShouldErr(t *testing.T) {
	t.Parallel()

	tdp := initDataPool()
	rtp, _ := NewValidatorInfoPreprocessor(
		&hashingMocks.HasherMock{},
		&marshallerMock.MarshalizerMock{},
		&testscommon.BlockSizeComputationStub{},
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	txHashes := make([][]byte, 0)
	mb1 := block.MiniBlock{
		TxHashes:        txHashes,
		ReceiverShardID: 1,
		SenderShardID:   0,
		Type:            0,
	}

	preProcessorExecutionInfoHandlerMock := &testscommon.PreProcessorExecutionInfoHandlerMock{
		GetNumOfCrossInterMbsAndTxsCalled: getNumOfCrossInterMbsAndTxsZero,
	}

	_, _, _, err := rtp.ProcessMiniBlock(&mb1, haveTimeTrue, haveAdditionalTimeFalse, false, false, -1, preProcessorExecutionInfoHandlerMock)
	assert.Equal(t, process.ErrWrongTypeInMiniBlock, err)
}

func TestNewValidatorInfoPreprocessor_ProcessMiniBlockShouldWork(t *testing.T) {
	t.Parallel()

	tdp := initDataPool()
	rtp, _ := NewValidatorInfoPreprocessor(
		&hashingMocks.HasherMock{},
		&marshallerMock.MarshalizerMock{},
		&testscommon.BlockSizeComputationStub{},
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	txHashes := make([][]byte, 0)
	mb1 := block.MiniBlock{
		TxHashes:        txHashes,
		ReceiverShardID: 1,
		SenderShardID:   core.MetachainShardId,
		Type:            block.PeerBlock,
	}

	preProcessorExecutionInfoHandlerMock := &testscommon.PreProcessorExecutionInfoHandlerMock{
		GetNumOfCrossInterMbsAndTxsCalled: getNumOfCrossInterMbsAndTxsZero,
	}

	_, _, _, err := rtp.ProcessMiniBlock(&mb1, haveTimeTrue, haveAdditionalTimeFalse, false, false, -1, preProcessorExecutionInfoHandlerMock)
	assert.Nil(t, err)
}

func TestNewValidatorInfoPreprocessor_ProcessMiniBlockNotFromMeta(t *testing.T) {
	t.Parallel()

	tdp := initDataPool()
	rtp, _ := NewValidatorInfoPreprocessor(
		&hashingMocks.HasherMock{},
		&marshallerMock.MarshalizerMock{},
		&testscommon.BlockSizeComputationStub{},
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	txHashes := make([][]byte, 0)
	mb1 := block.MiniBlock{
		TxHashes:        txHashes,
		ReceiverShardID: 1,
		SenderShardID:   0,
		Type:            block.PeerBlock,
	}

	preProcessorExecutionInfoHandlerMock := &testscommon.PreProcessorExecutionInfoHandlerMock{
		GetNumOfCrossInterMbsAndTxsCalled: getNumOfCrossInterMbsAndTxsZero,
	}

	_, _, _, err := rtp.ProcessMiniBlock(&mb1, haveTimeTrue, haveAdditionalTimeFalse, false, false, -1, preProcessorExecutionInfoHandlerMock)
	assert.Equal(t, process.ErrValidatorInfoMiniBlockNotFromMeta, err)
}

func TestNewValidatorInfoPreprocessor_RestorePeerBlockIntoPools(t *testing.T) {
	t.Parallel()

	hasher := &hashingMocks.HasherMock{}
	marshalizer := &marshallerMock.MarshalizerMock{}
	blockSizeComputation := &testscommon.BlockSizeComputationStub{}

	tdp := initDataPool()
	rtp, _ := NewValidatorInfoPreprocessor(
		hasher,
		marshalizer,
		blockSizeComputation,
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	txHashes := [][]byte{[]byte("tx_hash1")}
	mb1 := block.MiniBlock{
		TxHashes:        txHashes,
		ReceiverShardID: core.AllShardId,
		SenderShardID:   core.MetachainShardId,
		Type:            block.PeerBlock,
	}

	blockBody := &block.Body{}
	blockBody.MiniBlocks = append(blockBody.MiniBlocks, &mb1)
	miniBlockPool := testscommon.NewCacherMock()

	marshalizedMb, _ := marshalizer.Marshal(mb1)
	mbHash := hasher.Compute(string(marshalizedMb))

	foundMb, ok := miniBlockPool.Get(mbHash)
	assert.Nil(t, foundMb)
	assert.False(t, ok)

	numRestoredTxs, err := rtp.RestoreBlockDataIntoPools(blockBody, miniBlockPool)
	assert.Equal(t, 1, numRestoredTxs)
	assert.Nil(t, err)

	foundMb, ok = miniBlockPool.Get(mbHash)
	assert.NotNil(t, foundMb)
	assert.True(t, ok)
}

func TestNewValidatorInfoPreprocessor_RestoreOtherBlockTypeIntoPoolsShouldNotRestore(t *testing.T) {
	t.Parallel()

	hasher := &hashingMocks.HasherMock{}
	marshalizer := &marshallerMock.MarshalizerMock{}
	blockSizeComputation := &testscommon.BlockSizeComputationStub{}

	tdp := initDataPool()
	rtp, _ := NewValidatorInfoPreprocessor(
		hasher,
		marshalizer,
		blockSizeComputation,
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	txHashes := [][]byte{[]byte("tx_hash1")}
	mb1 := block.MiniBlock{
		TxHashes:        txHashes,
		ReceiverShardID: core.AllShardId,
		SenderShardID:   core.MetachainShardId,
		Type:            block.TxBlock,
	}

	blockBody := &block.Body{}
	blockBody.MiniBlocks = append(blockBody.MiniBlocks, &mb1)
	miniBlockPool := testscommon.NewCacherMock()

	marshalizedMb, _ := marshalizer.Marshal(mb1)
	mbHash := hasher.Compute(string(marshalizedMb))

	foundMb, ok := miniBlockPool.Get(mbHash)
	assert.Nil(t, foundMb)
	assert.False(t, ok)

	numRestoredTxs, err := rtp.RestoreBlockDataIntoPools(blockBody, miniBlockPool)
	assert.Equal(t, 0, numRestoredTxs)
	assert.Nil(t, err)

	foundMb, ok = miniBlockPool.Get(mbHash)
	assert.Nil(t, foundMb)
	assert.False(t, ok)
}

func TestNewValidatorInfoPreprocessor_RemovePeerBlockFromPool(t *testing.T) {
	t.Parallel()

	hasher := &hashingMocks.HasherMock{}
	marshalizer := &marshallerMock.MarshalizerMock{}
	blockSizeComputation := &testscommon.BlockSizeComputationStub{}

	tdp := initDataPool()
	rtp, _ := NewValidatorInfoPreprocessor(
		hasher,
		marshalizer,
		blockSizeComputation,
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	txHashes := [][]byte{[]byte("tx_hash1")}
	mb1 := block.MiniBlock{
		TxHashes:        txHashes,
		ReceiverShardID: core.AllShardId,
		SenderShardID:   core.MetachainShardId,
		Type:            block.PeerBlock,
	}

	marshalizedMb, _ := marshalizer.Marshal(mb1)
	mbHash := hasher.Compute(string(marshalizedMb))

	blockBody := &block.Body{}
	blockBody.MiniBlocks = append(blockBody.MiniBlocks, &mb1)
	miniBlockPool := testscommon.NewCacherMock()
	miniBlockPool.Put(mbHash, marshalizedMb, len(marshalizedMb))

	foundMb, ok := miniBlockPool.Get(mbHash)
	assert.NotNil(t, foundMb)
	assert.True(t, ok)

	err := rtp.RemoveBlockDataFromPools(blockBody, miniBlockPool)
	assert.Nil(t, err)

	foundMb, ok = miniBlockPool.Get(mbHash)
	assert.Nil(t, foundMb)
	assert.False(t, ok)
}

func TestNewValidatorInfoPreprocessor_RemoveOtherBlockTypeFromPoolShouldNotRemove(t *testing.T) {
	t.Parallel()

	hasher := &hashingMocks.HasherMock{}
	marshalizer := &marshallerMock.MarshalizerMock{}
	blockSizeComputation := &testscommon.BlockSizeComputationStub{}

	tdp := initDataPool()
	rtp, _ := NewValidatorInfoPreprocessor(
		hasher,
		marshalizer,
		blockSizeComputation,
		tdp.ValidatorsInfo(),
		genericMocks.NewChainStorerMock(0),
		&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	)

	txHashes := [][]byte{[]byte("tx_hash1")}
	mb1 := block.MiniBlock{
		TxHashes:        txHashes,
		ReceiverShardID: core.AllShardId,
		SenderShardID:   core.MetachainShardId,
		Type:            block.TxBlock,
	}

	marshalizedMb, _ := marshalizer.Marshal(mb1)
	mbHash := hasher.Compute(string(marshalizedMb))

	blockBody := &block.Body{}
	blockBody.MiniBlocks = append(blockBody.MiniBlocks, &mb1)
	miniBlockPool := testscommon.NewCacherMock()
	miniBlockPool.Put(mbHash, marshalizedMb, len(marshalizedMb))

	foundMb, ok := miniBlockPool.Get(mbHash)
	assert.NotNil(t, foundMb)
	assert.True(t, ok)

	err := rtp.RemoveBlockDataFromPools(blockBody, miniBlockPool)
	assert.Nil(t, err)

	foundMb, ok = miniBlockPool.Get(mbHash)
	assert.NotNil(t, foundMb)
	assert.True(t, ok)
}

func TestNewValidatorInfoPreprocessor_RestoreValidatorsInfo(t *testing.T) {
	t.Parallel()

	t.Run("restore validators info with not all txs found in storage", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("error")
		hasher := &hashingMocks.HasherMock{}
		marshalizer := &marshallerMock.MarshalizerMock{}
		blockSizeComputation := &testscommon.BlockSizeComputationStub{}
		storer := &storage.ChainStorerStub{
			GetAllCalled: func(unitType dataRetriever.UnitType, keys [][]byte) (map[string][]byte, error) {
				return nil, expectedErr
			},
		}
		tdp := initDataPool()
		rtp, _ := NewValidatorInfoPreprocessor(
			hasher,
			marshalizer,
			blockSizeComputation,
			tdp.ValidatorsInfo(),
			storer,
			&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
		)

		miniBlock := &block.MiniBlock{}
		err := rtp.restoreValidatorsInfo(miniBlock)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("restore validators info with all txs found in storage", func(t *testing.T) {
		t.Parallel()

		hasher := &hashingMocks.HasherMock{}
		marshalizer := &marshallerMock.MarshalizerMock{}
		blockSizeComputation := &testscommon.BlockSizeComputationStub{}
		shardValidatorInfoHash := []byte("hash")
		shardValidatorInfo := &state.ShardValidatorInfo{
			PublicKey: []byte("x"),
		}
		marshalledShardValidatorInfo, _ := marshalizer.Marshal(shardValidatorInfo)
		storer := &storage.ChainStorerStub{
			GetAllCalled: func(unitType dataRetriever.UnitType, keys [][]byte) (map[string][]byte, error) {
				allShardValidatorsInfo := make(map[string][]byte)
				allShardValidatorsInfo[string(shardValidatorInfoHash)] = marshalledShardValidatorInfo
				return allShardValidatorsInfo, nil
			},
		}
		tdp := initDataPool()
		wasCalledWithExpectedKey := false
		tdp.ValidatorsInfoCalled = func() dataRetriever.ShardedDataCacherNotifier {
			return &testscommon.ShardedDataStub{
				AddDataCalled: func(key []byte, data interface{}, sizeInBytes int, cacheID string) {
					if bytes.Equal(key, shardValidatorInfoHash) {
						wasCalledWithExpectedKey = true
					}
				},
			}
		}
		rtp, _ := NewValidatorInfoPreprocessor(
			hasher,
			marshalizer,
			blockSizeComputation,
			tdp.ValidatorsInfo(),
			storer,
			&enableEpochsHandlerMock.EnableEpochsHandlerStub{},
		)

		miniBlock := &block.MiniBlock{}
		err := rtp.restoreValidatorsInfo(miniBlock)
		assert.Nil(t, err)
		assert.True(t, wasCalledWithExpectedKey)
	})
}
