package handler

import (
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/helper"
	"individual-project-hacktiv8-p2/internal/model/rentBook"
	"net/http"

	"github.com/labstack/echo/v5"
)

type rentBookHandler struct {
	rentBookUseCase domain.RentBookUseCase
}

func NewrentBookHandler(RentBookUseCase domain.RentBookUseCase) domain.RentBookHandler {
	return &rentBookHandler{rentBookUseCase: RentBookUseCase}
}

func (h *rentBookHandler) PostRentBook(c *echo.Context) error {
	//clains email
	claims := helper.GetUserFromToken(c)
	email := claims.Email
	var rentBookRequest rentBook.RentBookRequest
	err := c.Bind(&rentBookRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rentBook.RentBookResponHandler{
			Status:  http.StatusText(http.StatusBadGateway),
			Message: err.Error(),
		})
	}
	err = h.rentBookUseCase.PostRentBook(c.Request().Context(), email, rentBookRequest.BookName, rentBookRequest.PaymentMethod, rentBookRequest.Duration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rentBook.RentBookResponHandler{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, rentBook.RentBookResponHandler{
		Status:  http.StatusText(http.StatusOK),
		Message: "succesfully rent book",
	})
}
func (h *rentBookHandler) GetRentBook(c *echo.Context) error {
	//clains email
	claims := helper.GetUserFromToken(c)
	email := claims.Email
	respon, err := h.rentBookUseCase.GetRentBookByEmail(c.Request().Context(), email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, rentBook.RentBookResponHandler{
			Status:  http.StatusText(http.StatusUnauthorized),
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, rentBook.GetrentBookRespon{
		Status:  http.StatusText(http.StatusOK),
		Message: "succesfully get rent books data",
		Data:    &respon,
	})
}
