package city

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type City struct {
	DisplayName string `json:"display_name"`
	AddressType string `json:"addresstype"`
	LatStr      string `json:"lat"`
	Lat         float64
	LonStr      string `json:"lon"`
	Lon         float64
}

type Cities []City

func GetCityCoordinates(cityName string) (city City, err error) {
	// Properly escape the city parameter to ensure the URL is correctly formatted
	escapedCity := url.QueryEscape(cityName)
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?addressdetails=1&q=%s&format=jsonv2&limit=10", escapedCity)

	// Create a new HTTP request
	req, err := http.Get(url)
	if err != nil {
		return City{}, fmt.Errorf("Error getting city: %s", err)
	}
	defer req.Body.Close()

	// Decode the JSON response into our Cities struct
	var cities Cities
	if err := json.NewDecoder(req.Body).Decode(&cities); err != nil {
		return City{}, fmt.Errorf("Error decoding city: %s", err)
	}

	if len(cities) == 0 {
		return City{}, fmt.Errorf("No city found: %s", err)
	}

	city = cities[0]
	city.Lat, err = strconv.ParseFloat(city.LatStr, 64)
	if err != nil {
		return City{}, fmt.Errorf("Invalid lat: %s", err)
	}

	city.Lon, err = strconv.ParseFloat(city.LonStr, 64)
	if err != nil {
		return City{}, fmt.Errorf("Invalid lon: %s", err)
	}

	return city, nil
}
