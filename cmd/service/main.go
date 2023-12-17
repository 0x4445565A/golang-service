package main

import (
	"os"

	"github.com/0x4445565a/golang-service/pkg/api"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func GetEnv(key string) string {
	log.Debugf("Loading %s env key", key)
	s := os.Getenv(key)
	if s == "" {
		log.Warnf("Missing %s env key", key)
	}

	return s
}

func main() {
	log.Info("Connecting to Redis...")
	service := api.Service{
		Redis: redis.NewClient(&redis.Options{
			Addr:     GetEnv("REDIS_URL"),
			Password: GetEnv("REDIS_PASSWORD"),
		}),
	}

	_, err := service.Redis.Ping().Result()
	if err != nil {
		log.Error("Error Connecting: ", err)
		log.Error("Continuing without cache...")
	}

	e := service.Init()
	e.Use(middleware.Logger())
	e.Start(":8080")
}
