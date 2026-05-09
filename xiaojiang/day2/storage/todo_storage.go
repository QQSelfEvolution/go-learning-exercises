package storage

import (
	"sync"
	"time"

	"github.com/QQSelfEvolution/go-learning-exercises/xiaojiang/day2/models"
)

// TodoStorage implements in-memory storage for todos
type TodoStorage struct {
	mu    sync.RWMutex
	todos map[int]*models.Todo
	nextID int
}

// NewTodoStorage creates a new TodoStorage instance
func NewTodoStorage() *TodoStorage {
	storage := &TodoStorage{
		todos:  make(map[int]*models.Todo),
		nextID: 1,
	}
	// Add some sample todos
	storage.todos[1] = &models.Todo{
		ID:          1,
		Title:       "Learn Go",
		Description: "Study Go programming basics",
		Completed:   false,
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	storage.nextID = 2
	return storage
}

// GetAll returns all todos
func (s *TodoStorage) GetAll() []*models.Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*models.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos
}

// GetByID returns a todo by ID
func (s *TodoStorage) GetByID(id int) (*models.Todo, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, ok := s.todos[id]
	return todo, ok
}

// Create creates a new todo
func (s *TodoStorage) Create(req *models.CreateTodoRequest) *models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := &models.Todo{
		ID:          s.nextID,
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.todos[s.nextID] = todo
	s.nextID++
	return todo
}

// Update updates an existing todo
func (s *TodoStorage) Update(id int, req *models.UpdateTodoRequest) (*models.Todo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.todos[id]
	if !ok {
		return nil, false
	}

	if req.Title != "" {
		todo.Title = req.Title
	}
	if req.Description != "" {
		todo.Description = req.Description
	}
	if req.Completed != nil {
		todo.Completed = *req.Completed
	}
	if req.Priority != 0 {
		todo.Priority = req.Priority
	}
	todo.UpdatedAt = time.Now()

	return todo, true
}

// Delete deletes a todo by ID
func (s *TodoStorage) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return false
	}
	delete(s.todos, id)
	return true
}
