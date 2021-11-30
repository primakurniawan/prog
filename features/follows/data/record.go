package data

import (
	"prog/features/articles"
	"prog/features/follows"

	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	ID              int
	FollowingUserId int
	FollowingUser   User
	FollowersUserId int
	FollowersUser   User
}

type User struct {
	gorm.Model
	ID       int
	Email    string
	Fullname string
	Image    string
}

func toFollowRecord(data follows.Core) Follow {
	return Follow{
		FollowingUserId: data.FollowingUserId,
		FollowersUserId: data.FollowersUserId,
	}
}

func toUserRecord(user articles.UserCore) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func toUserCore(user User) follows.UserCore {
	return follows.UserCore{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func toUserFollowersCoreList(fList []Follow) []follows.UserCore {
	convertedUser := []follows.UserCore{}

	for _, follow := range fList {
		convertedUser = append(convertedUser, toUserCore(follow.FollowersUser))
	}

	return convertedUser
}

func toUserFollowingCoreList(fList []Follow) []follows.UserCore {
	convertedUser := []follows.UserCore{}

	for _, follow := range fList {
		convertedUser = append(convertedUser, toUserCore(follow.FollowingUser))
	}

	return convertedUser
}
