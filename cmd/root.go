// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/alphayan/goinit/create"

	"github.com/spf13/cobra"
)

var dir, frame, db string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goinit",
	Short: "A tool to create a web application quickly with mysql,orm,redis,config,router",
	Run: func(cmd *cobra.Command, args []string) {
		if err := create.Create(dir, frame, db); err != nil {
			fmt.Println(err)
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
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&dir, "create", "c",
		"default", "creat a new project with the specified name;default is named default")

	rootCmd.Flags().StringVarP(&db, "orm", "o", "default", `select an orm from:
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
	    👉 http://bit.ly/iriscandothat1 or even http://bit.ly/iriscandothat2 https://iris-go.com)
	go-json-rest:
		(A quick and easy way to setup a RESTful JSON API https://ant0ine.github.io/go-json-rest/)
	default is use net/http`)
}
