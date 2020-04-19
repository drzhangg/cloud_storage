package model

type User struct {
	Id       int    `json:"id",db:"id"`
	UserName string `json:"user_name",db:"user_name"`
	UserPwd  string `json:"user_pwd",db:"user_pwd"`
}
