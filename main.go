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
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"echo-todo/internal/config"
	"echo-todo/internal/handlers"
	"echo-todo/internal/repository"
	"echo-todo/internal/services"
	_ "echo-todo/docs"
)

var echoApp *echo.Echo

func init() {
	// Initialize Echo app once during cold start
	cfg := config.Load()

	todoRepo, err := repository.NewDynamoDBTodoRepository(cfg.TableName)
	if err != nil {
		panic("Failed to initialize todo repository: " + err.Error())
	}

	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	echoApp = echo.New()
	echoApp.Use(middleware.Logger())
	echoApp.Use(middleware.Recover())
	echoApp.Use(middleware.CORS())

	// Root endpoints
	echoApp.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Echo TODO API on AWS Lambda",
			"status":  "ready",
		})
	})

	echoApp.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	// Swagger endpoint
	echoApp.GET("/swagger/*", echoSwagger.WrapHandler)

	// API routes
	api := echoApp.Group("/api/v1")

	// TODO routes
	todos := api.Group("/todos")
	todos.POST("", todoHandler.CreateTodo)
	todos.GET("", todoHandler.GetAllTodos)
	todos.GET("/:id", todoHandler.GetTodo)
	todos.PUT("/:id", todoHandler.UpdateTodo)
	todos.DELETE("/:id", todoHandler.DeleteTodo)
}

func handler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	// Function URLの場合、RawPathを使用
	path := req.RawPath
	if path == "" {
		path = "/"
	}
	
	// Debug logging
	println("DEBUG - HTTP Method:", req.RequestContext.HTTP.Method)
	println("DEBUG - RawPath:", req.RawPath)
	println("DEBUG - Final Path:", path)
	
	// Convert Lambda Function URL request to HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, req.RequestContext.HTTP.Method, path, strings.NewReader(req.Body))
	if err != nil {
		return events.LambdaFunctionURLResponse{StatusCode: 500}, err
	}

	// Set headers
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	// Set query parameters
	q := httpReq.URL.Query()
	for key, value := range req.QueryStringParameters {
		q.Add(key, value)
	}
	httpReq.URL.RawQuery = q.Encode()

	// Create response recorder
	w := httptest.NewRecorder()

	// Process request with Echo
	echoApp.ServeHTTP(w, httpReq)

	// Convert response
	resp := events.LambdaFunctionURLResponse{
		StatusCode: w.Code,
		Headers:    make(map[string]string),
		Body:       w.Body.String(),
	}

	// Convert headers
	for key, values := range w.Header() {
		resp.Headers[key] = strings.Join(values, ",")
	}

	// Handle binary content if needed
	if isBinary(w.Header().Get("Content-Type")) {
		resp.Body = base64.StdEncoding.EncodeToString(w.Body.Bytes())
		resp.IsBase64Encoded = true
	}

	return resp, nil
}

func isBinary(contentType string) bool {
	// Define binary content types
	binaryTypes := []string{
		"image/",
		"application/pdf",
		"application/zip",
		"application/gzip",
		"application/octet-stream",
	}

	for _, binaryType := range binaryTypes {
		if strings.HasPrefix(contentType, binaryType) {
			return true
		}
	}
	return false
}

func main() {
	lambda.Start(handler)
}