package plugins

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Rdb     *redis.Client
	Options *RedisClientOptions
}

type RedisClientOptions struct {
	Timeout time.Duration
}

func CreateRedisClient(options RedisClientOptions) (*RedisClient, error) {
	var client *RedisClient

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		return client, errors.New("hasn't redis host")
	}
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		return client, errors.New("hasn't redis port")
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisUsername := os.Getenv("REDIS_USERNAME")

	var rdb *redis.Client
	var err error

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Username: redisUsername,
		Password: redisPassword,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
	defer cancel()

	if _, err = rdb.Ping(ctx).Result(); err != nil {
		return client, err
	}

	client = &RedisClient{
		Rdb:     rdb,
		Options: &options,
	}

	return client, nil
}

func (client *RedisClient) Get(key string) (string, error) {
	var result string
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), client.Options.Timeout)
	defer cancel()

	result, err = client.Rdb.Get(ctx, key).Result()

	return result, err

}

func (client *RedisClient) Set(key string, value string, expiration time.Duration) error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), client.Options.Timeout)
	defer cancel()

	err = client.Rdb.Set(ctx, key, value, expiration).Err()

	return err
}

func (client *RedisClient) Del(keys ...string) error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), client.Options.Timeout)
	defer cancel()

	err = client.Rdb.Del(ctx, keys...).Err()

	return err

}
