package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Klef99/bhs-task/internal/entity"
	"github.com/Klef99/bhs-task/internal/usecase"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

type createAssetTest struct {
	name string
	ast  entity.Asset
	mock func()
	res  bool
	err  error
}

type deleteAssetTest struct {
	name string
	user entity.User
	id   int64
	mock func()
	res  bool
	err  error
}

type userAssetsListTest struct {
	name string
	user entity.User
	mock func()
	res  []entity.Asset
	err  error
}

type getAssetsToBuyingTest struct {
	name string
	user entity.User
	mock func()
	res  []entity.Asset
	err  error
}

type buyAssetTest struct {
	name string
	user entity.User
	id   int64
	mock func()
	res  bool
	err  error
}

type getPurchasedAssetTest struct {
	name string
	user entity.User
	mock func()
	res  []entity.Asset
	err  error
}

type getAssetByIdTest struct {
	name string
	id   int64
	mock func()
	res  entity.Asset
	err  error
}

func AssetUseCase(t *testing.T) (*usecase.AssetUseCase, *MockAssetRepository) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockAssetRepository(mockCtl)

	UserUseCase := usecase.NewAssetUseCase(repo)
	return UserUseCase, repo
}

func TestCreateAsset(t *testing.T) {
	t.Parallel()

	asset, repo := AssetUseCase(t)
	tests := []createAssetTest{
		{
			name: "empty asset",
			ast:  entity.Asset{},
			mock: func() {
				repo.EXPECT().Store(context.Background(), entity.Asset{}).Return(false, errInternalServErr)
			},
			res: false,
			err: fmt.Errorf("AssetUseCase - CreateAsset - invalid asset data"),
		},
		{
			name: "success",
			ast:  entity.Asset{Owner_id: 1, Name: "Sword", Description: "Rare", Price: 100},
			mock: func() {
				repo.EXPECT().Store(context.Background(), entity.Asset{Owner_id: 1, Name: "Sword", Description: "Rare", Price: 100}).Return(true, nil)
			},
			res: true,
			err: nil,
		},
		{
			name: "success without description",
			ast:  entity.Asset{Owner_id: 1, Name: "Sword", Price: 100},
			mock: func() {
				repo.EXPECT().Store(context.Background(), entity.Asset{Owner_id: 1, Name: "Sword", Price: 100}).Return(true, nil)
			},
			res: true,
			err: nil,
		},
		{
			name: "success with only name",
			ast:  entity.Asset{Owner_id: 1, Name: "Sword"},
			mock: func() {
				repo.EXPECT().Store(context.Background(), entity.Asset{Owner_id: 1, Name: "Sword"}).Return(true, nil)
			},
			res: true,
			err: nil,
		},
		{
			name: "without owner id",
			ast:  entity.Asset{Name: "Sword"},
			mock: func() {
				repo.EXPECT().Store(context.Background(), entity.Asset{Name: "Sword"}).Return(false, errInternalServErr)
			},
			res: false,
			err: fmt.Errorf("AssetUseCase - CreateAsset - invalid asset data"),
		},
		{
			name: "negative price",
			ast:  entity.Asset{Owner_id: 1, Name: "Sword", Price: -1},
			mock: func() {
				repo.EXPECT().Store(context.Background(), entity.Asset{Owner_id: 1, Name: "Sword", Price: -1}).Return(false, errInternalServErr)
			},
			res: false,
			err: errInternalServErr,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := asset.CreateAsset(context.Background(), tc.ast)
			require.Equal(t, res, tc.res)
			if err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestDeleteAsset(t *testing.T) {
	t.Parallel()

	asset, repo := AssetUseCase(t)
	tests := []deleteAssetTest{
		{
			name: "empty user",
			user: entity.User{},
			id:   1,
			mock: func() {
				repo.EXPECT().Erase(context.Background(), entity.User{}, int64(1)).Return(false, errInternalServErr)
			},
			res: false,
			err: fmt.Errorf("AssetUseCase - DeleteAsset - invalid user or asset id"),
		},
		{
			name: "invalid asset id",
			user: entity.User{Id: 1, Username: "test"},
			id:   0,
			mock: func() {
				repo.EXPECT().Erase(context.Background(), entity.User{Id: 1, Username: "test"}, int64(0)).Return(false, errInternalServErr)
			},
			res: false,
			err: fmt.Errorf("AssetUseCase - DeleteAsset - invalid user or asset id"),
		},
		{
			name: "invalid user id",
			user: entity.User{Id: 0, Username: "test"},
			id:   1,
			mock: func() {
				repo.EXPECT().Erase(context.Background(), entity.User{Id: 0, Username: "test"}, int64(1)).Return(false, errInternalServErr)
			},
			res: false,
			err: fmt.Errorf("AssetUseCase - DeleteAsset - invalid user or asset id"),
		},
		{
			name: "success",
			user: entity.User{Id: 1, Username: "test"},
			id:   1,
			mock: func() {
				repo.EXPECT().Erase(context.Background(), entity.User{Id: 1, Username: "test"}, int64(1)).Return(true, nil)
			},
			res: true,
			err: nil,
		},
		{
			name: "asset not exist",
			user: entity.User{Id: 2, Username: "test"},
			id:   3,
			mock: func() {
				repo.EXPECT().Erase(context.Background(), entity.User{Id: 2, Username: "test"}, int64(3)).Return(false, errInternalServErr)
			},
			res: false,
			err: errInternalServErr,
		},
		{
			name: "user not owner",
			user: entity.User{Id: 1, Username: "test"},
			id:   2,
			mock: func() {
				repo.EXPECT().Erase(context.Background(), entity.User{Id: 1, Username: "test"}, int64(2)).Return(false, errInternalServErr)
			},
			res: false,
			err: errInternalServErr,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := asset.DeleteAsset(context.Background(), tc.user, tc.id)
			require.Equal(t, res, tc.res)
			if err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestUserAssetsList(t *testing.T) {
	t.Parallel()

	asset, repo := AssetUseCase(t)
	tests := []userAssetsListTest{
		{
			name: "empty user",
			user: entity.User{},
			mock: func() {
				repo.EXPECT().UserAssetsList(context.Background(), entity.User{}).Return([]entity.Asset{}, errInternalServErr)
			},
			res: []entity.Asset{},
			err: fmt.Errorf("AssetUseCase - AssetsList - invalid user id"),
		},
		{
			name: "invalid user id",
			user: entity.User{Id: 0, Username: "test"},
			mock: func() {
				repo.EXPECT().UserAssetsList(context.Background(), entity.User{Id: 0, Username: "test"}).Return([]entity.Asset{}, errInternalServErr)
			},
			res: []entity.Asset{},
			err: fmt.Errorf("AssetUseCase - AssetsList - invalid user id"),
		},
		{
			name: "success",
			user: entity.User{Id: 1, Username: "test"},
			mock: func() {
				repo.EXPECT().UserAssetsList(context.Background(), entity.User{Id: 1, Username: "test"}).Return([]entity.Asset{{Id: 1, Name: "test", Owner_id: 1}}, nil)
			},
			res: []entity.Asset{{Id: 1, Name: "test", Owner_id: 1}},
			err: nil,
		},
		{
			name: "user not exist",
			user: entity.User{Id: 3, Username: "test3"},
			mock: func() {
				repo.EXPECT().UserAssetsList(context.Background(), entity.User{Id: 3, Username: "test3"}).Return([]entity.Asset{}, errInternalServErr)
			},
			res: []entity.Asset{},
			err: errInternalServErr,
		},
		{
			name: "assets not found",
			user: entity.User{Id: 2, Username: "test"},
			mock: func() {
				repo.EXPECT().UserAssetsList(context.Background(), entity.User{Id: 2, Username: "test"}).Return([]entity.Asset{}, nil)
			},
			res: []entity.Asset{},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := asset.UserAssetsList(context.Background(), tc.user)

			// Adjust comparison for the slice result
			require.ElementsMatch(t, res, tc.res)
			if err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestGetAssetsToBuying(t *testing.T) {
	t.Parallel()

	asset, repo := AssetUseCase(t)
	tests := []getAssetsToBuyingTest{
		{
			name: "empty user",
			user: entity.User{},
			mock: func() {
				repo.EXPECT().GetOtherUsersAssets(context.Background(), entity.User{}).Return([]entity.Asset{}, errInternalServErr)
			},
			res: []entity.Asset{},
			err: fmt.Errorf("AssetUseCase - GetAssetsToBuying - invalid user id"),
		},
		{
			name: "invalid user id",
			user: entity.User{Id: 0, Username: "test"},
			mock: func() {
				repo.EXPECT().GetOtherUsersAssets(context.Background(), entity.User{Id: 0, Username: "test"}).Return([]entity.Asset{}, errInternalServErr)
			},
			res: []entity.Asset{},
			err: fmt.Errorf("AssetUseCase - GetAssetsToBuying - invalid user id"),
		},
		{
			name: "success",
			user: entity.User{Id: 1, Username: "test"},
			mock: func() {
				repo.EXPECT().GetOtherUsersAssets(context.Background(), entity.User{Id: 1, Username: "test"}).Return([]entity.Asset{{Id: 1, Name: "test", Owner_id: 1}}, nil)
			},
			res: []entity.Asset{{Id: 1, Name: "test", Owner_id: 1}},
			err: nil,
		},
		{
			name: "user not exist",
			user: entity.User{Id: 3, Username: "test3"},
			mock: func() {
				repo.EXPECT().GetOtherUsersAssets(context.Background(), entity.User{Id: 3, Username: "test3"}).Return([]entity.Asset{}, errInternalServErr)
			},
			res: []entity.Asset{},
			err: errInternalServErr,
		},
		{
			name: "assets not found",
			user: entity.User{Id: 2, Username: "test"},
			mock: func() {
				repo.EXPECT().GetOtherUsersAssets(context.Background(), entity.User{Id: 2, Username: "test"}).Return([]entity.Asset{}, nil)
			},
			res: []entity.Asset{},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := asset.GetAssetsToBuying(context.Background(), tc.user)
			require.ElementsMatch(t, res, tc.res)
			if err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestBuyAsset(t *testing.T) {
	t.Parallel()

	asset, repo := AssetUseCase(t)
	tests := []buyAssetTest{
		{
			name: "empty user",
			user: entity.User{},
			id:   1,
			mock: func() {
				repo.EXPECT().BuyAsset(context.Background(), entity.User{}, int64(1)).Return(false, errInternalServErr)
			},
			res: false,
			err: fmt.Errorf("AssetUseCase - BuyAsset - invalid user id"),
		},
		{
			name: "invalid asset id",
			user: entity.User{Id: 1, Username: "test"},
			id:   0,
			mock: func() {
				repo.EXPECT().BuyAsset(context.Background(), entity.User{Id: 1, Username: "test"}, int64(0)).Return(false, errInternalServErr)
			},
			res: false,
			err: fmt.Errorf("AssetUseCase - BuyAsset - invalid asset id"),
		},
		{
			name: "invalid user id",
			user: entity.User{Id: 0, Username: "test"},
			id:   1,
			mock: func() {
				repo.EXPECT().BuyAsset(context.Background(), entity.User{Id: 0, Username: "test"}, int64(1)).Return(false, errInternalServErr)
			},
			res: false,
			err: fmt.Errorf("AssetUseCase - BuyAsset - invalid user id"),
		},
		{
			name: "success",
			user: entity.User{Id: 1, Username: "test"},
			id:   1,
			mock: func() {
				repo.EXPECT().BuyAsset(context.Background(), entity.User{Id: 1, Username: "test"}, int64(1)).Return(true, nil)
			},
			res: true,
			err: nil,
		},
		{
			name: "asset not exist",
			user: entity.User{Id: 1, Username: "test"},
			id:   3,
			mock: func() {
				repo.EXPECT().BuyAsset(context.Background(), entity.User{Id: 1, Username: "test"}, int64(3)).Return(false, errInternalServErr)
			},
			res: false,
			err: errInternalServErr,
		},
		{
			name: "user is owner",
			user: entity.User{Id: 1, Username: "test"},
			id:   2,
			mock: func() {
				repo.EXPECT().BuyAsset(context.Background(), entity.User{Id: 1, Username: "test"}, int64(2)).Return(false, errInternalServErr)
			},
			res: false,
			err: errInternalServErr,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := asset.BuyAsset(context.Background(), tc.user, tc.id)
			require.Equal(t, res, tc.res)
			if err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestGetPurchasedAsset(t *testing.T) {
	t.Parallel()

	asset, repo := AssetUseCase(t)
	tests := []getPurchasedAssetTest{
		{
			name: "empty user",
			user: entity.User{},
			mock: func() {
				repo.EXPECT().GetPurchasedAssets(context.Background(), entity.User{}).Return([]entity.Asset{}, errInternalServErr)
			},
			res: []entity.Asset{},
			err: fmt.Errorf("AssetUseCase - GetPurchasedAssets - invalid user id"),
		},
		{
			name: "invalid user id",
			user: entity.User{Id: 0, Username: "test"},
			mock: func() {
				repo.EXPECT().GetPurchasedAssets(context.Background(), entity.User{Id: 0, Username: "test"}).Return([]entity.Asset{}, errInternalServErr)
			},
			res: []entity.Asset{},
			err: fmt.Errorf("AssetUseCase - GetPurchasedAssets - invalid user id"),
		},
		{
			name: "user not exist",
			user: entity.User{Id: 3, Username: "test3"},
			mock: func() {
				repo.EXPECT().GetPurchasedAssets(context.Background(), entity.User{Id: 3, Username: "test3"}).Return([]entity.Asset{}, errInternalServErr)
			},
			res: []entity.Asset{},
			err: errInternalServErr,
		},
		{
			name: "success",
			user: entity.User{Id: 1, Username: "test"},
			mock: func() {
				repo.EXPECT().GetPurchasedAssets(context.Background(), entity.User{Id: 1, Username: "test"}).Return([]entity.Asset{{Id: 1, Name: "test", Owner_id: 1}}, nil)
			},
			res: []entity.Asset{{Id: 1, Name: "test", Owner_id: 1}},
			err: nil,
		},
		{
			name: "assets not found",
			user: entity.User{Id: 2, Username: "test"},
			mock: func() {
				repo.EXPECT().GetPurchasedAssets(context.Background(), entity.User{Id: 2, Username: "test"}).Return([]entity.Asset{}, nil)
			},
			res: []entity.Asset{},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := asset.GetPurchasedAssets(context.Background(), tc.user)
			require.ElementsMatch(t, res, tc.res)
			if err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestGetAssetById(t *testing.T) {
	t.Parallel()

	asset, repo := AssetUseCase(t)
	tests := []getAssetByIdTest{
		{
			name: "success",
			id:   1,
			mock: func() {
				repo.EXPECT().GetAssetById(context.Background(), int64(1)).Return(entity.Asset{Id: 1, Name: "Sword"}, nil)
			},
			res: entity.Asset{Id: 1, Name: "Sword"},
			err: nil,
		},
		{
			name: "invalid asset id",
			id:   -1,
			mock: func() {
				repo.EXPECT().GetAssetById(context.Background(), int64(-1)).Return(entity.Asset{}, errInternalServErr)
			},
			res: entity.Asset{},
			err: fmt.Errorf("AssetUseCase - GetAssetById - invalid asset id"),
		},
		{
			name: "assets not found",
			id:   2,
			mock: func() {
				repo.EXPECT().GetAssetById(context.Background(), int64(2)).Return(entity.Asset{}, errInternalServErr)
			},
			res: entity.Asset{},
			err: errInternalServErr,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := asset.GetAssetById(context.Background(), tc.id)
			require.Equal(t, res, tc.res)
			if err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
