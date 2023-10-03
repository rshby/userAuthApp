package service

import (
	"context"
	"database/sql"
	"errors"
	"time"
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
func NewAccountService(db *sql.DB, accountRepo repository.InterfaceAccountRepository) InterfaceAccountService {
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

	// get data by email
	account, err := a.AccountRepository.GetByEmail(ctx, tx, request.UserName)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("username not found in our database")
	}

	// cek password match
	isMatch, err := helper.CheckPasword(account.Password, request.Password)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// if password not match
	if !isMatch {
		tx.Rollback()
		return nil, errors.New("password not match")
	}

	// generate token
	token, err := helper.GenerateToken(account.UserName)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response := dto.LoginResponse{
		LoginAt: time.Now().Format("2006-01-02 15:04:05"),
		Token:   token,
	}

	tx.Commit()
	return &response, nil
}
