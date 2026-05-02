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

func (rentDB *rentBookDBconnection) GetRentBook(userId int) ([]rentBook.GetRentBook, error) {
	var payments []rentBook.GetRentBook
	err := rentDB.db.
		Table("rents_books").
		Select(`
        rents_books.rent_id,
        rents_books.user_id,
        rents_books.book_id,
        rents_books.payment_record_id,
        books.name AS book_name,
        books.rental_cost,
        books.category,
        users.name AS user_name,
        users.email,
        payments_records.payment_id,
        payments.name AS payment_method,
        payments_records.amount,
        payments_records.use_for,
		payments_records.created_at,
		rents_books.duration,
		rents_books.created_at,
		rents_books.due_date_rent
    `).
		Joins("JOIN users ON users.user_id = rents_books.user_id").
		Joins("JOIN books ON books.book_id = rents_books.book_id").
		Joins("JOIN payments_records ON payments_records.payment_record_id = rents_books.payment_record_id").
		Joins("JOIN payments ON payments.payment_id = payments_records.payment_id").
		Where("users.user_id = ?", userId).
		Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}
