package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)
var myWeatherAPIKEY string

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key="+myWeatherAPIKEY+"&q=London&aqi=no")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}