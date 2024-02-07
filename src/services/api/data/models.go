package data

import (
	"time"

	"github.com/lib/pq"
)

type Address struct {
	Id          int64  `json:"id,omitempty"`
	FullAddress string `json:"fullAddress" db:"full_address"`
	LatLng      string `json:"latLng" db:"lat_lng"`
}

type Restaurant struct {
	Id             int64          `json:"id,omitempty"`
	Name           string         `json:"name"`
	Address        Address        `json:"address"`
	MenuCategories []MenuCategory `json:"menuCategories" db:"menu_categories"`
}

type MenuCategory struct {
	Id           int64  `json:"id,omitempty"`
	Name         string `json:"name"`
	RestaurantId int64  `json:"restaurantId" db:"restaurant_id"`
	Items        []Item `json:"items"`
}

type Item struct {
	Id           int64   `json:"id,omitempty"`
	CategoryId   int64   `json:"categoryId" db:"category_id"`
	CategoryName string  `json:"categoryName,omitempty" db:"category_name"`
	Name         string  `json:"name"`
	Order        int32   `json:"order" db:"order_num"`
	Price        float64 `json:"price"`
	Description  string  `json:"description"`
	Image        string  `json:"image"`
}

type Order struct {
	Id            int64         `json:"id"`
	UserId        int64         `json:"userId" db:"user_id"`
	RestaurantId  int64         `json:"restaurantId" db:"restaurant_id"`
	DestinationId int64         `json:"destinationId" db:"destination_id"`
	Status        int64         `json:"status"`
	DateCreated   time.Time     `json:"dateCreated" db:"date_created"`
	Items         pq.Int64Array `json:"items"`
}

func (o Order) WithItems(i []Item) OrderWithItems {
	var orderItems []OrderItem
	for _, item := range i {
		orderItems = append(orderItems, OrderItem{
			Id:      item.Id,
			Name:    item.Name,
			Order:   item.Order,
			Price:   item.Price,
			Image:   item.Image,
			OrderId: o.Id,
		})
	}
	return OrderWithItems{
		Id:          o.Id,
		UserId:      o.UserId,
		Status:      o.Status,
		DateCreated: o.DateCreated,
		Items:       orderItems,
	}
}
