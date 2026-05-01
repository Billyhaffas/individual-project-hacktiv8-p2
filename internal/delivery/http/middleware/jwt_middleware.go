package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"

	"individual-project-hacktiv8-p2/internal/infrastructure/security"
)

func jwtClaims(c *echo.Context) jwt.Claims {
	return new(security.JWTCustomClaims)
}
func jwtErrorHandler(c *echo.Context, err error) error {
	return (c).JSON(401, map[string]string{
		"message": "unauthorized",
	})
}
func AuthMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(security.JWTSecret()),
		SigningMethod: "HS256",
		ContextKey:    "user",

		NewClaimsFunc: jwtClaims,
		ErrorHandler:  jwtErrorHandler,
	})
}
