package utility

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"time"

	"github.com/go-redis/redis/v8"
)

const CacheDuration = 24 * time.Hour // 24 hours

type ICache interface {
	ConnectRedis() (client *redis.Client)
	SetCache(ctx context.Context, key string, value interface{}) error
	GetCache(ctx context.Context, key string) ([]byte, error)
	GetOrSetCache(ctx context.Context, key string, value interface{}) ([]byte, error)
}

func extractHostPortFromRedisURL(redisURL string) (string, error) {
	parsedURL, err := url.Parse(redisURL)
	if err != nil {
		return "", err
	}
	return parsedURL.Host, nil
}

type RedisClient struct {
	host   string
	db     int
	exp    time.Duration
	client *redis.Client
}

func NewRedisClient(host string, db int, exp time.Duration) *RedisClient {
	return &RedisClient{
		host: host,
		db:   db,
		exp:  exp,
	}
}

func (r *RedisClient) ConnectRedis() *redis.Client {
	opt, _ := redis.ParseURL("rediss://default:a3677c1a7b84402eb34efd55ad3cf059@golden-colt-33790.upstash.io:33790")
	//client := redis.NewClient(opt)
	newClient := redis.NewClient(opt)

	pong, err := newClient.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Error connecting to redis")
	}
	log.Println("                                                        ")
	log.Println("=====  =====  =====  =====  =====  =====  =====  =======")
	log.Println("=========== Connecting To Redis on port 6379 ===========")
	log.Printf("============== Ready For Cache: %s ===================", pong)
	log.Println("=====  =====  =====  =====  =====  =====  =====  =======")
	log.Println("                                                        ")
	r.client = newClient
	return newClient
}

func (r *RedisClient) SetCache(ctx context.Context, key string, value interface{}) error {
	log.Println("Setting cache for key", key)
	if r.client == nil {
		return errors.New("redis client is nil")
	}
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	log.Println("Setting cache for key", key)
	//log.Println("Setting cache for value", jsonValue)
	return r.client.Set(ctx, key, jsonValue, r.exp).Err()
}

func (r *RedisClient) GetCache(ctx context.Context, key string) ([]byte, error) {
	log.Println("Getting cache for key", key)
	if r.client == nil {
		return nil, errors.New("redis client is nil")
	}
	val, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (r *RedisClient) GetOrSetCache(ctx context.Context, key string, value interface{}) ([]byte, error) {
	if r.client == nil {
		return nil, errors.New("redis client is nil")
	}
	val, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil || val == nil {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		err = r.client.Set(ctx, key, jsonValue, r.exp).Err()
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return val, nil
}