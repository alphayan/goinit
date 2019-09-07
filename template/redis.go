package temp

const REDIS = `package main

import (
	"time"

	"github.com/go-redis/redis"
)

var connectionRedis *redis.Client

func connectToRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost + ":" + conf.RedisPort,
		Password: conf.RedisPassword, // no password set
		DB:       conf.RedisDB,       // use default DB
	})
}

func Redis() *redis.Client {
	if connectionRedis == nil { // during startup - if it does not exist, create it
		connectionRedis = connectToRedis()
	}
	_, connected := connectionRedis.Ping().Result()
	for connected != nil { // reconnect if we lost connection
		logger.Error().Msg(connected.Error())
		logger.Info().Msg("Connection to Mysql was lost. Waiting for 5s...")
		connectionRedis.Close()
		time.Sleep(5 * time.Second)
		logger.Info().Msg("Reconnecting...")
		connectionRedis = connectToRedis()
		_, connected = connectionRedis.Ping().Result()
	}
	return connectionRedis
}
`
