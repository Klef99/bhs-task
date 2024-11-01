package usecase

import (
	"context"
	"fmt"

	"github.com/Klef99/bhs-task/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	repo UserRepository
}

var _ User = (*UserUseCase)(nil)

// New -.
func NewUserUseCase(r UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc *UserUseCase) Register(ctx context.Context, crd entity.Credentials) (bool, error) {
	if crd.Password == "" || crd.Username == "" {
		return false, fmt.Errorf("UserUseCase - Register - crd.Validate: invalid credentials")
	}
	status, err := uc.repo.CreateUser(ctx, crd)
	if err != nil {
		return false, fmt.Errorf("UserUseCase - Register - s.repo.CreateUser: %w", err)
	}
	return status, err
}

func (uc *UserUseCase) Login(ctx context.Context, crd entity.Credentials) (entity.User, error) {
	resp := entity.User{}
	if crd.Password == "" || crd.Username == "" {
		return resp, fmt.Errorf("UserUseCase - Login - crd.Validate: invalid credentials")
	}
	id, err := uc.repo.LoginUser(ctx, crd)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserUseCase - Login - s.repo.LoginUser: %w", err)
	}
	if id != -1 {
		resp.Id = id
		resp.Username = crd.Username
	}
	return resp, err
}

// Deposit -.
func (uc *UserUseCase) MakeDeposit(ctx context.Context, user entity.User, amount float64) (float64, error) {
	if user.Id < 1 {
		return -1, fmt.Errorf("UserUseCase - MakeDeposit - invalid input: user id must be provided")
	}
	if amount <= 0 {
		return -1, fmt.Errorf("UserUseCase - MakeDeposit - invalid input: amount must be greater than zero")
	}
	balance, err := uc.repo.MakeDeposit(ctx, user, amount)
	if err != nil {
		return -1, fmt.Errorf("UserUseCase - Deposit - uc.repo.Deposit: %w", err)
	}
	return balance, err
}

// CheckDeposit -.
func (uc *UserUseCase) CheckDeposit(ctx context.Context, user entity.User) (float64, error) {
	if user.Id < 1 {
		return 0, fmt.Errorf("UserUseCase - CheckBalance - invalid input: user id must be provided")
	}
	balance, err := uc.repo.CheckDeposit(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("UserUseCase - CheckBalance - uc.repo.CheckDeposit: %w", err)
	}
	return balance, err
}
