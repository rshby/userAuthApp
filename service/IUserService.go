package service

import (
	"context"
	"userAuthApp/model/dto"
)

type InterfaceUserService interface {
	InsertUser(ctx context.Context, request *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	Getuser(ctx context.Context) (*dto.UserDetailResponse, error)
}
