package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/QQSelfEvolution/go-learning-exercises/xiaojiang/day2/models"
	"github.com/QQSelfEvolution/go-learning-exercises/xiaojiang/day2/handlers"
	"github.com/QQSelfEvolution/go-learning-exercises/xiaojiang/day2/middleware"
	"github.com/QQSelfEvolution/go-learning-exercises/xiaojiang/day2/storage"
)

func main() {
	// Initialize storage
	todoStorage := storage.NewTodoStorage()

	// Initialize handlers
	todoHandler := handlers.NewTodoHandler(todoStorage)

	// Setup routes
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/api/health", handlers.HealthCheck)

	// Todo routes
	mux.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.GetAllTodos(w, r)
		case http.MethodPost:
			todoHandler.CreateTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/todos/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/api/todos/"):]
		switch r.Method {
		case http.MethodGet:
			todoHandler.GetTodo(w, r, id)
		case http.MethodPut:
			todoHandler.UpdateTodo(w, r, id)
		case http.MethodDelete:
			todoHandler.DeleteTodo(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Apply middleware
	handler := middleware.Logger(middleware.ErrorHandler(middleware.CORS(middleware.ContentType(mux))))

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
