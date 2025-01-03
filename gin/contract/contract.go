package contract

import (
	"context"
	"errors"
	"math/big"
	"web3/gin/blockchain"
	"web3/gin/contract/backend"
	"web3/gin/contract/mytoken"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type TransactParams struct {
	TxFrom  string `json:"txFrom"`
	TxTo    string `json:"txTo"`
	TxValue uint64 `json:"txValue"`
}
type contract struct {
	client *blockchain.Client
}

func NewContract(client *blockchain.Client) *contract {
	return &contract{client: client}
}

func (c *contract) DeployContract() (string, string, error) {
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		return "", "", err
	}
	chainId, err := c.client.Eth.ChainID(context.Background())
	if err != nil {
		return "", "", err
	}
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return "", "", err
	}
	bk := backend.NewMyTokenCB(c.client.Eth)
	// TODO: 部署合约
	contractAddress, txHash, _, err := mytoken.DeployMytoken(txOpts, *bk)
	return contractAddress.Hex(), txHash.Hash().Hex(), err
}

func (c *contract) Call(addr string, category string, params interface{}) (interface{}, error) {
	var resp interface{}
	var err error

	instance, err := mytoken.NewMytoken(common.HexToAddress(addr), c.client.Eth)
	if err != nil {
		return nil, err
	}
	switch category {
	case "BalanceOf":
		callOpts := &bind.CallOpts{
			Pending: false,
			Context: context.Background(),
		}
		if account, ok := params.(string); ok {
			bBlance, err := instance.BalanceOf(callOpts, common.HexToAddress(account))
			if err != nil {
				return nil, err
			}
			resp = bBlance.String()
		} else {
			err = errors.New("invalid params")
		}
	case "Transfer":
		// 使用私钥生成 ECDSA 密钥对
		privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
		if err != nil {
			return nil, errors.New("failed to get private key")
		}
		// 获取当前账户的地址
		pubAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
		// 获取当前账户的nonce值
		nonce, err := c.client.Eth.PendingNonceAt(context.Background(), pubAddr)
		if err != nil {
			return nil, errors.New("failed to get nonce")
		}
		// 获取当前推荐的gasPrice
		gasPrice, err := c.client.Eth.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, errors.New("failed to get gas price")
		}
		gasLimit := uint64(30000000)
		txOpts := &bind.TransactOpts{
			From:  pubAddr,
			Nonce: big.NewInt(int64(nonce)),
			Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
				chainId, err := c.client.Eth.ChainID(context.Background())
				if err != nil {
					return nil, errors.New("failed to get chain id")
				}
				return types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)

			},
			Value:    nil,
			GasPrice: gasPrice,
			GasLimit: gasLimit,
			Context:  context.Background(),
		}
		if txParams, ok := params.(TransactParams); ok {
			resp, err = instance.TransferFrom(txOpts, common.HexToAddress(txParams.TxFrom),
				common.HexToAddress(txParams.TxTo), big.NewInt(int64(txParams.TxValue)))
		} else {
			return nil, errors.New("invalid params")
		}

	default:
		return nil, errors.New("unsupported category")
	}
	return resp, err
}
