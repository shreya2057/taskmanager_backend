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
	repo     repository.UserRepository
	validate validator.Validate
}

func NewUserHandler(repo repository.UserRepository, validate *validator.Validate) *UserHandler {
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		pass := fl.Field().String()
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pass)
		hasNumber := regexp.MustCompile(`[0-9]`).MatchString(pass)
		hasSymbol := regexp.MustCompile(`[!@#~$%^&*()+|_.,<>?/{}\-]`).MatchString(pass)
		return len(pass) >= 8 && hasUpper && hasNumber && hasSymbol
	})
	return &UserHandler{repo: repo, validate: *validate}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.repo.GetAllUsers()

	userList := make([]views.GetUsers, len(users))
	for i, user := range users {
		userList[i] = views.GetUsers{
			ID:             user.ID,
			UserName:       user.UserName,
			Email:          user.Email,
			FullName:       user.FullName,
			Role:           user.Role,
			IsActive:       user.IsActive,
			ProfilePicture: user.ProfilePicture,
			CreatedAt:      user.CreatedAt,
			UpdatedAt:      user.UpdatedAt,
		}
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Message: "Could not fetch users", Errors: err.Error()})
	}
	return c.JSON(http.StatusOK, utils.Response{Message: "Users fetched successfully", Data: userList})
}

func (h *UserHandler) AddUser(c echo.Context) error {
	var user views.CreateUser

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, utils.Response{Message: "Invalid request", Errors: err.Error()})
	}

	code, err := utils.Validate(&h.validate, &user)

	if err != nil {
		return c.JSON(code, utils.Response{
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
	}

	if err := h.repo.CreateUser(&userModal); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Message: "User cannot be created", Errors: err.Error()})
	}
	return c.JSON(201, utils.Response{Message: "User created successfully"})
}

func (h *UserHandler) UpdateUser(c echo.Context) error {

	existingUser, err := h.repo.FindExistingUser(c.Param("id"), "id")
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.Response{Message: "User not found", Errors: "No user with the given ID"})
	}

	if existingUser == nil {
		return c.JSON(http.StatusNotFound, utils.Response{Message: "User not found", Errors: "No user with the given ID"})
	}

	var user views.UpdateUser
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, utils.Response{Message: "Invalid request", Errors: err.Error()})
	}

	if code, err := utils.Validate(&h.validate, &user); err != nil {
		return c.JSON(code, utils.Response{
			Message: "validation failed",
			Errors:  err,
		})
	}

	userModal := &models.User{
		ID:             c.Param("id"),
		UserName:       user.UserName,
		Email:          user.Email,
		Password:       user.Password,
		FullName:       user.FullName,
		Role:           user.Role,
		IsActive:       user.IsActive,
		ProfilePicture: user.ProfilePicture,
		UpdatedAt:      time.Now().Local().String(),
	}

	if err := h.repo.UpdateUser(userModal); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Message: "User cannot be updated", Errors: "Internal server error"})
	}
	return c.JSON(200, utils.Response{Message: "User updated successfully"})

}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	existingUser, err := h.repo.FindExistingUser(c.Param("id"), "ID")

	if err != nil {
		return c.JSON(http.StatusNotFound, utils.Response{Message: "User not found", Errors: "No user with the given ID"})
	}

	if existingUser == nil {
		return c.JSON(http.StatusNotFound, utils.Response{Message: "User not found", Errors: "No user with the given ID"})
	}

	if err := h.repo.DeleteUser(c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Message: "User cannot be deleted", Errors: "Internal server error"})
	}
	return c.JSON(200, utils.Response{Message: "User deleted successfully"})
}
