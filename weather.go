package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp      float64 `json:"temp"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Cod int `json:"cod"`
}

func getWeather(city string) (*WeatherData, error) {
	apiKey := "your_api_key_here"
	baseURL := "http://api.openweathermap.org/data/2.5/weather"
	params := map[string]string{
		"q":     city,
		"appid": apiKey,
		"units": "metric",
	}

	resp, err := http.Get(baseURL + "?" + encodeParams(params))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	// Check if the API request was successful
	if weatherData.Cod != 200 {
		return nil, fmt.Errorf("error fetching weather data (code %d)", weatherData.Cod)
	}

	if len(weatherData.Weather) == 0 {
		return nil, fmt.Errorf("no weather data available for %s", weatherData.Name)
	}

	return &weatherData, nil
}

func encodeParams(params map[string]string) string {
	paramStr := ""
	for key, value := range params {
		paramStr += key + "=" + value + "&"
	}
	return paramStr[:len(paramStr)-1]
}

func getWeatherIcon(weatherDescription string) string {
	switch weatherDescription {
	case "clear sky":
		return "☀️"
	case "few clouds", "scattered clouds", "broken clouds":
		return "⛅"
	case "overcast clouds":
		return "☁️"
	case "light rain", "moderate rain", "heavy intensity rain", "very heavy rain", "extreme rain":
		return "🌧️"
	case "light snow", "snow":
		return "❄️"
	case "thunderstorm":
		return "⚡"
	case "mist", "haze", "fog":
		return "🌫️"
	default:
		return "❓"
	}
}
func main() {
	var city string
	fmt.Print("Enter a city: ")
	fmt.Scanln(&city)

	weatherData, err := getWeather(city)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\n  Current weather in %s\n", weatherData.Name)
	fmt.Printf("  %s %.1f°C\n", getWeatherIcon(weatherData.Weather[0].Description), weatherData.Main.Temp)
	fmt.Printf("  Feels like: %.1f°C\n", weatherData.Main.FeelsLike)
	fmt.Printf("  Min/Max: %.1f°C/%.1f°C\n", weatherData.Main.TempMin, weatherData.Main.TempMax)
	fmt.Printf("  Humidity: %d%%\n\n", weatherData.Main.Humidity)
}
