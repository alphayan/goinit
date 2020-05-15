package temp

const REDIS = `package main

import (
	"time"

	redis "github.com/go-redis/redis/v7"
)

var connectionRedis *redis.Client

func initRedis() {
	for i := 0; i < 5; i++ {
		conn := redis.NewClient(&redis.Options{
			Addr:     conf.RedisHost + ":" + conf.RedisPort,
			Password: conf.RedisPassword, // no password set
			DB:       conf.RedisDB,       // use default DB
		})
		if _, err := conn.Ping().Result(); err != nil {
			logger.Error().Msgf("Init  Redis error: %s", err.Error())
			continue
		}
		connectionRedis = conn
		return
	}
	logger.Fatal().Msg("Init  Redis 5 times error,exit")

}
func connectToRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost + ":" + conf.RedisPort,
		Password: conf.RedisPassword, // no password set
		DB:       conf.RedisDB,       // use default DB
	})
}

func Redis() *redis.Client {
	if connectionRedis == nil {
		connectionRedis = connectToRedis()
	}
	_, connected := connectionRedis.Ping().Result()
	i := 0
	for connected != nil {
		if i > 4 {
			logger.Fatal().Msg("Connection to Redis 5 times error,exit")
		}
		i++
		logger.Error().Msg(connected.Error())
		logger.Info().Msg("Connection to Redis was lost. Waiting for 3s...")
		connectionRedis.Close()
		time.Sleep(3 * time.Second)
		logger.Info().Msg("Reconnecting...")
		connectionRedis = connectToRedis()
		_, connected = connectionRedis.Ping().Result()
	}
	return connectionRedis
}
`
