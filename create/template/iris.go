package temp

const IRIS = `package main

import "github.com/kataras/iris/v12"

func initRouter() {
	app := iris.Default()
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})
	app.Run(iris.Addr(conf.Port))
}`
