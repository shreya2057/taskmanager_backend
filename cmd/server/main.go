package main

import (
	"todoapp/internal/config"
	"todoapp/internal/handlers"
	"todoapp/internal/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	config.DBConnect()

	repo := repository.NewTaskRepository()
	tasks := handlers.NewTaskHandler(repo)

	e := echo.New()
	e.POST("/tasks", tasks.AddTasks)
	e.GET("/tasks", tasks.GetTasks)

	e.Logger.Fatal(e.Start(":8080"))
}
