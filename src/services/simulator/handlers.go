package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kova98/spiza/services/simulator/data"
)

func LatLngToLocation(latlng string) Location {
	split := strings.Split(latlng, ",")
	lat, _ := strconv.ParseFloat(split[0], 64)
	lng, _ := strconv.ParseFloat(split[1], 64)
	return Location{Lat: lat, Lng: lng}
}

func NewCourierAssignedHandler(l *log.Logger, repo *data.OrderRepo) mqtt.MessageHandler {
	// currentLoc := Location{
	// 	Lat: 0,
	// 	Lng: 0,
	// }

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

		loc := "45.800169905837784,15.943209331950337"
		path, _ := calculatePath(loc, destLatLng)
		locTopic := "order/" + strconv.FormatInt(msg.OrderId, 10) + "/courier-location"
		for _, loc := range path {
			tempLoc := loc
			json, _ := json.Marshal(loc)
			t := client.Publish(locTopic, 0, false, json)
			go func() {
				_ = t.Done()
				if t.Error() != nil {
					l.Println("Error publishing message", tempLoc, "to topic", locTopic, ":", t.Error())
				}
			}()
			time.Sleep(1 * time.Second)
		}
	}
}

func calculatePath(startLatLng string, endLatLng string) ([]Location, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	directionsApi := "https://maps.googleapis.com/maps/api/directions/json?origin=" + startLatLng +
		"&destination=" + endLatLng +
		"&key=" + apiKey

	response, err := http.Get(directionsApi)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return ParseLocationsFromResponse(body)
}
