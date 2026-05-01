package db

import (
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/model/user"

	"gorm.io/gorm"
)

type authDBconnection struct {
	db *gorm.DB
}

func NewAuthDBconnection(db *gorm.DB) domain.AuthRepository {
	return &authDBconnection{db: db}
}

func (auth *authDBconnection) Register(data user.User) error {
	err := auth.db.Select("name", "email", "password").Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (auth *authDBconnection) Login(data user.User) (*user.User, error) {
	var res user.User
	err := auth.db.
		Select("email", "password").
		Where("email = ?", data.Email).
		First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (auth *authDBconnection) GetMe(email string) (*user.User, error) {
	var res user.User
	err := auth.db.
		Select("user_id", "email", "name", "deposit_amount", "jwt_token").
		Where("email = ?", email).
		First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (auth *authDBconnection) UpdateJwt(token, email string) error {
	err := auth.db.Table("users").Where("email =?", email).UpdateColumn("jwt_token", token).Error
	if err != nil {
		return err
	}
	return nil
}

func (auth *authDBconnection) TopUpBalance(DepositAmount float32, email string) error {
	err := auth.db.Table("users").Where("email =?", email).Update("deposit_amount", gorm.Expr("deposit_amount + ?", DepositAmount)).Error
	if err != nil {
		return err
	}
	return nil
}

func (auth *authDBconnection) PaymentBalance(DepositAmount float32, email string) error {
	err := auth.db.
		Table("users").
		Where("email = ?", email).
		Update("deposit_amount", gorm.Expr("deposit_amount - ?", DepositAmount)).Error
	if err != nil {
		return err
	}
	return nil
}
