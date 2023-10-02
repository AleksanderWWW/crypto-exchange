package orderbook

import "fmt"

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

func (l *Limit) String() string {
	return fmt.Sprintf("[price: %.2f | volume: %.2f]", l.Price, l.TotalVolume)
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
