package temp

const GORM = `package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"time"
)

var connectionDB *gorm.DB // this is "internal", as in: we should NOT use this directly

func connectToDB() *gorm.DB {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8"
	conn, err := gorm.Open("mysql", path)
	if err != nil {
		return &gorm.DB{}
	}

	return conn
}

func Database() *gorm.DB {
	if connectionDB == nil { // during startup - if it does not exist, create it
		connectionDB = connectToDB()
	}
	connected := connectionDB.DB().Ping()
	for connected != nil { // reconnect if we lost connection
		logger.Error().Msg(connected.Error())
		logger.Info().Msg("Connection to Mysql was lost. Waiting for 5s...")
		connectionDB.Close()
		time.Sleep(5 * time.Second)
		logger.Info().Msg("Reconnecting...")
		connectionDB = connectToDB()
		connected = connectionDB.DB().Ping()
	}
	connectionDB.LogMode(true)
	return connectionDB
}
`
