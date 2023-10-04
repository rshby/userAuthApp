package testing

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"userAuthApp/mock/repo"
	"userAuthApp/model/dto"
	"userAuthApp/service"
)

var db, _ = NewMock()
var accountRepo = &repo.InterfaceAccountRepository{}
var accountService = service.NewAccountService(db, accountRepo)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return db, mock
}

func TestLoginFailed(t *testing.T) {
	ctx := context.Background()
	tx, _ := db.Begin()

	accountRepo.On("GetByEmail", ctx, tx, "reoshby@gmail.com").Return(nil, errors.New("username not found in our database"))
	_, err := accountService.Login(ctx, &dto.LoginRequest{
		UserName: "reo@gmail.com",
		Password: "123456789012",
	})

	assert.NotNil(t, err)
}

func TestInsertFailed(t *testing.T) {
	ctx := context.Background()
	accountRepo.On("GetByEmail", ctx, &sql.Tx{}, "reo@gmail.com").Return(nil, errors.New("username not found in our database"))
	_, err := accountService.InsertAccount(ctx, &dto.CreateAccountRequest{
		UserName: "reo@gmail.com",
		Password: "123456789012",
	})

	assert.NotNil(t, err)
}
