package main

import (
	"flag"
	"log"
	"os"

	eth "web3/Calls/ethcallers"
)

const (
	avalancheFujiURL   = "https://api.avax-test.network/ext/bc/C/rpc" // Avalanche Fuji节点
	contractAddressHex = "0x8Bf0438A0c1D77412F33459C27a2aB2F935931bC" // 部署的合约地址
)

// main()
func main() {
	var privateKey string
	flag.StringVar(&privateKey, "privateKey", "private.pem", "Private key for the sender account")
	flag.Parse()

	privKey, _ := os.ReadFile(privateKey)

	ethcaller, err := eth.NewEthCaller(avalancheFujiURL, contractAddressHex)
	if err != nil {
		log.Fatalf("Failed to create EthCaller: %v", err)
	}
	defer ethcaller.Close()

	// totalSupplyCaller := eth.NewTotalSupplyCaller(ethcaller)
	// ethcaller.View(totalSupplyCaller)

	// nameCaller := eth.NewNameCaller(ethcaller)
	// ethcaller.View(nameCaller)

	// balanceofCaller := eth.NewBalanceofCaller(ethcaller, "0xff10a3a7cb9007abb4e1a61f2680572d5fe5d489")
	// ethcaller.View(balanceofCaller)

	// allowance := eth.NewAllowance(ethcaller, "0xddef358f8a1d3e2d6d986ba358af4c6e435b921b", "0xf01b149542d1284d138d1bf7e59a252f47db3cc2")
	// ethcaller.View(allowance)

	approval := eth.NewApproval(ethcaller, "0xf01b149542d1284d138d1bf7e59a252f47db3cc2", 1100000000000000000)
	ethcaller.Trans(approval, string(privKey))

}

// func main() {
// 	// 连接到Avalanche Fuji网络
// 	client, err := ethclient.Dial(avalancheFujiURL)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the Avalanche client: %v", err)
// 	}

// 	// 合约ABI（根据您的合约生成的ABI）
// 	abiByte, err := os.ReadFile("./abi.json")
// 	if err != nil {
// 		log.Fatalf("os.ReadFile error , %v", err)
// 	}
// 	contractABI := string(abiByte)
// 	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
// 	if err != nil {
// 		log.Fatalf("Failed to parse contract ABI: %v", err)
// 	}

// 	// 合约地址
// 	contractAddress := common.HexToAddress(contractAddressHex)

// 	// 调用合约方法
// 	var result *big.Int
// 	callData, err := parsedABI.Pack("totalSupply") // 替换为您的方法名
// 	if err != nil {
// 		log.Fatalf("Failed to pack data for contract call: %v", err)
// 	}

// 	msg := ethereum.CallMsg{
// 		To:   &contractAddress,
// 		Data: callData,
// 	}

// 	response, err := client.CallContract(context.Background(), msg, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to call contract: %v", err)
// 	}

// 	err = parsedABI.UnpackIntoInterface(&result, "totalSupply", response)
// 	if err != nil {
// 		log.Fatalf("Failed to unpack contract response: %v", err)
// 	}

// 	fmt.Printf("Result: %s\n", result.String())
// }
