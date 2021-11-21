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

func Login(email, password string) int {
	var user User
	db.DB.Where("email = ? AND password = ?", email, password).First(&user)
	return user.ID
}
