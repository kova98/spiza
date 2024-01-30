package domain_test

import (
	"github.com/kova98/spiza/services/simulator/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLatLngToLocation(t *testing.T) {
	latlng := " 11.22222, 33.44444 "

	loc := domain.LatLngToLocation(latlng)

	assert.Equal(t, 11.22222, loc.Lat)
	assert.Equal(t, 33.44444, loc.Lng)
}

func TestToLatLng(t *testing.T) {
	loc := domain.Location{Lat: 11.22222, Lng: 33.44444}

	latlng := loc.ToLatLng()

	assert.Equal(t, "11.222220,33.444440", latlng)
}
