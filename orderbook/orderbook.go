package orderbook

type Match struct {
	Ask        *Order
	Bid        *Order
	SizeFilled float64
}

type OrderBook struct {
	Asks []*Limit
	Bids []*Limit

	AskLimits map[float64]*Limit
	BidLimits map[float64]*Limit
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		Asks: []*Limit{},
		Bids: []*Limit{},

		AskLimits: map[float64]*Limit{},
		BidLimits: map[float64]*Limit{},
	}
}

func (ob *OrderBook) PlaceOrder(price float64, o *Order) []Match {
	if o.Size > 0.0 {
		ob.add(price, o)
	}

	return []Match{}
}

func (ob *OrderBook) add(price float64, o *Order) {
	var limit *Limit
	var ok bool

	if o.Bid {
		_, ok = ob.BidLimits[price]
	} else {
		_, ok = ob.AskLimits[price]
	}

	if !ok {
		limit = NewLimit(price)
		limit.AddOrder(o)
		if o.Bid {
			ob.Bids = append(ob.Bids, limit)
			ob.BidLimits[price] = limit
		} else {
			ob.Asks = append(ob.Asks, limit)
			ob.AskLimits[price] = limit
		}
	}
}
