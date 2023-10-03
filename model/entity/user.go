package entity

type User struct {
	Id        int    `json:"id,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	Address   string `json:"address,omitempty"`
	AccountId int    `json:"account_id,omitempty"`
}
