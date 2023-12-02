package redis

import (
	"context"
	"fmt"
	"github.com/bullean-ai/hexa-neural-net/config"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	db  *redis.Client
	Ctx context.Context
	err error
)

// NewRedisClient Returns new redis client
func NewRedisClient(ctx context.Context, cfg *config.Config) (db *redis.Client, err error) {
	println("Driver Redis Initialized")

	redisHost := fmt.Sprintf("%s:%d", cfg.Redis.HOST, cfg.Redis.PORT)

	db = redis.NewClient(&redis.Options{
		Addr:         redisHost,
		Username:     cfg.Redis.USER,
		Password:     cfg.Redis.PASS,
		MinIdleConns: cfg.Redis.MIN_IDLE_CONN,
		PoolSize:     cfg.Redis.POOL_SIZE,
		PoolTimeout:  time.Second * time.Duration(cfg.Redis.POOL_TIMEOUT),
		DB:           cfg.Redis.DEFAULT_DB,
	})

	err = db.Ping(ctx).Err()

	return
}

/*
// NewRedisClient Returns new redis client
func NewRedisClient(cfg *config.Config) *redis.Client {
	println("Driver Redis Initialized")

	redisHost := fmt.Sprintf("%s:%d", cfg.Redis.HOST, cfg.Redis.PORT)

	db = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: cfg.Redis.PASS,
	})

	return db
}
*/
