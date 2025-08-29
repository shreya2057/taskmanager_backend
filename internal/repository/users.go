package repository

import (
	"context"
	"time"
	"todoapp/internal/config"
	"todoapp/internal/models"

	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type userRepoGorm struct{}

type UserRepository interface {
	CreateUser(user *models.User) error
}

func NewUserRepository() UserRepository {
	return &userRepoGorm{}
}

func (r *userRepoGorm) CreateUser(user *models.User) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.ID = xid.New().String()
	user.Password = string(hashPassword)

	// Sets a timeout for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Adds a new user to the database
	return config.DB.WithContext(ctx).Create(user).Error
}
