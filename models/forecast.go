package models

type Forecast struct {
	Type       string             `json:"type"`
	Geometry   forecastGeometry   `json:"geometry"`
	Properties forecastProperties `json:"properties"`
}

type forecastGeometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float32 `json:"coordinates"`
}

type forecastProperties struct {
	Updated           string            `json:"updated"`
	Units             string            `json:"units"`
	ForecastGenerator string            `json:"forecastGenerator"`
	GeneratedAt       string            `json:"generatedAt"`
	UpdateTime        string            `json:"updateTime"`
	ValidTimes        string            `json:"validTimes"`
	Elevation         forecastElevation `json:"elevation"`
	Periods           []forecastPeriod  `json:"periods"`
}

type forecastElevation struct {
	UnitCode string  `json:"unitCode"`
	Value    float32 `json:"value"`
}

type forecastPeriod struct {
	Number           int    `json:"number"`
	Name             string `json:"name"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	IsDaytime        bool   `json:"isDaytime"`
	Temperature      int    `json:"temperature"`
	TemperatureUnit  string `json:"temperatureUnit"`
	TemperatureTrend string `json:"temperatureTrend"`
	WindSpeed        string `json:"windSpeed"`
	ShortForecast    string `json:"shortForecast"`
	DetailedForecast string `json:"detailedForecast"`
}
