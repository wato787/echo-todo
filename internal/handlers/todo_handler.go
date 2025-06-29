package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"echo-todo/internal/services"
)

type TodoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

// CreateTodo creates a new todo
func (h *TodoHandler) CreateTodo(c echo.Context) error {
	// TODO: Implement create todo logic
	return c.JSON(http.StatusCreated, map[string]string{"message": "TODO: Implement CreateTodo"})
}

// GetTodo retrieves a todo by ID
func (h *TodoHandler) GetTodo(c echo.Context) error {
	// TODO: Implement get todo logic
	return c.JSON(http.StatusOK, map[string]string{"message": "TODO: Implement GetTodo"})
}

// GetAllTodos retrieves all todos
func (h *TodoHandler) GetAllTodos(c echo.Context) error {
	// TODO: Implement get all todos logic
	return c.JSON(http.StatusOK, map[string]string{"message": "TODO: Implement GetAllTodos"})
}

// UpdateTodo updates an existing todo
func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	// TODO: Implement update todo logic
	return c.JSON(http.StatusOK, map[string]string{"message": "TODO: Implement UpdateTodo"})
}

// DeleteTodo deletes a todo by ID
func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	// TODO: Implement delete todo logic
	return c.JSON(http.StatusOK, map[string]string{"message": "TODO: Implement DeleteTodo"})
}