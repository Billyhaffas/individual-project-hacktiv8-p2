package rentBook

import (
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

type PostrentBookRespon struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GetrentBookRespon struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    *[]RentBookRespon `json:"data"`
}
