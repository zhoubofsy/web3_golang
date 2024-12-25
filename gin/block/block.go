package block

import (
	"context"
	"math/big"
	"web3/gin/blockchain"
)

type OpBlock struct {
	client *blockchain.Client
}

func NewOpBlock(client *blockchain.Client) *OpBlock {
	return &OpBlock{
		client: client,
	}
}

func (op *OpBlock) GetBlockNumber() (uint64, error) {
	return op.client.Eth.BlockNumber(context.Background())
}

func (op *OpBlock) GetBlockInfo(number uint64) (BlockInfo, error) {
	blkInfo, err := op.client.Eth.BlockByNumber(context.Background(), big.NewInt(int64(number)))
	if err != nil {
		return BlockInfo{}, err
	}
	return BlockInfo{
		Hash:       blkInfo.Hash().Hex(),
		Height:     blkInfo.Number().Uint64(),
		Timestamp:  blkInfo.Time(),
		Difficulty: blkInfo.Difficulty().Uint64(),
		Nonce:      blkInfo.Nonce(),
		Miner:      blkInfo.Coinbase().Hex(),
		TransCount: uint64(len(blkInfo.Transactions())),
	}, err
}

func (op *OpBlock) ListBlocks(from, to uint64) ([]BlockInfo, error) {
	blkMaxNum, err := op.GetBlockNumber()
	if err != nil {
		return nil, err
	}
	start := max(from, 0)
	end := min(to, blkMaxNum)

	blocks := make([]BlockInfo, 0)
	for i := start; i <= end; i++ {
		bi, err := op.GetBlockInfo(i)
		if err != nil {
			continue
		}
		blocks = append(blocks, bi)
	}
	return blocks, nil
}
