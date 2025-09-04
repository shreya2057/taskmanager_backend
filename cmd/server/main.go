package main

import (
	"fmt"
	"log"
	"os"
	"todoapp/internal/config"
	"todoapp/internal/handlers"
	"todoapp/internal/models"
	"todoapp/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: failed to load .env file, continuing with system env vars")
		} else {
			log.Println(".env file loaded")
		}
	}
}

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

	e.GET("/users", users.GetAllUsers)
	e.GET("/users/:id", users.GetSingleUser)
	e.POST("/users", users.AddUser)
	e.PATCH("/users/:id", users.UpdateUser)
	e.DELETE("/users/:id", users.DeleteUser)

	e.POST("/upload-image", handlers.NewUploadHandler(*validate).GetPresignedURL)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
