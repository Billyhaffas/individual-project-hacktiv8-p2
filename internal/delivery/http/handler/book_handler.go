package handler

import (
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/model/book"
	"net/http"

	"github.com/labstack/echo/v5"
)

type bookHandler struct {
	bookUseCase domain.BookUseCase
}

func NewbookHandler(BookUseCase domain.BookUseCase) domain.BookHandler {
	return &bookHandler{bookUseCase: BookUseCase}
}

func (h *bookHandler) GetbookByName(c *echo.Context) error {
	bookName := c.Param("book_name")
	respon, err := h.bookUseCase.GetbookByName(c.Request().Context(), bookName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, book.Getbook{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, book.Getbook{
		Status:  http.StatusText(http.StatusOK),
		Message: "success get cart",
		Data:    respon,
	})
}
