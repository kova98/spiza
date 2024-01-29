package domain

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type Courier struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Loc            Location
	CurrentOrderId int64
	bus            Bus
	l              *log.Logger
}

func (c *Courier) AssignToOrder(orderId int64) {
	// TODO: make this thread safe?
	c.CurrentOrderId = orderId
}

func (c *Courier) Travel(orderId int64, path []Location) {
	// TODO: think about what happens if this happens before the courier has arrived at the restaurant
	locTopic := "order/" + strconv.FormatInt(orderId, 10) + "/courier-location"
	for _, loc := range path {
		c.bus.Publish(locTopic, loc)
		time.Sleep(1 * time.Second)
	}
	c.Loc = path[len(path)-1]
}

func (c *Courier) PickUpOrder(id int64) {
	statusMsg := OrderUpdated{
		Status:       OrderStatusPickedUp,
		DeliveryTime: calculateDeliveryTime(),
	}
	msg, _ := json.Marshal(statusMsg)
	c.bus.Publish("order/"+strconv.FormatInt(id, 10), msg)
	c.l.Println("Order" + strconv.FormatInt(id, 10) + " picked up")
}

func (c *Courier) CompleteOrder(id int64, msg OrderUpdated) {
	// complete order
	statusMsg := OrderUpdated{
		Status:       OrderStatusDelivered,
		DeliveryTime: msg.DeliveryTime,
	}
	marshalled, _ := json.Marshal(statusMsg)
	c.bus.Publish("order/"+strconv.FormatInt(id, 10), marshalled)
	c.l.Println("Order" + strconv.FormatInt(id, 10) + " completed")
}

func calculateDeliveryTime() time.Time {
	return time.Now().UTC().Add(15 * time.Minute)
}
