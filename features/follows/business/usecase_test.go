package business

import (
	"errors"
	"os"
	"prog/features/follows"
	"prog/features/follows/mocks"
	"prog/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	followsRepo     mocks.Data
	followsBusiness follows.Business
	followData      follows.Core
	followsData     []follows.Core
	usersData       []users.Core
)

func TestMain(m *testing.M) {
	followsBusiness = NewFollowsBusiness(&followsRepo)

	followData = follows.Core{
		FollowingUserId: 1,
		FollowingUser: users.Core{
			ID:       1,
			Email:    "prima@gmail.com",
			Password: "admin",
			Image:    "prima_image.jpg",
		},
		FollowersUserId: 2,
		FollowersUser: users.Core{
			ID:       2,
			Email:    "dwi@gmail.com",
			Password: "admin",
			Image:    "dwi_image.jpg",
		},
	}

	followsData = []follows.Core{
		{
			FollowingUserId: 3,
			FollowingUser: users.Core{
				ID:       3,
				Email:    "maria@gmail.com",
				Password: "admin",
				Image:    "maria_image.jpg",
			},
			FollowersUserId: 2,
			FollowersUser: users.Core{
				ID:       2,
				Email:    "dwi@gmail.com",
				Password: "admin",
				Image:    "dwi_image.jpg",
			}},
		{
			FollowingUserId: 3,
			FollowingUser: users.Core{
				ID:       3,
				Email:    "maria@gmail.com",
				Password: "admin",
				Image:    "maria_image.jpg",
			},
			FollowersUserId: 2,
			FollowersUser: users.Core{
				ID:       2,
				Email:    "dwi@gmail.com",
				Password: "admin",
				Image:    "dwi_image.jpg",
			},
		},
	}

	usersData = []users.Core{
		{
			ID:       1,
			Email:    "prima@gmail.com",
			Password: "admin",
			Image:    "prima_image.jpg",
		},
		{
			ID:       2,
			Email:    "dwi@gmail.com",
			Password: "admin",
			Image:    "dwi_image.jpg",
		},
	}

	os.Exit(m.Run())
}

func TestFollowUser(t *testing.T) {

	t.Run("create tags success", func(t *testing.T) {
		followsRepo.On("FollowUser", mock.AnythingOfType("follows.Core")).Return(nil).Once()
		err := followsBusiness.FollowUser(followData)
		assert.Nil(t, err)
	})

	t.Run("follow user fail", func(t *testing.T) {
		followsRepo.On("FollowUser", mock.AnythingOfType("follows.Core")).Return(errors.New("follow user fail")).Once()
		err := followsBusiness.FollowUser(followData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "follow user fail")
	})

}

func TestUnfollowUser(t *testing.T) {

	t.Run("create tags success", func(t *testing.T) {
		followsRepo.On("UnfollowUser", mock.AnythingOfType("follows.Core")).Return(nil).Once()
		err := followsBusiness.UnfollowUser(followData)
		assert.Nil(t, err)
	})

	t.Run("follow user fail", func(t *testing.T) {
		followsRepo.On("UnfollowUser", mock.AnythingOfType("follows.Core")).Return(errors.New("follow user fail")).Once()
		err := followsBusiness.UnfollowUser(followData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "follow user fail")
	})

}

func TestGetFollowingUsers(t *testing.T) {

	t.Run("get following users success", func(t *testing.T) {
		followsRepo.On("GetFollowingUsers", mock.AnythingOfType("int")).Return(usersData, nil).Once()
		resp, err := followsBusiness.GetFollowingUsers(1)
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("get following users fail", func(t *testing.T) {
		followsRepo.On("GetFollowingUsers", mock.AnythingOfType("int")).Return(nil, errors.New("get following users fail")).Once()
		resp, err := followsBusiness.GetFollowingUsers(0)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, err.Error(), "get following users fail")
	})

}

func TestGetFollowersUsers(t *testing.T) {

	t.Run("get followers users success", func(t *testing.T) {
		followsRepo.On("GetFollowersUsers", mock.AnythingOfType("int")).Return(usersData, nil).Once()
		resp, err := followsBusiness.GetFollowersUsers(1)
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("get followers users fail", func(t *testing.T) {
		followsRepo.On("GetFollowersUsers", mock.AnythingOfType("int")).Return(nil, errors.New("get followers users fail")).Once()
		resp, err := followsBusiness.GetFollowersUsers(0)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, err.Error(), "get followers users fail")
	})

}
