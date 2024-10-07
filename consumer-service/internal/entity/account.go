package entity

type Account struct {
	AccountId      string
	AccountBalance float64
}

func NewAccount(id string, balance float64) *Account {
	return &Account{
		AccountId:      id,
		AccountBalance: balance,
	}
}
