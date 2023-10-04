package repository

import (
	"context"
	"database/sql"
	"userAuthApp/model/dto"
	"userAuthApp/model/entity"
)

type InterfaceUserRepository interface {
	InsertUser(ctx context.Context, tx *sql.Tx, entity *entity.User) (*entity.User, error)
	GetByAccountId(ctx context.Context, tx *sql.Tx, accountId int) (*entity.User, error)
	GetUserDetail(ctx context.Context, tx *sql.Tx, accountId int) (*dto.UserDetailResponse, error)
}
