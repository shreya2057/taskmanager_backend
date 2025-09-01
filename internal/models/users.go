package models

type User struct {
	ID             string `gorm:"primaryKey; size:20" json:"id"`
	UserName       string `gorm:"not null" json:"user_name" validate:"required"`
	Email          string `gorm:"uniqueIndex" json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,password"`
	FullName       string `json:"full_name" validate:"required"`
	Role           string `json:"role" validate:"required,oneof=admin user"`
	IsActive       bool   `json:"is_active"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}
