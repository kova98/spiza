package main

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kova98/spiza/services/simulator/data"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func LatLngToLocation(latlng string) Location {
	split := strings.Split(latlng, ",")
	lat, _ := strconv.ParseFloat(split[0], 64)
	lng, _ := strconv.ParseFloat(split[1], 64)
	return Location{Lat: lat, Lng: lng}
}

func NewCourierAssignedHandler(l *log.Logger, repo *data.OrderRepo) mqtt.MessageHandler {
	currentLoc := Location{
		Lat: 0,
		Lng: 0,
	}

	return func(client mqtt.Client, mqttMsg mqtt.Message) {
		l.Println("Handle MSG order/+/courier-assigned")

		var msg CourierAssigned
		err := json.Unmarshal(mqttMsg.Payload(), &msg)
		if err != nil {
			l.Println("Unmarshal Error:", err)
			return
		}

		destLatLng, err := repo.GetOrderDestinationLatLng(msg.OrderId)
		if err != nil {
			l.Println("Error getting order:", err)
			return
		}

		path := calculatePath(currentLoc, destLatLng)
		locTopic := "order/" + strconv.FormatInt(msg.OrderId, 10) + "/courier-location"
		for _, loc := range path {
			currentLoc = loc
			tempLoc := loc
			t := client.Publish(locTopic, 0, false, loc)
			go func() {
				_ = t.Done()
				if t.Error() != nil {
					l.Println("Error publishing message", tempLoc, "to topic", locTopic, ":", t.Error())
				}
			}()
		}
	}
}

func calculatePath(currentLoc Location, orderDest string) []Location {
	return []Location{}
}
