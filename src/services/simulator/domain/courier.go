package domain

type Courier struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Loc            Location
	CurrentOrderId int64
}
