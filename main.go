package main

import (
	"flag"
	"fmt"
	"os"
)

// GOPATHSRC gopath/src路径
var GOPATHSRC string

func init() {
	GOPATHSRC = os.Getenv("GOPATH") + "/src/"
}

func main() {
	//获取命令行参数
	args := os.Args
	//自定义flag
	f := flag.NewFlagSet("goinit", 0)
	c := f.String("c", "", "creat a new project with the specified name")
	frame := f.String("frame", "", `select a web frame from:
	gin (Gin is a HTTP web framework written in Go (Golang) https://gin-gonic.github.io/gin/)
	echo (High performance, minimalist Go web framework https://echo.labstack.com)
	go-json-rest (A quick and easy way to setup a RESTful JSON API https://ant0ine.github.io/go-json-rest/)`)
	rm := f.String("rm", "", "remove a project with the specified name")
	//需要使用第二个参数之后的切片
	f.Parse(args[1:])
	if *c != "" {
		fmt.Println(gocreate(*c, *frame))
		return
	}
	if *rm != "" {
		fmt.Println(goremove(*rm))
		return
	}
	if len(args) < 2 {
		fmt.Println("use 'goinit -h' for help")
		return
	}
}
