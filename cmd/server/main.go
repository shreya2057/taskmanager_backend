package main

import (
	"todoapp/internal/config"
	"todoapp/internal/handlers"
	"todoapp/internal/middleware"
	"todoapp/internal/models"
	"todoapp/internal/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	config.DBConnect()

	models.Migrate(config.DB)

	taskRepo := repository.NewTaskRepository()
	tasks := handlers.NewTaskHandler(taskRepo)

	userRepo := repository.NewUserRepository()
	users := handlers.NewUserHandler(userRepo)

	e := echo.New()

	e.POST("/tasks", tasks.AddTasks)
	e.GET("/tasks", tasks.GetTasks)

	e.POST("/users", users.AddUser, middleware.UserMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}
