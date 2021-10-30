package core

import "fmt"

// Creates a new blockchain and sets the default difficulty
func NewBlockchain() (*Blockchain, error) {
	genesisBlock, err := NewGenesisBlock()
	if err != nil {
		return &Blockchain{}, fmt.Errorf("failed to generate genesis block: %v", err)
	}

	chain := &Blockchain{
		Blocks:     []Block{genesisBlock},
		Difficulty: 5,
	}
	return chain, nil
}

// TODO add networking functions
func (blockchain *Blockchain) AddBlock(block Block) error {

	blockchain.Blocks = append(blockchain.Blocks, block)

	return nil
}
