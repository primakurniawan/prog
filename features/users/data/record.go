package data

import (
	"prog/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int
	Email    string
	Password string
	Fullname string
	Image    string
}

func toUserRecord(user users.Core) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func toUserCore(user User) users.Core {
	return users.Core{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func toUserCoreList(uList []User) []users.Core {
	convertedUser := []users.Core{}

	for _, user := range uList {
		convertedUser = append(convertedUser, toUserCore(user))
	}

	return convertedUser
}

func toUserRecordList(uList []users.Core) []User {
	convertedUser := []User{}

	for _, user := range uList {
		convertedUser = append(convertedUser, toUserRecord(users.Core{
			ID:       user.ID,
			Email:    user.Email,
			Fullname: user.Fullname,
			Image:    user.Image,
		}))
	}

	return convertedUser
}
