package main

import (
	"encoding/json"
	"io"
	"math"
	"math/big"
	"net/http"
	"strconv"

	"web3/gin/account"
	"web3/gin/block"
	"web3/gin/blockchain"
	"web3/gin/contract"
	"web3/gin/event"
	"web3/gin/signature"
	trans "web3/gin/transaction"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

var Client *blockchain.Client
var WSClient *event.ContractEvent

func main() {
	Client = blockchain.NewClient("ethereum", "http://localhost:8545")
	WSClient := event.NewEvent("ws://127.0.0.1:8545")
	if WSClient == nil {
		panic("Failed to create event")
	}
	go WSClient.Run("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0")

	r := gin.Default()

	// Account routes
	r.POST("/accounts", createAccount)
	r.GET("/accounts", listAccounts)
	r.GET("/accounts/:id", getAccount)
	r.DELETE("/accounts/:id", deleteAccount)
	r.GET("/account/balance/:id", getAccountBalance)

	// Block routes
	r.GET("/blocks", listBlocks)
	r.GET("/blocks/count", getBlockCount)

	// Transaction routes
	r.GET("/transactions", listTransactions)
	r.POST("/transfer", transfer)
	r.GET("/transactions/:id", getTransaction)
	r.GET("/transactions/count", getTransactionCount)

	// Contract routes
	r.POST("/contracts", deployContract)
	r.POST("/contracts/:id/:method", callContract)
	r.GET("/contracts", listContracts)
	r.GET("/contracts/:id", getContract)

	// Event routes
	r.GET("/events", listEvents)
	r.GET("/events/:id", getEvent)
	r.GET("/events/listen", listenEvents)

	// Signature routes
	r.POST("/sign/:method", sign)

	r.Run(":8080")
}

// Signature handlers
func sign(c *gin.Context) {
	method := c.Param("method")
	privateKey := c.Request.Header.Get("privateKey")
	op := signature.NewSigner(privateKey)
	if op == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create signer"})
		return
	}
	switch method {
	case "make":
		byteBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
			return
		}
		sig, err := op.MakeSignature(byteBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"signatureHex": hexutil.Encode(sig)})
	case "verify":
		sig := c.Query("sig")
		byteBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
			return
		}

		byteSig, err := hexutil.Decode(sig)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode signature"})
			return
		}
		if op.VerfiySignature(byteBody, byteSig) {
			c.JSON(http.StatusOK, gin.H{"result": "Signature verified"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signature"})
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid method"})
	}
}

// Account handlers
func createAccount(c *gin.Context) {
	op := account.NewOpAccount(Client)
	if op == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
	}

	val := c.Request.Header.Get("category")
	// if len(vals) > 0 {
	// 	val = vals[0]
	// }
	if val == "account" {
		acc, err := op.CreateAccount("12345")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"address": acc})
	} else if val == "keys" {

		privKey, pubKey, err := op.GenerateKeys()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"publicKey": pubKey, "privateKey": privKey})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid account type"})
	}
}

func listAccounts(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"accounts": []string{}})
}

func getAccount(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"account": "account details"})
}

func deleteAccount(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
}

func getAccountBalance(c *gin.Context) {
	accountID := c.Param("id")

	var balance *big.Int
	var err error
	op := account.NewOpAccount(Client)
	if !op.IsAccount(accountID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account is a contract."})
		return
	}
	val := c.Query("pending")
	if val == "1" {
		balance, err = op.GetPendingBalance(accountID)
	} else {
		balance, err = op.GetBalance(accountID, nil)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fbalance, _ := new(big.Float).SetString(balance.String())
	if fbalance == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert balance to float"})
		return
	}
	eths := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	c.JSON(http.StatusOK, gin.H{"balance": eths})
}

// Block handlers
func listBlocks(c *gin.Context) {
	var start, end uint64
	op := block.NewOpBlock(Client)
	blkNum, err := op.GetBlockNumber()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	end = blkNum
	blks, err := op.ListBlocks(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blocks": blks})
}

func getBlockCount(c *gin.Context) {
	// ...existing code...
	op := block.NewOpBlock(Client)
	count, err := op.GetBlockNumber()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count + 1})
}

// Transaction handlers
func listTransactions(c *gin.Context) {
	blkHash := c.Query("blockhash")
	op := trans.NewOpTrans(Client)
	txInfos, err := op.ListTX(blkHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transactions": txInfos})
}

func transfer(c *gin.Context) {
	txTo := c.Request.Header.Get("TxTo")
	val, err := strconv.Atoi(c.Request.Header.Get("Value"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value"})
		return
	}
	op := trans.NewOpTrans(Client)
	tx, err := op.Transfer(txTo, uint64(val))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Transfer": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"Transfer": "success!", "txHash": tx})
	}
}

func getTransaction(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"transaction": "transaction details"})
}

func getTransactionCount(c *gin.Context) {
	op := trans.NewOpTrans(Client)
	count, err := op.GetHeaderTransactionCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

// Contract handlers
func deployContract(c *gin.Context) {
	op := contract.NewContract(Client)
	addr, tx, err := op.DeployContract()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contract deployed", "address": addr, "txHash": tx})
}

func callContract(c *gin.Context) {
	id := c.Param("id")
	method := c.Param("method")
	var params interface{}
	switch method {
	case "BalanceOf":
		params = c.Request.Header.Get("Address")
	case "Transfer":
		jsonParams := contract.TransactParams{}
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
			return
		}
		if nil != json.Unmarshal(body, &jsonParams) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
			return
		}
		params = jsonParams
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid method"})
		return
	}

	op := contract.NewContract(Client)
	resp, err := op.Call(id, method, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if balance, ok := resp.(string); ok {
		c.JSON(http.StatusOK, gin.H{"message": "success", "balance": balance})
	} else if txHash, ok := resp.(*types.Transaction); ok {
		c.JSON(http.StatusOK, gin.H{"message": "success", "txHash": txHash.Hash().Hex()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "unknow resp."})
	}
}

func listContracts(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"contracts": []string{}})
}

func getContract(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"contract": "contract details"})
}

// Event handlers
func listEvents(c *gin.Context) {
	cEventLog := event.NewEvent("ws://127.0.0.1:8545")
	cAddr := c.Query("contractAddr")
	logs, err := cEventLog.ListWithBlkId(cAddr, 0, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": logs})
}

func getEvent(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"event": "event details"})
}

func listenEvents(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"message": "Listening to events"})
}
