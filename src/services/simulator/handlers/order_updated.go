package handlers

import (
	"github.com/kova98/spiza/services/simulator/domain"
	"log"
)

type OrderUpdatedHandler struct {
	l       *log.Logger
	db      domain.Db
	courier *domain.Courier
	maps    domain.Map
}

func NewOrderUpdatedHandler(l *log.Logger, db domain.Db, c *domain.Courier, m domain.Map) *OrderUpdatedHandler {
	return &OrderUpdatedHandler{
		l:       l,
		db:      db,
		courier: c,
		maps:    m,
	}
}

func (h *OrderUpdatedHandler) Handle(msg domain.OrderUpdated) {
	h.l.Println("Handle MSG order/+")

	if msg.Id != h.courier.CurrentOrderId {
		// TODO: dynamically subscribe/unsubscribe to topics to avoid this?
		return
	}

	if msg.Status != domain.OrderStatusReady {
		// TODO: handle other statuses
		return
	}

	h.courier.PickUpOrder(msg.Id)

	dest, err := h.db.GetOrderDestinationLatLng(msg.Id)
	if err != nil {
		h.l.Println("Error getting order:", err)
		return
	}
	path, err := h.maps.GetPath(h.courier.Loc, dest)
	if err != nil {
		h.l.Println("Error calculating path:", err)
		return
	}
	h.courier.Travel(msg.Id, path)
	h.courier.CompleteOrder(msg.Id, msg)
}
