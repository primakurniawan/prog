package data

import (
	"prog/features/follows"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlFollowRepository struct {
	Conn *gorm.DB
}

func NewMysqlFollowRepository(conn *gorm.DB) follows.Data {
	return &mysqlFollowRepository{
		Conn: conn,
	}
}

func (alr *mysqlFollowRepository) FollowUser(followingUserId, followersUserId int) error {
	follows := Follow{
		FollowingUserId: followingUserId,
		FollowersUserId: followersUserId,
	}

	err := alr.Conn.Create(&follows).Error
	if err != nil {
		return err
	}
	return nil

}

func (alr *mysqlFollowRepository) UnfollowUser(followingUserId, followersUserId int) error {

	err := alr.Conn.Where("followers_user_id = ? AND following_user_id = ?", followersUserId, followingUserId).Delete(&Follow{}).Error
	if err != nil {
		return err
	}
	return nil

}

func (alr *mysqlFollowRepository) GetFollowingUsers(followersUserId int) ([]follows.UserCore, error) {

	var followingUsers []Follow
	err := alr.Conn.Preload(clause.Associations).Joins("FollowersUser").Where("followers_user_id = ?", followersUserId).Find(&followingUsers).Error
	if err != nil {
		return []follows.UserCore{}, err
	}
	return toUserFollowingCoreList(followingUsers), nil
}

func (alr *mysqlFollowRepository) GetFollowersUsers(followingUserId int) ([]follows.UserCore, error) {

	var followersUsers []Follow
	err := alr.Conn.Preload(clause.Associations).Joins("FollowingUser").Where("followers_user_id = ?", followingUserId).Find(&followersUsers).Error
	if err != nil {
		return []follows.UserCore{}, err
	}
	return toUserFollowersCoreList(followersUsers), nil
}
