package main

import (
	"log"
	"net/http"
	"os"
	"weather/internal/city"
	"weather/internal/handler"
	"weather/internal/weather"

	"github.com/go-redis/redis"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = "7001"
	}
	addr := "0.0.0.0:" + port

	cityClient := city.New()
	weatherClient := weather.New()
	redisClient := redis.NewClient(&redis.Options{
		Addr: "twem-proxy:16379",
	})
	if err := redisClient.Ping().Err(); err != nil {
		log.Fatalf("Can't ping redis client: %v", err)
	}

	http.HandleFunc("/forecast", handler.HandleTemperatureRequest(cityClient, weatherClient, redisClient))
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Starting server at %v ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Can't start http server: %v\n", err)
	}
}
