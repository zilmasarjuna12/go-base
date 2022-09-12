package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
}

func (User) TableName() string {
	return "users"
}
