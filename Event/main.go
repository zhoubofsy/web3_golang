package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	//avalancheFujiURL = "https://api.avax-test.network/ext/bc/C/rpc" // Avalanche Fuji节点
	avalancheFujiURL   = "wss://api.avax-test.network/ext/bc/C/ws"
	contractAddressHex = "0x8Bf0438A0c1D77412F33459C27a2aB2F935931bC" // 部署的合约地址
	fromBlockNumber    = 36424122                                     // 从哪个区块开始查询，若不指定区块，默认从0区块开始，最大支持2048个区块
)

type EventProcess interface {
	DisplayEvent(types.Log)
}

type EthEvent struct {
	peth  *ethclient.Client
	pABI  *abi.ABI
	query ethereum.FilterQuery
}

func (e *EthEvent) ListEvents(ep EventProcess) {
	logs, err := e.peth.FilterLogs(context.Background(), e.query)
	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog)
		ep.DisplayEvent(vLog)
	}
}

func (e *EthEvent) SubscribeEvents(ep EventProcess) {
	go func() {
		// 创建一个订阅器
		logs := make(chan types.Log)
		sub, err := e.peth.SubscribeFilterLogs(context.Background(), e.query, logs)
		if err != nil {
			log.Fatalf("Failed to subscribe logs: %v", err)
		}

		// 处理日志
		for {
			select {
			case err := <-sub.Err():
				log.Fatalf("Failed to subscribe logs: %v", err)
			case vLog := <-logs:
				ep.DisplayEvent(vLog)
			}
		}
	}()
}

type TransferEthEvent struct {
	EthEvent
}

func (t *TransferEthEvent) DisplayEvent(vLog types.Log) {
	var originalData struct {
		from  common.Address
		to    common.Address
		Value *big.Int
	}
	err := t.pABI.UnpackIntoInterface(&originalData, "Transfer", vLog.Data)
	if err != nil {
		log.Fatalf("Failed to unpack event data: %v", err)
	}

	fmt.Printf("From: %s\n", vLog.Topics[1].Hex())
	fmt.Printf("To: %s\n", vLog.Topics[2].Hex())
	fmt.Printf("Value: %s\n", originalData.Value.String())
}

func NewTransferEthEvent(peth *ethclient.Client, pABI *abi.ABI) *TransferEthEvent {
	return &TransferEthEvent{
		EthEvent: EthEvent{
			peth: peth,
			pABI: pABI,
			query: ethereum.FilterQuery{
				FromBlock: big.NewInt(fromBlockNumber),
				Addresses: []common.Address{common.HexToAddress(contractAddressHex)},
				Topics:    [][]common.Hash{{pABI.Events["Transfer"].ID}},
			},
		},
	}
}

type ApprovalEthEvent struct {
	EthEvent
}

func (a *ApprovalEthEvent) DisplayEvent(vLog types.Log) {
	var originalData struct {
		owner   common.Address
		spender common.Address
		Value   *big.Int
	}
	err := a.pABI.UnpackIntoInterface(&originalData, "Approval", vLog.Data)
	if err != nil {
		log.Fatalf("Failed to unpack event data: %v", err)
	}

	fmt.Printf("Owner: %s\n", vLog.Topics[1].Hex())
	fmt.Printf("Spender: %s\n", vLog.Topics[2].Hex())
	fmt.Printf("Value: %s\n", originalData.Value.String())
}

func NewApprovalEthEvent(peth *ethclient.Client, pABI *abi.ABI) *ApprovalEthEvent {
	return &ApprovalEthEvent{
		EthEvent: EthEvent{
			peth: peth,
			pABI: pABI,
			query: ethereum.FilterQuery{
				FromBlock: big.NewInt(fromBlockNumber),
				Addresses: []common.Address{common.HexToAddress(contractAddressHex)},
				Topics:    [][]common.Hash{{pABI.Events["Approval"].ID}},
			},
		},
	}
}

func main() {
	client, err := ethclient.Dial(avalancheFujiURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Avalanche client: %v", err)
	}

	// load abi.json
	contractABI, err := os.ReadFile("./abi.json")
	if err != nil {
		log.Fatalf("os.ReadFile error , %v", err)
	}
	contractABIstr := string(contractABI)

	// 解析合约ABI，获取事件签名
	parsedABI, err := abi.JSON(strings.NewReader(contractABIstr))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	// List Approval events
	approvalEvent := NewApprovalEthEvent(client, &parsedABI)
	approvalEvent.ListEvents(approvalEvent)

	// List Transfer events
	transferEvent := NewTransferEthEvent(client, &parsedABI)
	transferEvent.ListEvents(transferEvent)

	approvalEvent.SubscribeEvents(approvalEvent)
	transferEvent.SubscribeEvents(transferEvent)

	for {
		time.Sleep(1 * time.Second)
	}
}

func ListApprovalEvents(client *ethclient.Client, pabi abi.ABI) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(36424122),
		Addresses: []common.Address{common.HexToAddress(contractAddressHex)},
		Topics:    [][]common.Hash{{pabi.Events["Approval"].ID}},
		//Topics:    [][]common.Hash{{pabi.Events["Approval"].ID, pabi.Events["Transfer"].ID}}, // Approval or Transfer
		//Topics:    [][]common.Hash{{pabi.Events["Approval"].ID}, {pabi.Events["Transfer"].ID}},	// Approval and Transfer
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}
	for _, vLog := range logs {
		var originalData struct {
			from  common.Address
			to    common.Address
			Value *big.Int
		}
		err := pabi.UnpackIntoInterface(&originalData, "Approval", vLog.Data)
		if err != nil {
			log.Fatalf("Failed to unpack event data: %v", err)
		}

		fmt.Printf("Owner: %s\n", vLog.Topics[1].Hex())
		fmt.Printf("Spender: %s\n", vLog.Topics[2].Hex())
		fmt.Printf("Value: %s\n", originalData.Value.String())
	}
}

func ListTransferEvents(client *ethclient.Client, pabi abi.ABI) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(36424122),
		Addresses: []common.Address{common.HexToAddress(contractAddressHex)},
		Topics:    [][]common.Hash{{pabi.Events["Transfer"].ID}},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}
	for _, vLog := range logs {
		var originalData struct {
			from  common.Address
			to    common.Address
			Value *big.Int
		}
		err := pabi.UnpackIntoInterface(&originalData, "Transfer", vLog.Data)
		if err != nil {
			log.Fatalf("Failed to unpack event data: %v", err)
		}

		fmt.Printf("From: %s\n", vLog.Topics[1].Hex())
		fmt.Printf("To: %s\n", vLog.Topics[2].Hex())
		fmt.Printf("Value: %s\n", originalData.Value.String())
	}
}
