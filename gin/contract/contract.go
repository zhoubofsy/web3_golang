package contract

import (
	"context"
	"web3/gin/blockchain"
	"web3/gin/contract/backend"
	"web3/gin/contract/mytoken"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

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