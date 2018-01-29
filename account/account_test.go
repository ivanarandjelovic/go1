package account

import (
	"testing"
)

func TestAccountProperties(t *testing.T) {
	account := Account{Balance: 1.0, Currency: "USD"}
	if account.Balance != 1.0 {
		t.Error("Account Balance incorrect")
	}
	if account.Currency != "USD" {
		t.Error("Account Currency incorrect")
	}
}
