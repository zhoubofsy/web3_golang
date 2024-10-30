package ethcallers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthCaller struct {
	client *ethclient.Client
	pABI   *abi.ABI
	addr   common.Address
}

func NewEthCaller(url string, contract string) (*EthCaller, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	abiByte, err := os.ReadFile("./abi.json")
	if err != nil {
		return nil, err
	}
	pABI, err := abi.JSON(strings.NewReader(string(abiByte)))

	return &EthCaller{
		client: client,
		pABI:   &pABI,
		addr:   common.HexToAddress(contract),
	}, err
}

type EthViewInterface interface {
	MakeCallData() ([]byte, error)
	ProcessResponse(response []byte) error
}

type EthTransInterface interface {
	MakeCallData() ([]byte, error)
}

func (e *EthCaller) View(ei EthViewInterface) error {
	// 调用合约方法
	callData, err := ei.MakeCallData()
	if err != nil {
		log.Fatalf("Failed to make call data for contract call: %v", err)
		return err
	}

	msg := ethereum.CallMsg{
		To:   &e.addr,
		Data: callData,
	}

	response, err := e.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatalf("Failed to call contract: %v", err)
		return err
	}

	err = ei.ProcessResponse(response)
	if err != nil {
		log.Fatalf("Failed to process contract response: %v", err)
	}
	return err
}

func (e *EthCaller) Close() {
	e.client.Close()
}

func (e *EthCaller) Trans(ei EthTransInterface, privateKeyHex string) error {
	//hash, err := e.ClassicTrans(ei, privateKeyHex)
	hash, err := e.NewTrans(ei, privateKeyHex)
	if err != nil {
		return err
	}
	return e.WaitForConfirm(hash)
}

func (e *EthCaller) NewTrans(ei EthTransInterface, privateKeyHex string) (common.Hash, error) {
	// 加载私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:]) // 去掉开头的0x
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// 获取公钥和地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to assert public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 构造交易
	nonce, err := e.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// 签署交易
	chainId, err := e.client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	header, err := e.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get header: %v", err)
	}
	gasTipCap, err := e.client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas tip cap: %v", err)
	}

	data, err := ei.MakeCallData()
	if err != nil {
		log.Fatalf("Failed to make call data for contract call: %v", err)
	}
	tx := &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		To:        &e.addr,
		Value:     big.NewInt(0),
		Gas:       30000,
		GasFeeCap: header.BaseFee,
		GasTipCap: gasTipCap,
		Data:      data,
	}

	signedTX, err := types.SignNewTx(privateKey, types.LatestSignerForChainID(chainId), tx)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	err = e.client.SendTransaction(context.Background(), signedTX)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}
	log.Printf("Transaction sent: %s", signedTX.Hash().Hex())
	return signedTX.Hash(), err
}

func (e *EthCaller) ClassicTrans(ei EthTransInterface, privateKeyHex string) (common.Hash, error) {
	// 加载私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:]) // 去掉开头的0x
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// 获取公钥和地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to assert public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 构造交易
	nonce, err := e.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	// 创建合约调用数据
	data, err := ei.MakeCallData()
	if err != nil {
		log.Fatalf("Failed to make call data: %v", err)
	}

	// 构建交易
	gasLimit := uint64(30000) // 如何才能获取一个合理的Gas Limit ？
	tx := types.NewTransaction(nonce, e.addr, big.NewInt(0), gasLimit, gasPrice, data)

	// 签署交易
	chainID, err := e.client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}
	// 使用 types.SignNewTx() 更符合新版本推荐要求
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// 发送交易
	err = e.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	// 打印交易哈希
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	return signedTx.Hash(), err
}

func (e *EthCaller) WaitForConfirm(txHash common.Hash) error {
	receipt, err := e.client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("transaction reverted , hash: %s", receipt.TxHash.Hex())
	}
	return nil
}
