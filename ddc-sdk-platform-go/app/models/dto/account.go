package dto

type Account struct {
	address    string
	publicKey  string
	privateKey string
	mnemonic   string
}

func (a *Account) Address() string {
	return a.address
}

func NewAccount(address string, publicKey string, privateKey string, mnemonic string) *Account {
	return &Account{address: address, publicKey: publicKey, privateKey: privateKey, mnemonic: mnemonic}
}

func (a *Account) SetAddress(address string) {
	a.address = address
}

func (a *Account) PublicKey() string {
	return a.publicKey
}

func (a *Account) SetPublicKey(publicKey string) {
	a.publicKey = publicKey
}

func (a *Account) PrivateKey() string {
	return a.privateKey
}

func (a *Account) SetPrivateKey(privateKey string) {
	a.privateKey = privateKey
}

func (a *Account) Mnemonic() string {
	return a.mnemonic
}

func (a *Account) SetMnemonic(mnemonic string) {
	a.mnemonic = mnemonic
}
