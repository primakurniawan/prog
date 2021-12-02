package data

import (
	"prog/features/auth"
	"prog/features/users"
	userData "prog/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Authentication struct {
	gorm.Model
	Token     string
	CreatedAt time.Time
}

func toAuthRecord(auth auth.Core) Authentication {
	return Authentication{
		Token: auth.Token,
	}
}

func toUserRecord(user users.Core) userData.User {
	return userData.User{
		ID:       user.ID,
		Password: user.Password,
	}
}
