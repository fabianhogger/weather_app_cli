package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var myWeatherAPIKEY string

// Define the top-level struct for the entire JSON
type WeatherData struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type Forecast struct {
	forecastday []ForecastDay `json:"forecastday"`
}

type ForecastDay struct {
	Date string `json:"date,omitempty"`
	Day  struct {
		MaxTempC          float64   `json:"maxtemp_c"`
		MinTempC          float64   `json:"mintemp_c"`
		AvgTempC          float64   `json:"avgtemp_c"`
		AvgTempF          float64   `json:"avgtemp_f"`
		MaxWindKph        float64   `json:"maxwind_kph"`
		TotalSnowCm       float64   `json:"totalsnow_cm"`
		AvgVisKm          float64   `json:"avgvis_km"`
		AvgHumidity       int       `json:"avghumidity"`
		DailyWillItRain   int       `json:"daily_will_it_rain"`
		DailyChanceOfRain int       `json:"daily_chance_of_rain"`
		DailyWillItSnow   int       `json:"daily_will_it_snow"`
		DailyChanceOfSnow int       `json:"daily_chance_of_snow"`
		Condition         Condition `json:"condition"`
		UV               float64    `json:"uv"`
	} `json:"day,omitempty"`
}

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
}

// Define the struct for the "location" object
type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Day struct {
	MaxtempC          float64 `json:"maxtemp_c"`
	MintempC          float64 `json:"mintemp_c"`
	AvgtempC          float64 `json:"avgtemp_c"`
	AvgtempF          float64 `json:"avgtemp_f"`
	TotalsnowCm       float64 `json:"totalsnow_cm"`
	AvgvisKm          float64 `json:"avgvis_km"`
	Avghumidity       float64 `json:"avghumidity"`
	DailyWillItRain   int     `json:"daily_will_it_rain"`
	DailyChanceOfRain int     `json:"daily_chance_of_rain"`
	DailyWillItSnow   int     `json:"daily_will_it_snow"`
	DailyChanceOfSnow int     `json:"daily_chance_of_snow"`
}





func (d Day) temp() string {
	formated := fmt.Sprintf("MaxtempC: %f\n", d.MaxtempC)
	formated += fmt.Sprintf("MintempC: %f\n", d.MintempC)
	formated += fmt.Sprintf("AvgtempC: %f\n", d.AvgtempC)
	return formated
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
	requestSlice := []string{"http://api.weatherapi.com/v1/forecast.json?key=", myWeatherAPIKEY, "&q=Milan&days=1&aqi=no&alerts=no"}
	url := strings.Join(requestSlice, "")

	res, errRequest := http.Get(url)
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
	fmt.Println(string(body))
	if json_err := json.Unmarshal([]byte(body), &wdata); json_err != nil {
		log.Fatal(json_err)
	}
	// Print the results
	fmt.Println("Location:", wdata.Location.Name, "-", wdata.Location.Country)
	fmt.Println("Temperature:", wdata.Current.TempC, "°C")
	fmt.Println("Condition:", wdata.Current.Condition.Text)
	// Access the forecastday data
	forecast := wdata.Forecast
	if len(forecast.forecastday) == 0 {
		fmt.Println("No forecast days found!")
		return
	}

	for _, day := range forecast.forecastday {
		fmt.Printf("Date: %s\n", day.Date)
		fmt.Printf("Max Temp: %.1f°C\n", day.Day )
		fmt.Printf("Min Temp: %.1f°C\n", day.Day )
		fmt.Printf("Condition: %s\n", day.Day.Condition.Text)
	}
 
}
 