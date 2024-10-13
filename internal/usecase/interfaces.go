package usecase

import (
	"context"

	"github.com/Klef99/bhs-task/internal/entity"
)

type (
	User interface {
		Register(ctx context.Context, crd entity.Credentials) (bool, error)
		Login(ctx context.Context, crd entity.Credentials) (entity.User, error)

		MakeDeposit(ctx context.Context, user entity.User, amount float64) (float64, error)
		CheckDeposit(ctx context.Context, user entity.User) (float64, error)
	}

	UserRepository interface {
		CreateUser(ctx context.Context, crd entity.Credentials) (bool, error)
		LoginUser(ctx context.Context, crd entity.Credentials) (int64, error)

		MakeDeposit(ctx context.Context, user entity.User, amount float64) (float64, error)
		CheckDeposit(ctx context.Context, user entity.User) (float64, error)
	}

	Asset interface {
		CreateAsset(ctx context.Context, ast entity.Asset) (bool, error)
		DeleteAsset(ctx context.Context, user entity.User, id int64) (bool, error)
		UserAssetsList(ctx context.Context, user entity.User) ([]entity.Asset, error)

		GetAssetsToBuying(ctx context.Context, user entity.User) ([]entity.Asset, error)
		BuyAsset(ctx context.Context, user entity.User, id int64) (bool, error)
		GetAllAvaliableAsset(ctx context.Context, user entity.User) ([]entity.Asset, error)
	}

	AssetRepository interface {
		Store(ctx context.Context, ast entity.Asset) (bool, error)
		Erase(ctx context.Context, user entity.User, id int64) (bool, error)
		UserAssetsList(ctx context.Context, user entity.User) ([]entity.Asset, error)

		GetOtherUsersAssets(ctx context.Context, user entity.User) ([]entity.Asset, error)
		BuyAsset(ctx context.Context, user entity.User, id int64) (bool, error)
		GetPurchasedAssets(ctx context.Context, user entity.User) ([]entity.Asset, error)
	}
)
