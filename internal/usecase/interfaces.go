package usecase

import (
	"context"

	"github.com/Klef99/bhs-task/internal/entity"
)

type (
	User interface {
		Register(ctx context.Context, crd entity.Credentials) (bool, error)
		Login(ctx context.Context, crd entity.Credentials) (bool, error)
	}

	UserRepository interface {
		CreateUser(ctx context.Context, crd entity.Credentials) (bool, error)
		LoginUser(ctx context.Context, crd entity.Credentials) (bool, error)
	}
)
