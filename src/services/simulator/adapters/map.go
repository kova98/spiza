package adapters

import (
	"encoding/json"
	"github.com/kova98/spiza/services/simulator/domain"
	"github.com/twpayne/go-polyline"
	"io"
	"log"
	"net/http"
)

type GoogleMaps struct {
	l      *log.Logger
	apiKey string
}

func NewGoogleMaps(logger *log.Logger, apiKey string) *GoogleMaps {
	return &GoogleMaps{
		l:      logger,
		apiKey: apiKey,
	}
}

func (t *GoogleMaps) GetPath(start domain.Location, dest domain.Location) ([]domain.Location, error) {
	directionsUri := "https://maps.googleapis.com/maps/api/directions/json?origin=" + start.ToLatLng() +
		"&destination=" + dest.ToLatLng() +
		"&key=" + t.apiKey

	response, err := http.Get(directionsUri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return parseAndCalculatePath(body)
}

func parseAndCalculatePath(jsonData []byte) ([]domain.Location, error) {
	var response GeoResponse
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}
	var path []domain.Location
	for _, route := range response.Routes {
		for _, leg := range route.Legs {
			for _, step := range leg.Steps {
				buf := []byte(step.Polyline.Points)
				coords, _, _ := polyline.DecodeCoords(buf)
				for _, coord := range coords {
					loc := domain.Location{Lat: coord[0], Lng: coord[1]}
					path = append(path, loc)
				}
			}
		}
	}
	return path, nil
}

type GeoResponse struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	Legs []Leg `json:"legs"`
}

type Leg struct {
	Steps []Step `json:"steps"`
}

type Step struct {
	Polyline Polyline `json:"polyline"`
}

type Polyline struct {
	Points string `json:"points"`
}
