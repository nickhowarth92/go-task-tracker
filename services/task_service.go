package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-task-tracker/models"
	"go-task-tracker/repositories"
)

func HandleGetTasks() (interface{}, int) {
	tasks := repositories.GetAllTasks()
	return tasks, http.StatusOK
}

func HandleCreateTask(r *http.Request) (interface{}, int) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		return map[string]string{"error": "Invalid JSON"}, http.StatusBadRequest
	}
	repositories.CreateTask(task)
	return task, http.StatusCreated
}

func HandleUpdateTask(r *http.Request, idStr string) (interface{}, int) {
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return map[string]string{"error": "Invalid ID"}, http.StatusBadRequest
	}
	id := uint(id64)

	var updated models.Task
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		return map[string]string{"error": "Invalid JSON"}, http.StatusBadRequest
	}

	task, err := repositories.GetTaskByID(id)
	if err != nil {
		return map[string]string{"error": "Task not found"}, http.StatusNotFound
	}

	task.Title = updated.Title
	task.Description = updated.Description
	task.Completed = updated.Completed
	task.UserID = updated.UserID

	repositories.UpdateTask(task)
	return task, http.StatusOK
}

func HandleDeleteTask(idStr string) (interface{}, int) {
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return map[string]string{"error": "Invalid ID"}, http.StatusBadRequest
	}
	id := uint(id64)

	if err := repositories.DeleteTask(id); err != nil {
		return map[string]string{"error": "Failed to delete task"}, http.StatusInternalServerError
	}

	return map[string]string{"message": "Task deleted"}, http.StatusNoContent
}
