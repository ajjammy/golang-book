package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Weather struct {
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
		TempMin  int     `json:"temp_min"`
		TempMax  int     `json:"temp_max"`
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

func WeatherHandle(w http.ResponseWriter, r *http.Request) {
	cities := []string{"hobart", "newyork", "kupang", "naibori", "bangkok"}
	vars := mux.Vars(r)
	city := vars["city"]
	weather := new(Weather)

	if vars["city"] == "all" {
		for _, c := range cities {
			PrintWeather(w, weather, c)
		}
	} else {
		PrintWeather(w, weather, city)
	}
}

func PrintWeather(w http.ResponseWriter, weather *Weather, city string) {
	res, _ := http.Get("http://localhost:8882/api/v1/weather/" + city)

	json.NewDecoder(res.Body).Decode(weather)

	fmt.Fprintf(w, "%s\n", weather.Name)
	fmt.Fprintf(w, "%.fc %s\n\n", weather.Main.Temp, weather.Weather[0].Description)
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", WeatherHandle).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":3000", NewRouter())
}
