package getApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	api = "23a923037adc48fca74203916252709"
	baseURL = "https://api.weatherapi.com/v1"
)

type WeatherResp struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}


func WeatherHandler(city string) (*WeatherResp, error) {
	op := "internal/api/getApi.go"
	
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s",baseURL, api, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%scouldn`t get data from api: %w",op, err)
	} 
	defer resp.Body.Close()

	var data WeatherResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("%scouldn`t decode body: %w", op, err)
	}

	return &data, nil
}