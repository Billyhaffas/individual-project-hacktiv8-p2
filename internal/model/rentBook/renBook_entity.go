package rentBook

import (
	"time"
)

type OriginRentBook struct {
	rentBookId int
	UserId     int
	TotalPrice float32
	CreatedAt  time.Time
}
type rentBook struct {
	UserId     int
	TotalPrice float32
	CreatedAt  time.Time
}

type rentBookJoin struct {
	rentBookID int
	UserID     int
	UserName   string
	UserEmail  string
	TotalPrice float32
	CreatedAt  time.Time
}

type RentBookJoinResult struct {
	UserID          int
	BookID          int
	PaymentRecordID int
}

type RentsBooks struct {
	UserID          int
	BookID          int
	PaymentRecordID int
	Duration        int
	CreatedAt       time.Time
	DueDateRent     time.Time
}

type GetRentBook struct {
	RentID            int
	UserID            int
	BookID            int
	PaymentRecordID   int
	BookName          string
	RentalCost        float32
	Category          string
	UserName          string
	Email             string
	PaymentID         int
	PaymentMethod     string
	Amount            float32
	UseFor            string
	CreatedAtPayment  time.Time
	Duration          int
	CreatedAtRentBook time.Time
	DueDateRent       time.Time
	CountDate         string
}
