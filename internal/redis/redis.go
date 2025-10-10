package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	getApi "go.mod/internal/api"
)

var (
	rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "QtyMqYWUwVeTysnCnnoneXMHRIgBoci/uA==",
	DB: 0,
	MaxRetries: 5,
	ReadTimeout: 2 * time.Second,
	WriteTimeout: 2 * time.Second,
	DialTimeout: 3 * time.Second,
});
	ctx = context.Background()
)

func InitRedis() error {
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("error connecting on redis: %w", err)
	}
	fmt.Println("Connected to redis")
	return nil
}

func SaveKey(handler *getApi.WeatherResp) (*getApi.WeatherResp, error) {
	// Запись данных 
	// Set(контекст, ключ, значение, время в бд)
	if err := rdb.Set(ctx, handler.Location.Name, handler.Current.TempC, time.Hour).Err(); err != nil {
		fmt.Printf("failed to set data, error: %s", err.Error())
	}

	return handler, nil
}

func GetKey(handler *getApi.WeatherResp) (string, error) {
	val, err := rdb.Get(ctx, handler.Location.Name).Result()
	if err == redis.Nil {
		return  "", fmt.Errorf("error value not found: %w", err)
	} else if err != nil {
		return "", fmt.Errorf("error get data in redis: %w", err)
	}

	return val, nil
}

func DeleteKey(handler *getApi.WeatherResp) error {
	if err := rdb.Del(ctx, handler.Location.Name).Err(); err != nil {
		return fmt.Errorf("error delete in redis: %v", err)
	}
	
	fmt.Println("successfully removed")

	return nil
}