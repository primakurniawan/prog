package follows

import "prog/features/users"

type Core struct {
	FollowingUserId int
	FollowingUser   users.Core
	FollowersUserId int
	FollowersUser   users.Core
}

type Business interface {
	FollowUser(data Core) error
	GetFollowingUsers(followersUserId int) ([]users.Core, error)
	GetFollowersUsers(followingUserId int) ([]users.Core, error)
	UnfollowUser(data Core) error
}

type Data interface {
	FollowUser(data Core) error
	GetFollowingUsers(followersUserId int) ([]users.Core, error)
	GetFollowersUsers(followingUserId int) ([]users.Core, error)
	UnfollowUser(data Core) error
}
