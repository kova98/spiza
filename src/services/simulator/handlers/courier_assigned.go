package handlers

import (
	"github.com/kova98/spiza/services/simulator/domain"
	"log"
)

type CourierAssignedHandler struct {
	l       *log.Logger
	db      domain.Db
	courier *domain.Courier
	maps    domain.Map
}

func NewCourierAssignedHandler(logger *log.Logger, db domain.Db, c *domain.Courier, t domain.Map) *CourierAssignedHandler {
	return &CourierAssignedHandler{
		l:       logger,
		db:      db,
		courier: c,
		maps:    t,
	}
}

func (h *CourierAssignedHandler) Handle(msg domain.CourierAssigned) {
	h.l.Println("Handle Courier Assigned")

	h.courier.AssignToOrder(msg.OrderId)

	dest, err := h.db.GetOrderRestaurantLocation(msg.OrderId)
	if err != nil {
		h.l.Println("Error getting order", msg.OrderId, ":", err)
		//TODO unassign? try again?
		return
	}

	path, err := h.maps.GetPath(h.courier.Loc, dest)
	if err != nil {
		h.l.Println("Error calculating path:", err)
		return
	}

	h.courier.Travel(msg.OrderId, path)
}
