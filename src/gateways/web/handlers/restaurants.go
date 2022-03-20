package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kova98/spiza/gateways/web/data"
)

type RestaurantsHandler struct {
	l *log.Logger
}

func NewRestaurantsHandler(l *log.Logger) *RestaurantsHandler {
	return &RestaurantsHandler{l}
}

func (rh *RestaurantsHandler) GetRestaurants(rw http.ResponseWriter, r *http.Request) {
	rh.l.Println("Handle GET Restaurants")

	restaurants, err := data.GetRestaurants()
	if err != nil {
		http.Error(rw, "Unable to get restaurants", http.StatusInternalServerError)
	}

	err = data.ToJSON(restaurants, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (rh *RestaurantsHandler) GetRestaurant(rw http.ResponseWriter, r *http.Request) {
	rh.l.Println("Handle GET Restaurant")

	vars := mux.Vars(r)
	id := vars["id"]

	restaurant, err := data.GetRestaurant(id)
	if err != nil {
		http.Error(rw, "Unable to get restaurant", http.StatusInternalServerError)
	}

	err = data.ToJSON(restaurant, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
