package trader

// Package trader implements one trading platform which deals with one security type
// (usually one currency or one symbol)

import (
	"errors"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

func New(currency string) (*traderType, error) {
	if currency == "" {
		return nil, errors.New("Currency must not be empty or nil")
	}
	inverseFloat64Comparator := func(a, b interface{}) int {
		return -utils.Float64Comparator(a, b)
	}
	buyOrders := binaryheap.NewWith(inverseFloat64Comparator)
	sellOrders := binaryheap.NewWith(utils.Float64Comparator)
	return &traderType{Currency: currency, BuyOrders: buyOrders, SellOrders: sellOrders}, nil
}

type traderType struct {
	Currency   string
	BuyOrders  *binaryheap.Heap
	SellOrders *binaryheap.Heap
}
