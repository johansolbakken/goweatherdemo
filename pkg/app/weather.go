package app

import (
	"fmt"
	"johansolbakken.no/weatherdemo/pkg/image"
	"net/http"

	"johansolbakken.no/weatherdemo/pkg/city"
	"johansolbakken.no/weatherdemo/pkg/weather"
)

func (server *Server) getWeather(w http.ResponseWriter, r *http.Request) {
	cityName := r.URL.Query().Get("city")
	if cityName == "" {
		http.Error(w, "City is required", http.StatusBadRequest)
		return
	}

	cityCoordinates, err := city.GetCityCoordinates(cityName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	weatherData, err := weather.GetWeatherFromCoords(cityCoordinates.Lat, cityCoordinates.Lon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	imageSrc, err := image.GetFirstImage(weatherData.Next6Hours.Summary.SymbolCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tempratureColor := "text-red-500"
	if weatherData.Instant.Details.AirTemperature < 0 {
		tempratureColor = "text-blue-500"
	}

	html := fmt.Sprintf(` <h2 class="text-xl font-bold text-gray-800 mb-2">%s</h2>
	<img src="%s" class="w-full">
    <p class="text-gray-700">Temperature: <span class="font-semibold %s">%.1f°C</span></p>
    <p class="text-gray-700">Wind speed: <span class="font-semibold">%.1f m/s</span></p>`,
		cityCoordinates.DisplayName, imageSrc, tempratureColor, weatherData.Instant.Details.AirTemperature, weatherData.Instant.Details.WindSpeed)
	_, err = fmt.Fprintf(w, html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
