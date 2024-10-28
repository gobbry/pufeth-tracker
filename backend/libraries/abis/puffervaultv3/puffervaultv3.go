// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PufferVaultV3

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IEigenLayerQueuedWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IEigenLayerQueuedWithdrawal struct {
	Strategies           []common.Address
	Shares               []*big.Int
	Depositor            common.Address
	WithdrawerAndNonce   IEigenLayerWithdrawerAndNonce
	WithdrawalStartBlock uint32
	DelegatedAddress     common.Address
}

// IEigenLayerWithdrawerAndNonce is an auto generated low-level Go binding around an user-defined struct.
type IEigenLayerWithdrawerAndNonce struct {
	Withdrawer common.Address
	Nonce      *big.Int
}

// PufferVaultV3MetaData contains all meta data concerning the PufferVaultV3 contract.
var PufferVaultV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"stETH\",\"type\":\"address\"},{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"contractILidoWithdrawalQueue\",\"name\":\"lidoWithdrawalQueue\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"stETHStrategy\",\"type\":\"address\"},{\"internalType\":\"contractIEigenLayer\",\"name\":\"eigenStrategyManager\",\"type\":\"address\"},{\"internalType\":\"contractIPufferOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"contractIDelegationManager\",\"name\":\"delegationManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"AccessManagedInvalidAuthority\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"name\":\"AccessManagedRequiredDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AccessManagedUnauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositAndWithdrawalForbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExitFeeBasisPoints\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalsAreDisabled\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalAmount\",\"type\":\"uint256\"}],\"name\":\"AssetsWithdrawnToday\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"name\":\"ClaimedWithdrawals\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DailyWithdrawalLimitReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"oldLimit\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"newLimit\",\"type\":\"uint96\"}],\"name\":\"DailyWithdrawalLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"ExitFeeBasisPointsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedWithdrawal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualWithdrawal\",\"type\":\"uint256\"}],\"name\":\"LidoWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"name\":\"RequestedWithdrawals\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferredETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousTotalRewardsAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalRewardsAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedETHAmount\",\"type\":\"uint256\"}],\"name\":\"UpdatedTotalRewardsAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUFFER_ORACLE\",\"outputs\":[{\"internalType\":\"contractIPufferOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"}],\"internalType\":\"structIEigenLayer.WithdrawerAndNonce\",\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"delegatedAddress\",\"type\":\"address\"}],\"internalType\":\"structIEigenLayer.QueuedWithdrawal\",\"name\":\"queuedWithdrawal\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\"}],\"name\":\"claimWithdrawalFromEigenLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"}],\"internalType\":\"structIEigenLayer.WithdrawerAndNonce\",\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"delegatedAddress\",\"type\":\"address\"}],\"internalType\":\"structIEigenLayer.QueuedWithdrawal\",\"name\":\"queuedWithdrawal\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"claimWithdrawalFromEigenLayerM2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"name\":\"claimWithdrawalsFromLido\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToSharesUp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"depositETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stETHSharesAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"depositStETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToEigenLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getELBackingEthAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExitFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingLidoETHAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemainingAssetsDailyWithdrawalLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRewardDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRewardMintAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"initiateETHWithdrawalsFromLido\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesToWithdraw\",\"type\":\"uint256\"}],\"name\":\"initiateStETHWithdrawalFromEigenLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isConsumingScheduledOp\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAssets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardsAmount\",\"type\":\"uint256\"}],\"name\":\"mintRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethToPufETHRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pufETHAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pufETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"}],\"name\":\"revertMintRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"newLimit\",\"type\":\"uint96\"}],\"name\":\"setDailyWithdrawalLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newExitFeeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setExitFeeBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"}],\"name\":\"transferETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PufferVaultV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use PufferVaultV3MetaData.ABI instead.
var PufferVaultV3ABI = PufferVaultV3MetaData.ABI

// PufferVaultV3 is an auto generated Go binding around an Ethereum contract.
type PufferVaultV3 struct {
	PufferVaultV3Caller     // Read-only binding to the contract
	PufferVaultV3Transactor // Write-only binding to the contract
	PufferVaultV3Filterer   // Log filterer for contract events
}

// PufferVaultV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type PufferVaultV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PufferVaultV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PufferVaultV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PufferVaultV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PufferVaultV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PufferVaultV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PufferVaultV3Session struct {
	Contract     *PufferVaultV3    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PufferVaultV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PufferVaultV3CallerSession struct {
	Contract *PufferVaultV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PufferVaultV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PufferVaultV3TransactorSession struct {
	Contract     *PufferVaultV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PufferVaultV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type PufferVaultV3Raw struct {
	Contract *PufferVaultV3 // Generic contract binding to access the raw methods on
}

// PufferVaultV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PufferVaultV3CallerRaw struct {
	Contract *PufferVaultV3Caller // Generic read-only contract binding to access the raw methods on
}

// PufferVaultV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PufferVaultV3TransactorRaw struct {
	Contract *PufferVaultV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPufferVaultV3 creates a new instance of PufferVaultV3, bound to a specific deployed contract.
func NewPufferVaultV3(address common.Address, backend bind.ContractBackend) (*PufferVaultV3, error) {
	contract, err := bindPufferVaultV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3{PufferVaultV3Caller: PufferVaultV3Caller{contract: contract}, PufferVaultV3Transactor: PufferVaultV3Transactor{contract: contract}, PufferVaultV3Filterer: PufferVaultV3Filterer{contract: contract}}, nil
}

// NewPufferVaultV3Caller creates a new read-only instance of PufferVaultV3, bound to a specific deployed contract.
func NewPufferVaultV3Caller(address common.Address, caller bind.ContractCaller) (*PufferVaultV3Caller, error) {
	contract, err := bindPufferVaultV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3Caller{contract: contract}, nil
}

// NewPufferVaultV3Transactor creates a new write-only instance of PufferVaultV3, bound to a specific deployed contract.
func NewPufferVaultV3Transactor(address common.Address, transactor bind.ContractTransactor) (*PufferVaultV3Transactor, error) {
	contract, err := bindPufferVaultV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3Transactor{contract: contract}, nil
}

// NewPufferVaultV3Filterer creates a new log filterer instance of PufferVaultV3, bound to a specific deployed contract.
func NewPufferVaultV3Filterer(address common.Address, filterer bind.ContractFilterer) (*PufferVaultV3Filterer, error) {
	contract, err := bindPufferVaultV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3Filterer{contract: contract}, nil
}

// bindPufferVaultV3 binds a generic wrapper to an already deployed contract.
func bindPufferVaultV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PufferVaultV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PufferVaultV3 *PufferVaultV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PufferVaultV3.Contract.PufferVaultV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PufferVaultV3 *PufferVaultV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.PufferVaultV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PufferVaultV3 *PufferVaultV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.PufferVaultV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PufferVaultV3 *PufferVaultV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PufferVaultV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PufferVaultV3 *PufferVaultV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PufferVaultV3 *PufferVaultV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PufferVaultV3 *PufferVaultV3Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PufferVaultV3 *PufferVaultV3Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _PufferVaultV3.Contract.DOMAINSEPARATOR(&_PufferVaultV3.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PufferVaultV3 *PufferVaultV3CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _PufferVaultV3.Contract.DOMAINSEPARATOR(&_PufferVaultV3.CallOpts)
}

// PUFFERORACLE is a free data retrieval call binding the contract method 0x0195c505.
//
// Solidity: function PUFFER_ORACLE() view returns(address)
func (_PufferVaultV3 *PufferVaultV3Caller) PUFFERORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "PUFFER_ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PUFFERORACLE is a free data retrieval call binding the contract method 0x0195c505.
//
// Solidity: function PUFFER_ORACLE() view returns(address)
func (_PufferVaultV3 *PufferVaultV3Session) PUFFERORACLE() (common.Address, error) {
	return _PufferVaultV3.Contract.PUFFERORACLE(&_PufferVaultV3.CallOpts)
}

// PUFFERORACLE is a free data retrieval call binding the contract method 0x0195c505.
//
// Solidity: function PUFFER_ORACLE() view returns(address)
func (_PufferVaultV3 *PufferVaultV3CallerSession) PUFFERORACLE() (common.Address, error) {
	return _PufferVaultV3.Contract.PUFFERORACLE(&_PufferVaultV3.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PufferVaultV3 *PufferVaultV3Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PufferVaultV3 *PufferVaultV3Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _PufferVaultV3.Contract.UPGRADEINTERFACEVERSION(&_PufferVaultV3.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PufferVaultV3 *PufferVaultV3CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PufferVaultV3.Contract.UPGRADEINTERFACEVERSION(&_PufferVaultV3.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.Allowance(&_PufferVaultV3.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.Allowance(&_PufferVaultV3.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_PufferVaultV3 *PufferVaultV3Caller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_PufferVaultV3 *PufferVaultV3Session) Asset() (common.Address, error) {
	return _PufferVaultV3.Contract.Asset(&_PufferVaultV3.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_PufferVaultV3 *PufferVaultV3CallerSession) Asset() (common.Address, error) {
	return _PufferVaultV3.Contract.Asset(&_PufferVaultV3.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_PufferVaultV3 *PufferVaultV3Caller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_PufferVaultV3 *PufferVaultV3Session) Authority() (common.Address, error) {
	return _PufferVaultV3.Contract.Authority(&_PufferVaultV3.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_PufferVaultV3 *PufferVaultV3CallerSession) Authority() (common.Address, error) {
	return _PufferVaultV3.Contract.Authority(&_PufferVaultV3.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.BalanceOf(&_PufferVaultV3.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.BalanceOf(&_PufferVaultV3.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.ConvertToAssets(&_PufferVaultV3.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.ConvertToAssets(&_PufferVaultV3.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.ConvertToShares(&_PufferVaultV3.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.ConvertToShares(&_PufferVaultV3.CallOpts, assets)
}

// ConvertToSharesUp is a free data retrieval call binding the contract method 0x699beb59.
//
// Solidity: function convertToSharesUp(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) ConvertToSharesUp(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "convertToSharesUp", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToSharesUp is a free data retrieval call binding the contract method 0x699beb59.
//
// Solidity: function convertToSharesUp(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) ConvertToSharesUp(assets *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.ConvertToSharesUp(&_PufferVaultV3.CallOpts, assets)
}

// ConvertToSharesUp is a free data retrieval call binding the contract method 0x699beb59.
//
// Solidity: function convertToSharesUp(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) ConvertToSharesUp(assets *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.ConvertToSharesUp(&_PufferVaultV3.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_PufferVaultV3 *PufferVaultV3Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_PufferVaultV3 *PufferVaultV3Session) Decimals() (uint8, error) {
	return _PufferVaultV3.Contract.Decimals(&_PufferVaultV3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_PufferVaultV3 *PufferVaultV3CallerSession) Decimals() (uint8, error) {
	return _PufferVaultV3.Contract.Decimals(&_PufferVaultV3.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PufferVaultV3 *PufferVaultV3Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PufferVaultV3 *PufferVaultV3Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _PufferVaultV3.Contract.Eip712Domain(&_PufferVaultV3.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PufferVaultV3 *PufferVaultV3CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _PufferVaultV3.Contract.Eip712Domain(&_PufferVaultV3.CallOpts)
}

// GetELBackingEthAmount is a free data retrieval call binding the contract method 0xed344a22.
//
// Solidity: function getELBackingEthAmount() view returns(uint256 ethAmount)
func (_PufferVaultV3 *PufferVaultV3Caller) GetELBackingEthAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "getELBackingEthAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetELBackingEthAmount is a free data retrieval call binding the contract method 0xed344a22.
//
// Solidity: function getELBackingEthAmount() view returns(uint256 ethAmount)
func (_PufferVaultV3 *PufferVaultV3Session) GetELBackingEthAmount() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetELBackingEthAmount(&_PufferVaultV3.CallOpts)
}

// GetELBackingEthAmount is a free data retrieval call binding the contract method 0xed344a22.
//
// Solidity: function getELBackingEthAmount() view returns(uint256 ethAmount)
func (_PufferVaultV3 *PufferVaultV3CallerSession) GetELBackingEthAmount() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetELBackingEthAmount(&_PufferVaultV3.CallOpts)
}

// GetExitFeeBasisPoints is a free data retrieval call binding the contract method 0x8fd356ff.
//
// Solidity: function getExitFeeBasisPoints() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) GetExitFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "getExitFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExitFeeBasisPoints is a free data retrieval call binding the contract method 0x8fd356ff.
//
// Solidity: function getExitFeeBasisPoints() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) GetExitFeeBasisPoints() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetExitFeeBasisPoints(&_PufferVaultV3.CallOpts)
}

// GetExitFeeBasisPoints is a free data retrieval call binding the contract method 0x8fd356ff.
//
// Solidity: function getExitFeeBasisPoints() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) GetExitFeeBasisPoints() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetExitFeeBasisPoints(&_PufferVaultV3.CallOpts)
}

// GetPendingLidoETHAmount is a free data retrieval call binding the contract method 0xf7fe1bd1.
//
// Solidity: function getPendingLidoETHAmount() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) GetPendingLidoETHAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "getPendingLidoETHAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingLidoETHAmount is a free data retrieval call binding the contract method 0xf7fe1bd1.
//
// Solidity: function getPendingLidoETHAmount() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) GetPendingLidoETHAmount() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetPendingLidoETHAmount(&_PufferVaultV3.CallOpts)
}

// GetPendingLidoETHAmount is a free data retrieval call binding the contract method 0xf7fe1bd1.
//
// Solidity: function getPendingLidoETHAmount() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) GetPendingLidoETHAmount() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetPendingLidoETHAmount(&_PufferVaultV3.CallOpts)
}

// GetRemainingAssetsDailyWithdrawalLimit is a free data retrieval call binding the contract method 0xaf24278c.
//
// Solidity: function getRemainingAssetsDailyWithdrawalLimit() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) GetRemainingAssetsDailyWithdrawalLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "getRemainingAssetsDailyWithdrawalLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingAssetsDailyWithdrawalLimit is a free data retrieval call binding the contract method 0xaf24278c.
//
// Solidity: function getRemainingAssetsDailyWithdrawalLimit() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) GetRemainingAssetsDailyWithdrawalLimit() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetRemainingAssetsDailyWithdrawalLimit(&_PufferVaultV3.CallOpts)
}

// GetRemainingAssetsDailyWithdrawalLimit is a free data retrieval call binding the contract method 0xaf24278c.
//
// Solidity: function getRemainingAssetsDailyWithdrawalLimit() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) GetRemainingAssetsDailyWithdrawalLimit() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetRemainingAssetsDailyWithdrawalLimit(&_PufferVaultV3.CallOpts)
}

// GetTotalRewardDepositAmount is a free data retrieval call binding the contract method 0x9d0d951e.
//
// Solidity: function getTotalRewardDepositAmount() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) GetTotalRewardDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "getTotalRewardDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRewardDepositAmount is a free data retrieval call binding the contract method 0x9d0d951e.
//
// Solidity: function getTotalRewardDepositAmount() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) GetTotalRewardDepositAmount() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetTotalRewardDepositAmount(&_PufferVaultV3.CallOpts)
}

// GetTotalRewardDepositAmount is a free data retrieval call binding the contract method 0x9d0d951e.
//
// Solidity: function getTotalRewardDepositAmount() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) GetTotalRewardDepositAmount() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetTotalRewardDepositAmount(&_PufferVaultV3.CallOpts)
}

// GetTotalRewardMintAmount is a free data retrieval call binding the contract method 0x5612ba22.
//
// Solidity: function getTotalRewardMintAmount() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) GetTotalRewardMintAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "getTotalRewardMintAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRewardMintAmount is a free data retrieval call binding the contract method 0x5612ba22.
//
// Solidity: function getTotalRewardMintAmount() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) GetTotalRewardMintAmount() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetTotalRewardMintAmount(&_PufferVaultV3.CallOpts)
}

// GetTotalRewardMintAmount is a free data retrieval call binding the contract method 0x5612ba22.
//
// Solidity: function getTotalRewardMintAmount() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) GetTotalRewardMintAmount() (*big.Int, error) {
	return _PufferVaultV3.Contract.GetTotalRewardMintAmount(&_PufferVaultV3.CallOpts)
}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_PufferVaultV3 *PufferVaultV3Caller) IsConsumingScheduledOp(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "isConsumingScheduledOp")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_PufferVaultV3 *PufferVaultV3Session) IsConsumingScheduledOp() ([4]byte, error) {
	return _PufferVaultV3.Contract.IsConsumingScheduledOp(&_PufferVaultV3.CallOpts)
}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_PufferVaultV3 *PufferVaultV3CallerSession) IsConsumingScheduledOp() ([4]byte, error) {
	return _PufferVaultV3.Contract.IsConsumingScheduledOp(&_PufferVaultV3.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.MaxDeposit(&_PufferVaultV3.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.MaxDeposit(&_PufferVaultV3.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.MaxMint(&_PufferVaultV3.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.MaxMint(&_PufferVaultV3.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 maxShares)
func (_PufferVaultV3 *PufferVaultV3Caller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 maxShares)
func (_PufferVaultV3 *PufferVaultV3Session) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.MaxRedeem(&_PufferVaultV3.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 maxShares)
func (_PufferVaultV3 *PufferVaultV3CallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.MaxRedeem(&_PufferVaultV3.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 maxAssets)
func (_PufferVaultV3 *PufferVaultV3Caller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 maxAssets)
func (_PufferVaultV3 *PufferVaultV3Session) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.MaxWithdraw(&_PufferVaultV3.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 maxAssets)
func (_PufferVaultV3 *PufferVaultV3CallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.MaxWithdraw(&_PufferVaultV3.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PufferVaultV3 *PufferVaultV3Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PufferVaultV3 *PufferVaultV3Session) Name() (string, error) {
	return _PufferVaultV3.Contract.Name(&_PufferVaultV3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PufferVaultV3 *PufferVaultV3CallerSession) Name() (string, error) {
	return _PufferVaultV3.Contract.Name(&_PufferVaultV3.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) Nonces(owner common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.Nonces(&_PufferVaultV3.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _PufferVaultV3.Contract.Nonces(&_PufferVaultV3.CallOpts, owner)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.PreviewDeposit(&_PufferVaultV3.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.PreviewDeposit(&_PufferVaultV3.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.PreviewMint(&_PufferVaultV3.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.PreviewMint(&_PufferVaultV3.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.PreviewRedeem(&_PufferVaultV3.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.PreviewRedeem(&_PufferVaultV3.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.PreviewWithdraw(&_PufferVaultV3.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _PufferVaultV3.Contract.PreviewWithdraw(&_PufferVaultV3.CallOpts, assets)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PufferVaultV3 *PufferVaultV3Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PufferVaultV3 *PufferVaultV3Session) ProxiableUUID() ([32]byte, error) {
	return _PufferVaultV3.Contract.ProxiableUUID(&_PufferVaultV3.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PufferVaultV3 *PufferVaultV3CallerSession) ProxiableUUID() ([32]byte, error) {
	return _PufferVaultV3.Contract.ProxiableUUID(&_PufferVaultV3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PufferVaultV3 *PufferVaultV3Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PufferVaultV3 *PufferVaultV3Session) Symbol() (string, error) {
	return _PufferVaultV3.Contract.Symbol(&_PufferVaultV3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PufferVaultV3 *PufferVaultV3CallerSession) Symbol() (string, error) {
	return _PufferVaultV3.Contract.Symbol(&_PufferVaultV3.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) TotalAssets() (*big.Int, error) {
	return _PufferVaultV3.Contract.TotalAssets(&_PufferVaultV3.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) TotalAssets() (*big.Int, error) {
	return _PufferVaultV3.Contract.TotalAssets(&_PufferVaultV3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PufferVaultV3.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) TotalSupply() (*big.Int, error) {
	return _PufferVaultV3.Contract.TotalSupply(&_PufferVaultV3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PufferVaultV3 *PufferVaultV3CallerSession) TotalSupply() (*big.Int, error) {
	return _PufferVaultV3.Contract.TotalSupply(&_PufferVaultV3.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PufferVaultV3 *PufferVaultV3Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PufferVaultV3 *PufferVaultV3Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Approve(&_PufferVaultV3.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Approve(&_PufferVaultV3.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 shares) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) Burn(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "burn", shares)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 shares) returns()
func (_PufferVaultV3 *PufferVaultV3Session) Burn(shares *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Burn(&_PufferVaultV3.TransactOpts, shares)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 shares) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Burn(shares *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Burn(&_PufferVaultV3.TransactOpts, shares)
}

// ClaimWithdrawalFromEigenLayer is a paid mutator transaction binding the contract method 0xfdcffbac.
//
// Solidity: function claimWithdrawalFromEigenLayer((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) ClaimWithdrawalFromEigenLayer(opts *bind.TransactOpts, queuedWithdrawal IEigenLayerQueuedWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "claimWithdrawalFromEigenLayer", queuedWithdrawal, tokens, middlewareTimesIndex)
}

// ClaimWithdrawalFromEigenLayer is a paid mutator transaction binding the contract method 0xfdcffbac.
//
// Solidity: function claimWithdrawalFromEigenLayer((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex) returns()
func (_PufferVaultV3 *PufferVaultV3Session) ClaimWithdrawalFromEigenLayer(queuedWithdrawal IEigenLayerQueuedWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.ClaimWithdrawalFromEigenLayer(&_PufferVaultV3.TransactOpts, queuedWithdrawal, tokens, middlewareTimesIndex)
}

// ClaimWithdrawalFromEigenLayer is a paid mutator transaction binding the contract method 0xfdcffbac.
//
// Solidity: function claimWithdrawalFromEigenLayer((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) ClaimWithdrawalFromEigenLayer(queuedWithdrawal IEigenLayerQueuedWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.ClaimWithdrawalFromEigenLayer(&_PufferVaultV3.TransactOpts, queuedWithdrawal, tokens, middlewareTimesIndex)
}

// ClaimWithdrawalFromEigenLayerM2 is a paid mutator transaction binding the contract method 0x34201a64.
//
// Solidity: function claimWithdrawalFromEigenLayerM2((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, uint256 nonce) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) ClaimWithdrawalFromEigenLayerM2(opts *bind.TransactOpts, queuedWithdrawal IEigenLayerQueuedWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "claimWithdrawalFromEigenLayerM2", queuedWithdrawal, tokens, middlewareTimesIndex, nonce)
}

// ClaimWithdrawalFromEigenLayerM2 is a paid mutator transaction binding the contract method 0x34201a64.
//
// Solidity: function claimWithdrawalFromEigenLayerM2((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, uint256 nonce) returns()
func (_PufferVaultV3 *PufferVaultV3Session) ClaimWithdrawalFromEigenLayerM2(queuedWithdrawal IEigenLayerQueuedWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.ClaimWithdrawalFromEigenLayerM2(&_PufferVaultV3.TransactOpts, queuedWithdrawal, tokens, middlewareTimesIndex, nonce)
}

// ClaimWithdrawalFromEigenLayerM2 is a paid mutator transaction binding the contract method 0x34201a64.
//
// Solidity: function claimWithdrawalFromEigenLayerM2((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, uint256 nonce) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) ClaimWithdrawalFromEigenLayerM2(queuedWithdrawal IEigenLayerQueuedWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.ClaimWithdrawalFromEigenLayerM2(&_PufferVaultV3.TransactOpts, queuedWithdrawal, tokens, middlewareTimesIndex, nonce)
}

// ClaimWithdrawalsFromLido is a paid mutator transaction binding the contract method 0x677a11e9.
//
// Solidity: function claimWithdrawalsFromLido(uint256[] requestIds) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) ClaimWithdrawalsFromLido(opts *bind.TransactOpts, requestIds []*big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "claimWithdrawalsFromLido", requestIds)
}

// ClaimWithdrawalsFromLido is a paid mutator transaction binding the contract method 0x677a11e9.
//
// Solidity: function claimWithdrawalsFromLido(uint256[] requestIds) returns()
func (_PufferVaultV3 *PufferVaultV3Session) ClaimWithdrawalsFromLido(requestIds []*big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.ClaimWithdrawalsFromLido(&_PufferVaultV3.TransactOpts, requestIds)
}

// ClaimWithdrawalsFromLido is a paid mutator transaction binding the contract method 0x677a11e9.
//
// Solidity: function claimWithdrawalsFromLido(uint256[] requestIds) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) ClaimWithdrawalsFromLido(requestIds []*big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.ClaimWithdrawalsFromLido(&_PufferVaultV3.TransactOpts, requestIds)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Transactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Deposit(&_PufferVaultV3.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Deposit(&_PufferVaultV3.TransactOpts, assets, receiver)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address receiver) payable returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Transactor) DepositETH(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "depositETH", receiver)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address receiver) payable returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) DepositETH(receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.DepositETH(&_PufferVaultV3.TransactOpts, receiver)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address receiver) payable returns(uint256)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) DepositETH(receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.DepositETH(&_PufferVaultV3.TransactOpts, receiver)
}

// DepositRewards is a paid mutator transaction binding the contract method 0x152111f7.
//
// Solidity: function depositRewards() payable returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) DepositRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "depositRewards")
}

// DepositRewards is a paid mutator transaction binding the contract method 0x152111f7.
//
// Solidity: function depositRewards() payable returns()
func (_PufferVaultV3 *PufferVaultV3Session) DepositRewards() (*types.Transaction, error) {
	return _PufferVaultV3.Contract.DepositRewards(&_PufferVaultV3.TransactOpts)
}

// DepositRewards is a paid mutator transaction binding the contract method 0x152111f7.
//
// Solidity: function depositRewards() payable returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) DepositRewards() (*types.Transaction, error) {
	return _PufferVaultV3.Contract.DepositRewards(&_PufferVaultV3.TransactOpts)
}

// DepositStETH is a paid mutator transaction binding the contract method 0xf6dbd16f.
//
// Solidity: function depositStETH(uint256 stETHSharesAmount, address receiver) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Transactor) DepositStETH(opts *bind.TransactOpts, stETHSharesAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "depositStETH", stETHSharesAmount, receiver)
}

// DepositStETH is a paid mutator transaction binding the contract method 0xf6dbd16f.
//
// Solidity: function depositStETH(uint256 stETHSharesAmount, address receiver) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) DepositStETH(stETHSharesAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.DepositStETH(&_PufferVaultV3.TransactOpts, stETHSharesAmount, receiver)
}

// DepositStETH is a paid mutator transaction binding the contract method 0xf6dbd16f.
//
// Solidity: function depositStETH(uint256 stETHSharesAmount, address receiver) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) DepositStETH(stETHSharesAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.DepositStETH(&_PufferVaultV3.TransactOpts, stETHSharesAmount, receiver)
}

// DepositToEigenLayer is a paid mutator transaction binding the contract method 0x008e0590.
//
// Solidity: function depositToEigenLayer(uint256 amount) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) DepositToEigenLayer(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "depositToEigenLayer", amount)
}

// DepositToEigenLayer is a paid mutator transaction binding the contract method 0x008e0590.
//
// Solidity: function depositToEigenLayer(uint256 amount) returns()
func (_PufferVaultV3 *PufferVaultV3Session) DepositToEigenLayer(amount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.DepositToEigenLayer(&_PufferVaultV3.TransactOpts, amount)
}

// DepositToEigenLayer is a paid mutator transaction binding the contract method 0x008e0590.
//
// Solidity: function depositToEigenLayer(uint256 amount) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) DepositToEigenLayer(amount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.DepositToEigenLayer(&_PufferVaultV3.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PufferVaultV3 *PufferVaultV3Session) Initialize() (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Initialize(&_PufferVaultV3.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Initialize() (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Initialize(&_PufferVaultV3.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address accessManager) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) Initialize0(opts *bind.TransactOpts, accessManager common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "initialize0", accessManager)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address accessManager) returns()
func (_PufferVaultV3 *PufferVaultV3Session) Initialize0(accessManager common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Initialize0(&_PufferVaultV3.TransactOpts, accessManager)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address accessManager) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Initialize0(accessManager common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Initialize0(&_PufferVaultV3.TransactOpts, accessManager)
}

// InitiateETHWithdrawalsFromLido is a paid mutator transaction binding the contract method 0x593961de.
//
// Solidity: function initiateETHWithdrawalsFromLido(uint256[] amounts) returns(uint256[] requestIds)
func (_PufferVaultV3 *PufferVaultV3Transactor) InitiateETHWithdrawalsFromLido(opts *bind.TransactOpts, amounts []*big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "initiateETHWithdrawalsFromLido", amounts)
}

// InitiateETHWithdrawalsFromLido is a paid mutator transaction binding the contract method 0x593961de.
//
// Solidity: function initiateETHWithdrawalsFromLido(uint256[] amounts) returns(uint256[] requestIds)
func (_PufferVaultV3 *PufferVaultV3Session) InitiateETHWithdrawalsFromLido(amounts []*big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.InitiateETHWithdrawalsFromLido(&_PufferVaultV3.TransactOpts, amounts)
}

// InitiateETHWithdrawalsFromLido is a paid mutator transaction binding the contract method 0x593961de.
//
// Solidity: function initiateETHWithdrawalsFromLido(uint256[] amounts) returns(uint256[] requestIds)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) InitiateETHWithdrawalsFromLido(amounts []*big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.InitiateETHWithdrawalsFromLido(&_PufferVaultV3.TransactOpts, amounts)
}

// InitiateStETHWithdrawalFromEigenLayer is a paid mutator transaction binding the contract method 0x402064a7.
//
// Solidity: function initiateStETHWithdrawalFromEigenLayer(uint256 sharesToWithdraw) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) InitiateStETHWithdrawalFromEigenLayer(opts *bind.TransactOpts, sharesToWithdraw *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "initiateStETHWithdrawalFromEigenLayer", sharesToWithdraw)
}

// InitiateStETHWithdrawalFromEigenLayer is a paid mutator transaction binding the contract method 0x402064a7.
//
// Solidity: function initiateStETHWithdrawalFromEigenLayer(uint256 sharesToWithdraw) returns()
func (_PufferVaultV3 *PufferVaultV3Session) InitiateStETHWithdrawalFromEigenLayer(sharesToWithdraw *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.InitiateStETHWithdrawalFromEigenLayer(&_PufferVaultV3.TransactOpts, sharesToWithdraw)
}

// InitiateStETHWithdrawalFromEigenLayer is a paid mutator transaction binding the contract method 0x402064a7.
//
// Solidity: function initiateStETHWithdrawalFromEigenLayer(uint256 sharesToWithdraw) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) InitiateStETHWithdrawalFromEigenLayer(sharesToWithdraw *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.InitiateStETHWithdrawalFromEigenLayer(&_PufferVaultV3.TransactOpts, sharesToWithdraw)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Transactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Mint(&_PufferVaultV3.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Mint(&_PufferVaultV3.TransactOpts, shares, receiver)
}

// MintRewards is a paid mutator transaction binding the contract method 0x4828ced9.
//
// Solidity: function mintRewards(uint256 rewardsAmount) returns(uint256 ethToPufETHRate, uint256 pufETHAmount)
func (_PufferVaultV3 *PufferVaultV3Transactor) MintRewards(opts *bind.TransactOpts, rewardsAmount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "mintRewards", rewardsAmount)
}

// MintRewards is a paid mutator transaction binding the contract method 0x4828ced9.
//
// Solidity: function mintRewards(uint256 rewardsAmount) returns(uint256 ethToPufETHRate, uint256 pufETHAmount)
func (_PufferVaultV3 *PufferVaultV3Session) MintRewards(rewardsAmount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.MintRewards(&_PufferVaultV3.TransactOpts, rewardsAmount)
}

// MintRewards is a paid mutator transaction binding the contract method 0x4828ced9.
//
// Solidity: function mintRewards(uint256 rewardsAmount) returns(uint256 ethToPufETHRate, uint256 pufETHAmount)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) MintRewards(rewardsAmount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.MintRewards(&_PufferVaultV3.TransactOpts, rewardsAmount)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_PufferVaultV3 *PufferVaultV3Transactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_PufferVaultV3 *PufferVaultV3Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.OnERC721Received(&_PufferVaultV3.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.OnERC721Received(&_PufferVaultV3.TransactOpts, arg0, arg1, arg2, arg3)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PufferVaultV3 *PufferVaultV3Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Permit(&_PufferVaultV3.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Permit(&_PufferVaultV3.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Transactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Redeem(&_PufferVaultV3.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Redeem(&_PufferVaultV3.TransactOpts, shares, receiver, owner)
}

// RevertMintRewards is a paid mutator transaction binding the contract method 0xe7fc6f73.
//
// Solidity: function revertMintRewards(uint256 pufETHAmount, uint256 ethAmount) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) RevertMintRewards(opts *bind.TransactOpts, pufETHAmount *big.Int, ethAmount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "revertMintRewards", pufETHAmount, ethAmount)
}

// RevertMintRewards is a paid mutator transaction binding the contract method 0xe7fc6f73.
//
// Solidity: function revertMintRewards(uint256 pufETHAmount, uint256 ethAmount) returns()
func (_PufferVaultV3 *PufferVaultV3Session) RevertMintRewards(pufETHAmount *big.Int, ethAmount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.RevertMintRewards(&_PufferVaultV3.TransactOpts, pufETHAmount, ethAmount)
}

// RevertMintRewards is a paid mutator transaction binding the contract method 0xe7fc6f73.
//
// Solidity: function revertMintRewards(uint256 pufETHAmount, uint256 ethAmount) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) RevertMintRewards(pufETHAmount *big.Int, ethAmount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.RevertMintRewards(&_PufferVaultV3.TransactOpts, pufETHAmount, ethAmount)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) SetAuthority(opts *bind.TransactOpts, newAuthority common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "setAuthority", newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_PufferVaultV3 *PufferVaultV3Session) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.SetAuthority(&_PufferVaultV3.TransactOpts, newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.SetAuthority(&_PufferVaultV3.TransactOpts, newAuthority)
}

// SetDailyWithdrawalLimit is a paid mutator transaction binding the contract method 0x9e9406dc.
//
// Solidity: function setDailyWithdrawalLimit(uint96 newLimit) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) SetDailyWithdrawalLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "setDailyWithdrawalLimit", newLimit)
}

// SetDailyWithdrawalLimit is a paid mutator transaction binding the contract method 0x9e9406dc.
//
// Solidity: function setDailyWithdrawalLimit(uint96 newLimit) returns()
func (_PufferVaultV3 *PufferVaultV3Session) SetDailyWithdrawalLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.SetDailyWithdrawalLimit(&_PufferVaultV3.TransactOpts, newLimit)
}

// SetDailyWithdrawalLimit is a paid mutator transaction binding the contract method 0x9e9406dc.
//
// Solidity: function setDailyWithdrawalLimit(uint96 newLimit) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) SetDailyWithdrawalLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.SetDailyWithdrawalLimit(&_PufferVaultV3.TransactOpts, newLimit)
}

// SetExitFeeBasisPoints is a paid mutator transaction binding the contract method 0x5b63b05c.
//
// Solidity: function setExitFeeBasisPoints(uint256 newExitFeeBasisPoints) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) SetExitFeeBasisPoints(opts *bind.TransactOpts, newExitFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "setExitFeeBasisPoints", newExitFeeBasisPoints)
}

// SetExitFeeBasisPoints is a paid mutator transaction binding the contract method 0x5b63b05c.
//
// Solidity: function setExitFeeBasisPoints(uint256 newExitFeeBasisPoints) returns()
func (_PufferVaultV3 *PufferVaultV3Session) SetExitFeeBasisPoints(newExitFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.SetExitFeeBasisPoints(&_PufferVaultV3.TransactOpts, newExitFeeBasisPoints)
}

// SetExitFeeBasisPoints is a paid mutator transaction binding the contract method 0x5b63b05c.
//
// Solidity: function setExitFeeBasisPoints(uint256 newExitFeeBasisPoints) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) SetExitFeeBasisPoints(newExitFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.SetExitFeeBasisPoints(&_PufferVaultV3.TransactOpts, newExitFeeBasisPoints)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PufferVaultV3 *PufferVaultV3Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PufferVaultV3 *PufferVaultV3Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Transfer(&_PufferVaultV3.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Transfer(&_PufferVaultV3.TransactOpts, to, value)
}

// TransferETH is a paid mutator transaction binding the contract method 0x7b1a4909.
//
// Solidity: function transferETH(address to, uint256 ethAmount) returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) TransferETH(opts *bind.TransactOpts, to common.Address, ethAmount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "transferETH", to, ethAmount)
}

// TransferETH is a paid mutator transaction binding the contract method 0x7b1a4909.
//
// Solidity: function transferETH(address to, uint256 ethAmount) returns()
func (_PufferVaultV3 *PufferVaultV3Session) TransferETH(to common.Address, ethAmount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.TransferETH(&_PufferVaultV3.TransactOpts, to, ethAmount)
}

// TransferETH is a paid mutator transaction binding the contract method 0x7b1a4909.
//
// Solidity: function transferETH(address to, uint256 ethAmount) returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) TransferETH(to common.Address, ethAmount *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.TransferETH(&_PufferVaultV3.TransactOpts, to, ethAmount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PufferVaultV3 *PufferVaultV3Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PufferVaultV3 *PufferVaultV3Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.TransferFrom(&_PufferVaultV3.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.TransferFrom(&_PufferVaultV3.TransactOpts, from, to, value)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PufferVaultV3 *PufferVaultV3Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.UpgradeToAndCall(&_PufferVaultV3.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.UpgradeToAndCall(&_PufferVaultV3.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Transactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3Session) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Withdraw(&_PufferVaultV3.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Withdraw(&_PufferVaultV3.TransactOpts, assets, receiver, owner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PufferVaultV3 *PufferVaultV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PufferVaultV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PufferVaultV3 *PufferVaultV3Session) Receive() (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Receive(&_PufferVaultV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PufferVaultV3 *PufferVaultV3TransactorSession) Receive() (*types.Transaction, error) {
	return _PufferVaultV3.Contract.Receive(&_PufferVaultV3.TransactOpts)
}

// PufferVaultV3ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PufferVaultV3 contract.
type PufferVaultV3ApprovalIterator struct {
	Event *PufferVaultV3Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3Approval represents a Approval event raised by the PufferVaultV3 contract.
type PufferVaultV3Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PufferVaultV3ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3ApprovalIterator{contract: _PufferVaultV3.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PufferVaultV3Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3Approval)
				if err := _PufferVaultV3.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseApproval(log types.Log) (*PufferVaultV3Approval, error) {
	event := new(PufferVaultV3Approval)
	if err := _PufferVaultV3.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3AssetsWithdrawnTodayIterator is returned from FilterAssetsWithdrawnToday and is used to iterate over the raw logs and unpacked data for AssetsWithdrawnToday events raised by the PufferVaultV3 contract.
type PufferVaultV3AssetsWithdrawnTodayIterator struct {
	Event *PufferVaultV3AssetsWithdrawnToday // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3AssetsWithdrawnTodayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3AssetsWithdrawnToday)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3AssetsWithdrawnToday)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3AssetsWithdrawnTodayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3AssetsWithdrawnTodayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3AssetsWithdrawnToday represents a AssetsWithdrawnToday event raised by the PufferVaultV3 contract.
type PufferVaultV3AssetsWithdrawnToday struct {
	WithdrawalAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAssetsWithdrawnToday is a free log retrieval operation binding the contract event 0x139f9ee0762f3b0c92a4b8c7b8fe8be6b12aaece4b9b22de6bf1ba1094dcd998.
//
// Solidity: event AssetsWithdrawnToday(uint256 withdrawalAmount)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterAssetsWithdrawnToday(opts *bind.FilterOpts) (*PufferVaultV3AssetsWithdrawnTodayIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "AssetsWithdrawnToday")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3AssetsWithdrawnTodayIterator{contract: _PufferVaultV3.contract, event: "AssetsWithdrawnToday", logs: logs, sub: sub}, nil
}

// WatchAssetsWithdrawnToday is a free log subscription operation binding the contract event 0x139f9ee0762f3b0c92a4b8c7b8fe8be6b12aaece4b9b22de6bf1ba1094dcd998.
//
// Solidity: event AssetsWithdrawnToday(uint256 withdrawalAmount)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchAssetsWithdrawnToday(opts *bind.WatchOpts, sink chan<- *PufferVaultV3AssetsWithdrawnToday) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "AssetsWithdrawnToday")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3AssetsWithdrawnToday)
				if err := _PufferVaultV3.contract.UnpackLog(event, "AssetsWithdrawnToday", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetsWithdrawnToday is a log parse operation binding the contract event 0x139f9ee0762f3b0c92a4b8c7b8fe8be6b12aaece4b9b22de6bf1ba1094dcd998.
//
// Solidity: event AssetsWithdrawnToday(uint256 withdrawalAmount)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseAssetsWithdrawnToday(log types.Log) (*PufferVaultV3AssetsWithdrawnToday, error) {
	event := new(PufferVaultV3AssetsWithdrawnToday)
	if err := _PufferVaultV3.contract.UnpackLog(event, "AssetsWithdrawnToday", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3AuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the PufferVaultV3 contract.
type PufferVaultV3AuthorityUpdatedIterator struct {
	Event *PufferVaultV3AuthorityUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3AuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3AuthorityUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3AuthorityUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3AuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3AuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3AuthorityUpdated represents a AuthorityUpdated event raised by the PufferVaultV3 contract.
type PufferVaultV3AuthorityUpdated struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterAuthorityUpdated(opts *bind.FilterOpts) (*PufferVaultV3AuthorityUpdatedIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "AuthorityUpdated")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3AuthorityUpdatedIterator{contract: _PufferVaultV3.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *PufferVaultV3AuthorityUpdated) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "AuthorityUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3AuthorityUpdated)
				if err := _PufferVaultV3.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorityUpdated is a log parse operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseAuthorityUpdated(log types.Log) (*PufferVaultV3AuthorityUpdated, error) {
	event := new(PufferVaultV3AuthorityUpdated)
	if err := _PufferVaultV3.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3ClaimedWithdrawalsIterator is returned from FilterClaimedWithdrawals and is used to iterate over the raw logs and unpacked data for ClaimedWithdrawals events raised by the PufferVaultV3 contract.
type PufferVaultV3ClaimedWithdrawalsIterator struct {
	Event *PufferVaultV3ClaimedWithdrawals // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3ClaimedWithdrawalsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3ClaimedWithdrawals)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3ClaimedWithdrawals)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3ClaimedWithdrawalsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3ClaimedWithdrawalsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3ClaimedWithdrawals represents a ClaimedWithdrawals event raised by the PufferVaultV3 contract.
type PufferVaultV3ClaimedWithdrawals struct {
	RequestIds []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimedWithdrawals is a free log retrieval operation binding the contract event 0xfe1f3a60946e634f858dc1f2f911c04cba9dc419a19abcb1bb1ce905ed79e325.
//
// Solidity: event ClaimedWithdrawals(uint256[] requestIds)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterClaimedWithdrawals(opts *bind.FilterOpts) (*PufferVaultV3ClaimedWithdrawalsIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "ClaimedWithdrawals")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3ClaimedWithdrawalsIterator{contract: _PufferVaultV3.contract, event: "ClaimedWithdrawals", logs: logs, sub: sub}, nil
}

// WatchClaimedWithdrawals is a free log subscription operation binding the contract event 0xfe1f3a60946e634f858dc1f2f911c04cba9dc419a19abcb1bb1ce905ed79e325.
//
// Solidity: event ClaimedWithdrawals(uint256[] requestIds)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchClaimedWithdrawals(opts *bind.WatchOpts, sink chan<- *PufferVaultV3ClaimedWithdrawals) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "ClaimedWithdrawals")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3ClaimedWithdrawals)
				if err := _PufferVaultV3.contract.UnpackLog(event, "ClaimedWithdrawals", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedWithdrawals is a log parse operation binding the contract event 0xfe1f3a60946e634f858dc1f2f911c04cba9dc419a19abcb1bb1ce905ed79e325.
//
// Solidity: event ClaimedWithdrawals(uint256[] requestIds)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseClaimedWithdrawals(log types.Log) (*PufferVaultV3ClaimedWithdrawals, error) {
	event := new(PufferVaultV3ClaimedWithdrawals)
	if err := _PufferVaultV3.contract.UnpackLog(event, "ClaimedWithdrawals", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3DailyWithdrawalLimitResetIterator is returned from FilterDailyWithdrawalLimitReset and is used to iterate over the raw logs and unpacked data for DailyWithdrawalLimitReset events raised by the PufferVaultV3 contract.
type PufferVaultV3DailyWithdrawalLimitResetIterator struct {
	Event *PufferVaultV3DailyWithdrawalLimitReset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3DailyWithdrawalLimitResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3DailyWithdrawalLimitReset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3DailyWithdrawalLimitReset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3DailyWithdrawalLimitResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3DailyWithdrawalLimitResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3DailyWithdrawalLimitReset represents a DailyWithdrawalLimitReset event raised by the PufferVaultV3 contract.
type PufferVaultV3DailyWithdrawalLimitReset struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDailyWithdrawalLimitReset is a free log retrieval operation binding the contract event 0x190567136e3dd93d29bef98a7c7c87cff34ee88e71d634b52f5fb3b531085f40.
//
// Solidity: event DailyWithdrawalLimitReset()
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterDailyWithdrawalLimitReset(opts *bind.FilterOpts) (*PufferVaultV3DailyWithdrawalLimitResetIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "DailyWithdrawalLimitReset")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3DailyWithdrawalLimitResetIterator{contract: _PufferVaultV3.contract, event: "DailyWithdrawalLimitReset", logs: logs, sub: sub}, nil
}

// WatchDailyWithdrawalLimitReset is a free log subscription operation binding the contract event 0x190567136e3dd93d29bef98a7c7c87cff34ee88e71d634b52f5fb3b531085f40.
//
// Solidity: event DailyWithdrawalLimitReset()
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchDailyWithdrawalLimitReset(opts *bind.WatchOpts, sink chan<- *PufferVaultV3DailyWithdrawalLimitReset) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "DailyWithdrawalLimitReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3DailyWithdrawalLimitReset)
				if err := _PufferVaultV3.contract.UnpackLog(event, "DailyWithdrawalLimitReset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDailyWithdrawalLimitReset is a log parse operation binding the contract event 0x190567136e3dd93d29bef98a7c7c87cff34ee88e71d634b52f5fb3b531085f40.
//
// Solidity: event DailyWithdrawalLimitReset()
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseDailyWithdrawalLimitReset(log types.Log) (*PufferVaultV3DailyWithdrawalLimitReset, error) {
	event := new(PufferVaultV3DailyWithdrawalLimitReset)
	if err := _PufferVaultV3.contract.UnpackLog(event, "DailyWithdrawalLimitReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3DailyWithdrawalLimitSetIterator is returned from FilterDailyWithdrawalLimitSet and is used to iterate over the raw logs and unpacked data for DailyWithdrawalLimitSet events raised by the PufferVaultV3 contract.
type PufferVaultV3DailyWithdrawalLimitSetIterator struct {
	Event *PufferVaultV3DailyWithdrawalLimitSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3DailyWithdrawalLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3DailyWithdrawalLimitSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3DailyWithdrawalLimitSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3DailyWithdrawalLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3DailyWithdrawalLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3DailyWithdrawalLimitSet represents a DailyWithdrawalLimitSet event raised by the PufferVaultV3 contract.
type PufferVaultV3DailyWithdrawalLimitSet struct {
	OldLimit *big.Int
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDailyWithdrawalLimitSet is a free log retrieval operation binding the contract event 0x8d5f7487ce1fd25059bd15204a55ea2c293160362b849a6f9244aec7d5a3700b.
//
// Solidity: event DailyWithdrawalLimitSet(uint96 oldLimit, uint96 newLimit)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterDailyWithdrawalLimitSet(opts *bind.FilterOpts) (*PufferVaultV3DailyWithdrawalLimitSetIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "DailyWithdrawalLimitSet")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3DailyWithdrawalLimitSetIterator{contract: _PufferVaultV3.contract, event: "DailyWithdrawalLimitSet", logs: logs, sub: sub}, nil
}

// WatchDailyWithdrawalLimitSet is a free log subscription operation binding the contract event 0x8d5f7487ce1fd25059bd15204a55ea2c293160362b849a6f9244aec7d5a3700b.
//
// Solidity: event DailyWithdrawalLimitSet(uint96 oldLimit, uint96 newLimit)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchDailyWithdrawalLimitSet(opts *bind.WatchOpts, sink chan<- *PufferVaultV3DailyWithdrawalLimitSet) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "DailyWithdrawalLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3DailyWithdrawalLimitSet)
				if err := _PufferVaultV3.contract.UnpackLog(event, "DailyWithdrawalLimitSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDailyWithdrawalLimitSet is a log parse operation binding the contract event 0x8d5f7487ce1fd25059bd15204a55ea2c293160362b849a6f9244aec7d5a3700b.
//
// Solidity: event DailyWithdrawalLimitSet(uint96 oldLimit, uint96 newLimit)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseDailyWithdrawalLimitSet(log types.Log) (*PufferVaultV3DailyWithdrawalLimitSet, error) {
	event := new(PufferVaultV3DailyWithdrawalLimitSet)
	if err := _PufferVaultV3.contract.UnpackLog(event, "DailyWithdrawalLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PufferVaultV3 contract.
type PufferVaultV3DepositIterator struct {
	Event *PufferVaultV3Deposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3Deposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3Deposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3Deposit represents a Deposit event raised by the PufferVaultV3 contract.
type PufferVaultV3Deposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*PufferVaultV3DepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3DepositIterator{contract: _PufferVaultV3.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PufferVaultV3Deposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3Deposit)
				if err := _PufferVaultV3.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseDeposit(log types.Log) (*PufferVaultV3Deposit, error) {
	event := new(PufferVaultV3Deposit)
	if err := _PufferVaultV3.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the PufferVaultV3 contract.
type PufferVaultV3EIP712DomainChangedIterator struct {
	Event *PufferVaultV3EIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3EIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3EIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3EIP712DomainChanged represents a EIP712DomainChanged event raised by the PufferVaultV3 contract.
type PufferVaultV3EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*PufferVaultV3EIP712DomainChangedIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3EIP712DomainChangedIterator{contract: _PufferVaultV3.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *PufferVaultV3EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3EIP712DomainChanged)
				if err := _PufferVaultV3.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseEIP712DomainChanged(log types.Log) (*PufferVaultV3EIP712DomainChanged, error) {
	event := new(PufferVaultV3EIP712DomainChanged)
	if err := _PufferVaultV3.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3ExitFeeBasisPointsSetIterator is returned from FilterExitFeeBasisPointsSet and is used to iterate over the raw logs and unpacked data for ExitFeeBasisPointsSet events raised by the PufferVaultV3 contract.
type PufferVaultV3ExitFeeBasisPointsSetIterator struct {
	Event *PufferVaultV3ExitFeeBasisPointsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3ExitFeeBasisPointsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3ExitFeeBasisPointsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3ExitFeeBasisPointsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3ExitFeeBasisPointsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3ExitFeeBasisPointsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3ExitFeeBasisPointsSet represents a ExitFeeBasisPointsSet event raised by the PufferVaultV3 contract.
type PufferVaultV3ExitFeeBasisPointsSet struct {
	PreviousFee *big.Int
	NewFee      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExitFeeBasisPointsSet is a free log retrieval operation binding the contract event 0xb10a745484e9798f0014ea028d76169706f92e7eea5d5bb66001c1400769785d.
//
// Solidity: event ExitFeeBasisPointsSet(uint256 previousFee, uint256 newFee)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterExitFeeBasisPointsSet(opts *bind.FilterOpts) (*PufferVaultV3ExitFeeBasisPointsSetIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "ExitFeeBasisPointsSet")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3ExitFeeBasisPointsSetIterator{contract: _PufferVaultV3.contract, event: "ExitFeeBasisPointsSet", logs: logs, sub: sub}, nil
}

// WatchExitFeeBasisPointsSet is a free log subscription operation binding the contract event 0xb10a745484e9798f0014ea028d76169706f92e7eea5d5bb66001c1400769785d.
//
// Solidity: event ExitFeeBasisPointsSet(uint256 previousFee, uint256 newFee)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchExitFeeBasisPointsSet(opts *bind.WatchOpts, sink chan<- *PufferVaultV3ExitFeeBasisPointsSet) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "ExitFeeBasisPointsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3ExitFeeBasisPointsSet)
				if err := _PufferVaultV3.contract.UnpackLog(event, "ExitFeeBasisPointsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExitFeeBasisPointsSet is a log parse operation binding the contract event 0xb10a745484e9798f0014ea028d76169706f92e7eea5d5bb66001c1400769785d.
//
// Solidity: event ExitFeeBasisPointsSet(uint256 previousFee, uint256 newFee)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseExitFeeBasisPointsSet(log types.Log) (*PufferVaultV3ExitFeeBasisPointsSet, error) {
	event := new(PufferVaultV3ExitFeeBasisPointsSet)
	if err := _PufferVaultV3.contract.UnpackLog(event, "ExitFeeBasisPointsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PufferVaultV3 contract.
type PufferVaultV3InitializedIterator struct {
	Event *PufferVaultV3Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3Initialized represents a Initialized event raised by the PufferVaultV3 contract.
type PufferVaultV3Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterInitialized(opts *bind.FilterOpts) (*PufferVaultV3InitializedIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3InitializedIterator{contract: _PufferVaultV3.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PufferVaultV3Initialized) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3Initialized)
				if err := _PufferVaultV3.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseInitialized(log types.Log) (*PufferVaultV3Initialized, error) {
	event := new(PufferVaultV3Initialized)
	if err := _PufferVaultV3.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3LidoWithdrawalIterator is returned from FilterLidoWithdrawal and is used to iterate over the raw logs and unpacked data for LidoWithdrawal events raised by the PufferVaultV3 contract.
type PufferVaultV3LidoWithdrawalIterator struct {
	Event *PufferVaultV3LidoWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3LidoWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3LidoWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3LidoWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3LidoWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3LidoWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3LidoWithdrawal represents a LidoWithdrawal event raised by the PufferVaultV3 contract.
type PufferVaultV3LidoWithdrawal struct {
	ExpectedWithdrawal *big.Int
	ActualWithdrawal   *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLidoWithdrawal is a free log retrieval operation binding the contract event 0xb5cd6ba4df0e50a9991fc91db91ea56e2f134e498a70fc7224ad61d123e5bbb0.
//
// Solidity: event LidoWithdrawal(uint256 expectedWithdrawal, uint256 actualWithdrawal)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterLidoWithdrawal(opts *bind.FilterOpts) (*PufferVaultV3LidoWithdrawalIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "LidoWithdrawal")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3LidoWithdrawalIterator{contract: _PufferVaultV3.contract, event: "LidoWithdrawal", logs: logs, sub: sub}, nil
}

// WatchLidoWithdrawal is a free log subscription operation binding the contract event 0xb5cd6ba4df0e50a9991fc91db91ea56e2f134e498a70fc7224ad61d123e5bbb0.
//
// Solidity: event LidoWithdrawal(uint256 expectedWithdrawal, uint256 actualWithdrawal)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchLidoWithdrawal(opts *bind.WatchOpts, sink chan<- *PufferVaultV3LidoWithdrawal) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "LidoWithdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3LidoWithdrawal)
				if err := _PufferVaultV3.contract.UnpackLog(event, "LidoWithdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLidoWithdrawal is a log parse operation binding the contract event 0xb5cd6ba4df0e50a9991fc91db91ea56e2f134e498a70fc7224ad61d123e5bbb0.
//
// Solidity: event LidoWithdrawal(uint256 expectedWithdrawal, uint256 actualWithdrawal)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseLidoWithdrawal(log types.Log) (*PufferVaultV3LidoWithdrawal, error) {
	event := new(PufferVaultV3LidoWithdrawal)
	if err := _PufferVaultV3.contract.UnpackLog(event, "LidoWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3RequestedWithdrawalsIterator is returned from FilterRequestedWithdrawals and is used to iterate over the raw logs and unpacked data for RequestedWithdrawals events raised by the PufferVaultV3 contract.
type PufferVaultV3RequestedWithdrawalsIterator struct {
	Event *PufferVaultV3RequestedWithdrawals // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3RequestedWithdrawalsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3RequestedWithdrawals)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3RequestedWithdrawals)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3RequestedWithdrawalsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3RequestedWithdrawalsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3RequestedWithdrawals represents a RequestedWithdrawals event raised by the PufferVaultV3 contract.
type PufferVaultV3RequestedWithdrawals struct {
	RequestIds []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequestedWithdrawals is a free log retrieval operation binding the contract event 0x7dac5a8ab4fe1710dfba58441ca15750a9c71877b358e90aac49fc80b293e617.
//
// Solidity: event RequestedWithdrawals(uint256[] requestIds)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterRequestedWithdrawals(opts *bind.FilterOpts) (*PufferVaultV3RequestedWithdrawalsIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "RequestedWithdrawals")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3RequestedWithdrawalsIterator{contract: _PufferVaultV3.contract, event: "RequestedWithdrawals", logs: logs, sub: sub}, nil
}

// WatchRequestedWithdrawals is a free log subscription operation binding the contract event 0x7dac5a8ab4fe1710dfba58441ca15750a9c71877b358e90aac49fc80b293e617.
//
// Solidity: event RequestedWithdrawals(uint256[] requestIds)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchRequestedWithdrawals(opts *bind.WatchOpts, sink chan<- *PufferVaultV3RequestedWithdrawals) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "RequestedWithdrawals")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3RequestedWithdrawals)
				if err := _PufferVaultV3.contract.UnpackLog(event, "RequestedWithdrawals", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestedWithdrawals is a log parse operation binding the contract event 0x7dac5a8ab4fe1710dfba58441ca15750a9c71877b358e90aac49fc80b293e617.
//
// Solidity: event RequestedWithdrawals(uint256[] requestIds)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseRequestedWithdrawals(log types.Log) (*PufferVaultV3RequestedWithdrawals, error) {
	event := new(PufferVaultV3RequestedWithdrawals)
	if err := _PufferVaultV3.contract.UnpackLog(event, "RequestedWithdrawals", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PufferVaultV3 contract.
type PufferVaultV3TransferIterator struct {
	Event *PufferVaultV3Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3Transfer represents a Transfer event raised by the PufferVaultV3 contract.
type PufferVaultV3Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PufferVaultV3TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3TransferIterator{contract: _PufferVaultV3.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PufferVaultV3Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3Transfer)
				if err := _PufferVaultV3.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseTransfer(log types.Log) (*PufferVaultV3Transfer, error) {
	event := new(PufferVaultV3Transfer)
	if err := _PufferVaultV3.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3TransferredETHIterator is returned from FilterTransferredETH and is used to iterate over the raw logs and unpacked data for TransferredETH events raised by the PufferVaultV3 contract.
type PufferVaultV3TransferredETHIterator struct {
	Event *PufferVaultV3TransferredETH // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3TransferredETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3TransferredETH)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3TransferredETH)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3TransferredETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3TransferredETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3TransferredETH represents a TransferredETH event raised by the PufferVaultV3 contract.
type PufferVaultV3TransferredETH struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferredETH is a free log retrieval operation binding the contract event 0xba7bb5aa419c34d8776b86cc0e9d41e72d74a893a511f361a11af6c05e920c3d.
//
// Solidity: event TransferredETH(address indexed to, uint256 amount)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterTransferredETH(opts *bind.FilterOpts, to []common.Address) (*PufferVaultV3TransferredETHIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "TransferredETH", toRule)
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3TransferredETHIterator{contract: _PufferVaultV3.contract, event: "TransferredETH", logs: logs, sub: sub}, nil
}

// WatchTransferredETH is a free log subscription operation binding the contract event 0xba7bb5aa419c34d8776b86cc0e9d41e72d74a893a511f361a11af6c05e920c3d.
//
// Solidity: event TransferredETH(address indexed to, uint256 amount)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchTransferredETH(opts *bind.WatchOpts, sink chan<- *PufferVaultV3TransferredETH, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "TransferredETH", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3TransferredETH)
				if err := _PufferVaultV3.contract.UnpackLog(event, "TransferredETH", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferredETH is a log parse operation binding the contract event 0xba7bb5aa419c34d8776b86cc0e9d41e72d74a893a511f361a11af6c05e920c3d.
//
// Solidity: event TransferredETH(address indexed to, uint256 amount)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseTransferredETH(log types.Log) (*PufferVaultV3TransferredETH, error) {
	event := new(PufferVaultV3TransferredETH)
	if err := _PufferVaultV3.contract.UnpackLog(event, "TransferredETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3UpdatedTotalRewardsAmountIterator is returned from FilterUpdatedTotalRewardsAmount and is used to iterate over the raw logs and unpacked data for UpdatedTotalRewardsAmount events raised by the PufferVaultV3 contract.
type PufferVaultV3UpdatedTotalRewardsAmountIterator struct {
	Event *PufferVaultV3UpdatedTotalRewardsAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3UpdatedTotalRewardsAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3UpdatedTotalRewardsAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3UpdatedTotalRewardsAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3UpdatedTotalRewardsAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3UpdatedTotalRewardsAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3UpdatedTotalRewardsAmount represents a UpdatedTotalRewardsAmount event raised by the PufferVaultV3 contract.
type PufferVaultV3UpdatedTotalRewardsAmount struct {
	PreviousTotalRewardsAmount *big.Int
	NewTotalRewardsAmount      *big.Int
	DepositedETHAmount         *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTotalRewardsAmount is a free log retrieval operation binding the contract event 0x3a278b4e83c8793751d35f41b90435c742acf0dfdd54a8cbe09aa59720db93a5.
//
// Solidity: event UpdatedTotalRewardsAmount(uint256 previousTotalRewardsAmount, uint256 newTotalRewardsAmount, uint256 depositedETHAmount)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterUpdatedTotalRewardsAmount(opts *bind.FilterOpts) (*PufferVaultV3UpdatedTotalRewardsAmountIterator, error) {

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "UpdatedTotalRewardsAmount")
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3UpdatedTotalRewardsAmountIterator{contract: _PufferVaultV3.contract, event: "UpdatedTotalRewardsAmount", logs: logs, sub: sub}, nil
}

// WatchUpdatedTotalRewardsAmount is a free log subscription operation binding the contract event 0x3a278b4e83c8793751d35f41b90435c742acf0dfdd54a8cbe09aa59720db93a5.
//
// Solidity: event UpdatedTotalRewardsAmount(uint256 previousTotalRewardsAmount, uint256 newTotalRewardsAmount, uint256 depositedETHAmount)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchUpdatedTotalRewardsAmount(opts *bind.WatchOpts, sink chan<- *PufferVaultV3UpdatedTotalRewardsAmount) (event.Subscription, error) {

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "UpdatedTotalRewardsAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3UpdatedTotalRewardsAmount)
				if err := _PufferVaultV3.contract.UnpackLog(event, "UpdatedTotalRewardsAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedTotalRewardsAmount is a log parse operation binding the contract event 0x3a278b4e83c8793751d35f41b90435c742acf0dfdd54a8cbe09aa59720db93a5.
//
// Solidity: event UpdatedTotalRewardsAmount(uint256 previousTotalRewardsAmount, uint256 newTotalRewardsAmount, uint256 depositedETHAmount)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseUpdatedTotalRewardsAmount(log types.Log) (*PufferVaultV3UpdatedTotalRewardsAmount, error) {
	event := new(PufferVaultV3UpdatedTotalRewardsAmount)
	if err := _PufferVaultV3.contract.UnpackLog(event, "UpdatedTotalRewardsAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PufferVaultV3 contract.
type PufferVaultV3UpgradedIterator struct {
	Event *PufferVaultV3Upgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3Upgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3Upgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3Upgraded represents a Upgraded event raised by the PufferVaultV3 contract.
type PufferVaultV3Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PufferVaultV3UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3UpgradedIterator{contract: _PufferVaultV3.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PufferVaultV3Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3Upgraded)
				if err := _PufferVaultV3.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseUpgraded(log types.Log) (*PufferVaultV3Upgraded, error) {
	event := new(PufferVaultV3Upgraded)
	if err := _PufferVaultV3.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PufferVaultV3WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PufferVaultV3 contract.
type PufferVaultV3WithdrawIterator struct {
	Event *PufferVaultV3Withdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PufferVaultV3WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PufferVaultV3Withdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PufferVaultV3Withdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PufferVaultV3WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PufferVaultV3WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PufferVaultV3Withdraw represents a Withdraw event raised by the PufferVaultV3 contract.
type PufferVaultV3Withdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_PufferVaultV3 *PufferVaultV3Filterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*PufferVaultV3WithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PufferVaultV3.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PufferVaultV3WithdrawIterator{contract: _PufferVaultV3.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_PufferVaultV3 *PufferVaultV3Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PufferVaultV3Withdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PufferVaultV3.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PufferVaultV3Withdraw)
				if err := _PufferVaultV3.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_PufferVaultV3 *PufferVaultV3Filterer) ParseWithdraw(log types.Log) (*PufferVaultV3Withdraw, error) {
	event := new(PufferVaultV3Withdraw)
	if err := _PufferVaultV3.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
