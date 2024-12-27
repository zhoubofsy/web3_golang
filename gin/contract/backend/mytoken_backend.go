package backend

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MyTokenCB struct {
	c *ethclient.Client
}

func NewMyTokenCB(c *ethclient.Client) *MyTokenCB {
	return &MyTokenCB{c: c}
}

func (m MyTokenCB) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return m.c.CodeAt(ctx, contract, blockNumber)
}

func (m MyTokenCB) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return m.c.CallContract(ctx, call, blockNumber)
}

func (m MyTokenCB) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return m.c.EstimateGas(ctx, call)
}

func (m MyTokenCB) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return m.c.SuggestGasPrice(ctx)
}

func (m MyTokenCB) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return m.c.SuggestGasTipCap(ctx)
}

func (m MyTokenCB) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return m.c.SendTransaction(ctx, tx)
}

func (m MyTokenCB) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return m.c.HeaderByNumber(ctx, number)
}

func (m MyTokenCB) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return m.c.PendingCodeAt(ctx, account)
}

func (m MyTokenCB) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return m.c.PendingNonceAt(ctx, account)
}

func (m MyTokenCB) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return m.c.FilterLogs(ctx, q)
}

func (m MyTokenCB) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return m.c.SubscribeFilterLogs(ctx, q, ch)
}
