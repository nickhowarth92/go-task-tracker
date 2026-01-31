package repositories

import (
	"go-task-tracker/config"
	"go-task-tracker/models"
)

func GetAllUsers() []models.User {
	var users []models.User
	config.DB.Preload("Tasks").Find(&users)
	return users
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := config.DB.Preload("Tasks").First(&user, id)
	return user, result.Error
}

func CreateUser(user models.User) models.User {
	config.DB.Create(&user)
	return user
}

func UpdateUser(user models.User) models.User {
	config.DB.Save(&user)
	return user
}

func DeleteUser(id uint) error {
	result := config.DB.Delete(&models.User{}, id)
	return result.Error
}
