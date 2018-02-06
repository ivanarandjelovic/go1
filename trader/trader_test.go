package trader

import (
	"testing"
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
