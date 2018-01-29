package order

import (
	"github.com/ivanarandjelovic/go1/account"
)

// Order represents a trading order
type Order interface {
	getAmount() float64
	getDebitingAccount() account.Account
	getCreditingAccount() account.Account
}

type orderType struct {
	amount           float64
	debitingAccount  account.Account
	creditingAccount account.Account
}

func (order *orderType) getAmount() float64 {
	return order.amount
}

func (order *orderType) getDebitingAccount() account.Account {
	return order.debitingAccount
}

func (order *orderType) getCreditingAccount() account.Account {
	return order.creditingAccount
}

func createOrder(amount float64, debitingAccount account.Account, creditingAccount account.Account) Order {
	return &orderType{amount: amount, debitingAccount: debitingAccount, creditingAccount: creditingAccount}
}
