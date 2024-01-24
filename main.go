package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Weather struct {
	Now  int64 `json="now"`
	Info struct {
		Lat int    `json="lat"`
		Lon int    `json="lon"`
		Url string `json="url"`
	}
	Fact struct {
		Temp       int `json="temp"`
		Feels_like int `json="feels_like"`
	}
	Forecast struct {
		Date string `json="date"`
	}
}

func main() {
	var api_string string
	flag.StringVar(&api_string, "api", "", "enter yours api key")
	flag.Parse()

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://api.weather.yandex.ru/v2/informers?lat=55.4424&lon=37.3636&land=ru_RU", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-Yandex-API-Key", api_string)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var w Weather
	json.Unmarshal(body, &w)
	formattedPrint(&w)
}

func formattedPrint(w *Weather) {
	t, temp, feels_like, url := w.Now, w.Fact.Temp, w.Fact.Feels_like, w.Info.Url
	weather_time := time.Unix(t, 0)
	fmt.Printf("Время: %s\n", weather_time.UTC())
	fmt.Printf("Погода: %d, ощущается как  %d\n", temp, feels_like)
	fmt.Printf("Больше информации: %s\n", url)
}
