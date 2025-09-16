package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
	"time"
	"todoapp/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user *models.User, secretKey string, expiryTime int) (*string, error) {

	tokenClaims := jwt.MapClaims{
		"email":     user.Email,
		"user_name": user.UserName,
		"id":        user.ID,
		"exp":       time.Now().Add(time.Minute * time.Duration(expiryTime)).Unix(),
	}

	// Fix: handle escaped newlines from .env
	pemKey := strings.ReplaceAll(secretKey, `\n`, "\n")

	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, fmt.Errorf("invalid ECDSA PEM key")
	}

	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	tokenDetails := jwt.NewWithClaims(jwt.SigningMethodES256, tokenClaims)

	token, err := tokenDetails.SignedString(privKey)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
