package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Feature represents the top-level object in the JSON.
type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

// Geometry represents the geographical coordinates.
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"` // Assuming all coordinates are float64
}

// Properties contains metadata and the timeseries data.
type Properties struct {
	Meta       Meta        `json:"meta"`
	Timeseries []Timeserie `json:"timeseries"`
}

// Meta contains updated_at timestamp and units of measurement.
type Meta struct {
	UpdatedAt string            `json:"updated_at"`
	Units     map[string]string `json:"units"`
}

// Timeserie represents each entry in the timeseries array.
type Timeserie struct {
	Time string `json:"time"`
	Data Data   `json:"data"`
}

// Data contains all the weather data points: instant, next_12_hours, etc.
type Data struct {
	Instant     InstantDetail   `json:"instant"`
	Next12Hours ForecastSummary `json:"next_12_hours"`
	Next1Hours  ForecastSummary `json:"next_1_hours"`
	Next6Hours  ForecastSummary `json:"next_6_hours"`
}

// InstantDetail contains the immediate weather details.
type InstantDetail struct {
	Details WeatherDetails `json:"details"`
}

// ForecastSummary contains the summary and details for forecasts.
type ForecastSummary struct {
	Summary SymbolCode         `json:"summary"`
	Details map[string]float64 `json:"details"` // Assuming all detail values are float64
}

// WeatherDetails holds the specific weather measurement details.
type WeatherDetails struct {
	AirPressureAtSeaLevel float64 `json:"air_pressure_at_sea_level"`
	AirTemperature        float64 `json:"air_temperature"`
	CloudAreaFraction     float64 `json:"cloud_area_fraction"`
	RelativeHumidity      float64 `json:"relative_humidity"`
	WindFromDirection     float64 `json:"wind_from_direction"`
	WindSpeed             float64 `json:"wind_speed"`
	PrecipitationAmount   float64 `json:"precipitation_amount,omitempty"` // Optional, may not be present
}

// SymbolCode contains the symbol code for weather condition summary.
type SymbolCode struct {
	SymbolCode string `json:"symbol_code"`
}

type WeatherResponse struct {
	Type       string     `json:"type"`
	Features   []Feature  `json:"features"`
	Properties Properties `json:"properties"`
}

func GetWeatherFromCoords(lat float64, lon float64) (Data, error) {
	url := fmt.Sprintf("https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=%f&lon=%f", lat, lon)

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Data{}, fmt.Errorf("Error creating request: %s", err)
	}

	// Set a unique User-Agent header
	req.Header.Set("User-Agent", "MyUniqueWeatherApp/1.0")

	// Perform the HTTP request using our client
	resp, err := client.Do(req)
	if err != nil {
		return Data{}, fmt.Errorf("Error getting weather: %s", err)
	}
	defer resp.Body.Close()

	// Decode the JSON response into our WeatherResponse struct
	var weatherResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return Data{}, fmt.Errorf("Error decoding weather: %s", err)
	}

	// Return the first weather data
	if len(weatherResponse.Properties.Timeseries) == 0 {
		return Data{}, fmt.Errorf("No weather data available")
	}

	return weatherResponse.Properties.Timeseries[0].Data, nil
}
