package entity

type Accounts struct {
	Id       int    `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}
