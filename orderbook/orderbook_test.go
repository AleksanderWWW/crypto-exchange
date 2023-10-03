package orderbook

import (
	"testing"
)

func TestPlaceLimitOrder(t *testing.T) {
	ob := NewOrderBook()

	buyOrderA := NewOrder(10, true)
	buyOrderB := NewOrder(2000, true)
	sellOrderA := NewOrder(1000, false)
	sellOrderB := NewOrder(2000, false)

	ob.PlaceLimitOrder(15_000, buyOrderA)
	ob.PlaceLimitOrder(18_000, buyOrderB)
	ob.PlaceLimitOrder(14_000, sellOrderA)
	ob.PlaceLimitOrder(14_000, sellOrderB)

	if len(ob.Bids) != 2 {
		t.Errorf("Invalid bids length: %d", len(ob.Bids))
	}

	if len(ob.Asks) != 1 {
		t.Errorf("Invalid asks length: %d", len(ob.Asks))
	}

	if len(ob.AskLimits) != 1 {
		t.Errorf("Invalid ask limit length: %d", len(ob.AskLimits))
	}

	if len(ob.BidLimits) != 2 {
		t.Errorf("Invalid bid limit length: %d", len(ob.BidLimits))
	}

	sellLimit := ob.AskLimits[14_000]
	if sellLimit.Price != 14_000 || sellLimit.TotalVolume != 1000 {
		t.Error("Invalid sell limit")
	}

}
