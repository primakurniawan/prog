package business

import (
	"prog/features/follows"
	"prog/features/users"
)

type FollowsUsecase struct {
	FollowData follows.Data
}

func NewFollowsBusiness(followsData follows.Data) follows.Business {
	return &FollowsUsecase{FollowData: followsData}
}

func (alu *FollowsUsecase) FollowUser(data follows.Core) error {
	err := alu.FollowData.FollowUser(data)
	if err != nil {
		return err
	}
	return nil
}

func (alu *FollowsUsecase) UnfollowUser(data follows.Core) error {
	err := alu.FollowData.UnfollowUser(data)
	if err != nil {
		return err
	}
	return nil
}

func (alu *FollowsUsecase) GetFollowingUsers(followersUserId int) ([]users.Core, error) {
	data, err := alu.FollowData.GetFollowingUsers(followersUserId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (alu *FollowsUsecase) GetFollowersUsers(followingUserId int) ([]users.Core, error) {
	data, err := alu.FollowData.GetFollowersUsers(followingUserId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
