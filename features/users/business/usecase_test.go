package business

import (
	"errors"
	"os"
	"testing"

	"prog/features/users"
	"prog/features/users/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepo     mocks.Data
	userBusiness users.Business
	usersData    []users.Core
	userData     users.Core
	userUpdate   users.Core
)

func TestMain(m *testing.M) {
	userBusiness = NewUserBusiness(&userRepo)

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

	userData = users.Core{
		ID:       3,
		Email:    "maria@gmail.com",
		Password: "admin",
		Image:    "maria_image.jpg",
	}

	userUpdate = users.Core{
		ID:       4,
		Email:    "chris@gmail.com",
		Password: "admin",
		Image:    "chris_image.jpg"}
	os.Exit(m.Run())
}

func TestGetAllUsers(t *testing.T) {
	t.Run("get all users", func(t *testing.T) {
		userRepo.On("GetAllUsers", mock.AnythingOfType("users.Core")).Return(usersData, nil).Once()
		resp, err := userBusiness.GetAllUsers()
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("error get all users", func(t *testing.T) {
		userRepo.On("GetAllUsers", mock.AnythingOfType("users.Core")).Return(nil, errors.New("error on db"))
		resp, err := userBusiness.GetAllUsers()
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func TestRegisterUser(t *testing.T) {
	t.Run("Register user success", func(t *testing.T) {
		userRepo.On("CreateUser", mock.AnythingOfType("users.Core")).Return(1, nil).Once()
		userId, err := userBusiness.RegisterUser(userData)
		assert.Nil(t, err)
		assert.Equal(t, userId, 1)
	})

}

func TestGetUserByID(t *testing.T) {
	t.Run("Get user by id success", func(t *testing.T) {
		userRepo.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		data, err := userBusiness.GetUserById(1)
		assert.Equal(t, userData.ID, data.ID)
		assert.Nil(t, err)
	})

	t.Run("Get user by id error", func(t *testing.T) {
		userRepo.On("GetUserById", mock.AnythingOfType("int")).Return(users.Core{}, errors.New("error get user")).Once()
		data, err := userBusiness.GetUserById(1)
		assert.Empty(t, data)
		assert.NotNil(t, err)
		assert.Equal(t, "error get user", err.Error())
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Update user success", func(t *testing.T) {
		userRepo.On("UpdateUserById", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userBusiness.UpdateUserById(3, userUpdate)
		assert.Nil(t, err)
	})

	t.Run("Update user error UpdateUserData", func(t *testing.T) {
		userRepo.On("UpdateUserById", mock.AnythingOfType("users.Core")).Return(errors.New("error update user")).Once()
		err := userBusiness.UpdateUserById(4, userUpdate)
		assert.NotNil(t, err)
		assert.Equal(t, "error update user", err.Error())
	})

}
