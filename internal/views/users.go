package views

type CreateUser struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name" validate:"required"`
	Email          string `json:"email" validate:"email"`
	Password       string `json:"password" validate:"password"`
	FullName       string `json:"full_name" validate:"required"`
	Role           string `json:"role"`
	IsActive       bool   `json:"is_active"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}

type UpdateUser struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name" validate:"omitempty,required"`
	Email          string `json:"email" validate:"omitempty,email"`
	Password       string `json:"password" validate:"omitempty,password"`
	FullName       string `json:"full_name" validate:"omitempty"`
	Role           string `json:"role" validate:"oneof=admin user,omitempty"`
	IsActive       bool   `json:"is_active"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}
