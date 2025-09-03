package views

import "mime/multipart"

type ImageUpload struct {
	Image    *multipart.FileHeader `json:"image" form:"image" validate:"required"`
	Category string                `json:"category" form:"category" validate:"required"`
}
