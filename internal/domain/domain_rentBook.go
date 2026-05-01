package domain

import (
	"context"
	"individual-project-hacktiv8-p2/internal/model/rentBook"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type rentBookDBconnection struct {
	db *gorm.DB
}

type RentBookRepository interface {
	PostRentBookDependencies(ctx context.Context, request rentBook.RentsBooks) error
}
type RentBookUseCase interface {
	PostRentBook(ctx context.Context, email string, bookname string, paymentMethod string, duration int) error
}
type RentBookHandler interface {
	PostRentBook(c *echo.Context) error
}
