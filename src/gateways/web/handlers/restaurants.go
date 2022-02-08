package handlers

import (
	"log"
	"net/http"

	"github.com/kova98/spiza/gateways/web/data"
)

type RestaurantsHandler struct {
	l *log.Logger
}

func NewRestaurantsHandler(l *log.Logger) *RestaurantsHandler {
	return &RestaurantsHandler{l}
}

func (rh *RestaurantsHandler) GetRestaurants(rw http.ResponseWriter, r *http.Request) {
	rh.l.Println("Handle GET Products")

	restaurants := data.GetRestaurants()

	// serialize the list to JSON
	err := data.ToJSON(restaurants, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
