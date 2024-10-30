package ethcallers

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Approval struct {
	peth     *EthCaller
	funcName string
	spender  string
	amount   uint64
}

func NewApproval(e *EthCaller, spender string, amount uint64) *Approval {
	return &Approval{
		peth:     e,
		funcName: "approve",
		spender:  spender,
		amount:   amount,
	}
}

func (b *Approval) MakeCallData() ([]byte, error) {
	return b.peth.pABI.Pack(b.funcName, common.HexToAddress(b.spender), big.NewInt(int64(b.amount)))
}
