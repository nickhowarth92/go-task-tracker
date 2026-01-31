package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-task-tracker/models"
	"go-task-tracker/repositories"
)

func HandleGetUsers() (interface{}, int) {
	users := repositories.GetAllUsers()
	return users, http.StatusOK
}

func HandleCreateUser(r *http.Request) (interface{}, int) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return map[string]string{"error": "Invalid JSON"}, http.StatusBadRequest
	}
	repositories.CreateUser(user)
	return user, http.StatusCreated
}

func HandleUpdateUser(r *http.Request, idStr string) (interface{}, int) {
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return map[string]string{"error": "Invalid ID"}, http.StatusBadRequest
	}
	id := uint(id64)

	var updated models.User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		return map[string]string{"error": "Invalid JSON"}, http.StatusBadRequest
	}

	user, err := repositories.GetUserByID(id)
	if err != nil {
		return map[string]string{"error": "User not found"}, http.StatusNotFound
	}

	user.Name = updated.Name
	user.Email = updated.Email

	repositories.UpdateUser(user)
	return user, http.StatusOK
}

func HandleDeleteUser(idStr string) (interface{}, int) {
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return map[string]string{"error": "Invalid ID"}, http.StatusBadRequest
	}
	id := uint(id64)

	if err := repositories.DeleteUser(id); err != nil {
		return map[string]string{"error": "Failed to delete user"}, http.StatusInternalServerError
	}

	return map[string]string{"message": "User deleted"}, http.StatusNoContent
}
