package main

import (
	"log"
	"net/http"
	"os"
	"weather/internal/city"
	"weather/internal/handler"
	"weather/internal/weather"

	"github.com/go-redis/redis"
)

func main() {
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = "7001"
	}
	addr := "0.0.0.0:" + port

	cityClient := city.New()
	weatherClient := weather.New()
	// redisClient := redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs: []string{
	// 		"redis-1:6379",
	// 		"redis-2:6379",
	// 		"redis-3:6379",
	// 		"redis-4:6379",
	// 		"redis-5:6379",
	// 		"redis-6:6379",
	// 	},
	// })
	redisClient := redis.NewClient(&redis.Options{
		Addr: "twem-proxy:16379",
	})
	if err := redisClient.Ping().Err(); err != nil {
		log.Fatalf("Can't ping redis client: %v", err)
	}

	http.HandleFunc("/forecast", handler.HandleTemperatureRequest(cityClient, weatherClient, redisClient))

	log.Printf("Starting server at %v ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Can't start http server: %v\n", err)
	}
}
