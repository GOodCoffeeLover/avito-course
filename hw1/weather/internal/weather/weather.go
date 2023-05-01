package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Weather struct {
	apiUrl string
}

func New() *Weather {
	url := os.Getenv("WEATHER_API_URL")
	if url == "" {
		url = "https://api.open-meteo.com/v1/forecast"
	}
	return &Weather{
		apiUrl: url,
	}
}

func (w *Weather) GetTemperature(lat, lng float64, timestamp time.Time) (float64, error) {
	temperatures, err := w.getTemperaturesByHour(lat, lng, timestamp)
	if err != nil {
		return 0.0, fmt.Errorf("can't get temperatures for coordinates (%v, %v) : %v", lat, lng, err)
	}
	if len(temperatures) == 0 {
		return 0.0, fmt.Errorf("can't get temperature for day %v", timestamp.Format(time.DateOnly))
	}
	return temperatures[timestamp.Hour()], nil
}

func (w *Weather) getTemperaturesByHour(lat, lng float64, timestamp time.Time) ([]float64, error) {

	t := timestamp.Format("2006-01-02")
	req := fmt.Sprintf("%v?latitude=%v&longitude=%v&hourly=temperature_2m&start_date=%v&end_date=%v", w.apiUrl, lat, lng, t, t)

	log.Printf("[weatherer] Making request %v", req)

	rsp, err := http.Get(req)
	if err != nil {
		return []float64{}, fmt.Errorf("can't get resp for %v : %v", req, err)
	}

	body_str, err := io.ReadAll(rsp.Body)
	if err != nil {
		return []float64{}, fmt.Errorf("can't read body due to error: %v", err)

	}

	tmps := &struct {
		Hourly struct {
			Temperature_2m []float64 `json:"temperature_2m"`
		} `json:"hourly"`
	}{}

	err = json.Unmarshal(body_str, tmps)
	if err != nil {
		return []float64{}, fmt.Errorf("can't unmarshal body %v due to error: %v", string(body_str), err)
	}
	return tmps.Hourly.Temperature_2m, nil
}
