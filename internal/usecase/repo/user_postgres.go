package repo

import (
	"context"
	"fmt"

	"github.com/Klef99/bhs-task/internal/entity"
	"github.com/Klef99/bhs-task/internal/usecase"
	"github.com/Klef99/bhs-task/pkg/postgres"
	sq "github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository -.
type UserRepository struct {
	*postgres.Postgres
}

var _ usecase.UserRepository = (*UserRepository)(nil)

// New -.
func New(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

// CreateUser -.
func (r *UserRepository) CreateUser(ctx context.Context, crd entity.Credentials) (status bool, err error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(crd.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, fmt.Errorf("UserRepository - CreateUser - bcrypt.GenerateFromPassword: %w", err)
	}
	sql, args, err := r.Builder.
		Insert("users").
		Columns("username, password_hash").
		Values(crd.Username, hashedBytes).
		ToSql()

	if err != nil {
		return false, fmt.Errorf("UserRepository - CreateUser - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return false, fmt.Errorf("UserRepository - CreateUser - r.Pool.Exec: %w", err)
	}

	return true, nil
}

// LoginUser -.
func (r *UserRepository) LoginUser(ctx context.Context, crd entity.Credentials) (bool, error) {
	sql, args, err := r.Builder.Select("password_hash").From("users").Where(sq.Eq{"username": crd.Username}).ToSql()
	if err != nil {
		return false, fmt.Errorf("UserRepository - LoginUser - r.Builder: %w", err)
	}
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return false, fmt.Errorf("UserRepository - LoginUser - r.Pool.Query: %w", err)
	}
	var passwordHash string
	for rows.Next() {
		rows.Scan(&passwordHash)
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(crd.Password))
	if err != nil {
		return false, fmt.Errorf("UserRepository - LoginUser - bcrypt.CompareHashAndPassword: %w", err)
	}
	return true, nil
}
