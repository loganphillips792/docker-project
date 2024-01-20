package main

import (
	"context"
	"github.com/loganphillips792/kubernetes-project/components"
	"github.com/loganphillips792/kubernetes-project/config"
	"log/slog"
	"os"
	"fmt"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"database/sql"
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

	db := initializeDatabase()
	defer db.Close()


	e := echo.New()


	e.GET("/hello", func(c echo.Context) error {
		component := components.Hello("John")
		return component.Render(context.Background(), c.Response().Writer) 
	})

	e.GET("/", func(c echo.Context) error {
		component := components.Page(5,5)
		return component.Render(context.Background(), c.Response().Writer) 
	})

	e.POST("/", func(c echo.Context) error  {
		component := components.Page(6,5)
		return component.Render(context.Background(), c.Response().Writer)
	})
	
	logger.Info("Listening on :3000")
	e.Logger.Fatal(e.Start(":3000"))

}

func initializeDatabase() *sqlx.DB {
	slog.Info("Initializing SQL Lite database...")

	file, openFileErr := os.Open("data.db")

	if openFileErr != nil {
		slog.Info(openFileErr.Error())
	}

	if errors.Is(openFileErr, os.ErrNotExist) {
		file, _ = os.Create("data.db")
	}

	file.Close()

	db, err := sql.Open("sqlite3", "data.db")

	if err != nil {
		slog.Error(err.Error())
	}

	sqlxDb := sqlx.NewDb(db, "sqlite3")

	// create tables and seed data
	if errors.Is(openFileErr, os.ErrNotExist) {
		c, err := ioutil.ReadFile("script.sql")

		if err != nil {
			slog.Error(err.Error())
		}

		sql := string(c)

		_, err = db.Exec(sql)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
	return sqlxDb
}