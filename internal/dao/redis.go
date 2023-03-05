package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	Cache *redis.Client
	Nil   = redis.Nil
)

func InitRedis() (err error) {
	Cache = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis-local.address"),
		Password:     viper.GetString("redis-local.passowrd"),
		DB:           viper.GetInt("redis-local.db"),
		PoolSize:     viper.GetInt("redis-local.pool_size"),
		MinIdleConns: viper.GetInt("redis-local.min_idle_conns"),
	})
	_, err = Cache.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}
