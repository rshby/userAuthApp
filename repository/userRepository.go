package repository

import (
	"context"
	"database/sql"
	"errors"
	"userAuthApp/model/entity"
)

type UserRepository struct {
	DB *sql.DB
}

// create function provider
func NewUserRepository(db *sql.DB) InterfaceUserRepository {
	return &UserRepository{
		db,
	}
}

// method cerate user
func (u *UserRepository) InsertUser(ctx context.Context, tx *sql.Tx, entity *entity.User) (*entity.User, error) {
	query := "INSERT INTO users(full_name, address, account_id) VALUES($1, $2, $3)"
	res, err := tx.ExecContext(ctx, query, entity.FullName, entity.Address, entity.AccountId)
	if err != nil {
		return nil, err
	}

	if row, _ := res.RowsAffected(); row == 0 {
		return nil, errors.New("error cant insert user")
	}

	return nil, nil
}
