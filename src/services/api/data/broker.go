package data

import (
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
)

type Subscription struct {
	topics []string
	conn   *websocket.Conn
}

func (s Subscription) isSubscribed(topic string) bool {
	for _, t := range s.topics {
		if t == topic {
			return true
		}
	}
	return false
}

type Broker struct {
	subs []Subscription
	l    *log.Logger
	bus  mqtt.Client
}

func NewBroker(l *log.Logger) *Broker {
	return &Broker{[]Subscription{}, l, InitBusClient(l)}
}

func (b *Broker) Unsubscribe(conn *websocket.Conn) {
	var newSubs []Subscription
	for _, sub := range b.subs {
		if sub.conn != conn {
			newSubs = append(newSubs, Subscription{sub.topics, sub.conn})
		}
	}
	b.subs = newSubs
}

func (b *Broker) Subscribe(conn *websocket.Conn) {
	// TODO: support multiple topics
	// TODO: support multiple restaurants
	b.subs = append(b.subs, Subscription{[]string{}, conn})
}

func (b *Broker) Publish(topic string, msg interface{}) {
	marshalled, err := json.Marshal(msg)
	if err != nil {
		b.l.Println("Marshal Error:", err)
		return
	}

	// Publish to ws subscribers
	for _, sub := range b.subs {
		// TODO: implement routing and filtering
		err = sub.conn.WriteMessage(websocket.TextMessage, marshalled)
		if err != nil {
			b.l.Println("Write Error:", err)
		}
	}

	// Publish to bus
	t := b.bus.Publish(topic, 0, false, marshalled)
	go func() {
		_ = t.Done()
		if t.Error() != nil {
			b.l.Println("Error publishing message", marshalled, "to topic", topic, ":", t.Error())
		}
	}()
}

func InitBusClient(l *log.Logger) mqtt.Client {
	broker := "localhost"
	port := 1883
	opts := mqtt.NewClientOptions()
	addr := fmt.Sprintf("tcp://%s:%d", broker, port)
	opts.AddBroker(addr)
	opts.SetClientID("spiza_api")
	//opts.SetUsername("")
	//opts.SetPassword("")
	opts.SetDefaultPublishHandler(NewMessagePubHandler(l))
	opts.OnConnect = NewConnectHandler(l, addr)
	opts.OnConnectionLost = NewConnectionLostHandler(l)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func NewMessagePubHandler(l *log.Logger) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		l.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}
}

func NewConnectHandler(l *log.Logger, addr string) mqtt.OnConnectHandler {
	return func(client mqtt.Client) {
		l.Println("Connected to broker on address", addr)
	}
}

func NewConnectionLostHandler(l *log.Logger) mqtt.ConnectionLostHandler {
	return func(client mqtt.Client, err error) {
		l.Println("Connection to broker lost:", err)
	}
}
