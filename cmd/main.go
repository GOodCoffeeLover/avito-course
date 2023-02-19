package main

import (
	city "avito-course/internal/city"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	c := city.New("Moscow")
	lng, lat, err := c.GetLocation()
	fmt.Printf("Long: %v, Lat: %v, err: %v", lng, lat, err)
}
