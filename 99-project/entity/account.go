package entity

import (
	"time"
)

type Account struct {
	ID            ID
	AccountType   string
	AccountNumber string
	Amount        float64
	Limit         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewAccount(accountType string, accountNumber string) (*Account, error) {
	b := &Account{
		ID:            NewID(),
		AccountType:   accountType,
		AccountNumber: accountNumber,
		CreatedAt:     time.Now(),
	}
	err := b.Validate()
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (account *Account) Validate() error {

	if account.AccountNumber == "" {
		return ErrAccountNumberInvalid
	}
	return nil
}


