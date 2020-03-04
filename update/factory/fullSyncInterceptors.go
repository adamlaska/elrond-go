package factory

import (
	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/core/throttler"
	"github.com/ElrondNetwork/elrond-go/crypto"
	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/ElrondNetwork/elrond-go/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/hashing"
	"github.com/ElrondNetwork/elrond-go/marshal"
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/process/dataValidators"
	"github.com/ElrondNetwork/elrond-go/process/factory"
	"github.com/ElrondNetwork/elrond-go/process/interceptors"
	interceptorFactory "github.com/ElrondNetwork/elrond-go/process/interceptors/factory"
	"github.com/ElrondNetwork/elrond-go/process/interceptors/processor"
	"github.com/ElrondNetwork/elrond-go/process/mock"
	"github.com/ElrondNetwork/elrond-go/sharding"
	"github.com/ElrondNetwork/elrond-go/update"
)

var _ process.InterceptorsContainerFactory = (*fullSyncInterceptorsContainerFactory)(nil)

const numGoRoutines = 2000

// fullSyncInterceptorsContainerFactory will handle the creation the interceptors container for shards
type fullSyncInterceptorsContainerFactory struct {
	container              process.InterceptorsContainer
	shardCoordinator       sharding.Coordinator
	accounts               state.AccountsAdapter
	marshalizer            marshal.Marshalizer
	hasher                 hashing.Hasher
	store                  dataRetriever.StorageService
	dataPool               dataRetriever.PoolsHolder
	messenger              process.TopicHandler
	multiSigner            crypto.MultiSigner
	nodesCoordinator       sharding.NodesCoordinator
	blackList              process.BlackListHandler
	argInterceptorFactory  *interceptorFactory.ArgInterceptedDataFactory
	globalThrottler        process.InterceptorThrottler
	maxTxNonceDeltaAllowed int
	keyGen                 crypto.KeyGenerator
	singleSigner           crypto.SingleSigner
	addrConverter          state.AddressConverter
	whiteListHandler       update.WhiteListHandler
}

// ArgsNewFullSyncInterceptorsContainerFactory holds the arguments needed for fullSyncInterceptorsContainerFactory
type ArgsNewFullSyncInterceptorsContainerFactory struct {
	Accounts               state.AccountsAdapter
	ShardCoordinator       sharding.Coordinator
	NodesCoordinator       sharding.NodesCoordinator
	Messenger              process.TopicHandler
	Store                  dataRetriever.StorageService
	Marshalizer            marshal.Marshalizer
	Hasher                 hashing.Hasher
	KeyGen                 crypto.KeyGenerator
	BlockSignKeyGen        crypto.KeyGenerator
	SingleSigner           crypto.SingleSigner
	BlockSingleSigner      crypto.SingleSigner
	MultiSigner            crypto.MultiSigner
	DataPool               dataRetriever.PoolsHolder
	AddrConverter          state.AddressConverter
	MaxTxNonceDeltaAllowed int
	TxFeeHandler           process.FeeHandler
	BlackList              process.BlackListHandler
	HeaderSigVerifier      process.InterceptedHeaderSigVerifier
	ChainID                []byte
	SizeCheckDelta         uint32
	ValidityAttester       process.ValidityAttester
	EpochStartTrigger      process.EpochStartTriggerHandler
	WhiteListHandler       update.WhiteListHandler
	InterceptorsContainer  process.InterceptorsContainer
}

// NewFullSyncInterceptorsContainerFactory is responsible for creating a new interceptors factory object
func NewFullSyncInterceptorsContainerFactory(
	args ArgsNewFullSyncInterceptorsContainerFactory,
) (*fullSyncInterceptorsContainerFactory, error) {
	if args.SizeCheckDelta > 0 {
		args.Marshalizer = marshal.NewSizeCheckUnmarshalizer(args.Marshalizer, args.SizeCheckDelta)
	}
	err := checkBaseParams(
		args.ShardCoordinator,
		args.Accounts,
		args.Marshalizer,
		args.Hasher,
		args.Store,
		args.DataPool,
		args.Messenger,
		args.MultiSigner,
		args.NodesCoordinator,
		args.BlackList,
	)
	if err != nil {
		return nil, err
	}

	if check.IfNil(args.KeyGen) {
		return nil, process.ErrNilKeyGen
	}
	if check.IfNil(args.SingleSigner) {
		return nil, process.ErrNilSingleSigner
	}
	if check.IfNil(args.AddrConverter) {
		return nil, process.ErrNilAddressConverter
	}
	if check.IfNil(args.TxFeeHandler) {
		return nil, process.ErrNilEconomicsFeeHandler
	}
	if check.IfNil(args.BlockSignKeyGen) {
		return nil, process.ErrNilKeyGen
	}
	if check.IfNil(args.BlockSingleSigner) {
		return nil, process.ErrNilSingleSigner
	}
	if check.IfNil(args.HeaderSigVerifier) {
		return nil, process.ErrNilHeaderSigVerifier
	}
	if len(args.ChainID) == 0 {
		return nil, process.ErrInvalidChainID
	}
	if check.IfNil(args.ValidityAttester) {
		return nil, process.ErrNilValidityAttester
	}
	if check.IfNil(args.EpochStartTrigger) {
		return nil, process.ErrNilEpochStartTrigger
	}
	if check.IfNil(args.InterceptorsContainer) {
		return nil, update.ErrNilInterceptorsContainer
	}
	if check.IfNil(args.WhiteListHandler) {
		return nil, update.ErrNilWhiteListHandler
	}

	argInterceptorFactory := &interceptorFactory.ArgInterceptedDataFactory{
		Marshalizer:       args.Marshalizer,
		Hasher:            args.Hasher,
		ShardCoordinator:  args.ShardCoordinator,
		MultiSigVerifier:  args.MultiSigner,
		NodesCoordinator:  args.NodesCoordinator,
		KeyGen:            args.KeyGen,
		BlockKeyGen:       args.BlockSignKeyGen,
		Signer:            args.SingleSigner,
		BlockSigner:       args.BlockSingleSigner,
		AddrConv:          args.AddrConverter,
		FeeHandler:        args.TxFeeHandler,
		HeaderSigVerifier: args.HeaderSigVerifier,
		ChainID:           args.ChainID,
		ValidityAttester:  args.ValidityAttester,
		EpochStartTrigger: args.EpochStartTrigger,
	}

	icf := &fullSyncInterceptorsContainerFactory{
		container:              args.InterceptorsContainer,
		accounts:               args.Accounts,
		shardCoordinator:       args.ShardCoordinator,
		messenger:              args.Messenger,
		store:                  args.Store,
		marshalizer:            args.Marshalizer,
		hasher:                 args.Hasher,
		multiSigner:            args.MultiSigner,
		dataPool:               args.DataPool,
		nodesCoordinator:       args.NodesCoordinator,
		argInterceptorFactory:  argInterceptorFactory,
		blackList:              args.BlackList,
		maxTxNonceDeltaAllowed: args.MaxTxNonceDeltaAllowed,
		keyGen:                 args.KeyGen,
		singleSigner:           args.SingleSigner,
		addrConverter:          args.AddrConverter,
		whiteListHandler:       args.WhiteListHandler,
	}

	icf.globalThrottler, err = throttler.NewNumGoRoutineThrottler(numGoRoutines)
	if err != nil {
		return nil, err
	}

	return icf, nil
}

// Create returns an interceptor container that will hold all interceptors in the system
func (ficf *fullSyncInterceptorsContainerFactory) Create() (process.InterceptorsContainer, error) {
	err := ficf.generateTxInterceptors()
	if err != nil {
		return nil, err
	}

	err = ficf.generateUnsignedTxsInterceptors()
	if err != nil {
		return nil, err
	}

	err = ficf.generateRewardTxInterceptors()
	if err != nil {
		return nil, err
	}

	err = ficf.generateMiniBlocksInterceptors()
	if err != nil {
		return nil, err
	}

	err = ficf.generateMetachainHeaderInterceptors()
	if err != nil {
		return nil, err
	}

	err = ficf.generateShardHeaderInterceptors()
	if err != nil {
		return nil, err
	}

	err = ficf.generateTrieNodesInterceptors()
	if err != nil {
		return nil, err
	}

	err = ficf.setWhiteListHandlerToInterceptors()
	if err != nil {
		return nil, err
	}

	return ficf.container, nil
}

func checkBaseParams(
	shardCoordinator sharding.Coordinator,
	accounts state.AccountsAdapter,
	marshalizer marshal.Marshalizer,
	hasher hashing.Hasher,
	store dataRetriever.StorageService,
	dataPool dataRetriever.PoolsHolder,
	messenger process.TopicHandler,
	multiSigner crypto.MultiSigner,
	nodesCoordinator sharding.NodesCoordinator,
	blackList process.BlackListHandler,
) error {
	if check.IfNil(shardCoordinator) {
		return process.ErrNilShardCoordinator
	}
	if check.IfNil(messenger) {
		return process.ErrNilMessenger
	}
	if check.IfNil(store) {
		return process.ErrNilStore
	}
	if check.IfNil(marshalizer) {
		return process.ErrNilMarshalizer
	}
	if check.IfNil(hasher) {
		return process.ErrNilHasher
	}
	if check.IfNil(multiSigner) {
		return process.ErrNilMultiSigVerifier
	}
	if check.IfNil(dataPool) {
		return process.ErrNilDataPoolHolder
	}
	if check.IfNil(nodesCoordinator) {
		return process.ErrNilNodesCoordinator
	}
	if check.IfNil(accounts) {
		return process.ErrNilAccountsAdapter
	}
	if check.IfNil(blackList) {
		return process.ErrNilBlackListHandler
	}

	return nil
}

func (ficf *fullSyncInterceptorsContainerFactory) checkIfInterceptorExists(identifier string) bool {
	_, err := ficf.container.Get(identifier)
	return err == nil
}

func (ficf *fullSyncInterceptorsContainerFactory) generateShardHeaderInterceptors() error {
	noOfShards := ficf.shardCoordinator.NumberOfShards()
	tmpSC, err := sharding.NewMultiShardCoordinator(noOfShards, core.MetachainShardId)
	if err != nil {
		return err
	}

	keys := make([]string, noOfShards)
	interceptorsSlice := make([]process.Interceptor, noOfShards)

	//wire up to topics: shardBlocks_0_META, shardBlocks_1_META ...
	for idx := uint32(0); idx < noOfShards; idx++ {
		identifierHeader := factory.ShardBlocksTopic + tmpSC.CommunicationIdentifier(idx)
		if ficf.checkIfInterceptorExists(identifierHeader) {
			continue
		}

		interceptor, err := ficf.createOneShardHeaderInterceptor(identifierHeader)
		if err != nil {
			return err
		}

		keys[int(idx)] = identifierHeader
		interceptorsSlice[int(idx)] = interceptor
	}

	return ficf.container.AddMultiple(keys, interceptorsSlice)
}

func (ficf *fullSyncInterceptorsContainerFactory) createOneShardHeaderInterceptor(topic string) (process.Interceptor, error) {
	hdrValidator, err := dataValidators.NewNilHeaderValidator()
	if err != nil {
		return nil, err
	}

	hdrFactory, err := interceptorFactory.NewInterceptedShardHeaderDataFactory(ficf.argInterceptorFactory)
	if err != nil {
		return nil, err
	}

	argProcessor := &processor.ArgHdrInterceptorProcessor{
		Headers:      ficf.dataPool.Headers(),
		HdrValidator: hdrValidator,
		BlackList:    ficf.blackList,
	}
	hdrProcessor, err := processor.NewHdrInterceptorProcessor(argProcessor)
	if err != nil {
		return nil, err
	}

	interceptor, err := interceptors.NewSingleDataInterceptor(
		hdrFactory,
		hdrProcessor,
		ficf.globalThrottler,
	)
	if err != nil {
		return nil, err
	}

	return ficf.createTopicAndAssignHandler(topic, interceptor, true)
}

func (ficf *fullSyncInterceptorsContainerFactory) generateUnsignedTxsInterceptors() error {
	shardC := ficf.shardCoordinator

	noOfShards := shardC.NumberOfShards()

	keys := make([]string, noOfShards+1)
	interceptorsSlice := make([]process.Interceptor, noOfShards+1)

	for idx := uint32(0); idx < noOfShards; idx++ {
		identifierScr := factory.UnsignedTransactionTopic + shardC.CommunicationIdentifier(idx)
		if ficf.checkIfInterceptorExists(identifierScr) {
			continue
		}

		interceptor, err := ficf.createOneUnsignedTxInterceptor(identifierScr)
		if err != nil {
			return err
		}

		keys[int(idx)] = identifierScr
		interceptorsSlice[int(idx)] = interceptor
	}

	identifierScr := factory.UnsignedTransactionTopic + shardC.CommunicationIdentifier(core.MetachainShardId)
	if !ficf.checkIfInterceptorExists(identifierScr) {
		interceptor, err := ficf.createOneUnsignedTxInterceptor(identifierScr)
		if err != nil {
			return err
		}

		keys[noOfShards] = identifierScr
		interceptorsSlice[noOfShards] = interceptor
	}

	return ficf.container.AddMultiple(keys, interceptorsSlice)
}

func (ficf *fullSyncInterceptorsContainerFactory) generateTrieNodesInterceptors() error {
	noOfShards := ficf.shardCoordinator.NumberOfShards()

	keys := make([]string, 0)
	trieInterceptors := make([]process.Interceptor, 0)

	for i := uint32(0); i < noOfShards; i++ {
		identifierTrieNodes := factory.AccountTrieNodesTopic + core.CommunicationIdentifierBetweenShards(i, core.MetachainShardId)
		if ficf.checkIfInterceptorExists(identifierTrieNodes) {
			continue
		}

		interceptor, err := ficf.createOneTrieNodesInterceptor(identifierTrieNodes)
		if err != nil {
			return err
		}

		keys = append(keys, identifierTrieNodes)
		trieInterceptors = append(trieInterceptors, interceptor)
	}

	identifierTrieNodes := factory.ValidatorTrieNodesTopic + core.CommunicationIdentifierBetweenShards(core.MetachainShardId, core.MetachainShardId)
	if !ficf.checkIfInterceptorExists(identifierTrieNodes) {
		interceptor, err := ficf.createOneTrieNodesInterceptor(identifierTrieNodes)
		if err != nil {
			return err
		}

		keys = append(keys, identifierTrieNodes)
		trieInterceptors = append(trieInterceptors, interceptor)
	}

	identifierTrieNodes = factory.AccountTrieNodesTopic + core.CommunicationIdentifierBetweenShards(core.MetachainShardId, core.MetachainShardId)
	if !ficf.checkIfInterceptorExists(identifierTrieNodes) {
		interceptor, err := ficf.createOneTrieNodesInterceptor(identifierTrieNodes)
		if err != nil {
			return err
		}

		keys = append(keys, identifierTrieNodes)
		trieInterceptors = append(trieInterceptors, interceptor)
	}

	return ficf.container.AddMultiple(keys, trieInterceptors)
}

func (ficf *fullSyncInterceptorsContainerFactory) createTopicAndAssignHandler(
	topic string,
	interceptor process.Interceptor,
	createChannel bool,
) (process.Interceptor, error) {

	err := ficf.messenger.CreateTopic(topic, createChannel)
	if err != nil {
		return nil, err
	}

	return interceptor, ficf.messenger.RegisterMessageProcessor(topic, interceptor)
}

func (ficf *fullSyncInterceptorsContainerFactory) generateTxInterceptors() error {
	shardC := ficf.shardCoordinator

	noOfShards := shardC.NumberOfShards()

	keys := make([]string, noOfShards)
	interceptorSlice := make([]process.Interceptor, noOfShards)

	for idx := uint32(0); idx < noOfShards; idx++ {
		identifierTx := factory.TransactionTopic + shardC.CommunicationIdentifier(idx)
		if ficf.checkIfInterceptorExists(identifierTx) {
			continue
		}

		interceptor, err := ficf.createOneTxInterceptor(identifierTx)
		if err != nil {
			return err
		}

		keys[int(idx)] = identifierTx
		interceptorSlice[int(idx)] = interceptor
	}

	//tx interceptor for metachain topic
	identifierTx := factory.TransactionTopic + shardC.CommunicationIdentifier(core.MetachainShardId)
	if !ficf.checkIfInterceptorExists(identifierTx) {
		interceptor, err := ficf.createOneTxInterceptor(identifierTx)
		if err != nil {
			return err
		}

		keys = append(keys, identifierTx)
		interceptorSlice = append(interceptorSlice, interceptor)
	}

	return ficf.container.AddMultiple(keys, interceptorSlice)
}

func (ficf *fullSyncInterceptorsContainerFactory) createOneTxInterceptor(topic string) (process.Interceptor, error) {
	txValidator, err := dataValidators.NewTxValidator(ficf.accounts, ficf.shardCoordinator, ficf.maxTxNonceDeltaAllowed)
	if err != nil {
		return nil, err
	}

	argProcessor := &processor.ArgTxInterceptorProcessor{
		ShardedDataCache: ficf.dataPool.Transactions(),
		TxValidator:      txValidator,
	}
	txProcessor, err := processor.NewTxInterceptorProcessor(argProcessor)
	if err != nil {
		return nil, err
	}

	txFactory, err := interceptorFactory.NewInterceptedTxDataFactory(ficf.argInterceptorFactory)
	if err != nil {
		return nil, err
	}

	interceptor, err := interceptors.NewMultiDataInterceptor(
		ficf.marshalizer,
		txFactory,
		txProcessor,
		ficf.globalThrottler,
	)
	if err != nil {
		return nil, err
	}

	return ficf.createTopicAndAssignHandler(topic, interceptor, true)
}

func (ficf *fullSyncInterceptorsContainerFactory) createOneUnsignedTxInterceptor(topic string) (process.Interceptor, error) {
	txValidator, err := mock.NewNilTxValidator()
	if err != nil {
		return nil, err
	}

	argProcessor := &processor.ArgTxInterceptorProcessor{
		ShardedDataCache: ficf.dataPool.UnsignedTransactions(),
		TxValidator:      txValidator,
	}
	txProcessor, err := processor.NewTxInterceptorProcessor(argProcessor)
	if err != nil {
		return nil, err
	}

	txFactory, err := interceptorFactory.NewInterceptedUnsignedTxDataFactory(ficf.argInterceptorFactory)
	if err != nil {
		return nil, err
	}

	interceptor, err := interceptors.NewMultiDataInterceptor(
		ficf.marshalizer,
		txFactory,
		txProcessor,
		ficf.globalThrottler,
	)
	if err != nil {
		return nil, err
	}

	return ficf.createTopicAndAssignHandler(topic, interceptor, true)
}

func (ficf *fullSyncInterceptorsContainerFactory) createOneRewardTxInterceptor(topic string) (process.Interceptor, error) {
	txValidator, err := mock.NewNilTxValidator()
	if err != nil {
		return nil, err
	}

	argProcessor := &processor.ArgTxInterceptorProcessor{
		ShardedDataCache: ficf.dataPool.RewardTransactions(),
		TxValidator:      txValidator,
	}
	txProcessor, err := processor.NewTxInterceptorProcessor(argProcessor)
	if err != nil {
		return nil, err
	}

	txFactory, err := interceptorFactory.NewInterceptedRewardTxDataFactory(ficf.argInterceptorFactory)
	if err != nil {
		return nil, err
	}

	interceptor, err := interceptors.NewMultiDataInterceptor(
		ficf.marshalizer,
		txFactory,
		txProcessor,
		ficf.globalThrottler,
	)
	if err != nil {
		return nil, err
	}

	return ficf.createTopicAndAssignHandler(topic, interceptor, true)
}

func (ficf *fullSyncInterceptorsContainerFactory) generateMiniBlocksInterceptors() error {
	shardC := ficf.shardCoordinator
	noOfShards := shardC.NumberOfShards()
	keys := make([]string, noOfShards+1)
	interceptorsSlice := make([]process.Interceptor, noOfShards+1)

	for idx := uint32(0); idx < noOfShards; idx++ {
		identifierMiniBlocks := factory.MiniBlocksTopic + shardC.CommunicationIdentifier(idx)
		if ficf.checkIfInterceptorExists(identifierMiniBlocks) {
			continue
		}

		interceptor, err := ficf.createOneMiniBlocksInterceptor(identifierMiniBlocks)
		if err != nil {
			return err
		}

		keys[int(idx)] = identifierMiniBlocks
		interceptorsSlice[int(idx)] = interceptor
	}

	identifierMiniBlocks := factory.MiniBlocksTopic + shardC.CommunicationIdentifier(core.MetachainShardId)
	if !ficf.checkIfInterceptorExists(identifierMiniBlocks) {
		interceptor, err := ficf.createOneMiniBlocksInterceptor(identifierMiniBlocks)
		if err != nil {
			return err
		}

		keys[noOfShards] = identifierMiniBlocks
		interceptorsSlice[noOfShards] = interceptor
	}

	return ficf.container.AddMultiple(keys, interceptorsSlice)
}

func (ficf *fullSyncInterceptorsContainerFactory) createOneMiniBlocksInterceptor(topic string) (process.Interceptor, error) {
	argProcessor := &processor.ArgTxBodyInterceptorProcessor{
		MiniblockCache:   ficf.dataPool.MiniBlocks(),
		Marshalizer:      ficf.marshalizer,
		Hasher:           ficf.hasher,
		ShardCoordinator: ficf.shardCoordinator,
	}
	txBlockBodyProcessor, err := processor.NewTxBodyInterceptorProcessor(argProcessor)
	if err != nil {
		return nil, err
	}

	txFactory, err := interceptorFactory.NewInterceptedTxBlockBodyDataFactory(ficf.argInterceptorFactory)
	if err != nil {
		return nil, err
	}

	interceptor, err := interceptors.NewSingleDataInterceptor(
		txFactory,
		txBlockBodyProcessor,
		ficf.globalThrottler,
	)
	if err != nil {
		return nil, err
	}

	return ficf.createTopicAndAssignHandler(topic, interceptor, true)
}

func (ficf *fullSyncInterceptorsContainerFactory) generateMetachainHeaderInterceptors() error {
	identifierHdr := factory.MetachainBlocksTopic
	if ficf.checkIfInterceptorExists(identifierHdr) {
		return nil
	}

	hdrValidator, err := dataValidators.NewNilHeaderValidator()
	if err != nil {
		return err
	}

	hdrFactory, err := interceptorFactory.NewInterceptedMetaHeaderDataFactory(ficf.argInterceptorFactory)
	if err != nil {
		return err
	}

	argProcessor := &processor.ArgHdrInterceptorProcessor{
		Headers:      ficf.dataPool.Headers(),
		HdrValidator: hdrValidator,
		BlackList:    ficf.blackList,
	}
	hdrProcessor, err := processor.NewHdrInterceptorProcessor(argProcessor)
	if err != nil {
		return err
	}

	//only one metachain header topic
	interceptor, err := interceptors.NewSingleDataInterceptor(
		hdrFactory,
		hdrProcessor,
		ficf.globalThrottler,
	)
	if err != nil {
		return err
	}

	_, err = ficf.createTopicAndAssignHandler(identifierHdr, interceptor, true)
	if err != nil {
		return err
	}

	return ficf.container.Add(identifierHdr, interceptor)
}

func (ficf *fullSyncInterceptorsContainerFactory) createOneTrieNodesInterceptor(topic string) (process.Interceptor, error) {
	trieNodesProcessor, err := processor.NewTrieNodesInterceptorProcessor(ficf.dataPool.TrieNodes())
	if err != nil {
		return nil, err
	}

	trieNodesFactory, err := interceptorFactory.NewInterceptedTrieNodeDataFactory(ficf.argInterceptorFactory)
	if err != nil {
		return nil, err
	}

	interceptor, err := interceptors.NewMultiDataInterceptor(
		ficf.marshalizer,
		trieNodesFactory,
		trieNodesProcessor,
		ficf.globalThrottler,
	)
	if err != nil {
		return nil, err
	}

	return ficf.createTopicAndAssignHandler(topic, interceptor, true)
}

func (ficf *fullSyncInterceptorsContainerFactory) generateRewardTxInterceptors() error {
	noOfShards := ficf.shardCoordinator.NumberOfShards()

	tmpSC, err := sharding.NewMultiShardCoordinator(noOfShards, core.MetachainShardId)
	if err != nil {
		return err
	}

	keys := make([]string, noOfShards)
	interceptorSlice := make([]process.Interceptor, noOfShards)

	for idx := uint32(0); idx < noOfShards; idx++ {
		identifierScr := factory.RewardsTransactionTopic + tmpSC.CommunicationIdentifier(idx)
		if ficf.checkIfInterceptorExists(identifierScr) {
			return nil
		}

		interceptor, err := ficf.createOneRewardTxInterceptor(identifierScr)
		if err != nil {
			return err
		}

		keys[int(idx)] = identifierScr
		interceptorSlice[int(idx)] = interceptor
	}

	return ficf.container.AddMultiple(keys, interceptorSlice)
}

func (ficf *fullSyncInterceptorsContainerFactory) setWhiteListHandlerToInterceptors() error {
	var err error

	ficf.container.Iterate(func(key string, interceptor process.Interceptor) bool {
		errFound := interceptor.SetIsDataForCurrentShardVerifier(ficf.whiteListHandler)
		if errFound != nil {
			err = errFound
			return false
		}
		return true
	})

	return err
}

// IsInterfaceNil returns true if there is no value under the interface
func (ficf *fullSyncInterceptorsContainerFactory) IsInterfaceNil() bool {
	return ficf == nil
}
