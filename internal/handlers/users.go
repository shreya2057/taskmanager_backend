package handlers

import (
	"net/http"
	"regexp"
	"time"
	"todoapp/internal/models"
	"todoapp/internal/repository"
	"todoapp/internal/utils"
	"todoapp/internal/views"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo      repository.UserRepository
	validator validator.Validate
}

func NewUserHandler(repo repository.UserRepository, validator *validator.Validate) *UserHandler {
	return &UserHandler{repo: repo, validator: *validator}
}

func (h *UserHandler) AddUser(c echo.Context) error {
	var user views.CreateUser

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, Response{Message: "Invalid request", Errors: err.Error()})
	}

	h.validator.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		pass := fl.Field().String()
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pass)
		hasNumber := regexp.MustCompile(`[0-9]`).MatchString(pass)
		hasSymbol := regexp.MustCompile(`[!@#~$%^&*()+|_.,<>?/{}\-]`).MatchString(pass)
		return len(pass) >= 8 && hasUpper && hasNumber && hasSymbol
	})

	code, err := utils.Validate(&h.validator, &user)

	if err != nil {
		return c.JSON(code, Response{
			Message: "validation failed",
			Errors:  err,
		})
	}

	userModal := models.User{
		ID:             user.ID,
		UserName:       user.UserName,
		Email:          user.Email,
		Password:       user.Password,
		FullName:       user.FullName,
		Role:           user.Role,
		IsActive:       user.IsActive,
		ProfilePicture: user.ProfilePicture,
		CreatedAt:      time.Now().Local().String(),
		UpdatedAt:      time.Now().Local().String(),
		DeletedAt:      user.DeletedAt,
	}

	if err := h.repo.CreateUser(&userModal); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "User cannot be created", Errors: err.Error()})
	}
	return c.JSON(201, Response{Message: "User created successfully"})
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, Response{Message: "Invalid request", Errors: err.Error()})
	}
	err := h.repo.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "User cannot be updated", Errors: "Internal server error"})
	}
	return c.JSON(200, Response{Message: "User updated successfully"})

}
