package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func JwtMiddleWare() echo.MiddlewareFunc {
	key := os.Getenv("JWT_SECRET_KEY")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(key),
	})
}
