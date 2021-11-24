package data

import (
	"prog/features/auth"
	"prog/features/users"
	"time"

	"gorm.io/gorm"
)

type Authentication struct {
	gorm.Model
	Token     string
	CreatedAt time.Time
}

type User struct {
	gorm.Model
	ID       int
	Password string
}

func toAuthRecord(auth auth.Core) Authentication {
	return Authentication{
		Token: auth.Token,
	}
}

func toUserRecord(user users.Core) User {
	return User{
		ID:       user.ID,
		Password: user.Password,
	}
}

func toUserCore(user User) users.Core {
	return users.Core{
		ID: user.ID,
	}
}
