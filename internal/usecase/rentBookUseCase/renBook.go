package rentBookUseCase

import (
	"context"
	"errors"
	"fmt"
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/model/paymentRecordEntity"
	"individual-project-hacktiv8-p2/internal/model/rentBook"
	"strings"
	"time"

	"gorm.io/gorm"
)

type rentBookUseCase struct {
	db                      *gorm.DB
	userRepository          domain.AuthRepository
	rentBookRepository      domain.RentBookRepository
	paymentRecordRepository domain.PaymentRecordRepository
	bookRepository          domain.BookRepository
}

func NewRentBookUseCase(transaction *gorm.DB, userRepo domain.AuthRepository, rentBookRepo domain.RentBookRepository, paymentRecordRepo domain.PaymentRecordRepository, bookRepo domain.BookRepository) domain.RentBookUseCase {
	return &rentBookUseCase{
		db:                      transaction,
		userRepository:          userRepo,
		rentBookRepository:      rentBookRepo,
		paymentRecordRepository: paymentRecordRepo,
		bookRepository:          bookRepo,
	}
}

func (rentBookUC *rentBookUseCase) PostRentBook(ctx context.Context, email string, bookname string, paymentMethod string, duration int) error {
	var PaymentsRecords paymentRecordEntity.PaymentsRecords
	var bookId int
	err := rentBookUC.db.Transaction(func(tx *gorm.DB) error {
		userReq, err := rentBookUC.userRepository.GetMe(email)
		fmt.Println(userReq)
		if err != nil {
			return err
		}
		bookRent, err := rentBookUC.bookRepository.GetbookByNameTx(tx, bookname)
		bookId = bookRent.BookId
		fmt.Println(bookRent)
		if err != nil {
			return err
		}
		err = rentBookUC.bookRepository.UpdateStockForRentTx(tx, bookRent.BookId)
		if err != nil {
			return err
		}
		getPayment, err := rentBookUC.paymentRecordRepository.GetPaymentMethod(paymentMethod)
		fmt.Println(getPayment)
		if err != nil {
			return err
		}
		PaymentsRecords = paymentRecordEntity.PaymentsRecords{
			UserId:    userReq.UserId,
			PaymentId: getPayment.PaymentId,
			Amount:    bookRent.RentalCost * float32(duration),
			UseFor:    "payment for rent book",
			CreatedAt: time.Now(),
		}
		err = rentBookUC.userRepository.PaymentBalance(PaymentsRecords.Amount, email)
		if err != nil {
			if strings.Contains(err.Error(), "deposit_amount_check") {
				return errors.New("you don't have enough amount")
			}
			return err
		}
		paymentRecordID, err := rentBookUC.paymentRecordRepository.PaymentActivity(&PaymentsRecords)
		if err != nil {
			return err
		}

		requestRentBook := rentBook.RentsBooks{
			UserID:          PaymentsRecords.UserId,
			BookID:          bookId,
			PaymentRecordID: *paymentRecordID,
			Duration:        duration,
			CreatedAt:       time.Now(),
			DueDateRent:     time.Now().Add(time.Duration(duration) * 24 * time.Hour),
		}
		err = rentBookUC.rentBookRepository.PostRentBookDependencies(ctx, requestRentBook)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (rentBookUC *rentBookUseCase) GetRentBookByEmail(ctx context.Context, email string) ([]rentBook.GetRentBookRespon, error) {

	userReq, err := rentBookUC.userRepository.GetMe(email)
	if err != nil {
		return nil, err
	}

	repoRentBooks, err := rentBookUC.rentBookRepository.GetRentBook(userReq.UserId)
	if err != nil {
		return nil, err
	}

	responses := make([]rentBook.GetRentBookRespon, 0, len(repoRentBooks))

	for _, rentBookRepo := range repoRentBooks {
		responses = append(
			responses,
			mapRentBookJoinToResponse(rentBookRepo),
		)
	}

	return responses, nil
}
