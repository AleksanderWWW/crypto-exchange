package orderbook

import "testing"

func TestOrderBook(t *testing.T) {
	ob := NewOrderBook()

	buyOrderA := NewOrder(10, true)
	buyOrderB := NewOrder(2000, true)
	ob.PlaceOrder(15_000, buyOrderA)
	ob.PlaceOrder(18_000, buyOrderB)

	if len(ob.Bids) != 2 {
		t.Errorf("Invalid bids length: %d", len(ob.Bids))
	}

}
