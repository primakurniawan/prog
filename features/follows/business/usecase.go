package business

import (
	"prog/features/follows"
)

type FollowsUsecase struct {
	FollowData follows.Data
}

func NewFollowsBusiness(followsData follows.Data) follows.Business {
	return &FollowsUsecase{FollowData: followsData}
}

func (alu *FollowsUsecase) FollowUser(followingUserId, followersUserId int) error {
	err := alu.FollowData.FollowUser(followingUserId, followersUserId)
	if err != nil {
		return err
	}
	return nil
}

func (alu *FollowsUsecase) UnfollowUser(followingUserId, followersUserId int) error {
	err := alu.FollowData.UnfollowUser(followingUserId, followersUserId)
	if err != nil {
		return err
	}
	return nil
}

func (alu *FollowsUsecase) GetFollowingUsers(followersUserId int) ([]follows.UserCore, error) {
	data, err := alu.FollowData.GetFollowingUsers(followersUserId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (alu *FollowsUsecase) GetFollowersUsers(followingUserId int) ([]follows.UserCore, error) {
	data, err := alu.FollowData.GetFollowersUsers(followingUserId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
