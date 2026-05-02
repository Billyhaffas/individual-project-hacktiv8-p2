package googlebooks

import "time"

type PaymentRecord struct {
	PaymentRecordId int
	UserId          int
	PaymentId       int
	Amount          float32
	UseFor          string
	CreatedAt       time.Time
}

type PaymentRecordHelper struct {
	PaymentRecordId int
	UserId          int
	UserName        string
	UserSaldo       float32
	PaymentId       int
	PaymentName     string
	BankCover       string
	Amount          float32
	UseFor          string
	CreatedAt       time.Time
}

type PaymentsRecords struct {
	UserId    int
	PaymentId int
	Amount    float32
	UseFor    string
	CreatedAt time.Time
}
