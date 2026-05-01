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
