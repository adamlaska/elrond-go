package staking

import (
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go-core/data/typeConverters/uint64ByteSlice"
	"github.com/ElrondNetwork/elrond-go-core/hashing/sha256"
	"github.com/ElrondNetwork/elrond-go/common/forking"
	"github.com/ElrondNetwork/elrond-go/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/dataRetriever/blockchain"
	"github.com/ElrondNetwork/elrond-go/epochStart/notifier"
	factory2 "github.com/ElrondNetwork/elrond-go/factory"
	"github.com/ElrondNetwork/elrond-go/integrationTests"
	mock2 "github.com/ElrondNetwork/elrond-go/integrationTests/mock"
	factory3 "github.com/ElrondNetwork/elrond-go/node/mock/factory"
	"github.com/ElrondNetwork/elrond-go/process/mock"
	"github.com/ElrondNetwork/elrond-go/sharding"
	"github.com/ElrondNetwork/elrond-go/state/factory"
	"github.com/ElrondNetwork/elrond-go/statusHandler"
	"github.com/ElrondNetwork/elrond-go/testscommon"
	dataRetrieverMock "github.com/ElrondNetwork/elrond-go/testscommon/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/testscommon/mainFactoryMocks"
	"github.com/ElrondNetwork/elrond-go/trie"
)

func createComponentHolders(numOfShards uint32) (
	factory2.CoreComponentsHolder,
	factory2.DataComponentsHolder,
	factory2.BootstrapComponentsHolder,
	factory2.StatusComponentsHolder,
	factory2.StateComponentsHandler,
) {
	coreComponents := createCoreComponents()
	statusComponents := createStatusComponents()
	dataComponents := createDataComponents(coreComponents)
	stateComponents := createStateComponents(coreComponents)
	boostrapComponents := createBootstrapComponents(numOfShards)

	return coreComponents, dataComponents, boostrapComponents, statusComponents, stateComponents
}

func createCoreComponents() factory2.CoreComponentsHolder {
	return &mock2.CoreComponentsStub{
		InternalMarshalizerField:           &testscommon.MarshalizerMock{},
		HasherField:                        sha256.NewSha256(),
		Uint64ByteSliceConverterField:      uint64ByteSlice.NewBigEndianConverter(),
		StatusHandlerField:                 statusHandler.NewStatusMetrics(),
		RoundHandlerField:                  &mock.RoundHandlerMock{RoundTimeDuration: time.Second},
		EpochStartNotifierWithConfirmField: notifier.NewEpochStartSubscriptionHandler(),
		EpochNotifierField:                 forking.NewGenericEpochNotifier(),
		RaterField:                         &testscommon.RaterMock{Chance: 5},
		AddressPubKeyConverterField:        &testscommon.PubkeyConverterMock{},
		EconomicsDataField:                 createEconomicsData(),
	}
}

func createDataComponents(coreComponents factory2.CoreComponentsHolder) factory2.DataComponentsHolder {
	blockChain, _ := blockchain.NewMetaChain(coreComponents.StatusHandler())
	genesisBlock := createGenesisMetaBlock()
	genesisBlockHash, _ := coreComponents.InternalMarshalizer().Marshal(genesisBlock)
	genesisBlockHash = coreComponents.Hasher().Compute(string(genesisBlockHash))
	_ = blockChain.SetGenesisHeader(createGenesisMetaBlock())
	blockChain.SetGenesisHeaderHash(genesisBlockHash)

	chainStorer := dataRetriever.NewChainStorer()
	chainStorer.AddStorer(dataRetriever.BootstrapUnit, integrationTests.CreateMemUnit())
	chainStorer.AddStorer(dataRetriever.MetaHdrNonceHashDataUnit, integrationTests.CreateMemUnit())
	chainStorer.AddStorer(dataRetriever.MetaBlockUnit, integrationTests.CreateMemUnit())
	return &factory3.DataComponentsMock{
		Store:         chainStorer,
		DataPool:      dataRetrieverMock.NewPoolsHolderMock(),
		BlockChain:    blockChain,
		EconomicsData: createEconomicsData(),
	}
}

func createBootstrapComponents(numOfShards uint32) factory2.BootstrapComponentsHolder {
	shardCoordinator, _ := sharding.NewMultiShardCoordinator(numOfShards, core.MetachainShardId)

	return &mainFactoryMocks.BootstrapComponentsStub{
		ShCoordinator:        shardCoordinator,
		HdrIntegrityVerifier: &mock.HeaderIntegrityVerifierStub{},
		VersionedHdrFactory: &testscommon.VersionedHeaderFactoryStub{
			CreateCalled: func(epoch uint32) data.HeaderHandler {
				return &block.MetaBlock{Epoch: epoch}
			},
		},
	}
}

func createStateComponents(coreComponents factory2.CoreComponentsHolder) factory2.StateComponentsHandler {
	trieFactoryManager, _ := trie.NewTrieStorageManagerWithoutPruning(integrationTests.CreateMemUnit())
	userAccountsDB := createAccountsDB(coreComponents.Hasher(), coreComponents.InternalMarshalizer(), factory.NewAccountCreator(), trieFactoryManager)
	peerAccountsDB := createAccountsDB(coreComponents.Hasher(), coreComponents.InternalMarshalizer(), factory.NewPeerAccountCreator(), trieFactoryManager)
	return &testscommon.StateComponentsMock{
		PeersAcc: peerAccountsDB,
		Accounts: userAccountsDB,
	}
}

func createStatusComponents() factory2.StatusComponentsHolder {
	return &mock2.StatusComponentsStub{
		Outport: &testscommon.OutportStub{},
	}
}
