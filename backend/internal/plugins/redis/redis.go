package redisPlugin

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	rdb     *redis.Client
	options *Options
}

type Options struct {
	Timeout  time.Duration
	Host     string
	Port     string
	Password string
	Username string
}

func CreateClient(options Options) (*Client, error) {
	var client *Client

	if options.Host == "" {
		return client, errors.New("hasn't redis host")
	}
	if options.Port == "" {
		return client, errors.New("hasn't redis port")
	}

	var rdb *redis.Client
	var err error

	rdb = redis.NewClient(&redis.Options{
		Addr:     options.Host + ":" + options.Port,
		Username: options.Username,
		Password: options.Password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
	defer cancel()

	if _, err = rdb.Ping(ctx).Result(); err != nil {
		return client, err
	}

	client = &Client{
		rdb:     rdb,
		options: &options,
	}

	return client, nil
}

func (client *Client) Get(key string) (string, error) {
	var result string
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), client.options.Timeout)
	defer cancel()

	result, err = client.rdb.Get(ctx, key).Result()

	return result, err

}

func (client *Client) Set(key string, value string, expiration time.Duration) error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), client.options.Timeout)
	defer cancel()

	err = client.rdb.Set(ctx, key, value, expiration).Err()

	return err
}

func (client *Client) Del(keys ...string) error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), client.options.Timeout)
	defer cancel()

	err = client.rdb.Del(ctx, keys...).Err()

	return err

}
