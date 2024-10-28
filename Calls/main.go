package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	avalancheFujiURL   = "https://api.avax-test.network/ext/bc/C/rpc" // Avalanche Fuji节点
	contractAddressHex = "0xYourContractAddress"                      // 部署的合约地址
)

func main() {
	// 连接到Avalanche Fuji网络
	client, err := ethclient.Dial(avalancheFujiURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Avalanche client: %v", err)
	}

	// 合约ABI（根据您的合约生成的ABI）
	contractABI := `[{"constant":true,"inputs":[],"name":"yourMethodName","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	// 合约地址
	contractAddress := common.HexToAddress(contractAddressHex)

	// 调用合约方法
	var result *big.Int
	callData, err := parsedABI.Pack("yourMethodName") // 替换为您的方法名
	if err != nil {
		log.Fatalf("Failed to pack data for contract call: %v", err)
	}

	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}

	response, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatalf("Failed to call contract: %v", err)
	}

	err = parsedABI.UnpackIntoInterface(&result, "yourMethodName", response)
	if err != nil {
		log.Fatalf("Failed to unpack contract response: %v", err)
	}

	fmt.Printf("Result: %s\n", result.String())
}
