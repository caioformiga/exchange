package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var prices = []float64{
	2947,
	2946.06,
	2945.52,
	2945.34,
	2946,
}

var amounts = []float64{
	0.10089033,
	1.2,
	4.16479668,
	19.9976434,
}

var exchanges = []string{"coinex", "kucoin", "binance", "klever"}

var symbols = []string{"USDT/KLV", "BTC/KLV", "USDT/BTC", "ETH/BTC"}

func TestNewOrderBook(t *testing.T) {
	var i = 0
	ex := exchanges[i]
	s := symbols[i]
	emptyBids := []BookEntry{}

	var ob OrderBook
	var err error

	ob, err = NewOrderBook(ex, s, emptyBids)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, ex, ob.ExchangeName)
	assert.Equal(t, s, ob.Symbol)
	assert.Equal(t, emptyBids, ob.Bids)

	//cant have nil as parameter []BookEntry{nil} != of []BookEntry{}
	ob, err = NewOrderBook(ex, s, nil)
	assert.NotNil(t, err, "err should not be nil")
	print(err)
}

func TestAddNewBookEntry(t *testing.T) {
	ob, err := NewOrderBook("coinex", "USDT/KLV", []BookEntry{})
	assert.Nil(t, err, "err should be nil")

	i := 0
	price := prices[i]
	amount := amounts[i]
	ob.AddBookEntry(price, amount)

	assert.Equal(t, prices[i], ob.Bids[i].Price)
	assert.Equal(t, amounts[i], ob.Bids[i].Amount)
}
