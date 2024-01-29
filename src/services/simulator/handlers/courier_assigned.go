package handlers

import (
	"encoding/json"
	"github.com/kova98/spiza/services/simulator/domain"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kova98/spiza/services/simulator/adapters"
	"github.com/kova98/spiza/services/simulator/data"
)

type CourierAssignedHandler struct {
	l        *log.Logger
	repo     data.Repo
	courier  *domain.Courier
	traveler *adapters.Traveler
}

func NewCourierAssignedHandler(logger *log.Logger, repo data.Repo, c *domain.Courier, t *adapters.Traveler) *CourierAssignedHandler {
	return &CourierAssignedHandler{
		l:        logger,
		repo:     repo,
		courier:  c,
		traveler: t,
	}
}

func (h *CourierAssignedHandler) Handle(client mqtt.Client, mqttMsg mqtt.Message) {
	h.l.Println("Handle MSG order/+/courier-assigned")

	var msg data.CourierAssigned
	err := json.Unmarshal(mqttMsg.Payload(), &msg)
	if err != nil {
		h.l.Println("Unmarshal Error:", err)
		return
	}

	// TODO: make this thread safe?
	h.courier.CurrentOrderId = msg.OrderId

	destLatLng, err := h.repo.GetOrderRestaurantLocationLatLng(msg.OrderId)
	if err != nil {
		h.l.Println("Error getting order:", err)
		return
	}

	loc := h.courier.Loc.ToLatLng()
	path, err := h.traveler.GetPath(loc, destLatLng)
	if err != nil {
		h.l.Println("Error calculating path:", err)
		return
	}

	h.traveler.Travel(msg.OrderId, path)
	h.courier.Loc = domain.LatLngToLocation(destLatLng)
}
