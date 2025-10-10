package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	getApi "go.mod/internal/api"
	"go.mod/internal/config"
)

var ( 
	cfg config.HTTPServer
)

func SaveDataRedis(handler *getApi.WeatherResp) (*getApi.WeatherResp, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "1234",
		DB: 0,
		MaxRetries: 5,
		ReadTimeout: cfg.TimeOut,
		WriteTimeout: cfg.TimeOut,
		DialTimeout: cfg.IdleTimeout,
	})

	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting on redis: %w", err)
	}

	// Запись данных 
	// Set(контекст, ключ, значение, время в бд)
	if err := rdb.Set(ctx, handler.Location.Name, handler.Current.TempC, time.Hour).Err(); err != nil {
		fmt.Printf("failed to set data, error: %s", err.Error())
	}

	// Получение данных
	val, err := rdb.Get(context.Background(), handler.Location.Name).Result()
	if err == redis.Nil {
		fmt.Println("value not found")
	} else if err != nil {
		fmt.Println("failed to get value, err:", err)
	}

	fVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Printf("failed parse float: %s", err.Error())
	}

	handler.Current.TempC = fVal

	return handler, nil
}