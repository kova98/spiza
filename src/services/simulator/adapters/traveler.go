package adapters

import (
	"encoding/json"
	"github.com/kova98/spiza/services/simulator/domain"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Traveler struct {
	l      *log.Logger
	client mqtt.Client
	apiKey string
}

func NewTraveler(logger *log.Logger, client mqtt.Client, apiKey string) *Traveler {
	return &Traveler{
		l:      logger,
		client: client,
		apiKey: apiKey,
	}
}

func (t *Traveler) Travel(orderId int64, path []domain.Location) {
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

func (t *Traveler) GetPath(startLatLng string, endLatLng string) ([]domain.Location, error) {
	directionsApi := "https://maps.googleapis.com/maps/api/directions/json?origin=" + startLatLng +
		"&destination=" + endLatLng +
		"&key=" + t.apiKey

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
