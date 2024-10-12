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
func New(r UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc *UserUseCase) Register(ctx context.Context, crd entity.Credentials) (bool, error) {
	status, err := uc.repo.CreateUser(ctx, crd)
	if err != nil {
		return false, fmt.Errorf("UserUseCase - Register - s.repo.CreateUser: %w", err)
	}
	return status, err
}

func (uc *UserUseCase) Login(ctx context.Context, crd entity.Credentials) (bool, error) {
	status, err := uc.repo.LoginUser(ctx, crd)
	if err != nil {
		return false, fmt.Errorf("UserUseCase - Login - s.repo.LoginUser: %w", err)
	}
	return status, err
}

// Logout(ctx context.Context) (status bool, err error)
