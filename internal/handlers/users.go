package handlers

import (
	"todoapp/internal/models"
	"todoapp/internal/repository"
	"todoapp/services"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo    repository.UserRepository
	service services.UserService
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo, service: services.NewUserService(repo)}
}

func (h *UserHandler) AddUser(c echo.Context) error {
	user := c.Get("user").(*models.User)

	code, err := h.service.FindExistingUser(user)
	if err != nil {
		return c.JSON(code, Response{Message: "User cannot be created", Errors: err.Error()})
	}
	return c.JSON(201, Response{Message: "User created successfully"})
}
