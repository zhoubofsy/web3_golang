package blockchain

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	bcType string
	Eth    *ethclient.Client
}

func NewClient(bcType string, url string) *Client {
	switch bcType {
	case "ethereum":
		ethClient, err := ethclient.Dial(url)
		if err != nil {
			panic(err)
		}
		return &Client{
			bcType: bcType,
			Eth:    ethClient,
		}
	default:
		panic("unsupported blockchain type")
	}
}

func (c *Client) Close() {
	switch c.bcType {
	case "ethereum":
		c.Eth.Close()
	}
}
