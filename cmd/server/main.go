package main

import (
	"todoapp/internal/config"
	"todoapp/internal/handlers"
	"todoapp/internal/models"
	"todoapp/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	config.DBConnect()
	validate := validator.New()

	models.Migrate(config.DB)

	taskRepo := repository.NewTaskRepository()
	tasks := handlers.NewTaskHandler(taskRepo, validate)

	userRepo := repository.NewUserRepository()
	users := handlers.NewUserHandler(userRepo, validate)

	e := echo.New()

	e.POST("/tasks", tasks.AddTasks)
	e.GET("/tasks", tasks.GetTasks)

	e.POST("/users", users.AddUser)

	e.Logger.Fatal(e.Start(":8080"))
}
