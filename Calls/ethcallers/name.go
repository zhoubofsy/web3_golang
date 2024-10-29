package ethcallers

import "fmt"

// 定义一个结构体来调用name方法
type Name struct {
	peth     *EthCaller
	funcName string
}

func NewNameCaller(e *EthCaller) *Name {
	return &Name{
		peth:     e,
		funcName: "name",
	}
}

func (n *Name) MakeCallData() ([]byte, error) {
	return n.peth.pABI.Pack(n.funcName)
}

func (n *Name) ProcessResponse(response []byte) error {
	var result string
	err := n.peth.pABI.UnpackIntoInterface(&result, n.funcName, response)
	if err != nil {
		return err
	}

	fmt.Printf("Call Function: %s, Result: %s\n", n.funcName, result)
	return err
}
