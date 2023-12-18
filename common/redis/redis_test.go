package redis

import (
	"context"
	"fmt"
	"testing"

	redisLib "github.com/go-redis/redis/v8"
)

func TestNewInstance(t *testing.T) {
	redisUrl := "redis://localhost:6379/0"
	redisClient, err := GetRedisInstance(redisUrl)
	if err != nil {
		t.Error(err)
		return
	}

	rdb := redisClient.Redis

	val, err := rdb.Get(context.Background(), "key").Result()
	if err != nil {
		if err == redisLib.Nil {
			return
		}
		t.Error(err)
		return
	}
	fmt.Println("key", val)
}
