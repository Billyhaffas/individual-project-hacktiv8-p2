package db

import (
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/model/payment"
	"individual-project-hacktiv8-p2/internal/model/paymentRecordEntity"

	"gorm.io/gorm"
)

type paymentRecordDBconnection struct {
	db *gorm.DB
}

func NewPaymentRecordDBConnection(db *gorm.DB) domain.PaymentRecordRepository {
	return &paymentRecordDBconnection{db: db}
}

func (recordPaymentDB *paymentRecordDBconnection) PaymentActivity(PaymentsRecords *paymentRecordEntity.PaymentsRecords) (*int, error) {
	err := recordPaymentDB.db.Create(&PaymentsRecords).Error
	if err != nil {
		return nil, err
	}
	var PaymentResponId int
	err = recordPaymentDB.db.Table("payments_records").Where("payment_id", PaymentsRecords.PaymentId).Pluck("payment_record_id", &PaymentResponId).Error
	if err != nil {
		return nil, err
	}
	return &PaymentResponId, nil
}
func (recordPaymentDB *paymentRecordDBconnection) GetPaymentMethod(name string) (*payment.Payment, error) {
	var payment payment.Payment
	err := recordPaymentDB.db.
		Table("payments").
		Select("payment_id").
		Where("name = ?", name).
		First(&payment).Error
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
