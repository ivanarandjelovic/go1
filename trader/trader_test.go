package trader

import (
	"testing"

	"github.com/ivanarandjelovic/go1/account"
	"github.com/ivanarandjelovic/go1/order"
	//"github.com/ivanarandjelovic/go1/account"
)

const (
	currency = "USD"
)

func TestTraderConstructorOk(t *testing.T) {
	trader, err := New(currency)
	if err != nil {
		t.Error("Unexpected error: " + err.Error())
	}
	if trader.Currency != currency {
		t.Error("Unexpected currency: " + trader.Currency + ", " + currency + " expected")
	}
	if trader.BuyOrders == nil || trader.SellOrders == nil {
		t.Error("Order queues are empty!")
	}
}

func TestTraderConstructorEmptyCurrency(t *testing.T) {
	trader, err := New("")
	if err == nil {
		t.Error("Expected error for empty currency!")
	}
	if trader != nil {
		t.Error("Unexpected object, nil expected")
	}
}

const (
	orderAmount     = 1.1
	orderAmountHigh = 999999
)

var (
	moneyAccountUSD    = account.Account{Balance: 1.0, Currency: "USD"}
	securityAccountUSD = account.Account{Balance: 3.0, Currency: "USD"}
	securityAccountBTC = account.Account{Balance: 2.0, Currency: "BTC"}
)

func TestComparator(t *testing.T) {
	order1, err := order.New(0, orderAmount, securityAccountBTC, moneyAccountUSD)
	if err != nil {

	}
	order2, err := order.New(0, orderAmount+1, securityAccountBTC, moneyAccountUSD)
	if orderComparator(order1, order2) >= 0 {
		t.Error("order1 to be smaller")
	}
	if inverseorderComparator(order1, order2) < 0 {
		t.Error("order1 to be larger")
	}
}
