package repository

import (
	"context"
	"fmt"
	"time"
	"todoapp/internal/config"
	"todoapp/internal/models"

	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type userRepoGorm struct{}

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
	FindExistingUser(identifier, fieldName string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

func NewUserRepository() UserRepository {
	return &userRepoGorm{}
}

func (r *userRepoGorm) GetAllUsers() ([]models.User, error) {
	var users []models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := config.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepoGorm) CreateUser(user *models.User) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.ID = xid.New().String()
	user.Password = string(hashPassword)
	user.IsActive = true

	// Sets a timeout for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Adds a new user to the database
	return config.DB.WithContext(ctx).Create(user).Error
}

func (r *userRepoGorm) FindExistingUser(identifier, fieldName string) (*models.User, error) {

	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := config.DB.WithContext(ctx).Where(fmt.Sprintf("%s = ?", fieldName), identifier).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepoGorm) UpdateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return config.DB.WithContext(ctx).
		Model(&models.User{}).Where("id = ?", user.ID).
		Updates(user).Error
}

func (r *userRepoGorm) DeleteUser(id string) error {

	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return config.DB.WithContext(ctx).Delete(&user, "id = ?", id).Error
}
