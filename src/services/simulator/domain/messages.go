package domain

import "time"

type CourierAssigned struct {
	OrderId   int64 `json:"orderId"`
	CourierId int64 `json:"courierId"`
}

type OrderUpdated struct {
	Id           int64     `json:"id"`
	Status       int       `json:"status"`
	DeliveryTime time.Time `json:"deliveryTime"`
}

const (
	OrderStatusCreated   = 0
	OrderStatusAccepted  = 1
	OrderStatusRejected  = 2
	OrderStatusReady     = 3
	OrderStatusPickedUp  = 4
	OrderStatusDelivered = 5
)
