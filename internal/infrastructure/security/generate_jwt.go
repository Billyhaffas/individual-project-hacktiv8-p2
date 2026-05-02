package security

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JWTSecret() string {
	key := os.Getenv("JWT_SECRET")
	return key
}

type JWTCustomClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int, email string) (string, error) {
	claims := JWTCustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "billyhaffas_p2_lc02",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret()))
}
