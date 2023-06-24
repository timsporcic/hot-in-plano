package hot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const owUrl = "https://api.openweathermap.org/data/2.5/weather?q=plano,tx,usa&units=imperial&APPID=%s"

type OpenWeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type Weather struct {
	Temperature float64
	Humidity    int
	FeelsLike   float64
}

func GetWeather() Weather {

	apiKey, ok := os.LookupEnv("OW_API_KEY")

	if !ok {
		log.Fatal("OW_API_KEY is missing from the environment")
	}

	url := fmt.Sprintf(owUrl, apiKey)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Failed to retrieve OpenWeather data")
	}

	var openWx OpenWeatherData

	if err := json.NewDecoder(resp.Body).Decode(&openWx); err != nil {
		log.Fatal("Failed to deserialize OpenWeather response")
	}

	w := Weather{
		Temperature: openWx.Main.Temp,
		FeelsLike:   openWx.Main.FeelsLike,
		Humidity:    openWx.Main.Humidity,
	}

	return w
}
