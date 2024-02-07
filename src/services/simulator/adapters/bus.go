package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kova98/spiza/services/simulator/domain"
	"log"
	"strconv"
	"strings"
)

type MqttBus struct {
	l      *log.Logger
	Client mqtt.Client
}

func NewMqttBus(l *log.Logger) *MqttBus {
	return &MqttBus{l: l, Client: initClient(l)}
}

func (b *MqttBus) Publish(topic string, msg interface{}) {
	msgJson, _ := json.Marshal(msg)
	token := b.Client.Publish(topic, 0, false, msgJson)
	_ = token.Done()
	if token.Error() != nil {
		b.l.Println("Error publishing message", msgJson, "to topic", topic, ":", token.Error())
	}
	b.l.Println("Published to", topic, ":", msg)
}

func (b *MqttBus) SubscribeCourierAssigned(handle func(msg domain.CourierAssigned)) {
	topic := "order/+/courier-assigned"
	token := b.Client.Subscribe(topic, 0, func(client mqtt.Client, mqttMsg mqtt.Message) {
		var msg domain.CourierAssigned
		err := json.Unmarshal(mqttMsg.Payload(), &msg)
		if err != nil {
			b.l.Println("Unmarshal Error:", err)
			return
		}
		handle(msg)
	})
	_ = token.Done()
	if token.Error() != nil {
		b.l.Println("Error subscribing to topic", topic, ":", token.Error())
	}
}

func (b *MqttBus) SubscribeOrderUpdated(handle func(msg domain.OrderUpdated)) {
	topic := "order/+"
	token := b.Client.Subscribe(topic, 0, func(client mqtt.Client, mqttMsg mqtt.Message) {
		var msg domain.OrderUpdated
		err := json.Unmarshal(mqttMsg.Payload(), &msg)
		if err != nil {
			b.l.Println("Unmarshal Error:", err)
			return
		}

		orderId, err := parseOrderIdFromTopic(mqttMsg.Topic())
		if err != nil {
			b.l.Println("Error parsing order id from topic:", err)
			return
		}
		msg.Id = orderId

		handle(msg)
	})
	_ = token.Done()
	if token.Error() != nil {
		b.l.Println("Error subscribing to topic", topic, ":", token.Error())
	}
}

func parseOrderIdFromTopic(topic string) (int64, error) {
	parts := strings.Split(topic, "/")
	if len(parts) != 2 {
		return 0, errors.New("invalid topic: " + topic)
	}

	id, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func initClient(l *log.Logger) mqtt.Client {
	broker := "localhost"
	port := 1883
	opts := mqtt.NewClientOptions()
	addr := fmt.Sprintf("tcp://%s:%d", broker, port)
	opts.AddBroker(addr)
	opts.SetClientID("simulator")
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
