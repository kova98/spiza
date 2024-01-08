package data

import "time"

type OrderStatusUpdated struct {
	Status       int       `json:"status"`
	DeliveryTime time.Time `json:"delivery_time"`
}
