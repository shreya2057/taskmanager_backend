package repository

import (
	"context"
	"time"
	"todoapp/internal/config"
	"todoapp/internal/models"
)

type authRepoGorm struct{}

type AuthRepository interface {
	LoginUser(loginDetail *models.Login) (*models.User, error)
}

func NewAuthRepository() AuthRepository {
	return &authRepoGorm{}
}

func (r *authRepoGorm) LoginUser(loginDetail *models.Login) (*models.User, error) {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := config.DB.WithContext(ctx).Where("email = ?", loginDetail.Email).
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}
