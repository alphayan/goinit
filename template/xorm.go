package temp

const XORM_MYSQL = `package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var connectionDB *xorm.Engine

func initDB() {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8&timeout=5s"
	for i := 0; i < 5; i++ {
		conn, err := xorm.NewEngine("mysql", path)
		if err != nil {
			logger.Error().Msgf("Init Mysql error: %s\n", err.Error())
			continue
		}
		connectionDB = conn
		return
	}
	logger.Fatal().Msgf("Init Mysql 5 times error,exist")
}
func connectToDB() *xorm.Engine {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8&timeout=5s"
	conn, err := xorm.NewEngine("mysql", path)
	if err != nil {
		logger.Error().Msgf("Connection to Mysql error: %s\n", err.Error())
		return nil
	}
	return conn
}

func Database() *xorm.Engine {
	if connectionDB == nil { 
		connectionDB = connectToDB()
	}
	connected := connectionDB.Ping()
	i := 0
	for connected != nil { 
		if i > 4 {
			logger.Fatal().Msgf("Connection to Mysql 5 times error,exist")
		}
		logger.Error().Msg(connected.Error())
		logger.Info().Msg("Connection to Mysql was lost. Waiting for 3s...")
		connectionDB.Close()
		time.Sleep(3 * time.Second)
		logger.Info().Msg("Reconnecting...")
		connectionDB = connectToDB()
		connected = connectionDB.Ping()
	}
	connectionDB.ShowSQL(true)
	return connectionDB
}
`
const XORM_POSTGRESQL = ``
