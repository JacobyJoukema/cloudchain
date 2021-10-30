package main

import (
	"github.com/inflowml/logger"
	"github.com/jacobyjoukema/cloudchain/core"
	"github.com/jacobyjoukema/cloudchain/mine"
)

func main() {
	chain, err := core.NewBlockchain() // Initialize the blockchain
	if err != nil {
		logger.Error("failed to create new blockchain: %v", err)
		panic(err)
	}

	wallet := mine.NewWallet("jacobyjoukema")

	logger.Info("Starting Mining Process")
	mineError := mine.Mine(wallet, chain)
	logger.Info("error stopped mining process: %v", <-mineError)
}
