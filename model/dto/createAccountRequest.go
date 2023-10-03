package dto

type CreateAccountRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
