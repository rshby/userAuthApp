package dto

type LoginResponse struct {
	LoginAt string `json:"login_at"`
	Token   string `json:"token"`
}
