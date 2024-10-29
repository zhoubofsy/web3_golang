package ethcallers

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BalanceOf struct {
	peth     *EthCaller
	funcName string
	addr     string
}

func NewBalanceofCaller(e *EthCaller, addr string) *BalanceOf {
	return &BalanceOf{
		peth:     e,
		funcName: "balanceOf",
		addr:     addr,
	}
}

func (b *BalanceOf) MakeCallData() ([]byte, error) {
	return b.peth.pABI.Pack(b.funcName, common.HexToAddress(b.addr))
}

func (b *BalanceOf) ProcessResponse(response []byte) error {
	var result *big.Int
	err := b.peth.pABI.UnpackIntoInterface(&result, b.funcName, response)
	if err != nil {
		return err
	}

	fmt.Printf("Call Function: %s, Result: %s\n", b.funcName, result.String())
	return err
}
