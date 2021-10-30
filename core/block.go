package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/inflowml/logger"
)

func (block *Block) CalcHash() string {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	nonce := []byte(strconv.FormatInt(block.Nonce, 10))
	index := []byte(strconv.FormatInt(block.Index, 10))
	headers := bytes.Join([][]byte{index, timestamp, []byte(block.PrevHash), block.Data, nonce}, []byte{})
	hash := sha256.Sum256(headers)

	return hex.EncodeToString(hash[:])
}

func (block *Block) SetHash() {
	block.Hash = block.CalcHash()[:]
}

func NewBlock(data []byte, index int64, prevBlockHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now().Unix(),
		PrevHash:  prevBlockHash,
		Hash:      "",
		Data:      data,
		Nonce:     0,
	}
	block.SetHash()
	return block
}

func NewGenesisBlock() (Block, error) {
	tran := Transaction{
		From:   "",
		To:     "jacobyjoukema",
		Amount: 1,
	}

	data, err := MarshalTransaction(tran)
	if err != nil {
		return Block{}, fmt.Errorf("failed to marshal genesis transaction: %v", err)
	}

	return NewBlock(data, 0, ""), nil
}

func CreateBlock(header string, body string) {

	logger.Info("Header: %s, Body: %s", header, body)
}

func MarshalTransaction(transaction Transaction) ([]byte, error) {
	data, err := json.Marshal(transaction)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to marshal transaction into bytes: %v", err)
	}

	return data, nil
}
