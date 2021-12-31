package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type userInfoDto struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
