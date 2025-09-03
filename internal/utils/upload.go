package utils

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func BoolPtr(b bool) *bool { return &b }

func UploadImage(fileHeader *multipart.FileHeader, folderName string) (string, error) {

	if fileHeader == nil {
		return "", nil
	}
	cloudName := os.Getenv("CLOUD_NAME")
	cloudApiKey := os.Getenv("CLOUD_API_KEY")
	cloudApiSecret := os.Getenv("CLOUD_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, cloudApiKey, cloudApiSecret)

	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	uploadParams := uploader.UploadParams{
		Folder:         folderName,
		UniqueFilename: BoolPtr(true),
		Overwrite:      BoolPtr(false),
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		log.Fatalf("Failed to upload image: %v", err)
	}
	return uploadResult.SecureURL, nil
}
