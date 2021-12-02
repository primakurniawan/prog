package request

import "prog/features/follows"

type FollowRequest struct {
	FollowingUserId int `param:"userId"`
	FollowersUserId int
}

func (data *FollowRequest) ToFollowCore() follows.Core {
	return follows.Core{
		FollowingUserId: data.FollowingUserId,
		FollowersUserId: data.FollowersUserId,
	}
}
