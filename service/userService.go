package service

import (
	"context"
	"database/sql"
	"errors"
	"userAuthApp/model/dto"
	"userAuthApp/model/entity"
	"userAuthApp/repository"
)

type UserService struct {
	DB                *sql.DB
	UserRepository    repository.InterfaceUserRepository
	AccountRepository repository.InterfaceAccountRepository
}

// create function provider
func NewUserService(db *sql.DB, userRepo repository.InterfaceUserRepository, accountRepo repository.InterfaceAccountRepository) *UserService {
	return &UserService{
		DB:                db,
		UserRepository:    userRepo,
		AccountRepository: accountRepo,
	}
}

// method insert
func (u *UserService) InsertUser(ctx context.Context, request *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// cek account_id
	account, err := u.AccountRepository.GetById(ctx, tx, request.AccountId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = u.UserRepository.GetByAccountId(ctx, tx, account.Id)
	if err == nil {
		tx.Rollback()
		return nil, errors.New("user with related account_id already exist")
	}

	// create entity
	newUser := entity.User{
		FullName:  request.FullName,
		Address:   request.Address,
		AccountId: request.AccountId,
	}

	// call procedure insert in repository
	result, err := u.UserRepository.InsertUser(ctx, tx, &newUser)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// create response
	response := dto.CreateUserResponse{
		Id:       result.Id,
		FullName: result.FullName,
		Address:  result.Address,
		Account:  account,
	}

	tx.Commit()
	return &response, nil
}

// method get users
func (u *UserService) Getuser(ctx context.Context) (*dto.UserDetailResponse, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// get account by email
	username := ctx.Value("username")
	userDetail, err := u.UserRepository.GetUserDetail(ctx, tx, username.(string))
	if err != nil {
		return nil, err
	}

	// success get user detail
	return userDetail, nil
}
