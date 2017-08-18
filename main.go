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
	args := os.Args                    //获取命令行参数
	fl := flag.NewFlagSet("goinit", 0) //自定义flag
	c := fl.String("c", "", "creat a new project with the specified name")
	rm := fl.String("rm", "", "remove a project with the specified name")
	frame := fl.String("frame", "", `select a web frame from:
	gin
	beego
	echo
	go-json-rest
	`)
	fl.Parse(args[1:]) //需要使用第二个参数之后的切片
	if *c != "" {
		fmt.Println(gocreat(*c, *frame))
		goinitfmt(*c)
		os.Exit(0)
	}
	if *rm != "" {
		fmt.Println(goremove(*rm))
		os.Exit(0)
	}

	if len(args) < 2 {
		fmt.Println("use goinit -h for help")
		return
	}
}
