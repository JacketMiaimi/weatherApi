package main

import (
	"fmt"

	"go.mod/internal/redis"
)

// "fmt"
// "log"

// getApi "go.mod/internal/api"
// "go.mod/internal/config"

func main() {
	// _ = config.LoadConfig("C:\\IT\\Go\\petProject\\unfinished\\weatherApi\\configs\\local.yaml")
	
	// city, err := getApi.WeatherHandler("London")
	// if err != nil {
	// 	log.Fatal("Error getting weather:", err)
	// }

	// fmt.Printf("city: %s\ntemp: %.1f", city.Location.Name, city.Current.TempC)

	value,err := redis.SaveDataRedis()
	if err != nil {
		fmt.Printf("Error in redis: %v", err)
	}

	fmt.Println(value)
}