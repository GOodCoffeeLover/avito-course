package main

import (
	city "avito-course/internal/city"
	"avito-course/internal/handler"
	"avito-course/internal/weather"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = "7001"
	}
	cityClient := city.New()
	weatherClient := weather.New()
	http.HandleFunc("/forecast", handler.HandleTemperatureRequest(cityClient, weatherClient))

	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		log.Fatalf("Can't start http server: %v", err)
	}
}
