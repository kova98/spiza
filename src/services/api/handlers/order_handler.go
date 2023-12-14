package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/kova98/spiza/services/api/data"
	"log"
	"net/http"
	"strconv"
)

type OrderWithItems struct {
	Id           int64       `json:"id"`
	UserId       int64       `json:"user_id"`
	RestaurantId int64       `json:"restaurant_id" `
	Status       int64       `json:"status"`
	Items        []data.Item `json:"items" `
}

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
		oh.l.Println(err)
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
		oh.l.Println(err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (oh *OrderHandler) GetOrders(rw http.ResponseWriter, r *http.Request) {
	oh.l.Println("Handle GET Orders")

	vars := mux.Vars(r)
	idString := vars["id"]
	restaurantId, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	orders, err := oh.repo.GetOrders(restaurantId)
	if err != nil {
		oh.l.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := data.ToJSON(orders, rw); err != nil {
		oh.l.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

func (oh *OrderHandler) HandleOrderWebSocket(rw http.ResponseWriter, r *http.Request) {
	oh.l.Println("Handle GET OrderWebSocket")

	vars := mux.Vars(r)
	idString := vars["id"]
	restaurantId, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		oh.l.Print("upgrade: ", err)
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		orders, err := oh.repo.GetOrders(restaurantId)
		if err != nil {
			log.Println(err)
			break
		}
		msg, err := json.Marshal(orders)
		err = conn.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
