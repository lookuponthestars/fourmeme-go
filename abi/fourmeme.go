// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fourmeme_abi

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

// TokenManager3TokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenManager3TokenInfo struct {
	Base        common.Address
	Quote       common.Address
	Template    *big.Int
	TotalSupply *big.Int
	MaxOffers   *big.Int
	MaxRaising  *big.Int
	LaunchTime  *big.Int
	Offers      *big.Int
	Funds       *big.Int
	LastPrice   *big.Int
	K           *big.Int
	T           *big.Int
	Status      *big.Int
}

// FourMemeMetaData contains all meta data concerning the FourMeme contract.
var FourMemeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTradingFee\",\"type\":\"uint256\"}],\"name\":\"addTemplate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"origin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFunds\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFunds\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"origin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFunds\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFunds\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"origin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"buyTokenAMAP\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"buyTokenAMAP\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"buyTokenAMAP\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"origin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"buyTokenAMAP\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Max\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Max\",\"type\":\"uint128\"}],\"name\":\"collectFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"createToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referralRewardKeeper\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"launchFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"}],\"name\":\"LPLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offers\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"origin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minFunds\",\"type\":\"uint256\"}],\"name\":\"sellToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"origin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sellToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minFunds\",\"type\":\"uint256\"}],\"name\":\"sellToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sellToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setLaunchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"template\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setMinTradingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTradingFee\",\"type\":\"uint256\"}],\"name\":\"setMinTradingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setReferralRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"adminRole\",\"type\":\"bytes32\"}],\"name\":\"setRoleAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setTradingFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"suspendTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"TemplateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"launchTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"launchFee\",\"type\":\"uint256\"}],\"name\":\"TokenCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offers\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"origin\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offers\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"}],\"name\":\"TokenSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"origin\",\"type\":\"uint256\"}],\"name\":\"TokenSale2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TradeStop\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_feeCollectors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_launchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_referralRewardKeeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_referralRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_templateCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_templates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRaising\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOffers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTradingFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_tokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_tokenCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenInfoExs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"founder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"template\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOffers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRaising\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"launchTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_tradingFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_tradingHalt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"template\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOffers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRaising\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"launchTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structTokenManager3.TokenInfo\",\"name\":\"ti\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"}],\"name\":\"calcBuyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"template\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOffers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRaising\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"launchTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structTokenManager3.TokenInfo\",\"name\":\"ti\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calcBuyCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"template\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOffers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRaising\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"launchTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structTokenManager3.TokenInfo\",\"name\":\"ti\",\"type\":\"tuple\"}],\"name\":\"calcLastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"template\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOffers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRaising\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"launchTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structTokenManager3.TokenInfo\",\"name\":\"ti\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calcSellCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"template\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOffers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRaising\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"launchTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structTokenManager3.TokenInfo\",\"name\":\"ti\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"}],\"name\":\"calcTradingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLICK_FUN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDesiredAmounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_LOCKER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"nameOfRole\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PANCAKE_POSITION_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_DEPLOYER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_OPERATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STATUS_ADDING_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STATUS_COMPLETED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STATUS_HALT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STATUS_TRADING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WBNB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FourMemeABI is the input ABI used to generate the binding from.
// Deprecated: Use FourMemeMetaData.ABI instead.
var FourMemeABI = FourMemeMetaData.ABI

// FourMeme is an auto generated Go binding around an Ethereum contract.
type FourMeme struct {
	FourMemeCaller     // Read-only binding to the contract
	FourMemeTransactor // Write-only binding to the contract
	FourMemeFilterer   // Log filterer for contract events
}

// FourMemeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FourMemeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FourMemeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FourMemeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FourMemeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FourMemeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FourMemeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FourMemeSession struct {
	Contract     *FourMeme         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FourMemeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FourMemeCallerSession struct {
	Contract *FourMemeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FourMemeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FourMemeTransactorSession struct {
	Contract     *FourMemeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FourMemeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FourMemeRaw struct {
	Contract *FourMeme // Generic contract binding to access the raw methods on
}

// FourMemeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FourMemeCallerRaw struct {
	Contract *FourMemeCaller // Generic read-only contract binding to access the raw methods on
}

// FourMemeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FourMemeTransactorRaw struct {
	Contract *FourMemeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFourMeme creates a new instance of FourMeme, bound to a specific deployed contract.
func NewFourMeme(address common.Address, backend bind.ContractBackend) (*FourMeme, error) {
	contract, err := bindFourMeme(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FourMeme{FourMemeCaller: FourMemeCaller{contract: contract}, FourMemeTransactor: FourMemeTransactor{contract: contract}, FourMemeFilterer: FourMemeFilterer{contract: contract}}, nil
}

// NewFourMemeCaller creates a new read-only instance of FourMeme, bound to a specific deployed contract.
func NewFourMemeCaller(address common.Address, caller bind.ContractCaller) (*FourMemeCaller, error) {
	contract, err := bindFourMeme(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FourMemeCaller{contract: contract}, nil
}

// NewFourMemeTransactor creates a new write-only instance of FourMeme, bound to a specific deployed contract.
func NewFourMemeTransactor(address common.Address, transactor bind.ContractTransactor) (*FourMemeTransactor, error) {
	contract, err := bindFourMeme(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FourMemeTransactor{contract: contract}, nil
}

// NewFourMemeFilterer creates a new log filterer instance of FourMeme, bound to a specific deployed contract.
func NewFourMemeFilterer(address common.Address, filterer bind.ContractFilterer) (*FourMemeFilterer, error) {
	contract, err := bindFourMeme(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FourMemeFilterer{contract: contract}, nil
}

// bindFourMeme binds a generic wrapper to an already deployed contract.
func bindFourMeme(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FourMemeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FourMeme *FourMemeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FourMeme.Contract.FourMemeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FourMeme *FourMemeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FourMeme.Contract.FourMemeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FourMeme *FourMemeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FourMeme.Contract.FourMemeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FourMeme *FourMemeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FourMeme.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FourMeme *FourMemeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FourMeme.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FourMeme *FourMemeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FourMeme.Contract.contract.Transact(opts, method, params...)
}

// CLICKFUN is a free data retrieval call binding the contract method 0x0b7a8b65.
//
// Solidity: function CLICK_FUN() view returns(address)
func (_FourMeme *FourMemeCaller) CLICKFUN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "CLICK_FUN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CLICKFUN is a free data retrieval call binding the contract method 0x0b7a8b65.
//
// Solidity: function CLICK_FUN() view returns(address)
func (_FourMeme *FourMemeSession) CLICKFUN() (common.Address, error) {
	return _FourMeme.Contract.CLICKFUN(&_FourMeme.CallOpts)
}

// CLICKFUN is a free data retrieval call binding the contract method 0x0b7a8b65.
//
// Solidity: function CLICK_FUN() view returns(address)
func (_FourMeme *FourMemeCallerSession) CLICKFUN() (common.Address, error) {
	return _FourMeme.Contract.CLICKFUN(&_FourMeme.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FourMeme *FourMemeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FourMeme *FourMemeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FourMeme.Contract.DEFAULTADMINROLE(&_FourMeme.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FourMeme *FourMemeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FourMeme.Contract.DEFAULTADMINROLE(&_FourMeme.CallOpts)
}

// LPLOCKER is a free data retrieval call binding the contract method 0xcdf05fbc.
//
// Solidity: function LP_LOCKER() view returns(address)
func (_FourMeme *FourMemeCaller) LPLOCKER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "LP_LOCKER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LPLOCKER is a free data retrieval call binding the contract method 0xcdf05fbc.
//
// Solidity: function LP_LOCKER() view returns(address)
func (_FourMeme *FourMemeSession) LPLOCKER() (common.Address, error) {
	return _FourMeme.Contract.LPLOCKER(&_FourMeme.CallOpts)
}

// LPLOCKER is a free data retrieval call binding the contract method 0xcdf05fbc.
//
// Solidity: function LP_LOCKER() view returns(address)
func (_FourMeme *FourMemeCallerSession) LPLOCKER() (common.Address, error) {
	return _FourMeme.Contract.LPLOCKER(&_FourMeme.CallOpts)
}

// PANCAKEPOSITIONMANAGER is a free data retrieval call binding the contract method 0x412f7b60.
//
// Solidity: function PANCAKE_POSITION_MANAGER() view returns(address)
func (_FourMeme *FourMemeCaller) PANCAKEPOSITIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "PANCAKE_POSITION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PANCAKEPOSITIONMANAGER is a free data retrieval call binding the contract method 0x412f7b60.
//
// Solidity: function PANCAKE_POSITION_MANAGER() view returns(address)
func (_FourMeme *FourMemeSession) PANCAKEPOSITIONMANAGER() (common.Address, error) {
	return _FourMeme.Contract.PANCAKEPOSITIONMANAGER(&_FourMeme.CallOpts)
}

// PANCAKEPOSITIONMANAGER is a free data retrieval call binding the contract method 0x412f7b60.
//
// Solidity: function PANCAKE_POSITION_MANAGER() view returns(address)
func (_FourMeme *FourMemeCallerSession) PANCAKEPOSITIONMANAGER() (common.Address, error) {
	return _FourMeme.Contract.PANCAKEPOSITIONMANAGER(&_FourMeme.CallOpts)
}

// ROLEDEPLOYER is a free data retrieval call binding the contract method 0x2ee7cbad.
//
// Solidity: function ROLE_DEPLOYER() view returns(bytes32)
func (_FourMeme *FourMemeCaller) ROLEDEPLOYER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "ROLE_DEPLOYER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLEDEPLOYER is a free data retrieval call binding the contract method 0x2ee7cbad.
//
// Solidity: function ROLE_DEPLOYER() view returns(bytes32)
func (_FourMeme *FourMemeSession) ROLEDEPLOYER() ([32]byte, error) {
	return _FourMeme.Contract.ROLEDEPLOYER(&_FourMeme.CallOpts)
}

// ROLEDEPLOYER is a free data retrieval call binding the contract method 0x2ee7cbad.
//
// Solidity: function ROLE_DEPLOYER() view returns(bytes32)
func (_FourMeme *FourMemeCallerSession) ROLEDEPLOYER() ([32]byte, error) {
	return _FourMeme.Contract.ROLEDEPLOYER(&_FourMeme.CallOpts)
}

// ROLEOPERATOR is a free data retrieval call binding the contract method 0x98a1b397.
//
// Solidity: function ROLE_OPERATOR() view returns(bytes32)
func (_FourMeme *FourMemeCaller) ROLEOPERATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "ROLE_OPERATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLEOPERATOR is a free data retrieval call binding the contract method 0x98a1b397.
//
// Solidity: function ROLE_OPERATOR() view returns(bytes32)
func (_FourMeme *FourMemeSession) ROLEOPERATOR() ([32]byte, error) {
	return _FourMeme.Contract.ROLEOPERATOR(&_FourMeme.CallOpts)
}

// ROLEOPERATOR is a free data retrieval call binding the contract method 0x98a1b397.
//
// Solidity: function ROLE_OPERATOR() view returns(bytes32)
func (_FourMeme *FourMemeCallerSession) ROLEOPERATOR() ([32]byte, error) {
	return _FourMeme.Contract.ROLEOPERATOR(&_FourMeme.CallOpts)
}

// STATUSADDINGLIQUIDITY is a free data retrieval call binding the contract method 0xe63e65d8.
//
// Solidity: function STATUS_ADDING_LIQUIDITY() view returns(uint256)
func (_FourMeme *FourMemeCaller) STATUSADDINGLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "STATUS_ADDING_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STATUSADDINGLIQUIDITY is a free data retrieval call binding the contract method 0xe63e65d8.
//
// Solidity: function STATUS_ADDING_LIQUIDITY() view returns(uint256)
func (_FourMeme *FourMemeSession) STATUSADDINGLIQUIDITY() (*big.Int, error) {
	return _FourMeme.Contract.STATUSADDINGLIQUIDITY(&_FourMeme.CallOpts)
}

// STATUSADDINGLIQUIDITY is a free data retrieval call binding the contract method 0xe63e65d8.
//
// Solidity: function STATUS_ADDING_LIQUIDITY() view returns(uint256)
func (_FourMeme *FourMemeCallerSession) STATUSADDINGLIQUIDITY() (*big.Int, error) {
	return _FourMeme.Contract.STATUSADDINGLIQUIDITY(&_FourMeme.CallOpts)
}

// STATUSCOMPLETED is a free data retrieval call binding the contract method 0x724f27ee.
//
// Solidity: function STATUS_COMPLETED() view returns(uint256)
func (_FourMeme *FourMemeCaller) STATUSCOMPLETED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "STATUS_COMPLETED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STATUSCOMPLETED is a free data retrieval call binding the contract method 0x724f27ee.
//
// Solidity: function STATUS_COMPLETED() view returns(uint256)
func (_FourMeme *FourMemeSession) STATUSCOMPLETED() (*big.Int, error) {
	return _FourMeme.Contract.STATUSCOMPLETED(&_FourMeme.CallOpts)
}

// STATUSCOMPLETED is a free data retrieval call binding the contract method 0x724f27ee.
//
// Solidity: function STATUS_COMPLETED() view returns(uint256)
func (_FourMeme *FourMemeCallerSession) STATUSCOMPLETED() (*big.Int, error) {
	return _FourMeme.Contract.STATUSCOMPLETED(&_FourMeme.CallOpts)
}

// STATUSHALT is a free data retrieval call binding the contract method 0xb7836d11.
//
// Solidity: function STATUS_HALT() view returns(uint256)
func (_FourMeme *FourMemeCaller) STATUSHALT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "STATUS_HALT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STATUSHALT is a free data retrieval call binding the contract method 0xb7836d11.
//
// Solidity: function STATUS_HALT() view returns(uint256)
func (_FourMeme *FourMemeSession) STATUSHALT() (*big.Int, error) {
	return _FourMeme.Contract.STATUSHALT(&_FourMeme.CallOpts)
}

// STATUSHALT is a free data retrieval call binding the contract method 0xb7836d11.
//
// Solidity: function STATUS_HALT() view returns(uint256)
func (_FourMeme *FourMemeCallerSession) STATUSHALT() (*big.Int, error) {
	return _FourMeme.Contract.STATUSHALT(&_FourMeme.CallOpts)
}

// STATUSTRADING is a free data retrieval call binding the contract method 0x0a4e8445.
//
// Solidity: function STATUS_TRADING() view returns(uint256)
func (_FourMeme *FourMemeCaller) STATUSTRADING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "STATUS_TRADING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STATUSTRADING is a free data retrieval call binding the contract method 0x0a4e8445.
//
// Solidity: function STATUS_TRADING() view returns(uint256)
func (_FourMeme *FourMemeSession) STATUSTRADING() (*big.Int, error) {
	return _FourMeme.Contract.STATUSTRADING(&_FourMeme.CallOpts)
}

// STATUSTRADING is a free data retrieval call binding the contract method 0x0a4e8445.
//
// Solidity: function STATUS_TRADING() view returns(uint256)
func (_FourMeme *FourMemeCallerSession) STATUSTRADING() (*big.Int, error) {
	return _FourMeme.Contract.STATUSTRADING(&_FourMeme.CallOpts)
}

// WBNB is a free data retrieval call binding the contract method 0x8dd95002.
//
// Solidity: function WBNB() view returns(address)
func (_FourMeme *FourMemeCaller) WBNB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "WBNB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WBNB is a free data retrieval call binding the contract method 0x8dd95002.
//
// Solidity: function WBNB() view returns(address)
func (_FourMeme *FourMemeSession) WBNB() (common.Address, error) {
	return _FourMeme.Contract.WBNB(&_FourMeme.CallOpts)
}

// WBNB is a free data retrieval call binding the contract method 0x8dd95002.
//
// Solidity: function WBNB() view returns(address)
func (_FourMeme *FourMemeCallerSession) WBNB() (common.Address, error) {
	return _FourMeme.Contract.WBNB(&_FourMeme.CallOpts)
}

// FeeCollectors is a free data retrieval call binding the contract method 0x53415191.
//
// Solidity: function _feeCollectors(address ) view returns(bool)
func (_FourMeme *FourMemeCaller) FeeCollectors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_feeCollectors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FeeCollectors is a free data retrieval call binding the contract method 0x53415191.
//
// Solidity: function _feeCollectors(address ) view returns(bool)
func (_FourMeme *FourMemeSession) FeeCollectors(arg0 common.Address) (bool, error) {
	return _FourMeme.Contract.FeeCollectors(&_FourMeme.CallOpts, arg0)
}

// FeeCollectors is a free data retrieval call binding the contract method 0x53415191.
//
// Solidity: function _feeCollectors(address ) view returns(bool)
func (_FourMeme *FourMemeCallerSession) FeeCollectors(arg0 common.Address) (bool, error) {
	return _FourMeme.Contract.FeeCollectors(&_FourMeme.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0xa342f238.
//
// Solidity: function _feeRecipient() view returns(address)
func (_FourMeme *FourMemeCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0xa342f238.
//
// Solidity: function _feeRecipient() view returns(address)
func (_FourMeme *FourMemeSession) FeeRecipient() (common.Address, error) {
	return _FourMeme.Contract.FeeRecipient(&_FourMeme.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0xa342f238.
//
// Solidity: function _feeRecipient() view returns(address)
func (_FourMeme *FourMemeCallerSession) FeeRecipient() (common.Address, error) {
	return _FourMeme.Contract.FeeRecipient(&_FourMeme.CallOpts)
}

// LaunchFee is a free data retrieval call binding the contract method 0x009523a2.
//
// Solidity: function _launchFee() view returns(uint256)
func (_FourMeme *FourMemeCaller) LaunchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_launchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaunchFee is a free data retrieval call binding the contract method 0x009523a2.
//
// Solidity: function _launchFee() view returns(uint256)
func (_FourMeme *FourMemeSession) LaunchFee() (*big.Int, error) {
	return _FourMeme.Contract.LaunchFee(&_FourMeme.CallOpts)
}

// LaunchFee is a free data retrieval call binding the contract method 0x009523a2.
//
// Solidity: function _launchFee() view returns(uint256)
func (_FourMeme *FourMemeCallerSession) LaunchFee() (*big.Int, error) {
	return _FourMeme.Contract.LaunchFee(&_FourMeme.CallOpts)
}

// ReferralRewardKeeper is a free data retrieval call binding the contract method 0x38c2e3ac.
//
// Solidity: function _referralRewardKeeper() view returns(address)
func (_FourMeme *FourMemeCaller) ReferralRewardKeeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_referralRewardKeeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReferralRewardKeeper is a free data retrieval call binding the contract method 0x38c2e3ac.
//
// Solidity: function _referralRewardKeeper() view returns(address)
func (_FourMeme *FourMemeSession) ReferralRewardKeeper() (common.Address, error) {
	return _FourMeme.Contract.ReferralRewardKeeper(&_FourMeme.CallOpts)
}

// ReferralRewardKeeper is a free data retrieval call binding the contract method 0x38c2e3ac.
//
// Solidity: function _referralRewardKeeper() view returns(address)
func (_FourMeme *FourMemeCallerSession) ReferralRewardKeeper() (common.Address, error) {
	return _FourMeme.Contract.ReferralRewardKeeper(&_FourMeme.CallOpts)
}

// ReferralRewardRate is a free data retrieval call binding the contract method 0x40ca6a69.
//
// Solidity: function _referralRewardRate() view returns(uint256)
func (_FourMeme *FourMemeCaller) ReferralRewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_referralRewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferralRewardRate is a free data retrieval call binding the contract method 0x40ca6a69.
//
// Solidity: function _referralRewardRate() view returns(uint256)
func (_FourMeme *FourMemeSession) ReferralRewardRate() (*big.Int, error) {
	return _FourMeme.Contract.ReferralRewardRate(&_FourMeme.CallOpts)
}

// ReferralRewardRate is a free data retrieval call binding the contract method 0x40ca6a69.
//
// Solidity: function _referralRewardRate() view returns(uint256)
func (_FourMeme *FourMemeCallerSession) ReferralRewardRate() (*big.Int, error) {
	return _FourMeme.Contract.ReferralRewardRate(&_FourMeme.CallOpts)
}

// TemplateCount is a free data retrieval call binding the contract method 0xe2176efe.
//
// Solidity: function _templateCount() view returns(uint256)
func (_FourMeme *FourMemeCaller) TemplateCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_templateCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TemplateCount is a free data retrieval call binding the contract method 0xe2176efe.
//
// Solidity: function _templateCount() view returns(uint256)
func (_FourMeme *FourMemeSession) TemplateCount() (*big.Int, error) {
	return _FourMeme.Contract.TemplateCount(&_FourMeme.CallOpts)
}

// TemplateCount is a free data retrieval call binding the contract method 0xe2176efe.
//
// Solidity: function _templateCount() view returns(uint256)
func (_FourMeme *FourMemeCallerSession) TemplateCount() (*big.Int, error) {
	return _FourMeme.Contract.TemplateCount(&_FourMeme.CallOpts)
}

// Templates is a free data retrieval call binding the contract method 0x4ed7fdcd.
//
// Solidity: function _templates(uint256 ) view returns(address quote, uint256 initialLiquidity, uint256 maxRaising, uint256 totalSupply, uint256 maxOffers, uint256 minTradingFee)
func (_FourMeme *FourMemeCaller) Templates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Quote            common.Address
	InitialLiquidity *big.Int
	MaxRaising       *big.Int
	TotalSupply      *big.Int
	MaxOffers        *big.Int
	MinTradingFee    *big.Int
}, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_templates", arg0)

	outstruct := new(struct {
		Quote            common.Address
		InitialLiquidity *big.Int
		MaxRaising       *big.Int
		TotalSupply      *big.Int
		MaxOffers        *big.Int
		MinTradingFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quote = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.InitialLiquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxRaising = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxOffers = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MinTradingFee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Templates is a free data retrieval call binding the contract method 0x4ed7fdcd.
//
// Solidity: function _templates(uint256 ) view returns(address quote, uint256 initialLiquidity, uint256 maxRaising, uint256 totalSupply, uint256 maxOffers, uint256 minTradingFee)
func (_FourMeme *FourMemeSession) Templates(arg0 *big.Int) (struct {
	Quote            common.Address
	InitialLiquidity *big.Int
	MaxRaising       *big.Int
	TotalSupply      *big.Int
	MaxOffers        *big.Int
	MinTradingFee    *big.Int
}, error) {
	return _FourMeme.Contract.Templates(&_FourMeme.CallOpts, arg0)
}

// Templates is a free data retrieval call binding the contract method 0x4ed7fdcd.
//
// Solidity: function _templates(uint256 ) view returns(address quote, uint256 initialLiquidity, uint256 maxRaising, uint256 totalSupply, uint256 maxOffers, uint256 minTradingFee)
func (_FourMeme *FourMemeCallerSession) Templates(arg0 *big.Int) (struct {
	Quote            common.Address
	InitialLiquidity *big.Int
	MaxRaising       *big.Int
	TotalSupply      *big.Int
	MaxOffers        *big.Int
	MinTradingFee    *big.Int
}, error) {
	return _FourMeme.Contract.Templates(&_FourMeme.CallOpts, arg0)
}

// TokenCount is a free data retrieval call binding the contract method 0x1eef9d2c.
//
// Solidity: function _tokenCount() view returns(uint256)
func (_FourMeme *FourMemeCaller) TokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_tokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCount is a free data retrieval call binding the contract method 0x1eef9d2c.
//
// Solidity: function _tokenCount() view returns(uint256)
func (_FourMeme *FourMemeSession) TokenCount() (*big.Int, error) {
	return _FourMeme.Contract.TokenCount(&_FourMeme.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x1eef9d2c.
//
// Solidity: function _tokenCount() view returns(uint256)
func (_FourMeme *FourMemeCallerSession) TokenCount() (*big.Int, error) {
	return _FourMeme.Contract.TokenCount(&_FourMeme.CallOpts)
}

// TokenCreator is a free data retrieval call binding the contract method 0x892e73ef.
//
// Solidity: function _tokenCreator() view returns(address)
func (_FourMeme *FourMemeCaller) TokenCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_tokenCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenCreator is a free data retrieval call binding the contract method 0x892e73ef.
//
// Solidity: function _tokenCreator() view returns(address)
func (_FourMeme *FourMemeSession) TokenCreator() (common.Address, error) {
	return _FourMeme.Contract.TokenCreator(&_FourMeme.CallOpts)
}

// TokenCreator is a free data retrieval call binding the contract method 0x892e73ef.
//
// Solidity: function _tokenCreator() view returns(address)
func (_FourMeme *FourMemeCallerSession) TokenCreator() (common.Address, error) {
	return _FourMeme.Contract.TokenCreator(&_FourMeme.CallOpts)
}

// TokenInfoExs is a free data retrieval call binding the contract method 0xe91b2f3b.
//
// Solidity: function _tokenInfoExs(address ) view returns(address creator, address founder, uint256 reserves)
func (_FourMeme *FourMemeCaller) TokenInfoExs(opts *bind.CallOpts, arg0 common.Address) (struct {
	Creator  common.Address
	Founder  common.Address
	Reserves *big.Int
}, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_tokenInfoExs", arg0)

	outstruct := new(struct {
		Creator  common.Address
		Founder  common.Address
		Reserves *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Creator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Founder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Reserves = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenInfoExs is a free data retrieval call binding the contract method 0xe91b2f3b.
//
// Solidity: function _tokenInfoExs(address ) view returns(address creator, address founder, uint256 reserves)
func (_FourMeme *FourMemeSession) TokenInfoExs(arg0 common.Address) (struct {
	Creator  common.Address
	Founder  common.Address
	Reserves *big.Int
}, error) {
	return _FourMeme.Contract.TokenInfoExs(&_FourMeme.CallOpts, arg0)
}

// TokenInfoExs is a free data retrieval call binding the contract method 0xe91b2f3b.
//
// Solidity: function _tokenInfoExs(address ) view returns(address creator, address founder, uint256 reserves)
func (_FourMeme *FourMemeCallerSession) TokenInfoExs(arg0 common.Address) (struct {
	Creator  common.Address
	Founder  common.Address
	Reserves *big.Int
}, error) {
	return _FourMeme.Contract.TokenInfoExs(&_FourMeme.CallOpts, arg0)
}

// TokenInfos is a free data retrieval call binding the contract method 0xe684626b.
//
// Solidity: function _tokenInfos(address ) view returns(address base, address quote, uint256 template, uint256 totalSupply, uint256 maxOffers, uint256 maxRaising, uint256 launchTime, uint256 offers, uint256 funds, uint256 lastPrice, uint256 K, uint256 T, uint256 status)
func (_FourMeme *FourMemeCaller) TokenInfos(opts *bind.CallOpts, arg0 common.Address) (struct {
	Base        common.Address
	Quote       common.Address
	Template    *big.Int
	TotalSupply *big.Int
	MaxOffers   *big.Int
	MaxRaising  *big.Int
	LaunchTime  *big.Int
	Offers      *big.Int
	Funds       *big.Int
	LastPrice   *big.Int
	K           *big.Int
	T           *big.Int
	Status      *big.Int
}, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_tokenInfos", arg0)

	outstruct := new(struct {
		Base        common.Address
		Quote       common.Address
		Template    *big.Int
		TotalSupply *big.Int
		MaxOffers   *big.Int
		MaxRaising  *big.Int
		LaunchTime  *big.Int
		Offers      *big.Int
		Funds       *big.Int
		LastPrice   *big.Int
		K           *big.Int
		T           *big.Int
		Status      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Base = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Quote = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Template = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxOffers = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MaxRaising = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LaunchTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Offers = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Funds = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LastPrice = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.K = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.T = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenInfos is a free data retrieval call binding the contract method 0xe684626b.
//
// Solidity: function _tokenInfos(address ) view returns(address base, address quote, uint256 template, uint256 totalSupply, uint256 maxOffers, uint256 maxRaising, uint256 launchTime, uint256 offers, uint256 funds, uint256 lastPrice, uint256 K, uint256 T, uint256 status)
func (_FourMeme *FourMemeSession) TokenInfos(arg0 common.Address) (struct {
	Base        common.Address
	Quote       common.Address
	Template    *big.Int
	TotalSupply *big.Int
	MaxOffers   *big.Int
	MaxRaising  *big.Int
	LaunchTime  *big.Int
	Offers      *big.Int
	Funds       *big.Int
	LastPrice   *big.Int
	K           *big.Int
	T           *big.Int
	Status      *big.Int
}, error) {
	return _FourMeme.Contract.TokenInfos(&_FourMeme.CallOpts, arg0)
}

// TokenInfos is a free data retrieval call binding the contract method 0xe684626b.
//
// Solidity: function _tokenInfos(address ) view returns(address base, address quote, uint256 template, uint256 totalSupply, uint256 maxOffers, uint256 maxRaising, uint256 launchTime, uint256 offers, uint256 funds, uint256 lastPrice, uint256 K, uint256 T, uint256 status)
func (_FourMeme *FourMemeCallerSession) TokenInfos(arg0 common.Address) (struct {
	Base        common.Address
	Quote       common.Address
	Template    *big.Int
	TotalSupply *big.Int
	MaxOffers   *big.Int
	MaxRaising  *big.Int
	LaunchTime  *big.Int
	Offers      *big.Int
	Funds       *big.Int
	LastPrice   *big.Int
	K           *big.Int
	T           *big.Int
	Status      *big.Int
}, error) {
	return _FourMeme.Contract.TokenInfos(&_FourMeme.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xed70554d.
//
// Solidity: function _tokens(uint256 ) view returns(address)
func (_FourMeme *FourMemeCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0xed70554d.
//
// Solidity: function _tokens(uint256 ) view returns(address)
func (_FourMeme *FourMemeSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _FourMeme.Contract.Tokens(&_FourMeme.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xed70554d.
//
// Solidity: function _tokens(uint256 ) view returns(address)
func (_FourMeme *FourMemeCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _FourMeme.Contract.Tokens(&_FourMeme.CallOpts, arg0)
}

// TradingFeeRate is a free data retrieval call binding the contract method 0x3472aee7.
//
// Solidity: function _tradingFeeRate() view returns(uint256)
func (_FourMeme *FourMemeCaller) TradingFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_tradingFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TradingFeeRate is a free data retrieval call binding the contract method 0x3472aee7.
//
// Solidity: function _tradingFeeRate() view returns(uint256)
func (_FourMeme *FourMemeSession) TradingFeeRate() (*big.Int, error) {
	return _FourMeme.Contract.TradingFeeRate(&_FourMeme.CallOpts)
}

// TradingFeeRate is a free data retrieval call binding the contract method 0x3472aee7.
//
// Solidity: function _tradingFeeRate() view returns(uint256)
func (_FourMeme *FourMemeCallerSession) TradingFeeRate() (*big.Int, error) {
	return _FourMeme.Contract.TradingFeeRate(&_FourMeme.CallOpts)
}

// TradingHalt is a free data retrieval call binding the contract method 0x088c5d0b.
//
// Solidity: function _tradingHalt() view returns(bool)
func (_FourMeme *FourMemeCaller) TradingHalt(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "_tradingHalt")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TradingHalt is a free data retrieval call binding the contract method 0x088c5d0b.
//
// Solidity: function _tradingHalt() view returns(bool)
func (_FourMeme *FourMemeSession) TradingHalt() (bool, error) {
	return _FourMeme.Contract.TradingHalt(&_FourMeme.CallOpts)
}

// TradingHalt is a free data retrieval call binding the contract method 0x088c5d0b.
//
// Solidity: function _tradingHalt() view returns(bool)
func (_FourMeme *FourMemeCallerSession) TradingHalt() (bool, error) {
	return _FourMeme.Contract.TradingHalt(&_FourMeme.CallOpts)
}

// CalcBuyAmount is a free data retrieval call binding the contract method 0x22810b19.
//
// Solidity: function calcBuyAmount((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 funds) pure returns(uint256)
func (_FourMeme *FourMemeCaller) CalcBuyAmount(opts *bind.CallOpts, ti TokenManager3TokenInfo, funds *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "calcBuyAmount", ti, funds)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcBuyAmount is a free data retrieval call binding the contract method 0x22810b19.
//
// Solidity: function calcBuyAmount((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 funds) pure returns(uint256)
func (_FourMeme *FourMemeSession) CalcBuyAmount(ti TokenManager3TokenInfo, funds *big.Int) (*big.Int, error) {
	return _FourMeme.Contract.CalcBuyAmount(&_FourMeme.CallOpts, ti, funds)
}

// CalcBuyAmount is a free data retrieval call binding the contract method 0x22810b19.
//
// Solidity: function calcBuyAmount((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 funds) pure returns(uint256)
func (_FourMeme *FourMemeCallerSession) CalcBuyAmount(ti TokenManager3TokenInfo, funds *big.Int) (*big.Int, error) {
	return _FourMeme.Contract.CalcBuyAmount(&_FourMeme.CallOpts, ti, funds)
}

// CalcBuyCost is a free data retrieval call binding the contract method 0x467dffe6.
//
// Solidity: function calcBuyCost((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 amount) pure returns(uint256)
func (_FourMeme *FourMemeCaller) CalcBuyCost(opts *bind.CallOpts, ti TokenManager3TokenInfo, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "calcBuyCost", ti, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcBuyCost is a free data retrieval call binding the contract method 0x467dffe6.
//
// Solidity: function calcBuyCost((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 amount) pure returns(uint256)
func (_FourMeme *FourMemeSession) CalcBuyCost(ti TokenManager3TokenInfo, amount *big.Int) (*big.Int, error) {
	return _FourMeme.Contract.CalcBuyCost(&_FourMeme.CallOpts, ti, amount)
}

// CalcBuyCost is a free data retrieval call binding the contract method 0x467dffe6.
//
// Solidity: function calcBuyCost((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 amount) pure returns(uint256)
func (_FourMeme *FourMemeCallerSession) CalcBuyCost(ti TokenManager3TokenInfo, amount *big.Int) (*big.Int, error) {
	return _FourMeme.Contract.CalcBuyCost(&_FourMeme.CallOpts, ti, amount)
}

// CalcLastPrice is a free data retrieval call binding the contract method 0xe87bc493.
//
// Solidity: function calcLastPrice((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti) pure returns(uint256)
func (_FourMeme *FourMemeCaller) CalcLastPrice(opts *bind.CallOpts, ti TokenManager3TokenInfo) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "calcLastPrice", ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcLastPrice is a free data retrieval call binding the contract method 0xe87bc493.
//
// Solidity: function calcLastPrice((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti) pure returns(uint256)
func (_FourMeme *FourMemeSession) CalcLastPrice(ti TokenManager3TokenInfo) (*big.Int, error) {
	return _FourMeme.Contract.CalcLastPrice(&_FourMeme.CallOpts, ti)
}

// CalcLastPrice is a free data retrieval call binding the contract method 0xe87bc493.
//
// Solidity: function calcLastPrice((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti) pure returns(uint256)
func (_FourMeme *FourMemeCallerSession) CalcLastPrice(ti TokenManager3TokenInfo) (*big.Int, error) {
	return _FourMeme.Contract.CalcLastPrice(&_FourMeme.CallOpts, ti)
}

// CalcSellCost is a free data retrieval call binding the contract method 0xcf141ecd.
//
// Solidity: function calcSellCost((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 amount) pure returns(uint256)
func (_FourMeme *FourMemeCaller) CalcSellCost(opts *bind.CallOpts, ti TokenManager3TokenInfo, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "calcSellCost", ti, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSellCost is a free data retrieval call binding the contract method 0xcf141ecd.
//
// Solidity: function calcSellCost((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 amount) pure returns(uint256)
func (_FourMeme *FourMemeSession) CalcSellCost(ti TokenManager3TokenInfo, amount *big.Int) (*big.Int, error) {
	return _FourMeme.Contract.CalcSellCost(&_FourMeme.CallOpts, ti, amount)
}

// CalcSellCost is a free data retrieval call binding the contract method 0xcf141ecd.
//
// Solidity: function calcSellCost((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 amount) pure returns(uint256)
func (_FourMeme *FourMemeCallerSession) CalcSellCost(ti TokenManager3TokenInfo, amount *big.Int) (*big.Int, error) {
	return _FourMeme.Contract.CalcSellCost(&_FourMeme.CallOpts, ti, amount)
}

// CalcTradingFee is a free data retrieval call binding the contract method 0x5cd616d3.
//
// Solidity: function calcTradingFee((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 funds) view returns(uint256)
func (_FourMeme *FourMemeCaller) CalcTradingFee(opts *bind.CallOpts, ti TokenManager3TokenInfo, funds *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "calcTradingFee", ti, funds)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTradingFee is a free data retrieval call binding the contract method 0x5cd616d3.
//
// Solidity: function calcTradingFee((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 funds) view returns(uint256)
func (_FourMeme *FourMemeSession) CalcTradingFee(ti TokenManager3TokenInfo, funds *big.Int) (*big.Int, error) {
	return _FourMeme.Contract.CalcTradingFee(&_FourMeme.CallOpts, ti, funds)
}

// CalcTradingFee is a free data retrieval call binding the contract method 0x5cd616d3.
//
// Solidity: function calcTradingFee((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) ti, uint256 funds) view returns(uint256)
func (_FourMeme *FourMemeCallerSession) CalcTradingFee(ti TokenManager3TokenInfo, funds *big.Int) (*big.Int, error) {
	return _FourMeme.Contract.CalcTradingFee(&_FourMeme.CallOpts, ti, funds)
}

// GetDesiredAmounts is a free data retrieval call binding the contract method 0xd3db6cd1.
//
// Solidity: function getDesiredAmounts(address token) view returns(address token0, address token1, uint256 amount0, uint256 amount1)
func (_FourMeme *FourMemeCaller) GetDesiredAmounts(opts *bind.CallOpts, token common.Address) (struct {
	Token0  common.Address
	Token1  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "getDesiredAmounts", token)

	outstruct := new(struct {
		Token0  common.Address
		Token1  common.Address
		Amount0 *big.Int
		Amount1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDesiredAmounts is a free data retrieval call binding the contract method 0xd3db6cd1.
//
// Solidity: function getDesiredAmounts(address token) view returns(address token0, address token1, uint256 amount0, uint256 amount1)
func (_FourMeme *FourMemeSession) GetDesiredAmounts(token common.Address) (struct {
	Token0  common.Address
	Token1  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _FourMeme.Contract.GetDesiredAmounts(&_FourMeme.CallOpts, token)
}

// GetDesiredAmounts is a free data retrieval call binding the contract method 0xd3db6cd1.
//
// Solidity: function getDesiredAmounts(address token) view returns(address token0, address token1, uint256 amount0, uint256 amount1)
func (_FourMeme *FourMemeCallerSession) GetDesiredAmounts(token common.Address) (struct {
	Token0  common.Address
	Token1  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _FourMeme.Contract.GetDesiredAmounts(&_FourMeme.CallOpts, token)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FourMeme *FourMemeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FourMeme *FourMemeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FourMeme.Contract.GetRoleAdmin(&_FourMeme.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FourMeme *FourMemeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FourMeme.Contract.GetRoleAdmin(&_FourMeme.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FourMeme *FourMemeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FourMeme *FourMemeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FourMeme.Contract.HasRole(&_FourMeme.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FourMeme *FourMemeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FourMeme.Contract.HasRole(&_FourMeme.CallOpts, role, account)
}

// NameOfRole is a free data retrieval call binding the contract method 0x92e00ada.
//
// Solidity: function nameOfRole(bytes32 role) pure returns(string)
func (_FourMeme *FourMemeCaller) NameOfRole(opts *bind.CallOpts, role [32]byte) (string, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "nameOfRole", role)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NameOfRole is a free data retrieval call binding the contract method 0x92e00ada.
//
// Solidity: function nameOfRole(bytes32 role) pure returns(string)
func (_FourMeme *FourMemeSession) NameOfRole(role [32]byte) (string, error) {
	return _FourMeme.Contract.NameOfRole(&_FourMeme.CallOpts, role)
}

// NameOfRole is a free data retrieval call binding the contract method 0x92e00ada.
//
// Solidity: function nameOfRole(bytes32 role) pure returns(string)
func (_FourMeme *FourMemeCallerSession) NameOfRole(role [32]byte) (string, error) {
	return _FourMeme.Contract.NameOfRole(&_FourMeme.CallOpts, role)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FourMeme *FourMemeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FourMeme *FourMemeSession) Owner() (common.Address, error) {
	return _FourMeme.Contract.Owner(&_FourMeme.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FourMeme *FourMemeCallerSession) Owner() (common.Address, error) {
	return _FourMeme.Contract.Owner(&_FourMeme.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FourMeme *FourMemeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FourMeme *FourMemeSession) ProxiableUUID() ([32]byte, error) {
	return _FourMeme.Contract.ProxiableUUID(&_FourMeme.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FourMeme *FourMemeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _FourMeme.Contract.ProxiableUUID(&_FourMeme.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FourMeme *FourMemeCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FourMeme *FourMemeSession) Signer() (common.Address, error) {
	return _FourMeme.Contract.Signer(&_FourMeme.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FourMeme *FourMemeCallerSession) Signer() (common.Address, error) {
	return _FourMeme.Contract.Signer(&_FourMeme.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FourMeme *FourMemeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FourMeme.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FourMeme *FourMemeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FourMeme.Contract.SupportsInterface(&_FourMeme.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FourMeme *FourMemeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FourMeme.Contract.SupportsInterface(&_FourMeme.CallOpts, interfaceId)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x55110640.
//
// Solidity: function addLiquidity(address tokenAddress, uint160 sqrtPriceX96) returns()
func (_FourMeme *FourMemeTransactor) AddLiquidity(opts *bind.TransactOpts, tokenAddress common.Address, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "addLiquidity", tokenAddress, sqrtPriceX96)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x55110640.
//
// Solidity: function addLiquidity(address tokenAddress, uint160 sqrtPriceX96) returns()
func (_FourMeme *FourMemeSession) AddLiquidity(tokenAddress common.Address, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.AddLiquidity(&_FourMeme.TransactOpts, tokenAddress, sqrtPriceX96)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x55110640.
//
// Solidity: function addLiquidity(address tokenAddress, uint160 sqrtPriceX96) returns()
func (_FourMeme *FourMemeTransactorSession) AddLiquidity(tokenAddress common.Address, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.AddLiquidity(&_FourMeme.TransactOpts, tokenAddress, sqrtPriceX96)
}

// AddTemplate is a paid mutator transaction binding the contract method 0xc59e2598.
//
// Solidity: function addTemplate(address quote, uint256 minTradingFee) returns()
func (_FourMeme *FourMemeTransactor) AddTemplate(opts *bind.TransactOpts, quote common.Address, minTradingFee *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "addTemplate", quote, minTradingFee)
}

// AddTemplate is a paid mutator transaction binding the contract method 0xc59e2598.
//
// Solidity: function addTemplate(address quote, uint256 minTradingFee) returns()
func (_FourMeme *FourMemeSession) AddTemplate(quote common.Address, minTradingFee *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.AddTemplate(&_FourMeme.TransactOpts, quote, minTradingFee)
}

// AddTemplate is a paid mutator transaction binding the contract method 0xc59e2598.
//
// Solidity: function addTemplate(address quote, uint256 minTradingFee) returns()
func (_FourMeme *FourMemeTransactorSession) AddTemplate(quote common.Address, minTradingFee *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.AddTemplate(&_FourMeme.TransactOpts, quote, minTradingFee)
}

// BuyToken is a paid mutator transaction binding the contract method 0x43fc3523.
//
// Solidity: function buyToken(uint256 origin, address token, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeTransactor) BuyToken(opts *bind.TransactOpts, origin *big.Int, token common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "buyToken", origin, token, amount, maxFunds)
}

// BuyToken is a paid mutator transaction binding the contract method 0x43fc3523.
//
// Solidity: function buyToken(uint256 origin, address token, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeSession) BuyToken(origin *big.Int, token common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyToken(&_FourMeme.TransactOpts, origin, token, amount, maxFunds)
}

// BuyToken is a paid mutator transaction binding the contract method 0x43fc3523.
//
// Solidity: function buyToken(uint256 origin, address token, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeTransactorSession) BuyToken(origin *big.Int, token common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyToken(&_FourMeme.TransactOpts, origin, token, amount, maxFunds)
}

// BuyToken0 is a paid mutator transaction binding the contract method 0x7d17ff3d.
//
// Solidity: function buyToken(address token, address to, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeTransactor) BuyToken0(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "buyToken0", token, to, amount, maxFunds)
}

// BuyToken0 is a paid mutator transaction binding the contract method 0x7d17ff3d.
//
// Solidity: function buyToken(address token, address to, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeSession) BuyToken0(token common.Address, to common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyToken0(&_FourMeme.TransactOpts, token, to, amount, maxFunds)
}

// BuyToken0 is a paid mutator transaction binding the contract method 0x7d17ff3d.
//
// Solidity: function buyToken(address token, address to, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeTransactorSession) BuyToken0(token common.Address, to common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyToken0(&_FourMeme.TransactOpts, token, to, amount, maxFunds)
}

// BuyToken1 is a paid mutator transaction binding the contract method 0x90702094.
//
// Solidity: function buyToken(uint256 origin, address token, address to, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeTransactor) BuyToken1(opts *bind.TransactOpts, origin *big.Int, token common.Address, to common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "buyToken1", origin, token, to, amount, maxFunds)
}

// BuyToken1 is a paid mutator transaction binding the contract method 0x90702094.
//
// Solidity: function buyToken(uint256 origin, address token, address to, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeSession) BuyToken1(origin *big.Int, token common.Address, to common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyToken1(&_FourMeme.TransactOpts, origin, token, to, amount, maxFunds)
}

// BuyToken1 is a paid mutator transaction binding the contract method 0x90702094.
//
// Solidity: function buyToken(uint256 origin, address token, address to, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeTransactorSession) BuyToken1(origin *big.Int, token common.Address, to common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyToken1(&_FourMeme.TransactOpts, origin, token, to, amount, maxFunds)
}

// BuyToken2 is a paid mutator transaction binding the contract method 0xe671499b.
//
// Solidity: function buyToken(address token, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeTransactor) BuyToken2(opts *bind.TransactOpts, token common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "buyToken2", token, amount, maxFunds)
}

// BuyToken2 is a paid mutator transaction binding the contract method 0xe671499b.
//
// Solidity: function buyToken(address token, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeSession) BuyToken2(token common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyToken2(&_FourMeme.TransactOpts, token, amount, maxFunds)
}

// BuyToken2 is a paid mutator transaction binding the contract method 0xe671499b.
//
// Solidity: function buyToken(address token, uint256 amount, uint256 maxFunds) payable returns()
func (_FourMeme *FourMemeTransactorSession) BuyToken2(token common.Address, amount *big.Int, maxFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyToken2(&_FourMeme.TransactOpts, token, amount, maxFunds)
}

// BuyTokenAMAP is a paid mutator transaction binding the contract method 0x5e56c39a.
//
// Solidity: function buyTokenAMAP(uint256 origin, address token, address to, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeTransactor) BuyTokenAMAP(opts *bind.TransactOpts, origin *big.Int, token common.Address, to common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "buyTokenAMAP", origin, token, to, funds, minAmount)
}

// BuyTokenAMAP is a paid mutator transaction binding the contract method 0x5e56c39a.
//
// Solidity: function buyTokenAMAP(uint256 origin, address token, address to, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeSession) BuyTokenAMAP(origin *big.Int, token common.Address, to common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyTokenAMAP(&_FourMeme.TransactOpts, origin, token, to, funds, minAmount)
}

// BuyTokenAMAP is a paid mutator transaction binding the contract method 0x5e56c39a.
//
// Solidity: function buyTokenAMAP(uint256 origin, address token, address to, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeTransactorSession) BuyTokenAMAP(origin *big.Int, token common.Address, to common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyTokenAMAP(&_FourMeme.TransactOpts, origin, token, to, funds, minAmount)
}

// BuyTokenAMAP0 is a paid mutator transaction binding the contract method 0x7f79f6df.
//
// Solidity: function buyTokenAMAP(address token, address to, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeTransactor) BuyTokenAMAP0(opts *bind.TransactOpts, token common.Address, to common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "buyTokenAMAP0", token, to, funds, minAmount)
}

// BuyTokenAMAP0 is a paid mutator transaction binding the contract method 0x7f79f6df.
//
// Solidity: function buyTokenAMAP(address token, address to, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeSession) BuyTokenAMAP0(token common.Address, to common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyTokenAMAP0(&_FourMeme.TransactOpts, token, to, funds, minAmount)
}

// BuyTokenAMAP0 is a paid mutator transaction binding the contract method 0x7f79f6df.
//
// Solidity: function buyTokenAMAP(address token, address to, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeTransactorSession) BuyTokenAMAP0(token common.Address, to common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyTokenAMAP0(&_FourMeme.TransactOpts, token, to, funds, minAmount)
}

// BuyTokenAMAP1 is a paid mutator transaction binding the contract method 0x87f27655.
//
// Solidity: function buyTokenAMAP(address token, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeTransactor) BuyTokenAMAP1(opts *bind.TransactOpts, token common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "buyTokenAMAP1", token, funds, minAmount)
}

// BuyTokenAMAP1 is a paid mutator transaction binding the contract method 0x87f27655.
//
// Solidity: function buyTokenAMAP(address token, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeSession) BuyTokenAMAP1(token common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyTokenAMAP1(&_FourMeme.TransactOpts, token, funds, minAmount)
}

// BuyTokenAMAP1 is a paid mutator transaction binding the contract method 0x87f27655.
//
// Solidity: function buyTokenAMAP(address token, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeTransactorSession) BuyTokenAMAP1(token common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyTokenAMAP1(&_FourMeme.TransactOpts, token, funds, minAmount)
}

// BuyTokenAMAP2 is a paid mutator transaction binding the contract method 0xedf9e251.
//
// Solidity: function buyTokenAMAP(uint256 origin, address token, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeTransactor) BuyTokenAMAP2(opts *bind.TransactOpts, origin *big.Int, token common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "buyTokenAMAP2", origin, token, funds, minAmount)
}

// BuyTokenAMAP2 is a paid mutator transaction binding the contract method 0xedf9e251.
//
// Solidity: function buyTokenAMAP(uint256 origin, address token, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeSession) BuyTokenAMAP2(origin *big.Int, token common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyTokenAMAP2(&_FourMeme.TransactOpts, origin, token, funds, minAmount)
}

// BuyTokenAMAP2 is a paid mutator transaction binding the contract method 0xedf9e251.
//
// Solidity: function buyTokenAMAP(uint256 origin, address token, uint256 funds, uint256 minAmount) payable returns()
func (_FourMeme *FourMemeTransactorSession) BuyTokenAMAP2(origin *big.Int, token common.Address, funds *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.BuyTokenAMAP2(&_FourMeme.TransactOpts, origin, token, funds, minAmount)
}

// CollectFees is a paid mutator transaction binding the contract method 0x59f3f952.
//
// Solidity: function collectFees(uint256 lockId, address recipient, uint128 amount0Max, uint128 amount1Max) returns(uint256 amount0, uint256 amount1, uint256 fee0, uint256 fee1)
func (_FourMeme *FourMemeTransactor) CollectFees(opts *bind.TransactOpts, lockId *big.Int, recipient common.Address, amount0Max *big.Int, amount1Max *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "collectFees", lockId, recipient, amount0Max, amount1Max)
}

// CollectFees is a paid mutator transaction binding the contract method 0x59f3f952.
//
// Solidity: function collectFees(uint256 lockId, address recipient, uint128 amount0Max, uint128 amount1Max) returns(uint256 amount0, uint256 amount1, uint256 fee0, uint256 fee1)
func (_FourMeme *FourMemeSession) CollectFees(lockId *big.Int, recipient common.Address, amount0Max *big.Int, amount1Max *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.CollectFees(&_FourMeme.TransactOpts, lockId, recipient, amount0Max, amount1Max)
}

// CollectFees is a paid mutator transaction binding the contract method 0x59f3f952.
//
// Solidity: function collectFees(uint256 lockId, address recipient, uint128 amount0Max, uint128 amount1Max) returns(uint256 amount0, uint256 amount1, uint256 fee0, uint256 fee1)
func (_FourMeme *FourMemeTransactorSession) CollectFees(lockId *big.Int, recipient common.Address, amount0Max *big.Int, amount1Max *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.CollectFees(&_FourMeme.TransactOpts, lockId, recipient, amount0Max, amount1Max)
}

// CreateToken is a paid mutator transaction binding the contract method 0x519ebb10.
//
// Solidity: function createToken(bytes args, bytes signature) payable returns()
func (_FourMeme *FourMemeTransactor) CreateToken(opts *bind.TransactOpts, args []byte, signature []byte) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "createToken", args, signature)
}

// CreateToken is a paid mutator transaction binding the contract method 0x519ebb10.
//
// Solidity: function createToken(bytes args, bytes signature) payable returns()
func (_FourMeme *FourMemeSession) CreateToken(args []byte, signature []byte) (*types.Transaction, error) {
	return _FourMeme.Contract.CreateToken(&_FourMeme.TransactOpts, args, signature)
}

// CreateToken is a paid mutator transaction binding the contract method 0x519ebb10.
//
// Solidity: function createToken(bytes args, bytes signature) payable returns()
func (_FourMeme *FourMemeTransactorSession) CreateToken(args []byte, signature []byte) (*types.Transaction, error) {
	return _FourMeme.Contract.CreateToken(&_FourMeme.TransactOpts, args, signature)
}

// GrantDeployer is a paid mutator transaction binding the contract method 0xd83722f0.
//
// Solidity: function grantDeployer(address account) returns()
func (_FourMeme *FourMemeTransactor) GrantDeployer(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "grantDeployer", account)
}

// GrantDeployer is a paid mutator transaction binding the contract method 0xd83722f0.
//
// Solidity: function grantDeployer(address account) returns()
func (_FourMeme *FourMemeSession) GrantDeployer(account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.GrantDeployer(&_FourMeme.TransactOpts, account)
}

// GrantDeployer is a paid mutator transaction binding the contract method 0xd83722f0.
//
// Solidity: function grantDeployer(address account) returns()
func (_FourMeme *FourMemeTransactorSession) GrantDeployer(account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.GrantDeployer(&_FourMeme.TransactOpts, account)
}

// GrantOperator is a paid mutator transaction binding the contract method 0xe348da13.
//
// Solidity: function grantOperator(address account) returns()
func (_FourMeme *FourMemeTransactor) GrantOperator(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "grantOperator", account)
}

// GrantOperator is a paid mutator transaction binding the contract method 0xe348da13.
//
// Solidity: function grantOperator(address account) returns()
func (_FourMeme *FourMemeSession) GrantOperator(account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.GrantOperator(&_FourMeme.TransactOpts, account)
}

// GrantOperator is a paid mutator transaction binding the contract method 0xe348da13.
//
// Solidity: function grantOperator(address account) returns()
func (_FourMeme *FourMemeTransactorSession) GrantOperator(account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.GrantOperator(&_FourMeme.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FourMeme *FourMemeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FourMeme *FourMemeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.GrantRole(&_FourMeme.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FourMeme *FourMemeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.GrantRole(&_FourMeme.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FourMeme *FourMemeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FourMeme *FourMemeSession) Initialize() (*types.Transaction, error) {
	return _FourMeme.Contract.Initialize(&_FourMeme.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FourMeme *FourMemeTransactorSession) Initialize() (*types.Transaction, error) {
	return _FourMeme.Contract.Initialize(&_FourMeme.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address signer, address feeRecipient, address tokenCreator, address referralRewardKeeper, uint256 launchFee) returns()
func (_FourMeme *FourMemeTransactor) Initialize0(opts *bind.TransactOpts, signer common.Address, feeRecipient common.Address, tokenCreator common.Address, referralRewardKeeper common.Address, launchFee *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "initialize0", signer, feeRecipient, tokenCreator, referralRewardKeeper, launchFee)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address signer, address feeRecipient, address tokenCreator, address referralRewardKeeper, uint256 launchFee) returns()
func (_FourMeme *FourMemeSession) Initialize0(signer common.Address, feeRecipient common.Address, tokenCreator common.Address, referralRewardKeeper common.Address, launchFee *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.Initialize0(&_FourMeme.TransactOpts, signer, feeRecipient, tokenCreator, referralRewardKeeper, launchFee)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address signer, address feeRecipient, address tokenCreator, address referralRewardKeeper, uint256 launchFee) returns()
func (_FourMeme *FourMemeTransactorSession) Initialize0(signer common.Address, feeRecipient common.Address, tokenCreator common.Address, referralRewardKeeper common.Address, launchFee *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.Initialize0(&_FourMeme.TransactOpts, signer, feeRecipient, tokenCreator, referralRewardKeeper, launchFee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FourMeme *FourMemeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FourMeme *FourMemeSession) RenounceOwnership() (*types.Transaction, error) {
	return _FourMeme.Contract.RenounceOwnership(&_FourMeme.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FourMeme *FourMemeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FourMeme.Contract.RenounceOwnership(&_FourMeme.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FourMeme *FourMemeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FourMeme *FourMemeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.RenounceRole(&_FourMeme.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FourMeme *FourMemeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.RenounceRole(&_FourMeme.TransactOpts, role, account)
}

// RevokeDeployer is a paid mutator transaction binding the contract method 0x2e4934dd.
//
// Solidity: function revokeDeployer(address account) returns()
func (_FourMeme *FourMemeTransactor) RevokeDeployer(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "revokeDeployer", account)
}

// RevokeDeployer is a paid mutator transaction binding the contract method 0x2e4934dd.
//
// Solidity: function revokeDeployer(address account) returns()
func (_FourMeme *FourMemeSession) RevokeDeployer(account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.RevokeDeployer(&_FourMeme.TransactOpts, account)
}

// RevokeDeployer is a paid mutator transaction binding the contract method 0x2e4934dd.
//
// Solidity: function revokeDeployer(address account) returns()
func (_FourMeme *FourMemeTransactorSession) RevokeDeployer(account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.RevokeDeployer(&_FourMeme.TransactOpts, account)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address account) returns()
func (_FourMeme *FourMemeTransactor) RevokeOperator(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "revokeOperator", account)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address account) returns()
func (_FourMeme *FourMemeSession) RevokeOperator(account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.RevokeOperator(&_FourMeme.TransactOpts, account)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address account) returns()
func (_FourMeme *FourMemeTransactorSession) RevokeOperator(account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.RevokeOperator(&_FourMeme.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FourMeme *FourMemeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FourMeme *FourMemeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.RevokeRole(&_FourMeme.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FourMeme *FourMemeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.RevokeRole(&_FourMeme.TransactOpts, role, account)
}

// SellToken is a paid mutator transaction binding the contract method 0x0da74935.
//
// Solidity: function sellToken(uint256 origin, address token, uint256 amount, uint256 minFunds) returns()
func (_FourMeme *FourMemeTransactor) SellToken(opts *bind.TransactOpts, origin *big.Int, token common.Address, amount *big.Int, minFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "sellToken", origin, token, amount, minFunds)
}

// SellToken is a paid mutator transaction binding the contract method 0x0da74935.
//
// Solidity: function sellToken(uint256 origin, address token, uint256 amount, uint256 minFunds) returns()
func (_FourMeme *FourMemeSession) SellToken(origin *big.Int, token common.Address, amount *big.Int, minFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SellToken(&_FourMeme.TransactOpts, origin, token, amount, minFunds)
}

// SellToken is a paid mutator transaction binding the contract method 0x0da74935.
//
// Solidity: function sellToken(uint256 origin, address token, uint256 amount, uint256 minFunds) returns()
func (_FourMeme *FourMemeTransactorSession) SellToken(origin *big.Int, token common.Address, amount *big.Int, minFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SellToken(&_FourMeme.TransactOpts, origin, token, amount, minFunds)
}

// SellToken0 is a paid mutator transaction binding the contract method 0x15a68a10.
//
// Solidity: function sellToken(uint256 origin, address token, uint256 amount) returns()
func (_FourMeme *FourMemeTransactor) SellToken0(opts *bind.TransactOpts, origin *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "sellToken0", origin, token, amount)
}

// SellToken0 is a paid mutator transaction binding the contract method 0x15a68a10.
//
// Solidity: function sellToken(uint256 origin, address token, uint256 amount) returns()
func (_FourMeme *FourMemeSession) SellToken0(origin *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SellToken0(&_FourMeme.TransactOpts, origin, token, amount)
}

// SellToken0 is a paid mutator transaction binding the contract method 0x15a68a10.
//
// Solidity: function sellToken(uint256 origin, address token, uint256 amount) returns()
func (_FourMeme *FourMemeTransactorSession) SellToken0(origin *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SellToken0(&_FourMeme.TransactOpts, origin, token, amount)
}

// SellToken1 is a paid mutator transaction binding the contract method 0x3e11741f.
//
// Solidity: function sellToken(address token, uint256 amount, uint256 minFunds) returns()
func (_FourMeme *FourMemeTransactor) SellToken1(opts *bind.TransactOpts, token common.Address, amount *big.Int, minFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "sellToken1", token, amount, minFunds)
}

// SellToken1 is a paid mutator transaction binding the contract method 0x3e11741f.
//
// Solidity: function sellToken(address token, uint256 amount, uint256 minFunds) returns()
func (_FourMeme *FourMemeSession) SellToken1(token common.Address, amount *big.Int, minFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SellToken1(&_FourMeme.TransactOpts, token, amount, minFunds)
}

// SellToken1 is a paid mutator transaction binding the contract method 0x3e11741f.
//
// Solidity: function sellToken(address token, uint256 amount, uint256 minFunds) returns()
func (_FourMeme *FourMemeTransactorSession) SellToken1(token common.Address, amount *big.Int, minFunds *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SellToken1(&_FourMeme.TransactOpts, token, amount, minFunds)
}

// SellToken2 is a paid mutator transaction binding the contract method 0xf464e7db.
//
// Solidity: function sellToken(address token, uint256 amount) returns()
func (_FourMeme *FourMemeTransactor) SellToken2(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "sellToken2", token, amount)
}

// SellToken2 is a paid mutator transaction binding the contract method 0xf464e7db.
//
// Solidity: function sellToken(address token, uint256 amount) returns()
func (_FourMeme *FourMemeSession) SellToken2(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SellToken2(&_FourMeme.TransactOpts, token, amount)
}

// SellToken2 is a paid mutator transaction binding the contract method 0xf464e7db.
//
// Solidity: function sellToken(address token, uint256 amount) returns()
func (_FourMeme *FourMemeTransactorSession) SellToken2(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SellToken2(&_FourMeme.TransactOpts, token, amount)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x9a4a3b90.
//
// Solidity: function setFeeCollector(address account, bool enabled) returns()
func (_FourMeme *FourMemeTransactor) SetFeeCollector(opts *bind.TransactOpts, account common.Address, enabled bool) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "setFeeCollector", account, enabled)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x9a4a3b90.
//
// Solidity: function setFeeCollector(address account, bool enabled) returns()
func (_FourMeme *FourMemeSession) SetFeeCollector(account common.Address, enabled bool) (*types.Transaction, error) {
	return _FourMeme.Contract.SetFeeCollector(&_FourMeme.TransactOpts, account, enabled)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x9a4a3b90.
//
// Solidity: function setFeeCollector(address account, bool enabled) returns()
func (_FourMeme *FourMemeTransactorSession) SetFeeCollector(account common.Address, enabled bool) (*types.Transaction, error) {
	return _FourMeme.Contract.SetFeeCollector(&_FourMeme.TransactOpts, account, enabled)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address v) returns()
func (_FourMeme *FourMemeTransactor) SetFeeRecipient(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "setFeeRecipient", v)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address v) returns()
func (_FourMeme *FourMemeSession) SetFeeRecipient(v common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.SetFeeRecipient(&_FourMeme.TransactOpts, v)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address v) returns()
func (_FourMeme *FourMemeTransactorSession) SetFeeRecipient(v common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.SetFeeRecipient(&_FourMeme.TransactOpts, v)
}

// SetLaunchFee is a paid mutator transaction binding the contract method 0x5313be2c.
//
// Solidity: function setLaunchFee(uint256 v) returns()
func (_FourMeme *FourMemeTransactor) SetLaunchFee(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "setLaunchFee", v)
}

// SetLaunchFee is a paid mutator transaction binding the contract method 0x5313be2c.
//
// Solidity: function setLaunchFee(uint256 v) returns()
func (_FourMeme *FourMemeSession) SetLaunchFee(v *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SetLaunchFee(&_FourMeme.TransactOpts, v)
}

// SetLaunchFee is a paid mutator transaction binding the contract method 0x5313be2c.
//
// Solidity: function setLaunchFee(uint256 v) returns()
func (_FourMeme *FourMemeTransactorSession) SetLaunchFee(v *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SetLaunchFee(&_FourMeme.TransactOpts, v)
}

// SetMinTradingFee is a paid mutator transaction binding the contract method 0x244af72d.
//
// Solidity: function setMinTradingFee(uint256 template, uint256 v) returns()
func (_FourMeme *FourMemeTransactor) SetMinTradingFee(opts *bind.TransactOpts, template *big.Int, v *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "setMinTradingFee", template, v)
}

// SetMinTradingFee is a paid mutator transaction binding the contract method 0x244af72d.
//
// Solidity: function setMinTradingFee(uint256 template, uint256 v) returns()
func (_FourMeme *FourMemeSession) SetMinTradingFee(template *big.Int, v *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SetMinTradingFee(&_FourMeme.TransactOpts, template, v)
}

// SetMinTradingFee is a paid mutator transaction binding the contract method 0x244af72d.
//
// Solidity: function setMinTradingFee(uint256 template, uint256 v) returns()
func (_FourMeme *FourMemeTransactorSession) SetMinTradingFee(template *big.Int, v *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SetMinTradingFee(&_FourMeme.TransactOpts, template, v)
}

// SetMinTradingFee0 is a paid mutator transaction binding the contract method 0x925b6d80.
//
// Solidity: function setMinTradingFee(address quote, uint256 minTradingFee) returns()
func (_FourMeme *FourMemeTransactor) SetMinTradingFee0(opts *bind.TransactOpts, quote common.Address, minTradingFee *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "setMinTradingFee0", quote, minTradingFee)
}

// SetMinTradingFee0 is a paid mutator transaction binding the contract method 0x925b6d80.
//
// Solidity: function setMinTradingFee(address quote, uint256 minTradingFee) returns()
func (_FourMeme *FourMemeSession) SetMinTradingFee0(quote common.Address, minTradingFee *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SetMinTradingFee0(&_FourMeme.TransactOpts, quote, minTradingFee)
}

// SetMinTradingFee0 is a paid mutator transaction binding the contract method 0x925b6d80.
//
// Solidity: function setMinTradingFee(address quote, uint256 minTradingFee) returns()
func (_FourMeme *FourMemeTransactorSession) SetMinTradingFee0(quote common.Address, minTradingFee *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SetMinTradingFee0(&_FourMeme.TransactOpts, quote, minTradingFee)
}

// SetReferralRewardRate is a paid mutator transaction binding the contract method 0x616fc7ad.
//
// Solidity: function setReferralRewardRate(uint256 v) returns()
func (_FourMeme *FourMemeTransactor) SetReferralRewardRate(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "setReferralRewardRate", v)
}

// SetReferralRewardRate is a paid mutator transaction binding the contract method 0x616fc7ad.
//
// Solidity: function setReferralRewardRate(uint256 v) returns()
func (_FourMeme *FourMemeSession) SetReferralRewardRate(v *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SetReferralRewardRate(&_FourMeme.TransactOpts, v)
}

// SetReferralRewardRate is a paid mutator transaction binding the contract method 0x616fc7ad.
//
// Solidity: function setReferralRewardRate(uint256 v) returns()
func (_FourMeme *FourMemeTransactorSession) SetReferralRewardRate(v *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SetReferralRewardRate(&_FourMeme.TransactOpts, v)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0x1e4e0091.
//
// Solidity: function setRoleAdmin(bytes32 role, bytes32 adminRole) returns()
func (_FourMeme *FourMemeTransactor) SetRoleAdmin(opts *bind.TransactOpts, role [32]byte, adminRole [32]byte) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "setRoleAdmin", role, adminRole)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0x1e4e0091.
//
// Solidity: function setRoleAdmin(bytes32 role, bytes32 adminRole) returns()
func (_FourMeme *FourMemeSession) SetRoleAdmin(role [32]byte, adminRole [32]byte) (*types.Transaction, error) {
	return _FourMeme.Contract.SetRoleAdmin(&_FourMeme.TransactOpts, role, adminRole)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0x1e4e0091.
//
// Solidity: function setRoleAdmin(bytes32 role, bytes32 adminRole) returns()
func (_FourMeme *FourMemeTransactorSession) SetRoleAdmin(role [32]byte, adminRole [32]byte) (*types.Transaction, error) {
	return _FourMeme.Contract.SetRoleAdmin(&_FourMeme.TransactOpts, role, adminRole)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address newSigner) returns()
func (_FourMeme *FourMemeTransactor) SetSigner(opts *bind.TransactOpts, newSigner common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "setSigner", newSigner)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address newSigner) returns()
func (_FourMeme *FourMemeSession) SetSigner(newSigner common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.SetSigner(&_FourMeme.TransactOpts, newSigner)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address newSigner) returns()
func (_FourMeme *FourMemeTransactorSession) SetSigner(newSigner common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.SetSigner(&_FourMeme.TransactOpts, newSigner)
}

// SetTradingFeeRate is a paid mutator transaction binding the contract method 0x1091f67c.
//
// Solidity: function setTradingFeeRate(uint256 v) returns()
func (_FourMeme *FourMemeTransactor) SetTradingFeeRate(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "setTradingFeeRate", v)
}

// SetTradingFeeRate is a paid mutator transaction binding the contract method 0x1091f67c.
//
// Solidity: function setTradingFeeRate(uint256 v) returns()
func (_FourMeme *FourMemeSession) SetTradingFeeRate(v *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SetTradingFeeRate(&_FourMeme.TransactOpts, v)
}

// SetTradingFeeRate is a paid mutator transaction binding the contract method 0x1091f67c.
//
// Solidity: function setTradingFeeRate(uint256 v) returns()
func (_FourMeme *FourMemeTransactorSession) SetTradingFeeRate(v *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.SetTradingFeeRate(&_FourMeme.TransactOpts, v)
}

// SuspendTrading is a paid mutator transaction binding the contract method 0x7ad4c7ab.
//
// Solidity: function suspendTrading(address token, bool value) returns()
func (_FourMeme *FourMemeTransactor) SuspendTrading(opts *bind.TransactOpts, token common.Address, value bool) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "suspendTrading", token, value)
}

// SuspendTrading is a paid mutator transaction binding the contract method 0x7ad4c7ab.
//
// Solidity: function suspendTrading(address token, bool value) returns()
func (_FourMeme *FourMemeSession) SuspendTrading(token common.Address, value bool) (*types.Transaction, error) {
	return _FourMeme.Contract.SuspendTrading(&_FourMeme.TransactOpts, token, value)
}

// SuspendTrading is a paid mutator transaction binding the contract method 0x7ad4c7ab.
//
// Solidity: function suspendTrading(address token, bool value) returns()
func (_FourMeme *FourMemeTransactorSession) SuspendTrading(token common.Address, value bool) (*types.Transaction, error) {
	return _FourMeme.Contract.SuspendTrading(&_FourMeme.TransactOpts, token, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FourMeme *FourMemeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FourMeme *FourMemeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.TransferOwnership(&_FourMeme.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FourMeme *FourMemeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.TransferOwnership(&_FourMeme.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FourMeme *FourMemeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FourMeme *FourMemeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.UpgradeTo(&_FourMeme.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FourMeme *FourMemeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FourMeme.Contract.UpgradeTo(&_FourMeme.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FourMeme *FourMemeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FourMeme *FourMemeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FourMeme.Contract.UpgradeToAndCall(&_FourMeme.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FourMeme *FourMemeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FourMeme.Contract.UpgradeToAndCall(&_FourMeme.TransactOpts, newImplementation, data)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount) returns()
func (_FourMeme *FourMemeTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "withdrawERC20", token, to, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount) returns()
func (_FourMeme *FourMemeSession) WithdrawERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.WithdrawERC20(&_FourMeme.TransactOpts, token, to, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount) returns()
func (_FourMeme *FourMemeTransactorSession) WithdrawERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.WithdrawERC20(&_FourMeme.TransactOpts, token, to, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address to, uint256 amount) returns()
func (_FourMeme *FourMemeTransactor) WithdrawEth(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.contract.Transact(opts, "withdrawEth", to, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address to, uint256 amount) returns()
func (_FourMeme *FourMemeSession) WithdrawEth(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.WithdrawEth(&_FourMeme.TransactOpts, to, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address to, uint256 amount) returns()
func (_FourMeme *FourMemeTransactorSession) WithdrawEth(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FourMeme.Contract.WithdrawEth(&_FourMeme.TransactOpts, to, amount)
}

// FourMemeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the FourMeme contract.
type FourMemeAdminChangedIterator struct {
	Event *FourMemeAdminChanged // Event containing the contract specifics and raw log

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
func (it *FourMemeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeAdminChanged)
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
		it.Event = new(FourMemeAdminChanged)
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
func (it *FourMemeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeAdminChanged represents a AdminChanged event raised by the FourMeme contract.
type FourMemeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FourMeme *FourMemeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*FourMemeAdminChangedIterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &FourMemeAdminChangedIterator{contract: _FourMeme.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FourMeme *FourMemeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *FourMemeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeAdminChanged)
				if err := _FourMeme.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FourMeme *FourMemeFilterer) ParseAdminChanged(log types.Log) (*FourMemeAdminChanged, error) {
	event := new(FourMemeAdminChanged)
	if err := _FourMeme.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the FourMeme contract.
type FourMemeBeaconUpgradedIterator struct {
	Event *FourMemeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *FourMemeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeBeaconUpgraded)
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
		it.Event = new(FourMemeBeaconUpgraded)
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
func (it *FourMemeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeBeaconUpgraded represents a BeaconUpgraded event raised by the FourMeme contract.
type FourMemeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FourMeme *FourMemeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*FourMemeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &FourMemeBeaconUpgradedIterator{contract: _FourMeme.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FourMeme *FourMemeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *FourMemeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeBeaconUpgraded)
				if err := _FourMeme.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FourMeme *FourMemeFilterer) ParseBeaconUpgraded(log types.Log) (*FourMemeBeaconUpgraded, error) {
	event := new(FourMemeBeaconUpgraded)
	if err := _FourMeme.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FourMeme contract.
type FourMemeInitializedIterator struct {
	Event *FourMemeInitialized // Event containing the contract specifics and raw log

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
func (it *FourMemeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeInitialized)
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
		it.Event = new(FourMemeInitialized)
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
func (it *FourMemeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeInitialized represents a Initialized event raised by the FourMeme contract.
type FourMemeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FourMeme *FourMemeFilterer) FilterInitialized(opts *bind.FilterOpts) (*FourMemeInitializedIterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FourMemeInitializedIterator{contract: _FourMeme.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FourMeme *FourMemeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FourMemeInitialized) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeInitialized)
				if err := _FourMeme.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FourMeme *FourMemeFilterer) ParseInitialized(log types.Log) (*FourMemeInitialized, error) {
	event := new(FourMemeInitialized)
	if err := _FourMeme.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeLPLockIterator is returned from FilterLPLock and is used to iterate over the raw logs and unpacked data for LPLock events raised by the FourMeme contract.
type FourMemeLPLockIterator struct {
	Event *FourMemeLPLock // Event containing the contract specifics and raw log

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
func (it *FourMemeLPLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeLPLock)
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
		it.Event = new(FourMemeLPLock)
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
func (it *FourMemeLPLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeLPLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeLPLock represents a LPLock event raised by the FourMeme contract.
type FourMemeLPLock struct {
	Token  common.Address
	NftId  *big.Int
	LockId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLPLock is a free log retrieval operation binding the contract event 0xa78d55aeb92a87db782ede05df51f62cd9c43f9c4ee844147e54d963cd30d37a.
//
// Solidity: event LPLock(address token, uint256 nftId, uint256 lockId)
func (_FourMeme *FourMemeFilterer) FilterLPLock(opts *bind.FilterOpts) (*FourMemeLPLockIterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "LPLock")
	if err != nil {
		return nil, err
	}
	return &FourMemeLPLockIterator{contract: _FourMeme.contract, event: "LPLock", logs: logs, sub: sub}, nil
}

// WatchLPLock is a free log subscription operation binding the contract event 0xa78d55aeb92a87db782ede05df51f62cd9c43f9c4ee844147e54d963cd30d37a.
//
// Solidity: event LPLock(address token, uint256 nftId, uint256 lockId)
func (_FourMeme *FourMemeFilterer) WatchLPLock(opts *bind.WatchOpts, sink chan<- *FourMemeLPLock) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "LPLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeLPLock)
				if err := _FourMeme.contract.UnpackLog(event, "LPLock", log); err != nil {
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

// ParseLPLock is a log parse operation binding the contract event 0xa78d55aeb92a87db782ede05df51f62cd9c43f9c4ee844147e54d963cd30d37a.
//
// Solidity: event LPLock(address token, uint256 nftId, uint256 lockId)
func (_FourMeme *FourMemeFilterer) ParseLPLock(log types.Log) (*FourMemeLPLock, error) {
	event := new(FourMemeLPLock)
	if err := _FourMeme.contract.UnpackLog(event, "LPLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the FourMeme contract.
type FourMemeLiquidityAddedIterator struct {
	Event *FourMemeLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *FourMemeLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeLiquidityAdded)
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
		it.Event = new(FourMemeLiquidityAdded)
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
func (it *FourMemeLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeLiquidityAdded represents a LiquidityAdded event raised by the FourMeme contract.
type FourMemeLiquidityAdded struct {
	Base   common.Address
	Offers *big.Int
	Quote  common.Address
	Funds  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0xc18aa71171b358b706fe3dd345299685ba21a5316c66ffa9e319268b033c44b0.
//
// Solidity: event LiquidityAdded(address base, uint256 offers, address quote, uint256 funds)
func (_FourMeme *FourMemeFilterer) FilterLiquidityAdded(opts *bind.FilterOpts) (*FourMemeLiquidityAddedIterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return &FourMemeLiquidityAddedIterator{contract: _FourMeme.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0xc18aa71171b358b706fe3dd345299685ba21a5316c66ffa9e319268b033c44b0.
//
// Solidity: event LiquidityAdded(address base, uint256 offers, address quote, uint256 funds)
func (_FourMeme *FourMemeFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *FourMemeLiquidityAdded) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeLiquidityAdded)
				if err := _FourMeme.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0xc18aa71171b358b706fe3dd345299685ba21a5316c66ffa9e319268b033c44b0.
//
// Solidity: event LiquidityAdded(address base, uint256 offers, address quote, uint256 funds)
func (_FourMeme *FourMemeFilterer) ParseLiquidityAdded(log types.Log) (*FourMemeLiquidityAdded, error) {
	event := new(FourMemeLiquidityAdded)
	if err := _FourMeme.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FourMeme contract.
type FourMemeOwnershipTransferredIterator struct {
	Event *FourMemeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FourMemeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeOwnershipTransferred)
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
		it.Event = new(FourMemeOwnershipTransferred)
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
func (it *FourMemeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeOwnershipTransferred represents a OwnershipTransferred event raised by the FourMeme contract.
type FourMemeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FourMeme *FourMemeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FourMemeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FourMemeOwnershipTransferredIterator{contract: _FourMeme.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FourMeme *FourMemeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FourMemeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeOwnershipTransferred)
				if err := _FourMeme.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FourMeme *FourMemeFilterer) ParseOwnershipTransferred(log types.Log) (*FourMemeOwnershipTransferred, error) {
	event := new(FourMemeOwnershipTransferred)
	if err := _FourMeme.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FourMeme contract.
type FourMemeRoleAdminChangedIterator struct {
	Event *FourMemeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FourMemeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeRoleAdminChanged)
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
		it.Event = new(FourMemeRoleAdminChanged)
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
func (it *FourMemeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeRoleAdminChanged represents a RoleAdminChanged event raised by the FourMeme contract.
type FourMemeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FourMeme *FourMemeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FourMemeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FourMemeRoleAdminChangedIterator{contract: _FourMeme.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FourMeme *FourMemeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FourMemeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeRoleAdminChanged)
				if err := _FourMeme.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FourMeme *FourMemeFilterer) ParseRoleAdminChanged(log types.Log) (*FourMemeRoleAdminChanged, error) {
	event := new(FourMemeRoleAdminChanged)
	if err := _FourMeme.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FourMeme contract.
type FourMemeRoleGrantedIterator struct {
	Event *FourMemeRoleGranted // Event containing the contract specifics and raw log

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
func (it *FourMemeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeRoleGranted)
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
		it.Event = new(FourMemeRoleGranted)
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
func (it *FourMemeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeRoleGranted represents a RoleGranted event raised by the FourMeme contract.
type FourMemeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FourMeme *FourMemeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FourMemeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FourMemeRoleGrantedIterator{contract: _FourMeme.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FourMeme *FourMemeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FourMemeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeRoleGranted)
				if err := _FourMeme.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FourMeme *FourMemeFilterer) ParseRoleGranted(log types.Log) (*FourMemeRoleGranted, error) {
	event := new(FourMemeRoleGranted)
	if err := _FourMeme.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FourMeme contract.
type FourMemeRoleRevokedIterator struct {
	Event *FourMemeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FourMemeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeRoleRevoked)
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
		it.Event = new(FourMemeRoleRevoked)
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
func (it *FourMemeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeRoleRevoked represents a RoleRevoked event raised by the FourMeme contract.
type FourMemeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FourMeme *FourMemeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FourMemeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FourMemeRoleRevokedIterator{contract: _FourMeme.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FourMeme *FourMemeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FourMemeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeRoleRevoked)
				if err := _FourMeme.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FourMeme *FourMemeFilterer) ParseRoleRevoked(log types.Log) (*FourMemeRoleRevoked, error) {
	event := new(FourMemeRoleRevoked)
	if err := _FourMeme.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeTemplateAddedIterator is returned from FilterTemplateAdded and is used to iterate over the raw logs and unpacked data for TemplateAdded events raised by the FourMeme contract.
type FourMemeTemplateAddedIterator struct {
	Event *FourMemeTemplateAdded // Event containing the contract specifics and raw log

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
func (it *FourMemeTemplateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeTemplateAdded)
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
		it.Event = new(FourMemeTemplateAdded)
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
func (it *FourMemeTemplateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeTemplateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeTemplateAdded represents a TemplateAdded event raised by the FourMeme contract.
type FourMemeTemplateAdded struct {
	Quote common.Address
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTemplateAdded is a free log retrieval operation binding the contract event 0x70300ab95e22eebeda3bb99011e4ddcd4e3012b97e60bf47c6f4d547dbb8d47a.
//
// Solidity: event TemplateAdded(address quote, uint256 index)
func (_FourMeme *FourMemeFilterer) FilterTemplateAdded(opts *bind.FilterOpts) (*FourMemeTemplateAddedIterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "TemplateAdded")
	if err != nil {
		return nil, err
	}
	return &FourMemeTemplateAddedIterator{contract: _FourMeme.contract, event: "TemplateAdded", logs: logs, sub: sub}, nil
}

// WatchTemplateAdded is a free log subscription operation binding the contract event 0x70300ab95e22eebeda3bb99011e4ddcd4e3012b97e60bf47c6f4d547dbb8d47a.
//
// Solidity: event TemplateAdded(address quote, uint256 index)
func (_FourMeme *FourMemeFilterer) WatchTemplateAdded(opts *bind.WatchOpts, sink chan<- *FourMemeTemplateAdded) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "TemplateAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeTemplateAdded)
				if err := _FourMeme.contract.UnpackLog(event, "TemplateAdded", log); err != nil {
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

// ParseTemplateAdded is a log parse operation binding the contract event 0x70300ab95e22eebeda3bb99011e4ddcd4e3012b97e60bf47c6f4d547dbb8d47a.
//
// Solidity: event TemplateAdded(address quote, uint256 index)
func (_FourMeme *FourMemeFilterer) ParseTemplateAdded(log types.Log) (*FourMemeTemplateAdded, error) {
	event := new(FourMemeTemplateAdded)
	if err := _FourMeme.contract.UnpackLog(event, "TemplateAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeTokenCreateIterator is returned from FilterTokenCreate and is used to iterate over the raw logs and unpacked data for TokenCreate events raised by the FourMeme contract.
type FourMemeTokenCreateIterator struct {
	Event *FourMemeTokenCreate // Event containing the contract specifics and raw log

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
func (it *FourMemeTokenCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeTokenCreate)
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
		it.Event = new(FourMemeTokenCreate)
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
func (it *FourMemeTokenCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeTokenCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeTokenCreate represents a TokenCreate event raised by the FourMeme contract.
type FourMemeTokenCreate struct {
	Creator     common.Address
	Token       common.Address
	RequestId   *big.Int
	Name        string
	Symbol      string
	TotalSupply *big.Int
	LaunchTime  *big.Int
	LaunchFee   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenCreate is a free log retrieval operation binding the contract event 0x396d5e902b675b032348d3d2e9517ee8f0c4a926603fbc075d3d282ff00cad20.
//
// Solidity: event TokenCreate(address creator, address token, uint256 requestId, string name, string symbol, uint256 totalSupply, uint256 launchTime, uint256 launchFee)
func (_FourMeme *FourMemeFilterer) FilterTokenCreate(opts *bind.FilterOpts) (*FourMemeTokenCreateIterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "TokenCreate")
	if err != nil {
		return nil, err
	}
	return &FourMemeTokenCreateIterator{contract: _FourMeme.contract, event: "TokenCreate", logs: logs, sub: sub}, nil
}

// WatchTokenCreate is a free log subscription operation binding the contract event 0x396d5e902b675b032348d3d2e9517ee8f0c4a926603fbc075d3d282ff00cad20.
//
// Solidity: event TokenCreate(address creator, address token, uint256 requestId, string name, string symbol, uint256 totalSupply, uint256 launchTime, uint256 launchFee)
func (_FourMeme *FourMemeFilterer) WatchTokenCreate(opts *bind.WatchOpts, sink chan<- *FourMemeTokenCreate) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "TokenCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeTokenCreate)
				if err := _FourMeme.contract.UnpackLog(event, "TokenCreate", log); err != nil {
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

// ParseTokenCreate is a log parse operation binding the contract event 0x396d5e902b675b032348d3d2e9517ee8f0c4a926603fbc075d3d282ff00cad20.
//
// Solidity: event TokenCreate(address creator, address token, uint256 requestId, string name, string symbol, uint256 totalSupply, uint256 launchTime, uint256 launchFee)
func (_FourMeme *FourMemeFilterer) ParseTokenCreate(log types.Log) (*FourMemeTokenCreate, error) {
	event := new(FourMemeTokenCreate)
	if err := _FourMeme.contract.UnpackLog(event, "TokenCreate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the FourMeme contract.
type FourMemeTokenPurchaseIterator struct {
	Event *FourMemeTokenPurchase // Event containing the contract specifics and raw log

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
func (it *FourMemeTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeTokenPurchase)
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
		it.Event = new(FourMemeTokenPurchase)
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
func (it *FourMemeTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeTokenPurchase represents a TokenPurchase event raised by the FourMeme contract.
type FourMemeTokenPurchase struct {
	Token   common.Address
	Account common.Address
	Price   *big.Int
	Amount  *big.Int
	Cost    *big.Int
	Fee     *big.Int
	Offers  *big.Int
	Funds   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x7db52723a3b2cdd6164364b3b766e65e540d7be48ffa89582956d8eaebe62942.
//
// Solidity: event TokenPurchase(address token, address account, uint256 price, uint256 amount, uint256 cost, uint256 fee, uint256 offers, uint256 funds)
func (_FourMeme *FourMemeFilterer) FilterTokenPurchase(opts *bind.FilterOpts) (*FourMemeTokenPurchaseIterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "TokenPurchase")
	if err != nil {
		return nil, err
	}
	return &FourMemeTokenPurchaseIterator{contract: _FourMeme.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x7db52723a3b2cdd6164364b3b766e65e540d7be48ffa89582956d8eaebe62942.
//
// Solidity: event TokenPurchase(address token, address account, uint256 price, uint256 amount, uint256 cost, uint256 fee, uint256 offers, uint256 funds)
func (_FourMeme *FourMemeFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *FourMemeTokenPurchase) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "TokenPurchase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeTokenPurchase)
				if err := _FourMeme.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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

// ParseTokenPurchase is a log parse operation binding the contract event 0x7db52723a3b2cdd6164364b3b766e65e540d7be48ffa89582956d8eaebe62942.
//
// Solidity: event TokenPurchase(address token, address account, uint256 price, uint256 amount, uint256 cost, uint256 fee, uint256 offers, uint256 funds)
func (_FourMeme *FourMemeFilterer) ParseTokenPurchase(log types.Log) (*FourMemeTokenPurchase, error) {
	event := new(FourMemeTokenPurchase)
	if err := _FourMeme.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeTokenPurchase2Iterator is returned from FilterTokenPurchase2 and is used to iterate over the raw logs and unpacked data for TokenPurchase2 events raised by the FourMeme contract.
type FourMemeTokenPurchase2Iterator struct {
	Event *FourMemeTokenPurchase2 // Event containing the contract specifics and raw log

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
func (it *FourMemeTokenPurchase2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeTokenPurchase2)
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
		it.Event = new(FourMemeTokenPurchase2)
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
func (it *FourMemeTokenPurchase2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeTokenPurchase2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeTokenPurchase2 represents a TokenPurchase2 event raised by the FourMeme contract.
type FourMemeTokenPurchase2 struct {
	Origin *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase2 is a free log retrieval operation binding the contract event 0x48063b1239b68b5d50123408787a6df1f644d9160f0e5f702fefddb9a855954d.
//
// Solidity: event TokenPurchase2(uint256 origin)
func (_FourMeme *FourMemeFilterer) FilterTokenPurchase2(opts *bind.FilterOpts) (*FourMemeTokenPurchase2Iterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "TokenPurchase2")
	if err != nil {
		return nil, err
	}
	return &FourMemeTokenPurchase2Iterator{contract: _FourMeme.contract, event: "TokenPurchase2", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase2 is a free log subscription operation binding the contract event 0x48063b1239b68b5d50123408787a6df1f644d9160f0e5f702fefddb9a855954d.
//
// Solidity: event TokenPurchase2(uint256 origin)
func (_FourMeme *FourMemeFilterer) WatchTokenPurchase2(opts *bind.WatchOpts, sink chan<- *FourMemeTokenPurchase2) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "TokenPurchase2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeTokenPurchase2)
				if err := _FourMeme.contract.UnpackLog(event, "TokenPurchase2", log); err != nil {
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

// ParseTokenPurchase2 is a log parse operation binding the contract event 0x48063b1239b68b5d50123408787a6df1f644d9160f0e5f702fefddb9a855954d.
//
// Solidity: event TokenPurchase2(uint256 origin)
func (_FourMeme *FourMemeFilterer) ParseTokenPurchase2(log types.Log) (*FourMemeTokenPurchase2, error) {
	event := new(FourMemeTokenPurchase2)
	if err := _FourMeme.contract.UnpackLog(event, "TokenPurchase2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeTokenSaleIterator is returned from FilterTokenSale and is used to iterate over the raw logs and unpacked data for TokenSale events raised by the FourMeme contract.
type FourMemeTokenSaleIterator struct {
	Event *FourMemeTokenSale // Event containing the contract specifics and raw log

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
func (it *FourMemeTokenSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeTokenSale)
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
		it.Event = new(FourMemeTokenSale)
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
func (it *FourMemeTokenSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeTokenSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeTokenSale represents a TokenSale event raised by the FourMeme contract.
type FourMemeTokenSale struct {
	Token   common.Address
	Account common.Address
	Price   *big.Int
	Amount  *big.Int
	Cost    *big.Int
	Fee     *big.Int
	Offers  *big.Int
	Funds   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenSale is a free log retrieval operation binding the contract event 0x0a5575b3648bae2210cee56bf33254cc1ddfbc7bf637c0af2ac18b14fb1bae19.
//
// Solidity: event TokenSale(address token, address account, uint256 price, uint256 amount, uint256 cost, uint256 fee, uint256 offers, uint256 funds)
func (_FourMeme *FourMemeFilterer) FilterTokenSale(opts *bind.FilterOpts) (*FourMemeTokenSaleIterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "TokenSale")
	if err != nil {
		return nil, err
	}
	return &FourMemeTokenSaleIterator{contract: _FourMeme.contract, event: "TokenSale", logs: logs, sub: sub}, nil
}

// WatchTokenSale is a free log subscription operation binding the contract event 0x0a5575b3648bae2210cee56bf33254cc1ddfbc7bf637c0af2ac18b14fb1bae19.
//
// Solidity: event TokenSale(address token, address account, uint256 price, uint256 amount, uint256 cost, uint256 fee, uint256 offers, uint256 funds)
func (_FourMeme *FourMemeFilterer) WatchTokenSale(opts *bind.WatchOpts, sink chan<- *FourMemeTokenSale) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "TokenSale")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeTokenSale)
				if err := _FourMeme.contract.UnpackLog(event, "TokenSale", log); err != nil {
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

// ParseTokenSale is a log parse operation binding the contract event 0x0a5575b3648bae2210cee56bf33254cc1ddfbc7bf637c0af2ac18b14fb1bae19.
//
// Solidity: event TokenSale(address token, address account, uint256 price, uint256 amount, uint256 cost, uint256 fee, uint256 offers, uint256 funds)
func (_FourMeme *FourMemeFilterer) ParseTokenSale(log types.Log) (*FourMemeTokenSale, error) {
	event := new(FourMemeTokenSale)
	if err := _FourMeme.contract.UnpackLog(event, "TokenSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeTokenSale2Iterator is returned from FilterTokenSale2 and is used to iterate over the raw logs and unpacked data for TokenSale2 events raised by the FourMeme contract.
type FourMemeTokenSale2Iterator struct {
	Event *FourMemeTokenSale2 // Event containing the contract specifics and raw log

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
func (it *FourMemeTokenSale2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeTokenSale2)
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
		it.Event = new(FourMemeTokenSale2)
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
func (it *FourMemeTokenSale2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeTokenSale2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeTokenSale2 represents a TokenSale2 event raised by the FourMeme contract.
type FourMemeTokenSale2 struct {
	Origin *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenSale2 is a free log retrieval operation binding the contract event 0x741ffc4605df23259462547defeab4f6e755bdc5fbb6d0820727d6d3400c7e0d.
//
// Solidity: event TokenSale2(uint256 origin)
func (_FourMeme *FourMemeFilterer) FilterTokenSale2(opts *bind.FilterOpts) (*FourMemeTokenSale2Iterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "TokenSale2")
	if err != nil {
		return nil, err
	}
	return &FourMemeTokenSale2Iterator{contract: _FourMeme.contract, event: "TokenSale2", logs: logs, sub: sub}, nil
}

// WatchTokenSale2 is a free log subscription operation binding the contract event 0x741ffc4605df23259462547defeab4f6e755bdc5fbb6d0820727d6d3400c7e0d.
//
// Solidity: event TokenSale2(uint256 origin)
func (_FourMeme *FourMemeFilterer) WatchTokenSale2(opts *bind.WatchOpts, sink chan<- *FourMemeTokenSale2) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "TokenSale2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeTokenSale2)
				if err := _FourMeme.contract.UnpackLog(event, "TokenSale2", log); err != nil {
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

// ParseTokenSale2 is a log parse operation binding the contract event 0x741ffc4605df23259462547defeab4f6e755bdc5fbb6d0820727d6d3400c7e0d.
//
// Solidity: event TokenSale2(uint256 origin)
func (_FourMeme *FourMemeFilterer) ParseTokenSale2(log types.Log) (*FourMemeTokenSale2, error) {
	event := new(FourMemeTokenSale2)
	if err := _FourMeme.contract.UnpackLog(event, "TokenSale2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeTradeStopIterator is returned from FilterTradeStop and is used to iterate over the raw logs and unpacked data for TradeStop events raised by the FourMeme contract.
type FourMemeTradeStopIterator struct {
	Event *FourMemeTradeStop // Event containing the contract specifics and raw log

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
func (it *FourMemeTradeStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeTradeStop)
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
		it.Event = new(FourMemeTradeStop)
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
func (it *FourMemeTradeStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeTradeStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeTradeStop represents a TradeStop event raised by the FourMeme contract.
type FourMemeTradeStop struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTradeStop is a free log retrieval operation binding the contract event 0x8f9ab4bd7eff0d085f91575d50cd83f97aa5258e24ded7630d4fd6739e857132.
//
// Solidity: event TradeStop(address token)
func (_FourMeme *FourMemeFilterer) FilterTradeStop(opts *bind.FilterOpts) (*FourMemeTradeStopIterator, error) {

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "TradeStop")
	if err != nil {
		return nil, err
	}
	return &FourMemeTradeStopIterator{contract: _FourMeme.contract, event: "TradeStop", logs: logs, sub: sub}, nil
}

// WatchTradeStop is a free log subscription operation binding the contract event 0x8f9ab4bd7eff0d085f91575d50cd83f97aa5258e24ded7630d4fd6739e857132.
//
// Solidity: event TradeStop(address token)
func (_FourMeme *FourMemeFilterer) WatchTradeStop(opts *bind.WatchOpts, sink chan<- *FourMemeTradeStop) (event.Subscription, error) {

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "TradeStop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeTradeStop)
				if err := _FourMeme.contract.UnpackLog(event, "TradeStop", log); err != nil {
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

// ParseTradeStop is a log parse operation binding the contract event 0x8f9ab4bd7eff0d085f91575d50cd83f97aa5258e24ded7630d4fd6739e857132.
//
// Solidity: event TradeStop(address token)
func (_FourMeme *FourMemeFilterer) ParseTradeStop(log types.Log) (*FourMemeTradeStop, error) {
	event := new(FourMemeTradeStop)
	if err := _FourMeme.contract.UnpackLog(event, "TradeStop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FourMemeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FourMeme contract.
type FourMemeUpgradedIterator struct {
	Event *FourMemeUpgraded // Event containing the contract specifics and raw log

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
func (it *FourMemeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FourMemeUpgraded)
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
		it.Event = new(FourMemeUpgraded)
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
func (it *FourMemeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FourMemeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FourMemeUpgraded represents a Upgraded event raised by the FourMeme contract.
type FourMemeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FourMeme *FourMemeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FourMemeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FourMeme.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FourMemeUpgradedIterator{contract: _FourMeme.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FourMeme *FourMemeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FourMemeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FourMeme.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FourMemeUpgraded)
				if err := _FourMeme.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_FourMeme *FourMemeFilterer) ParseUpgraded(log types.Log) (*FourMemeUpgraded, error) {
	event := new(FourMemeUpgraded)
	if err := _FourMeme.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
