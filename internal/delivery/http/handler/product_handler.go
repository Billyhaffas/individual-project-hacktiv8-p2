package http

// import (
// 	"individual-project-hacktiv8-p2/internal/model/book"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v5"
// )

// type bookHandler struct {
// 	bookUseCase domain.bookUseCase
// }

// func NewbookHandler(bookUseCase domain.bookUseCase) domain.bookHandler {
// 	return &bookHandler{bookUseCase: bookUseCase}
// }

// func (h *bookHandler) GetAllbook(c *echo.Context) error {
// 	respon, err := h.bookUseCase.GetAllbook(c.Request().Context())
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, book.Getbook{
// 			Status:  http.StatusText(http.StatusInternalServerError),
// 			Message: err.Error(),
// 			Data:    nil,
// 		})
// 	}
// 	return c.JSON(http.StatusOK, book.Getbook{
// 		Status:  http.StatusText(http.StatusOK),
// 		Message: "success get cart",
// 		Data:    respon,
// 	})
// }

// func (h *bookHandler) GetbookById(c *echo.Context) error {
// 	cartIDStr := c.Param("book_id")
// 	cartID, err := strconv.Atoi(cartIDStr)
// 	respon, err := h.bookUseCase.GetbookById(c.Request().Context(), cartID)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, book.GetbookById{
// 			Status:  http.StatusText(http.StatusInternalServerError),
// 			Message: err.Error(),
// 			Data:    nil,
// 		})
// 	}
// 	return c.JSON(http.StatusOK, book.GetbookById{
// 		Status:  http.StatusText(http.StatusOK),
// 		Message: "success get book",
// 		Data:    respon,
// 	})
// }
