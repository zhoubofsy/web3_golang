package account

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"web3/gin/blockchain"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type AccountOperations interface {
	getBalance(accountID string) (string, error)
}

type OpAccount struct {
	client *blockchain.Client
}

var ErrAccountIDRequired = errors.New("account ID is required")
var Store *keystore.KeyStore

func NewOpAccount(bcClient *blockchain.Client) *OpAccount {
	return &OpAccount{client: bcClient}
}

func init() {
	// 这里应该实现实际的初始化逻辑
	if Store != nil {
		return
	}
	Store = keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
}

func (op *OpAccount) GetBalance(accountID string, blkNum *big.Int) (*big.Int, error) {
	if accountID == "" {
		return nil, ErrAccountIDRequired
	}
	// 这里应该实现实际的获取余额逻辑
	return op.client.Eth.BalanceAt(context.Background(), common.HexToAddress(accountID), blkNum)
}

func (op *OpAccount) GetPendingBalance(accountID string) (*big.Int, error) {
	if accountID == "" {
		return nil, ErrAccountIDRequired
	}
	return op.client.Eth.PendingBalanceAt(context.Background(), common.HexToAddress(accountID))
}

func (op *OpAccount) GenerateKeys() (string, string, error) {
	// 这里应该实现实际的生成密钥对逻辑
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("failed to get public key")
	}

	pubAddr := crypto.PubkeyToAddress(*publicKey).Hex()

	return hexutil.Encode(crypto.FromECDSA(privateKey)), pubAddr, nil
}

func (op *OpAccount) CreateAccount(passwd string) (string, error) {
	account, err := Store.NewAccount(passwd)
	fmt.Printf("account address : %s", account.Address.Hex())
	return account.Address.Hex(), err
}

func (op *OpAccount) ImportAccount(file string, passwd string) error {
	keyJSON, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	account, err := Store.Import(keyJSON, passwd, passwd)
	fmt.Printf("account address : %s", account.Address.Hex())
	return err
}

func (op *OpAccount) IsAccount(accountID string) bool {
	byteCode, err := op.client.Eth.CodeAt(context.Background(), common.HexToAddress(accountID), nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(byteCode) <= 0 {
		return true
	}
	return false
}
