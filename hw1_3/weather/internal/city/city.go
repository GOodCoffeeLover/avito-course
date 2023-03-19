package city

import (
	"fmt"
	"log"

	openstreetmap "github.com/codingsince1985/geo-golang/openstreetmap"
)

type City struct {
}

func New() *City {
	return &City{}
}

func (c *City) GetLocationByAddress(address string) (latitude, longitude float64, err error) {
	coder := openstreetmap.Geocoder()
	loc, err := coder.Geocode(address)
	if err != nil {
		err = fmt.Errorf("error while getting location of %v: %v", address, err)
		return
	}
	if loc == nil {
		err = fmt.Errorf("can't get location of %v", address)
		return
	}
	log.Printf("[city] Get location of %v : (%v, %v)", address, loc.Lat, loc.Lng)
	return loc.Lat, loc.Lng, nil
}
