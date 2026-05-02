package domain

import (
	"context"
	"individual-project-hacktiv8-p2/internal/model/book"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type BookDBconnection struct {
	db *gorm.DB
}

type BookRepository interface {
	GetbookByNameTx(tx *gorm.DB, name string) (*book.BookRent, error)
	UpdateStockForRentTx(tx *gorm.DB, bookId int) error
	GetbookByName(name string) (*book.GetBook, string, error)
}
type BookUseCase interface {
	GetbookByName(ctx context.Context, bookName string) (*book.GetBookRespon, error)
}

type BookHandler interface {
	GetbookByName(c *echo.Context) error
}
