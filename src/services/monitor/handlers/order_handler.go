package handlers

import (
	"github.com/kova98/spiza/services/monitor/domain"
	"log"
)

type OrderHandler struct {
	l *log.Logger
	s *domain.State
}

func NewOrderHandler(l *log.Logger, s *domain.State) *OrderHandler {
	return &OrderHandler{l, s}
}

func (h *OrderHandler) HandleOrderCreated(msg domain.Order) {
	h.l.Println("Handle order created")
	h.s.ActiveOrders = append(h.s.ActiveOrders, msg)
}

func (h *OrderHandler) HandleOrderUpdated(msg domain.OrderUpdated) {
	h.l.Println("Handle order updated")
	for i, o := range h.s.ActiveOrders {
		if o.Id == msg.Id {
			h.s.ActiveOrders[i].Status = msg.Status
		}
	}
}

func (h *OrderHandler) HandleCourierAssigned(msg domain.CourierAssigned) {
	h.l.Println("Handle courier assigned")
	for i, o := range h.s.ActiveOrders {
		if o.Id == msg.OrderId {
			h.s.ActiveOrders[i].CourierId = msg.CourierId
		}
	}
}
