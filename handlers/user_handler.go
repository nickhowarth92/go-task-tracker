package handlers

import (
	"encoding/json"
	"go-task-tracker/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	resp, status := services.HandleGetUsers()
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	resp, status := services.HandleCreateUser(r)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	resp, status := services.HandleUpdateUser(r, id)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	resp, status := services.HandleDeleteUser(id)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
