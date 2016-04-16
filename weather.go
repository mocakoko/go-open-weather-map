package openweather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Weather ...
type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

// Main ...
type Main struct {
	Temp     float32 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float32 `json:"temp_min"`
	TempMax  float32 `json:"temp_max"`
}

// CurrentWeather ...
type CurrentWeather struct {
	Name    string    `json:"name"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
}

const (
	weatherAPI = "http://api.openweathermap.org/data/2.5/weather"
)

// ReqestCurrentWeather ...
func ReqestCurrentWeather(apiID string, lat float64, lon float64) *CurrentWeather {

	values := url.Values{}
	values.Add("appid", apiID)
	values.Add("lat", strconv.FormatFloat(lat, 'g', 10, 64))
	values.Add("lon", strconv.FormatFloat(lon, 'g', 10, 64))
	values.Add("units", "metric")

	req, _ := http.NewRequest("GET", weatherAPI, nil)
	req.URL.RawQuery = values.Encode()

	log.Print(req)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	if body != nil {
		w := new(CurrentWeather)
		err = json.Unmarshal(body, w)
		if err != nil {
			return nil
		}

		return w

	}

	return nil

}
