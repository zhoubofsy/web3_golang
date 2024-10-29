package ethcallers

import (
	"fmt"
	"math/big"
)

// 定义一个结构体来调用totalSupply方法
type TotalSupply struct {
	peth     *EthCaller
	funcName string
}

func NewTotalSupplyCaller(e *EthCaller) *TotalSupply {
	return &TotalSupply{
		peth:     e,
		funcName: "totalSupply",
	}
}

func (t *TotalSupply) MakeCallData() ([]byte, error) {
	return t.peth.pABI.Pack(t.funcName)
}

func (t *TotalSupply) ProcessResponse(response []byte) error {
	var result *big.Int
	err := t.peth.pABI.UnpackIntoInterface(&result, t.funcName, response)
	if err != nil {
		return err
	}

	fmt.Printf("Call Function: %s, Result: %s\n", t.funcName, result.String())
	return err
}
