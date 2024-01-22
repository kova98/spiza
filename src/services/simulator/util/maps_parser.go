package util

import (
	"encoding/json"
	"math"

	"github.com/kova98/spiza/services/simulator/data"
)

type GoogleMapsResponse struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	Legs []Leg `json:"legs"`
}

type Leg struct {
	Steps []Step `json:"steps"`
}

type Step struct {
	StartLocation data.Location `json:"start_location"`
	EndLocation   data.Location `json:"end_location"`
}

func lerp(start, end data.Location, t float64) data.Location {
	lat := start.Lat + (end.Lat-start.Lat)*t
	lng := start.Lng + (end.Lng-start.Lng)*t
	return data.Location{Lat: lat, Lng: lng}
}

func distance(loc1, loc2 data.Location) float64 {
	// Radius of the Earth in kilometers
	const earthRadius = 6371.0

	// Convert latitude and longitude from degrees to radians
	lat1 := loc1.Lat * (math.Pi / 180.0)
	lng1 := loc1.Lng * (math.Pi / 180.0)
	lat2 := loc2.Lat * (math.Pi / 180.0)
	lng2 := loc2.Lng * (math.Pi / 180.0)

	// Haversine formula
	dlat := lat2 - lat1
	dlng := lng2 - lng1
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}

func ParseAndCalculatePath(jsonData []byte) ([]data.Location, error) {
	var response GoogleMapsResponse
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}

	var smoothPath []data.Location

	for _, route := range response.Routes {
		for _, leg := range route.Legs {
			for i, step := range leg.Steps {
				smoothPath = append(smoothPath, step.StartLocation)
				if i < len(leg.Steps)-1 {
					// Calculate the number of steps based on distance
					numSteps := 10
					distance := distance(step.StartLocation, step.EndLocation)
					if distance > 0 {
						numSteps = int(distance / 0.1) // Adjust the 0.01 for desired step size
					}

					for j := 1; j <= numSteps; j++ {
						t := float64(j) / float64(numSteps)
						interpolatedPoint := lerp(step.StartLocation, step.EndLocation, t)
						smoothPath = append(smoothPath, interpolatedPoint)
					}
				}
			}
		}
	}
	return smoothPath, nil
}
