package domain

import (
	"strconv"
	"strings"
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

func (l Location) ToLatLng() string {
	return strconv.FormatFloat(l.Lat, 'f', 6, 64) + "," + strconv.FormatFloat(l.Lng, 'f', 6, 64)
}
