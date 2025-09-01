package services

import (
	"fmt"
	"todoapp/internal/models"
	"todoapp/internal/repository"

	"gorm.io/gorm"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo: repo}
}

func (s *UserService) FindExistingUser(user *models.User) (int, error) {

	check := func(field, fieldName string) (int, error) {
		existing, err := s.repo.FindExistingUser(field, fieldName)
		if err != nil && err != gorm.ErrRecordNotFound {
			return 500, err
		}
		if existing != nil {
			return 400, fmt.Errorf("user with this %s already exists", fieldName)
		}
		return 0, nil
	}
	if code, err := check(user.Email, "email"); err != nil {
		return code, err
	}
	if code, err := check(user.UserName, "user_name"); err != nil {
		return code, err
	}

	return 500, s.repo.CreateUser(user)
}
