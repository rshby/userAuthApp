package repository

import (
	"context"
	"database/sql"
	"userAuthApp/model/entity"
)

type InterfaceUserRepository interface {
	InsertUser(ctx context.Context, tx *sql.Tx, entity *entity.User) (*entity.User, error)
}
