package temp

const ECHO = `package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	echo "github.com/labstack/echo/v4"
)

func initRouter() {
	e := echo.New()
	e.HideBanner=true
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
    go func() {
		if err := e.Start(conf.Port); err != nil {
			logger.Error().Msg("start server error:" + err.Error())
			logger.Fatal().Msg("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal().Msg(err.Error())
	}
}
`
