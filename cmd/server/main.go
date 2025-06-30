// @title Echo TODO API
// @version 1.0
// @description A simple TODO API built with Echo framework and DynamoDB
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:1323
// @BasePath /
// @schemes http https

package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"echo-todo/internal/config"
	"echo-todo/internal/handlers"
	"echo-todo/internal/repository"
	"echo-todo/internal/services"
	_ "echo-todo/docs"
)

func main() {
	cfg := config.Load()

	todoRepo, err := repository.NewDynamoDBTodoRepository(cfg.TableName)
	if err != nil {
		log.Fatalf("Failed to initialize todo repository: %v", err)
	}

	// Initialize service layer
	todoService := services.NewTodoService(todoRepo)
	
	// Initialize handler layer
	todoHandler := handlers.NewTodoHandler(todoService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Echo TODO API on AWS Lambda with LWA",
			"status":  "ready",
		})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// API routes
	api := e.Group("/api/v1")
	
	// TODO routes
	todos := api.Group("/todos")
	todos.POST("", todoHandler.CreateTodo)
	todos.GET("", todoHandler.GetAllTodos)
	todos.GET("/:id", todoHandler.GetTodo)
	todos.PUT("/:id", todoHandler.UpdateTodo)
	todos.DELETE("/:id", todoHandler.DeleteTodo)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}