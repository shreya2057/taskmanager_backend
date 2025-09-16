package handlers

import (
	"net/http"
	"os"
	"todoapp/internal/models"
	"todoapp/internal/repository"
	"todoapp/internal/utils"
	"todoapp/internal/views"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	repo repository.AuthRepository
}

func NewAuthHandler(repo repository.AuthRepository) *AuthHandler {
	return &AuthHandler{repo: repo}
}

func (r *AuthHandler) UserLogin(c echo.Context) error {
	loginDetail := new(views.LoginDetails)
	err := c.Bind(&loginDetail)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Message: "Login Failed", Errors: "Invalid Request"})
	}

	existing, err := r.repo.LoginUser(&models.Login{
		Email:    loginDetail.Email,
		Password: loginDetail.Password,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Message: "Login Failed", Errors: "User doesn't exists"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(loginDetail.Password))

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Message: "Login Failed", Errors: "Password doesn't match"})
	}

	accessToken, err := utils.CreateToken(existing, os.Getenv("ACCESS_PRIVATE_KEY"), 15)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Message: "Login Failed", Errors: err.Error()})
	}
	refreshToken, err := utils.CreateToken(existing, os.Getenv("REFRESH_PRIVATE_KEY"), 60)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Message: "Login Failed", Errors: "Refresh token error"})
	}

	return c.JSON(http.StatusOK, utils.Response{Message: "Logged in successfully", Data: views.TokenDetails{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}})

}
