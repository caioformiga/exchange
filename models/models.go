package models

import (
	"fmt"
	"sort"
)

type BookEntry struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

type OrderBook struct {
	ExchangeName string      `json:"exchange_name"`
	Symbol       string      `json:"symbol"`
	Bids         []BookEntry `json:"bids"`
	//	Asks         []*BookEntry `json:"asks"`
}

func NewOrderBook(exName string, sym string, b []BookEntry) (ob OrderBook, err error) {
	if b == nil {
		err = fmt.Errorf("bids can't be nil, use BookEntry{} instead")
	}

	ob = OrderBook{
		ExchangeName: exName,
		Symbol:       sym,
		Bids:         b,
	}
	return
}

func (ob *OrderBook) AddBid(price float64, amount float64) (bid BookEntry, err error) {
	bid = BookEntry{
		Price:  price,
		Amount: amount,
	}

	if ob.Bids != nil {
		if price == 0 {
			err = fmt.Errorf("price can't be zero")
			return
		}

		if amount == 0 {
			err = fmt.Errorf("amout can't be zero")
			return
		}

		// Mutex lock
		ob.Bids = append(ob.Bids, bid)
		// Mutex unlock
	}

	err = fmt.Errorf("bids can't be nil")
	return
}

// SortBids sorts all BookEntry from Bids to asc mode
func (ob *OrderBook) SortBids(bids []BookEntry) (sorted []BookEntry) {
	sort.Slice(bids, func(i, j int) bool {
		return bids[i].Price > bids[j].Price
	})
	sorted = bids
	return
}

// SortAsks sorts all BookEntry from Bids to asc mode
func (ob *OrderBook) SortAsks(asks []BookEntry) (sorted []BookEntry) {
	sort.Slice(asks, func(i, j int) bool {
		return asks[i].Price < asks[j].Price
	})
	sorted = asks
	return
}
