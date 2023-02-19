package main

import (
	city "avito-course/internal/city"
	"avito-course/internal/weather"
	"time"

	"fmt"
)

func main() {
	c := city.New("Moscow")
	lat, lng, err := c.GetLocation()
	fmt.Printf("Long: %v, Lat: %v, err: %v\n", lng, lat, err)
	w := weather.New()
	tmp, err := w.GetTemperature(lat, lng, time.Now())
	fmt.Printf("Temperature: %v, err: %v", tmp, err)
}
