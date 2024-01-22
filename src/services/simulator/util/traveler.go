package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kova98/spiza/services/simulator/data"
)

type Traveler struct {
	l      *log.Logger
	client mqtt.Client
}

func NewTraveler(logger *log.Logger, client mqtt.Client) *Traveler {
	return &Traveler{
		l:      logger,
		client: client,
	}
}

func (t *Traveler) Travel(orderId int64, path []data.Location) {
	locTopic := "order/" + strconv.FormatInt(orderId, 10) + "/courier-location"
	for _, loc := range path {
		tempLoc := loc
		json, _ := json.Marshal(loc)
		token := t.client.Publish(locTopic, 0, false, json)
		_ = token.Done()
		if token.Error() != nil {
			t.l.Println("Error publishing message", tempLoc, "to topic", locTopic, ":", token.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

func (t *Traveler) CalculatePath(startLatLng string, endLatLng string) ([]data.Location, error) {
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

	return ParseAndCalculatePath(body)
}
