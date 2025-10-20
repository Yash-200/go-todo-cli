package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Yash-200/go-todo-cli/internal/database"
	"github.com/Yash-200/go-todo-cli/internal/models"
	"github.com/go-chi/chi/v5"
)

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameters from the URL
	filterName := r.URL.Query().Get("filter_name")
	filterStatus := r.URL.Query().Get("status") // Get the new status parameter
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")

	if order == "" {
		order = "asc"
	}

	// Use the updated function with the new parameter
	tasks, err := database.GetTasks(filterName, filterStatus, sortBy, order)
	if err != nil {
		log.Printf("Error getting tasks: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	// Decode the incoming JSON request body into the task struct
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if task.Name == "" {
		http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
		return
	}

	newTask, err := database.CreateTask(task.Name)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(newTask)
}

func UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	rowsAffected, err := database.UpdateTaskStatus(id)
	if err != nil {
		log.Printf("Error updating task status: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	rowsAffected, err := database.DeleteTask(id)
	if err != nil {
		log.Printf("Error deleting task: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
