package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"echo-todo/internal/services"
	"echo-todo/pkg/models"
	"echo-todo/pkg/utils"
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
// @Summary Create a new TODO
// @Description Create a new TODO item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.CreateTodoRequest true "Create TODO request"
// @Success 201 {object} utils.Response{data=models.Todo} "Successfully created"
// @Failure 400 {object} utils.Response "Bad request"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router /api/v1/todos [post]
func (h *TodoHandler) CreateTodo(c echo.Context) error {
	var req models.CreateTodoRequest
	
	// Bind request body
	if err := c.Bind(&req); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid request format")
	}
	
	// Validate request
	if err := utils.ValidateStruct(&req); err != nil {
		return utils.ValidationErrorResponse(c, err.Error())
	}
	
	// Create todo via service
	todo, err := h.todoService.CreateTodo(c.Request().Context(), &req)
	if err != nil {
		return utils.InternalErrorResponse(c, "Failed to create todo")
	}
	
	return utils.SuccessResponse(c, http.StatusCreated, "Todo created successfully", todo)
}

// GetTodo retrieves a todo by ID
// @Summary Get a TODO by ID
// @Description Get a specific TODO item by ID
// @Tags todos
// @Produce json
// @Param id path string true "TODO ID"
// @Success 200 {object} utils.Response{data=models.Todo} "Successfully retrieved"
// @Failure 400 {object} utils.Response "Bad request"
// @Failure 404 {object} utils.Response "TODO not found"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router /api/v1/todos/{id} [get]
func (h *TodoHandler) GetTodo(c echo.Context) error {
	// Get ID from URL parameter
	id := c.Param("id")
	if id == "" {
		return utils.ValidationErrorResponse(c, "ID is required")
	}
	
	// Get todo via service
	todo, err := h.todoService.GetTodoByID(c.Request().Context(), id)
	if err != nil {
		return utils.InternalErrorResponse(c, "Failed to get todo")
	}
	
	// Check if todo was found
	if todo == nil {
		return utils.NotFoundResponse(c, "Todo not found")
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Todo retrieved successfully", todo)
}

// GetAllTodos retrieves all todos
// @Summary Get all TODOs
// @Description Get all TODO items
// @Tags todos
// @Produce json
// @Success 200 {object} utils.Response{data=[]models.Todo} "Successfully retrieved"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router /api/v1/todos [get]
func (h *TodoHandler) GetAllTodos(c echo.Context) error {

	todos,err := h.todoService.GetAllTodos(c.Request().Context())
	if err != nil {
		return utils.InternalErrorResponse(c,"faild")
	}

	return utils.SuccessResponse(c,http.StatusOK,"OK",todos)
}

// UpdateTodo updates an existing todo
// @Summary Update a TODO
// @Description Update an existing TODO item
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "TODO ID"
// @Param todo body models.UpdateTodoRequest true "Update TODO request"
// @Success 200 {object} utils.Response{data=models.Todo} "Successfully updated"
// @Failure 400 {object} utils.Response "Bad request"
// @Failure 404 {object} utils.Response "TODO not found"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router /api/v1/todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	// Get ID from URL parameter
	id := c.Param("id")
	if id == "" {
		return utils.ValidationErrorResponse(c, "ID is required")
	}
	
	var req models.UpdateTodoRequest
	
	// Bind request body
	if err := c.Bind(&req); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid request format")
	}
	
	// Validate that at least one field is provided
	if req.Title == nil && req.Description == nil && req.Completed == nil {
		return utils.ValidationErrorResponse(c, "At least one field must be provided for update")
	}
	
	// Update todo via service
	todo, err := h.todoService.UpdateTodo(c.Request().Context(), id, &req)
	if err != nil {
		return utils.InternalErrorResponse(c, "Failed to update todo")
	}
	
	// Check if todo was found
	if todo == nil {
		return utils.NotFoundResponse(c, "Todo not found")
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Todo updated successfully", todo)
}

// DeleteTodo deletes a todo by ID
// @Summary Delete a TODO
// @Description Delete a TODO item by ID
// @Tags todos
// @Produce json
// @Param id path string true "TODO ID"
// @Success 200 {object} utils.Response "Successfully deleted"
// @Failure 400 {object} utils.Response "Bad request"
// @Failure 404 {object} utils.Response "TODO not found"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router /api/v1/todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	// Get ID from URL parameter
	id := c.Param("id")
	if id == "" {
		return utils.ValidationErrorResponse(c, "ID is required")
	}
	
	// Delete todo via service
	err := h.todoService.DeleteTodo(c.Request().Context(), id)
	if err != nil {
		// Check for specific error types
		if err.Error() == "todo not found" {
			return utils.NotFoundResponse(c, "Todo not found")
		}
		return utils.InternalErrorResponse(c, "Failed to delete todo")
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Todo deleted successfully", nil)
}