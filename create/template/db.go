package temp

// DB_MYSQL ...
const DB_MYSQL = `package main

import (
	"time"
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

var connectionDB *sql.DB 
func initDB() {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8&timeout=5s"
	for i := 0; i < 5; i++ {
		conn, err := sql.Open("mysql", path)
		if err != nil {
			logger.Error().Msgf("Init Mysql error: %s\n", err.Error())
			continue
		}
		connectionDB = conn
		return
	}
	logger.Fatal().Msgf("Init Mysql 5 times error,exit")
}
func connectToDB() *sql.DB {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8&timeout=5s"
	conn, err := sql.Open("mysql", path)
	if err != nil {
		return nil
	}
	return conn
}

func Database() *sql.DB {
	if connectionDB == nil { 
		connectionDB = connectToDB()
	}
	connected := connectionDB.Ping()
	i := 0
	for connected != nil {
		if i > 4 {
			logger.Fatal().Msgf("Connected Mysql 5 times error,exit")
		}
		i++
		logger.Error().Msg(connected.Error())
		logger.Info().Msg("Connection to Mysql was lost. Waiting for 3s...")
		connectionDB.Close()
		time.Sleep(3 * time.Second)
		logger.Info().Msg("Reconnecting...")
		connectionDB = connectToDB()
		connected = connectionDB.Ping()
	}
	return connectionDB
}
`

// DB_POSTGRESQL ...
const DB_POSTGRESQL = ``
