package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var connectionRedis *redis.Client

func initRedis() {
	for i := 0; i < 5; i++ {
		conn := redis.NewClient(&redis.Options{
			Addr:     conf.RedisHost + ":" + conf.RedisPort,
			Password: conf.RedisPassword, // no password set
			DB:       conf.RedisDB,       // use default DB
		})
		if _, err := conn.Ping(context.Background()).Result(); err != nil {
			logger.Error().Msgf("Init  Redis error: %s", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		connectionRedis = conn
		return
	}
	logger.Fatal().Msg("Init  Redis 5 times error,exit")
}
