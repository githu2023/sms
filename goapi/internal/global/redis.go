package global

import (
	"context"
	"fmt"
	"sms-platform/goapi/internal/config"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

// InitRedis initializes the Redis connection.
func InitRedis(cfg config.RedisConfig) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试连接
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	return nil
}

// GetRedis returns the Redis client instance.
// 任何地方都可以直接调用 global.GetRedis() 获取Redis客户端
func GetRedis() *redis.Client {
	return redisClient
}

// CloseRedis closes the Redis connection.
func CloseRedis() error {
	if redisClient != nil {
		return redisClient.Close()
	}
	return nil
}

