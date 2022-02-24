package models

type Point struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Properties pointProperties `json:"properties"`
	Geometry   pointGeometry   `json:"geometry"`
}

type pointProperties struct {
	// Id             string `json:"@id"`
	// Type           string `json:"@type"`
	// Cwa            string `json:"cwa"`
	// ForecastOffice string `json:"forecastOffice"`
	// GridId         int    `json:"gridId"`
	// GridX          string `json:"gridX"`
	// GridY          string `json:"gridY"`
	Forecast string `json:"forecast"`
	// ForecastHourly string `json:"forecastHourly"`
	// ForecastHourly string `json:"forecastGridData"`
}

type pointGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float32 `json:"coordinates"`
}
