package db

import (
	"errors"
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/model/book"

	"gorm.io/gorm"
)

type bookDBconnection struct {
	db *gorm.DB
}

func NewBookDBconnection(db *gorm.DB) domain.BookRepository {
	return &bookDBconnection{db: db}

}

func (bookDB *bookDBconnection) GetbookByNameTx(tx *gorm.DB, name string) (*book.BookRent, error) {
	var bookRent book.BookRent
	err := tx.Table("books").Select("book_id", "stock_availability", "rental_cost").Where("name = ?", name).First(&bookRent).Error
	if err != nil {
		return nil, err
	}
	return &bookRent, nil
}

func (bookDB *bookDBconnection) UpdateStockForRentTx(tx *gorm.DB, bookId int) error {
	result := tx.
		Table("books").
		Where("book_id = ? AND stock_availability > 0", bookId).
		UpdateColumn(
			"stock_availability",
			gorm.Expr("stock_availability - ?", 1),
		)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("stock not available")
	}
	return nil
}
