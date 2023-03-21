package main

import (
	"log"
	"net/http"
	"os"
	"weather/internal/city"
	"weather/internal/handler"
	"weather/internal/weather"
)

func main() {
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = "7001"
	}
	cityClient := city.New()
	weatherClient := weather.New()

	http.HandleFunc("/forecast", handler.HandleTemperatureRequest(cityClient, weatherClient))
	addr := "0.0.0.0:" + port

	log.Printf("Starting server at %v ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Can't start http server: %v\n", err)
	}
}
