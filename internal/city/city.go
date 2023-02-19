package ity

import (
	"fmt"

	openstreetmap "github.com/codingsince1985/geo-golang/openstreetmap"
)

type Citier struct {
	address string
}

func New(address string) *Citier {
	return &Citier{
		address: address,
	}
}

func (c *Citier) GetLocation() (latitude float64, longitude float64, err error) {
	coder := openstreetmap.Geocoder()
	loc, err := coder.Geocode(c.address)
	if err != nil {
		err = fmt.Errorf("can't get location of %v : %v", c.address, err)
		return
	}
	return loc.Lat, loc.Lng, nil
}

func (c *Citier) GetAddress() string {
	return c.address
}
