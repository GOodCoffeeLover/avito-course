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
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	middlewarestd "github.com/slok/go-http-metrics/middleware/std"
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
	// Create our middleware.
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	// Our handler.
	hf := http.HandlerFunc(handler.HandleTemperatureRequest(cityClient, weatherClient, redisClient))
	h := middlewarestd.Handler("my_mytrics", mdlw, hf)
	http.Handle("/forecast", h)
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Starting server at %v ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Can't start http server: %v\n", err)
	}
}
