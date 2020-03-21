package cmd

import (
	"fmt"
	"os"

	"github.com/alphayan/goinit/create"

	"github.com/spf13/cobra"
)

var (
	dir, frame, orm, db string
	module              bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goinit",
	Short: "a tool to create a web application quickly with mysql,orm,redis,config,router",
	Run: func(cmd *cobra.Command, args []string) {
		if err := create.Create(dir, frame, orm, db, module); err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("create", dir, "success")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&dir, "create", "c",
		"default", "creat a new project with the specified name;default is named default")

	rootCmd.Flags().StringVarP(&db, "database", "d", "mysql", `select an database from:
		mysql:
				(MySQL is the world's most popular open source database. Whether you are a fast growing web property, technology ISV 
				or large enterprise, MySQL can cost-effectively help you deliver high performance, scalable database applications.)
		postgresql:
				(The World's Most Advanced Open Source Relational Database)
		mongodb:
				(MongoDB is a general purpose, document-based, distributed database built for modern application developers 
				and for the cloud era. No database makes you more productive.)
		`)

	rootCmd.Flags().StringVarP(&orm, "orm", "o", "default", `select an orm from:
		xorm:
			(Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle http://xorm.io)
		gorm:
			(The fantastic ORM library for Golang, aims to be developer friendly http://gorm.io)
		default is use no orm`)

	rootCmd.Flags().StringVarP(&frame, "frame", `f`, "default", `select a web frame from:
		gin:
			(Gin is a HTTP web framework written in Go (Golang) https://gin-gonic.github.io/gin/)
		echo:
			(High performance, minimalist Go web framework https://echo.labstack.com)
		iris:
			(The fastest backend community-driven web framework on (THIS) Earth. Can your favourite web framework do that? 
	    	👉 http://bit.ly/iriscandothat1 orm even http://bit.ly/iriscandothat2 https://iris-go.com)
		go-json-rest:
			(A quick and easy way to setup a RESTful JSON API https://ant0ine.github.io/go-json-rest/)
			default is use net/http`)

	rootCmd.Flags().BoolVarP(&module, "module", "m",
		false, "creat a new project use go.mod")
}
