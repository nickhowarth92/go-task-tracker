package handlers

import (
	"encoding/json"
	"go-task-tracker/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	resp, status := services.HandleGetTasks()
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	resp, status := services.HandleCreateTask(r)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	resp, status := services.HandleUpdateTask(r, id)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	resp, status := services.HandleDeleteTask(id)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
