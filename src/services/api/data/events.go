package data

import "time"

type OrderStatusUpdated struct {
	Status       int
	DeliveryTime time.Time
}
