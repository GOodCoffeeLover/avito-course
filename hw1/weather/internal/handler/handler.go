package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"weather/pkg/auth"

	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	inputTimeFormat = time.DateOnly + "T" + time.TimeOnly
	keyTimeFormat   = time.DateOnly + "T15" // 15 - hours
)

type citier interface {
	GetLocationByAddress(addr string) (lat, lng float64, err error)
}

type weatherer interface {
	GetTemperature(lat, lng float64, timestamp time.Time) (temp float64, err error)
}

type errorResponse struct {
	Error string `json:"error"`
}

type response struct {
	City        string  `json:"city"`
	Unit        string  `json:"unit"`
	Temperature float64 `json:"temperature"`
}

func HandleTemperatureRequest(cityClient citier, weatherClient weatherer, redisClient *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		log.Printf("Get request for %v with method %v and params %v", r.URL.Path, r.Method, r.URL.Query())

		rw.Header().Set("Content-Type", "application/json")

		userName := r.URL.Query().Get("user")
		auth, err := checkAuth(userName)

		if err != nil {
			handleError(rw, 500, fmt.Sprintf("Can't auth: %v", err))
			return
		}
		if !auth {
			handleError(rw, 403, "Did not authed")
			return
		}

		city := r.URL.Query().Get("city")
		dt := r.URL.Query().Get("dt")

		if city == "" {
			handleError(rw, 400, "Did get empty city")
			return
		}

		log.Printf("City: %v", city)
		log.Printf("Timestamp: %v", dt)
		if dt == "" {
			dt = time.Now().Format(inputTimeFormat)
		}

		timestamp, err := time.Parse(inputTimeFormat, dt)
		if err != nil {
			handleError(rw, 400, fmt.Sprintf("Can't parse dt: %v", err))
			return
		}

		switch r.Method {
		case "GET":
			getWeather(rw, city, timestamp, redisClient)
		case "PUT":
			saveWeather(rw, city, timestamp, cityClient, weatherClient, redisClient)
		default:
			handleError(rw, 405, fmt.Sprintf("unknown method %v", r.Method))
		}

	}
}

func handleError(w http.ResponseWriter, exitCode int, err string) {
	w.WriteHeader(exitCode)
	errResp := errorResponse{
		Error: fmt.Sprint(err),
	}
	errBody, _ := json.Marshal(errResp)
	w.Write(errBody)
}

func checkAuth(name string) (bool, error) {
	conn, err := grpc.Dial("auth:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, fmt.Errorf("can't connect to auther: %v", err)
	}

	c := auth.NewAutherClient(conn)
	res, err := c.AuthByName(context.Background(), &auth.AuthByNameRequest{Name: name})
	if err != nil {
		return false, fmt.Errorf("error during auth: %v", err)
	}
	return res.GetAuthed(), nil
}

func saveWeather(rw http.ResponseWriter, city string, timestamp time.Time, cityClient citier, weatherClient weatherer, redisClient *redis.Client) {
	lat, lng, err := cityClient.GetLocationByAddress(city)
	if err != nil {
		handleError(rw, 500, fmt.Sprintf("Can't get location: %v", err))
		return
	}

	temp, err := weatherClient.GetTemperature(lat, lng, timestamp)
	if err != nil {
		handleError(rw, 500, fmt.Sprintf("Can't get temperature for city(%v): %v", city, err))
		return
	}

	key := formKey(city, timestamp)
	if err := redisClient.Set(key, temp, 0).Err(); err != nil {
		log.Printf("Can't save weather with key %v: %v", key, err)
		return
	}
	log.Printf("Succesfuly saved temperature(%v) for %v", temp, key)
}

func getWeather(rw http.ResponseWriter, city string, timestamp time.Time, redisClient *redis.Client) {

	key := formKey(city, timestamp)
	resCmd := redisClient.Get(key)
	if err := resCmd.Err(); err != nil {
		handleError(rw, 500, fmt.Sprintf("Can't get weather with key %v: %v", key, err))
		return
	}

	temp, err := resCmd.Float64()
	if err != nil {
		handleError(rw, 404, fmt.Sprintf("Can't find weather for %v: %v", key, err))
		return
	}

	log.Printf("Succesfuly get temperature(%v) for %v", temp, key)

	respBytes, _ := json.Marshal(response{
		City:        city,
		Unit:        "celsius",
		Temperature: temp,
	})
	rw.Write(respBytes)
}

func formKey(city string, timestamp time.Time) string {
	return city + "@" + timestamp.Format(keyTimeFormat)
}
