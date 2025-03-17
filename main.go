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
	"weather_cli/models"
)

func loadEnv() string {
	path, rerr := os.Getwd()
	if rerr != nil {
		panic(rerr)
	}
	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return os.Getenv("API_KEY")

}

func getData(myWeatherAPIKEY string) ([]byte, error) {
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
	apiKey := loadEnv()
	body, errBody := getData(apiKey)
	if errBody != nil {
		log.Fatal("error reading response")
	}
	var wdata models.WeatherData
	if json_err := json.Unmarshal([]byte(body), &wdata); json_err != nil {
		log.Fatal(json_err)
	}
	// Print the results
	fmt.Println("Location:", wdata.Location.Name, "-", wdata.Location.Country)
	fmt.Println("Temperature:", wdata.Current.TempC, "°C")
	fmt.Println("Condition:", wdata.Current.Condition.Text)

}
