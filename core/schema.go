package core

type Block struct {
	Index     int64  `json:"index"`
	Timestamp int64  `json:"timestamp"`
	PrevHash  []byte `json:"prevhash"`
	Hash      []byte `json:"hash"`
	Data      []byte `json:"data"` // Byte encoded transaction
	Nonce     int64  `json:"nonce"`
}

type Blockchain struct {
	Blocks     []Block
	Difficulty int
}

type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int64  `json:"amount"`
}
