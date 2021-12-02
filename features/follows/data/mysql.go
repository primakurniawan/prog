package data

import (
	"prog/features/follows"
	"prog/features/users"

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

func (alr *mysqlFollowRepository) FollowUser(data follows.Core) error {
	follow := ToFollowRecord(data)

	err := alr.Conn.Create(&follow).Error
	if err != nil {
		return err
	}
	return nil

}

func (alr *mysqlFollowRepository) UnfollowUser(data follows.Core) error {

	err := alr.Conn.Where("followers_user_id = ? AND following_user_id = ?", data.FollowersUserId, data.FollowingUserId).Delete(&Follow{}).Error
	if err != nil {
		return err
	}
	return nil

}

func (alr *mysqlFollowRepository) GetFollowingUsers(followersUserId int) ([]users.Core, error) {

	var followingUsers []Follow
	err := alr.Conn.Preload(clause.Associations).Joins("FollowersUser").Where("followers_user_id = ?", followersUserId).Find(&followingUsers).Error
	if err != nil {
		return []users.Core{}, err
	}
	return ToUserFollowingCoreList(followingUsers), nil
}

func (alr *mysqlFollowRepository) GetFollowersUsers(followingUserId int) ([]users.Core, error) {

	var followersUsers []Follow
	err := alr.Conn.Preload(clause.Associations).Joins("FollowingUser").Where("following_user_id = ?", followingUserId).Find(&followersUsers).Error
	if err != nil {
		return []users.Core{}, err
	}
	return ToUserFollowersCoreList(followersUsers), nil
}
