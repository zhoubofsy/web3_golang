package trans

import (
	"context"
	"log"
	"math/big"
	"web3/gin/blockchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type OpTrans struct {
	client *blockchain.Client
}

func NewOpTrans(bcClient *blockchain.Client) *OpTrans {
	return &OpTrans{client: bcClient}
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
