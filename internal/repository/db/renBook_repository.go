package db

import (
	"context"
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/model/rentBook"

	"gorm.io/gorm"
)

type rentBookDBconnection struct {
	db *gorm.DB
}

func NewRentBookDBconnection(db *gorm.DB) domain.RentBookRepository {
	return &rentBookDBconnection{db: db}
}

func (rentDB *rentBookDBconnection) PostRentBookDependencies(ctx context.Context, request rentBook.RentsBooks) error {
	err := rentDB.db.Table("rents_books").Create(&request)
	if err != nil {
		return err.Error
	}
	return nil
}
