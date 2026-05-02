package helper

import (
	"individual-project-hacktiv8-p2/internal/infrastructure/security"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func GetUserFromToken(c *echo.Context) *security.JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*security.JWTCustomClaims)
}
