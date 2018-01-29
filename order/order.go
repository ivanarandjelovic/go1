package order

import (
	"github.com/ivanarandjelovic/go1/account"
)

// Order represents a trading order
type Order interface {
}

type order struct {
	Amount           float64
	DebitingAccount  account.Account
	CreditingAccount account.Account
}

func createOrder(amount float64, debitingAccount account.Account, creditingAccount account.Account) *order {
	order := new(order)
	order.Amount = amount
	order.DebitingAccount = debitingAccount
	order.CreditingAccount = creditingAccount
	return order
}
