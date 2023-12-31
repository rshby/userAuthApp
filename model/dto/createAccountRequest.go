package dto

type CreateAccountRequest struct {
	UserName string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required,min=12"`
}
