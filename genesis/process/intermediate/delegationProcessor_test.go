package intermediate

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/genesis"
	"github.com/ElrondNetwork/elrond-go/genesis/data"
	"github.com/ElrondNetwork/elrond-go/genesis/mock"
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/sharding"
	vmcommon "github.com/ElrondNetwork/elrond-vm-common"
	"github.com/stretchr/testify/assert"
)

func createMockDelegationProcessorArg() ArgDelegationProcessor {
	return ArgDelegationProcessor{
		Executor:            &mock.TxExecutionProcessorStub{},
		ShardCoordinator:    &mock.ShardCoordinatorMock{},
		AccountsParser:      &mock.AccountsParserStub{},
		SmartContractParser: &mock.SmartContractParserStub{},
		NodesListSplitter:   &mock.NodesListSplitterStub{},
		QueryService:        &mock.QueryServiceStub{},
		NodePrice:           big.NewInt(10),
	}
}

func TestNewDelegationProcessor_NilExecutorShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockDelegationProcessorArg()
	arg.Executor = nil
	dp, err := NewDelegationProcessor(arg)

	assert.True(t, check.IfNil(dp))
	assert.Equal(t, genesis.ErrNilTxExecutionProcessor, err)
}

func TestNewDelegationProcessor_NilShardCoordinatorShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockDelegationProcessorArg()
	arg.ShardCoordinator = nil
	dp, err := NewDelegationProcessor(arg)

	assert.True(t, check.IfNil(dp))
	assert.Equal(t, genesis.ErrNilShardCoordinator, err)
}

func TestNewDelegationProcessor_NilAccountsParserShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockDelegationProcessorArg()
	arg.AccountsParser = nil
	dp, err := NewDelegationProcessor(arg)

	assert.True(t, check.IfNil(dp))
	assert.Equal(t, genesis.ErrNilAccountsParser, err)
}

func TestNewDelegationProcessor_NilSmartContractParserShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockDelegationProcessorArg()
	arg.SmartContractParser = nil
	dp, err := NewDelegationProcessor(arg)

	assert.True(t, check.IfNil(dp))
	assert.Equal(t, genesis.ErrNilSmartContractParser, err)
}

func TestNewDelegationProcessor_NilNodesSplitterShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockDelegationProcessorArg()
	arg.NodesListSplitter = nil
	dp, err := NewDelegationProcessor(arg)

	assert.True(t, check.IfNil(dp))
	assert.Equal(t, genesis.ErrNilNodesListSplitter, err)
}

func TestNewDelegationProcessor_NilQueryServiceShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockDelegationProcessorArg()
	arg.QueryService = nil
	dp, err := NewDelegationProcessor(arg)

	assert.True(t, check.IfNil(dp))
	assert.Equal(t, genesis.ErrNilQueryService, err)
}

func TestNewDelegationProcessor_NilNodePriceShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockDelegationProcessorArg()
	arg.NodePrice = nil
	dp, err := NewDelegationProcessor(arg)

	assert.True(t, check.IfNil(dp))
	assert.Equal(t, genesis.ErrNilInitialNodePrice, err)
}

func TestNewDelegationProcessor_ZeroNodePriceShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockDelegationProcessorArg()
	arg.NodePrice = big.NewInt(0)
	dp, err := NewDelegationProcessor(arg)

	assert.True(t, check.IfNil(dp))
	assert.Equal(t, genesis.ErrInvalidInitialNodePrice, err)
}

func TestNewDelegationProcessor_ShouldWork(t *testing.T) {
	t.Parallel()

	arg := createMockDelegationProcessorArg()
	dp, err := NewDelegationProcessor(arg)

	assert.False(t, check.IfNil(dp))
	assert.Nil(t, err)
}

//------- ExecuteDelegation

func TestDelegationProcessor_ExecuteDelegationSplitFailsShouldErr(t *testing.T) {
	t.Parallel()

	expectedErr := fmt.Errorf("expected error")
	arg := createMockDelegationProcessorArg()
	arg.Executor = &mock.TxExecutionProcessorStub{
		ExecuteTransactionCalled: func(nonce uint64, sndAddr []byte, rcvAddress []byte, value *big.Int, data []byte) error {
			assert.Fail(t, "should have not execute a transaction")

			return nil
		},
	}
	arg.SmartContractParser = &mock.SmartContractParserStub{
		InitialSmartContractsSplitOnOwnersShardsCalled: func(shardCoordinator sharding.Coordinator) (map[uint32][]genesis.InitialSmartContractHandler, error) {
			return nil, expectedErr
		},
	}

	dp, _ := NewDelegationProcessor(arg)

	result, err := dp.ExecuteDelegation()

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, genesis.DelegationResult{}, result)
}

func TestDelegationProcessor_ExecuteDelegationNoDelegationScShouldRetNil(t *testing.T) {
	t.Parallel()

	arg := createMockDelegationProcessorArg()
	arg.Executor = &mock.TxExecutionProcessorStub{
		ExecuteTransactionCalled: func(nonce uint64, sndAddr []byte, rcvAddress []byte, value *big.Int, data []byte) error {
			assert.Fail(t, "should have not execute a transaction")

			return nil
		},
	}
	arg.SmartContractParser = &mock.SmartContractParserStub{
		InitialSmartContractsSplitOnOwnersShardsCalled: func(shardCoordinator sharding.Coordinator) (map[uint32][]genesis.InitialSmartContractHandler, error) {
			return map[uint32][]genesis.InitialSmartContractHandler{
				0: {
					&data.InitialSmartContract{
						Type: "test",
					},
				},
			}, nil
		},
	}
	dp, _ := NewDelegationProcessor(arg)

	result, err := dp.ExecuteDelegation()

	assert.Nil(t, err)
	assert.Equal(t, genesis.DelegationResult{}, result)
}

func TestDelegationProcessor_ExecuteDelegationStakeShouldWork(t *testing.T) {
	t.Parallel()

	staker1 := []byte("stakerB")
	staker2 := []byte("stakerC")
	delegationSc := []byte("delegation SC")
	pubkey1 := []byte("pubkey1")
	pubkey2 := []byte("pubkey2")
	pubkey3 := []byte("pubkey3")

	arg := createMockDelegationProcessorArg()
	arg.Executor = &mock.TxExecutionProcessorStub{
		ExecuteTransactionCalled: func(nonce uint64, sndAddr []byte, rcvAddress []byte, value *big.Int, data []byte) error {
			isStakeCall := strings.Contains(string(data), "stake")
			isStaker := bytes.Equal(sndAddr, staker1) || bytes.Equal(sndAddr, staker2)
			if isStakeCall && !isStaker {
				assert.Fail(t, "stake should have been called by the one of the stakers")
			}

			return nil
		},
	}
	arg.ShardCoordinator = &mock.ShardCoordinatorMock{
		SelfShardId: 0,
		NumOfShards: 2,
	}
	arg.AccountsParser = &mock.AccountsParserStub{
		GetInitialAccountsForDelegatedCalled: func(addressBytes []byte) []genesis.InitialAccountHandler {
			if bytes.Equal(addressBytes, delegationSc) {
				ia1 := &data.InitialAccount{
					Delegation: &data.DelegationData{
						Value: big.NewInt(2),
					},
				}
				ia1.SetAddressBytes(staker1)
				ia1.Delegation.SetAddressBytes(delegationSc)

				ia2 := &data.InitialAccount{
					Delegation: &data.DelegationData{
						Value: big.NewInt(2),
					},
				}
				ia2.SetAddressBytes(staker2)
				ia2.Delegation.SetAddressBytes(delegationSc)

				return []genesis.InitialAccountHandler{ia1, ia2}
			}

			return make([]genesis.InitialAccountHandler, 0)
		},
	}
	arg.SmartContractParser = &mock.SmartContractParserStub{
		InitialSmartContractsSplitOnOwnersShardsCalled: func(shardCoordinator sharding.Coordinator) (map[uint32][]genesis.InitialSmartContractHandler, error) {
			sc := &data.InitialSmartContract{
				Type: genesis.DelegationType,
			}
			sc.SetAddressBytes(delegationSc)

			return map[uint32][]genesis.InitialSmartContractHandler{
				0: {sc},
			}, nil
		},
	}
	arg.QueryService = &mock.QueryServiceStub{
		ExecuteQueryCalled: func(query *process.SCQuery) (*vmcommon.VMOutput, error) {
			if query.FuncName == "getFilledStake" {
				return &vmcommon.VMOutput{
					ReturnData: [][]byte{big.NewInt(4).Bytes()},
				}, nil
			}
			if query.FuncName == "getBlsKeys" {
				return &vmcommon.VMOutput{
					ReturnData: [][]byte{pubkey2, pubkey3, pubkey1}, //random order should work
				}, nil
			}

			return nil, fmt.Errorf("unexpected function")
		},
	}
	arg.NodesListSplitter = &mock.NodesListSplitterStub{
		GetDelegatedNodesCalled: func(delegationScAddress []byte) []sharding.GenesisNodeInfoHandler {
			return []sharding.GenesisNodeInfoHandler{
				&mock.GenesisNodeInfoHandlerMock{
					AddressBytesValue: delegationSc,
					PubKeyBytesValue:  pubkey1,
				},
				&mock.GenesisNodeInfoHandlerMock{
					AddressBytesValue: delegationSc,
					PubKeyBytesValue:  pubkey2,
				},
				&mock.GenesisNodeInfoHandlerMock{
					AddressBytesValue: delegationSc,
					PubKeyBytesValue:  pubkey3,
				},
			}
		},
	}
	dp, _ := NewDelegationProcessor(arg)

	result, err := dp.ExecuteDelegation()

	expectedResult := genesis.DelegationResult{
		NumTotalDelegated: 3,
		NumTotalStaked:    2,
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

//------- SameElements

func TestSameElements_WrongNumberShouldErr(t *testing.T) {
	t.Parallel()

	scReturned := [][]byte{[]byte("buf1"), []byte("buf2"), []byte("buf3")}
	loaded := [][]byte{[]byte("buf1"), []byte("buf2")}

	dp := &delegationProcessor{}
	err := dp.sameElements(scReturned, loaded)

	assert.True(t, errors.Is(err, genesis.ErrWhileVerifyingDelegation))
}

func TestSameElements_MissingFromLoadedShouldErr(t *testing.T) {
	t.Parallel()

	scReturned := [][]byte{[]byte("buf5"), []byte("buf2"), []byte("buf3")}
	loaded := [][]byte{[]byte("buf1"), []byte("buf3"), []byte("buf2")}

	dp := &delegationProcessor{}
	err := dp.sameElements(scReturned, loaded)

	assert.True(t, errors.Is(err, genesis.ErrMissingElement))
}

func TestSameElements_DuplicateShouldErr(t *testing.T) {
	t.Parallel()

	scReturned := [][]byte{[]byte("buf2"), []byte("buf2"), []byte("buf3")}
	loaded := [][]byte{[]byte("buf2"), []byte("buf1"), []byte("buf1")}

	dp := &delegationProcessor{}
	err := dp.sameElements(scReturned, loaded)

	assert.True(t, errors.Is(err, genesis.ErrMissingElement))
}

func TestSameElements_ShouldWork(t *testing.T) {
	t.Parallel()

	scReturned := [][]byte{[]byte("buf1"), []byte("buf2"), []byte("buf3")}
	loaded := [][]byte{[]byte("buf2"), []byte("buf3"), []byte("buf1")}

	dp := &delegationProcessor{}
	err := dp.sameElements(scReturned, loaded)

	assert.Nil(t, err)
}
