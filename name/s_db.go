package main

import (
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
		conn.DB().SetMaxOpenConns(conf.DBMaxOpenConns)
		conn.DB().SetMaxIdleConns(conf.DBMaxIdleConns)
		connectionDB = conn
		return
	}
	logger.Fatal().Msgf("Init Mysql 5 times error,exist")
}
