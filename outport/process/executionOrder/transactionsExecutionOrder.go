package executionOrder

import (
	"encoding/hex"
	"fmt"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	txsSort "github.com/ElrondNetwork/elrond-go-core/core/transaction"
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/outport"
	"github.com/ElrondNetwork/elrond-go-core/data/smartContractResult"
	"github.com/ElrondNetwork/elrond-go-core/hashing"
	logger "github.com/ElrondNetwork/elrond-go-logger"
	"github.com/ElrondNetwork/elrond-go/process"
)

var log = logger.GetOrCreate("outport/process/executionOrder")

type sorter struct {
	hasher hashing.Hasher
}

// NewSorter will create a new instance of sorter
func NewSorter(hasher hashing.Hasher) (*sorter, error) {
	if check.IfNil(hasher) {
		return nil, process.ErrNilHasher
	}

	return &sorter{
		hasher: hasher,
	}, nil
}

// PutExecutionOrderInTransactionPool will put the execution order for every transaction and smart contract result
func (s *sorter) PutExecutionOrderInTransactionPool(
	pool *outport.Pool,
	header data.HeaderHandler,
	body data.BodyHandler,
) error {
	blockBody, ok := body.(*block.Body)
	if !ok {
		log.Warn("s.PutExecutionOrderInTransactionPool cannot cast BodyHandler to *Body")
		return nil
	}

	// already sorted
	transactionsToMe, scheduledTransactionsToMe, err := extractNormalTransactionAndScrsToMe(pool, blockBody, header)
	if err != nil {
		return err
	}

	// need to be sorted
	transactionsFromMe, scheduledTransactionsFromMe, err := extractNormalTransactionsAndInvalidFromMe(pool, blockBody, header)
	if err != nil {
		return err
	}

	txsSort.SortTransactionsBySenderAndNonceWithFrontRunningProtectionExtendedTransactions(transactionsFromMe, s.hasher, header.GetPrevRandSeed())

	rewardsTxs, err := getRewardsTxsFromMe(pool, blockBody, header)
	if err != nil {
		return err
	}

	// scheduled from me, need to be sorted
	txsSort.SortTransactionsBySenderAndNonceWithFrontRunningProtectionExtendedTransactions(scheduledTransactionsFromMe, s.hasher, header.GetPrevRandSeed())

	allTransaction := append(transactionsToMe, transactionsFromMe...)
	allTransaction = append(allTransaction, rewardsTxs...)
	allTransaction = append(allTransaction, scheduledTransactionsToMe...)
	allTransaction = append(allTransaction, scheduledTransactionsFromMe...)

	for idx, tx := range allTransaction {
		tx.SetExecutionOrder(idx)
	}

	setOrderSmartContractResults(pool)

	printPool(pool)
	return nil
}

func setOrderSmartContractResults(pool *outport.Pool) {
	for _, scrHandler := range pool.Scrs {
		scr, ok := scrHandler.GetTxHandler().(*smartContractResult.SmartContractResult)
		if !ok {
			continue
		}

		tx, found := pool.Txs[string(scr.OriginalTxHash)]
		if !found {
			continue
		}

		scrHandler.SetExecutionOrder(tx.GetExecutionOrder())
	}
}

func extractNormalTransactionsAndInvalidFromMe(pool *outport.Pool, blockBody *block.Body, header data.HeaderHandler) ([]data.TransactionHandlerWithGasUsedAndFee, []data.TransactionHandlerWithGasUsedAndFee, error) {
	transactionsFromMe := make([]data.TransactionHandlerWithGasUsedAndFee, 0)
	scheduledTransactionsFromMe := make([]data.TransactionHandlerWithGasUsedAndFee, 0)

	appendTxs := func(mbIndex int, txs []data.TransactionHandlerWithGasUsedAndFee) {
		if isMBScheduled(header, mbIndex) {
			scheduledTransactionsFromMe = append(scheduledTransactionsFromMe, txs...)
			return
		}
		transactionsFromMe = append(transactionsFromMe, txs...)
	}

	var err error
	var txs []data.TransactionHandlerWithGasUsedAndFee
	for mbIndex, mb := range blockBody.MiniBlocks {
		if err != nil {
			return nil, nil, err
		}

		if shouldIgnoreProcessedMBScheduled(header, mbIndex) {
			continue
		}

		isFromMe := mb.SenderShardID == header.GetShardID()
		if !isFromMe {
			continue
		}

		if mb.Type == block.TxBlock {
			txs, err = extractTxsFromMap(mb.TxHashes, pool.Txs)
			appendTxs(mbIndex, txs)
			continue
		}
		if mb.Type == block.InvalidBlock {
			txs, err = extractTxsFromMap(mb.TxHashes, pool.Invalid)
			appendTxs(mbIndex, txs)
			continue
		}
	}

	return transactionsFromMe, scheduledTransactionsFromMe, nil
}

func extractNormalTransactionAndScrsToMe(pool *outport.Pool, blockBody *block.Body, header data.HeaderHandler) ([]data.TransactionHandlerWithGasUsedAndFee, []data.TransactionHandlerWithGasUsedAndFee, error) {
	transactionsToMe := make([]data.TransactionHandlerWithGasUsedAndFee, 0)
	scheduledTransactionsToMe := make([]data.TransactionHandlerWithGasUsedAndFee, 0)

	appendTxs := func(mbIndex int, txs []data.TransactionHandlerWithGasUsedAndFee) {
		if isMBScheduled(header, mbIndex) {
			scheduledTransactionsToMe = append(scheduledTransactionsToMe, txs...)
			return
		}
		transactionsToMe = append(transactionsToMe, txs...)
	}

	var err error
	var txs []data.TransactionHandlerWithGasUsedAndFee
	for mbIndex, mb := range blockBody.MiniBlocks {
		if err != nil {
			return nil, nil, err
		}

		if shouldIgnoreProcessedMBScheduled(header, mbIndex) {
			continue
		}

		isToMeCross := mb.ReceiverShardID == header.GetShardID() && mb.SenderShardID != mb.ReceiverShardID
		if !isToMeCross {
			continue
		}

		executedTxsHashes := extractExecutedTxHashes(mbIndex, mb.TxHashes, header)
		if mb.Type == block.TxBlock {
			txs, err = extractTxsFromMap(executedTxsHashes, pool.Txs)
			appendTxs(mbIndex, txs)
			continue
		}
		if mb.Type == block.SmartContractResultBlock {
			txs, err = extractTxsFromMap(executedTxsHashes, pool.Scrs)
			appendTxs(mbIndex, txs)
			continue
		}
		if mb.Type == block.RewardsBlock {
			txs, err = extractTxsFromMap(executedTxsHashes, pool.Rewards)
			appendTxs(mbIndex, txs)
			continue
		}
	}

	return transactionsToMe, scheduledTransactionsToMe, nil
}

func getRewardsTxsFromMe(pool *outport.Pool, blockBody *block.Body, header data.HeaderHandler) ([]data.TransactionHandlerWithGasUsedAndFee, error) {
	rewardsTxsHashes := make([][]byte, 0)
	rewardsTxs := make([]data.TransactionHandlerWithGasUsedAndFee, 0)
	if header.GetShardID() != core.MetachainShardId {
		return rewardsTxs, nil
	}

	for _, mb := range blockBody.MiniBlocks {
		if mb.Type != block.RewardsBlock {
			continue
		}
		rewardsTxsHashes = append(rewardsTxsHashes, mb.TxHashes...)
	}

	return extractTxsFromMap(rewardsTxsHashes, pool.Rewards)
}

func extractTxsFromMap(txsHashes [][]byte, txs map[string]data.TransactionHandlerWithGasUsedAndFee) ([]data.TransactionHandlerWithGasUsedAndFee, error) {
	result := make([]data.TransactionHandlerWithGasUsedAndFee, 0, len(txsHashes))
	for _, txHash := range txsHashes {
		tx, found := txs[string(txHash)]
		if !found {
			return nil, fmt.Errorf("cannot find transaction in pool, txHash: %s", hex.EncodeToString(txHash))
		}
		result = append(result, tx)
	}

	return result, nil
}

func extractExecutedTxHashes(mbIndex int, mbTxHashes [][]byte, header data.HeaderHandler) [][]byte {
	miniblockHeaders := header.GetMiniBlockHeaderHandlers()
	if len(miniblockHeaders) <= mbIndex {
		return mbTxHashes
	}

	firstProcessed := miniblockHeaders[mbIndex].GetIndexOfFirstTxProcessed()
	lastProcessed := miniblockHeaders[mbIndex].GetIndexOfLastTxProcessed()

	return mbTxHashes[firstProcessed : lastProcessed+1]
}

func shouldIgnoreProcessedMBScheduled(header data.HeaderHandler, mbIndex int) bool {
	return getProcessingType(header, mbIndex) == int32(block.Processed)
}

func isMBScheduled(header data.HeaderHandler, mbIndex int) bool {
	return getProcessingType(header, mbIndex) == int32(block.Scheduled)
}

func getProcessingType(header data.HeaderHandler, mbIndex int) int32 {
	miniblockHeaders := header.GetMiniBlockHeaderHandlers()
	if len(miniblockHeaders) <= mbIndex {
		return int32(block.Normal)
	}

	return miniblockHeaders[mbIndex].GetProcessingType()
}

// TODO remove this after system test will pass
func printPool(pool *outport.Pool) {
	printMapTxs := func(txs map[string]data.TransactionHandlerWithGasUsedAndFee) {
		for hash, tx := range txs {
			log.Warn(hex.EncodeToString([]byte(hash)), "order", tx.GetExecutionOrder())
		}
	}

	total := len(pool.Txs) + len(pool.Invalid) + len(pool.Scrs) + len(pool.Rewards)
	if total > 0 {
		log.Warn("###################################")
	}

	if len(pool.Txs) > 0 {
		log.Warn("############### NORMAL TXS ####################")
		printMapTxs(pool.Txs)
	}
	if len(pool.Invalid) > 0 {
		log.Warn("############### INVALID ####################")
		printMapTxs(pool.Invalid)
	}

	if len(pool.Scrs) > 0 {
		log.Warn("############### SCRS ####################")
		printMapTxs(pool.Scrs)
	}

	if len(pool.Rewards) > 0 {
		log.Warn("############### REWARDS ####################")
		printMapTxs(pool.Rewards)
	}
	if total > 0 {
		log.Warn("###################################")
	}
}
