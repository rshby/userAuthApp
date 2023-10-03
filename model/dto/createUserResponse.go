package dto

import "userAuthApp/model/entity"

type CreateUserResponse struct {
	Id       int              `json:"id,omitempty"`
	FullName string           `json:"full_name,omitempty"`
	Address  string           `json:"address,omitempty"`
	Account  *entity.Accounts `json:"account,omitempty"`
}
