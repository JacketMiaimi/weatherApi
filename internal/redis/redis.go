package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	getApi "go.mod/internal/api"
	"go.mod/internal/config"
)

var ( 
	cfg config.HTTPServer
)
type WeatherMarshal getApi.WeatherResp

func (u WeatherMarshal) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func SaveDataRedis() (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Username: "user",
		Password: "1234",
		DB: 0,
		ReadTimeout: cfg.TimeOut,
		WriteTimeout: cfg.TimeOut,
		DialTimeout: cfg.IdleTimeout,
	})

	ctx := context.Background()

	_, err := db.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting on redis: %w", err)
	}

	// Запись данных 
	// Set(контекст, ключ, значение, время в бд)
	if err := db.Set(ctx, "key", "value",0).Err(); err != nil {
		fmt.Printf("failed to set data, error: %s", err.Error())
	}

	if err = db.Set(ctx, "Dexter", "Morgan", 30 * time.Second).Err(); err != nil {
		fmt.Printf("failed to set data, error: %s", err.Error())
	}

	// Получение данных

	val, err := db.Get(context.Background(), "key").Result()
	if err == redis.Nil {
		fmt.Println("value not found")
	} else if err != nil {
		fmt.Println("failed to get value, err:", err)
	}

	val2, err := db.Get(context.Background(), "Dexter").Result()
	if err == redis.Nil {
		fmt.Println("value not found")
	} else if err != nil {
		fmt.Println("failed to get value, err:", err)
	}

	fmt.Printf("value: %v\n", val)
	fmt.Printf("value: %v\n", val2)

	return db, nil

}