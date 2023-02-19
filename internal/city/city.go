package ity

import (
	"fmt"

	openstreetmap "github.com/codingsince1985/geo-golang/openstreetmap"
)

type City struct {
	address string
}

func New(address string) City {
	return City{
		address: address,
	}
}

func (c *City) GetLocation() (longitude float64, latitude float64, err error) {
	coder := openstreetmap.Geocoder()
	loc, err := coder.Geocode(c.address)
	if err != nil {
		err = fmt.Errorf("can't get location of %v : %v", c.address, err)
		return
	}
	return loc.Lng, loc.Lat, nil
}

func (c *City) GetAddress() string {
	return c.address
}
