package utility

import (
"github.com/redis/go-redis/v9"
"log"
"context"
"time"
)


func CreateRedis() *redis.Client{

 rdb := redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
	Password:"",
	DB: 0,
})
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

// Ping Redis
pong, err := rdb.Ping(ctx).Result()
if err != nil {
	log.Fatalf("Could not connect to Redis: %v", err)
}

log.Printf("Redis connected: %s", pong)
return rdb

}