package usecase

import (
	"context"
	"fmt"

	"github.com/Klef99/bhs-task/internal/entity"
)

// AssetUseCase -.
type AssetUseCase struct {
	repo AssetRepository
}

var _ Asset = (*AssetUseCase)(nil)

// New -.
func NewAssetUseCase(r AssetRepository) *AssetUseCase {
	return &AssetUseCase{repo: r}
}

func (uc *AssetUseCase) CreateAsset(ctx context.Context, ast entity.Asset) (bool, error) {
	if ast.Name == "" || ast.Owner_id <= 0 {
		return false, fmt.Errorf("AssetUseCase - CreateAsset - invalid asset data")
	}
	status, err := uc.repo.Store(ctx, ast)
	if err != nil {
		return false, fmt.Errorf("AssetUseCase - CreateAsset - uc.repo.Store: %w", err)
	}
	return status, err
}

func (uc *AssetUseCase) DeleteAsset(ctx context.Context, user entity.User, id int64) (bool, error) {
	if id <= 0 || user.Id <= 0 {
		return false, fmt.Errorf("AssetUseCase - DeleteAsset - invalid user or asset id")
	}
	status, err := uc.repo.Erase(ctx, user, id)
	if err != nil {
		return false, fmt.Errorf("AssetUseCase - DeleteAsset - uc.repo.Erase: %w", err)
	}
	return status, err
}

func (uc *AssetUseCase) UserAssetsList(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	if user.Id <= 0 {
		return []entity.Asset{}, fmt.Errorf("AssetUseCase - AssetsList - invalid user id")
	}
	assets, err := uc.repo.UserAssetsList(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("AssetUseCase - AssetsList - uc.repo.List: %w", err)
	}
	return assets, nil
}

func (uc *AssetUseCase) GetAssetsToBuying(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	if user.Id <= 0 {
		return []entity.Asset{}, fmt.Errorf("AssetUseCase - GetAssetsToBuying - invalid user id")
	}
	assets, err := uc.repo.GetOtherUsersAssets(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("AssetUseCase - GetAssetsToBuying - uc.repo.GetOtherUserAsset: %w", err)
	}
	return assets, nil
}

func (uc *AssetUseCase) BuyAsset(ctx context.Context, user entity.User, id int64) (bool, error) {
	if user.Id <= 0 {
		return false, fmt.Errorf("AssetUseCase - BuyAsset - invalid user id")
	}
	if id <= 0 {
		return false, fmt.Errorf("AssetUseCase - BuyAsset - invalid asset id")
	}
	status, err := uc.repo.BuyAsset(ctx, user, id)
	if err != nil {
		return false, fmt.Errorf("AssetUseCase - BuyAsset - uc.repo.BuyAsset: %w", err)
	}
	return status, err
}

func (uc *AssetUseCase) GetPurchasedAssets(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	if user.Id <= 0 {
		return []entity.Asset{}, fmt.Errorf("AssetUseCase - GetPurchasedAssets - invalid user id")
	}
	assets, err := uc.repo.GetPurchasedAssets(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("AssetUseCase - GetPurchasedAssets - uc.repo.GetPurchasedAssets: %w", err)
	}
	return assets, nil
}

func (uc *AssetUseCase) GetAssetById(ctx context.Context, id int64) (entity.Asset, error) {
	if id <= 0 {
		return entity.Asset{}, fmt.Errorf("AssetUseCase - GetAssetById - invalid asset id")
	}
	asset, err := uc.repo.GetAssetById(ctx, id)
	if err != nil {
		return entity.Asset{}, fmt.Errorf("AssetUseCase - GetAssetById - uc.repo.GetAssetById: %w", err)
	}
	return asset, nil
}
