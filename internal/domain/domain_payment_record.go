package domain

import (
	"context"
	"individual-project-hacktiv8-p2/internal/model/payment"
	"individual-project-hacktiv8-p2/internal/model/paymentRecordEntity"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type PaymentRecordDBconnection struct {
	db *gorm.DB
}

type PaymentRecordRepository interface {
	PaymentActivity(PaymentsRecords *paymentRecordEntity.PaymentsRecords) (*int, error)
	GetPaymentMethod(name string) (*payment.Payment, error)

	// GetbookById(id int) (book.book, error)
}
type PaymentRecordUseCase interface {
	PostPaymentRecord(ctx context.Context, email string, paymentMethod string, amount float32, useFor string) error
	// GetbookById(ctx context.Context, bookID int) (*book.bookRespon, error)
	// GetbookById(ctx context.Context, id int) (*book.bookRespon, error)
}

type PaymentRecordHandler interface {
	PostPaymentRecord(c *echo.Context) error
	// GetbookById(c *echo.Context) error
	// GetbookById(c *echo.Context) error
}
