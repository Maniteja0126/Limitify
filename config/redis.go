package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis(){
	RedisClient = redis.NewClient(&redis.Options{
		Addr : os.Getenv("REDIS_URL"),
		Password:  "",
		DB:      0,
	})

	_ , err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis " , err)
		os.Exit(1)
	}

	fmt.Println("Connected to Redis")
}
