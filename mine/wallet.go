package mine

type Wallet struct {
	Account []byte
}

// TODO figure out how to create a unique wallet
// TODO figure out how to split tokens
func NewWallet(seed string) Wallet {
	wallet := Wallet{
		Account: []byte(seed),
	}

	return wallet
}
