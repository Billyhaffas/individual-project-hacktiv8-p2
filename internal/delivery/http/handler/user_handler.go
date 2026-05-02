package handler

import (
	"errors"
	"fmt"
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/helper"
	"individual-project-hacktiv8-p2/internal/infrastructure/security"
	"individual-project-hacktiv8-p2/internal/model/user"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v5"
)

type authHandler struct {
	authUseCase domain.AuthUseCase
}

func AuthHandler(AuthUseCase domain.AuthUseCase) domain.AuthHandler {
	return &authHandler{authUseCase: AuthUseCase}
}

func (h *authHandler) Register(c *echo.Context) error {
	var request user.RegisterRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, user.RegisterRespon{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	}
	err = helper.ValidatePassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, user.RegisterRespon{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	}
	userUseCase := user.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	}

	err = h.authUseCase.Register(c.Request().Context(), userUseCase)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)" {
			var ErrEmailExists = errors.New("email already registered")
			return c.JSON(http.StatusInternalServerError, user.RegisterRespon{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: ErrEmailExists.Error(),
			})

		}
		return c.JSON(http.StatusInternalServerError, user.RegisterRespon{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, user.RegisterRespon{
		Status:  http.StatusText(http.StatusCreated),
		Message: "User has been successfully created",
	})
}

func (h *authHandler) Login(c *echo.Context) error {
	var request user.LoginRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, user.LoginRespon{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
			Token:   "doesn't generated token",
		})
	}
	token, err := h.authUseCase.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, user.LoginRespon{
			Status:  http.StatusText(http.StatusUnauthorized),
			Message: "invalid credentials",
			Token:   "doesn't generated token",
		})
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, user.LoginRespon{
		Status:  http.StatusText(http.StatusOK),
		Message: "Successfully Authentication",
		Token:   token,
	})
}

func (h *authHandler) GetMe(c *echo.Context) error {
	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(*security.JWTCustomClaims)

	email := claims.Email

	getMeRespon, err := h.authUseCase.GetMe(c.Request().Context(), email)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, user.LoginRespon{
			Status:  http.StatusText(http.StatusUnauthorized),
			Message: "invalid credentials",
			Token:   "Token doesn't claims",
		})
	}
	return c.JSON(http.StatusOK, user.GetMeRespon{
		Status:  http.StatusText(http.StatusOK),
		Message: "Successfully Authorization",
		Data:    getMeRespon,
	})

}
