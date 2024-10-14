package repo

import (
	"context"
	"fmt"

	"github.com/Klef99/bhs-task/internal/entity"
	"github.com/Klef99/bhs-task/internal/usecase"
	"github.com/Klef99/bhs-task/pkg/postgres"
	sq "github.com/Masterminds/squirrel"
)

// UserRepository -.
type AssetRepository struct {
	*postgres.Postgres
}

var _ usecase.AssetRepository = (*AssetRepository)(nil)

// New -.
func NewAssetRepository(pg *postgres.Postgres) *AssetRepository {
	return &AssetRepository{pg}
}

// Store -.
func (r *AssetRepository) Store(ctx context.Context, ast entity.Asset) (bool, error) {
	sql, args, err := r.Builder.
		Insert("assets").
		Columns("name", "description", "price", "owner_id").
		Values(ast.Name, ast.Description, ast.Price, ast.Owner_id).
		ToSql()

	if err != nil {
		return false, fmt.Errorf("AssetRepository - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return false, fmt.Errorf("AssetRepository - Store - r.Pool.Exec: %w", err)
	}

	return true, nil
}

// Erase -.
func (r *AssetRepository) Erase(ctx context.Context, user entity.User, id int64) (bool, error) {
	sql, args, err := r.Builder.
		Delete("assets").
		Where(sq.Eq{"id": id, "owner_id": user.Id}).
		ToSql()

	if err != nil {
		return false, fmt.Errorf("AssetRepository - Erase - r.Builder: %w", err)
	}

	res, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return false, fmt.Errorf("AssetRepository - Erase - r.Pool.Exec: %w", err)
	}

	rowsAffected := res.RowsAffected()
	return rowsAffected > 0, nil
}

// List -.
func (r *AssetRepository) UserAssetsList(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	sql, args, err := r.Builder.
		Select("*").
		From("assets").
		Where(sq.Eq{"owner_id": user.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AssetRepository - List - r.Builder: %w", err)
	}
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AssetRepository - List - r.Pool.Query: %w", err)
	}
	defer rows.Close()
	assets := make([]entity.Asset, 0)
	for rows.Next() {
		var ast entity.Asset
		err := rows.Scan(&ast.Id, &ast.Name, &ast.Description, &ast.Price, &ast.Owner_id)
		if err != nil {
			return nil, fmt.Errorf("AssetRepository - List - rows.Scan: %w", err)
		}
		assets = append(assets, ast)
	}
	return assets, nil
}

func (r *AssetRepository) GetOtherUsersAssets(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	sql, args, err := r.Builder.Select("id, name, description, price").
		From("assets").
		Where(sq.NotEq{"owner_id": user.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AssetRepository - GetOtherUserAssets - r.Builder: %w", err)
	}
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AssetRepository - GetOtherUserAssets - r.Pool.Query: %w", err)
	}
	defer rows.Close()
	assets := make([]entity.Asset, 0)
	for rows.Next() {
		ast := entity.Asset{Owner_id: user.Id}
		err := rows.Scan(&ast.Id, &ast.Name, &ast.Description, &ast.Price)
		if err != nil {
			return nil, fmt.Errorf("AssetRepository - GetOtherUserAssets - rows.Scan: %w", err)
		}
		assets = append(assets, ast)
	}
	return assets, nil
}

func (r *AssetRepository) BuyAsset(ctx context.Context, user entity.User, id int64) (bool, error) {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("AssetRepository - BuyAsset - r.Pool.Begin: %w", err)
	}
	sql, args, err := r.Builder.Select("price, owner_id").From("assets").Where(sq.Eq{"id": id}).Suffix("FOR UPDATE").ToSql()
	if err != nil {
		tx.Rollback(ctx)
		return false, fmt.Errorf("AssetRepository - BuyAsset - r.Builder.Select('price'): %w", err)
	}
	row := tx.QueryRow(ctx, sql, args...)
	var price float64
	var owner_id int64
	err = row.Scan(&price, &owner_id)
	if err != nil {
		tx.Rollback(ctx)
		return false, fmt.Errorf("AssetRepository - BuyAsset - row.Scan: %w", err)
	}
	if owner_id == user.Id {
		return false, fmt.Errorf("AssetRepository - BuyAsset - user can't buy their own asset")
	}
	sql, args, err = r.Builder.
		Update("users").
		Set("balance", sq.Expr("balance - ?", price)).
		Where(sq.Eq{"id": user.Id}).
		ToSql()
	if err != nil {
		tx.Rollback(ctx)
		return false, fmt.Errorf("AssetRepository - BuyAsset - r.Builder.Update('users'): %w", err)
	}
	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		tx.Rollback(ctx)
		return false, fmt.Errorf("AssetRepository - BuyAsset - tx.Exec: %w", err)
	}

	sql, args, err = r.Builder.
		Insert("access_assets").
		Columns("asset_id", "user_id").
		Values(id, user.Id).
		Suffix("on conflict (asset_id, user_id) do nothing").
		ToSql()
	if err != nil {
		tx.Rollback(ctx)
		return false, fmt.Errorf("AssetRepository - BuyAsset - r.Builder.Insert('acces_asset'): %w", err)
	}
	res, err := tx.Exec(ctx, sql, args...)
	if res.RowsAffected() == 0 || err != nil {
		tx.Rollback(ctx)
		return false, fmt.Errorf("AssetRepository - BuyAsset - tx.Exec: %w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return false, fmt.Errorf("AssetRepository - BuyAsset - tx.Commit: %w", err)
	}
	return true, nil
}

func (r *AssetRepository) GetPurchasedAssets(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	sql, args, err := r.Builder.
		Select("id, name, description, price, owner_id").
		From("assets").
		Join("access_assets ON assets.id = access_assets.asset_id").
		Where(sq.Eq{"access_assets.user_id": user.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AssetRepository - GetPurchasedAssets - r.Builder: %w", err)
	}
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AssetRepository - GetPurchasedAssets - r.Pool.Query: %w", err)
	}
	defer rows.Close()
	assets := make([]entity.Asset, 0)
	for rows.Next() {
		ast := entity.Asset{}
		err := rows.Scan(&ast.Id, &ast.Name, &ast.Description, &ast.Price, &ast.Owner_id)
		if err != nil {
			return nil, fmt.Errorf("AssetRepository - GetOtherUserAssets - rows.Scan: %w", err)
		}
		assets = append(assets, ast)
	}
	return assets, nil
}

func (r *AssetRepository) GetAssetById(ctx context.Context, id int64) (entity.Asset, error) {
	sql, args, err := r.Builder.
		Select("name, description, price, owner_id").
		From("assets").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return entity.Asset{}, fmt.Errorf("AssetRepository - GetAssetById - r.Builder: %w", err)
	}
	row := r.Pool.QueryRow(ctx, sql, args...)
	ast := entity.Asset{Id: id}
	err = row.Scan(&ast.Name, &ast.Description, &ast.Price, &ast.Owner_id)
	if err != nil {
		return entity.Asset{}, fmt.Errorf("AssetRepository - GetAssetById - row.Scan: %w", err)
	}
	return ast, nil
}

// func (r *AssetRepository) UpdateAssetById(ctx context.Context, asset entity.Asset) (entity.Asset, error) {
// 	sql, args, err := r.Builder.
// 		Update("assets").
// 		Set("balance", sq.Expr("balance +?", amount)).
// 		Where(sq.Eq{"id": asset.Id}).
// 		Suffix("RETURNING balance").
// 		ToSql()
// }
