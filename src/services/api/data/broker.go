package data

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Broker struct {
	oc   chan OrderWithItems
	subs []*websocket.Conn
	l    *log.Logger
}

func NewBroker(logger *log.Logger) *Broker {
	oc := make(chan OrderWithItems)
	return &Broker{oc, []*websocket.Conn{}, logger}
}

func (b *Broker) Unsubscribe(conn *websocket.Conn) {
	var newSubs []*websocket.Conn
	for _, sub := range b.subs {
		if sub != conn {
			newSubs = append(newSubs, conn)
		}
	}
	b.subs = newSubs
}

func (b *Broker) Subscribe(conn *websocket.Conn) {
	b.subs = append(b.subs, conn)
}

func (b *Broker) Publish(msg interface{}) {
	marshalled, err := json.Marshal(msg)
	for _, sub := range b.subs {
		err = sub.WriteMessage(websocket.TextMessage, marshalled)
		if err != nil {
			b.l.Println("write:", err)
		}
	}
}
