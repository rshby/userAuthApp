package repository

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"userAuthApp/model/entity"
)

type AccountRepository struct {
	DB *sql.DB
}

// create function provider
func NewAccountRepository(db *sql.DB) InterfaceAccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

// method create account
func (a *AccountRepository) InsertAccount(ctx context.Context, tx *sql.Tx, entity *entity.Accounts) (*entity.Accounts, error) {
	query := "INSERT INTO accounts(username, password) VALUES($1, $2)"
	result, err := tx.ExecContext(ctx, query, entity.UserName, entity.Password)
	if err != nil {
		return nil, err
	}

	if row, _ := result.RowsAffected(); row == 0 {
		return nil, errors.New("cant insert to database")
	}

	return entity, nil
}

// method get data by email
func (a *AccountRepository) GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.Accounts, error) {
	query := "SELECT id, username, password FROM accounts WHERE username=$1"
	row := tx.QueryRowContext(ctx, query, email)
	if row.Err() != nil {
		return nil, errors.New("record not found")
	}

	var account entity.Accounts
	err := row.Scan(&account.Id, &account.UserName, &account.Password)
	if err != nil {
		return nil, err
	}

	// success get data by email
	return &account, nil
}
