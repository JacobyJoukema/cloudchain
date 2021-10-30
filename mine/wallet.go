package mine

type Wallet struct {
	PublicKey  string
	PrivateKey string
}

// TODO figure out how to create a unique wallet
// TODO figure out how to split tokens
func NewWallet(seed string) Wallet {
	wallet := Wallet{
		PublicKey:  seed,
		PrivateKey: seed,
	}

	return wallet
}
