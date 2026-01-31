package routes

import (
	"go-task-tracker/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes registers all API endpoints and returns a configured router
func SetupRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// ---- API Routes ----
	api := router.PathPrefix("/api").Subrouter()

	// Task routes
	api.HandleFunc("/tasks", handlers.GetTasks).Methods(http.MethodGet)
	api.HandleFunc("/tasks", handlers.CreateTask).Methods(http.MethodPost)
	api.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods(http.MethodPut)
	api.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods(http.MethodDelete)

	// User routes
	api.HandleFunc("/users", handlers.GetUsers).Methods(http.MethodGet)
	api.HandleFunc("/users", handlers.CreateUser).Methods(http.MethodPost)
	api.HandleFunc("/users/{id}", handlers.UpdateUser).Methods(http.MethodPut)
	api.HandleFunc("/users/{id}", handlers.DeleteUser).Methods(http.MethodDelete)

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	}).Methods(http.MethodGet)

	return router
}
