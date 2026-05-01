package http

import (
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/helper"
	"individual-project-hacktiv8-p2/internal/model/paymentRecordEntity"
	"net/http"

	"github.com/labstack/echo/v5"
)

type paymentHandler struct {
	paymentUseCase domain.PaymentRecordUseCase
}

func PaymentHandler(PaymentUseCase domain.PaymentRecordUseCase) domain.PaymentRecordHandler {
	return &paymentHandler{paymentUseCase: PaymentUseCase}
}

func (handlerPyr *paymentHandler) PostPaymentRecord(c *echo.Context) error {
	claims := helper.GetUserFromToken(c)
	email := claims.Email

	var paymentRecordRequest paymentRecordEntity.PaymentRequest
	err := c.Bind(&paymentRecordRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, paymentRecordEntity.PostPaymentRecord{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}
	err = handlerPyr.paymentUseCase.PostPaymentRecord(c.Request().Context(), email, paymentRecordRequest.PaymentMethod, paymentRecordRequest.Amount, paymentRecordRequest.UseFor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, paymentRecordEntity.PostPaymentRecord{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, paymentRecordEntity.PostPaymentRecord{
		Status:  http.StatusText(http.StatusOK),
		Message: "success created top up saldo",
	})

}
