package main

import (
	"fmt"
	"go.mod/internal/redis"
	"log"
	getApi "go.mod/internal/api"
)

func main() {
	// _ = config.LoadConfig("C:\\IT\\Go\\petProject\\unfinished\\weatherApi\\configs\\local.yaml")
	
	City, err := getApi.WeatherHandler("London")
	if err != nil {
		log.Fatal("Error getting weather:", err)
	}

	// fmt.Printf("city: %s\ntemp: %f", city.Location.Name, city.Current.TempC)

	value,err := redis.SaveDataRedis(City)
	if err != nil {
		fmt.Printf("Error in redis: %v", err)
	}

	fmt.Println(
		value.Location.Name,
		value.Current.TempC,
	)
}