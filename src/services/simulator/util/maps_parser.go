package util

import (
	"encoding/json"
	"github.com/kova98/spiza/services/simulator/data"
	"github.com/twpayne/go-polyline"
)

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

func ParseAndCalculatePath(jsonData []byte) ([]data.Location, error) {
	var response GeoResponse
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}
	var path []data.Location
	for _, route := range response.Routes {
		for _, leg := range route.Legs {
			for _, step := range leg.Steps {
				buf := []byte(step.Polyline.Points)
				coords, _, _ := polyline.DecodeCoords(buf)
				for _, coord := range coords {
					loc := data.Location{Lat: coord[0], Lng: coord[1]}
					path = append(path, loc)
				}
			}
		}
	}
	return path, nil
}
