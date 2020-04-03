package temp

// GORM_MYSQL ...
const GORM_MYSQL = `package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var connectionDB *gorm.DB

func initDB() {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8&timeout=5s"
	for i := 0; i < 5; i++ {
		conn, err := gorm.Open("mysql", path)
		if err != nil {
			logger.Error().Msgf("Init Mysql error: %s\n", err.Error())
			continue
		}
		connectionDB = conn
		connectionDB.LogMode(conf.DBDebug)
		connectionDB.DB().SetMaxIdleConns(conf.DBMaxIdleConns)
		connectionDB.DB().SetMaxOpenConns(conf.DBMaxOpenConns)
		return
	}
	logger.Fatal().Msgf("Init Mysql 5 times error,exist")
}

func connectToDB() *gorm.DB {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8&timeout=5s"
	conn, err := gorm.Open("mysql", path)
	if err != nil {
		logger.Error().Msgf("Connection Mysql error: %s\n", err.Error())
		return nil
	}
	return conn
}

func Database() *gorm.DB {
	if connectionDB == nil {
		connectionDB = connectToDB()
	}
	connected := connectionDB.DB().Ping()
	i := 0
	for connected != nil {
		if i > 4 {
			logger.Fatal().Msgf("Connected Mysql 5 times error,exist")
		}
		i++
		logger.Error().Msg(connected.Error())
		logger.Info().Msg("Connection to Mysql was lost. Waiting for 3s...")
		connectionDB.Close()
		time.Sleep(3 * time.Second)
		logger.Info().Msg("Reconnecting...")
		connectionDB = connectToDB()
		connected = connectionDB.DB().Ping()
	}
	connectionDB.LogMode(conf.DBDebug)
	connectionDB.DB().SetMaxIdleConns(conf.DBMaxIdleConns)
	connectionDB.DB().SetMaxOpenConns(conf.DBMaxOpenConns)
	return connectionDB
}
`

// GORM_POSTGRESQL ...
const GORM_POSTGRESQL = ``
