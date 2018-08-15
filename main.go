package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

// GOPATHSRC gopath/src路径
var GOPATHSRC string

func init() {
	GOPATHSRC = path.Join(os.Getenv("GOPATH"), "/src/")
}

func main() {
	//获取命令行参数
	args := os.Args
	//自定义flag
	f := flag.NewFlagSet("goinit", 0)
	c := f.String("c", "default", `creat a new project with the specified name
	default is named default`)
	db := f.String("db", "", `select an orm from:
	xorm:
		(Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle http://xorm.io)
	gorm:
		(The fantastic ORM library for Golang, aims to be developer friendly http://gorm.io)
	default is no orm`)
	frame := f.String("frame", "", `select a web frame from:
	gin:
		(Gin is a HTTP web framework written in Go (Golang) https://gin-gonic.github.io/gin/)
	echo:
		(High performance, minimalist Go web framework https://echo.labstack.com)
	iris:
		(The fastest backend community-driven web framework on (THIS) Earth. Can your favourite web framework do that? 
	    👉 http://bit.ly/iriscandothat1 or even http://bit.ly/iriscandothat2 https://iris-go.com)
	go-json-rest:
		(A quick and easy way to setup a RESTful JSON API https://ant0ine.github.io/go-json-rest/)
	default is use net/http`)
	//需要使用第二个参数之后的切片
	f.Parse(args[1:])
	fmt.Println(gocreate(*c, *frame, *db))
}
