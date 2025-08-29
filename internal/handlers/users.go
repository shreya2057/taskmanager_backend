package handlers

import (
	"todoapp/internal/models"
	"todoapp/internal/repository"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) AddUser(c echo.Context) error {
	user := new(models.User)

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}
	err = h.repo.CreateUser(user)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create user"})
	}
	return c.JSON(201, user)
}
