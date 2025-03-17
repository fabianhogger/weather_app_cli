package models 
// Define the top-level struct for the entire JSON
type WeatherData struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

// Forecast represents the "forecast" object in the JSON
type Forecast struct {
	Forecastday []ForecastDay `json:"forecastday"` // Ensure this matches the JSON key
}

// ForecastDay represents a single day's forecast in the JSON
type ForecastDay struct {
	Date string `json:"date"`
	Day  Day    `json:"day"`
}

// Day represents the "day" object in the JSON
type Day struct {
	MaxtempC  float64   `json:"maxtemp_c"`
	MintempC  float64   `json:"mintemp_c"`
	Condition Condition `json:"condition"`
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

 