package city

import (
	"github.com/codingsince1985/geo-golang"
)

type City struct{
	address string 
}

func New(address string) City {
	return City{
		address: address
	}
}