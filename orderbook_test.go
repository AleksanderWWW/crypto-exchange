package main

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
