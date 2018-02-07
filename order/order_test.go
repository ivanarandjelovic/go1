package order

import (
	"testing"

	"github.com/ivanarandjelovic/go1/account"
)

const (
	orderAmount     = 1.1
	orderAmountHigh = 999999
)

var (
	moneyAccountUSD    = account.Account{Balance: 1.0, Currency: "USD"}
	securityAccountUSD = account.Account{Balance: 3.0, Currency: "USD"}
	securityAccountBTC = account.Account{Balance: 2.0, Currency: "BTC"}
)

func TestOrderConstructorOk(t *testing.T) {
	order, err := New(0, orderAmount, securityAccountBTC, moneyAccountUSD)
	if err != nil {
		t.Error("Order creation failed with: " + err.Error())
	}
	if order.Amount != orderAmount {
		t.Error("Order Amount incorrect")
	}
	if order.SecurityAccount != securityAccountBTC {
		t.Error("Order moneyAccount incorrect")
	}
	if order.MoneyAccount != moneyAccountUSD {
		t.Error("Order DebitingAccount incorrect")
	}
}

func TestOrderConstructorSameCurrency(t *testing.T) {
	order, err := New(1, orderAmount, securityAccountUSD, moneyAccountUSD)
	if err == nil {
		t.Error("Order creation expected to fail")
	}
	if order != nil {
		t.Error("Order object expected to be nil")
	}
}

func TestOrderConstructorWrongDirection(t *testing.T) {
	order, err := New(2, orderAmount, securityAccountUSD, moneyAccountUSD)
	if err == nil {
		t.Error("Order creation expected to fail due to invalid direction")
	}
	if order != nil {
		t.Error("Order object expected to be nil")
	}
}

func TestOrderConstructorBalanceLow(t *testing.T) {
	order, err := New(1, orderAmountHigh, securityAccountUSD, moneyAccountUSD)
	if err == nil {
		t.Error("Order creation expected to fail due to insufficient balance")
	}
	if order != nil {
		t.Error("Order object expected to be nil")
	}
}
