package testing

import (
	"context"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"userAuthApp/mock/repo"
	"userAuthApp/model/dto"
	"userAuthApp/service"
)

var userRepo = &repo.InterfaceUserRepository{}
var userService = service.NewUserService(db, userRepo, accountRepo)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	accountRepo.On("GetById", ctx, &sql.Tx{}, 99).Return(nil, errors.New("error"))
	_, err := userService.InsertUser(ctx, &dto.CreateUserRequest{
		FullName:  "Reo",
		Address:   "Jakarta Selatan",
		AccountId: 99,
	})
	assert.NotNil(t, err)
}
