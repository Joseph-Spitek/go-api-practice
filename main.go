package main

import (
	"fmt"
	"go-api-practice/models"
	"go-api-practice/utils"
	"log"
	"sync"
)

const (
	baseUrl string = "https://api.weather.gov"
)

func main() {
	locations, configErr := utils.LoadConfig()
	if configErr != nil {
		return
	}

	log.Printf("Found %d configs", len(locations))

	inputs := make(chan models.Forecast)
	outputs := make(chan []models.Forecast)
	var wg sync.WaitGroup

	go handleForecasts(inputs, outputs, &wg)
	defer close(outputs)

	for _, location := range locations {
		wg.Add(1)
		go fetchWeather(location, inputs)
	}

	wg.Wait()
	close(inputs)

	forecasts := <-outputs
	log.Printf("Completed with %d forecasts", len(forecasts))

}

func handleForecasts(input chan models.Forecast, outputs chan []models.Forecast, wg *sync.WaitGroup) {
	var forecasts []models.Forecast

	for forecast := range input {
		forecasts = append(forecasts, forecast)
		wg.Done()
	}

	outputs <- forecasts
}

func fetchWeather(location models.Location, inputs chan models.Forecast) {
	pointsUrl := fmt.Sprintf("%s/points/%.4f,%.4f", baseUrl, location.Latitude, location.Longitude)
	var point models.Point
	pointErr := utils.PerformGetRequest(pointsUrl, &point)
	if pointErr != nil {
		log.Printf("Error calling points API: %s", pointErr)
		inputs <- models.Forecast{}
		return
	}

	var forecast models.Forecast
	forecastErr := utils.PerformGetRequest(point.Properties.Forecast, &forecast)
	if forecastErr != nil {
		log.Printf("Error calling forecast API: %s", forecastErr)
		inputs <- models.Forecast{}
		return
	}

	inputs <- forecast
}
