package repository

import (
	"context"
	"time"
	"todoapp/internal/config"
	"todoapp/internal/models"
)

// taskRepoGorm implements TaskRepository using GORM
type taskRepoGorm struct{}

// TaskRepository defines the methods for task data access
type TaskRepository interface {
	CreateTask(tasks *models.Task) error
	GetTasks() ([]models.Task, error)
}

// NewTaskRepository creates a new instance of TaskRepository
func NewTaskRepository() TaskRepository {
	return &taskRepoGorm{}
}

func (r *taskRepoGorm) CreateTask(tasks *models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Adds a new task to the database
	return config.DB.WithContext(ctx).Create(tasks).Error
}

func (r *taskRepoGorm) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := config.DB.WithContext(ctx).Find(&tasks).Error
	return tasks, err
}
