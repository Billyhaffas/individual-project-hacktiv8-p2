package http

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
		return c.JSON(http.StatusBadRequest, rentBook.PostrentBookRespon{
			Status:  http.StatusText(http.StatusBadGateway),
			Message: err.Error(),
		})
	}
	err = h.rentBookUseCase.PostRentBook(c.Request().Context(), email, rentBookRequest.BookName, rentBookRequest.PaymentMethod, rentBookRequest.Duration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rentBook.PostrentBookRespon{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, rentBook.PostrentBookRespon{
		Status:  http.StatusText(http.StatusOK),
		Message: "succesfully rent book",
	})
}
