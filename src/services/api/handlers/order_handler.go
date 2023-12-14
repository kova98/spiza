package handlers

import (
	"github.com/kova98/spiza/services/api/data"
	"log"
	"net/http"
)

type OrderHandler struct {
	l    *log.Logger
	repo *data.OrderRepo
}

func NewOrderHandler(l *log.Logger, repo *data.OrderRepo) *OrderHandler {
	return &OrderHandler{l, repo}
}

func (oh *OrderHandler) CreateOrder(rw http.ResponseWriter, r *http.Request) {
	oh.l.Println("Handle POST Order")

	order := &data.Order{}
	err := data.FromJSON(order, r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	id, err := oh.repo.CreateOrder(order)
	if err != nil {
		oh.l.Println(err)
		http.Error(rw, "Unable to create order", http.StatusInternalServerError)
		return
	}

	order.Id = id
	err = data.ToJSON(order, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
