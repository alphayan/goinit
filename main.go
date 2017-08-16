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
	fla := flag.NewFlagSet("goinit", 1)
	fla.Parse([]string{"c", "rm"})
	args := os.Args
	c := flag.String("c", "", "gocreat a new project with the specified name")
	rm := flag.String("rm", "", "remove a project with the specified name")
	flag.Parse()
	if *c != "" {
		fmt.Println(gocreat(*c))
		goinitfmt(*c)
		os.Exit(0)
	}
	if *rm != "" {
		goremove(*rm)
		os.Exit(0)
	}
	if len(args) < 2 {
		fmt.Println("use goinit -h for help")
		return
	}
}
