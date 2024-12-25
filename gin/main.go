package main

import (
	"math"
	"math/big"
	"net/http"

	"web3/gin/account"
	"web3/gin/block"
	"web3/gin/blockchain"
	trans "web3/gin/transaction"

	"github.com/gin-gonic/gin"
)

var Client *blockchain.Client

func main() {
	Client = blockchain.NewClient("ethereum", "http://localhost:8545")

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
	r.GET("/transactions/:id", getTransaction)
	r.GET("/transactions/count", getTransactionCount)

	// Contract routes
	r.POST("/contracts", deployContract)
	r.POST("/contracts/:id/call", callContract)
	r.GET("/contracts", listContracts)
	r.GET("/contracts/:id", getContract)

	// Event routes
	r.GET("/events", listEvents)
	r.GET("/events/:id", getEvent)
	r.GET("/events/listen", listenEvents)

	r.Run(":8080")
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
		// 假设我们有一个函数 getBalance(accountID string, blkNum *big.Int) (string, error) 来获取余额
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
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"message": "Contract deployed"})
}

func callContract(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"message": "Contract called"})
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
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"events": []string{}})
}

func getEvent(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"event": "event details"})
}

func listenEvents(c *gin.Context) {
	// ...existing code...
	c.JSON(http.StatusOK, gin.H{"message": "Listening to events"})
}

// 假设这是一个获取余额的函数
func getBalance(accountID string) (string, error) {
	// 这里应该实现实际的获取余额逻辑
	return "1000", nil
}