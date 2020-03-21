package temp

const XORM_MYSQL = `package main

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"

	"time"
)

var connectionDB *xorm.Engine // this is "internal", as in: we should NOT use this directly

func connectToDB() *xorm.Engine {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8"
	conn, err := xorm.NewEngine("mysql", path)
	if err != nil {
		return &xorm.Engine{}
	}
	return conn
}

func Database() *xorm.Engine {
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
	connectionDB.ShowSQL(true)
	connectionDB.ShowExecTime(true)
	return connectionDB
}
`
const XORM_POSTGRESQL = ``
