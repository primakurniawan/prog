package follows

type Core struct {
	FollowingUserId int
	FollowingUser   UserCore
	FollowersUserId int
	FollowersUser   UserCore
}

type UserCore struct {
	ID       int
	Email    string
	Fullname string
	Image    string
}

type Business interface {
	FollowUser(followingUserId, followersUserId int) error
	GetFollowingUsers(followersUserId int) ([]UserCore, error)
	GetFollowersUsers(followingUserId int) ([]UserCore, error)
	UnfollowUser(followingUserId, followersUserId int) error
}

type Data interface {
	FollowUser(followingUserId, followersUserId int) error
	GetFollowingUsers(followersUserId int) ([]UserCore, error)
	GetFollowersUsers(followingUserId int) ([]UserCore, error)
	UnfollowUser(followingUserId, followersUserId int) error
}
