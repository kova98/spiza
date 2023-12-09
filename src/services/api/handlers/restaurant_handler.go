package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kova98/spiza/services/api/data"
)

type RestaurantHandler struct {
	logger *log.Logger
	repo   *data.RestaurantRepo
}

func NewRestaurantsHandler(l *log.Logger, r *data.RestaurantRepo) *RestaurantHandler {
	return &RestaurantHandler{l, r}
}

func (rh *RestaurantHandler) GetRestaurants(rw http.ResponseWriter, r *http.Request) {
	rh.logger.Println("Handle GET Restaurants")

	restaurants, err := rh.repo.GetRestaurants()
	if err != nil {
		http.Error(rw, "Unable to get restaurants", http.StatusInternalServerError)
	}

	err = data.ToJSON(restaurants, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (rh *RestaurantHandler) GetRestaurant(rw http.ResponseWriter, r *http.Request) {
	rh.logger.Println("Handle GET Restaurant")

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid id: "+idString, http.StatusInternalServerError)
	}

	restaurant, err := rh.repo.GetRestaurant(id)
	if err != nil {
		http.Error(rw, "Unable to get restaurant", http.StatusInternalServerError)
	}

	err = data.ToJSON(restaurant, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (rh *RestaurantHandler) CreateRestaurant(rw http.ResponseWriter, r *http.Request) {
	rh.logger.Println("Handle POST Restaurant")
	restaurant := &data.Restaurant{}
	err := data.FromJSON(restaurant, r.Body)
	if err != nil {
		http.Error(rw, "Error reading request body", http.StatusInternalServerError)
		return
	}

	id, err := rh.repo.CreateRestaurant(restaurant)
	if err != nil {
		http.Error(rw, "Error creating restaurant: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)

	response := struct {
		Id int64 `json:"id"`
	}{Id: id}
	data.ToJSON(response, rw)
}

func (rh *RestaurantHandler) DeleteRestaurant(rw http.ResponseWriter, r *http.Request) {
	rh.logger.Println("Handle DELETE Restaurant")
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid id: "+idString, http.StatusInternalServerError)
	}

	err = rh.repo.DeleteRestaurant(id)
	if err != nil {
		http.Error(rw, "Error deleting restaurant: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
