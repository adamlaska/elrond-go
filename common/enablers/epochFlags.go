package enablers

import (
	"github.com/multiversx/mx-chain-core-go/core/atomic"
)

type epochFlagsHolder struct {
	builtInFunctionOnMetaFlag                  *atomic.Flag
	computeRewardCheckpointFlag                *atomic.Flag
	scrSizeInvariantCheckFlag                  *atomic.Flag
	backwardCompSaveKeyValueFlag               *atomic.Flag
	esdtNFTCreateOnMultiShardFlag              *atomic.Flag
	metaESDTSetFlag                            *atomic.Flag
	addTokensToDelegationFlag                  *atomic.Flag
	multiESDTTransferFixOnCallBackFlag         *atomic.Flag
	optimizeGasUsedInCrossMiniBlocksFlag       *atomic.Flag
	correctFirstQueuedFlag                     *atomic.Flag
	deleteDelegatorAfterClaimRewardsFlag       *atomic.Flag
	fixOOGReturnCodeFlag                       *atomic.Flag
	removeNonUpdatedStorageFlag                *atomic.Flag
	optimizeNFTStoreFlag                       *atomic.Flag
	createNFTThroughExecByCallerFlag           *atomic.Flag
	stopDecreasingValidatorRatingWhenStuckFlag *atomic.Flag
	frontRunningProtectionFlag                 *atomic.Flag
	isPayableBySCFlag                          *atomic.Flag
	cleanUpInformativeSCRsFlag                 *atomic.Flag
	storageAPICostOptimizationFlag             *atomic.Flag
	esdtRegisterAndSetAllRolesFlag             *atomic.Flag
	scheduledMiniBlocksFlag                    *atomic.Flag
	correctJailedNotUnStakedEmptyQueueFlag     *atomic.Flag
	doNotReturnOldBlockInBlockchainHookFlag    *atomic.Flag
	addFailedRelayedTxToInvalidMBsFlag         *atomic.Flag
	scrSizeInvariantOnBuiltInResultFlag        *atomic.Flag
	checkCorrectTokenIDForTransferRoleFlag     *atomic.Flag
	failExecutionOnEveryAPIErrorFlag           *atomic.Flag
	isMiniBlockPartialExecutionFlag            *atomic.Flag
	managedCryptoAPIsFlag                      *atomic.Flag
	esdtMetadataContinuousCleanupFlag          *atomic.Flag
	disableExecByCallerFlag                    *atomic.Flag
	refactorContextFlag                        *atomic.Flag
	checkFunctionArgumentFlag                  *atomic.Flag
	checkExecuteOnReadOnlyFlag                 *atomic.Flag
	setSenderInEeiOutputTransferFlag           *atomic.Flag
	changeDelegationOwnerFlag                  *atomic.Flag
	refactorPeersMiniBlocksFlag                *atomic.Flag
	scProcessorV2Flag                          *atomic.Flag
	fixAsyncCallBackArgsList                   *atomic.Flag
	fixOldTokenLiquidity                       *atomic.Flag
	runtimeMemStoreLimitFlag                   *atomic.Flag
	runtimeCodeSizeFixFlag                     *atomic.Flag
	maxBlockchainHookCountersFlag              *atomic.Flag
	wipeSingleNFTLiquidityDecreaseFlag         *atomic.Flag
	alwaysSaveTokenMetaDataFlag                *atomic.Flag
	setGuardianFlag                            *atomic.Flag
	relayedNonceFixFlag                        *atomic.Flag
	keepExecOrderOnCreatedSCRsFlag             *atomic.Flag
	multiClaimOnDelegationFlag                 *atomic.Flag
	changeUsernameFlag                         *atomic.Flag
	consistentTokensValuesCheckFlag            *atomic.Flag
	autoBalanceDataTriesFlag                   *atomic.Flag
	fixDelegationChangeOwnerOnAccountFlag      *atomic.Flag
}

func newEpochFlagsHolder() *epochFlagsHolder {
	return &epochFlagsHolder{
		builtInFunctionOnMetaFlag:                  &atomic.Flag{},
		computeRewardCheckpointFlag:                &atomic.Flag{},
		scrSizeInvariantCheckFlag:                  &atomic.Flag{},
		backwardCompSaveKeyValueFlag:               &atomic.Flag{},
		esdtNFTCreateOnMultiShardFlag:              &atomic.Flag{},
		metaESDTSetFlag:                            &atomic.Flag{},
		addTokensToDelegationFlag:                  &atomic.Flag{},
		multiESDTTransferFixOnCallBackFlag:         &atomic.Flag{},
		optimizeGasUsedInCrossMiniBlocksFlag:       &atomic.Flag{},
		correctFirstQueuedFlag:                     &atomic.Flag{},
		deleteDelegatorAfterClaimRewardsFlag:       &atomic.Flag{},
		fixOOGReturnCodeFlag:                       &atomic.Flag{},
		removeNonUpdatedStorageFlag:                &atomic.Flag{},
		optimizeNFTStoreFlag:                       &atomic.Flag{},
		createNFTThroughExecByCallerFlag:           &atomic.Flag{},
		stopDecreasingValidatorRatingWhenStuckFlag: &atomic.Flag{},
		frontRunningProtectionFlag:                 &atomic.Flag{},
		isPayableBySCFlag:                          &atomic.Flag{},
		cleanUpInformativeSCRsFlag:                 &atomic.Flag{},
		storageAPICostOptimizationFlag:             &atomic.Flag{},
		esdtRegisterAndSetAllRolesFlag:             &atomic.Flag{},
		scheduledMiniBlocksFlag:                    &atomic.Flag{},
		correctJailedNotUnStakedEmptyQueueFlag:     &atomic.Flag{},
		doNotReturnOldBlockInBlockchainHookFlag:    &atomic.Flag{},
		addFailedRelayedTxToInvalidMBsFlag:         &atomic.Flag{},
		scrSizeInvariantOnBuiltInResultFlag:        &atomic.Flag{},
		checkCorrectTokenIDForTransferRoleFlag:     &atomic.Flag{},
		failExecutionOnEveryAPIErrorFlag:           &atomic.Flag{},
		isMiniBlockPartialExecutionFlag:            &atomic.Flag{},
		managedCryptoAPIsFlag:                      &atomic.Flag{},
		esdtMetadataContinuousCleanupFlag:          &atomic.Flag{},
		disableExecByCallerFlag:                    &atomic.Flag{},
		refactorContextFlag:                        &atomic.Flag{},
		checkFunctionArgumentFlag:                  &atomic.Flag{},
		checkExecuteOnReadOnlyFlag:                 &atomic.Flag{},
		setSenderInEeiOutputTransferFlag:           &atomic.Flag{},
		changeDelegationOwnerFlag:                  &atomic.Flag{},
		refactorPeersMiniBlocksFlag:                &atomic.Flag{},
		scProcessorV2Flag:                          &atomic.Flag{},
		fixAsyncCallBackArgsList:                   &atomic.Flag{},
		fixOldTokenLiquidity:                       &atomic.Flag{},
		runtimeMemStoreLimitFlag:                   &atomic.Flag{},
		runtimeCodeSizeFixFlag:                     &atomic.Flag{},
		maxBlockchainHookCountersFlag:              &atomic.Flag{},
		wipeSingleNFTLiquidityDecreaseFlag:         &atomic.Flag{},
		alwaysSaveTokenMetaDataFlag:                &atomic.Flag{},
		setGuardianFlag:                            &atomic.Flag{},
		relayedNonceFixFlag:                        &atomic.Flag{},
		keepExecOrderOnCreatedSCRsFlag:             &atomic.Flag{},
		consistentTokensValuesCheckFlag:            &atomic.Flag{},
		multiClaimOnDelegationFlag:                 &atomic.Flag{},
		changeUsernameFlag:                         &atomic.Flag{},
		autoBalanceDataTriesFlag:                   &atomic.Flag{},
		fixDelegationChangeOwnerOnAccountFlag:      &atomic.Flag{},
	}
}

// IsBuiltInFunctionOnMetaFlagEnabled returns true if builtInFunctionOnMetaFlag is enabled
func (holder *epochFlagsHolder) IsBuiltInFunctionOnMetaFlagEnabled() bool {
	return holder.builtInFunctionOnMetaFlag.IsSet()
}

// IsComputeRewardCheckpointFlagEnabled returns true if computeRewardCheckpointFlag is enabled
func (holder *epochFlagsHolder) IsComputeRewardCheckpointFlagEnabled() bool {
	return holder.computeRewardCheckpointFlag.IsSet()
}

// IsSCRSizeInvariantCheckFlagEnabled returns true if scrSizeInvariantCheckFlag is enabled
func (holder *epochFlagsHolder) IsSCRSizeInvariantCheckFlagEnabled() bool {
	return holder.scrSizeInvariantCheckFlag.IsSet()
}

// IsBackwardCompSaveKeyValueFlagEnabled returns true if backwardCompSaveKeyValueFlag is enabled
func (holder *epochFlagsHolder) IsBackwardCompSaveKeyValueFlagEnabled() bool {
	return holder.backwardCompSaveKeyValueFlag.IsSet()
}

// IsESDTNFTCreateOnMultiShardFlagEnabled returns true if esdtNFTCreateOnMultiShardFlag is enabled
func (holder *epochFlagsHolder) IsESDTNFTCreateOnMultiShardFlagEnabled() bool {
	return holder.esdtNFTCreateOnMultiShardFlag.IsSet()
}

// IsMetaESDTSetFlagEnabled returns true if metaESDTSetFlag is enabled
func (holder *epochFlagsHolder) IsMetaESDTSetFlagEnabled() bool {
	return holder.metaESDTSetFlag.IsSet()
}

// IsAddTokensToDelegationFlagEnabled returns true if addTokensToDelegationFlag is enabled
func (holder *epochFlagsHolder) IsAddTokensToDelegationFlagEnabled() bool {
	return holder.addTokensToDelegationFlag.IsSet()
}

// IsMultiESDTTransferFixOnCallBackFlagEnabled returns true if multiESDTTransferFixOnCallBackFlag is enabled
func (holder *epochFlagsHolder) IsMultiESDTTransferFixOnCallBackFlagEnabled() bool {
	return holder.multiESDTTransferFixOnCallBackFlag.IsSet()
}

// IsOptimizeGasUsedInCrossMiniBlocksFlagEnabled returns true if optimizeGasUsedInCrossMiniBlocksFlag is enabled
func (holder *epochFlagsHolder) IsOptimizeGasUsedInCrossMiniBlocksFlagEnabled() bool {
	return holder.optimizeGasUsedInCrossMiniBlocksFlag.IsSet()
}

// IsCorrectFirstQueuedFlagEnabled returns true if correctFirstQueuedFlag is enabled
func (holder *epochFlagsHolder) IsCorrectFirstQueuedFlagEnabled() bool {
	return holder.correctFirstQueuedFlag.IsSet()
}

// IsDeleteDelegatorAfterClaimRewardsFlagEnabled returns true if deleteDelegatorAfterClaimRewardsFlag is enabled
func (holder *epochFlagsHolder) IsDeleteDelegatorAfterClaimRewardsFlagEnabled() bool {
	return holder.deleteDelegatorAfterClaimRewardsFlag.IsSet()
}

// IsFixOOGReturnCodeFlagEnabled returns true if fixOOGReturnCodeFlag is enabled
func (holder *epochFlagsHolder) IsFixOOGReturnCodeFlagEnabled() bool {
	return holder.fixOOGReturnCodeFlag.IsSet()
}

// IsRemoveNonUpdatedStorageFlagEnabled returns true if removeNonUpdatedStorageFlag is enabled
func (holder *epochFlagsHolder) IsRemoveNonUpdatedStorageFlagEnabled() bool {
	return holder.removeNonUpdatedStorageFlag.IsSet()
}

// IsOptimizeNFTStoreFlagEnabled returns true if removeNonUpdatedStorageFlag is enabled
func (holder *epochFlagsHolder) IsOptimizeNFTStoreFlagEnabled() bool {
	return holder.optimizeNFTStoreFlag.IsSet()
}

// IsCreateNFTThroughExecByCallerFlagEnabled returns true if createNFTThroughExecByCallerFlag is enabled
func (holder *epochFlagsHolder) IsCreateNFTThroughExecByCallerFlagEnabled() bool {
	return holder.createNFTThroughExecByCallerFlag.IsSet()
}

// IsStopDecreasingValidatorRatingWhenStuckFlagEnabled returns true if stopDecreasingValidatorRatingWhenStuckFlag is enabled
func (holder *epochFlagsHolder) IsStopDecreasingValidatorRatingWhenStuckFlagEnabled() bool {
	return holder.stopDecreasingValidatorRatingWhenStuckFlag.IsSet()
}

// IsFrontRunningProtectionFlagEnabled returns true if frontRunningProtectionFlag is enabled
func (holder *epochFlagsHolder) IsFrontRunningProtectionFlagEnabled() bool {
	return holder.frontRunningProtectionFlag.IsSet()
}

// IsPayableBySCFlagEnabled returns true if isPayableBySCFlag is enabled
func (holder *epochFlagsHolder) IsPayableBySCFlagEnabled() bool {
	return holder.isPayableBySCFlag.IsSet()
}

// IsCleanUpInformativeSCRsFlagEnabled returns true if cleanUpInformativeSCRsFlag is enabled
func (holder *epochFlagsHolder) IsCleanUpInformativeSCRsFlagEnabled() bool {
	return holder.cleanUpInformativeSCRsFlag.IsSet()
}

// IsStorageAPICostOptimizationFlagEnabled returns true if storageAPICostOptimizationFlag is enabled
func (holder *epochFlagsHolder) IsStorageAPICostOptimizationFlagEnabled() bool {
	return holder.storageAPICostOptimizationFlag.IsSet()
}

// IsESDTRegisterAndSetAllRolesFlagEnabled returns true if esdtRegisterAndSetAllRolesFlag is enabled
func (holder *epochFlagsHolder) IsESDTRegisterAndSetAllRolesFlagEnabled() bool {
	return holder.esdtRegisterAndSetAllRolesFlag.IsSet()
}

// IsScheduledMiniBlocksFlagEnabled returns true if scheduledMiniBlocksFlag is enabled
func (holder *epochFlagsHolder) IsScheduledMiniBlocksFlagEnabled() bool {
	return holder.scheduledMiniBlocksFlag.IsSet()
}

// IsCorrectJailedNotUnStakedEmptyQueueFlagEnabled returns true if correctJailedNotUnStakedEmptyQueueFlag is enabled
func (holder *epochFlagsHolder) IsCorrectJailedNotUnStakedEmptyQueueFlagEnabled() bool {
	return holder.correctJailedNotUnStakedEmptyQueueFlag.IsSet()
}

// IsDoNotReturnOldBlockInBlockchainHookFlagEnabled returns true if doNotReturnOldBlockInBlockchainHookFlag is enabled
func (holder *epochFlagsHolder) IsDoNotReturnOldBlockInBlockchainHookFlagEnabled() bool {
	return holder.doNotReturnOldBlockInBlockchainHookFlag.IsSet()
}

// IsAddFailedRelayedTxToInvalidMBsFlag returns true if addFailedRelayedTxToInvalidMBsFlag is enabled
func (holder *epochFlagsHolder) IsAddFailedRelayedTxToInvalidMBsFlag() bool {
	return holder.addFailedRelayedTxToInvalidMBsFlag.IsSet()
}

// IsSCRSizeInvariantOnBuiltInResultFlagEnabled returns true if scrSizeInvariantOnBuiltInResultFlag is enabled
func (holder *epochFlagsHolder) IsSCRSizeInvariantOnBuiltInResultFlagEnabled() bool {
	return holder.scrSizeInvariantOnBuiltInResultFlag.IsSet()
}

// IsCheckCorrectTokenIDForTransferRoleFlagEnabled returns true if checkCorrectTokenIDForTransferRoleFlag is enabled
func (holder *epochFlagsHolder) IsCheckCorrectTokenIDForTransferRoleFlagEnabled() bool {
	return holder.checkCorrectTokenIDForTransferRoleFlag.IsSet()
}

// IsFailExecutionOnEveryAPIErrorFlagEnabled returns true if failExecutionOnEveryAPIErrorFlag is enabled
func (holder *epochFlagsHolder) IsFailExecutionOnEveryAPIErrorFlagEnabled() bool {
	return holder.failExecutionOnEveryAPIErrorFlag.IsSet()
}

// IsMiniBlockPartialExecutionFlagEnabled returns true if isMiniBlockPartialExecutionFlag is enabled
func (holder *epochFlagsHolder) IsMiniBlockPartialExecutionFlagEnabled() bool {
	return holder.isMiniBlockPartialExecutionFlag.IsSet()
}

// IsManagedCryptoAPIsFlagEnabled returns true if managedCryptoAPIsFlag is enabled
func (holder *epochFlagsHolder) IsManagedCryptoAPIsFlagEnabled() bool {
	return holder.managedCryptoAPIsFlag.IsSet()
}

// IsESDTMetadataContinuousCleanupFlagEnabled returns true if esdtMetadataContinuousCleanupFlag is enabled
func (holder *epochFlagsHolder) IsESDTMetadataContinuousCleanupFlagEnabled() bool {
	return holder.esdtMetadataContinuousCleanupFlag.IsSet()
}

// IsDisableExecByCallerFlagEnabled returns true if disableExecByCallerFlag is enabled
func (holder *epochFlagsHolder) IsDisableExecByCallerFlagEnabled() bool {
	return holder.disableExecByCallerFlag.IsSet()
}

// IsRefactorContextFlagEnabled returns true if refactorContextFlag is enabled
func (holder *epochFlagsHolder) IsRefactorContextFlagEnabled() bool {
	return holder.refactorContextFlag.IsSet()
}

// IsCheckFunctionArgumentFlagEnabled returns true if checkFunctionArgumentFlag is enabled
func (holder *epochFlagsHolder) IsCheckFunctionArgumentFlagEnabled() bool {
	return holder.checkFunctionArgumentFlag.IsSet()
}

// IsCheckExecuteOnReadOnlyFlagEnabled returns true if checkExecuteOnReadOnlyFlag is enabled
func (holder *epochFlagsHolder) IsCheckExecuteOnReadOnlyFlagEnabled() bool {
	return holder.checkExecuteOnReadOnlyFlag.IsSet()
}

// IsSetSenderInEeiOutputTransferFlagEnabled returns true if setSenderInEeiOutputTransferFlag is enabled
func (holder *epochFlagsHolder) IsSetSenderInEeiOutputTransferFlagEnabled() bool {
	return holder.setSenderInEeiOutputTransferFlag.IsSet()
}

// IsFixAsyncCallbackCheckFlagEnabled returns true if esdtMetadataContinuousCleanupFlag is enabled
// this is a duplicate for ESDTMetadataContinuousCleanupEnableEpoch needed for consistency into vm-common
func (holder *epochFlagsHolder) IsFixAsyncCallbackCheckFlagEnabled() bool {
	return holder.esdtMetadataContinuousCleanupFlag.IsSet()
}

// IsSaveToSystemAccountFlagEnabled returns true if optimizeNFTStoreFlag is enabled
// this is a duplicate for OptimizeNFTStoreEnableEpoch needed for consistency into vm-common
func (holder *epochFlagsHolder) IsSaveToSystemAccountFlagEnabled() bool {
	return holder.optimizeNFTStoreFlag.IsSet()
}

// IsCheckFrozenCollectionFlagEnabled returns true if optimizeNFTStoreFlag is enabled
// this is a duplicate for OptimizeNFTStoreEnableEpoch needed for consistency into vm-common
func (holder *epochFlagsHolder) IsCheckFrozenCollectionFlagEnabled() bool {
	return holder.optimizeNFTStoreFlag.IsSet()
}

// IsSendAlwaysFlagEnabled returns true if esdtMetadataContinuousCleanupFlag is enabled
// this is a duplicate for ESDTMetadataContinuousCleanupEnableEpoch needed for consistency into vm-common
func (holder *epochFlagsHolder) IsSendAlwaysFlagEnabled() bool {
	return holder.esdtMetadataContinuousCleanupFlag.IsSet()
}

// IsValueLengthCheckFlagEnabled returns true if optimizeNFTStoreFlag is enabled
// this is a duplicate for OptimizeNFTStoreEnableEpoch needed for consistency into vm-common
func (holder *epochFlagsHolder) IsValueLengthCheckFlagEnabled() bool {
	return holder.optimizeNFTStoreFlag.IsSet()
}

// IsCheckTransferFlagEnabled returns true if optimizeNFTStoreFlag is enabled
// this is a duplicate for OptimizeNFTStoreEnableEpoch needed for consistency into vm-common
func (holder *epochFlagsHolder) IsCheckTransferFlagEnabled() bool {
	return holder.optimizeNFTStoreFlag.IsSet()
}

// IsTransferToMetaFlagEnabled returns true if builtInFunctionOnMetaFlag is enabled
// this is a duplicate for BuiltInFunctionOnMetaEnableEpoch needed for consistency into vm-common
func (holder *epochFlagsHolder) IsTransferToMetaFlagEnabled() bool {
	return holder.builtInFunctionOnMetaFlag.IsSet()
}

// IsChangeDelegationOwnerFlagEnabled returns true if the change delegation owner feature is enabled
func (holder *epochFlagsHolder) IsChangeDelegationOwnerFlagEnabled() bool {
	return holder.changeDelegationOwnerFlag.IsSet()
}

// IsRefactorPeersMiniBlocksFlagEnabled returns true if refactorPeersMiniBlocksFlag is enabled
func (holder *epochFlagsHolder) IsRefactorPeersMiniBlocksFlagEnabled() bool {
	return holder.refactorPeersMiniBlocksFlag.IsSet()
}

// IsSCProcessorV2FlagEnabled returns true if scProcessorV2Flag is enabled
func (holder *epochFlagsHolder) IsSCProcessorV2FlagEnabled() bool {
	return holder.scProcessorV2Flag.IsSet()
}

// IsFixAsyncCallBackArgsListFlagEnabled returns true if fixAsyncCallBackArgsList is enabled
func (holder *epochFlagsHolder) IsFixAsyncCallBackArgsListFlagEnabled() bool {
	return holder.fixAsyncCallBackArgsList.IsSet()
}

// IsFixOldTokenLiquidityEnabled returns true if fixOldTokenLiquidity is enabled
func (holder *epochFlagsHolder) IsFixOldTokenLiquidityEnabled() bool {
	return holder.fixOldTokenLiquidity.IsSet()
}

// IsRuntimeMemStoreLimitEnabled returns true if runtimeMemStoreLimitFlag is enabled
func (holder *epochFlagsHolder) IsRuntimeMemStoreLimitEnabled() bool {
	return holder.runtimeMemStoreLimitFlag.IsSet()
}

// IsRuntimeCodeSizeFixEnabled returns true if runtimeCodeSizeFixFlag is enabled
func (holder *epochFlagsHolder) IsRuntimeCodeSizeFixEnabled() bool {
	return holder.runtimeCodeSizeFixFlag.IsSet()
}

// IsMaxBlockchainHookCountersFlagEnabled returns true if maxBlockchainHookCountersFlagEnabled is enabled
func (holder *epochFlagsHolder) IsMaxBlockchainHookCountersFlagEnabled() bool {
	return holder.maxBlockchainHookCountersFlag.IsSet()
}

// IsWipeSingleNFTLiquidityDecreaseEnabled returns true if wipeSingleNFTLiquidityDecreaseFlag is enabled
func (holder *epochFlagsHolder) IsWipeSingleNFTLiquidityDecreaseEnabled() bool {
	return holder.wipeSingleNFTLiquidityDecreaseFlag.IsSet()
}

// IsAlwaysSaveTokenMetaDataEnabled returns true if alwaysSaveTokenMetaDataFlag is enabled
func (holder *epochFlagsHolder) IsAlwaysSaveTokenMetaDataEnabled() bool {
	return holder.alwaysSaveTokenMetaDataFlag.IsSet()
}

// IsSetGuardianEnabled returns true if setGuardianFlag is enabled
func (holder *epochFlagsHolder) IsSetGuardianEnabled() bool {
	return holder.setGuardianFlag.IsSet()
}

// IsRelayedNonceFixEnabled returns true if relayedNonceFixFlag is enabled
func (holder *epochFlagsHolder) IsRelayedNonceFixEnabled() bool {
	return holder.relayedNonceFixFlag.IsSet()
}

// IsConsistentTokensValuesLengthCheckEnabled returns true if consistentTokensValuesCheckFlag is enabled
func (holder *epochFlagsHolder) IsConsistentTokensValuesLengthCheckEnabled() bool {
	return holder.consistentTokensValuesCheckFlag.IsSet()
}

// IsKeepExecOrderOnCreatedSCRsEnabled returns true if keepExecOrderOnCreatedSCRsFlag is enabled
func (holder *epochFlagsHolder) IsKeepExecOrderOnCreatedSCRsEnabled() bool {
	return holder.keepExecOrderOnCreatedSCRsFlag.IsSet()
}

// IsMultiClaimOnDelegationEnabled returns true if multi claim on delegation is enabled
func (holder *epochFlagsHolder) IsMultiClaimOnDelegationEnabled() bool {
	return holder.multiClaimOnDelegationFlag.IsSet()
}

// IsChangeUsernameEnabled returns true if changeUsernameFlag is enabled
func (holder *epochFlagsHolder) IsChangeUsernameEnabled() bool {
	return holder.changeUsernameFlag.IsSet()
}

// IsAutoBalanceDataTriesEnabled returns true if autoBalanceDataTriesFlag is enabled
func (holder *epochFlagsHolder) IsAutoBalanceDataTriesEnabled() bool {
	return holder.autoBalanceDataTriesFlag.IsSet()
}

// FixDelegationChangeOwnerOnAccountEnabled returns true if the fix for the delegation change owner on account is enabled
func (holder *epochFlagsHolder) FixDelegationChangeOwnerOnAccountEnabled() bool {
	return holder.fixDelegationChangeOwnerOnAccountFlag.IsSet()
}
