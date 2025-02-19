package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

var myWeatherAPIKEY string

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

// Define the struct for the "current" object
type Current struct {
	LastUpdatedEpoch int       `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PrecipMm         float64   `json:"precip_mm"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelslikeC       float64   `json:"feelslike_c"`
	WindchillC       float64   `json:"windchill_c"`
	HeatindexC       float64   `json:"heatindex_c"`
	DewpointC        float64   `json:"dewpoint_c"`
	VisKm            float64   `json:"vis_km"`
	UV               float64   `json:"uv"`
	GustKph          float64   `json:"gust_kph"`
}

// Define the struct for the "location" object
type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

// Define the top-level struct for the entire JSON
type WeatherData struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

func loadEnv() {
	path, rerr := os.Getwd()
	if rerr != nil {
		panic(rerr)
	}
	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	myWeatherAPIKEY = os.Getenv("API_KEY")

}

func getData() ([]byte, error) {
	res, errRequest := http.Get("http://api.weatherapi.com/v1/current.json?key=" + myWeatherAPIKEY + "&q=London&aqi=no")
	body, errBody := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatal("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if errRequest != nil {
		return nil, errRequest
	}
	if errBody != nil {
		return nil, errBody
	}
	res.Body.Close()
	return body, nil
}

func main() {
	var wdata WeatherData
	loadEnv()
	body, errBody := getData()
	if errBody != nil {
		log.Fatal("error reading response")
	}

	if json_err := json.Unmarshal(body, &wdata); json_err != nil {
		log.Fatal(json_err)
	}
	fmt.Println("%s", wdata.Location)
}
