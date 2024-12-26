package trans

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"web3/gin/blockchain"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type OpTrans struct {
	client *blockchain.Client
}

func NewOpTrans(bcClient *blockchain.Client) *OpTrans {
	return &OpTrans{client: bcClient}
}

func (op *OpTrans) Transfer(to string, value uint64) (string, error) {
	// 1. 使用私钥生成 ECDSA 密钥对
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		return "", err
	}
	// 2. 从私钥中获取公钥
	pubKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("failed to get public key")
	}
	// 3. 从公钥中获取From地址
	fromAddress := crypto.PubkeyToAddress(*pubKey)
	toAddress := common.HexToAddress(to)

	// 4. 获取当前账户(From地址)的nonce值
	nonce, err := op.client.Eth.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	// 5. 设置转账金额
	val := big.NewInt(int64(value * 1000000000000000000)) // in wei (1 eth)
	// 6. 设置gasLimit
	//gasLimit := uint64(21000) // in units
	gasLimit, err := op.client.Eth.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  fromAddress,
		To:    &toAddress,
		Value: val,
		Data:  nil,
	})
	if err != nil {
		return "", err
	}
	// 7. 获取当前推荐的gasPrice
	gasPrice, err := op.client.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	// 8. 获取当前网络的chainID
	chainId, err := op.client.Eth.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	// 9. 创建交易
	tx := types.NewTransaction(nonce, toAddress, val, gasLimit, gasPrice, nil)
	// 10. 签名交易
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainId.Int64())), privateKey)
	if err != nil {
		return "", err
	}
	// 11. 发送交易
	err = op.client.Eth.SendTransaction(context.Background(), signTx)
	if err != nil {
		return "", err
	}
	return signTx.Hash().Hex(), err
}

func (op *OpTrans) GetHeaderTransactionCount() (uint, error) {
	headerBlockNum, err := op.client.Eth.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockInfo, err := op.client.Eth.BlockByNumber(context.Background(), big.NewInt(int64(headerBlockNum.Number.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	return op.client.Eth.TransactionCount(context.Background(), blockInfo.Hash())
}

func (op *OpTrans) ListTX(blkHash string) ([]TXInfo, error) {
	block, err := op.client.Eth.BlockByHash(context.Background(), common.HexToHash(blkHash))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var txInfos []TXInfo

	for _, tx := range block.Transactions() {
		txHash := tx.Hash()
		receipt, err := op.client.Eth.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			continue
		}
		chainId := tx.ChainId()
		from, err := types.Sender(types.NewEIP155Signer(chainId), tx)
		if err != nil {
			continue
		}
		txInfo := TXInfo{
			TxHash:     txHash.Hex(),
			TxValue:    tx.Value().Uint64(),
			TxGas:      tx.Gas(),
			TxGasPrice: tx.GasPrice().Uint64(),
			TxNonce:    tx.Nonce(),
			TxData:     tx.Data(),
			TxTo:       tx.To().Hex(),
			TxReceipt:  uint8(receipt.Status),
			TxFrom:     from.Hex(),
		}
		txInfos = append(txInfos, txInfo)
	}
	return txInfos, nil
}
