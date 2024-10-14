// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=internal/usecase/interfaces.go -destination=internal/usecase/mock_test.go -package=usecase_test
//

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"

	entity "github.com/Klef99/bhs-task/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CheckDeposit mocks base method.
func (m *MockUser) CheckDeposit(ctx context.Context, user entity.User) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDeposit", ctx, user)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDeposit indicates an expected call of CheckDeposit.
func (mr *MockUserMockRecorder) CheckDeposit(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDeposit", reflect.TypeOf((*MockUser)(nil).CheckDeposit), ctx, user)
}

// Login mocks base method.
func (m *MockUser) Login(ctx context.Context, crd entity.Credentials) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, crd)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserMockRecorder) Login(ctx, crd any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUser)(nil).Login), ctx, crd)
}

// MakeDeposit mocks base method.
func (m *MockUser) MakeDeposit(ctx context.Context, user entity.User, amount float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDeposit", ctx, user, amount)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeDeposit indicates an expected call of MakeDeposit.
func (mr *MockUserMockRecorder) MakeDeposit(ctx, user, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDeposit", reflect.TypeOf((*MockUser)(nil).MakeDeposit), ctx, user, amount)
}

// Register mocks base method.
func (m *MockUser) Register(ctx context.Context, crd entity.Credentials) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, crd)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserMockRecorder) Register(ctx, crd any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUser)(nil).Register), ctx, crd)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CheckDeposit mocks base method.
func (m *MockUserRepository) CheckDeposit(ctx context.Context, user entity.User) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDeposit", ctx, user)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDeposit indicates an expected call of CheckDeposit.
func (mr *MockUserRepositoryMockRecorder) CheckDeposit(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDeposit", reflect.TypeOf((*MockUserRepository)(nil).CheckDeposit), ctx, user)
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(ctx context.Context, crd entity.Credentials) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, crd)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(ctx, crd any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), ctx, crd)
}

// LoginUser mocks base method.
func (m *MockUserRepository) LoginUser(ctx context.Context, crd entity.Credentials) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", ctx, crd)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockUserRepositoryMockRecorder) LoginUser(ctx, crd any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockUserRepository)(nil).LoginUser), ctx, crd)
}

// MakeDeposit mocks base method.
func (m *MockUserRepository) MakeDeposit(ctx context.Context, user entity.User, amount float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDeposit", ctx, user, amount)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeDeposit indicates an expected call of MakeDeposit.
func (mr *MockUserRepositoryMockRecorder) MakeDeposit(ctx, user, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDeposit", reflect.TypeOf((*MockUserRepository)(nil).MakeDeposit), ctx, user, amount)
}

// MockAsset is a mock of Asset interface.
type MockAsset struct {
	ctrl     *gomock.Controller
	recorder *MockAssetMockRecorder
}

// MockAssetMockRecorder is the mock recorder for MockAsset.
type MockAssetMockRecorder struct {
	mock *MockAsset
}

// NewMockAsset creates a new mock instance.
func NewMockAsset(ctrl *gomock.Controller) *MockAsset {
	mock := &MockAsset{ctrl: ctrl}
	mock.recorder = &MockAssetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAsset) EXPECT() *MockAssetMockRecorder {
	return m.recorder
}

// BuyAsset mocks base method.
func (m *MockAsset) BuyAsset(ctx context.Context, user entity.User, id int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyAsset", ctx, user, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuyAsset indicates an expected call of BuyAsset.
func (mr *MockAssetMockRecorder) BuyAsset(ctx, user, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyAsset", reflect.TypeOf((*MockAsset)(nil).BuyAsset), ctx, user, id)
}

// CreateAsset mocks base method.
func (m *MockAsset) CreateAsset(ctx context.Context, ast entity.Asset) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAsset", ctx, ast)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAsset indicates an expected call of CreateAsset.
func (mr *MockAssetMockRecorder) CreateAsset(ctx, ast any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAsset", reflect.TypeOf((*MockAsset)(nil).CreateAsset), ctx, ast)
}

// DeleteAsset mocks base method.
func (m *MockAsset) DeleteAsset(ctx context.Context, user entity.User, id int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAsset", ctx, user, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAsset indicates an expected call of DeleteAsset.
func (mr *MockAssetMockRecorder) DeleteAsset(ctx, user, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAsset", reflect.TypeOf((*MockAsset)(nil).DeleteAsset), ctx, user, id)
}

// GetAssetById mocks base method.
func (m *MockAsset) GetAssetById(ctx context.Context, id int64) (entity.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetById", ctx, id)
	ret0, _ := ret[0].(entity.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetById indicates an expected call of GetAssetById.
func (mr *MockAssetMockRecorder) GetAssetById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetById", reflect.TypeOf((*MockAsset)(nil).GetAssetById), ctx, id)
}

// GetAssetsToBuying mocks base method.
func (m *MockAsset) GetAssetsToBuying(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetsToBuying", ctx, user)
	ret0, _ := ret[0].([]entity.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetsToBuying indicates an expected call of GetAssetsToBuying.
func (mr *MockAssetMockRecorder) GetAssetsToBuying(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetsToBuying", reflect.TypeOf((*MockAsset)(nil).GetAssetsToBuying), ctx, user)
}

// GetPurchasedAsset mocks base method.
func (m *MockAsset) GetPurchasedAsset(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPurchasedAsset", ctx, user)
	ret0, _ := ret[0].([]entity.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPurchasedAsset indicates an expected call of GetPurchasedAsset.
func (mr *MockAssetMockRecorder) GetPurchasedAsset(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPurchasedAsset", reflect.TypeOf((*MockAsset)(nil).GetPurchasedAsset), ctx, user)
}

// UserAssetsList mocks base method.
func (m *MockAsset) UserAssetsList(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserAssetsList", ctx, user)
	ret0, _ := ret[0].([]entity.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserAssetsList indicates an expected call of UserAssetsList.
func (mr *MockAssetMockRecorder) UserAssetsList(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserAssetsList", reflect.TypeOf((*MockAsset)(nil).UserAssetsList), ctx, user)
}

// MockAssetRepository is a mock of AssetRepository interface.
type MockAssetRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAssetRepositoryMockRecorder
}

// MockAssetRepositoryMockRecorder is the mock recorder for MockAssetRepository.
type MockAssetRepositoryMockRecorder struct {
	mock *MockAssetRepository
}

// NewMockAssetRepository creates a new mock instance.
func NewMockAssetRepository(ctrl *gomock.Controller) *MockAssetRepository {
	mock := &MockAssetRepository{ctrl: ctrl}
	mock.recorder = &MockAssetRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssetRepository) EXPECT() *MockAssetRepositoryMockRecorder {
	return m.recorder
}

// BuyAsset mocks base method.
func (m *MockAssetRepository) BuyAsset(ctx context.Context, user entity.User, id int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyAsset", ctx, user, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuyAsset indicates an expected call of BuyAsset.
func (mr *MockAssetRepositoryMockRecorder) BuyAsset(ctx, user, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyAsset", reflect.TypeOf((*MockAssetRepository)(nil).BuyAsset), ctx, user, id)
}

// Erase mocks base method.
func (m *MockAssetRepository) Erase(ctx context.Context, user entity.User, id int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Erase", ctx, user, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Erase indicates an expected call of Erase.
func (mr *MockAssetRepositoryMockRecorder) Erase(ctx, user, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Erase", reflect.TypeOf((*MockAssetRepository)(nil).Erase), ctx, user, id)
}

// GetAssetById mocks base method.
func (m *MockAssetRepository) GetAssetById(ctx context.Context, id int64) (entity.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetById", ctx, id)
	ret0, _ := ret[0].(entity.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetById indicates an expected call of GetAssetById.
func (mr *MockAssetRepositoryMockRecorder) GetAssetById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetById", reflect.TypeOf((*MockAssetRepository)(nil).GetAssetById), ctx, id)
}

// GetOtherUsersAssets mocks base method.
func (m *MockAssetRepository) GetOtherUsersAssets(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOtherUsersAssets", ctx, user)
	ret0, _ := ret[0].([]entity.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOtherUsersAssets indicates an expected call of GetOtherUsersAssets.
func (mr *MockAssetRepositoryMockRecorder) GetOtherUsersAssets(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOtherUsersAssets", reflect.TypeOf((*MockAssetRepository)(nil).GetOtherUsersAssets), ctx, user)
}

// GetPurchasedAssets mocks base method.
func (m *MockAssetRepository) GetPurchasedAssets(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPurchasedAssets", ctx, user)
	ret0, _ := ret[0].([]entity.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPurchasedAssets indicates an expected call of GetPurchasedAssets.
func (mr *MockAssetRepositoryMockRecorder) GetPurchasedAssets(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPurchasedAssets", reflect.TypeOf((*MockAssetRepository)(nil).GetPurchasedAssets), ctx, user)
}

// Store mocks base method.
func (m *MockAssetRepository) Store(ctx context.Context, ast entity.Asset) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, ast)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockAssetRepositoryMockRecorder) Store(ctx, ast any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockAssetRepository)(nil).Store), ctx, ast)
}

// UserAssetsList mocks base method.
func (m *MockAssetRepository) UserAssetsList(ctx context.Context, user entity.User) ([]entity.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserAssetsList", ctx, user)
	ret0, _ := ret[0].([]entity.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserAssetsList indicates an expected call of UserAssetsList.
func (mr *MockAssetRepositoryMockRecorder) UserAssetsList(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserAssetsList", reflect.TypeOf((*MockAssetRepository)(nil).UserAssetsList), ctx, user)
}