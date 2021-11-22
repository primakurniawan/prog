package data

import (
	"errors"
	"prog/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		Conn: conn,
	}
}

func (ur *mysqlUserRepository) CreateUser(data users.Core) error {
	recordData := toUserRecord(data)
	err := ur.Conn.Create(&recordData)
	if err != nil {
		return err.Error
	}
	return nil
}

func (ur *mysqlUserRepository) GetUsersByFullname(fullname string) ([]users.Core, error) {

	var users []User

	err := ur.Conn.Where("fullname LIKE ?", "%"+fullname+"%").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return toUserCoreList(users), nil

}

func (ur *mysqlUserRepository) GetUserById(userId int) (users.Core, error) {
	var user User
	err := ur.Conn.First(&user, userId).Error

	if user.Fullname == "" && user.ID == 0 {
		return users.Core{}, errors.New("no existing user")
	}
	if err != nil {
		return users.Core{}, err
	}

	return toUserCore(user), nil

}

// func (ur *mysqlUserRepository) GetUserFollowing(userId int) ([]users.Core, error) {

// 	var usersFollowing []UserFollows

// 	err := ur.Conn.Raw("SELECT users.id, users.email, users.fullname, users.image FROM follows LEFT JOIN users ON follows.following_user_id = users.id WHERE follows.followers_user_id = ?", userId).Scan(&usersFollowing).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return toUserFollowsList(usersFollowing), nil

// }

// func (ur *mysqlUserRepository) GetUserFollowers(userId int) ([]users.Core, error) {

// 	var usersFollowers []UserFollows

// 	err := ur.Conn.Raw("SELECT users.id, users.email, users.fullname, users.image FROM follows LEFT JOIN users ON follows.followers_user_id = users.id WHERE follows.following_user_id = ?", userId).Scan(&usersFollowers).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return toUserFollowsList(usersFollowers), nil

// }

// func (ur *mysqlUserRepository) UpdateUserById(data users.Core) error {

// }
// func (ur *mysqlUserRepository) DeleteUserById(data users.Core) error {

// }
