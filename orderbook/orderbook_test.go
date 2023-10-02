package orderbook

import "testing"

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrderA := NewOrder(5, true)
	buyOrderB := NewOrder(8, true)
	buyOrderC := NewOrder(1, true)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	l.DeleteOrder(buyOrderB)

	if l.TotalVolume != 6 {
		t.Error("Invalid total volume")
	}

}

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
