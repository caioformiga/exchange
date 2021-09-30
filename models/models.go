package models

import "fmt"

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

func (ob *OrderBook) AddBookEntry(price float64, amount float64) {
	b := BookEntry{
		Price:  price,
		Amount: amount,
	}
	ob.Bids = append(ob.Bids, b)
}
