package service

import (
	"context"
	"userAuthApp/model/dto"
	"userAuthApp/model/entity"
)

type InterfaceAccountService interface {
	InsertAccount(ctx context.Context, request *dto.CreateAccountRequest) (*entity.Accounts, error)
	Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error)
}
