package repo

import (
	"context"
	"log"
	"shortLink/internal/config"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
	TTL time.Duration
)

func InitRedis() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Addr,
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.DB,
	})
	TTL = time.Duration(config.C.Redis.TTL) * time.Second
	if TTL <= 0 {
		TTL = 24 * time.Hour
	}

	// ping
	if err := RDB.Ping(context.Background()).Err(); err != nil {
		return err
	}
	log.Println("[redis] connected")
	return nil
}
