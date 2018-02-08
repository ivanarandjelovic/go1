package trader

// Package trader implements one trading platform which deals with one security type
// (usually one currency or one symbol)

import (
	"errors"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"

	"github.com/ivanarandjelovic/go1/order"
)

// orderComparator compares order by amount, lower ammount is first
func orderComparator(a, b interface{}) int {
	return utils.Float64Comparator(a.(*order.Order).Amount, b.(*order.Order).Amount)
}

// inverseorderComparator compares order by amount, higher ammount is first
func inverseorderComparator(a, b interface{}) int {
	return -orderComparator(a, b)
}

func New(currency string) (*traderType, error) {
	if currency == "" {
		return nil, errors.New("Currency must not be empty or nil")
	}

	buyOrders := binaryheap.NewWith(inverseorderComparator)
	sellOrders := binaryheap.NewWith(orderComparator)
	return &traderType{Currency: currency, BuyOrders: buyOrders, SellOrders: sellOrders}, nil
}

type traderType struct {
	Currency   string
	BuyOrders  *binaryheap.Heap
	SellOrders *binaryheap.Heap
}
