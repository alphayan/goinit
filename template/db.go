package temp

const DB_MYSQL = `package main

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"time"
)

var connectionDB *sql.DB // this is "internal", as in: we should NOT use this directly

func connectToDB() *sql.DB {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8"
	conn, err := sql.Open("mysql", path)
	if err != nil {
		return &sql.DB{}
	}
	return conn
}

func Database() *sql.DB {
	if connectionDB == nil { // during startup - if it does not exist, create it
		connectionDB = connectToDB()
	}
	connected := connectionDB.Ping()
	for connected != nil { // reconnect if we lost connection
		logger.Error().Msg(connected.Error())
		logger.Info().Msg("Connection to Mysql was lost. Waiting for 5s...")
		connectionDB.Close()
		time.Sleep(5 * time.Second)
		logger.Info().Msg("Reconnecting...")
		connectionDB = connectToDB()
		connected = connectionDB.Ping()
	}
	return connectionDB
}
`
const DB_POSTGRESQL = `package main

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"time"
)

var connectionDB *sql.DB // this is "internal", as in: we should NOT use this directly

func connectToDB() *sql.DB {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8"
	conn, err := sql.Open("mysql", path)
	if err != nil {
		return &sql.DB{}
	}
	return conn
}

func Database() *sql.DB {
	if connectionDB == nil { // during startup - if it does not exist, create it
		connectionDB = connectToDB()
	}
	connected := connectionDB.Ping()
	for connected != nil { // reconnect if we lost connection
		logger.Error().Msg(connected.Error())
		logger.Info().Msg("Connection to Mysql was lost. Waiting for 5s...")
		connectionDB.Close()
		time.Sleep(5 * time.Second)
		logger.Info().Msg("Reconnecting...")
		connectionDB = connectToDB()
		connected = connectionDB.Ping()
	}
	return connectionDB
}
`
