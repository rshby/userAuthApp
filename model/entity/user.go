package entity

type User struct {
	Id        int    `json:"id"`
	FullName  string `json:"full_name"`
	Address   string `json:"address"`
	AccountId string `json:"account_id"`
}
