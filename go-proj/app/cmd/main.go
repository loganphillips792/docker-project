package main

import (
	"context"
	"github.com/loganphillips792/kubernetes-project/components"
	"github.com/loganphillips792/kubernetes-project/config"
	"log/slog"
	"os"

	_ "github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)


func main() {

	cfg, configError := config.Init()

	
	if configError != nil {
		slog.Error("error when reading config file")
	}


	var logger *slog.Logger

	if cfg.AppEnvironment == "development" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}


	e := echo.New()


	e.GET("/hello", func(c echo.Context) error {
		component := components.Hello("John")
		return component.Render(context.Background(), c.Response().Writer) 
	})

	
	logger.Info(
		"incoming request",
		"method", "GET",
		"time_taken_ms", 158,
		"path", "/hello/world?q=search",
		"status", 200,
		"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
	)
	logger.Info("Listening on :3000")
	e.Logger.Fatal(e.Start(":3000"))

}