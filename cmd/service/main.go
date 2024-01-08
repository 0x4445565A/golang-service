package main

import (
	"fmt"
	"os"

	"github.com/0x4445565a/golang-service/internal/service"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	redisClient := redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_URL"),
		Password: GetEnv("REDIS_PASSWORD"),
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Error("Error Connecting: ", err)
		log.Error("Continuing without cache...")
	}

	log.Info("Connecting to Postgres...")
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("POSTGRES_HOST"),
		GetEnv("POSTGRES_PORT"),
		GetEnv("POSTGRES_USER"),
		GetEnv("POSTGRES_PASSWORD"),
		GetEnv("POSTGRES_DB"))
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	service := service.Service{
		Redis:   redisClient,
		ReadDB:  db,
		WriteDB: db,
	}

	e := service.Init()
	e.Use(middleware.Logger())
	e.Start(":8080")
}
