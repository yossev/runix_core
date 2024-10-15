package api

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RateLimiter struct {
	client *redis.Client
	limit  int
	window time.Duration
}

func NewRateLimiter(redisAddr string, limit int, window time.Duration) *RateLimiter {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &RateLimiter{
		client: rdb,
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiter) Allow(key string) bool {
	count, err := rl.client.Incr(ctx, key).Result() // INCR the Key entry

	if err != nil {
		return false
	}

	// Set the expiration window if the Rl is new
	if count == 1 {
		rl.client.Expire(ctx, key, rl.window)
	}
	return count <= int64(rl.limit)
}
