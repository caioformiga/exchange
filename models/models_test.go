package models

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	name     string
	exchange string
	symbol   string
	bids     []BookEntry
}

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

var higherBid = BookEntry{
	Price:  3000,
	Amount: 3000,
}

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

func TestAddBid(t *testing.T) {
	ob, err := NewOrderBook("coinex", "USDT/KLV", []BookEntry{})
	assert.Nil(t, err, "err should be nil")

	i := 0
	price := prices[i]
	amount := amounts[i]
	ob.AddBid(price, amount)

	assert.Equal(t, prices[i], ob.Bids[i].Price)
	assert.Equal(t, amounts[i], ob.Bids[i].Amount)
}

func TestSortBids(t *testing.T) {
	ob, err := NewOrderBook("coinex", "USDT/KLV", []BookEntry{})
	assert.Nil(t, err, "err should be nil")

	var testCaseScenarios []TestData = CreateTestCaseScenarios(higherBid)
	for _, testData := range testCaseScenarios {
		lastBid := testData.bids[len(testData.bids)-1]
		assert.Equal(t, lastBid, higherBid)

		ob.SortBids(testData.bids)
		firstBid := testData.bids[0]
		assert.Equal(t, firstBid, higherBid)

		isSorted := isDescSorted(testData.bids)
		assert.True(t, isSorted)
	}
}

func isDescSorted(bids []BookEntry) (isSorted bool) {
	isSorted = true
	for i, bid := range bids {		
		j := i + 1
		if j < len(bids) {
			if bid.Price < bids[j].Price {
				isSorted = false
				return

			}
		}
	}
	return
}

func isAscSorted(asks []BookEntry) (isSorted bool) {
	isSorted = true
	for i, bid := range asks {		
		j := i + 1
		if j < len(asks) {
			if bid.Price > asks[j].Price {
				isSorted = false
				return

			}
		}
	}
	return
}

func TestCreateTestCaseScenarios(t *testing.T) {
	var testCaseScenarios []TestData = CreateTestCaseScenarios(higherBid)
	assert.Equal(t, len(testCaseScenarios), 16)
	for _, testData := range testCaseScenarios {
		lastBid := testData.bids[len(testData.bids)-1]
		assert.Equal(t, lastBid, higherBid)
	}
}

func CreateTestCaseScenarios(higherBid BookEntry) []TestData {
	var testCaseScenarios []TestData
	for _, ex := range exchanges {

		for _, symbol := range symbols {

			testData := CreateTestDataCase(ex, symbol)
			testData.bids = append(testData.bids, higherBid)
			testCaseScenarios = append(testCaseScenarios, testData)
		}
	}
	return testCaseScenarios
}

func CreateTestDataCase(ex string, symbol string) (testData TestData) {
	ob, err := NewOrderBook(ex, symbol, []BookEntry{})
	if err != nil {
		fmt.Println(err)
	}

	var i int = 0
	var size = len(prices)
	for i < (size) {
		j := rand.Intn(size - 1)

		p := prices[j]
		a := amounts[j]

		ob.AddBid(p, a)

		i = i + 1
	}

	testData = TestData{
		name:     "random data",
		exchange: ob.ExchangeName,
		symbol:   ob.Symbol,
		bids:     ob.Bids,
	}
	return
}
