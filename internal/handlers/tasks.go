package handlers

import (
	"net/http"
	"todoapp/internal/models"
	"todoapp/internal/repository"

	"github.com/labstack/echo/v4"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	repo repository.TaskRepository
}

// NewTaskHandler creates a new instance of TaskHandler
func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

// AddTasks handles the creation of a new task
func (h *TaskHandler) AddTasks(c echo.Context) error {

	// Create a new Task instance
	tasks := new(models.Task)

	// Bind the incoming JSON to the Task struct
	err := c.Bind(&tasks)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	err = h.repo.CreateTask(tasks)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create task"})
	}
	return c.JSON(http.StatusCreated, tasks)

}

func (h *TaskHandler) GetTasks(c echo.Context) error {

	// Fetch tasks from the repository
	tasks, err := h.repo.GetTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}
