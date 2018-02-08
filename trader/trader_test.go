package trader

import (
	"fmt"
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
		t.Error("Unable to construct order: " + err.Error())
	}
	order2, err := order.New(0, orderAmount+1, securityAccountBTC, moneyAccountUSD)
	if orderComparator(order1, order2) >= 0 {
		t.Error("order1 to be smaller")
	}
	if inverseorderComparator(order1, order2) < 0 {
		t.Error("order1 to be larger")
	}
}

func TestBuyAndSellHeaps(t *testing.T) {
	order1, _ := order.New(0, orderAmount+1, securityAccountBTC, moneyAccountUSD)
	order2, _ := order.New(0, orderAmount+3, securityAccountBTC, moneyAccountUSD)
	order3, _ := order.New(0, orderAmount+4, securityAccountBTC, moneyAccountUSD)
	order4, _ := order.New(0, orderAmount+2, securityAccountBTC, moneyAccountUSD)

	trader, _ := New("BTC")

	trader.BuyOrders.Push(order1)
	trader.BuyOrders.Push(order2, order3, order4)

	if trader.BuyOrders.Size() != 4 {
		t.Error("Unexpected number of orders in Buy queue: " + string(trader.BuyOrders.Size()))
	}
	expectedAmount := orderAmount + 4
	for trader.BuyOrders.Size() > 0 {
		tmp, _ := trader.BuyOrders.Pop()
		o := tmp.(*order.Order)
		fmt.Println(o)
		if o.Amount != expectedAmount {
			t.Error("Unexpected amount :", o.Amount, " while expected: ", expectedAmount)
		} else {
			fmt.Println("Ok amount ", o.Amount)
		}
		expectedAmount -= 1.0
	}
}
