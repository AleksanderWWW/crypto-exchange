package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id        string
	Size      float64
	Bid       bool
	Limit     *Limit
	Timestamp int64
}

func NewOrder(size float64, bid bool) *Order {
	return &Order{
		Id:        uuid.New().String(),
		Size:      size,
		Bid:       bid,
		Timestamp: time.Now().UnixNano(),
	}
}

func (o *Order) String() string {
	return fmt.Sprintf("[size: %.2f]", o.Size)
}

type Limit struct {
	Price       float64
	Orders      map[string]*Order
	TotalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: map[string]*Order{},
	}
}

func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders[o.Id] = o
	l.TotalVolume += o.Size
}

func (l *Limit) DeleteOrder(o *Order) error {
	_, ok := l.Orders[o.Id]
	if !ok {
		return fmt.Errorf("Cannot delete order %s as it does not exist", o.Id)
	}

	delete(l.Orders, o.Id)
	l.TotalVolume -= o.Size
	o.Limit = nil

	return nil
}

type OrderBook struct {
	Asks []*Limit
	Bids []*Limit
}
