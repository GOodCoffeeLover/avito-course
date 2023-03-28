package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"weather/pkg/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type temeratureRequestBody map[string]interface{}

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

func HandleTemperatureRequest(c citier, w weatherer) func(w http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("Get request for %v with method %v and params %v", r.URL.Path, r.Method, r.URL.Query())
		rw.Header().Set("Content-Type", "application/json")
		city := r.URL.Query().Get("city")
		dt := r.URL.Query().Get("dt")
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

		if city == "" {
			handleError(rw, 400, "Did get empty city")
			return
		}

		log.Printf("City: %v", city)
		log.Printf("Timestamp: %v", dt)
		if dt == "" {
			dt = time.Now().Format("2006-01-02T03:04:05")
		}

		timestamp, err := time.Parse("2006-01-02T03:04:05", dt)
		if err != nil {
			handleError(rw, 400, fmt.Sprintf("Can't parse dt: %v", err))
			return
		}

		lat, lng, err := c.GetLocationByAddress(city)
		if err != nil {
			handleError(rw, 400, fmt.Sprintf("Can't get location: %v", err))
			return
		}
		temp, err := w.GetTemperature(lat, lng, timestamp)
		if err != nil {
			handleError(rw, 400, fmt.Sprintf("Can't get temperature for city(%v): %v", city, err))
			return
		}
		resp := response{
			City:        city,
			Unit:        "celsius",
			Temperature: temp,
		}
		respBytes, _ := json.Marshal(resp)
		rw.Write(respBytes)
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
