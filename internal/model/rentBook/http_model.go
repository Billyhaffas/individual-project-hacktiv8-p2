package rentBook

import (
	"individual-project-hacktiv8-p2/internal/model/book"
	"individual-project-hacktiv8-p2/internal/model/paymentRecordEntity"
	"individual-project-hacktiv8-p2/internal/model/user"
	"time"
)

type RentBookRespon struct {
	RentBookId int              `json:"rentBook_id"`
	UserId     int              `json:"user_id"`
	DetailUser *user.UserRespon `json:"detail_user"`
	TotalPrice float32          `json:"total_price"`
	CreatedAt  time.Time        `json:"created_at"`
}

type RentBookRequest struct {
	BookName      string `json:"book_name"`
	PaymentMethod string `json:"payment_method"`
	Duration      int    `json:"duration"`
}

type RentBookResponHandler struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GetrentBookRespon struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    *[]GetRentBookRespon `json:"data"`
}

type GetRentBookRespon struct {
	RentId          int                                           `json:"rent_id"`
	UserId          int                                           `json:"user_id"`
	BookId          int                                           `json:"book_id"`
	User            *user.UserRespon                              `json:"user_detail"`
	Book            *book.RentBookRespon                          `json:"book_detail"`
	PaymentRecordId int                                           `json:"payment_record_id"`
	PaymentRecord   *paymentRecordEntity.GetRentBookPaymentRespon `json:"payment_record_detail"`
	Duration        int                                           `json:"duration"`
	CreatedAt       time.Time                                     `json:"created_at"`
	DueDateRent     time.Time                                     `json:"due_date_rent"`
	CountDate       string                                        `json:"count_due_date"`
}
