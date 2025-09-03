package models

import "gorm.io/gorm"

type User struct {
	ID             string         `gorm:"primaryKey; size:20; uniqueIndex" json:"id"`
	UserName       string         `gorm:"not null; uniqueIndex" json:"user_name"`
	Email          string         `gorm:"uniqueIndex" json:"email"`
	Password       string         `json:"password"`
	FullName       string         `json:"full_name"`
	Role           string         `json:"role"`
	IsActive       bool           `json:"is_active"`
	ProfilePicture string         `json:"profile_picture"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
