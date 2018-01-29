package order

import (
	"testing"

	"github.com/ivanarandjelovic/go1/account"
)

const (
	orderAmount = 1.1
)

var (
	creditingAccount = account.Account{Balance: 1.0, Currency: "USD"}
	debitingAccount  = account.Account{Balance: 2.0, Currency: "BTC"}
)

func TestOrderConstructorOk(t *testing.T) {
	order := createOrder(orderAmount, debitingAccount, creditingAccount)
	if order.Amount != orderAmount {
		t.Error("Order Amount incorrect")
	}
	if order.CreditingAccount != creditingAccount {
		t.Error("Order CreditingAccount incorrect")
	}
	if order.DebitingAccount != debitingAccount {
		t.Error("Order DebitingAccount incorrect")
	}
}