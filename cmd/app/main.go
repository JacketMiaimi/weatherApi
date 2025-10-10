package main

import (
	"fmt"
	rds "go.mod/internal/redis"
	"log"
	getApi "go.mod/internal/api"
)

func main() {
	city := "Moscow"

	data, err := getApi.WeatherHandler(city)
	if err != nil {
		log.Fatal("error getting weather:", err)
	}

	if err := rds.InitRedis(); err != nil {
		log.Fatal("error to connected redis")
	}

	// data,err = rds.SaveKey(data)
	// if err != nil {
	// 	fmt.Printf("error save in redis: %v", err)
	// }

	// if err := rds.DeleteKey(data); err != nil {
	// 	fmt.Printf("error delete in redis: %v\n", err)
	// }

	res, err := rds.GetKey(data)
	if err != nil {
		fmt.Printf("error get in redis: %v\n", err)
	} 

	fmt.Printf("%s: %v\n",city,res)


}
