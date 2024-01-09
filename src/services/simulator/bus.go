package main

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Bus struct {
	Client mqtt.Client
	l      *log.Logger
}

func NewBus(l *log.Logger) *Bus {
	return &Bus{InitBusClient(l), l}
}

func InitBusClient(l *log.Logger) mqtt.Client {
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
