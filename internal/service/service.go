package service

import (
	
	"log/slog"
	"strconv"

	getApi "go.mod/internal/api"
	"go.mod/internal/redis"
)

func GetWeather(city string) (float64, error) {
	res, err := redis.Get(city)
	if err != nil{
		slog.Warn("city in redis not found","err", err)
	}

	if res == "" {
		weatherApi,err := getApi.WeatherHandler(city)
		if err != nil {
			slog.Error("failed to get city in api", "err", err)
		}
		
		temp := weatherApi.Current.TempC
		slog.Info("The value was obtained from API")
		
		if err = redis.Save(city,temp); err != nil {
			slog.Error("failed to save in redis", "err", err)
		}
		slog.Info("Value was save to redis")

		return temp, nil
	}

	temp,err := strconv.ParseFloat(res, 64)

	return temp, nil
}