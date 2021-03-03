package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/alphayan/goinit/create"

	"github.com/spf13/cobra"
)

var (
	dir, frame, orm string
	str             *[]string
	module          bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goinit",
	Short: "a tool to create a web application quickly with mysql,orm,redis,config,router",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.Version() < "go1.16" {
			fmt.Println("error:", "go version must above go1.16")
			return
		}
		if err := create.Create(dir, frame, orm, module, str); err != nil {
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
		"default", "creat a new project with the specified name")

	rootCmd.Flags().StringVarP(&orm, "orm", "o", "gorm", `select an orm from:
		xorm:
			(Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle http://xorm.io)
		gorm:
			(The fantastic ORM library for Golang, aims to be developer friendly http://gorm.io)`)

	rootCmd.Flags().StringVarP(&frame, "frame", `f`, "echo", `select a web frame from:
		gin:
			(Gin is a HTTP web framework written in Go (Golang) https://gin-gonic.github.io/gin/)
		echo:
			(High performance, minimalist Go web framework https://echo.labstack.com)`)

	rootCmd.Flags().BoolVarP(&module, "module", "m",
		true, "creat a new project use go.mod")
	str = rootCmd.Flags().StringSliceP("struct", "s", []string{},
		"creat a struct controller and model")
}
