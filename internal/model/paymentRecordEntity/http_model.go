package paymentRecordEntity

import (
	"individual-project-hacktiv8-p2/internal/model/payment"
	"individual-project-hacktiv8-p2/internal/model/user"
	"time"
)

type PaymentRespon struct {
	PaymentRecordId int                  `json:"payment_record_id"`
	UserId          int                  `json:"user_id"`
	PaymentId       int                  `json:"payment_id"`
	UserInfo        *user.GetMe          `json:"user_info"`
	PaymentInfo     *payment.PaymentInfo `json:"payment_info"`
	Amount          float32              `json:"amount"`
	UseFor          string               `json:"use_for"`
	CreatedAt       time.Time            `json:"created_at"`
}
type AllPayementRecord struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    *[]PaymentRespon `json:"data"`
}

type GetPaymentById struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    *PaymentRespon `json:"data"`
}

type PaymentRequest struct {
	PaymentMethod string  `json:"payment_name"`
	Amount        float32 `json:"amount"`
	UseFor        string  `json:"use_for"`
}
type PostPaymentRecord struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GetRentBookPaymentRespon struct {
	PaymentRecordId int                               `json:"payment_record_id"`
	PaymentId       int                               `json:"payment_id"`
	PaymentInfo     *payment.GetRentBookPaymentMethod `json:"payment_info"`
	Amount          float32                           `json:"amount"`
	UseFor          string                            `json:"use_for"`
	CreatedAt       time.Time                         `json:"created_at"`
}
