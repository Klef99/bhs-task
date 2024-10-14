package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Klef99/bhs-task/internal/entity"
	"github.com/Klef99/bhs-task/internal/usecase"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

var errInternalServErr = errors.New("internal server error")

type authTest struct {
	name string
	crd  entity.Credentials
	mock func()
	res  interface{}
	err  error
}

type checkDepositTest struct {
	name string
	user entity.User
	mock func()
	res  float64
	err  error
}

type makeDepositTest struct {
	name   string
	user   entity.User
	amount float64
	mock   func()
	res    float64
	err    error
}

func UserUseCase(t *testing.T) (*usecase.UserUseCase, *MockUserRepository) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockUserRepository(mockCtl)

	UserUseCase := usecase.NewUserUseCase(repo)

	return UserUseCase, repo
}

func TestRegister(t *testing.T) {
	t.Parallel()

	user, repo := UserUseCase(t)
	tests := []authTest{
		{
			name: "empty result",
			crd:  entity.Credentials{},
			mock: func() {
				repo.EXPECT().CreateUser(context.Background(), entity.Credentials{}).Return(false, errInternalServErr)
			},
			res: false,
			err: errInternalServErr,
		},
		{
			name: "success",
			crd:  entity.Credentials{Username: "test", Password: "pass"},
			mock: func() {
				repo.EXPECT().CreateUser(context.Background(), entity.Credentials{Username: "test", Password: "pass"}).Return(true, nil)
			},
			res: true,
			err: nil,
		},
		{
			name: "user exist",
			crd:  entity.Credentials{Username: "test", Password: "pass"},
			mock: func() {
				repo.EXPECT().CreateUser(context.Background(), entity.Credentials{Username: "test", Password: "pass"}).Return(false, errInternalServErr)
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

			res, err := user.Register(context.Background(), tc.crd)
			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestLogin(t *testing.T) {
	t.Parallel()

	user, repo := UserUseCase(t)
	tests := []authTest{
		{
			name: "empty result",
			crd:  entity.Credentials{},
			mock: func() {
				repo.EXPECT().LoginUser(context.Background(), entity.Credentials{}).Return(int64(-1), errInternalServErr)
			},
			res: entity.User{},
			err: errInternalServErr,
		},
		{
			name: "success",
			crd:  entity.Credentials{Username: "test", Password: "pass"},
			mock: func() {
				repo.EXPECT().LoginUser(context.Background(), entity.Credentials{Username: "test", Password: "pass"}).Return(int64(1), nil)
			},
			res: entity.User{Id: 1, Username: "test"},
			err: nil,
		},
		{
			name: "user not exist",
			crd:  entity.Credentials{Username: "test", Password: "pass"},
			mock: func() {
				repo.EXPECT().LoginUser(context.Background(), entity.Credentials{Username: "test", Password: "pass"}).Return(int64(-1), errInternalServErr)
			},
			res: entity.User{},
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := user.Login(context.Background(), tc.crd)
			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestCheckDeposit(t *testing.T) {
	t.Parallel()

	user, repo := UserUseCase(t)
	tests := []checkDepositTest{
		{
			name: "empty user",
			user: entity.User{},
			mock: func() {
				repo.EXPECT().CheckDeposit(context.Background(), entity.User{}).Return(float64(0), errInternalServErr)
			},
			res: 0,
			err: errInternalServErr,
		},
		{
			name: "success",
			user: entity.User{Username: "test", Id: 1},
			mock: func() {
				repo.EXPECT().CheckDeposit(context.Background(), entity.User{Username: "test", Id: 1}).Return(float64(100), nil)
			},
			res: float64(100),
			err: nil,
		},
		{
			name: "user not exist",
			user: entity.User{Username: "test", Id: 1},
			mock: func() {
				repo.EXPECT().CheckDeposit(context.Background(), entity.User{Username: "test", Id: 1}).Return(float64(0), errInternalServErr)
			},
			res: 0,
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := user.CheckDeposit(context.Background(), tc.user)
			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestMakeDeposit(t *testing.T) {
	t.Parallel()

	user, repo := UserUseCase(t)
	tests := []makeDepositTest{
		{
			name:   "empty user",
			user:   entity.User{},
			amount: float64(11),
			mock: func() {
				repo.EXPECT().MakeDeposit(context.Background(), entity.User{}, float64(11)).Return(float64(-1), errInternalServErr)
			},
			res: float64(-1),
			err: errInternalServErr,
		},
		{
			name:   "success",
			user:   entity.User{Username: "test", Id: 1},
			amount: float64(10),
			mock: func() {
				repo.EXPECT().MakeDeposit(context.Background(), entity.User{Username: "test", Id: 1}, float64(10)).Return(float64(110), nil)
			},
			res: float64(110),
			err: nil,
		},
		{
			name:   "user not exist",
			user:   entity.User{Username: "test", Id: 2},
			amount: 10,
			mock: func() {
				repo.EXPECT().MakeDeposit(context.Background(), entity.User{Username: "test", Id: 2}, float64(10)).Return(float64(-1), errInternalServErr)
			},
			res: float64(-1),
			err: errInternalServErr,
		},
		{
			name:   "amount is below 0",
			user:   entity.User{Username: "test", Id: 1},
			amount: -10,
			mock: func() {
				repo.EXPECT().MakeDeposit(context.Background(), entity.User{Username: "test", Id: 1}, float64(-10)).Return(float64(-1), errInternalServErr)
			},
			res: float64(-1),
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := user.MakeDeposit(context.Background(), tc.user, tc.amount)
			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
