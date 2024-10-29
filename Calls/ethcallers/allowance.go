package ethcallers

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Allowance struct {
	peth     *EthCaller
	funcName string
	owner    string
	spender  string
}

func NewAllowance(e *EthCaller, owner string, spender string) *Allowance {
	return &Allowance{
		peth:     e,
		funcName: "allowance",
		owner:    owner,
		spender:  spender,
	}
}

func (b *Allowance) MakeCallData() ([]byte, error) {
	return b.peth.pABI.Pack(b.funcName, common.HexToAddress(b.owner), common.HexToAddress(b.spender))
}

func (b *Allowance) ProcessResponse(response []byte) error {
	var result *big.Int
	err := b.peth.pABI.UnpackIntoInterface(&result, b.funcName, response)
	if err != nil {
		return err
	}

	fmt.Printf("Call Function: %s, Result: %s\n", b.funcName, result.String())
	return err
}
