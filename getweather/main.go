package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const APIKEY string = "buraya kendi api keyinizi yazın"

func main() {

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=izmir&units=metric&appid=" + APIKEY)

	if err != nil {
		log.Fatal(err)
	}

	weather_web_result, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var weatherResult WeatherResult
	err = json.Unmarshal(weather_web_result, &weatherResult)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(weatherResult.Name, weatherResult.Main.Temp, "derece")

}

type WeatherResult struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
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
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}
