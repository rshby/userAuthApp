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
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

// method create account
func (a *AccountRepository) InsertAccount(ctx context.Context, tx *sql.Tx, entity *entity.Accounts) (*entity.Accounts, error) {
	query := "INSERT INTO accounts(username, password) VALUES($1, $2) RETURNING id"
	result := tx.QueryRowContext(ctx, query, entity.UserName, entity.Password)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var id int
	err := result.Scan(&id)
	if err != nil {
		return nil, err
	}

	entity.Id = int(id)
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

// method get account by id
func (a *AccountRepository) GetById(ctx context.Context, tx *sql.Tx, id int) (*entity.Accounts, error) {
	query := "SELECT id, username, password FROM accounts WHERE id = $1"

	row := tx.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, errors.New("record not found")
	}

	var account entity.Accounts
	err := row.Scan(&account.Id, &account.UserName, &account.Password)
	if err != nil {
		return nil, errors.New("record account not found")
	}

	return &account, nil
}
