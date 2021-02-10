package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpazk/go-rest-api/middleware"
)

type Service interface {
	GetAccessToken(userID int) (string, error)
}

type authService struct {
}

func NewService() *authService {
	return &authService{}
}

func (s *authService) GetAccessToken(userID int) (string, error) {
	claims := &middleware.JwtCustomClaims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedKey, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return signedKey, err
	}

	return signedKey, nil
}
