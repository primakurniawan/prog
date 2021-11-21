package models

import "prog/db"

type User struct {
	ID       int
	Email    string `json:"email" form:"email"`
	Fullname string `json:"fullname" form:"fullname"`
}

func FindAllUsers() []User {
	users := []User{}
	db.DB.Find(&users)
	return users
}

func CreateUser(data User) User {
	user := data
	db.DB.Create(&user)
	return user
}
