package orderbook

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
	return fmt.Sprintf("[size: %.2f | timestamp: %d]", o.Size, o.Timestamp)
}
