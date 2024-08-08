# Weather CLI

A simple command-line interface (CLI) for fetching and displaying weather information for a given city.

## Features

- Allows users to enter a city and displays the current weather information
- Displays the current temperature, feels like temperature, minimum and maximum temperatures, and humidity
- Uses weather icons to visually represent the current weather conditions

## Prerequisites

- Go programming language (version 1.16 or later)
- An API key from [OpenWeatherMap](https://openweathermap.org/) to access the weather data

## Installation

1. Clone the repository:
https://github.com/QoSz/cli-weather-info.git

2. Navigate to the project directory:
cd cli-weather-info

3. Open the `weather.go` file and replace `"your_api_key_here"` with your actual OpenWeatherMap API key.

## Usage

1. Run the application:
go run weather.go

2. When prompted, enter the name of the city you want to get the weather information for.

3. The application will display the current weather information for the specified city, including:
- Current temperature
- Feels like temperature
- Minimum and maximum temperatures
- Humidity

Example output:
Enter a city: london
Current weather in London
⛅ 19.3°C
Feels like: 19.6°C
Min/Max: 18.2°C/20.1°C
Humidity: 90%

## License

This project is licensed under the [MIT License](LICENSE).