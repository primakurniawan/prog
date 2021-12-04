package business

import (
	"errors"
	"os"
	"prog/features/auth"
	"prog/features/auth/mocks"
	"prog/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authRepo     mocks.Data
	authBusiness auth.Business
	authData     auth.Core
	userData     users.Core
)

func TestMain(m *testing.M) {
	authBusiness = NewAuthBusiness(&authRepo)

	authData = auth.Core{
		Token: "12345abcde",
	}

	userData = users.Core{
		ID:       1,
		Email:    "prima@gmail.com",
		Password: "admin",
		Image:    "prima_image.jpg",
	}

	os.Exit(m.Run())
}

func TestAddRefreshToken(t *testing.T) {

	t.Run("add refresh token success", func(t *testing.T) {
		authRepo.On("AddRefreshToken", mock.AnythingOfType("auth.Core")).Return(nil).Once()
		err := authBusiness.AddRefreshToken(authData)
		assert.Nil(t, err)
	})

	t.Run("add refresh token fail", func(t *testing.T) {
		authRepo.On("AddRefreshToken", mock.AnythingOfType("auth.Core")).Return(errors.New("add refresh token fail")).Once()
		err := authBusiness.AddRefreshToken(authData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "add refresh token fail")
	})

}

func TestVerifyRefreshToken(t *testing.T) {

	t.Run("verify refresh token success", func(t *testing.T) {
		authRepo.On("VerifyRefreshToken", mock.AnythingOfType("auth.Core")).Return(nil).Once()
		err := authBusiness.VerifyRefreshToken(authData)
		assert.Nil(t, err)
	})

	t.Run("Verify refresh token fail", func(t *testing.T) {
		authRepo.On("VerifyRefreshToken", mock.AnythingOfType("auth.Core")).Return(errors.New("Verify refresh token fail")).Once()
		err := authBusiness.VerifyRefreshToken(authData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Verify refresh token fail")
	})

}

func TestDeleteRefreshToken(t *testing.T) {

	t.Run("Delete refresh token success", func(t *testing.T) {
		authRepo.On("DeleteRefreshToken", mock.AnythingOfType("auth.Core")).Return(nil).Once()
		err := authBusiness.DeleteRefreshToken(authData)
		assert.Nil(t, err)
	})

	t.Run("Delete refresh token fail", func(t *testing.T) {
		authRepo.On("DeleteRefreshToken", mock.AnythingOfType("auth.Core")).Return(errors.New("Delete refresh token fail")).Once()
		err := authBusiness.DeleteRefreshToken(authData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Delete refresh token fail")
	})

}

func TestVerifyUserCredential(t *testing.T) {

	t.Run("Verify user credentials success", func(t *testing.T) {
		authRepo.On("VerifyUserCredential", mock.AnythingOfType("users.Core")).Return(1, nil).Once()
		userId, err := authBusiness.VerifyUserCredential(userData)
		assert.Nil(t, err)
		assert.Equal(t, userId, 1)
	})

	t.Run("Verify user credentials fail", func(t *testing.T) {
		authRepo.On("VerifyUserCredential", mock.AnythingOfType("users.Core")).Return(0, errors.New("Verify user credentials")).Once()
		userId, err := authBusiness.VerifyUserCredential(userData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Verify user credentials")
		assert.Equal(t, userId, 0)
	})

}
