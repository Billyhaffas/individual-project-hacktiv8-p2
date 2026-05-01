package user

import (
	"context"
	"fmt"
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/infrastructure/security"
	"individual-project-hacktiv8-p2/internal/model/user"

	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	authRepository domain.AuthRepository
}

func AuthUseCase(AuthRepo domain.AuthRepository) domain.AuthUseCase {
	return &authUseCase{authRepository: AuthRepo}
}

func (auth *authUseCase) Register(ctx context.Context, req user.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := user.User{
		Email:         req.Email,
		Name:          req.Name,
		Password:      string(hashedPassword),
		DepositAmount: req.DepositAmount,
	}
	err = auth.authRepository.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func (auth *authUseCase) Login(ctx context.Context, request user.LoginRequest) (string, error) {
	var data user.User
	data.Email = request.Email
	repo, err := auth.authRepository.Login(data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	err = security.ComparePassword(repo.Password, request.Password)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	token, err := security.GenerateToken(repo.UserId, repo.Email)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}
	go auth.authRepository.UpdateJwt(token, repo.Email)
	return token, nil

}

func (auth *authUseCase) GetMe(ctx context.Context, email string) (*user.GetMe, error) {
	repo, err := auth.authRepository.GetMe(email)
	if err != nil {
		return nil, err
	}
	respon := user.GetMe{
		Id:            repo.UserId,
		Email:         repo.Email,
		Name:          repo.Name,
		DepositAmount: repo.DepositAmount,
	}
	return &respon, nil
}

// func (auth *authUseCase) Login(ctx context.Context, req user.LoginRequest) (string, error) {
// 	err := auth.authRepository.Register(req)
// }
