package utils

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
)

func GetCache(ctx context.Context, redisConn *redis.Client, keyRedis string) (map[string]any, error) {

	var resData map[string]any

	val, err := redisConn.Get(ctx, keyRedis).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == redis.Nil {
		return nil, nil
	}

	err = sonic.Unmarshal([]byte(val), &resData)
	if err != nil {
		return nil, err
	}

	return resData, nil
}
