package order

import (
	"testing"

	"github.com/ivanarandjelovic/go1/account"
)

const (
	orderAmount = 1.1
)

var (
	creditingAccount   = account.Account{Balance: 1.0, Currency: "USD"}
	debitingAccountUSD = account.Account{Balance: 3.0, Currency: "USD"}
	debitingAccount    = account.Account{Balance: 2.0, Currency: "BTC"}
)

func TestOrderConstructorOk(t *testing.T) {
	order, err := createOrder(orderAmount, debitingAccount, creditingAccount)
	if err != nil {
		t.Error("Order creation failed with: " + err.Error())
	}
	if order.getAmount() != orderAmount {
		t.Error("Order Amount incorrect")
	}
	if order.getCreditingAccount() != creditingAccount {
		t.Error("Order CreditingAccount incorrect")
	}
	if order.getDebitingAccount() != debitingAccount {
		t.Error("Order DebitingAccount incorrect")
	}
}

func TestOrderConstructorSameCurrency(t *testing.T) {
	order, err := createOrder(orderAmount, debitingAccountUSD, creditingAccount)
	if err == nil {
		t.Error("Order creation expected to fail")
	}
	if order != nil {
		t.Error("Order object expected to be nil")
	}
}
