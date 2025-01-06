package event

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"web3/gin/contract/mytoken"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	TransferTopic = "Transfer"
	ApproveTopic  = "Approval"
)

type ContractEvent struct {
	Client   *ethclient.Client
	TopicMap map[string]string
}

func NewEvent(addr string) *ContractEvent {
	client, err := ethclient.Dial(addr)
	if err != nil {
		return nil
	}
	c := &ContractEvent{Client: client}
	c.TopicMap = make(map[string]string)
	c.TopicMap[crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")).Hex()] = TransferTopic
	c.TopicMap[crypto.Keccak256Hash([]byte("Approval(address,address,uint256)")).Hex()] = ApproveTopic
	return c
}

func (e *ContractEvent) Run(contractAddr string) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddr)},
	}

	ch := make(chan types.Log)
	sp, err := e.Client.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		fmt.Printf("SubscribeFilterLogs err: %v\n", err)
		return
	}
	for {
		select {
		case err := <-sp.Err():
			fmt.Printf("sp err: %v\n", err)
			return
		case vLog := <-ch:
			logJSON, err := json.Marshal(vLog)
			if err != nil {
				fmt.Printf("json.Marshal err: %v\n", err)
				return
			}
			fmt.Printf("vLog: %s\n", logJSON)
		}
	}
}

type LogInfo struct {
	//Log         types.Log `json:"Log"`
	LogType     string `json:"LogType"`
	FromAddress string `json:"FromAddress"`
	ToAddress   string `json:"ToAddress"`
	ParseData   string `json:"ParseData"`
}

func (e *ContractEvent) ListWithBlkId(contractAddr string, fromBlock, toBlock uint64) ([]LogInfo, error) {
	if fromBlock > toBlock {
		return nil, fmt.Errorf("fromBlock > toBlock")
	}
	var fBlock *big.Int
	var tBlock *big.Int
	if fromBlock > 0 {
		fBlock = big.NewInt(int64(fromBlock))
	}
	if toBlock > 0 {
		tBlock = big.NewInt(int64(toBlock))
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddr)},
		FromBlock: fBlock,
		ToBlock:   tBlock,
	}

	logs, err := e.Client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	contractABI, err := abi.JSON(strings.NewReader(mytoken.MytokenABI))
	if err != nil {
		return nil, err
	}
	var logsInfo []LogInfo
	var logType string
	for _, vLog := range logs {
		if topic, ok := e.TopicMap[vLog.Topics[0].Hex()]; !ok {
			fmt.Printf("topic not found: %s\n", vLog.Topics[0].Hex())
			continue
		} else {
			logType = topic
		}
		parseData, err := contractABI.Unpack(logType, vLog.Data)
		if err != nil {
			if logJSON, err := json.Marshal(vLog); err == nil {
				fmt.Printf("Unpack err: %v\n %s", err, logJSON)
			} else {
				fmt.Printf("Unpack err: %v\n %v", err, vLog)
			}
			continue
		}
		var strData string
		switch logType {
		case TransferTopic:
			strData = parseData[0].(*big.Int).String()
		case ApproveTopic:
			strData = parseData[0].(*big.Int).String()
		default:
			fmt.Printf("Unknow parse data: %v\n", parseData)
			strData = "unknow"
		}
		logsInfo = append(logsInfo, LogInfo{
			//Log: vLog,
			LogType:     logType,
			FromAddress: common.HexToAddress(vLog.Topics[1].Hex()).Hex(),
			ToAddress:   common.HexToAddress(vLog.Topics[2].Hex()).Hex(),
			ParseData:   strData,
		})
	}
	return logsInfo, nil
}
