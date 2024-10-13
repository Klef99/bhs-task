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
	status, err := uc.repo.CreateUser(ctx, crd)
	if err != nil {
		return false, fmt.Errorf("UserUseCase - Register - s.repo.CreateUser: %w", err)
	}
	return status, err
}

func (uc *UserUseCase) Login(ctx context.Context, crd entity.Credentials) (entity.User, error) {
	resp := entity.User{}
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
func (uc *UserUseCase) MakeDeposit(ctx context.Context, user entity.User, amount float64) (bool, error) {
	status, err := uc.repo.MakeDeposit(ctx, user, amount)
	if err != nil {
		return false, fmt.Errorf("UserUseCase - Deposit - uc.repo.Deposit: %w", err)
	}
	return status, err
}

// CheckDeposit -.
func (uc *UserUseCase) CheckDeposit(ctx context.Context, user entity.User) (float64, error) {
	balance, err := uc.repo.CheckDeposit(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("UserUseCase - CheckBalance - uc.repo.CheckDeposit: %w", err)
	}
	return balance, err
}
