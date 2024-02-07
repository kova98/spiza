package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WebSocketAdapter struct {
	l       *log.Logger
	state   *State
	clients map[*websocket.Conn]bool
}

func NewWebsocketAdapter(l *log.Logger, state *State) *WebSocketAdapter {
	var clients = make(map[*websocket.Conn]bool)
	return &WebSocketAdapter{l, state, clients}
}

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func (ws *WebSocketAdapter) HandleWebsocketConnection(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ws.l.Println("Error upgrading to websocket: ", err)
		return
	}
	ws.clients[conn] = true
}

// TODO: move to order handler
func (ws *WebSocketAdapter) HandleOrderCreated(o Order) {
	ws.l.Println("Handle order created")
	msg := Message{Type: "OrderCreated", Data: o}
	ws.Broadcast(msg)
}

func (ws *WebSocketAdapter) HandleOrderUpdated(ou OrderUpdated) {
	ws.l.Println("Handle order updated")
	msg := Message{Type: "OrderUpdated", Data: ou}
	ws.Broadcast(msg)
}

func (ws *WebSocketAdapter) HandleCourierAssigned(ca CourierAssigned) {
	ws.l.Println("Handle courier assigned")
	msg := Message{Type: "CourierAssigned", Data: ca}
	ws.Broadcast(msg)
}

func (ws *WebSocketAdapter) HandleCourierLocationUpdated(clu CourierLocationUpdated) {
	ws.l.Println("Handle courier location updated")
	msg := Message{Type: "CourierLocationUpdated", Data: clu}
	ws.Broadcast(msg)
}

func (ws *WebSocketAdapter) Broadcast(msg Message) {
	marsh, err := json.Marshal(msg)
	if err != nil {
		ws.l.Println("Error marshalling message: ", err)
		return
	}
	for client := range ws.clients {
		err := client.WriteMessage(websocket.TextMessage, marsh)
		if err != nil {
			log.Printf("Error sending message to client: %v", err)
			client.Close()
			delete(ws.clients, client)
		}
	}
}
