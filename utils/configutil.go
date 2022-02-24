package utils

import (
	"encoding/json"
	"go-api-practice/models"
	"io/ioutil"
	"log"
	"os"
)

func LoadConfig() ([]models.Location, error) {
	jsonFile, err := os.Open("configs/places.json")
	if err != nil {
		log.Printf("unable to open config file due to error: %s", err)
		return nil, err
	}
	defer jsonFile.Close()

	data, readError := ioutil.ReadAll(jsonFile)
	if readError != nil {
		log.Printf("unable to read config file due to error: %s", readError)
		return nil, readError
	}

	locations := make([]models.Location, 0)
	jsonError := json.Unmarshal(data, &locations)
	if jsonError != nil {
		log.Printf("unable to unmarshal config into slice due to error: %s", jsonError)
		return nil, jsonError
	}

	return locations, nil

}
