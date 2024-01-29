package handlers

import (
	"encoding/json"
	"errors"
	"github.com/kova98/spiza/services/simulator/domain"
	"log"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kova98/spiza/services/simulator/adapters"
	"github.com/kova98/spiza/services/simulator/data"
)

const (
	OrderStatusCreated   = 0
	OrderStatusAccepted  = 1
	OrderStatusRejected  = 2
	OrderStatusReady     = 3
	OrderStatusPickedUp  = 4
	OrderStatusDelivered = 5
)

type OrderUpdatedHandler struct {
	l        *log.Logger
	repo     *data.DbRepo
	courier  *domain.Courier
	traveler *adapters.Traveler
}

func NewOrderUpdatedHandler(l *log.Logger, r *data.DbRepo, c *domain.Courier, t *adapters.Traveler) *OrderUpdatedHandler {
	return &OrderUpdatedHandler{
		l:        l,
		repo:     r,
		courier:  c,
		traveler: t,
	}
}

func (h *OrderUpdatedHandler) Handle(client mqtt.Client, mqttMsg mqtt.Message) {
	h.l.Println("Handle MSG order/+")

	orderId, err := parseOrderIdFromTopic(mqttMsg.Topic())
	if err != nil {
		h.l.Println("Error parsing order id from topic:", err)
		return
	}

	if orderId != h.courier.CurrentOrderId {
		// TODO: dynamically subscribe/unsubscribe to topics to avoid this?
		return
	}

	var msg data.OrderUpdated
	err = json.Unmarshal(mqttMsg.Payload(), &msg)
	if err != nil {
		h.l.Println("Unmarshal Error:", err)
		return
	}

	if msg.Status != OrderStatusReady {
		return
	}

	// pick up order

	statusMsg := data.OrderUpdated{
		Status:       OrderStatusPickedUp,
		DeliveryTime: msg.DeliveryTime,
	}
	marshalled, _ := json.Marshal(statusMsg)
	t := client.Publish("order/"+strconv.FormatInt(orderId, 10), 0, false, marshalled)
	_ = t.Done()
	if t.Error() != nil {
		h.l.Println("Error publishing message", statusMsg, "to topic", "order/"+strconv.FormatInt(orderId, 10), ":", t.Error())
		return
	}
	h.l.Println("Order" + strconv.FormatInt(orderId, 10) + " picked up")

	// travel from restaurant to delivery location
	// TODO: think about what happens if this happens before the courier has arrived at the restaurant

	latLng, err := h.repo.GetOrderDestinationLatLng(orderId)
	if err != nil {
		h.l.Println("Error getting order:", err)
		return
	}
	start := h.courier.Loc.ToLatLng()
	path, err := h.traveler.GetPath(start, latLng)
	if err != nil {
		h.l.Println("Error calculating path:", err)
		return
	}
	h.traveler.Travel(orderId, path)

	// complete order
	statusMsg = data.OrderUpdated{
		Status:       OrderStatusDelivered,
		DeliveryTime: msg.DeliveryTime,
	}
	marshalled, _ = json.Marshal(statusMsg)
	t = client.Publish("order/"+strconv.FormatInt(orderId, 10), 0, false, marshalled)
	_ = t.Done()
	if t.Error() != nil {
		h.l.Println("Error publishing message", statusMsg, "to topic", "order/"+strconv.FormatInt(orderId, 10), ":", t.Error())
		return
	}
	h.l.Println("Order" + strconv.FormatInt(orderId, 10) + " completed")
}

func parseOrderIdFromTopic(topic string) (int64, error) {
	parts := strings.Split(topic, "/")
	if len(parts) != 2 {
		return 0, errors.New("invalid topic: " + topic)
	}

	id, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
