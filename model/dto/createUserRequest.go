package dto

type CreateUserRequest struct {
	FullName  string `json:"full_name" binding:"required"`
	Address   string `json:"address" binding:"required"`
	AccountId int    `json:"account_id" binding:"required"`
}
