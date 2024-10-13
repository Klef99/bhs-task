package repo

import (
	"context"
	"fmt"

	"github.com/Klef99/bhs-task/internal/entity"
	"github.com/Klef99/bhs-task/internal/usecase"
	"github.com/Klef99/bhs-task/pkg/hasher"
	"github.com/Klef99/bhs-task/pkg/postgres"
	sq "github.com/Masterminds/squirrel"
)

// UserRepository -.
type UserRepository struct {
	*postgres.Postgres
	Hasher hasher.Interface
}

var _ usecase.UserRepository = (*UserRepository)(nil)

// New -.
func NewUserRepository(pg *postgres.Postgres, hs hasher.Interface) *UserRepository {
	return &UserRepository{pg, hs}
}

// CreateUser -.
func (r *UserRepository) CreateUser(ctx context.Context, crd entity.Credentials) (bool, error) {
	hashedBytes, err := r.Hasher.HashPassword(crd.Password)
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
func (r *UserRepository) LoginUser(ctx context.Context, crd entity.Credentials) (int64, error) {
	sql, args, err := r.Builder.Select("id", "password_hash").From("users").Where(sq.Eq{"username": crd.Username}).ToSql()
	if err != nil {
		return -1, fmt.Errorf("UserRepository - LoginUser - r.Builder: %w", err)
	}
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return -1, fmt.Errorf("UserRepository - LoginUser - r.Pool.Query: %w", err)
	}
	var passwordHash string
	var id int64
	for rows.Next() {
		rows.Scan(&id, &passwordHash)
	}
	err = r.Hasher.CompareHashAndPassword(passwordHash, crd.Password)
	if err != nil {
		return -1, fmt.Errorf("UserRepository - LoginUser - bcrypt.CompareHashAndPassword: %w", err)
	}
	return id, nil
}

// Deposit -.
func (r *UserRepository) MakeDeposit(ctx context.Context, user entity.User, amount float64) (bool, error) {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("UserRepository - Deposit - r.Pool.Begin: %w", err)
	}
	sql, args, err := r.Builder.
		Update("users").
		Set("balance", sq.Expr("balance +?", amount)).
		Where(sq.Eq{"id": user.Id}).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("UserRepository - Deposit - r.Builder: %w", err)
	}
	res, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return false, fmt.Errorf("UserRepository - Deposit - r.Pool.Exec: %w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return false, fmt.Errorf("UserRepository - Deposit - tx.Commit: %w", err)
	}
	rowsAffected := res.RowsAffected()
	return rowsAffected > 0, nil
}

// CheckDeposit -.
func (r *UserRepository) CheckDeposit(ctx context.Context, user entity.User) (float64, error) {
	sql, args, err := r.Builder.Select("balance").From("users").Where(sq.Eq{"id": user.Id}).ToSql()
	if err != nil {
		return 0, fmt.Errorf("UserRepository - CheckDeposit - r.Builder: %w", err)
	}
	row := r.Pool.QueryRow(ctx, sql, args...)
	var balance float64
	err = row.Scan(&balance)
	if err != nil {
		return 0, fmt.Errorf("UserRepository - CheckDeposit - row.Scan: %w", err)
	}
	return balance, nil
}
