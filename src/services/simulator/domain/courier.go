package domain

import (
	"log"
	"strconv"
	"time"
)

type Courier struct {
	Id             int64
	Name           string
	Loc            Location
	CurrentOrderId int64
	bus            Bus
	l              *log.Logger
}

const StepDurationInMiliseconds = 500

func NewCourier(id int64, name string, loc Location, bus Bus, l *log.Logger) *Courier {
	return &Courier{
		Id:             id,
		Name:           name,
		Loc:            loc,
		CurrentOrderId: 0,
		bus:            bus,
		l:              l,
	}
}

func (c *Courier) AssignToOrder(orderId int64) {
	// TODO: make this thread safe?
	c.CurrentOrderId = orderId
}

func (c *Courier) Travel(orderId int64, path []Location) {
	locTopic := "order/" + strconv.FormatInt(orderId, 10) + "/courier-location"
	for _, loc := range path {
		c.bus.Publish(locTopic, loc)
		time.Sleep(StepDurationInMiliseconds * time.Millisecond)
	}
	c.Loc = path[len(path)-1]
}

func (c *Courier) PickUpOrder(id int64) {
	statusMsg := OrderUpdated{
		Status:       OrderStatusPickedUp,
		DeliveryTime: calculateDeliveryTime(),
	}
	c.bus.Publish("order/"+strconv.FormatInt(id, 10), statusMsg)
	c.l.Println("Order" + strconv.FormatInt(id, 10) + " picked up")
}

func (c *Courier) CompleteOrder(id int64, msg OrderUpdated) {
	// complete order
	statusMsg := OrderUpdated{
		Status:       OrderStatusDelivered,
		DeliveryTime: msg.DeliveryTime,
	}
	c.bus.Publish("order/"+strconv.FormatInt(id, 10), statusMsg)
	c.l.Println("Order" + strconv.FormatInt(id, 10) + " completed")
}

func calculateDeliveryTime() time.Time {
	return time.Now().UTC().Add(15 * time.Minute)
}
