package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	DB *DatabaseClient
}

func main() {
	config := LoadConfig()

	db, err := NewDatabaseClient(config.TableName)
	if err != nil {
		log.Fatalf("Failed to connect to DynamoDB: %v", err)
	}

	app := &App{DB: db}
	_ = app // TODO: Will be used when implementing API endpoints

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Echo TODO API",
			"status":  "ready",
		})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	e.Logger.Fatal(e.Start(":" + config.Port))
}