package service

import (
	"context"
	"database/sql"
	"userAuthApp/helper"
	"userAuthApp/model/dto"
	"userAuthApp/model/entity"
	"userAuthApp/repository"
)

type AccountService struct {
	DB                *sql.DB
	AccountRepository repository.InterfaceAccountRepository
}

// create function provider
func NewAccountService(db *sql.DB, accountRepo repository.InterfaceAccountRepository) *AccountService {
	return &AccountService{
		DB:                db,
		AccountRepository: accountRepo,
	}
}

// method insert data account
func (a *AccountService) InsertAccount(ctx context.Context, request *dto.CreateAccountRequest) (*entity.Accounts, error) {
	tx, err := a.DB.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	// hashedPassword
	password, err := helper.HashPassword(request.Password)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// mapping data
	newAccount := entity.Accounts{
		UserName: request.UserName,
		Password: password,
	}

	// call function in repository
	result, err := a.AccountRepository.InsertAccount(ctx, tx, &newAccount)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response, err := a.AccountRepository.GetByEmail(ctx, tx, result.UserName)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return response, nil
}

// method login
func (a *AccountService) Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error) {
	tx, err := a.DB.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
}
