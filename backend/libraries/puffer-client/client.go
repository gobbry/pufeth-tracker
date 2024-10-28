package pufferclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gobbry/puffering/libraries/abis/multicall"
	PufferVaultV3 "github.com/gobbry/puffering/libraries/abis/puffervaultv3"
)

var PUFFER_ADDRESS = common.HexToAddress("0xd9a442856c234a39a81a089c06451ebaa4306a72")
var MULTICALL_ADDRESS = common.HexToAddress("0xca11bde05977b3631167028862be2a173976ca11")

type MulticallInput struct {
	Target       common.Address
	AllowFailure bool
	Calldata     []byte
}

type Aggregate3Response struct {
	Success    bool
	ReturnData []byte
}

type PufferConversionRate struct {
	TotalAssets    *big.Float
	TotalSupply    *big.Float
	ConversionRate *big.Float
}

func GetPufEthConversionRate(ctx context.Context, client *ethclient.Client, block *big.Int) (*PufferConversionRate, error) {
	pufferVaultAbi, err := PufferVaultV3.PufferVaultV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	multicallAbi, err := multicall.MulticallMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	totalSupplyCalldata, err := pufferVaultAbi.Pack("totalSupply")
	if err != nil {
		return nil, err
	}

	totalAssetsCallData, err := pufferVaultAbi.Pack("totalAssets")
	if err != nil {
		return nil, err
	}

	decimalsCallData, err := pufferVaultAbi.Pack("decimals")
	if err != nil {
		return nil, err
	}

	calls := []multicall.Multicall3Call3{
		{
			Target:       PUFFER_ADDRESS,
			AllowFailure: true,
			CallData:     totalSupplyCalldata,
		},
		{
			Target:       PUFFER_ADDRESS,
			AllowFailure: true,
			CallData:     totalAssetsCallData,
		},
		{
			Target:       PUFFER_ADDRESS,
			AllowFailure: true,
			CallData:     decimalsCallData,
		},
	}

	multicallData, err := multicallAbi.Pack("aggregate3", calls)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{To: &MULTICALL_ADDRESS, Data: multicallData}

	result, err := client.CallContract(ctx, msg, block)
	if err != nil {
		return nil, err
	}

	var multicallResults []multicall.Multicall3Result
	err = multicallAbi.UnpackIntoInterface(&multicallResults, "aggregate3", result)
	if err != nil {
		return nil, err
	}

	totalSupply := new(big.Int).SetBytes(multicallResults[0].ReturnData)
	totalAssets := new(big.Int).SetBytes(multicallResults[1].ReturnData)
	conversionRate := new(big.Float).Quo(new(big.Float).SetInt(totalAssets), new(big.Float).SetInt(totalSupply))

	decimals := new(big.Int).SetBytes(multicallResults[2].ReturnData)
	divisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), decimals, nil))

	return &PufferConversionRate{
		TotalAssets:    new(big.Float).Quo(new(big.Float).SetInt(totalAssets), divisor),
		TotalSupply:    new(big.Float).Quo(new(big.Float).SetInt(totalSupply), divisor),
		ConversionRate: conversionRate,
	}, nil
}
