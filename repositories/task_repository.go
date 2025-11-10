package repositories

import (
	"go-task-tracker/config"
	"go-task-tracker/models"
)

func GetAllTasks() []models.Task {
	var tasks []models.Task
	config.DB.Preload("User").Find(&tasks)
	return tasks
}

func GetTaskByID(id uint) (models.Task, error) {
	var task models.Task
	result := config.DB.Preload("User").First(&task, id)
	return task, result.Error
}

func CreateTask(task models.Task) models.Task {
	config.DB.Create(&task)
	return task
}

func UpdateTask(task models.Task) models.Task {
	config.DB.Save(&task)
	return task
}

func DeleteTask(id uint) error {
	result := config.DB.Delete(&models.Task{}, id)
	return result.Error
}
