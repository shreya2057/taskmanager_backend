package models

type User struct {
	ID             string `gorm:"primaryKey; size:20" json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	FullName       string `json:"full_name"`
	Role           string `json:"role"`
	Status         bool   `json:"status"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}
