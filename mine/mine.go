package mine

import (
	"fmt"

	"github.com/inflowml/logger"
	"github.com/jacobyjoukema/cloudchain/core"
)

func Mine(wallet Wallet, chain *core.Blockchain) chan error {
	r := make(chan error)
	go func() {
		for true {
			newBlock, err := GenerateBlock(wallet, chain)
			if err != nil {
				r <- fmt.Errorf("failed to generate new block: %v", err)
			}

			block, err := MineBlock(newBlock, chain)
			if err != nil {
				r <- fmt.Errorf("failed to mine block: %v", err)
			}

			logger.Info("Mined Block: %v", block.Index)
			chain.AddBlock(block)
		}
	}()

	return r
}

func MineBlock(block core.Block, blockchain *core.Blockchain) (core.Block, error) {
	compString := ""
	for i := 0; i < blockchain.Difficulty; i++ {
		compString = fmt.Sprintf("%s0", compString)
	}

	for block.Hash[:blockchain.Difficulty] != compString {
		block.Nonce += 1
		block.SetHash()
		//logger.Info(block.Hash)
	}

	logger.Info("Broke out of mining loop after %v iterations", block.Nonce)
	logger.Info("Legal Hash: %s", block.Hash)

	return block, nil
}

func GenerateBlock(wallet Wallet, chain *core.Blockchain) (core.Block, error) {
	tran := core.Transaction{
		From:   "",
		To:     "jacobyjoukema",
		Amount: 1,
	}

	data, err := core.MarshalTransaction(tran)
	if err != nil {
		return core.Block{}, fmt.Errorf("failed to marshal transaction data: %v", err)
	}

	index := int64(len(chain.Blocks))

	return core.NewBlock(data, index, chain.Blocks[index-1].Hash), nil
}
