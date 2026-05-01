package domain

import (
	"context"
	"individual-project-hacktiv8-p2/internal/model/user"

	"github.com/labstack/echo/v5"
)

type AuthRepository interface {
	Register(req user.User) error
	Login(data user.User) (*user.User, error)
	GetMe(email string) (*user.User, error)
	UpdateJwt(token, email string) error
	TopUpBalance(DepositAmount float32, email string) error
	PaymentBalance(DepositAmount float32, email string) error
	// GetPasswordByEmail(email string) (*string, error)
}
type AuthUseCase interface {
	Register(ctx context.Context, req user.User) error
	Login(ctx context.Context, request user.LoginRequest) (string, error)
	GetMe(ctx context.Context, email string) (*user.GetMe, error)
	// Login(ctx context.Context, req user.LoginRequest) (string, error)
}

type AuthHandler interface {
	Register(c *echo.Context) error
	Login(c *echo.Context) error
	GetMe(c *echo.Context) error
	// Login(c *echo.Context) error
}
