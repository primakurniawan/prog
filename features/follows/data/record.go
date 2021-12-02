package data

import (
	"prog/features/follows"
	"prog/features/users"
	userData "prog/features/users/data"

	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	ID              int
	FollowingUserId int
	FollowingUser   userData.User
	FollowersUserId int
	FollowersUser   userData.User
}

func ToFollowRecord(data follows.Core) Follow {
	return Follow{
		FollowingUserId: data.FollowingUserId,
		FollowersUserId: data.FollowersUserId,
	}
}

func ToUserFollowersCoreList(fList []Follow) []users.Core {
	convertedUser := []users.Core{}

	for _, follow := range fList {
		convertedUser = append(convertedUser, userData.ToUserCore(follow.FollowersUser))
	}

	return convertedUser
}

func ToUserFollowingCoreList(fList []Follow) []users.Core {
	convertedUser := []users.Core{}

	for _, follow := range fList {
		convertedUser = append(convertedUser, userData.ToUserCore(follow.FollowingUser))
	}

	return convertedUser
}
