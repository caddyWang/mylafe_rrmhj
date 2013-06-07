package models

type UserBase struct {
	UserId   int
	UserName string
	HeadImg  string
}

type User struct {
	UserBase
}
