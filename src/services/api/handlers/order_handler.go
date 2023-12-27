package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/kova98/spiza/services/api/data"
)

type OrderWithItems struct {
	Id           int64       `json:"id"`
	UserId       int64       `json:"user_id"`
	RestaurantId int64       `json:"restaurant_id" `
	Status       int64       `json:"status"`
	Items        []data.Item `json:"items" `
}

type OrderHandler struct {
	l        *log.Logger
	repo     *data.OrderRepo
	itemRepo *data.ItemRepo
	broker   *data.Broker
}

type UpdateOrderMsg struct {
	Id     int64
	Action string
}

const (
	OrderStatusCreated   = 0
	OrderStatusAccepted  = 1
	OrderStatusRejected  = 2
	OrderStatusReady     = 3
	OrderStatusPickedUp  = 4
	OrderStatusDelivered = 5
)

func NewOrderHandler(l *log.Logger, or *data.OrderRepo, ir *data.ItemRepo, b *data.Broker) *OrderHandler {
	return &OrderHandler{l, or, ir, b}
}

func (oh *OrderHandler) CreateOrder(rw http.ResponseWriter, r *http.Request) {
	oh.l.Println("Handle POST Order")

	order := data.Order{}
	err := data.FromJSON(&order, r.Body)
	if err != nil {
		oh.l.Println(err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	createdOrder, err := oh.repo.CreateOrder(&order)
	if err != nil {
		oh.l.Println(err)
		http.Error(rw, "Unable to create order", http.StatusInternalServerError)
		return
	}

	items, err := oh.itemRepo.GetByOrder(order.Id)
	if err != nil {
		oh.l.Println(err)
		return
	}
	topic := "order/" + strconv.FormatInt(order.Id, 10) + "/created"
	oh.broker.Publish(topic, createdOrder.WithItems(items))

	err = data.ToJSON(createdOrder, rw)
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

	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		oh.l.Print("upgrade: ", err)
		return
	}
	defer conn.Close()

	oh.broker.Subscribe(conn)
	defer oh.broker.Unsubscribe(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			oh.l.Println("read:", err)
			break
		}
		var msg UpdateOrderMsg
		if err := json.Unmarshal(message, &msg); err != nil {
			oh.l.Println(err)
			continue
		}
		if err := updateOrderStatus(oh, msg.Id, msg.Action); err != nil {
			oh.l.Println(err)
			continue
		}
	}
}

func updateOrderStatus(oh *OrderHandler, id int64, action string) error {
	oh.l.Println("Updating order", id, "to status", action)
	status, err := getStatus(action)
	if err != nil {
		return err
	}
	if err := oh.repo.UpdateOrderStatus(id, status); err != nil {
		return err
	}
	topic := "order/" + strconv.FormatInt(id, 10)

	deliveryTime := time.Now().UTC().Add(time.Hour)
	oh.broker.Publish(topic, data.OrderStatusUpdated{Status: status, DeliveryTime: deliveryTime})
	return nil
}

func getStatus(action string) (int, error) {
	switch action {
	case "accept":
		return OrderStatusAccepted, nil
	case "reject":
		return OrderStatusRejected, nil
	default:
		return 0, fmt.Errorf("Invalid action:" + action)
	}
}
