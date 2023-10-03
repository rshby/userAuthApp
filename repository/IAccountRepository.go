package repository

import (
	"context"
	"database/sql"
	"userAuthApp/model/entity"
)

type InterfaceAccountRepository interface {
	InsertAccount(ctx context.Context, tx *sql.Tx, entity *entity.Accounts) (*entity.Accounts, error)
	GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.Accounts, error)
	GetById(ctx context.Context, tx *sql.Tx, id int) (*entity.Accounts, error)
}
