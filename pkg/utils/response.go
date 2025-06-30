package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse returns a success response
func SuccessResponse(c echo.Context, code int, message string, data interface{}) error {
	return c.JSON(code, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse returns an error response
func ErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, Response{
		Success: false,
		Error:   message,
	})
}

// ValidationErrorResponse returns a validation error response
func ValidationErrorResponse(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusBadRequest, message)
}

// NotFoundResponse returns a not found response
func NotFoundResponse(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusNotFound, message)
}

// InternalErrorResponse returns an internal server error response
func InternalErrorResponse(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusInternalServerError, message)
}