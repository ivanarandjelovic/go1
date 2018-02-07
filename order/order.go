package order

import (
	"errors"

	"github.com/ivanarandjelovic/go1/account"
)

// Order is a structure representing single trading order
type Order struct {
	// direction 0 - buy, 1 - sell
	Direction       int
	Amount          float64
	SecurityAccount account.Account
	MoneyAccount    account.Account
}

func New(direction int, amount float64, securityAccount account.Account, moneyAccount account.Account) (*Order, error) {
	if direction != 0 && direction != 1 {
		return nil, errors.New("Invalid direction, allowed values: 0 -buy, 1 - sell")
	}
	if direction == 1 && amount > securityAccount.Balance {
		return nil, errors.New("Order amount over the balance")
	}
	if securityAccount.Currency == moneyAccount.Currency {
		return nil, errors.New("security and money account cannot have same currency")
	} else {
		return &Order{Amount: amount, SecurityAccount: securityAccount, MoneyAccount: moneyAccount}, nil
	}
}
