package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rdb = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "QtyMqYWUwVeTysnCnnoneXMHRIgBoci/uA==",
		DB:           0,
		MaxRetries:   5,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		DialTimeout:  3 * time.Second,
	})
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

func Save(key string, val float64) error {
	// Запись данных
	// Set(контекст, ключ, значение, время в бд)
	if err := rdb.Set(ctx, key, val, time.Hour).Err(); err != nil {
		slog.Error("failed to set data","error:", err.Error())
	}

	return nil
}

func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("error value not found: %w", err)
	} else if err != nil {
		return "", fmt.Errorf("error get data in redis: %w", err)
	}



	return val, nil
}

func Delete(key string) error {
	if err := rdb.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("error delete in redis: %v", err)
	}
	return nil
}
