package paymentRecord

import (
	"context"
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/model/paymentRecordEntity"
	"time"
)

type paymentRecordUseCase struct {
	paymentRecordRepository domain.PaymentRecordRepository
	userRepository          domain.AuthRepository
}

func NewPaymentRecordUseCase(PaymentRecordRepo domain.PaymentRecordRepository, UserRepo domain.AuthRepository) domain.PaymentRecordUseCase {
	return &paymentRecordUseCase{
		paymentRecordRepository: PaymentRecordRepo,
		userRepository:          UserRepo,
	}
}

func (paymentUC *paymentRecordUseCase) PostPaymentRecord(ctx context.Context, email string, paymentMethod string, amount float32, useFor string) error {
	user, err := paymentUC.userRepository.GetMe(email)
	if err != nil {
		return err
	}
	payment, err := paymentUC.paymentRecordRepository.GetPaymentMethod(paymentMethod)
	if err != nil {
		return err
	}
	paymentRequest := paymentRecordEntity.PaymentsRecords{
		UserId:    user.UserId,
		PaymentId: payment.PaymentId,
		Amount:    amount,
		UseFor:    useFor,
		CreatedAt: time.Now(),
	}
	_, err = paymentUC.paymentRecordRepository.PaymentActivity(&paymentRequest)
	if err != nil {
		return err
	}
	err = paymentUC.userRepository.TopUpBalance(amount, email)
	return nil
}
