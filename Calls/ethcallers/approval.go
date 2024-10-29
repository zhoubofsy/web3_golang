package ethcallers

import (
	"fmt"
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

func (b *Approval) ProcessResponse(response []byte) error {
	var result bool
	err := b.peth.pABI.UnpackIntoInterface(&result, b.funcName, response)
	if err != nil {
		return err
	}

	fmt.Printf("Call Function: %s, Result: %v\n", b.funcName, result)
	return err
}

func (b *Approval) GetPrivateKey() string {
	return ""
}
