package entity

type Accounts struct {
	Id       int    `json:"id,omitempty"`
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
