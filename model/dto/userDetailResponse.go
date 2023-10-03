package dto

import "userAuthApp/model/entity"

type UserDetail struct {
	AccountId int
	UserName  string
	UserId    int
	FullName  string
	Address   string
}

type UserDetailResponse struct {
	Account *entity.Accounts `json:"account,omitempty"`
	User    *entity.User     `json:"user,omitempty"`
}
