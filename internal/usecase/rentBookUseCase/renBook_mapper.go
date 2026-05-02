package rentBookUseCase

import (
	"fmt"
	"individual-project-hacktiv8-p2/internal/model/book"
	"individual-project-hacktiv8-p2/internal/model/payment"
	"individual-project-hacktiv8-p2/internal/model/paymentRecordEntity"
	"individual-project-hacktiv8-p2/internal/model/rentBook"
	"individual-project-hacktiv8-p2/internal/model/user"
	"time"
)

func mapRentBookJoinToResponse(repoRentBook rentBook.GetRentBook) rentBook.GetRentBookRespon {

	responPaymentMethod := payment.GetRentBookPaymentMethod{
		Name: repoRentBook.PaymentMethod,
	}
	userRespon := user.UserRespon{
		Name:  repoRentBook.UserName,
		Email: repoRentBook.Email,
	}

	bookRespon := book.RentBookRespon{
		Name:       repoRentBook.BookName,
		Cartegory:  repoRentBook.Category,
		RentalCost: repoRentBook.RentalCost,
	}

	responPayemntRecord := paymentRecordEntity.GetRentBookPaymentRespon{
		PaymentRecordId: repoRentBook.PaymentRecordID,
		PaymentId:       repoRentBook.PaymentID,
		PaymentInfo:     &responPaymentMethod,
		Amount:          repoRentBook.Amount,
		UseFor:          repoRentBook.UseFor,
		CreatedAt:       repoRentBook.CreatedAtPayment,
	}
	startTime := time.Now()
	endTime := repoRentBook.DueDateRent
	diff := endTime.Sub(startTime)
	totalHours := diff.Hours()
	days := int(totalHours / 24)
	remainingHours := int(totalHours) % 24
	countDate := fmt.Sprintf(
		"%d hari %d jam",
		days,
		remainingHours,
	)

	responRentBook := rentBook.GetRentBookRespon{
		RentId:          repoRentBook.RentID,
		UserId:          repoRentBook.UserID,
		BookId:          repoRentBook.BookID,
		PaymentRecordId: repoRentBook.PaymentRecordID,
		User:            &userRespon,
		Book:            &bookRespon,
		PaymentRecord:   &responPayemntRecord,
		Duration:        repoRentBook.Duration,
		CreatedAt:       repoRentBook.CreatedAtRentBook,
		DueDateRent:     repoRentBook.DueDateRent,
		CountDate:       countDate,
	}
	return responRentBook
}
