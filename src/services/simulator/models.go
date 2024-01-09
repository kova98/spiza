package main

type Courier struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CourierAssigned struct {
	OrderId   int64 `json:"order_id"`
	CourierId int64 `json:"courier_id"`
}
