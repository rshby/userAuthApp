package repository

import (
	"context"
	"database/sql"
	"userAuthApp/model/dto"
	"userAuthApp/model/entity"
)

type UserRepository struct {
	DB *sql.DB
}

// create function provider
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

// method cerate user
func (u *UserRepository) InsertUser(ctx context.Context, tx *sql.Tx, entity *entity.User) (*entity.User, error) {
	query := "INSERT INTO users(full_name, address, account_id) VALUES($1, $2, $3) RETURNING id"
	res := tx.QueryRowContext(ctx, query, entity.FullName, entity.Address, entity.AccountId)
	if res.Err() != nil {
		return nil, res.Err()
	}

	var id int
	err := res.Scan(&id)
	if err != nil {
		return nil, err
	}

	entity.Id = int(id)
	return entity, nil
}

// method get user by account id
func (u *UserRepository) GetByAccountId(ctx context.Context, tx *sql.Tx, accountId int) (*entity.User, error) {
	query := "SELECT id, full_name, address, account_id FROM users WHERE account_id = $1"

	row := tx.QueryRowContext(ctx, query, accountId)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user entity.User
	err := row.Scan(&user.Id, &user.FullName, &user.Address, &user.AccountId)
	if err != nil {
		return nil, err
	}

	// success get data user by account id
	return &user, nil
}

// method get user detail
func (u *UserRepository) GetUserDetail(ctx context.Context, tx *sql.Tx, username string) (*dto.UserDetailResponse, error) {
	query := "SELECT a.id, a.username, u.id, u.full_name, u.address FROM accounts a JOIN users u ON u.account_id=a.id WHERE a.username=$1"

	row := tx.QueryRowContext(ctx, query, username)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var data dto.UserDetail
	var userDetail dto.UserDetailResponse
	err := row.Scan(&data.AccountId, &data.UserName, &data.UserId, &data.FullName, &data.Address)
	if err != nil {
		return nil, err
	}

	userDetail = dto.UserDetailResponse{
		Account: &entity.Accounts{
			Id:       data.AccountId,
			UserName: data.UserName,
		},
		User: &entity.User{
			Id:       data.UserId,
			FullName: data.FullName,
			Address:  data.Address,
		},
	}

	// success get user detail
	return &userDetail, nil
}
