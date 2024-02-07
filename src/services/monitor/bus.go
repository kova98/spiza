package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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

const QosAtLeastOnce = 1

func (b *MqttBus) SubscribeOrderUpdated(handle func(msg OrderUpdated)) {
	topic := "order/+"
	token := b.Client.Subscribe(topic, QosAtLeastOnce, func(client mqtt.Client, mqttMsg mqtt.Message) {
		var msg OrderUpdated
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
	b.l.Println("Subscribed to topic", topic)
}

func (b *MqttBus) SubscribeOrderCreated(handle func(msg Order)) {
	topic := "order/+/created"
	token := b.Client.Subscribe(topic, QosAtLeastOnce, func(client mqtt.Client, mqttMsg mqtt.Message) {
		var msg Order
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
	b.l.Println("Subscribed to topic", topic)
}

func (b *MqttBus) SubscribeCourierAssigned(handle func(msg CourierAssigned)) {
	topic := "order/+/courier-assigned"
	token := b.Client.Subscribe(topic, QosAtLeastOnce, func(client mqtt.Client, mqttMsg mqtt.Message) {
		var msg CourierAssigned
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
	b.l.Println("Subscribed to topic", topic)
}

func (b *MqttBus) SubscribeCourierLocationUpdated(handle func(msg CourierLocationUpdated)) {
	topic := "order/+/courier-location"
	token := b.Client.Subscribe(topic, QosAtLeastOnce, func(client mqtt.Client, mqttMsg mqtt.Message) {
		var msg CourierLocationUpdated
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
		msg.OrderId = orderId

		handle(msg)
	})
	_ = token.Done()
	if token.Error() != nil {
		b.l.Println("Error subscribing to topic", topic, ":", token.Error())
	}
	b.l.Println("Subscribed to topic", topic)
}

func parseOrderIdFromTopic(topic string) (int64, error) {
	parts := strings.Split(topic, "/")
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
	opts.SetClientID("monitor")
	opts.SetDefaultPublishHandler(NewMessagePubHandler(l))
	opts.OnConnect = NewConnectHandler(l, addr)
	opts.OnConnectionLost = NewConnectionLostHandler(l)
	opts.SetKeepAlive(5)
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