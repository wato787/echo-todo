package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"echo-todo/internal/repository"
	"echo-todo/pkg/models"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
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
	// Get todo from repository
	todo, err := s.todoRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	// Return nil if todo not found (repository returns nil, nil for not found)
	if todo == nil {
		return nil, nil
	}
	
	return todo, nil
}

func (s *todoService) GetAllTodos(ctx context.Context) ([]models.Todo, error) {
	todos,err := s.todoRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, id string, req *models.UpdateTodoRequest) (*models.Todo, error) {
	// Get existing todo
	existingTodo, err := s.todoRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	// Check if todo exists
	if existingTodo == nil {
		return nil, nil
	}
	
	// Update fields if provided
	if req.Title != nil {
		existingTodo.Title = *req.Title
	}
	if req.Description != nil {
		existingTodo.Description = *req.Description
	}
	if req.Completed != nil {
		existingTodo.Completed = *req.Completed
	}
	
	// Update timestamp
	existingTodo.UpdatedAt = time.Now()
	
	// Save updated todo
	err = s.todoRepo.Update(ctx, existingTodo)
	if err != nil {
		return nil, err
	}
	
	return existingTodo, nil
}

func (s *todoService) DeleteTodo(ctx context.Context, id string) error {
	// Check if todo exists before deletion
	existingTodo, err := s.todoRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	// Return specific error if todo not found
	if existingTodo == nil {
		return ErrTodoNotFound
	}
	
	// Delete todo from repository
	err = s.todoRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	
	return nil
}

func generateID() string {
	// Generate UUID v4 for unique ID
	return uuid.New().String()
}