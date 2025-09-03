package handlers

import (
	"fmt"
	"os"
	"todoapp/internal/utils"
	"todoapp/internal/views"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UploadHandler struct {
	validate validator.Validate
}

func NewUploadHandler(validate validator.Validate) *UploadHandler {
	return &UploadHandler{validate: validate}
}

func (r *UploadHandler) GetPresignedURL(c echo.Context) error {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return c.JSON(400, utils.Response{Message: "Invalid request", Errors: err.Error()})
	}

	image := views.ImageUpload{
		Image:    fileHeader,
		Category: c.FormValue("category"),
	}

	if code, err := utils.Validate(&r.validate, &image); err != nil {
		return c.JSON(code, utils.Response{
			Message: "validation failed",
			Errors:  err,
		})
	}

	imageURL, err := utils.UploadImage(fileHeader, fmt.Sprintf("%s/%s", os.Getenv("FOLDER_NAME"), image.Category))
	if err != nil {
		return c.JSON(500, utils.Response{Message: "Failed to upload image", Errors: err.Error()})
	}

	return c.JSON(200, utils.Response{
		Message: "Image uploaded successfully",
		Data:    map[string]string{"image_url": imageURL},
	})
}
