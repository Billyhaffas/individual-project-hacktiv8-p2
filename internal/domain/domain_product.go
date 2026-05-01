package domain

import (
	"individual-project-hacktiv8-p2/internal/model/book"

	"gorm.io/gorm"
)

type bookDBconnection struct {
	db *gorm.DB
}

type BookRepository interface {
	GetbookByNameTx(tx *gorm.DB, name string) (*book.BookRent, error)
	UpdateStockForRentTx(tx *gorm.DB, bookId int) error
}
type bookUseCase interface {
}

type bookHandler interface {
}
