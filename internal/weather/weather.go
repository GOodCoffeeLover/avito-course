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

type Weatherer struct {
	apiUrl string
}

func New() *Weatherer {
	url := os.Getenv("WEATHER_API_URL")
	if url == "" {
		url = "api.open-meteo.com/v1/forecast"
	}
	url = "https://" + url
	return &Weatherer{
		apiUrl: url,
	}
}

func (w *Weatherer) GetTemperature(lat, lng float64, timestamp time.Time) (float64, error) {
	temperatures, err := w.getTemperaturesByHour(lat, lng, timestamp)
	if err != nil {
		return 0.0, fmt.Errorf("can't get temperatures for (%v, %v) : %v", lat, lng, err)
	}
	return temperatures[timestamp.Hour()], nil
}

func (w *Weatherer) getTemperaturesByHour(lat, lng float64, timestamp time.Time) ([]float64, error) {

	t := timestamp.Format("2006-01-02")
	req := fmt.Sprintf("%v?latitude=%v&longitude=%v&hourly=temperature_2m&start_date=%v&end_date=%v", w.apiUrl, lat, lng, t, t)

	log.Printf("[weatherer] making request %v", req)

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
