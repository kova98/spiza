package data

import "time"

type OrderStatusUpdated struct {
	Status       int       `json:"status"`
	DeliveryTime time.Time `json:"deliveryTime"`
}

type CourierLocationUpdated struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type CourierAssigned struct {
	OrderId   int64 `json:"orderId"`
	CourierId int64 `json:"courierId"`
}

type OrderDelivered struct {
	OrderId   int64 `json:"orderId"`
	CourierId int64 `json:"courierId"`
}
