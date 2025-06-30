package services

import (
	"context"
	"fmt"
	"time"

	"echo-todo/internal/repository"
	"echo-todo/pkg/models"
)

type TodoService interface {
	CreateTodo(ctx context.Context, req *models.CreateTodoRequest) (*models.Todo, error)
	GetTodoByID(ctx context.Context, id string) (*models.Todo, error)
	GetAllTodos(ctx context.Context) ([]models.Todo, error)
	UpdateTodo(ctx context.Context, id string, req *models.UpdateTodoRequest) (*models.Todo, error)
	DeleteTodo(ctx context.Context, id string) error
}

type todoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (s *todoService) CreateTodo(ctx context.Context, req *models.CreateTodoRequest) (*models.Todo, error) {
	// Generate unique ID
	id := generateID()
	
	// Create todo entity with timestamps
	todo := &models.Todo{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	// Save to repository
	err := s.todoRepo.Create(ctx, todo)
	if err != nil {
		return nil, err
	}
	
	return todo, nil
}

func (s *todoService) GetTodoByID(ctx context.Context, id string) (*models.Todo, error) {
	// TODO: Implement business logic for getting todo by ID
	return nil, nil
}

func (s *todoService) GetAllTodos(ctx context.Context) ([]models.Todo, error) {
	// TODO: Implement business logic for getting all todos
	return nil, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, id string, req *models.UpdateTodoRequest) (*models.Todo, error) {
	// TODO: Implement business logic for updating todo
	return nil, nil
}

func (s *todoService) DeleteTodo(ctx context.Context, id string) error {
	// TODO: Implement business logic for deleting todo
	return nil
}

func generateID() string {
	// Simple ID generation using timestamp
	return fmt.Sprintf("todo_%d", time.Now().UnixNano())
}